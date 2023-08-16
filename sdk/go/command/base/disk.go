// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package base

import (
	"context"
	"reflect"

	"errors"
	"github.com/fabianv-cloud/pulumi-vco-native/sdk/go/command/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Disk struct {
	pulumi.CustomResourceState

	Cloudspace_id pulumi.StringOutput  `pulumi:"cloudspace_id"`
	CustomerID    pulumi.StringOutput  `pulumi:"customerID"`
	Description   pulumi.StringOutput  `pulumi:"description"`
	Detach        pulumi.BoolPtrOutput `pulumi:"detach"`
	Disk_id       pulumi.IntOutput     `pulumi:"disk_id"`
	Disk_name     pulumi.StringOutput  `pulumi:"disk_name"`
	Disk_size     pulumi.IntOutput     `pulumi:"disk_size"`
	Disk_type     pulumi.StringOutput  `pulumi:"disk_type"`
	Exposed       pulumi.BoolOutput    `pulumi:"exposed"`
	Iops          pulumi.IntPtrOutput  `pulumi:"iops"`
	Location      pulumi.StringOutput  `pulumi:"location"`
	Model         pulumi.StringOutput  `pulumi:"model"`
	Order         pulumi.StringOutput  `pulumi:"order"`
	Permanently   pulumi.BoolPtrOutput `pulumi:"permanently"`
	Port          pulumi.IntOutput     `pulumi:"port"`
	Status        pulumi.StringOutput  `pulumi:"status"`
	Token         pulumi.StringOutput  `pulumi:"token"`
	Url           pulumi.StringOutput  `pulumi:"url"`
	Vm_id         pulumi.IntOutput     `pulumi:"vm_id"`
}

// NewDisk registers a new resource with the given unique name, arguments, and options.
func NewDisk(ctx *pulumi.Context,
	name string, args *DiskArgs, opts ...pulumi.ResourceOption) (*Disk, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CustomerID == nil {
		return nil, errors.New("invalid value for required argument 'CustomerID'")
	}
	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.Disk_name == nil {
		return nil, errors.New("invalid value for required argument 'Disk_name'")
	}
	if args.Disk_size == nil {
		return nil, errors.New("invalid value for required argument 'Disk_size'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Token == nil {
		return nil, errors.New("invalid value for required argument 'Token'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Disk
	err := ctx.RegisterResource("vco:base:Disk", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDisk gets an existing Disk resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDisk(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiskState, opts ...pulumi.ResourceOption) (*Disk, error) {
	var resource Disk
	err := ctx.ReadResource("vco:base:Disk", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Disk resources.
type diskState struct {
}

type DiskState struct {
}

func (DiskState) ElementType() reflect.Type {
	return reflect.TypeOf((*diskState)(nil)).Elem()
}

type diskArgs struct {
	CustomerID  string  `pulumi:"customerID"`
	Description string  `pulumi:"description"`
	Detach      *bool   `pulumi:"detach"`
	Disk_name   string  `pulumi:"disk_name"`
	Disk_size   int     `pulumi:"disk_size"`
	Disk_type   *string `pulumi:"disk_type"`
	Iops        *int    `pulumi:"iops"`
	Location    string  `pulumi:"location"`
	Permanently *bool   `pulumi:"permanently"`
	Token       string  `pulumi:"token"`
	Url         string  `pulumi:"url"`
	Vm_id       *string `pulumi:"vm_id"`
}

// The set of arguments for constructing a Disk resource.
type DiskArgs struct {
	CustomerID  pulumi.StringInput
	Description pulumi.StringInput
	Detach      pulumi.BoolPtrInput
	Disk_name   pulumi.StringInput
	Disk_size   pulumi.IntInput
	Disk_type   pulumi.StringPtrInput
	Iops        pulumi.IntPtrInput
	Location    pulumi.StringInput
	Permanently pulumi.BoolPtrInput
	Token       pulumi.StringInput
	Url         pulumi.StringInput
	Vm_id       pulumi.StringPtrInput
}

func (DiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*diskArgs)(nil)).Elem()
}

type DiskInput interface {
	pulumi.Input

	ToDiskOutput() DiskOutput
	ToDiskOutputWithContext(ctx context.Context) DiskOutput
}

func (*Disk) ElementType() reflect.Type {
	return reflect.TypeOf((**Disk)(nil)).Elem()
}

func (i *Disk) ToDiskOutput() DiskOutput {
	return i.ToDiskOutputWithContext(context.Background())
}

func (i *Disk) ToDiskOutputWithContext(ctx context.Context) DiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskOutput)
}

// DiskArrayInput is an input type that accepts DiskArray and DiskArrayOutput values.
// You can construct a concrete instance of `DiskArrayInput` via:
//
//	DiskArray{ DiskArgs{...} }
type DiskArrayInput interface {
	pulumi.Input

	ToDiskArrayOutput() DiskArrayOutput
	ToDiskArrayOutputWithContext(context.Context) DiskArrayOutput
}

type DiskArray []DiskInput

func (DiskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Disk)(nil)).Elem()
}

func (i DiskArray) ToDiskArrayOutput() DiskArrayOutput {
	return i.ToDiskArrayOutputWithContext(context.Background())
}

func (i DiskArray) ToDiskArrayOutputWithContext(ctx context.Context) DiskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskArrayOutput)
}

