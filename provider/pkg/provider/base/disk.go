package base

import (
	"bytes"
	"encoding/json"
	"fmt"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"net/http"
)

type Disk struct{}

type DiskState struct {
	DiskArgs
	Status          string `pulumi:"status" json:"status"`
	DiskSize        int    `pulumi:"disk_size" json:"disk_size"`
	DiskName        string `pulumi:"disk_name" json:"disk_name"`
	DiskDescription string `pulumi:"description" json:"description"`
	Exposed         bool   `pulumi:"exposed" json:"exposed"`
	DiskID          int    `pulumi:"disk_id" json:"disk_id"`
	DiskType        string `pulumi:"disk_type" json:"disk_type"`
	Order           string `pulumi:"order" json:"order"`
	VMID            int    `pulumi:"vm_id" json:"vm_id"`
	Port            int    `pulumi:"port" json:"port"`
	Model           string `pulumi:"model" json:"model"`
	CloudspaceID    string `pulumi:"cloudspace_id" json:"cloudspace_id"`
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

func (d Disk) WireDependencies(f infer.FieldSelector, args *DiskArgs, state *DiskState) {
	f.OutputField(&state.DiskName).DependsOn(f.InputField(&args.DiskName))
	f.OutputField(&state.DiskDescription).DependsOn(f.InputField(&args.DiskDescription))
	f.OutputField(&state.DiskSize).DependsOn(f.InputField(&args.DiskSize))
}

func (d Disk) Create(ctx p.Context, name string, input DiskArgs, preview bool) (string, DiskState, error) {
	state := DiskState{DiskArgs: input}
	id, err := resource.NewUniqueHex(name, 8, 0)
	if err != nil {
		return "", state, err
	}
	if preview {
		return name, state, nil
	}
	url := fmt.Sprintf("https://%s/api/1/customers/%s/locations/%s/disks?disk_name=%s&description=%s&disk_size=%d", input.URL, input.CustomerID, input.Location, input.DiskName, input.DiskDescription, input.DiskSize)

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

	diskID := result["disk_id"].(float64)
	state.DiskID = int(diskID)

	updatedState, err := d.Read(ctx, id, state, input)
	if err != nil {
		return "", state, err
	}

	return id, updatedState, nil
}

func (Disk) Read(ctx p.Context, id string, state DiskState, input DiskArgs) (DiskState, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/locations/%s/disks/%d", input.URL, input.CustomerID, input.Location, state.DiskID)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error making API request for %s: %v", id, err)
		return DiskState{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", input.Token))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error retrieving resource %s: %v\n", id, err)
		return DiskState{}, err
	}
	defer resp.Body.Close()

	var result DiskState
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Printf("Error decoding response body for %s: %v\n", id, err)
		return DiskState{}, err
	}

	return result, nil
}

func (Disk) Update(ctx p.Context, id string, state DiskState, input DiskArgs) (DiskState, error) {
	return DiskState{}, nil
}

func (Disk) Delete(ctx p.Context, id string, state DiskState, input DiskArgs) error {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/locations/%s/disks/%d", input.URL, input.CustomerID, input.Location, state.DiskID)

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
