// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package base

import (
	"context"
	"reflect"

	"errors"
	"github.com/fabianv-cloud/pulumi-vco-native/sdk/go/vco/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ObjectSpace struct {
	pulumi.CustomResourceState

	Access_key       pulumi.StringPtrOutput `pulumi:"access_key"`
	CloudspaceID     pulumi.StringPtrOutput `pulumi:"cloudspaceID"`
	Creation_time    pulumi.StringOutput    `pulumi:"creation_time"`
	CustomerID       pulumi.StringOutput    `pulumi:"customerID"`
	Domain           pulumi.StringPtrOutput `pulumi:"domain"`
	ExternalNetwork  pulumi.IntPtrOutput    `pulumi:"externalNetwork"`
	Letsencrypt      pulumi.BoolPtrOutput   `pulumi:"letsencrypt"`
	LetsencryptEmail pulumi.StringPtrOutput `pulumi:"letsencryptEmail"`
	Location         pulumi.StringOutput    `pulumi:"location"`
	Objectspace_id   pulumi.StringOutput    `pulumi:"objectspace_id"`
	Objectspace_name pulumi.StringOutput    `pulumi:"objectspace_name"`
	Secret           pulumi.StringPtrOutput `pulumi:"secret"`
	Status           pulumi.StringOutput    `pulumi:"status"`
	Subnet           pulumi.StringPtrOutput `pulumi:"subnet"`
	Token            pulumi.StringOutput    `pulumi:"token"`
	Update_time      pulumi.StringOutput    `pulumi:"update_time"`
	Url              pulumi.StringOutput    `pulumi:"url"`
}

// NewObjectSpace registers a new resource with the given unique name, arguments, and options.
func NewObjectSpace(ctx *pulumi.Context,
	name string, args *ObjectSpaceArgs, opts ...pulumi.ResourceOption) (*ObjectSpace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CustomerID == nil {
		return nil, errors.New("invalid value for required argument 'CustomerID'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Objectspace_name == nil {
		return nil, errors.New("invalid value for required argument 'Objectspace_name'")
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
	var resource ObjectSpace
	err := ctx.RegisterResource("vco:base:ObjectSpace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetObjectSpace gets an existing ObjectSpace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetObjectSpace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ObjectSpaceState, opts ...pulumi.ResourceOption) (*ObjectSpace, error) {
	var resource ObjectSpace
	err := ctx.ReadResource("vco:base:ObjectSpace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ObjectSpace resources.
type objectSpaceState struct {
}

type ObjectSpaceState struct {
}

func (ObjectSpaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*objectSpaceState)(nil)).Elem()
}

type objectSpaceArgs struct {
	CloudspaceID     *string `pulumi:"cloudspaceID"`
	CustomerID       string  `pulumi:"customerID"`
	Domain           *string `pulumi:"domain"`
	ExternalNetwork  *int    `pulumi:"externalNetwork"`
	Letsencrypt      *bool   `pulumi:"letsencrypt"`
	LetsencryptEmail *string `pulumi:"letsencryptEmail"`
	Location         string  `pulumi:"location"`
	Objectspace_name string  `pulumi:"objectspace_name"`
	Subnet           *string `pulumi:"subnet"`
	Token            string  `pulumi:"token"`
	Url              string  `pulumi:"url"`
}

// The set of arguments for constructing a ObjectSpace resource.
type ObjectSpaceArgs struct {
	CloudspaceID     pulumi.StringPtrInput
	CustomerID       pulumi.StringInput
	Domain           pulumi.StringPtrInput
	ExternalNetwork  pulumi.IntPtrInput
	Letsencrypt      pulumi.BoolPtrInput
	LetsencryptEmail pulumi.StringPtrInput
	Location         pulumi.StringInput
	Objectspace_name pulumi.StringInput
	Subnet           pulumi.StringPtrInput
	Token            pulumi.StringInput
	Url              pulumi.StringInput
}

func (ObjectSpaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*objectSpaceArgs)(nil)).Elem()
}

type ObjectSpaceInput interface {
	pulumi.Input

	ToObjectSpaceOutput() ObjectSpaceOutput
	ToObjectSpaceOutputWithContext(ctx context.Context) ObjectSpaceOutput
}

func (*ObjectSpace) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectSpace)(nil)).Elem()
}

func (i *ObjectSpace) ToObjectSpaceOutput() ObjectSpaceOutput {
	return i.ToObjectSpaceOutputWithContext(context.Background())
}

func (i *ObjectSpace) ToObjectSpaceOutputWithContext(ctx context.Context) ObjectSpaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectSpaceOutput)
}

// ObjectSpaceArrayInput is an input type that accepts ObjectSpaceArray and ObjectSpaceArrayOutput values.
// You can construct a concrete instance of `ObjectSpaceArrayInput` via:
//
//	ObjectSpaceArray{ ObjectSpaceArgs{...} }
type ObjectSpaceArrayInput interface {
	pulumi.Input

	ToObjectSpaceArrayOutput() ObjectSpaceArrayOutput
	ToObjectSpaceArrayOutputWithContext(context.Context) ObjectSpaceArrayOutput
}

type ObjectSpaceArray []ObjectSpaceInput

func (ObjectSpaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ObjectSpace)(nil)).Elem()
}

func (i ObjectSpaceArray) ToObjectSpaceArrayOutput() ObjectSpaceArrayOutput {
	return i.ToObjectSpaceArrayOutputWithContext(context.Background())
}

func (i ObjectSpaceArray) ToObjectSpaceArrayOutputWithContext(ctx context.Context) ObjectSpaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectSpaceArrayOutput)
}