// DiskMapInput is an input type that accepts DiskMap and DiskMapOutput values.
// You can construct a concrete instance of `DiskMapInput` via:
//
//	DiskMap{ "key": DiskArgs{...} }
type DiskMapInput interface {
	pulumi.Input

	ToDiskMapOutput() DiskMapOutput
	ToDiskMapOutputWithContext(context.Context) DiskMapOutput
}

type DiskMap map[string]DiskInput

func (DiskMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Disk)(nil)).Elem()
}

func (i DiskMap) ToDiskMapOutput() DiskMapOutput {
	return i.ToDiskMapOutputWithContext(context.Background())
}

func (i DiskMap) ToDiskMapOutputWithContext(ctx context.Context) DiskMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskMapOutput)
}

type DiskOutput struct{ *pulumi.OutputState }

func (DiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Disk)(nil)).Elem()
}

func (o DiskOutput) ToDiskOutput() DiskOutput {
	return o
}

func (o DiskOutput) ToDiskOutputWithContext(ctx context.Context) DiskOutput {
	return o
}

func (o DiskOutput) Cloudspace_id() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.Cloudspace_id }).(pulumi.StringOutput)
}

func (o DiskOutput) CustomerID() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.CustomerID }).(pulumi.StringOutput)
}

func (o DiskOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o DiskOutput) Detach() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.BoolPtrOutput { return v.Detach }).(pulumi.BoolPtrOutput)
}

func (o DiskOutput) Disk_id() pulumi.IntOutput {
	return o.ApplyT(func(v *Disk) pulumi.IntOutput { return v.Disk_id }).(pulumi.IntOutput)
}

func (o DiskOutput) Disk_name() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.Disk_name }).(pulumi.StringOutput)
}

func (o DiskOutput) Disk_size() pulumi.IntOutput {
	return o.ApplyT(func(v *Disk) pulumi.IntOutput { return v.Disk_size }).(pulumi.IntOutput)
}

func (o DiskOutput) Disk_type() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.Disk_type }).(pulumi.StringOutput)
}

func (o DiskOutput) Exposed() pulumi.BoolOutput {
	return o.ApplyT(func(v *Disk) pulumi.BoolOutput { return v.Exposed }).(pulumi.BoolOutput)
}

func (o DiskOutput) Iops() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.IntPtrOutput { return v.Iops }).(pulumi.IntPtrOutput)
}

func (o DiskOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o DiskOutput) Model() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.Model }).(pulumi.StringOutput)
}

func (o DiskOutput) Order() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.Order }).(pulumi.StringOutput)
}

func (o DiskOutput) Permanently() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.BoolPtrOutput { return v.Permanently }).(pulumi.BoolPtrOutput)
}

func (o DiskOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *Disk) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

func (o DiskOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o DiskOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

func (o DiskOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

func (o DiskOutput) Vm_id() pulumi.IntOutput {
	return o.ApplyT(func(v *Disk) pulumi.IntOutput { return v.Vm_id }).(pulumi.IntOutput)
}

type DiskArrayOutput struct{ *pulumi.OutputState }

func (DiskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Disk)(nil)).Elem()
}

func (o DiskArrayOutput) ToDiskArrayOutput() DiskArrayOutput {
	return o
}

func (o DiskArrayOutput) ToDiskArrayOutputWithContext(ctx context.Context) DiskArrayOutput {
	return o
}

func (o DiskArrayOutput) Index(i pulumi.IntInput) DiskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Disk {
		return vs[0].([]*Disk)[vs[1].(int)]
	}).(DiskOutput)
}

type DiskMapOutput struct{ *pulumi.OutputState }

func (DiskMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Disk)(nil)).Elem()
}

func (o DiskMapOutput) ToDiskMapOutput() DiskMapOutput {
	return o
}

func (o DiskMapOutput) ToDiskMapOutputWithContext(ctx context.Context) DiskMapOutput {
	return o
}

func (o DiskMapOutput) MapIndex(k pulumi.StringInput) DiskOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Disk {
		return vs[0].(map[string]*Disk)[vs[1].(string)]
	}).(DiskOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DiskInput)(nil)).Elem(), &Disk{})
	pulumi.RegisterInputType(reflect.TypeOf((*DiskArrayInput)(nil)).Elem(), DiskArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DiskMapInput)(nil)).Elem(), DiskMap{})
	pulumi.RegisterOutputType(DiskOutput{})
	pulumi.RegisterOutputType(DiskArrayOutput{})
	pulumi.RegisterOutputType(DiskMapOutput{})
}
