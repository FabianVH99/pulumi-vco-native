package base

import (
	"bytes"
	"encoding/json"
	"fmt"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"net/http"
)

type Cloudspace struct{}

func (c *CloudspaceState) run(state CloudspaceState) error {

	c.CloudSpaceID = state.CloudSpaceID
	c.Name = state.Name
	c.CloudSpaceID = state.CloudSpaceID
	c.Status = state.Status
	c.ExternalNetworkIP = state.ExternalNetworkIP
	c.ExternalNetworkID = state.ExternalNetworkID
	c.PrivateNetwork = state.PrivateNetwork
	c.LocalDomain = state.LocalDomain
	c.UpdateTime = state.UpdateTime
	c.CreationTime = state.CreationTime
	c.RouterType = state.RouterType
	c.Location = state.Location
	c.CloudspaceMode = state.CloudspaceMode

	return nil
}

func (c *CloudspaceState) id(id string) error {

	c.CloudSpaceID = id

	return nil
}

type CloudspaceState struct {
	CloudspaceArgs
	URL               string `json:"url" pulumi:"url"`
	Token             string `json:"token" pulumi:"token"`
	CustomerID        string `json:"customerID" pulumi:"customerID"`
	CloudSpaceID      string `json:"cloudspace_id" pulumi:"cloudspace_id"`
	Name              string `json:"name" pulumi:"name"`
	Status            string `json:"status" pulumi:"status"`
	ExternalNetworkIP string `json:"external_network_ip" pulumi:"external_network_ip"`
	ExternalNetworkID string `json:"external_network_id" pulumi:"external_network_id"`
	PrivateNetwork    string `json:"private_network" pulumi:"private_network"`
	LocalDomain       string `json:"local_domain" pulumi:"local_domain"`
	UpdateTime        int64  `json:"update_time" pulumi:"update_time"`
	CreationTime      int64  `json:"creation_time" pulumi:"creation_time"`
	RouterType        string `json:"router_type" pulumi:"router_type"`
	Location          string `json:"location" pulumi:"location"`
	CloudspaceMode    string `json:"cloudspace_mode" pulumi:"cloudspace_mode"`
}

type CloudspaceArgs struct {
	URL                string  `pulumi:"url" provider:"secret"`
	Token              string  `pulumi:"token" provider:"secret"`
	CustomerID         string  `pulumi:"customerID" provider:"secret"`
	Name               string  `pulumi:"name"`
	Location           string  `pulumi:"location"`
	PrivateNetwork     string  `pulumi:"private_network"`
	ExternalNetworkID  int     `pulumi:"external_network_id"`
	Private            bool    `pulumi:"private"`
	Host               *string `pulumi:"host,optional"`
	LocalDomain        *string `pulumi:"local_domain,optional"`
	VcpuQuota          *int    `pulumi:"vcpu_quota,optional"`
	VdiskSpaceQuota    *int    `pulumi:"vdisk_space_quota,optional"`
	MemoryQuota        *int    `pulumi:"memory_quota,optional"`
	PublicIPQuota      *int    `pulumi:"public_ip_quota,optional"`
	ParentCloudspaceID *string `pulumi:"parent_cloudspace_id,optional"`
	ImageID            *int    `pulumi:"image_id,optional"`
	CdromID            *int    `pulumi:"cdrom_id,optional"`
	Type               *string `pulumi:"type,optional"`
	DiskSize           *int    `pulumi:"disk_size,optional"`
	Vcpus              *int    `pulumi:"vcpus,optional"`
	Memory             *int    `pulumi:"memory,optional"`
}

