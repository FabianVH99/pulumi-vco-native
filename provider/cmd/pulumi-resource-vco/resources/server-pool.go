package resources

import (
	"bytes"
	"encoding/json"
	"fmt"
	p "github.com/pulumi/pulumi-go-provider"
	"net/http"
)

type ServerPool struct{}

type ServerPoolState struct {
	ServerPoolArgs
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
	URL          string `pulumi:"url"`
	Token        string `pulumi:"token"`
	CustomerID   string `pulumi:"customerID"`
	CloudSpaceID string `pulumi:"cloudspace_id"`
	Name         string `pulumi:"name"`
	Description  string `pulumi:"description"`
}

func (sv ServerPool) Create(ctx p.Context, name string, input ServerPoolArgs, preview bool) (string, ServerPoolState, error) {
	state := ServerPoolState{ServerPoolArgs: input}

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

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", state, err
	}

	state.ServerPoolID = result["id"].(string)

	updatedState, err := sv.Read(ctx, state, input)
	if err != nil {
		return "", state, err
	}

	return name, updatedState, nil
}

func (ServerPool) Read(ctx p.Context, state ServerPoolState, input ServerPoolArgs) (ServerPoolState, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/ingress/server-pools/%s", input.URL, input.CustomerID, input.CloudSpaceID, state.ServerPoolID)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ServerPoolState{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", input.Token))

	resp, err := client.Do(req)
	if err != nil {
		return ServerPoolState{}, err
	}
	defer resp.Body.Close()

	var result ServerPoolState
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return ServerPoolState{}, err
	}

	return result, nil
}

func (sv ServerPool) Update(ctx p.Context, state ServerPoolState, input ServerPoolArgs) (bool, ServerPoolState, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/ingress/server-pools/%s?name=%s&description=%s", input.URL, input.CustomerID, input.CloudSpaceID, state.ServerPoolID, input.Name, input.Description)

	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(nil))
	if err != nil {
		return false, state, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", input.Token))

	resp, err := client.Do(req)
	if err != nil {
		return false, state, err
	}
	defer resp.Body.Close()

	var status map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&status); err != nil {
		return false, state, err
	}

	success := status["success"].(bool)

	updatedState, err := sv.Read(ctx, state, input)
	if err != nil {
		return false, state, err
	}

	return success, updatedState, nil
}

func (ServerPool) Delete(ctx p.Context, state ServerPoolState, input ServerPoolArgs) (bool, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/ingress/server-pools/%s", input.URL, input.CustomerID, input.CloudSpaceID, state.ServerPoolID)

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(nil))
	if err != nil {
		return false, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", input.Token))

	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return false, err
	}

	status := response["success"].(bool)

	return status, nil
}
