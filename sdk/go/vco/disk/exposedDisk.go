// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package disk

import (
	"context"
	"reflect"

	"errors"
	"github.com/fabianv-cloud/pulumi-vco-native/sdk/go/vco/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ExposedDisk struct {
	pulumi.CustomResourceState

	Cloudspace_id   pulumi.StringOutput `pulumi:"cloudspace_id"`
	CustomerID      pulumi.StringOutput `pulumi:"customerID"`
	Disk_id         pulumi.IntOutput    `pulumi:"disk_id"`
	Endpoint        EndpointOutput      `pulumi:"endpoint"`
	Iops            pulumi.IntPtrOutput `pulumi:"iops"`
	Max_connections pulumi.IntPtrOutput `pulumi:"max_connections"`
	Protocol        pulumi.StringOutput `pulumi:"protocol"`
	Token           pulumi.StringOutput `pulumi:"token"`
	Url             pulumi.StringOutput `pulumi:"url"`
}

// NewExposedDisk registers a new resource with the given unique name, arguments, and options.
func NewExposedDisk(ctx *pulumi.Context,
	name string, args *ExposedDiskArgs, opts ...pulumi.ResourceOption) (*ExposedDisk, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Cloudspace_id == nil {
		return nil, errors.New("invalid value for required argument 'Cloudspace_id'")
	}
	if args.CustomerID == nil {
		return nil, errors.New("invalid value for required argument 'CustomerID'")
	}
	if args.Disk_id == nil {
		return nil, errors.New("invalid value for required argument 'Disk_id'")
	}
	if args.Token == nil {
		return nil, errors.New("invalid value for required argument 'Token'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	if args.CustomerID != nil {
		args.CustomerID = pulumi.ToSecret(args.CustomerID).(pulumi.StringInput)
	}
	if args.Token != nil {
		args.Token = pulumi.ToSecret(args.Token).(pulumi.StringInput)
	}
	if args.Url != nil {
		args.Url = pulumi.ToSecret(args.Url).(pulumi.StringInput)
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ExposedDisk
	err := ctx.RegisterResource("vco:disk:ExposedDisk", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExposedDisk gets an existing ExposedDisk resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExposedDisk(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExposedDiskState, opts ...pulumi.ResourceOption) (*ExposedDisk, error) {
	var resource ExposedDisk
	err := ctx.ReadResource("vco:disk:ExposedDisk", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExposedDisk resources.
type exposedDiskState struct {
}

type ExposedDiskState struct {
}

func (ExposedDiskState) ElementType() reflect.Type {
	return reflect.TypeOf((*exposedDiskState)(nil)).Elem()
}

type exposedDiskArgs struct {
	Cloudspace_id   string `pulumi:"cloudspace_id"`
	CustomerID      string `pulumi:"customerID"`
	Disk_id         int    `pulumi:"disk_id"`
	Iops            *int   `pulumi:"iops"`
	Max_connections *int   `pulumi:"max_connections"`
	Token           string `pulumi:"token"`
	Url             string `pulumi:"url"`
}

// The set of arguments for constructing a ExposedDisk resource.
type ExposedDiskArgs struct {
	Cloudspace_id   pulumi.StringInput
	CustomerID      pulumi.StringInput
	Disk_id         pulumi.IntInput
	Iops            pulumi.IntPtrInput
	Max_connections pulumi.IntPtrInput
	Token           pulumi.StringInput
	Url             pulumi.StringInput
}

func (ExposedDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*exposedDiskArgs)(nil)).Elem()
}

type ExposedDiskInput interface {
	pulumi.Input

	ToExposedDiskOutput() ExposedDiskOutput
	ToExposedDiskOutputWithContext(ctx context.Context) ExposedDiskOutput
}

func (*ExposedDisk) ElementType() reflect.Type {
	return reflect.TypeOf((**ExposedDisk)(nil)).Elem()
}

func (i *ExposedDisk) ToExposedDiskOutput() ExposedDiskOutput {
	return i.ToExposedDiskOutputWithContext(context.Background())
}

func (i *ExposedDisk) ToExposedDiskOutputWithContext(ctx context.Context) ExposedDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExposedDiskOutput)
}

// ExposedDiskArrayInput is an input type that accepts ExposedDiskArray and ExposedDiskArrayOutput values.
// You can construct a concrete instance of `ExposedDiskArrayInput` via:
//
//	ExposedDiskArray{ ExposedDiskArgs{...} }
type ExposedDiskArrayInput interface {
	pulumi.Input

	ToExposedDiskArrayOutput() ExposedDiskArrayOutput
	ToExposedDiskArrayOutputWithContext(context.Context) ExposedDiskArrayOutput
}

type ExposedDiskArray []ExposedDiskInput

func (ExposedDiskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExposedDisk)(nil)).Elem()
}

