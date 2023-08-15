package base

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

type Disk struct{}

type DiskState struct {
	DiskArgs
	URL             string `pulumi:"url"`
	Token           string `pulumi:"token"`
	CustomerID      string `pulumi:"customerID"`
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
	VmID            *int    `pulumi:"vm_id,optional"`
	Detach          *bool   `pulumi:"detach,optional"`
	Permanently     *bool   `pulumi:"permanently,optional"`
}

func (d Disk) WireDependencies(f infer.FieldSelector, args *DiskArgs, state *DiskState) {
	f.OutputField(&state.URL).DependsOn(f.InputField(&args.URL))
	f.OutputField(&state.Token).DependsOn(f.InputField(&args.Token))
	f.OutputField(&state.CustomerID).DependsOn(f.InputField(&args.CustomerID))
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
		fmt.Printf("Error decoding response body for %s: %v\n", id, err)
		return "", state, err
	}
	state.URL = input.URL
	state.CustomerID = input.CustomerID
	state.Token = input.Token
	state.Location = input.Location
	diskID := result["disk_id"].(float64)
	state.DiskID = int(diskID)

	updatedState, err := d.Read(ctx, id, state)
	if err != nil {
		return "", state, err
	}

	return id, updatedState, nil
}

func (Disk) Read(ctx p.Context, id string, state DiskState) (DiskState, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/locations/%s/disks/%d", state.URL, state.CustomerID, state.Location, state.DiskID)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error making API request for %s: %v", id, err)
		return DiskState{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", state.Token))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error retrieving resource %s: %v\n", id, err)
		return DiskState{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response body for %s: %v\n", id, err)
			return state, err
		}
		fmt.Printf("Error creating resource %s: received status code %d\n: %s\n", id, resp.StatusCode, string(body))
		return state, fmt.Errorf("received status code %d\n: %s\n", resp.StatusCode, string(body))
	}

	var result DiskState
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Printf("Error decoding response body for %s: %v\n", id, err)
		return DiskState{}, err
	}

	result.URL = state.URL
	result.CustomerID = state.CustomerID
	result.Token = state.Token
	result.Location = state.Location

	return result, nil
}

func (Disk) Update(ctx p.Context, id string, state DiskState, input DiskArgs, preview bool) (DiskState, error) {
	if preview {
		return state, nil
	}
	return state, nil
}

func (Disk) Delete(ctx p.Context, id string, state DiskState) error {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/locations/%s/disks/%d", state.URL, state.CustomerID, state.Location, state.DiskID)

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
