// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package resources

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"internal"
)

type VirtualMachine struct {
	pulumi.CustomResourceState

	Acronis                 pulumi.BoolPtrOutput        `pulumi:"acronis"`
	Agent_status            pulumi.StringOutput         `pulumi:"agent_status"`
	All_vm_disks            pulumi.BoolPtrOutput        `pulumi:"all_vm_disks"`
	Anti_affinity_group_ids pulumi.StringArrayOutput    `pulumi:"anti_affinity_group_ids"`
	Appliance               pulumi.BoolOutput           `pulumi:"appliance"`
	Boot_disk_id            pulumi.IntOutput            `pulumi:"boot_disk_id"`
	Boot_type               pulumi.StringOutput         `pulumi:"boot_type"`
	Cdrom_id                pulumi.IntPtrOutput         `pulumi:"cdrom_id"`
	Cloudspace_id           pulumi.StringOutput         `pulumi:"cloudspace_id"`
	Cpu_topology            CpuTopologyOutput           `pulumi:"cpu_topology"`
	Cpus_pin_status         pulumi.BoolOutput           `pulumi:"cpus_pin_status"`
	Creation_time           pulumi.Float64Output        `pulumi:"creation_time"`
	CustomerID              pulumi.StringOutput         `pulumi:"customerID"`
	Data_disks              pulumi.IntArrayOutput       `pulumi:"data_disks"`
	Description             pulumi.StringOutput         `pulumi:"description"`
	Disk_size               pulumi.IntPtrOutput         `pulumi:"disk_size"`
	Disks                   VmDiskArrayOutput           `pulumi:"disks"`
	Enable_vm_agent         pulumi.BoolPtrOutput        `pulumi:"enable_vm_agent"`
	Hostname                pulumi.StringOutput         `pulumi:"hostname"`
	Image_id                pulumi.IntOutput            `pulumi:"image_id"`
	Impact_alert_hook       pulumi.StringOutput         `pulumi:"impact_alert_hook"`
	Locked                  pulumi.BoolOutput           `pulumi:"locked"`
	Memory                  pulumi.IntOutput            `pulumi:"memory"`
	Name                    pulumi.StringOutput         `pulumi:"name"`
	Network_interfaces      NetworkInterfaceArrayOutput `pulumi:"network_interfaces"`
	Os_accounts             OsAccountArrayOutput        `pulumi:"os_accounts"`
	Os_image_name           pulumi.StringOutput         `pulumi:"os_image_name"`
	Os_name                 pulumi.StringPtrOutput      `pulumi:"os_name"`
	Os_type                 pulumi.StringPtrOutput      `pulumi:"os_type"`
	Private_ip              pulumi.StringPtrOutput      `pulumi:"private_ip"`
	Snapshot_id             pulumi.StringPtrOutput      `pulumi:"snapshot_id"`
	Stack_id                pulumi.IntOutput            `pulumi:"stack_id"`
	Status                  pulumi.StringOutput         `pulumi:"status"`
	Storage                 pulumi.IntOutput            `pulumi:"storage"`
	Token                   pulumi.StringOutput         `pulumi:"token"`
	Update_time             pulumi.Float64Output        `pulumi:"update_time"`
	Url                     pulumi.StringOutput         `pulumi:"url"`
	User_data               pulumi.StringPtrOutput      `pulumi:"user_data"`
	Vcpus                   pulumi.IntOutput            `pulumi:"vcpus"`
	Vm_id                   pulumi.IntOutput            `pulumi:"vm_id"`
}