// ObjectSpaceMapInput is an input type that accepts ObjectSpaceMap and ObjectSpaceMapOutput values.
// You can construct a concrete instance of `ObjectSpaceMapInput` via:
//
//	ObjectSpaceMap{ "key": ObjectSpaceArgs{...} }
type ObjectSpaceMapInput interface {
	pulumi.Input

	ToObjectSpaceMapOutput() ObjectSpaceMapOutput
	ToObjectSpaceMapOutputWithContext(context.Context) ObjectSpaceMapOutput
}

type ObjectSpaceMap map[string]ObjectSpaceInput

func (ObjectSpaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ObjectSpace)(nil)).Elem()
}

func (i ObjectSpaceMap) ToObjectSpaceMapOutput() ObjectSpaceMapOutput {
	return i.ToObjectSpaceMapOutputWithContext(context.Background())
}

func (i ObjectSpaceMap) ToObjectSpaceMapOutputWithContext(ctx context.Context) ObjectSpaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectSpaceMapOutput)
}

type ObjectSpaceOutput struct{ *pulumi.OutputState }

func (ObjectSpaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectSpace)(nil)).Elem()
}

func (o ObjectSpaceOutput) ToObjectSpaceOutput() ObjectSpaceOutput {
	return o
}

func (o ObjectSpaceOutput) ToObjectSpaceOutputWithContext(ctx context.Context) ObjectSpaceOutput {
	return o
}

func (o ObjectSpaceOutput) Access_key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ObjectSpace) pulumi.StringPtrOutput { return v.Access_key }).(pulumi.StringPtrOutput)
}

func (o ObjectSpaceOutput) CloudspaceID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ObjectSpace) pulumi.StringPtrOutput { return v.CloudspaceID }).(pulumi.StringPtrOutput)
}

func (o ObjectSpaceOutput) Creation_time() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectSpace) pulumi.StringOutput { return v.Creation_time }).(pulumi.StringOutput)
}

func (o ObjectSpaceOutput) CustomerID() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectSpace) pulumi.StringOutput { return v.CustomerID }).(pulumi.StringOutput)
}

func (o ObjectSpaceOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ObjectSpace) pulumi.StringPtrOutput { return v.Domain }).(pulumi.StringPtrOutput)
}

func (o ObjectSpaceOutput) ExternalNetwork() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ObjectSpace) pulumi.IntPtrOutput { return v.ExternalNetwork }).(pulumi.IntPtrOutput)
}

func (o ObjectSpaceOutput) Letsencrypt() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ObjectSpace) pulumi.BoolPtrOutput { return v.Letsencrypt }).(pulumi.BoolPtrOutput)
}

func (o ObjectSpaceOutput) LetsencryptEmail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ObjectSpace) pulumi.StringPtrOutput { return v.LetsencryptEmail }).(pulumi.StringPtrOutput)
}

func (o ObjectSpaceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectSpace) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ObjectSpaceOutput) Objectspace_id() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectSpace) pulumi.StringOutput { return v.Objectspace_id }).(pulumi.StringOutput)
}

func (o ObjectSpaceOutput) Objectspace_name() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectSpace) pulumi.StringOutput { return v.Objectspace_name }).(pulumi.StringOutput)
}

func (o ObjectSpaceOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ObjectSpace) pulumi.StringPtrOutput { return v.Secret }).(pulumi.StringPtrOutput)
}

func (o ObjectSpaceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectSpace) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o ObjectSpaceOutput) Subnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ObjectSpace) pulumi.StringPtrOutput { return v.Subnet }).(pulumi.StringPtrOutput)
}

func (o ObjectSpaceOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectSpace) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

func (o ObjectSpaceOutput) Update_time() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectSpace) pulumi.StringOutput { return v.Update_time }).(pulumi.StringOutput)
}

func (o ObjectSpaceOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectSpace) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type ObjectSpaceArrayOutput struct{ *pulumi.OutputState }

func (ObjectSpaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ObjectSpace)(nil)).Elem()
}

func (o ObjectSpaceArrayOutput) ToObjectSpaceArrayOutput() ObjectSpaceArrayOutput {
	return o
}

func (o ObjectSpaceArrayOutput) ToObjectSpaceArrayOutputWithContext(ctx context.Context) ObjectSpaceArrayOutput {
	return o
}

func (o ObjectSpaceArrayOutput) Index(i pulumi.IntInput) ObjectSpaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ObjectSpace {
		return vs[0].([]*ObjectSpace)[vs[1].(int)]
	}).(ObjectSpaceOutput)
}

type ObjectSpaceMapOutput struct{ *pulumi.OutputState }

func (ObjectSpaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ObjectSpace)(nil)).Elem()
}

func (o ObjectSpaceMapOutput) ToObjectSpaceMapOutput() ObjectSpaceMapOutput {
	return o
}

func (o ObjectSpaceMapOutput) ToObjectSpaceMapOutputWithContext(ctx context.Context) ObjectSpaceMapOutput {
	return o
}

func (o ObjectSpaceMapOutput) MapIndex(k pulumi.StringInput) ObjectSpaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ObjectSpace {
		return vs[0].(map[string]*ObjectSpace)[vs[1].(string)]
	}).(ObjectSpaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectSpaceInput)(nil)).Elem(), &ObjectSpace{})
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectSpaceArrayInput)(nil)).Elem(), ObjectSpaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectSpaceMapInput)(nil)).Elem(), ObjectSpaceMap{})
	pulumi.RegisterOutputType(ObjectSpaceOutput{})
	pulumi.RegisterOutputType(ObjectSpaceArrayOutput{})
	pulumi.RegisterOutputType(ObjectSpaceMapOutput{})
}
