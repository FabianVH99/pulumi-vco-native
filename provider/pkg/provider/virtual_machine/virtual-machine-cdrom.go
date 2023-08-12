package virtual_machine

import (
	"bytes"
	"encoding/json"
	"fmt"
	p "github.com/pulumi/pulumi-go-provider"
	"net/http"
)

type VirtualMachineCD struct{}

type VirtualMachineCDState struct {
	VirtualMachineCDArgs
	CdRomID          int    `pulumi:"cdrom_id" json:"cdrom_id"`
	VirtualMachineID int    `pulumi:"vm_id" json:"vm_id"`
	Name             string `pulumi:"name" json:"name"`
	Description      string `pulumi:"description" json:"description"`
	Status           string `pulumi:"status" json:"status"`
	DiskSize         string `pulumi:"disk_size" json:"disk_size"`
}

type VirtualMachineCDArgs struct {
	URL              string `pulumi:"url"`
	Token            string `pulumi:"token"`
	CustomerID       string `pulumi:"customerID"`
	CloudSpaceID     string `pulumi:"cloudspace_id"`
	VirtualMachineID int    `pulumi:"vm_id"`
	CdRomID          int    `pulumi:"cdrom_id"`
}

func (cd VirtualMachineCD) Create(ctx p.Context, name string, input VirtualMachineCDArgs, preview bool) (string, VirtualMachineCDState, error) {
	state := VirtualMachineCDState{VirtualMachineCDArgs: input}

	if preview {
		return name, state, nil
	}

	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/vms/%d/cdrom-images?cdrom_id=%d", input.URL, input.CustomerID, input.CloudSpaceID, input.VirtualMachineID, input.CdRomID)

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

	state.CdRomID = input.CdRomID
	state.VirtualMachineID = input.VirtualMachineID

	updatedState, err := cd.Read(ctx, state, input)
	if err != nil {
		return "", state, err
	}

	return name, updatedState, nil
}

func (VirtualMachineCD) Read(ctx p.Context, state VirtualMachineCDState, input VirtualMachineCDArgs) (VirtualMachineCDState, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/vms/%d/cdrom-images/%d", input.URL, input.CustomerID, input.CloudSpaceID, input.VirtualMachineID, state.CdRomID)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return VirtualMachineCDState{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", input.Token))

	resp, err := client.Do(req)
	if err != nil {
		return VirtualMachineCDState{}, err
	}
	defer resp.Body.Close()

	var result VirtualMachineCDState
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return VirtualMachineCDState{}, err
	}

	result.VirtualMachineCDArgs = input
	return result, nil
}

func (VirtualMachineCD) Delete(ctx p.Context, state VirtualMachineCDState, input VirtualMachineCDArgs) (bool, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/vms/%d/cdrom-images/%d", input.URL, input.CustomerID, input.CloudSpaceID, input.VirtualMachineID, state.CdRomID)
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
