package virtual_machine

import (
	"bytes"
	"encoding/json"
	"fmt"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type VirtualMachineNIC struct{}

type VirtualMachineNICState struct {
	VirtualMachineNICArgs
	VirtualMachineID     int    `pulumi:"vm_id" json:"vm_id"`
	DeviceName           string `pulumi:"device_name" json:"device_name"`
	MacAddress           string `pulumi:"mac_address" json:"mac_address"`
	IPAddress            string `pulumi:"ip_address" json:"ip_address"`
	NetworkID            int    `pulumi:"network_id" json:"network_id"`
	NicType              string `pulumi:"nic_type" json:"nic_type"`
	Model                string `pulumi:"model" json:"model"`
	ExternalCloudspaceID string `pulumi:"external_cloudspace_id" json:"external_cloudspace_id"`
}

type VirtualMachineNICArgs struct {
	URL                  string  `pulumi:"url" provider:"secret"`
	Token                string  `pulumi:"token" provider:"secret"`
	CustomerID           string  `pulumi:"customerID" provider:"secret"`
	CloudSpaceID         string  `pulumi:"cloudspace_id"`
	VirtualMachineID     int     `pulumi:"vm_id"`
	ExternalNetworkID    int     `pulumi:"external_network_id"`
	ExternalNetworkIP    *string `pulumi:"external_network_ip,optional"`
	Model                *string `pulumi:"model,optional"`
	ExternalNetworkType  *string `pulumi:"external_network_type,optional"`
	ExternalCloudspaceID *string `pulumi:"external_cloudspace_id,optional"`
}

func (nic VirtualMachineNIC) WireDependencies(f infer.FieldSelector, args *VirtualMachineNICArgs, state *VirtualMachineNICState) {
	f.OutputField(&state.VirtualMachineID).DependsOn(f.InputField(&args.VirtualMachineID))
	f.OutputField(&state.NetworkID).DependsOn(f.InputField(&args.ExternalNetworkID))
}

func (nic VirtualMachineNIC) Create(ctx p.Context, name string, input VirtualMachineNICArgs, preview bool) (string, VirtualMachineNICState, error) {
	state := VirtualMachineNICState{VirtualMachineNICArgs: input}
	id, err := resource.NewUniqueHex(name, 8, 0)
	if err != nil {
		return "", state, err
	}
	if preview {
		return name, state, nil
	}

	u, err := url.Parse(fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/vms/%d/external-nics", input.URL, input.CustomerID, input.CloudSpaceID, input.VirtualMachineID))
	if err != nil {
		return "", state, err
	}

	q := u.Query()
	q.Set("external_network_id", strconv.Itoa(input.ExternalNetworkID))

	if input.ExternalNetworkIP != nil {
		q.Set("external_network_ip", *input.ExternalNetworkIP)
	}

	if input.Model != nil {
		q.Set("model", *input.Model)
	}

	if input.ExternalNetworkType != nil {
		q.Set("external_network_type", *input.ExternalNetworkType)
	}

	if input.ExternalCloudspaceID != nil {
		q.Set("external_cloudspace_id", *input.ExternalCloudspaceID)
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

	url = fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/vms/%d/external-nics", input.URL, input.CustomerID, input.CloudSpaceID, input.VirtualMachineID)
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return "", state, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", input.Token))

	resp, err = client.Do(req)
	if err != nil {
		return "", state, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", state, err
	}

	if nics, ok := result["result"].([]interface{}); ok && len(nics) > 0 {
		if nic, ok := nics[len(nics)-1].(map[string]interface{}); ok {
			if ipAddress, ok := nic["ip_address"].(string); ok {
				state.IPAddress = ipAddress
			}
		}
	}

	state.VirtualMachineID = input.VirtualMachineID

	updatedState, err := nic.Read(ctx, id, state, input)
	if err != nil {
		return "", state, err
	}

	return name, updatedState, nil
}

func (VirtualMachineNIC) Read(ctx p.Context, id string, state VirtualMachineNICState, input VirtualMachineNICArgs) (VirtualMachineNICState, error) {
	ipAddress := strings.Split(state.IPAddress, "/")[0]

	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/vms/%d/external-nics/%s", input.URL, input.CustomerID, input.CloudSpaceID, input.VirtualMachineID, ipAddress)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return VirtualMachineNICState{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", input.Token))

	resp, err := client.Do(req)
	if err != nil {
		return VirtualMachineNICState{}, err
	}
	defer resp.Body.Close()

	var result VirtualMachineNICState
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return VirtualMachineNICState{}, err
	}

	return result, nil
}

func (VirtualMachineNIC) Delete(ctx p.Context, id string, state VirtualMachineNICState, input VirtualMachineNICArgs) error {
	ipAddress := strings.Split(state.IPAddress, "/")[0]

	u, err := url.Parse(fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/vms/%d/external-nics/%s", input.URL, input.CustomerID, input.CloudSpaceID, input.VirtualMachineID, ipAddress))
	if err != nil {
		return err
	}

	q := u.Query()
	q.Set("external_network_id", strconv.Itoa(state.NetworkID))

	q.Set("external_network_type", strings.ToLower(state.NicType))

	if input.ExternalCloudspaceID != nil {
		q.Set("external_cloudspace_id", state.ExternalCloudspaceID)
	}

	u.RawQuery = q.Encode()
	url := u.String()

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
