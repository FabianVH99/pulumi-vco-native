package resources

import (
	"bytes"
	"encoding/json"
	"fmt"
	p "github.com/pulumi/pulumi-go-provider"
	"net/http"
)

type VirtualMachineDisk struct{}

type VirtualMachineDiskState struct {
	VirtualMachineDiskArgs
	VirtualMachineID int  `pulumi:"vm_id" json:"vm_id"`
	DiskID           int  `pulumi:"disk_id" json:"disk_id"`
	Success          bool `pulumi:"success" json:"success"`
}

type VirtualMachineDiskArgs struct {
	URL              string `pulumi:"url"`
	Token            string `pulumi:"token"`
	CustomerID       string `pulumi:"customerID"`
	CloudSpaceID     string `pulumi:"cloudspace_id"`
	VirtualMachineID int    `pulumi:"vm_id"`
	DiskID           int    `pulumi:"disk_id"`
}

func (VirtualMachineDisk) Create(ctx p.Context, name string, input VirtualMachineDiskArgs, preview bool) (string, VirtualMachineDiskState, error) {
	state := VirtualMachineDiskState{VirtualMachineDiskArgs: input}

	if preview {
		return name, state, nil
	}

	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/vms/%d/disks?disk_id=%d", input.URL, input.CustomerID, input.CloudSpaceID, input.VirtualMachineID, input.DiskID)

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

	success := result["success"].(bool)

	state.DiskID = input.DiskID
	state.VirtualMachineID = input.VirtualMachineID
	state.Success = success

	return name, state, nil
}

func (VirtualMachineDisk) Delete(ctx p.Context, state VirtualMachineDiskState, input VirtualMachineDiskArgs) (bool, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/vms/%d/disks/%d", input.URL, input.CustomerID, input.CloudSpaceID, input.VirtualMachineID, state.DiskID)

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
