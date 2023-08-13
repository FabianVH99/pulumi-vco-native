package ingress

import (
	"bytes"
	"encoding/json"
	"fmt"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"io/ioutil"
	"net/http"
)

type Host struct{}

type HostState struct {
	HostArgs
	URL          string `pulumi:"url"`
	Token        string `pulumi:"token"`
	CustomerID   string `pulumi:"customerID"`
	CloudSpaceID string `pulumi:"cloudspace_id"`
	HostID       string `json:"host_id" pulumi:"host_id"`
	Address      string `json:"address" pulumi:"address"`
}

type HostArgs struct {
	URL          string `pulumi:"url" provider:"secret"`
	Token        string `pulumi:"token" provider:"secret"`
	CustomerID   string `pulumi:"customerID" provider:"secret"`
	CloudSpaceID string `pulumi:"cloudspace_id"`
	ServerPoolID string `pulumi:"serverpool_id"`
	Address      string `pulumi:"address"`
}

func (exd Host) WireDependencies(f infer.FieldSelector, args *HostArgs, state *HostState) {
	f.OutputField(&state.URL).DependsOn(f.InputField(&args.URL))
	f.OutputField(&state.Token).DependsOn(f.InputField(&args.Token))
	f.OutputField(&state.CustomerID).DependsOn(f.InputField(&args.CustomerID))
	f.OutputField(&state.CloudSpaceID).DependsOn(f.InputField(&args.CloudSpaceID))
	f.OutputField(&state.Address).DependsOn(f.InputField(&args.Address))
}

func (Host) Create(ctx p.Context, name string, input HostArgs, preview bool) (string, HostState, error) {
	state := HostState{HostArgs: input}
	id, err := resource.NewUniqueHex(name, 8, 0)
	if err != nil {
		return "", state, err
	}
	if preview {
		return name, state, nil
	}

	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/ingress/server-pools/%s/hosts?address=%s", input.URL, input.CustomerID, input.CloudSpaceID, input.ServerPoolID, input.Address)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(nil))
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
	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response body for %s: %v\n", id, err)
			return "", state, err
		}
		fmt.Printf("Error creating resource %s: received status code %d\n: %s\n", id, resp.StatusCode, string(body))
		return "", state, fmt.Errorf("received status code %d\n: %s\n", resp.StatusCode, string(body))
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", state, err
	}
	state.URL = input.URL
	state.CustomerID = input.CustomerID
	state.Token = input.Token
	state.CloudSpaceID = input.CloudSpaceID
	state.ServerPoolID = input.ServerPoolID
	state.HostID = result["host_id"].(string)
	state.Address = result["address"].(string)

	return id, state, nil
}

func (Host) Update(ctx p.Context, id string, state HostState, input HostArgs, preview bool) (HostState, error) {
	if preview {
		return state, nil
	}
	return state, nil
}

func (Host) Delete(ctx p.Context, id string, state HostState) error {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/ingress/server-pools/%s/hosts/%s", state.URL, state.CustomerID, state.CloudSpaceID, state.ServerPoolID, state.HostID)

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
	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response body for %s: %v\n", id, err)
			return err
		}
		fmt.Printf("Error creating resource %s: received status code %d\n: %s\n", id, resp.StatusCode, string(body))
		return fmt.Errorf("received status code %d\n: %s\n", resp.StatusCode, string(body))
	}

	return nil
}
