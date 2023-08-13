package ingress

import (
	"bytes"
	"encoding/json"
	"fmt"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"net/http"
)

type LoadBalancer struct{}

type LoadBalancerState struct {
	LoadBalancerArgs
	URL            string       `pulumi:"url"`
	Token          string       `pulumi:"token"`
	CustomerID     string       `pulumi:"customerID"`
	LoadBalancerID string       `pulumi:"loadbalancer_id" json:"loadbalancer_id"`
	CloudSpaceID   string       `pulumi:"cloudspace_id"`
	Name           string       `pulumi:"name" json:"name"`
	Description    string       `pulumi:"description" json:"description"`
	Type           string       `pulumi:"type" json:"type"`
	FrontEnd       FrontEnd     `pulumi:"front_end" json:"front_end"`
	BackEnd        BackEndState `pulumi:"back_end" json:"back_end"`
}

type BackEndState struct {
	ServerpoolID   string `pulumi:"serverpool_id" json:"serverpool_id"`
	ServerpoolName string `pulumi:"serverpool_name" json:"serverpool_name"`
	TargetPort     int    `pulumi:"target_port" json:"target_port"`
}

type LoadBalancerArgs struct {
	URL          string  `pulumi:"url" provider:"secret"`
	Token        string  `pulumi:"token" provider:"secret"`
	CustomerID   string  `pulumi:"customerID" provider:"secret"`
	CloudSpaceID string  `pulumi:"cloudspace_id"`
	Name         string  `pulumi:"name"`
	Description  *string `pulumi:"description,optional"`
	Type         string  `pulumi:"type"`
	FrontEnd     FrontEnd
	BackEnd      BackEnd
}

type FrontEnd struct {
	Port      int     `pulumi:"port" json:"port"`
	IPAddress *string `pulumi:"ip_address,optional" json:"ip_address"`
	TLS       *TLS    `pulumi:"tls,optional" json:"tls"`
}

type TLS struct {
	IsEnabled      *bool   `pulumi:"is_enabled,optional" json:"is_enabled"`
	Domain         *string `pulumi:"domain,optional" json:"domain"`
	TLSTermination *bool   `pulumi:"tls_termination,optional" json:"tls_termination"`
}

type BackEnd struct {
	ServerpoolID string `pulumi:"serverpool_id"`
	TargetPort   int    `pulumi:"target_port"`
}

func (lb LoadBalancer) WireDependencies(f infer.FieldSelector, args *LoadBalancerArgs, state *LoadBalancerState) {
	f.OutputField(&state.URL).DependsOn(f.InputField(&args.URL))
	f.OutputField(&state.Token).DependsOn(f.InputField(&args.Token))
	f.OutputField(&state.CustomerID).DependsOn(f.InputField(&args.CustomerID))
	f.OutputField(&state.CloudSpaceID).DependsOn(f.InputField(&args.CloudSpaceID))
	f.OutputField(&state.Name).DependsOn(f.InputField(&args.Name))
	f.OutputField(&state.Type).DependsOn(f.InputField(&args.Type))
	f.OutputField(&state.FrontEnd).DependsOn(f.InputField(&args.FrontEnd))
}

func (lb LoadBalancer) Create(ctx p.Context, name string, input LoadBalancerArgs, preview bool) (string, LoadBalancerState, error) {
	state := LoadBalancerState{LoadBalancerArgs: input}
	id, err := resource.NewUniqueHex(name, 8, 0)
	if err != nil {
		return "", state, err
	}
	if preview {
		return name, state, nil
	}

	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/ingress/load-balancers", input.URL, input.CustomerID, input.CloudSpaceID)
	payload := map[string]interface{}{
		"name": input.Name,
		"type": input.Type,
		"front_end": map[string]interface{}{
			"port": input.FrontEnd.Port,
		},
		"back_end": map[string]interface{}{
			"serverpool_id": input.BackEnd.ServerpoolID,
			"target_port":   input.BackEnd.TargetPort,
		},
	}

	if input.Description != nil {
		payload["description"] = *input.Description
	}

	if input.FrontEnd.IPAddress != nil {
		payload["front_end"].(map[string]interface{})["ip_address"] = *input.FrontEnd.IPAddress
	}

	if input.FrontEnd.TLS != nil {
		if input.FrontEnd.TLS.IsEnabled != nil || input.FrontEnd.TLS.Domain != nil || input.FrontEnd.TLS.TLSTermination != nil {
			tls := map[string]interface{}{}
			if input.FrontEnd.TLS.IsEnabled != nil {
				tls["is_enabled"] = *input.FrontEnd.TLS.IsEnabled
			}
			if input.FrontEnd.TLS.Domain != nil {
				tls["domain"] = *input.FrontEnd.TLS.Domain
			}
			if input.FrontEnd.TLS.TLSTermination != nil {
				tls["tls_termination"] = *input.FrontEnd.TLS.TLSTermination
			}
			payload["front_end"].(map[string]interface{})["tls"] = tls
		}
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", state, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Printf("Error making API request for %s: %v", id, err)
		return "", state, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", input.Token))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error creating resource %s: %v", id, err)
		return "", state, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", state, err
	}
	state.URL = input.URL
	state.CustomerID = input.CustomerID
	state.Token = input.Token
	state.CloudSpaceID = input.CloudSpaceID
	state.LoadBalancerID = result["id"].(string)

	updatedState, err := lb.Read(ctx, id, state)
	if err != nil {
		return "", state, err
	}

	return id, updatedState, nil
}

