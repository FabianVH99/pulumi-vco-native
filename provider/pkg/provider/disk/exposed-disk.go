package disk

import (
	"bytes"
	"encoding/json"
	"fmt"
	p "github.com/pulumi/pulumi-go-provider"
	"net/http"
)

type ExposedDisk struct{}

type ExposedDiskState struct {
	ExposedDiskArgs
	DiskID   int      `json:"disk_id" pulumi:"disk_id"`
	Protocol string   `json:"protocol" pulumi:"protocol"`
	Endpoint Endpoint `json:"endpoint" pulumi:"endpoint"`
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
	URL            string `pulumi:"url"`
	Token          string `pulumi:"token"`
	CustomerID     string `pulumi:"customerID"`
	CloudSpaceID   string `pulumi:"cloudspace_id"`
	DiskID         int    `pulumi:"disk_id"`
	IOPS           *int   `pulumi:"iops,optional"`
	MaxConnections *int   `pulumi:"max_connections,optional"`
}

func (ExposedDisk) Create(ctx p.Context, name string, input ExposedDiskArgs, preview bool) (string, ExposedDiskState, error) {
	state := ExposedDiskState{ExposedDiskArgs: input}

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
		return "", state, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
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

	var result ExposedDiskState
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return name, ExposedDiskState{}, err
	}

	result.DiskID = input.DiskID

	return name, result, nil
}

func (ExposedDisk) Delete(ctx p.Context, state ExposedDiskState, input ExposedDiskArgs) (bool, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/exposed-disks/%d", input.URL, input.CustomerID, input.CloudSpaceID, state.DiskID)

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
