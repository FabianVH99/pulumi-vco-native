package resources

import (
	"bytes"
	"encoding/json"
	"fmt"
	p "github.com/pulumi/pulumi-go-provider"
	"net/http"
)

type AntiAffinityGroupVM struct{}

type AntiAffinityGroupVMState struct {
	AntiAffinityGroupVMArgs
	VirtualMachineID int    `pulumi:"vm_id" json:"vm_id"`
	Status           string `pulumi:"status" json:"status"`
}

type AntiAffinityGroupVMArgs struct {
	URL              string `pulumi:"url"`
	Token            string `pulumi:"token"`
	CustomerID       string `pulumi:"customerID"`
	CloudSpaceID     string `pulumi:"cloudspace_id"`
	GroupID          string `pulumi:"group_id"`
	VirtualMachineID int    `pulumi:"vm_id"`
}

func (agvm AntiAffinityGroupVM) Create(ctx p.Context, name string, input AntiAffinityGroupVMArgs, preview bool) (string, AntiAffinityGroupVMState, error) {
	state := AntiAffinityGroupVMState{AntiAffinityGroupVMArgs: input}

	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/anti-affinity-groups/%s/vms?vm_id=%d", input.URL, input.CustomerID, input.CloudSpaceID, input.GroupID, input.VirtualMachineID)

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

	updatedState, err := agvm.Read(ctx, input)
	if err != nil {
		return "", state, err
	}

	return name, updatedState, nil
}

func (AntiAffinityGroupVM) Read(ctx p.Context, input AntiAffinityGroupVMArgs) (AntiAffinityGroupVMState, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/anti-affinity-groups/%s/vms", input.URL, input.CustomerID, input.CloudSpaceID, input.GroupID)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return AntiAffinityGroupVMState{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", input.Token))

	resp, err := client.Do(req)
	if err != nil {
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

func (AntiAffinityGroupVM) Delete(ctx p.Context, state AntiAffinityGroupVMState, input AntiAffinityGroupVMArgs) (bool, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/anti-affinity-groups/%s/vms/%d", input.URL, input.CustomerID, input.CloudSpaceID, input.GroupID, state.VirtualMachineID)
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
