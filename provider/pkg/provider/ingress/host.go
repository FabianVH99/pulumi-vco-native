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

type Host struct{}

type HostState struct {
	HostArgs
	HostID  string `json:"host_id" pulumi:"host_id"`
	Address string `json:"address" pulumi:"address"`
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

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", state, err
	}

	state.HostID = result["host_id"].(string)
	state.Address = result["address"].(string)

	return id, state, nil
}

func (Host) Update(ctx p.Context, id string, state HostState, input HostArgs) (HostState, error) {
	return HostState{}, nil
}

func (Host) Delete(ctx p.Context, id string, state HostState, input HostArgs) error {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/ingress/server-pools/%s/hosts/%s", input.URL, input.CustomerID, input.CloudSpaceID, input.ServerPoolID, state.HostID)

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

	return nil
}
