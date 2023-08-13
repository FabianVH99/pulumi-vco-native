package anti_affinity_group

import (
	"bytes"
	"encoding/json"
	"fmt"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"net/http"
)

type AntiAffinityGroupVM struct{}

type AntiAffinityGroupVMState struct {
	AntiAffinityGroupVMArgs
	VirtualMachineID int    `pulumi:"vm_id" json:"vm_id"`
	Status           string `pulumi:"status" json:"status"`
}

type AntiAffinityGroupVMArgs struct {
	URL              string `pulumi:"url" provider:"secret"`
	Token            string `pulumi:"token" provider:"secret"`
	CustomerID       string `pulumi:"customerID" provider:"secret"`
	CloudSpaceID     string `pulumi:"cloudspace_id"`
	GroupID          string `pulumi:"group_id"`
	VirtualMachineID int    `pulumi:"vm_id"`
}

func (agvm AntiAffinityGroupVM) Create(ctx p.Context, name string, input AntiAffinityGroupVMArgs, preview bool) (string, AntiAffinityGroupVMState, error) {
	state := AntiAffinityGroupVMState{AntiAffinityGroupVMArgs: input}
	id, err := resource.NewUniqueHex(name, 8, 0)
	if err != nil {
		return "", state, err
	}
	if preview {
		return name, state, nil
	}

	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/anti-affinity-groups/%s/vms?vm_id=%d", input.URL, input.CustomerID, input.CloudSpaceID, input.GroupID, input.VirtualMachineID)

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

	updatedState, err := agvm.Read(ctx, id, input)
	if err != nil {
		return "", state, err
	}

	return id, updatedState, nil
}

func (AntiAffinityGroupVM) Read(ctx p.Context, id string, input AntiAffinityGroupVMArgs) (AntiAffinityGroupVMState, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/anti-affinity-groups/%s/vms", input.URL, input.CustomerID, input.CloudSpaceID, input.GroupID)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error making API request for %s: %v", id, err)
		return AntiAffinityGroupVMState{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", input.Token))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error retrieving resource %s: %v\n", id, err)
		return AntiAffinityGroupVMState{}, err
	}
	defer resp.Body.Close()

	var result struct {
		VMS []struct {
			VirtualMachineID int    `json:"vm_id"`
			Status           string `json:"status"`
		} `json:"vms"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Printf("Error decoding response body for %s: %v\n", id, err)
		return AntiAffinityGroupVMState{}, err
	}

	var state AntiAffinityGroupVMState
	for _, v := range result.VMS {
		if v.VirtualMachineID == input.VirtualMachineID {
			state.VirtualMachineID = v.VirtualMachineID
			state.Status = v.Status
			break
		}
	}

	return state, nil
}

func (AntiAffinityGroupVM) Update(ctx p.Context, id string, state AntiAffinityGroupVMState, input AntiAffinityGroupVMArgs) (AntiAffinityGroupVMState, error) {
	return AntiAffinityGroupVMState{}, nil
}

func (AntiAffinityGroupVM) Delete(ctx p.Context, id string, state AntiAffinityGroupVMState, input AntiAffinityGroupVMArgs) error {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/anti-affinity-groups/%s/vms/%d", input.URL, input.CustomerID, input.CloudSpaceID, input.GroupID, state.VirtualMachineID)
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
