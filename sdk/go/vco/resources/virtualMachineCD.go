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

type VirtualMachineCD struct {
	pulumi.CustomResourceState

	Cdrom_id      pulumi.IntOutput    `pulumi:"cdrom_id"`
	Cloudspace_id pulumi.StringOutput `pulumi:"cloudspace_id"`
	CustomerID    pulumi.StringOutput `pulumi:"customerID"`
	Description   pulumi.StringOutput `pulumi:"description"`
	Disk_size     pulumi.StringOutput `pulumi:"disk_size"`
	Name          pulumi.StringOutput `pulumi:"name"`
	Status        pulumi.StringOutput `pulumi:"status"`
	Token         pulumi.StringOutput `pulumi:"token"`
	Url           pulumi.StringOutput `pulumi:"url"`
	Vm_id         pulumi.IntOutput    `pulumi:"vm_id"`
}

// NewVirtualMachineCD registers a new resource with the given unique name, arguments, and options.
func NewVirtualMachineCD(ctx *pulumi.Context,
	name string, args *VirtualMachineCDArgs, opts ...pulumi.ResourceOption) (*VirtualMachineCD, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Cdrom_id == nil {
		return nil, errors.New("invalid value for required argument 'Cdrom_id'")
	}
	if args.Cloudspace_id == nil {
		return nil, errors.New("invalid value for required argument 'Cloudspace_id'")
	}
	if args.CustomerID == nil {
		return nil, errors.New("invalid value for required argument 'CustomerID'")
	}
	if args.Token == nil {
		return nil, errors.New("invalid value for required argument 'Token'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	if args.Vm_id == nil {
		return nil, errors.New("invalid value for required argument 'Vm_id'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VirtualMachineCD
	err := ctx.RegisterResource("vco:resources:VirtualMachineCD", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualMachineCD gets an existing VirtualMachineCD resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualMachineCD(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineCDState, opts ...pulumi.ResourceOption) (*VirtualMachineCD, error) {
	var resource VirtualMachineCD
	err := ctx.ReadResource("vco:resources:VirtualMachineCD", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualMachineCD resources.
type virtualMachineCDState struct {
}

type VirtualMachineCDState struct {
}

func (VirtualMachineCDState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineCDState)(nil)).Elem()
}

type virtualMachineCDArgs struct {
	Cdrom_id      int    `pulumi:"cdrom_id"`
	Cloudspace_id string `pulumi:"cloudspace_id"`
	CustomerID    string `pulumi:"customerID"`
	Token         string `pulumi:"token"`
	Url           string `pulumi:"url"`
	Vm_id         int    `pulumi:"vm_id"`
}

// The set of arguments for constructing a VirtualMachineCD resource.
type VirtualMachineCDArgs struct {
	Cdrom_id      pulumi.IntInput
	Cloudspace_id pulumi.StringInput
	CustomerID    pulumi.StringInput
	Token         pulumi.StringInput
	Url           pulumi.StringInput
	Vm_id         pulumi.IntInput
}

func (VirtualMachineCDArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineCDArgs)(nil)).Elem()
}

type VirtualMachineCDInput interface {
	pulumi.Input

	ToVirtualMachineCDOutput() VirtualMachineCDOutput
	ToVirtualMachineCDOutputWithContext(ctx context.Context) VirtualMachineCDOutput
}

func (*VirtualMachineCD) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineCD)(nil)).Elem()
}

func (i *VirtualMachineCD) ToVirtualMachineCDOutput() VirtualMachineCDOutput {
	return i.ToVirtualMachineCDOutputWithContext(context.Background())
}

func (i *VirtualMachineCD) ToVirtualMachineCDOutputWithContext(ctx context.Context) VirtualMachineCDOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineCDOutput)
}

type VirtualMachineCDOutput struct{ *pulumi.OutputState }

func (VirtualMachineCDOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineCD)(nil)).Elem()
}

func (o VirtualMachineCDOutput) ToVirtualMachineCDOutput() VirtualMachineCDOutput {
	return o
}

func (o VirtualMachineCDOutput) ToVirtualMachineCDOutputWithContext(ctx context.Context) VirtualMachineCDOutput {
	return o
}

func (o VirtualMachineCDOutput) Cdrom_id() pulumi.IntOutput {
	return o.ApplyT(func(v *VirtualMachineCD) pulumi.IntOutput { return v.Cdrom_id }).(pulumi.IntOutput)
}

func (o VirtualMachineCDOutput) Cloudspace_id() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineCD) pulumi.StringOutput { return v.Cloudspace_id }).(pulumi.StringOutput)
}

func (o VirtualMachineCDOutput) CustomerID() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineCD) pulumi.StringOutput { return v.CustomerID }).(pulumi.StringOutput)
}

func (o VirtualMachineCDOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineCD) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o VirtualMachineCDOutput) Disk_size() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineCD) pulumi.StringOutput { return v.Disk_size }).(pulumi.StringOutput)
}

func (o VirtualMachineCDOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineCD) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualMachineCDOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineCD) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o VirtualMachineCDOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineCD) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

func (o VirtualMachineCDOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineCD) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

func (o VirtualMachineCDOutput) Vm_id() pulumi.IntOutput {
	return o.ApplyT(func(v *VirtualMachineCD) pulumi.IntOutput { return v.Vm_id }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualMachineCDInput)(nil)).Elem(), &VirtualMachineCD{})
	pulumi.RegisterOutputType(VirtualMachineCDOutput{})
}