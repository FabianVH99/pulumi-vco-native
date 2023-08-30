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

type VirtualMachine struct{}

type VmDisk struct {
	Status      string `pulumi:"status" json:"status"`
	DiskSize    int    `pulumi:"disk_size" json:"disk_size"`
	DiskName    string `pulumi:"disk_name" json:"disk_name"`
	Description string `pulumi:"description" json:"description"`
	Exposed     bool   `pulumi:"exposed" json:"exposed"`
	PciBus      int    `pulumi:"pci_bus" json:"pci_bus"`
	PciSlot     int    `pulumi:"pci_slot" json:"pci_slot"`
	DiskID      int    `pulumi:"disk_id" json:"disk_id"`
	DiskType    string `pulumi:"disk_type" json:"disk_type"`
	Order       string `pulumi:"order" json:"order"`
}

type OsAccount struct {
	Login    string `pulumi:"login" json:"login"`
	Password string `pulumi:"password" json:"password"`
}

type NetworkInterface struct {
	DeviceName           string `pulumi:"device_name" json:"device_name"`
	MacAddress           string `pulumi:"mac_address" json:"mac_address"`
	IPAddress            string `pulumi:"ip_address" json:"ip_address"`
	NetworkID            int    `pulumi:"network_id" json:"network_id"`
	NicType              string `pulumi:"nic_type" json:"nic_type"`
	Model                string `pulumi:"model" json:"model"`
	ExternalCloudspaceID string `pulumi:"external_cloudspace_id,optional" json:"external_cloudspace_id"`
}

type CpuTopology struct {
	Cores   int `pulumi:"cores" json:"cores"`
	Threads int `pulumi:"threads" json:"threads"`
	Sockets int `pulumi:"sockets" json:"sockets"`
}

type VirtualMachineState struct {
	VirtualMachineArgs
	URL               string             `pulumi:"url"`
	Token             string             `pulumi:"token"`
	CustomerID        string             `pulumi:"customerID"`
	CloudSpaceID      string             `pulumi:"cloudspace_id"`
	VirtualMachineID  int                `pulumi:"vm_id" json:"vm_id"`
	Name              string             `pulumi:"name" json:"name"`
	Description       string             `pulumi:"description" json:"description"`
	BootDiskID        int                `pulumi:"boot_disk_id" json:"boot_disk_id"`
	CloudspaceID      string             `pulumi:"cloudspace_id" json:"cloudspace_id"`
	Status            string             `pulumi:"status" json:"status"`
	UpdateTime        float64            `pulumi:"update_time" json:"update_time"`
	CreationTime      float64            `pulumi:"creation_time" json:"creation_time"`
	Vcpus             int                `pulumi:"vcpus" json:"vcpus"`
	Storage           int                `pulumi:"storage" json:"storage"`
	Memory            int                `pulumi:"memory" json:"memory"`
	AgentStatus       string             `pulumi:"agent_status" json:"agent_status"`
	ImageID           int                `pulumi:"image_id" json:"image_id"`
	Locked            bool               `pulumi:"locked" json:"locked"`
	StackID           int                `pulumi:"stack_id" json:"stack_id"`
	ImpactAlertHook   string             `pulumi:"impact_alert_hook,optional" json:"impact_alert_hook"`
	OsImageName       string             `pulumi:"os_image_name" json:"os_image_name"`
	Hostname          string             `pulumi:"hostname" json:"hostname"`
	Disks             []VmDisk           `pulumi:"disks" json:"disks"`
	OsAccounts        []OsAccount        `pulumi:"os_accounts" json:"os_accounts"`
	NetworkInterfaces []NetworkInterface `pulumi:"network_interfaces" json:"network_interfaces"`
	Appliance         bool               `pulumi:"appliance" json:"appliance"`
	CpusPinStatus     bool               `pulumi:"cpus_pin_status" json:"cpus_pin_status"`
	BootType          string             `pulumi:"boot_type" json:"boot_type"`
	CpuTopology       CpuTopology        `pulumi:"cpu_topology" json:"cpu_topology"`
}

type VirtualMachineArgs struct {
	URL           string  `pulumi:"url" provider:"secret"`
	Token         string  `pulumi:"token" provider:"secret"`
	CustomerID    string  `pulumi:"customerID" provider:"secret"`
	CloudSpaceID  string  `pulumi:"cloudspace_id"`
	Name          string  `pulumi:"name"`
	Description   string  `pulumi:"description"`
	DataDisks     *[]int  `pulumi:"data_disks,optional"`
	Vcpus         int     `pulumi:"vcpus"`
	Memory        int     `pulumi:"memory"`
	PrivateIP     *string `pulumi:"private_ip,optional"`
	UserData      *string `pulumi:"user_data,optional"`
	ImageID       *int    `pulumi:"image_id,optional"`
	DiskSize      *int    `pulumi:"disk_size,optional"`
	CdromID       *int    `pulumi:"cdrom_id,optional"`
	BootDiskID    *int    `pulumi:"boot_disk_id,optional"`
	OsType        *string `pulumi:"os_type,optional"`
	OsName        *string `pulumi:"os_name,optional"`
	EnableVMAgent *bool   `pulumi:"enable_vm_agent,optional"`
	SnapshotID    *string `pulumi:"snapshot_id,optional"`
	AllVMDisks    *bool   `pulumi:"all_vm_disks,optional"`
	Acronis       *bool   `pulumi:"acronis,optional"`
	BootType      *string `pulumi:"boot_type,optional"`
}