// NewVirtualMachine registers a new resource with the given unique name, arguments, and options.
func NewVirtualMachine(ctx *pulumi.Context,
	name string, args *VirtualMachineArgs, opts ...pulumi.ResourceOption) (*VirtualMachine, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Cloudspace_id == nil {
		return nil, errors.New("invalid value for required argument 'Cloudspace_id'")
	}
	if args.CustomerID == nil {
		return nil, errors.New("invalid value for required argument 'CustomerID'")
	}
	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.Memory == nil {
		return nil, errors.New("invalid value for required argument 'Memory'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.Token == nil {
		return nil, errors.New("invalid value for required argument 'Token'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	if args.Vcpus == nil {
		return nil, errors.New("invalid value for required argument 'Vcpus'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VirtualMachine
	err := ctx.RegisterResource("vco:resources:VirtualMachine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualMachine gets an existing VirtualMachine resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualMachine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineState, opts ...pulumi.ResourceOption) (*VirtualMachine, error) {
	var resource VirtualMachine
	err := ctx.ReadResource("vco:resources:VirtualMachine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualMachine resources.
type virtualMachineState struct {
}

type VirtualMachineState struct {
}

func (VirtualMachineState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineState)(nil)).Elem()
}

type virtualMachineArgs struct {
	Acronis         *bool   `pulumi:"acronis"`
	All_vm_disks    *bool   `pulumi:"all_vm_disks"`
	Boot_disk_id    *int    `pulumi:"boot_disk_id"`
	Boot_type       *string `pulumi:"boot_type"`
	Cdrom_id        *int    `pulumi:"cdrom_id"`
	Cloudspace_id   string  `pulumi:"cloudspace_id"`
	CustomerID      string  `pulumi:"customerID"`
	Data_disks      []int   `pulumi:"data_disks"`
	Description     string  `pulumi:"description"`
	Disk_size       *int    `pulumi:"disk_size"`
	Enable_vm_agent *bool   `pulumi:"enable_vm_agent"`
	Image_id        *int    `pulumi:"image_id"`
	Memory          int     `pulumi:"memory"`
	Name            string  `pulumi:"name"`
	Os_name         *string `pulumi:"os_name"`
	Os_type         *string `pulumi:"os_type"`
	Private_ip      *string `pulumi:"private_ip"`
	Snapshot_id     *string `pulumi:"snapshot_id"`
	Token           string  `pulumi:"token"`
	Url             string  `pulumi:"url"`
	User_data       *string `pulumi:"user_data"`
	Vcpus           int     `pulumi:"vcpus"`
}

// The set of arguments for constructing a VirtualMachine resource.
type VirtualMachineArgs struct {
	Acronis         pulumi.BoolPtrInput
	All_vm_disks    pulumi.BoolPtrInput
	Boot_disk_id    pulumi.IntPtrInput
	Boot_type       pulumi.StringPtrInput
	Cdrom_id        pulumi.IntPtrInput
	Cloudspace_id   pulumi.StringInput
	CustomerID      pulumi.StringInput
	Data_disks      pulumi.IntArrayInput
	Description     pulumi.StringInput
	Disk_size       pulumi.IntPtrInput
	Enable_vm_agent pulumi.BoolPtrInput
	Image_id        pulumi.IntPtrInput
	Memory          pulumi.IntInput
	Name            pulumi.StringInput
	Os_name         pulumi.StringPtrInput
	Os_type         pulumi.StringPtrInput
	Private_ip      pulumi.StringPtrInput
	Snapshot_id     pulumi.StringPtrInput
	Token           pulumi.StringInput
	Url             pulumi.StringInput
	User_data       pulumi.StringPtrInput
	Vcpus           pulumi.IntInput
}

func (VirtualMachineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineArgs)(nil)).Elem()
}

type VirtualMachineInput interface {
	pulumi.Input

	ToVirtualMachineOutput() VirtualMachineOutput
	ToVirtualMachineOutputWithContext(ctx context.Context) VirtualMachineOutput
}

func (*VirtualMachine) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachine)(nil)).Elem()
}

func (i *VirtualMachine) ToVirtualMachineOutput() VirtualMachineOutput {
	return i.ToVirtualMachineOutputWithContext(context.Background())
}

func (i *VirtualMachine) ToVirtualMachineOutputWithContext(ctx context.Context) VirtualMachineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineOutput)
}

type VirtualMachineOutput struct{ *pulumi.OutputState }

func (VirtualMachineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachine)(nil)).Elem()
}

func (o VirtualMachineOutput) ToVirtualMachineOutput() VirtualMachineOutput {
	return o
}

func (o VirtualMachineOutput) ToVirtualMachineOutputWithContext(ctx context.Context) VirtualMachineOutput {
	return o
}

func (o VirtualMachineOutput) Acronis() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.BoolPtrOutput { return v.Acronis }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineOutput) Agent_status() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.Agent_status }).(pulumi.StringOutput)
}

func (o VirtualMachineOutput) All_vm_disks() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.BoolPtrOutput { return v.All_vm_disks }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineOutput) Anti_affinity_group_ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringArrayOutput { return v.Anti_affinity_group_ids }).(pulumi.StringArrayOutput)
}

func (o VirtualMachineOutput) Appliance() pulumi.BoolOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.BoolOutput { return v.Appliance }).(pulumi.BoolOutput)
}

