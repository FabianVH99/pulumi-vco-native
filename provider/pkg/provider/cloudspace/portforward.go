package cloudspace

import (
	"bytes"
	"encoding/json"
	"fmt"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type PortForward struct{}

type PortForwardState struct {
	PortForwardArgs
	URL              string  `pulumi:"url"`
	Token            string  `pulumi:"token"`
	CustomerID       string  `pulumi:"customerID"`
	CloudSpaceID     string  `pulumi:"cloudspace_id"`
	PortForwardID    string  `json:"portforward_id" pulumi:"portforward_id"`
	LocalPort        int     `json:"local_port" pulumi:"local_port"`
	PublicPort       int     `json:"public_port" pulumi:"public_port"`
	Protocol         string  `json:"protocol" pulumi:"protocol"`
	VirtualMachineID int     `json:"vm_id" pulumi:"vm_id"`
	PublicIP         string  `json:"public_ip" pulumi:"public_ip"`
	TillPublicPort   *int    `json:"till_public_port" pulumi:"till_public_port"`
	NestedCSID       *string `json:"nested_cs_id" pulumi:"nested_cs_id"`
}

type PortForwardArgs struct {
	URL              string  `pulumi:"url" provider:"secret"`
	Token            string  `pulumi:"token" provider:"secret"`
	CustomerID       string  `pulumi:"customerID" provider:"secret"`
	CloudSpaceID     string  `pulumi:"cloudspace_id"`
	LocalPort        int     `pulumi:"local_port"`
	PublicPort       int     `pulumi:"public_port"`
	Protocol         string  `pulumi:"protocol"`
	VirtualMachineID int     `pulumi:"vm_id"`
	PublicIP         string  `pulumi:"public_ip"`
	TillPublicPort   *int    `pulumi:"till_public_port,optional"`
	NestedCSID       *string `pulumi:"nested_cs_id,optional"`
}

func (pf PortForward) WireDependencies(f infer.FieldSelector, args *PortForwardArgs, state *PortForwardState) {
	f.OutputField(&state.URL).DependsOn(f.InputField(&args.URL))
	f.OutputField(&state.Token).DependsOn(f.InputField(&args.Token))
	f.OutputField(&state.CustomerID).DependsOn(f.InputField(&args.CustomerID))
	f.OutputField(&state.CloudSpaceID).DependsOn(f.InputField(&args.CloudSpaceID))
	f.OutputField(&state.LocalPort).DependsOn(f.InputField(&args.LocalPort))
	f.OutputField(&state.PublicPort).DependsOn(f.InputField(&args.PublicPort))
	f.OutputField(&state.Protocol).DependsOn(f.InputField(&args.Protocol))
	f.OutputField(&state.VirtualMachineID).DependsOn(f.InputField(&args.VirtualMachineID))
	f.OutputField(&state.PublicIP).DependsOn(f.InputField(&args.PublicIP))
}

func (pf PortForward) Create(ctx p.Context, name string, input PortForwardArgs, preview bool) (string, PortForwardState, error) {
	state := PortForwardState{PortForwardArgs: input}
	id, err := resource.NewUniqueHex(name, 8, 0)
	if err != nil {
		return "", state, err
	}
	if preview {
		return name, state, nil
	}

	u, err := url.Parse(fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/portforwards", input.URL, input.CustomerID, input.CloudSpaceID))
	if err != nil {
		fmt.Printf("Error making API request for %s: %v", id, err)
		return "", state, err
	}

	q := u.Query()
	q.Set("local_port", strconv.Itoa(input.LocalPort))
	q.Set("public_port", strconv.Itoa(input.PublicPort))
	q.Set("protocol", input.Protocol)
	q.Set("vm_id", strconv.Itoa(input.VirtualMachineID))
	q.Set("public_ip", input.PublicIP)

	if input.TillPublicPort != nil {
		q.Set("till_public_port", strconv.Itoa(*input.TillPublicPort))
	}

	if input.NestedCSID != nil {
		q.Set("nested_cs_id", *input.NestedCSID)
	}

	u.RawQuery = q.Encode()
	url := u.String()

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(nil))
	if err != nil {
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
		return "", state, err
	}
	state.URL = input.URL
	state.CustomerID = input.CustomerID
	state.Token = input.Token
	state.CloudSpaceID = input.CloudSpaceID
	state.PortForwardID = result["portforward_id"].(string)

	updatedState, err := pf.Read(ctx, id, state)
	if err != nil {
		return "", state, err
	}

	return id, updatedState, nil
}

func (PortForward) Read(ctx p.Context, id string, state PortForwardState) (PortForwardState, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/portforwards/%s", state.URL, state.CustomerID, state.CloudSpaceID, state.PortForwardID)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error making API request for %s: %v", id, err)
		return PortForwardState{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", state.Token))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error updating resource %s: %v\n", id, err)
		return PortForwardState{}, err
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

	var result PortForwardState
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return PortForwardState{}, err
	}

	result.URL = state.URL
	result.CustomerID = state.CustomerID
	result.Token = state.Token
	result.CloudSpaceID = state.CloudSpaceID
	result.PortForwardID = state.PortForwardID

	return result, nil
}