func (c Cloudspace) WireDependencies(f infer.FieldSelector, args *CloudspaceArgs, state *CloudspaceState) {
	f.OutputField(&state.URL).DependsOn(f.InputField(&args.URL))
	f.OutputField(&state.Token).DependsOn(f.InputField(&args.Token))
	f.OutputField(&state.CustomerID).DependsOn(f.InputField(&args.CustomerID))
	f.OutputField(&state.Name).DependsOn(f.InputField(&args.Name))
	f.OutputField(&state.ExternalNetworkID).DependsOn(f.InputField(&args.ExternalNetworkID))
	f.OutputField(&state.Location).DependsOn(f.InputField(&args.Location))
	f.OutputField(&state.PrivateNetwork).DependsOn(f.InputField(&args.PrivateNetwork))
}

func (c Cloudspace) Create(ctx p.Context, name string, input CloudspaceArgs, preview bool) (string, CloudspaceState, error) {
	state := CloudspaceState{CloudspaceArgs: input}
	id, err := resource.NewUniqueHex(name, 8, 0)
	if err != nil {
		return "", state, err
	}
	if preview {
		return name, state, nil
	}

	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces", input.URL, input.CustomerID)
	payload := map[string]interface{}{
		"name":            input.Name,
		"location":        input.Location,
		"private_network": input.PrivateNetwork,
	}

	if input.Host != nil {
		payload["host"] = *input.Host
	}
	if input.LocalDomain != nil {
		payload["local_domain"] = *input.LocalDomain
	}
	if input.VcpuQuota != nil {
		payload["vcpu_quota"] = *input.VcpuQuota
	}
	if input.VdiskSpaceQuota != nil {
		payload["vdisk_space_quota"] = *input.VdiskSpaceQuota
	}
	if input.MemoryQuota != nil {
		payload["memory_quota"] = *input.MemoryQuota
	}
	if input.PublicIPQuota != nil {
		payload["public_ip_quota"] = *input.PublicIPQuota
	}
	firewall := map[string]interface{}{
		"external_network_id": input.ExternalNetworkID,
		"private":             input.Private,
	}
	if input.ParentCloudspaceID != nil {
		firewall["parent_cloudspace_id"] = *input.ParentCloudspaceID
	}

	custom := map[string]interface{}{}
	if input.ImageID != nil {
		custom["image_id"] = *input.ImageID
	}
	if input.CdromID != nil {
		custom["cdrom_id"] = *input.CdromID
	}
	if input.Type != nil {
		custom["type"] = *input.Type
	}
	if input.DiskSize != nil {
		custom["disk_size"] = *input.DiskSize
	}
	if input.Vcpus != nil {
		custom["vcpus"] = *input.Vcpus
	}
	if input.Memory != nil {
		custom["memory"] = *input.Memory
	}

	firewall["custom"] = custom

	payload["firewall"] = firewall

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", state, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Printf("Error making API request: %v", err)
		return "", state, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", input.Token))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making API request: %v", err)
		return "", state, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", state, err
	}

	state.CloudSpaceID = result["cloudspace_id"].(string)
	state.URL = input.URL
	state.CustomerID = input.CustomerID
	state.Token = input.Token
	state.id(result["cloudspace_id"].(string))

	updatedState, err := c.Read(nil, id, state)
	if err != nil {
		fmt.Printf("Error retrieving resource: %v", err)
		return "", state, err
	}

	return id, updatedState, nil
}

func (Cloudspace) Read(ctx p.Context, id string, state CloudspaceState) (CloudspaceState, error) {
	fmt.Printf("Delete method called!!")
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s", state.URL, state.CustomerID, state.CloudSpaceID)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return CloudspaceState{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", state.Token))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making API request: %v", err)
		return CloudspaceState{}, err
	}
	defer resp.Body.Close()

	var result CloudspaceState
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return CloudspaceState{}, err
	}

	result.run(result)

	result.URL = state.URL
	result.CustomerID = state.CustomerID
	result.Token = state.Token

	return result, nil
}

func (Cloudspace) Delete(ctx p.Context, id string, props CloudspaceState) error {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s", props.URL, props.CustomerID, props.CloudSpaceID)
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(nil))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", props.Token))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making API request: %v\n", err)
		return err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	return nil
}