func (i ExposedDiskArray) ToExposedDiskArrayOutput() ExposedDiskArrayOutput {
	return i.ToExposedDiskArrayOutputWithContext(context.Background())
}

func (i ExposedDiskArray) ToExposedDiskArrayOutputWithContext(ctx context.Context) ExposedDiskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExposedDiskArrayOutput)
}

// ExposedDiskMapInput is an input type that accepts ExposedDiskMap and ExposedDiskMapOutput values.
// You can construct a concrete instance of `ExposedDiskMapInput` via:
//
//	ExposedDiskMap{ "key": ExposedDiskArgs{...} }
type ExposedDiskMapInput interface {
	pulumi.Input

	ToExposedDiskMapOutput() ExposedDiskMapOutput
	ToExposedDiskMapOutputWithContext(context.Context) ExposedDiskMapOutput
}

type ExposedDiskMap map[string]ExposedDiskInput

func (ExposedDiskMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExposedDisk)(nil)).Elem()
}

func (i ExposedDiskMap) ToExposedDiskMapOutput() ExposedDiskMapOutput {
	return i.ToExposedDiskMapOutputWithContext(context.Background())
}

func (i ExposedDiskMap) ToExposedDiskMapOutputWithContext(ctx context.Context) ExposedDiskMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExposedDiskMapOutput)
}

type ExposedDiskOutput struct{ *pulumi.OutputState }

func (ExposedDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExposedDisk)(nil)).Elem()
}

func (o ExposedDiskOutput) ToExposedDiskOutput() ExposedDiskOutput {
	return o
}

func (o ExposedDiskOutput) ToExposedDiskOutputWithContext(ctx context.Context) ExposedDiskOutput {
	return o
}

func (o ExposedDiskOutput) Cloudspace_id() pulumi.StringOutput {
	return o.ApplyT(func(v *ExposedDisk) pulumi.StringOutput { return v.Cloudspace_id }).(pulumi.StringOutput)
}

func (o ExposedDiskOutput) CustomerID() pulumi.StringOutput {
	return o.ApplyT(func(v *ExposedDisk) pulumi.StringOutput { return v.CustomerID }).(pulumi.StringOutput)
}

func (o ExposedDiskOutput) Disk_id() pulumi.IntOutput {
	return o.ApplyT(func(v *ExposedDisk) pulumi.IntOutput { return v.Disk_id }).(pulumi.IntOutput)
}

func (o ExposedDiskOutput) Endpoint() EndpointOutput {
	return o.ApplyT(func(v *ExposedDisk) EndpointOutput { return v.Endpoint }).(EndpointOutput)
}

func (o ExposedDiskOutput) Iops() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ExposedDisk) pulumi.IntPtrOutput { return v.Iops }).(pulumi.IntPtrOutput)
}

func (o ExposedDiskOutput) Max_connections() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ExposedDisk) pulumi.IntPtrOutput { return v.Max_connections }).(pulumi.IntPtrOutput)
}

func (o ExposedDiskOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *ExposedDisk) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

func (o ExposedDiskOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *ExposedDisk) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

func (o ExposedDiskOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *ExposedDisk) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type ExposedDiskArrayOutput struct{ *pulumi.OutputState }

func (ExposedDiskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExposedDisk)(nil)).Elem()
}

func (o ExposedDiskArrayOutput) ToExposedDiskArrayOutput() ExposedDiskArrayOutput {
	return o
}

func (o ExposedDiskArrayOutput) ToExposedDiskArrayOutputWithContext(ctx context.Context) ExposedDiskArrayOutput {
	return o
}

func (o ExposedDiskArrayOutput) Index(i pulumi.IntInput) ExposedDiskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ExposedDisk {
		return vs[0].([]*ExposedDisk)[vs[1].(int)]
	}).(ExposedDiskOutput)
}

type ExposedDiskMapOutput struct{ *pulumi.OutputState }

func (ExposedDiskMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExposedDisk)(nil)).Elem()
}

func (o ExposedDiskMapOutput) ToExposedDiskMapOutput() ExposedDiskMapOutput {
	return o
}

func (o ExposedDiskMapOutput) ToExposedDiskMapOutputWithContext(ctx context.Context) ExposedDiskMapOutput {
	return o
}

func (o ExposedDiskMapOutput) MapIndex(k pulumi.StringInput) ExposedDiskOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ExposedDisk {
		return vs[0].(map[string]*ExposedDisk)[vs[1].(string)]
	}).(ExposedDiskOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExposedDiskInput)(nil)).Elem(), &ExposedDisk{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExposedDiskArrayInput)(nil)).Elem(), ExposedDiskArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExposedDiskMapInput)(nil)).Elem(), ExposedDiskMap{})
	pulumi.RegisterOutputType(ExposedDiskOutput{})
	pulumi.RegisterOutputType(ExposedDiskArrayOutput{})
	pulumi.RegisterOutputType(ExposedDiskMapOutput{})
}