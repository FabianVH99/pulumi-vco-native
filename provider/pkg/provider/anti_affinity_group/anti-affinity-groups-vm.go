package anti_affinity_group

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

type AntiAffinityGroupVM struct{}

type AntiAffinityGroupVMState struct {
	AntiAffinityGroupVMArgs
	URL              string `pulumi:"url"`
	Token            string `pulumi:"token"`
	CustomerID       string `pulumi:"customerID"`
	CloudSpaceID     string `pulumi:"cloudspace_id"`
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

func (agvm AntiAffinityGroupVM) WireDependencies(f infer.FieldSelector, args *AntiAffinityGroupVMArgs, state *AntiAffinityGroupVMState) {
	f.OutputField(&state.URL).DependsOn(f.InputField(&args.URL))
	f.OutputField(&state.Token).DependsOn(f.InputField(&args.Token))
	f.OutputField(&state.CustomerID).DependsOn(f.InputField(&args.CustomerID))
	f.OutputField(&state.CloudSpaceID).DependsOn(f.InputField(&args.CloudSpaceID))
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
	state.CloudSpaceID = input.CloudSpaceID
	state.GroupID = input.GroupID
	state.VirtualMachineID = input.VirtualMachineID

	updatedState, err := agvm.Read(ctx, id, state)
	if err != nil {
		return "", state, err
	}

	return id, updatedState, nil
}

func (AntiAffinityGroupVM) Read(ctx p.Context, id string, state AntiAffinityGroupVMState) (AntiAffinityGroupVMState, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/anti-affinity-groups/%s/vms", state.URL, state.CustomerID, state.CloudSpaceID, state.GroupID)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error making API request for %s: %v", id, err)
		return AntiAffinityGroupVMState{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", state.Token))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error retrieving resource %s: %v\n", id, err)
		return AntiAffinityGroupVMState{}, err
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

	var newState AntiAffinityGroupVMState
	for _, v := range result.VMS {
		if v.VirtualMachineID == state.VirtualMachineID {
			state.VirtualMachineID = v.VirtualMachineID
			state.Status = v.Status
			break
		}
	}

	return newState, nil
}

func (AntiAffinityGroupVM) Update(ctx p.Context, id string, state AntiAffinityGroupVMState, input AntiAffinityGroupVMArgs, preview bool) (AntiAffinityGroupVMState, error) {
	if preview {
		return state, nil
	}
	return state, nil
}

func (AntiAffinityGroupVM) Delete(ctx p.Context, id string, state AntiAffinityGroupVMState) error {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/anti-affinity-groups/%s/vms/%d", state.URL, state.CustomerID, state.CloudSpaceID, state.GroupID, state.VirtualMachineID)
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
