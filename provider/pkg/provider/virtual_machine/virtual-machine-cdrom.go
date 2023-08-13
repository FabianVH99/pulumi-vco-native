package virtual_machine

import (
	"bytes"
	"encoding/json"
	"fmt"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"net/http"
)

type VirtualMachineCD struct{}

type VirtualMachineCDState struct {
	VirtualMachineCDArgs
	URL              string `pulumi:"url"`
	Token            string `pulumi:"token"`
	CustomerID       string `pulumi:"customerID"`
	CloudSpaceID     string `pulumi:"cloudspace_id"`
	CdRomID          int    `pulumi:"cdrom_id" json:"cdrom_id"`
	VirtualMachineID int    `pulumi:"vm_id" json:"vm_id"`
	Name             string `pulumi:"name" json:"name"`
	Description      string `pulumi:"description" json:"description"`
	Status           string `pulumi:"status" json:"status"`
	DiskSize         string `pulumi:"disk_size" json:"disk_size"`
}

type VirtualMachineCDArgs struct {
	URL              string `pulumi:"url" provider:"secret"`
	Token            string `pulumi:"token" provider:"secret"`
	CustomerID       string `pulumi:"customerID" provider:"secret"`
	CloudSpaceID     string `pulumi:"cloudspace_id"`
	VirtualMachineID int    `pulumi:"vm_id"`
	CdRomID          int    `pulumi:"cdrom_id"`
}

func (cd VirtualMachineCD) WireDependencies(f infer.FieldSelector, args *VirtualMachineCDArgs, state *VirtualMachineCDState) {
	f.OutputField(&state.URL).DependsOn(f.InputField(&args.URL))
	f.OutputField(&state.Token).DependsOn(f.InputField(&args.Token))
	f.OutputField(&state.CustomerID).DependsOn(f.InputField(&args.CustomerID))
	f.OutputField(&state.CloudSpaceID).DependsOn(f.InputField(&args.CloudSpaceID))
	f.OutputField(&state.VirtualMachineID).DependsOn(f.InputField(&args.VirtualMachineID))
	f.OutputField(&state.CdRomID).DependsOn(f.InputField(&args.CdRomID))
}

func (cd VirtualMachineCD) Create(ctx p.Context, name string, input VirtualMachineCDArgs, preview bool) (string, VirtualMachineCDState, error) {
	state := VirtualMachineCDState{VirtualMachineCDArgs: input}
	id, err := resource.NewUniqueHex(name, 8, 0)
	if err != nil {
		return "", state, err
	}
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
	state.URL = input.URL
	state.CustomerID = input.CustomerID
	state.Token = input.Token
	state.CloudSpaceID = input.CloudSpaceID
	state.CdRomID = input.CdRomID
	state.VirtualMachineID = input.VirtualMachineID

	updatedState, err := cd.Read(ctx, id, state)
	if err != nil {
		return "", state, err
	}

	return id, updatedState, nil
}

func (VirtualMachineCD) Read(ctx p.Context, id string, state VirtualMachineCDState) (VirtualMachineCDState, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/vms/%d/cdrom-images/%d", state.URL, state.CustomerID, state.CloudSpaceID, state.VirtualMachineID, state.CdRomID)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return VirtualMachineCDState{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", state.Token))

	resp, err := client.Do(req)
	if err != nil {
		return VirtualMachineCDState{}, err
	}
	defer resp.Body.Close()

	var result VirtualMachineCDState
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return VirtualMachineCDState{}, err
	}

	result.URL = state.URL
	result.CustomerID = state.CustomerID
	result.Token = state.Token
	result.CloudSpaceID = state.CloudSpaceID
	result.VirtualMachineID = state.VirtualMachineID

	return result, nil
}

func (VirtualMachineCD) Update(ctx p.Context, id string, state VirtualMachineCDState, input VirtualMachineCDArgs) (VirtualMachineCDState, error) {
	return VirtualMachineCDState{}, nil
}

func (VirtualMachineCD) Delete(ctx p.Context, id string, state VirtualMachineCDState) error {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/vms/%d/cdrom-images/%d", state.URL, state.CustomerID, state.CloudSpaceID, state.VirtualMachineID, state.CdRomID)
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
