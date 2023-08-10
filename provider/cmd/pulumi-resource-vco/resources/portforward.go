package resources

import (
	"bytes"
	"encoding/json"
	"fmt"
	p "github.com/pulumi/pulumi-go-provider"
	"net/http"
	"net/url"
	"strconv"
)

type PortForward struct{}

type PortForwardState struct {
	PortForwardArgs
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
	URL              string  `pulumi:"url"`
	Token            string  `pulumi:"token"`
	CustomerID       string  `pulumi:"customerID"`
	CloudSpaceID     string  `pulumi:"cloudspace_id"`
	LocalPort        int     `pulumi:"local_port"`
	PublicPort       int     `pulumi:"public_port"`
	Protocol         string  `pulumi:"protocol"`
	VirtualMachineID int     `pulumi:"vm_id"`
	PublicIP         string  `pulumi:"public_ip"`
	TillPublicPort   *int    `pulumi:"till_public_port,optional"`
	NestedCSID       *string `pulumi:"nested_cs_id,optional"`
}

func (pf PortForward) Create(ctx p.Context, name string, input PortForwardArgs, preview bool) (string, PortForwardState, error) {
	state := PortForwardState{PortForwardArgs: input}

	u, err := url.Parse(fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/portforwards", input.URL, input.CustomerID, input.CloudSpaceID))
	if err != nil {
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
		return "", state, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", state, err
	}

	state.PortForwardID = result["portforward_id"].(string)

	updatedState, err := pf.Read(ctx, state, input)
	if err != nil {
		return "", state, err
	}

	return name, updatedState, nil
}

func (PortForward) Read(ctx p.Context, state PortForwardState, input PortForwardArgs) (PortForwardState, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/portforwards/%s", input.URL, input.CustomerID, input.CloudSpaceID, state.PortForwardID)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PortForwardState{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", input.Token))

	resp, err := client.Do(req)
	if err != nil {
		return PortForwardState{}, err
	}
	defer resp.Body.Close()

	var result PortForwardState
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return PortForwardState{}, err
	}

	return result, nil
}

func (pf PortForward) Update(ctx p.Context, state PortForwardState, input PortForwardArgs) (PortForwardState, error) {
	u, err := url.Parse(fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/portforwards/%s", input.URL, input.CustomerID, input.CloudSpaceID, state.PortForwardID))
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
		return PortForwardState{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", input.Token))

	resp, err := client.Do(req)
	if err != nil {
		return PortForwardState{}, err
	}
	defer resp.Body.Close()

	getURL := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/portforwards", input.URL, input.CustomerID, input.CloudSpaceID)
	getReq, err := http.NewRequest("GET", getURL, nil)
	if err != nil {
		return PortForwardState{}, err
	}
	getReq.Header.Set("Content-Type", "application/json")
	getReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", input.Token))

	getResp, err := client.Do(getReq)
	if err != nil {
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

	updatedState, err := pf.Read(ctx, state, input)
	if err != nil {
		return state, err
	}

	return updatedState, nil
}

func (PortForward) Delete(ctx p.Context, state PortForwardState, input PortForwardArgs) (int, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/portforwards/%s", input.URL, input.CustomerID, input.CloudSpaceID, state.PortForwardID)

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(nil))
	if err != nil {
		return 0, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", input.Token))

	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}
