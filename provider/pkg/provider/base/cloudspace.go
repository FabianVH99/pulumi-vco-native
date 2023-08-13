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

type CloudspaceState struct {
	CloudspaceArgs
	CloudSpaceID      string `pulumi:"cloudspace_id" json:"cloudspace_id"`
	Name              string `pulumi:"name" json:"name"`
	Status            string `pulumi:"status" json:"status"`
	ExternalNetworkIP string `pulumi:"external_network_ip" json:"external_network_ip"`
	ExternalNetworkID string `pulumi:"external_network_id" json:"external_network_id"`
	PrivateNetwork    string `pulumi:"private_network" json:"private_network"`
	LocalDomain       string `pulumi:"local_domain" json:"local_domain"`
	UpdateTime        int64  `pulumi:"update_time" json:"update_time"`
	CreationTime      int64  `pulumi:"creation_time" json:"creation_time"`
	RouterType        string `pulumi:"router_type" json:"router_type"`
	Location          string `pulumi:"location" json:"location"`
	CloudspaceMode    string `pulumi:"cloudspace_mode" json:"cloudspace_mode"`
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

	state.CloudSpaceID = result["cloudspace_id"].(string)

	updatedState, err := c.Read(nil, id, input, state)
	if err != nil {
		return "", state, err
	}

	return id, updatedState, nil
}

func (Cloudspace) Read(ctx p.Context, id string, input CloudspaceArgs, state CloudspaceState) (CloudspaceState, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s", input.URL, input.CustomerID, state.CloudSpaceID)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error making API request for %s: %v", id, err)
		return CloudspaceState{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", input.Token))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error retrieving resource %s: %v\n", id, err)
		return CloudspaceState{}, err
	}
	defer resp.Body.Close()

	var result CloudspaceState
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Printf("Error decoding response body for %s: %v\n", id, err)
		return CloudspaceState{}, err
	}

	return result, nil
}

func (Cloudspace) Update(ctx p.Context, id string, state CloudspaceState, input CloudspaceArgs) (CloudspaceState, error) {
	return CloudspaceState{}, nil
}

func (Cloudspace) Delete(ctx p.Context, id string, input CloudspaceArgs, state CloudspaceState) error {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s", input.URL, input.CustomerID, state.CloudSpaceID)
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
