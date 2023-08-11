package resources

import (
	"bytes"
	"encoding/json"
	"fmt"
	p "github.com/pulumi/pulumi-go-provider"
	"net/http"
)

type Host struct{}

type HostState struct {
	HostArgs
	HostID  string `json:"host_id" pulumi:"host_id"`
	Address string `json:"address" pulumi:"address"`
}

type HostArgs struct {
	URL          string `pulumi:"url"`
	Token        string `pulumi:"token"`
	CustomerID   string `pulumi:"customerID"`
	CloudSpaceID string `pulumi:"cloudspace_id"`
	ServerPoolID string `pulumi:"serverpool_id"`
	Address      string `pulumi:"address"`
}

func (Host) Create(ctx p.Context, name string, input HostArgs, preview bool) (string, HostState, error) {
	state := HostState{HostArgs: input}

	if preview {
		return name, state, nil
	}

	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/ingress/server-pools/%s/hosts?address=%s", input.URL, input.CustomerID, input.CloudSpaceID, input.ServerPoolID, input.Address)

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

	state.HostID = result["host_id"].(string)
	state.Address = result["address"].(string)

	return name, state, nil
}

func (Host) Delete(ctx p.Context, state HostState, input HostArgs) (bool, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/ingress/server-pools/%s/hosts/%s", input.URL, input.CustomerID, input.CloudSpaceID, input.ServerPoolID, state.HostID)

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

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	status := result["success"].(bool)

	return status, nil
}
