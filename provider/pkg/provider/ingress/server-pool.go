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

type ServerPool struct{}

type ServerPoolState struct {
	ServerPoolArgs
	URL          string           `pulumi:"url"`
	Token        string           `pulumi:"token"`
	CustomerID   string           `pulumi:"customerID"`
	CloudSpaceID string           `pulumi:"cloudspace_id"`
	ServerPoolID string           `json:"serverpool_id" pulumi:"serverpool_id"`
	Name         string           `json:"name" pulumi:"name"`
	Description  string           `json:"description" pulumi:"description"`
	Hosts        []ServerPoolHost `json:"hosts" pulumi:"hosts"`
}

type ServerPoolHost struct {
	HostID  string `json:"host_id" pulumi:"host_id"`
	Address string `json:"address" pulumi:"address"`
}

type ServerPoolArgs struct {
	URL          string `pulumi:"url" provider:"secret"`
	Token        string `pulumi:"token" provider:"secret"`
	CustomerID   string `pulumi:"customerID" provider:"secret"`
	CloudSpaceID string `pulumi:"cloudspace_id"`
	Name         string `pulumi:"name"`
	Description  string `pulumi:"description"`
}

func (sv ServerPool) WireDependencies(f infer.FieldSelector, args *ServerPoolArgs, state *ServerPoolState) {
	f.OutputField(&state.URL).DependsOn(f.InputField(&args.URL))
	f.OutputField(&state.Token).DependsOn(f.InputField(&args.Token))
	f.OutputField(&state.CustomerID).DependsOn(f.InputField(&args.CustomerID))
	f.OutputField(&state.CloudSpaceID).DependsOn(f.InputField(&args.CloudSpaceID))
	f.OutputField(&state.Name).DependsOn(f.InputField(&args.Name))
	f.OutputField(&state.Description).DependsOn(f.InputField(&args.Description))
}

func (sv ServerPool) Create(ctx p.Context, name string, input ServerPoolArgs, preview bool) (string, ServerPoolState, error) {
	state := ServerPoolState{ServerPoolArgs: input}
	id, err := resource.NewUniqueHex(name, 8, 0)
	if err != nil {
		return "", state, err
	}
	if preview {
		return name, state, nil
	}

	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/ingress/server-pools?name=%s&description=%s", input.URL, input.CustomerID, input.CloudSpaceID, input.Name, input.Description)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(nil))
	if err != nil {
		return "", state, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", input.Token))

	resp, err := client.Do(req)
	if err != nil {
		return "", state, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error creating resource %s: received status code %d", id, resp.StatusCode)
		return "", state, fmt.Errorf("received status code %d", resp.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", state, err
	}
	state.URL = input.URL
	state.CustomerID = input.CustomerID
	state.Token = input.Token
	state.CloudSpaceID = input.CloudSpaceID
	state.ServerPoolID = result["id"].(string)

	updatedState, err := sv.Read(ctx, id, state)
	if err != nil {
		return "", state, err
	}

	return id, updatedState, nil
}

func (ServerPool) Read(ctx p.Context, id string, state ServerPoolState) (ServerPoolState, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/ingress/server-pools/%s", state.URL, state.CustomerID, state.CloudSpaceID, state.ServerPoolID)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ServerPoolState{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", state.Token))

	resp, err := client.Do(req)
	if err != nil {
		return ServerPoolState{}, err
	}
	defer resp.Body.Close()

	var result ServerPoolState
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return ServerPoolState{}, err
	}

	result.URL = state.URL
	result.CustomerID = state.CustomerID
	result.Token = state.Token
	result.CloudSpaceID = state.CloudSpaceID
	result.ServerPoolID = state.ServerPoolID

	return result, nil
}

func (sv ServerPool) Update(ctx p.Context, id string, state ServerPoolState, input ServerPoolArgs, preview bool) (ServerPoolState, error) {
	if preview {
		return state, nil
	}
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/ingress/server-pools/%s?name=%s&description=%s", state.URL, state.CustomerID, state.CloudSpaceID, state.ServerPoolID, input.Name, input.Description)

	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(nil))
	if err != nil {
		return state, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", input.Token))

	resp, err := client.Do(req)
	if err != nil {
		return state, err
	}
	defer resp.Body.Close()

	updatedState, err := sv.Read(ctx, id, state)
	if err != nil {
		return state, err
	}

	return updatedState, nil
}

func (ServerPool) Delete(ctx p.Context, id string, state ServerPoolState) error {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/ingress/server-pools/%s", state.URL, state.CustomerID, state.CloudSpaceID, state.ServerPoolID)

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