func (vm VirtualMachine) WireDependencies(f infer.FieldSelector, args *VirtualMachineArgs, state *VirtualMachineState) {
	f.OutputField(&state.URL).DependsOn(f.InputField(&args.URL))
	f.OutputField(&state.Token).DependsOn(f.InputField(&args.Token))
	f.OutputField(&state.CustomerID).DependsOn(f.InputField(&args.CustomerID))
	f.OutputField(&state.CloudSpaceID).DependsOn(f.InputField(&args.CloudSpaceID))
	f.OutputField(&state.CloudSpaceID).DependsOn(f.InputField(&args.CloudSpaceID))
	f.OutputField(&state.Name).DependsOn(f.InputField(&args.Name))
	f.OutputField(&state.Description).DependsOn(f.InputField(&args.Description))
	f.OutputField(&state.Vcpus).DependsOn(f.InputField(&args.Vcpus))
	f.OutputField(&state.Memory).DependsOn(f.InputField(&args.Memory))
}

func (vm VirtualMachine) Create(ctx p.Context, name string, input VirtualMachineArgs, preview bool) (string, VirtualMachineState, error) {
	state := VirtualMachineState{VirtualMachineArgs: input}
	id, err := resource.NewUniqueHex(name, 8, 0)
	if err != nil {
		return "", state, err
	}
	if preview {
		return name, state, nil
	}

	u, err := url.Parse(fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/vms", input.URL, input.CustomerID, input.CloudSpaceID))
	if err != nil {
		fmt.Printf("Error making API request for %s: %v", id, err)
		return "", state, err
	}

	q := u.Query()
	q.Set("name", input.Name)
	q.Set("description", input.Description)
	q.Set("vcpus", strconv.Itoa(input.Vcpus))
	q.Set("memory", strconv.Itoa(input.Memory))

	if input.DataDisks != nil {
		dataDisks := *input.DataDisks
		q.Set("data_disks", fmt.Sprintf("%v", dataDisks))
	}

	if input.PrivateIP != nil {
		q.Set("private_ip", *input.PrivateIP)
	}

	if input.ImageID != nil {
		q.Set("image_id", strconv.Itoa(*input.ImageID))
	}

	if input.DiskSize != nil {
		q.Set("disk_size", strconv.Itoa(*input.DiskSize))
	}

	if input.CdromID != nil {
		q.Set("cdrom_id", strconv.Itoa(*input.CdromID))
	}

	if input.BootDiskID != nil {
		q.Set("boot_disk_id", strconv.Itoa(*input.BootDiskID))
	}

	if input.OsType != nil {
		q.Set("os_type", *input.OsType)
	}

	if input.OsName != nil {
		q.Set("os_name", *input.OsName)
	}

	if input.EnableVMAgent != nil {
		q.Set("enable_vm_agent", strconv.FormatBool(*input.EnableVMAgent))
	}

	if input.SnapshotID != nil {
		q.Set("snapshot_id", *input.SnapshotID)
	}

	if input.AllVMDisks != nil {
		q.Set("all_vm_disks", strconv.FormatBool(*input.AllVMDisks))
	}

	if input.Acronis != nil {
		q.Set("acronis", strconv.FormatBool(*input.Acronis))
	}

	if input.BootType != nil {
		q.Set("boot_type", *input.BootType)
	}

	u.RawQuery = q.Encode()
	url := u.String()

	payload := map[string]interface{}{}

	if input.UserData != nil {
		payload["user_data"] = *input.UserData
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("Error constructing JSON payload %s: %v", id, err)
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
		fmt.Printf("Error creating resource %s: %v", id, err)
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
	virtualMachineID := result["vm_id"].(float64)
	state.VirtualMachineID = int(virtualMachineID)

	updatedState, err := vm.Read(ctx, id, state)
	if err != nil {
		return "", state, err
	}

	return id, updatedState, nil
}

func (VirtualMachine) Read(ctx p.Context, id string, state VirtualMachineState) (VirtualMachineState, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/vms/%d", state.URL, state.CustomerID, state.CloudSpaceID, state.VirtualMachineID)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error making API request for %s: %v", id, err)
		return VirtualMachineState{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", state.Token))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error updating resource %s: %v\n", id, err)
		return VirtualMachineState{}, err
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

	var result VirtualMachineState
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return VirtualMachineState{}, err
	}

	result.URL = state.URL
	result.CustomerID = state.CustomerID
	result.Token = state.Token
	result.CloudSpaceID = state.CloudSpaceID

	return result, nil
}

func (VirtualMachine) Update(ctx p.Context, id string, state VirtualMachineState, input VirtualMachineArgs, preview bool) (VirtualMachineState, error) {
	if preview {
		return state, nil
	}
	state.Token = input.Token
	return state, nil
}

func (VirtualMachine) Delete(ctx p.Context, id string, state VirtualMachineState) error {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/vms/%d", state.URL, state.CustomerID, state.CloudSpaceID, state.VirtualMachineID)
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