func (o VirtualMachineOutput) Boot_disk_id() pulumi.IntOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.IntOutput { return v.Boot_disk_id }).(pulumi.IntOutput)
}

func (o VirtualMachineOutput) Boot_type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.Boot_type }).(pulumi.StringOutput)
}

func (o VirtualMachineOutput) Cdrom_id() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.IntPtrOutput { return v.Cdrom_id }).(pulumi.IntPtrOutput)
}

func (o VirtualMachineOutput) Cloudspace_id() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.Cloudspace_id }).(pulumi.StringOutput)
}

func (o VirtualMachineOutput) Cpu_topology() CpuTopologyOutput {
	return o.ApplyT(func(v *VirtualMachine) CpuTopologyOutput { return v.Cpu_topology }).(CpuTopologyOutput)
}

func (o VirtualMachineOutput) Cpus_pin_status() pulumi.BoolOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.BoolOutput { return v.Cpus_pin_status }).(pulumi.BoolOutput)
}

func (o VirtualMachineOutput) Creation_time() pulumi.Float64Output {
	return o.ApplyT(func(v *VirtualMachine) pulumi.Float64Output { return v.Creation_time }).(pulumi.Float64Output)
}

func (o VirtualMachineOutput) CustomerID() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.CustomerID }).(pulumi.StringOutput)
}

func (o VirtualMachineOutput) Data_disks() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.IntArrayOutput { return v.Data_disks }).(pulumi.IntArrayOutput)
}

func (o VirtualMachineOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o VirtualMachineOutput) Disk_size() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.IntPtrOutput { return v.Disk_size }).(pulumi.IntPtrOutput)
}

func (o VirtualMachineOutput) Disks() VmDiskArrayOutput {
	return o.ApplyT(func(v *VirtualMachine) VmDiskArrayOutput { return v.Disks }).(VmDiskArrayOutput)
}

func (o VirtualMachineOutput) Enable_vm_agent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.BoolPtrOutput { return v.Enable_vm_agent }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.Hostname }).(pulumi.StringOutput)
}

func (o VirtualMachineOutput) Image_id() pulumi.IntOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.IntOutput { return v.Image_id }).(pulumi.IntOutput)
}

func (o VirtualMachineOutput) Impact_alert_hook() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.Impact_alert_hook }).(pulumi.StringOutput)
}

func (o VirtualMachineOutput) Locked() pulumi.BoolOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.BoolOutput { return v.Locked }).(pulumi.BoolOutput)
}

func (o VirtualMachineOutput) Memory() pulumi.IntOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.IntOutput { return v.Memory }).(pulumi.IntOutput)
}

func (o VirtualMachineOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualMachineOutput) Network_interfaces() NetworkInterfaceArrayOutput {
	return o.ApplyT(func(v *VirtualMachine) NetworkInterfaceArrayOutput { return v.Network_interfaces }).(NetworkInterfaceArrayOutput)
}

func (o VirtualMachineOutput) Os_accounts() OsAccountArrayOutput {
	return o.ApplyT(func(v *VirtualMachine) OsAccountArrayOutput { return v.Os_accounts }).(OsAccountArrayOutput)
}

func (o VirtualMachineOutput) Os_image_name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.Os_image_name }).(pulumi.StringOutput)
}

func (o VirtualMachineOutput) Os_name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.Os_name }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineOutput) Os_type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.Os_type }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineOutput) Private_ip() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.Private_ip }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineOutput) Snapshot_id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.Snapshot_id }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineOutput) Stack_id() pulumi.IntOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.IntOutput { return v.Stack_id }).(pulumi.IntOutput)
}

func (o VirtualMachineOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o VirtualMachineOutput) Storage() pulumi.IntOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.IntOutput { return v.Storage }).(pulumi.IntOutput)
}

func (o VirtualMachineOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

func (o VirtualMachineOutput) Update_time() pulumi.Float64Output {
	return o.ApplyT(func(v *VirtualMachine) pulumi.Float64Output { return v.Update_time }).(pulumi.Float64Output)
}

func (o VirtualMachineOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

func (o VirtualMachineOutput) User_data() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.User_data }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineOutput) Vcpus() pulumi.IntOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.IntOutput { return v.Vcpus }).(pulumi.IntOutput)
}

func (o VirtualMachineOutput) Vm_id() pulumi.IntOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.IntOutput { return v.Vm_id }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualMachineInput)(nil)).Elem(), &VirtualMachine{})
	pulumi.RegisterOutputType(VirtualMachineOutput{})
}