func (LoadBalancer) Read(ctx p.Context, id string, state LoadBalancerState) (LoadBalancerState, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/ingress/load-balancers/%s", state.URL, state.CustomerID, state.CloudSpaceID, state.LoadBalancerID)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error making API request for %s: %v", id, err)
		return LoadBalancerState{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", state.Token))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error retrieving resource %s: %v", id, err)
		return LoadBalancerState{}, err
	}
	defer resp.Body.Close()

	var result LoadBalancerState
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return LoadBalancerState{}, err
	}

	result.URL = state.URL
	result.CustomerID = state.CustomerID
	result.Token = state.Token
	result.CloudSpaceID = state.CloudSpaceID
	result.LoadBalancerID = state.LoadBalancerID

	return result, nil
}

func (lb LoadBalancer) Update(ctx p.Context, id string, state LoadBalancerState, input LoadBalancerArgs) (LoadBalancerState, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/ingress/load-balancers/%s", state.URL, state.CustomerID, state.CloudSpaceID, state.LoadBalancerID)
	payload := map[string]interface{}{
		"name": input.Name,
		"type": input.Type,
		"front_end": map[string]interface{}{
			"port": input.FrontEnd.Port,
		},
		"back_end": map[string]interface{}{
			"serverpool_id": input.BackEnd.ServerpoolID,
			"target_port":   input.BackEnd.TargetPort,
		},
	}

	if input.Description != nil {
		payload["description"] = *input.Description
	}

	if input.FrontEnd.IPAddress != nil {
		payload["front_end"].(map[string]interface{})["ip_address"] = *input.FrontEnd.IPAddress
	}

	if input.FrontEnd.TLS != nil {
		if input.FrontEnd.TLS.IsEnabled != nil || input.FrontEnd.TLS.Domain != nil || input.FrontEnd.TLS.TLSTermination != nil {
			tls := map[string]interface{}{}
			if input.FrontEnd.TLS.IsEnabled != nil {
				tls["is_enabled"] = *input.FrontEnd.TLS.IsEnabled
			}
			if input.FrontEnd.TLS.Domain != nil {
				tls["domain"] = *input.FrontEnd.TLS.Domain
			}
			if input.FrontEnd.TLS.TLSTermination != nil {
				tls["tls_termination"] = *input.FrontEnd.TLS.TLSTermination
			}
			payload["front_end"].(map[string]interface{})["tls"] = tls
		}
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return LoadBalancerState{}, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Printf("Error making API request for %s: %v", id, err)
		return state, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", state.Token))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error updating resource %s: %v", id, err)
		return state, err
	}
	defer resp.Body.Close()

	updatedState, err := lb.Read(ctx, id, state)
	if err != nil {
		return state, err
	}

	return updatedState, nil
}

func (LoadBalancer) Delete(ctx p.Context, id string, state LoadBalancerState) error {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/ingress/load-balancers/%s", state.URL, state.CustomerID, state.CloudSpaceID, state.LoadBalancerID)
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(nil))
	if err != nil {
		fmt.Printf("Error making API request for %s: %v", id, err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", state.Token))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error deleting resource %s: %v\n", id, err)
		return err
	}
	defer resp.Body.Close()

	return nil
}