func (pf PortForward) Update(ctx p.Context, id string, state PortForwardState, input PortForwardArgs, preview bool) (PortForwardState, error) {
	if preview {
		return state, nil
	}
	state.Token = input.Token
	u, err := url.Parse(fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/portforwards/%s", state.URL, state.CustomerID, state.CloudSpaceID, state.PortForwardID))
	if err != nil {
		return state, err
	}

	q := u.Query()
	q.Set("local_port", strconv.Itoa(input.LocalPort))
	q.Set("public_port", strconv.Itoa(input.PublicPort))
	q.Set("protocol", input.Protocol)
	q.Set("vm_id", strconv.Itoa(input.VirtualMachineID))
	q.Set("public_ip", input.PublicIP)

	if input.TillPublicPort != nil {
		q.Set("till_public_port", strconv.Itoa(*input.TillPublicPort))
	}

	if input.NestedCSID != nil {
		q.Set("nested_cs_id", *input.NestedCSID)
	}

	u.RawQuery = q.Encode()
	url := u.String()

	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(nil))
	if err != nil {
		fmt.Printf("Error making API request for %s: %v", id, err)
		return PortForwardState{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", input.Token))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error updating resource %s: %v\n", id, err)
		return PortForwardState{}, err
	}
	defer resp.Body.Close()

	getURL := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/portforwards", input.URL, input.CustomerID, input.CloudSpaceID)
	getReq, err := http.NewRequest("GET", getURL, nil)
	if err != nil {
		fmt.Printf("Error making API request for %s: %v", id, err)
		return PortForwardState{}, err
	}
	getReq.Header.Set("Content-Type", "application/json")
	getReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", input.Token))

	getResp, err := client.Do(getReq)
	if err != nil {
		fmt.Printf("Error retrieving resource %s: %v\n", id, err)
		return PortForwardState{}, err
	}
	defer getResp.Body.Close()

	var result struct {
		Result []struct {
			PortForwardID    string  `json:"portforward_id"`
			LocalPort        int     `json:"local_port"`
			PublicPort       int     `json:"public_port"`
			Protocol         string  `json:"protocol"`
			VirtualMachineID int     `json:"vm_id"`
			PublicIP         string  `json:"public_ip"`
			TillPublicPort   *int    `json:"till_public_port"`
			NestedCSID       *string `json:"nested_cs_id"`
		} `json:"result"`
	}
	if err := json.NewDecoder(getResp.Body).Decode(&result); err != nil {
		return PortForwardState{}, err
	}

	for _, pf := range result.Result {
		if pf.LocalPort == input.LocalPort && pf.PublicPort == input.PublicPort {
			state.PortForwardID = pf.PortForwardID
			break
		}
	}

	updatedState, err := pf.Read(ctx, id, state)
	if err != nil {
		return state, err
	}

	return updatedState, nil
}

func (PortForward) Delete(ctx p.Context, id string, state PortForwardState) error {
	pfUrl := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/portforwards/%s", state.URL, state.CustomerID, state.CloudSpaceID, state.PortForwardID)

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", pfUrl, bytes.NewBuffer(nil))
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
