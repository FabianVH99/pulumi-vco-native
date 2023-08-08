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
	ExternalCloudspaceID string `pulumi:"external_cloudspace_id" json:"external_cloudspace_id"`
}

type CpuTopology struct {
	Cores   int `pulumi:"cores" json:"cores"`
	Threads int `pulumi:"threads" json:"threads"`
	Sockets int `pulumi:"sockets" json:"sockets"`
}

type VirtualMachineState struct {
	VirtualMachineArgs
	VirtualMachineID     int                `pulumi:"vm_id" json:"vm_id"`
	Name                 string             `pulumi:"name" json:"name"`
	Description          string             `pulumi:"description" json:"description"`
	BootDiskID           int                `pulumi:"boot_disk_id" json:"boot_disk_id"`
	CloudspaceID         string             `pulumi:"cloudspace_id" json:"cloudspace_id"`
	Status               string             `pulumi:"status" json:"status"`
	UpdateTime           float64            `pulumi:"update_time" json:"update_time"`
	CreationTime         float64            `pulumi:"creation_time" json:"creation_time"`
	Vcpus                int                `pulumi:"vcpus" json:"vcpus"`
	Storage              int                `pulumi:"storage" json:"storage"`
	Memory               int                `pulumi:"memory" json:"memory"`
	AgentStatus          string             `pulumi:"agent_status" json:"agent_status"`
	ImageID              int                `pulumi:"image_id" json:"image_id"`
	Locked               bool               `pulumi:"locked" json:"locked"`
	StackID              int                `pulumi:"stack_id" json:"stack_id"`
	ImpactAlertHook      string             `pulumi:"impact_alert_hook" json:"impact_alert_hook"`
	OsImageName          string             `pulumi:"os_image_name" json:"os_image_name"`
	Hostname             string             `pulumi:"hostname" json:"hostname"`
	AntiAffinityGroupIDs []string           `pulumi:"anti_affinity_group_ids" json:"anti_affinity_group_ids"`
	Disks                []VmDisk           `pulumi:"disks" json:"disks"`
	OsAccounts           []OsAccount        `pulumi:"os_accounts" json:"os_accounts"`
	NetworkInterfaces    []NetworkInterface `pulumi:"network_interfaces" json:"network_interfaces"`
	Appliance            bool               `pulumi:"appliance" json:"appliance"`
	CpusPinStatus        bool               `pulumi:"cpus_pin_status" json:"cpus_pin_status"`
	BootType             string             `pulumi:"boot_type" json:"boot_type"`
	CpuTopology          CpuTopology        `pulumi:"cpu_topology" json:"cpu_topology"`
}

type VirtualMachineArgs struct {
	URL           string  `pulumi:"url"`
	Token         string  `pulumi:"token"`
	CustomerID    string  `pulumi:"customerID"`
	CloudSpaceID  string  `pulumi:"cloudspace_id"`
	Name          string  `pulumi:"name"`
	Description   string  `pulumi:"description"`
	DataDisks     *[]int  `pulumi:"data_disks"`
	Vcpus         int     `pulumi:"vcpus"`
	Memory        int     `pulumi:"memory"`
	PrivateIP     *string `pulumi:"private_ip"`
	UserData      *string `pulumi:"user_data"`
	ImageID       *int    `pulumi:"image_id"`
	DiskSize      *int    `pulumi:"disk_size"`
	CdromID       *int    `pulumi:"cdrom_id"`
	BootDiskID    *int    `pulumi:"boot_disk_id"`
	OsType        *string `pulumi:"os_type"`
	OsName        *string `pulumi:"os_name"`
	EnableVMAgent *bool   `pulumi:"enable_vm_agent"`
	SnapshotID    *string `pulumi:"snapshot_id"`
	AllVMDisks    *bool   `pulumi:"all_vm_disks"`
	Acronis       *bool   `pulumi:"acronis"`
	BootType      *string `pulumi:"boot_type"`
}

func (vm VirtualMachine) Create(ctx p.Context, name string, input VirtualMachineArgs, preview bool) (string, VirtualMachineState, error) {
	state := VirtualMachineState{VirtualMachineArgs: input}

	u, err := url.Parse(fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/vms", input.URL, input.CustomerID, input.CloudSpaceID))
	if err != nil {
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

	virtualMachineID := result["vm_id"].(float64)
	state.VirtualMachineID = int(virtualMachineID)

	updatedState, err := vm.Read(ctx, state, input)
	if err != nil {
		return "", state, err
	}

	return name, updatedState, nil
}

func (VirtualMachine) Read(ctx p.Context, state VirtualMachineState, input VirtualMachineArgs) (VirtualMachineState, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/vms/%d", input.URL, input.CustomerID, input.CloudSpaceID, state.VirtualMachineID)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return VirtualMachineState{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", input.Token))

	resp, err := client.Do(req)
	if err != nil {
		return VirtualMachineState{}, err
	}
	defer resp.Body.Close()

	var result VirtualMachineState
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return VirtualMachineState{}, err
	}

	return result, nil
}

func (VirtualMachine) Delete(ctx p.Context, state VirtualMachineState, input VirtualMachineArgs) (bool, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/vms/%d", input.URL, input.CustomerID, input.CloudSpaceID, state.VirtualMachineID)
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
