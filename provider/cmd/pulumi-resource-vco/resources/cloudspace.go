package resources

import (
	"bytes"
	"encoding/json"
	"fmt"
	p "github.com/pulumi/pulumi-go-provider"
	"net/http"
)

type Cloudspace struct{}

type ResourceLimits struct {
	MemoryQuota          float64 `json:"memory_quota" pulumi:"memory_quota"`
	VdiskSpaceQuota      float64 `json:"vdisk_space_quota" pulumi:"vdisk_space_quota"`
	PublicIPQuota        float64 `json:"public_ip_quota" pulumi:"public_ip_quota"`
	VcpuQuota            float64 `json:"vcpu_quota" pulumi:"vcpu_quota"`
	ExternalNetworkQuota float64 `json:"external_network_quota" pulumi:"external_network_quota"`
}

type CloudspaceState struct {
	CloudspaceArgs
	CloudSpaceID      string         `json:"cloudspace_id" pulumi:"cloudspace_id"`
	Name              string         `json:"name" pulumi:"name"`
	Status            string         `json:"status" pulumi:"status"`
	ExternalNetworkIP string         `json:"external_network_ip" pulumi:"external_network_ip"`
	ExternalNetworkID string         `json:"external_network_id" pulumi:"external_network_id"`
	PrivateNetwork    string         `json:"private_network" pulumi:"private_network"`
	LocalDomain       string         `json:"local_domain" pulumi:"local_domain"`
	UpdateTime        int64          `json:"update_time" pulumi:"update_time"`
	CreationTime      int64          `json:"creation_time" pulumi:"creation_time"`
	RouterType        string         `json:"router_type" pulumi:"router_type"`
	Location          string         `json:"location" pulumi:"location"`
	CloudspaceMode    string         `json:"cloudspace_mode" pulumi:"cloudspace_mode"`
	ResourceLimits    ResourceLimits `json:"resource_limits" pulumi:"resource_limits"`
}

type FirewallCustom struct {
	ImageID  *int    `pulumi:"image_id"`
	CdromID  *int    `pulumi:"cdrom_id"`
	Type     *string `pulumi:"type"`
	DiskSize *int    `pulumi:"disk_size"`
	Vcpus    *int    `pulumi:"vcpus"`
	Memory   *int    `pulumi:"memory"`
}

type CloudspaceArgs struct {
	URL                string          `pulumi:"url"`
	Token              string          `pulumi:"token"`
	CustomerID         string          `pulumi:"customerID"`
	Name               string          `pulumi:"name"`
	Location           string          `pulumi:"location"`
	PrivateNetwork     string          `pulumi:"privateNetwork"`
	ExternalNetworkID  int             `pulumi:"externalNetworkID"`
	Private            bool            `pulumi:"private"`
	Host               *string         `pulumi:"host"`
	LocalDomain        *string         `pulumi:"local_domain"`
	VcpuQuota          *int            `pulumi:"vcpu_quota"`
	VdiskSpaceQuota    *int            `pulumi:"vdisk_space_quota"`
	MemoryQuota        *int            `pulumi:"memory_quota"`
	PublicIPQuota      *int            `pulumi:"public_ip_quota"`
	ParentCloudspaceID *string         `pulumi:"parent_cloudspace_id"`
	FirewallCustom     *FirewallCustom `pulumi:"firewall_custom"`
}

func (c Cloudspace) Create(ctx p.Context, name string, input CloudspaceArgs, preview bool) (string, CloudspaceState, error) {
	state := CloudspaceState{CloudspaceArgs: input}

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
	if input.FirewallCustom != nil {
		custom := map[string]interface{}{}
		if input.FirewallCustom.ImageID != nil {
			custom["image_id"] = *input.FirewallCustom.ImageID
		}
		if input.FirewallCustom.CdromID != nil {
			custom["cdrom_id"] = *input.FirewallCustom.CdromID
		}
		if input.FirewallCustom.Type != nil {
			custom["type"] = *input.FirewallCustom.Type
		}
		if input.FirewallCustom.DiskSize != nil {
			custom["disk_size"] = *input.FirewallCustom.DiskSize
		}
		if input.FirewallCustom.Vcpus != nil {
			custom["vcpus"] = *input.FirewallCustom.Vcpus
		}
		if input.FirewallCustom.Memory != nil {
			custom["memory"] = *input.FirewallCustom.Memory
		}

		firewall["custom"] = custom
	}
	payload["firewall"] = firewall

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", state, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
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

	state.CloudSpaceID = result["cloudspace_id"].(string)

	updatedState, err := c.Read(ctx, state, input)
	if err != nil {
		return "", state, err
	}

	return name, updatedState, nil
}

func (Cloudspace) Read(ctx p.Context, state CloudspaceState, input CloudspaceArgs) (CloudspaceState, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s", input.URL, input.CustomerID, state.CloudSpaceID)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return CloudspaceState{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", input.Token))

	resp, err := client.Do(req)
	if err != nil {
		return CloudspaceState{}, err
	}
	defer resp.Body.Close()

	var result CloudspaceState
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return CloudspaceState{}, err
	}

	return result, nil
}

func (Cloudspace) Delete(ctx p.Context, state CloudspaceState, input CloudspaceArgs) (bool, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s", input.URL, input.CustomerID, state.CloudSpaceID)
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
