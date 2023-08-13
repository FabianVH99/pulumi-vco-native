package cloudspace

import (
	"bytes"
	"encoding/json"
	"fmt"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"net/http"
)

type AntiAffinityGroup struct{}

type AntiAffinityGroupState struct {
	AntiAffinityGroupArgs
	URL          string `json:"url" pulumi:"url"`
	Token        string `json:"token" pulumi:"token"`
	CustomerID   string `json:"customerID" pulumi:"customerID"`
	CloudSpaceID string `json:"cloudspace_id" pulumi:"cloudspace_id"`
	GroupID      string `json:"group_id" pulumi:"group_id"`
	Spread       int    `json:"spread" pulumi:"spread"`
	Status       string `json:"status" pulumi:"status"`
}

type AntiAffinityGroupArgs struct {
	URL          string `pulumi:"url" provider:"secret"`
	Token        string `pulumi:"token" provider:"secret"`
	CustomerID   string `pulumi:"customerID" provider:"secret"`
	CloudSpaceID string `pulumi:"cloudspace_id"`
	GroupID      string `pulumi:"group_id"`
	Spread       int    `pulumi:"spread"`
}

func (ag AntiAffinityGroup) WireDependencies(f infer.FieldSelector, args *AntiAffinityGroupArgs, state *AntiAffinityGroupState) {
	f.OutputField(&state.URL).DependsOn(f.InputField(&args.URL))
	f.OutputField(&state.Token).DependsOn(f.InputField(&args.Token))
	f.OutputField(&state.CustomerID).DependsOn(f.InputField(&args.CustomerID))
	f.OutputField(&state.CloudSpaceID).DependsOn(f.InputField(&args.CloudSpaceID))
	f.OutputField(&state.GroupID).DependsOn(f.InputField(&args.GroupID))
	f.OutputField(&state.Spread).DependsOn(f.InputField(&args.Spread))
}

func (ag AntiAffinityGroup) Create(ctx p.Context, name string, input AntiAffinityGroupArgs, preview bool) (string, AntiAffinityGroupState, error) {
	state := AntiAffinityGroupState{AntiAffinityGroupArgs: input}
	id, err := resource.NewUniqueHex(name, 8, 0)
	if err != nil {
		return "", state, err
	}
	if preview {
		return name, state, nil
	}

	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/anti-affinity-groups?spread=%d&group_id=%s", input.URL, input.CustomerID, input.CloudSpaceID, input.Spread, input.GroupID)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Printf("Error making API request for %s: %v", id, err)
		return "", state, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", input.Token))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error creating resource %s: %v\n", id, err)
		return "", state, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Printf("Error decoding response body for %s: %v\n", id, err)
		return "", state, err
	}

	updatedState, err := ag.Read(ctx, id, input)
	if err != nil {
		return "", state, err
	}

	return id, updatedState, nil
}

func (AntiAffinityGroup) Read(ctx p.Context, id string, input AntiAffinityGroupArgs) (AntiAffinityGroupState, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/anti-affinity-groups/%s", input.URL, input.CustomerID, input.CloudSpaceID, input.GroupID)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error making API request for %s: %v", id, err)
		return AntiAffinityGroupState{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", input.Token))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error retrieving resource %s: %v\n", id, err)
		return AntiAffinityGroupState{}, err
	}
	defer resp.Body.Close()

	var result AntiAffinityGroupState
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Printf("Error decoding response body for %s: %v\n", id, err)
		return AntiAffinityGroupState{}, err
	}

	return result, nil
}

func (ag AntiAffinityGroup) Update(ctx p.Context, id string, state AntiAffinityGroupState, input AntiAffinityGroupArgs, preview bool) (AntiAffinityGroupState, error) {
	if preview {
		return state, nil
	}
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/anti-affinity-groups/%s?spread=%d", input.URL, input.CustomerID, input.CloudSpaceID, input.GroupID, input.Spread)

	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(nil))
	if err != nil {
		fmt.Printf("Error making API request for %s: %v", id, err)
		return state, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", input.Token))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error updating resource %s: %v\n", id, err)
		return state, err
	}
	defer resp.Body.Close()

	updatedState, err := ag.Read(ctx, id, input)
	if err != nil {
		return state, err
	}

	return updatedState, nil
}

func (AntiAffinityGroup) Delete(ctx p.Context, id string, input AntiAffinityGroupArgs) error {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/anti-affinity-groups/%s", input.URL, input.CustomerID, input.CloudSpaceID, input.GroupID)
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(nil))
	if err != nil {
		fmt.Printf("Error making API request for %s: %v", id, err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", input.Token))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error deleting resource %s: %v\n", id, err)
		return err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Printf("Error decoding response body for %s: %v\n", id, err)
		return err
	}

	return nil
}
