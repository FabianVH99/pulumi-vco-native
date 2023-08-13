package disk

import (
	"bytes"
	"encoding/json"
	"fmt"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"net/http"
)

type ExposedDisk struct{}

type ExposedDiskState struct {
	ExposedDiskArgs
	URL          string   `pulumi:"url"`
	Token        string   `pulumi:"token"`
	CustomerID   string   `pulumi:"customerID"`
	CloudSpaceID string   `pulumi:"cloudspace_id"`
	DiskID       int      `json:"disk_id" pulumi:"disk_id"`
	Protocol     string   `json:"protocol" pulumi:"protocol"`
	Endpoint     Endpoint `json:"endpoint" pulumi:"endpoint"`
}

type Endpoint struct {
	PSK            string `json:"psk" pulumi:"psk"`
	User           string `json:"user" pulumi:"user"`
	Port           int    `json:"port" pulumi:"port"`
	Name           string `json:"name" pulumi:"name"`
	Address        string `json:"address" pulumi:"address"`
	PrivatePort    int    `json:"private_port" pulumi:"private_port"`
	PrivateAddress string `json:"private_address" pulumi:"private_address"`
}

type ExposedDiskArgs struct {
	URL            string `pulumi:"url" provider:"secret"`
	Token          string `pulumi:"token" provider:"secret"`
	CustomerID     string `pulumi:"customerID" provider:"secret"`
	CloudSpaceID   string `pulumi:"cloudspace_id"`
	DiskID         int    `pulumi:"disk_id"`
	IOPS           *int   `pulumi:"iops,optional"`
	MaxConnections *int   `pulumi:"max_connections,optional"`
}

func (exd ExposedDisk) WireDependencies(f infer.FieldSelector, args *ExposedDiskArgs, state *ExposedDiskState) {
	f.OutputField(&state.URL).DependsOn(f.InputField(&args.URL))
	f.OutputField(&state.Token).DependsOn(f.InputField(&args.Token))
	f.OutputField(&state.CustomerID).DependsOn(f.InputField(&args.CustomerID))
	f.OutputField(&state.CloudSpaceID).DependsOn(f.InputField(&args.CloudSpaceID))
	f.OutputField(&state.DiskID).DependsOn(f.InputField(&args.DiskID))
}

func (ExposedDisk) Create(ctx p.Context, name string, input ExposedDiskArgs, preview bool) (string, ExposedDiskState, error) {
	state := ExposedDiskState{ExposedDiskArgs: input}
	id, err := resource.NewUniqueHex(name, 8, 0)
	if err != nil {
		return "", state, err
	}
	if preview {
		return name, state, nil
	}

	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/exposed-disks", input.URL, input.CustomerID, input.CloudSpaceID)
	payload := map[string]interface{}{
		"disk_id": input.DiskID,
	}

	if input.IOPS != nil {
		payload["iops"] = *input.IOPS
	}

	if input.MaxConnections != nil {
		payload["max_connections"] = *input.MaxConnections
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("Error constructing JSON payload %s: %v", id, err)
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
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error creating resource %s: received status code %d", id, resp.StatusCode)
		return "", state, fmt.Errorf("received status code %d", resp.StatusCode)
	}
	state.URL = input.URL
	state.CustomerID = input.CustomerID
	state.Token = input.Token
	state.CloudSpaceID = input.CloudSpaceID

	var result ExposedDiskState
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return name, ExposedDiskState{}, err
	}

	result.DiskID = input.DiskID

	return id, result, nil
}

func (ExposedDisk) Update(ctx p.Context, id string, state ExposedDiskState, input ExposedDiskArgs, preview bool) (ExposedDiskState, error) {
	if preview {
		return state, nil
	}
	return state, nil
}

func (ExposedDisk) Delete(ctx p.Context, id string, state ExposedDiskState) error {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/exposed-disks/%d", state.URL, state.CustomerID, state.CloudSpaceID, state.DiskID)

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
