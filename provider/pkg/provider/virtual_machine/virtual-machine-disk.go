package virtual_machine

import (
	"bytes"
	"fmt"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"net/http"
)

type VirtualMachineDisk struct{}

type VirtualMachineDiskState struct {
	VirtualMachineDiskArgs
	VirtualMachineID int `pulumi:"vm_id" json:"vm_id"`
	DiskID           int `pulumi:"disk_id" json:"disk_id"`
}

type VirtualMachineDiskArgs struct {
	URL              string `pulumi:"url" provider:"secret"`
	Token            string `pulumi:"token" provider:"secret"`
	CustomerID       string `pulumi:"customerID" provider:"secret"`
	CloudSpaceID     string `pulumi:"cloudspace_id"`
	VirtualMachineID int    `pulumi:"vm_id"`
	DiskID           int    `pulumi:"disk_id"`
}

func (vmdsk VirtualMachineDisk) WireDependencies(f infer.FieldSelector, args *VirtualMachineDiskArgs, state *VirtualMachineDiskState) {
	f.OutputField(&state.VirtualMachineID).DependsOn(f.InputField(&args.VirtualMachineID))
	f.OutputField(&state.DiskID).DependsOn(f.InputField(&args.DiskID))
}

func (VirtualMachineDisk) Create(ctx p.Context, name string, input VirtualMachineDiskArgs, preview bool) (string, VirtualMachineDiskState, error) {
	state := VirtualMachineDiskState{VirtualMachineDiskArgs: input}
	id, err := resource.NewUniqueHex(name, 8, 0)
	if err != nil {
		return "", state, err
	}
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

	state.DiskID = input.DiskID
	state.VirtualMachineID = input.VirtualMachineID

	return id, state, nil
}

func (VirtualMachineDisk) Delete(ctx p.Context, id string, state VirtualMachineDiskState, input VirtualMachineDiskArgs) error {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/vms/%d/disks/%d", input.URL, input.CustomerID, input.CloudSpaceID, input.VirtualMachineID, state.DiskID)

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
