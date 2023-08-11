package resources

import (
	"bytes"
	"encoding/json"
	"fmt"
	p "github.com/pulumi/pulumi-go-provider"
	"net/http"
)

type Disk struct{}

type DiskState struct {
	DiskArgs
	Status       string `json:"status" pulumi:"status"`
	DiskSize     int    `json:"disk_size" pulumi:"disk_size"`
	DiskName     string `json:"disk_name" pulumi:"disk_name"`
	Description  string `json:"description" pulumi:"description"`
	Exposed      bool   `json:"exposed" pulumi:"exposed"`
	DiskID       int    `json:"disk_id" pulumi:"disk_id"`
	DiskType     string `json:"disk_type" pulumi:"disk_type"`
	Order        string `json:"order" pulumi:"order"`
	VMID         int    `json:"vm_id" pulumi:"vm_id"`
	Port         int    `json:"port" pulumi:"port"`
	Model        string `json:"model" pulumi:"model"`
	IOTune       IOTune `json:"iotune" pulumi:"iotune"`
	CloudspaceID string `json:"cloudspace_id" pulumi:"cloudspace_id"`
}

type IOTune struct {
	IOPS float64 `json:"iops" pulumi:"iops"`
}

type DiskArgs struct {
	URL             string  `pulumi:"url"`
	Token           string  `pulumi:"token"`
	CustomerID      string  `pulumi:"customerID"`
	Location        string  `pulumi:"location"`
	DiskName        string  `pulumi:"disk_name"`
	DiskDescription string  `pulumi:"description"`
	DiskSize        int     `pulumi:"disk_size"`
	Iops            *int    `pulumi:"iops,optional"`
	DiskType        *string `pulumi:"disk_type,optional"`
	VmID            *string `pulumi:"vm_id,optional"`
	Detach          *bool   `pulumi:"detach,optional"`
	Permanently     *bool   `pulumi:"permanently,optional"`
}

func (d Disk) Create(ctx p.Context, name string, input DiskArgs, preview bool) (string, DiskState, error) {
	state := DiskState{DiskArgs: input}
	url := fmt.Sprintf("https://%s/api/1/customers/%s/locations/%s/disks?disk_name=%s&description=%s&disk_size=%d", input.URL, input.CustomerID, input.Location, input.DiskName, input.DiskDescription, input.DiskSize)

	if preview {
		return name, state, nil
	}

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

	diskID := result["disk_id"].(float64)
	state.DiskID = int(diskID)

	updatedState, err := d.Read(ctx, state, input)
	if err != nil {
		return "", state, err
	}

	return name, updatedState, nil
}

func (Disk) Read(ctx p.Context, state DiskState, input DiskArgs) (DiskState, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/locations/%s/disks/%d", input.URL, input.CustomerID, input.Location, state.DiskID)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return DiskState{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", input.Token))

	resp, err := client.Do(req)
	if err != nil {
		return DiskState{}, err
	}
	defer resp.Body.Close()

	var result DiskState
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return DiskState{}, err
	}

	return result, nil
}

func (Disk) Delete(ctx p.Context, state DiskState, input DiskArgs) (bool, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/locations/%s/disks/%d", input.URL, input.CustomerID, input.Location, state.DiskID)

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
