// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package objectspace

import (
	"context"
	"reflect"

	"errors"
	"github.com/fabianv-cloud/pulumi-vco-native/sdk/go/vco/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Link struct {
	pulumi.CustomResourceState

	Cloudspace_id  pulumi.StringOutput `pulumi:"cloudspace_id"`
	CustomerID     pulumi.StringOutput `pulumi:"customerID"`
	Objectspace_id pulumi.StringOutput `pulumi:"objectspace_id"`
	Token          pulumi.StringOutput `pulumi:"token"`
	Url            pulumi.StringOutput `pulumi:"url"`
}

// NewLink registers a new resource with the given unique name, arguments, and options.
func NewLink(ctx *pulumi.Context,
	name string, args *LinkArgs, opts ...pulumi.ResourceOption) (*Link, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Cloudspace_id == nil {
		return nil, errors.New("invalid value for required argument 'Cloudspace_id'")
	}
	if args.CustomerID == nil {
		return nil, errors.New("invalid value for required argument 'CustomerID'")
	}
	if args.Objectspace_id == nil {
		return nil, errors.New("invalid value for required argument 'Objectspace_id'")
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
	var resource Link
	err := ctx.RegisterResource("vco:objectspace:Link", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLink gets an existing Link resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkState, opts ...pulumi.ResourceOption) (*Link, error) {
	var resource Link
	err := ctx.ReadResource("vco:objectspace:Link", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Link resources.
type linkState struct {
}

type LinkState struct {
}

func (LinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkState)(nil)).Elem()
}

type linkArgs struct {
	Cloudspace_id  string `pulumi:"cloudspace_id"`
	CustomerID     string `pulumi:"customerID"`
	Objectspace_id string `pulumi:"objectspace_id"`
	Token          string `pulumi:"token"`
	Url            string `pulumi:"url"`
}

// The set of arguments for constructing a Link resource.
type LinkArgs struct {
	Cloudspace_id  pulumi.StringInput
	CustomerID     pulumi.StringInput
	Objectspace_id pulumi.StringInput
	Token          pulumi.StringInput
	Url            pulumi.StringInput
}

func (LinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkArgs)(nil)).Elem()
}

type LinkInput interface {
	pulumi.Input

	ToLinkOutput() LinkOutput
	ToLinkOutputWithContext(ctx context.Context) LinkOutput
}

func (*Link) ElementType() reflect.Type {
	return reflect.TypeOf((**Link)(nil)).Elem()
}

func (i *Link) ToLinkOutput() LinkOutput {
	return i.ToLinkOutputWithContext(context.Background())
}

func (i *Link) ToLinkOutputWithContext(ctx context.Context) LinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkOutput)
}

// LinkArrayInput is an input type that accepts LinkArray and LinkArrayOutput values.
// You can construct a concrete instance of `LinkArrayInput` via:
//
//	LinkArray{ LinkArgs{...} }
type LinkArrayInput interface {
	pulumi.Input

	ToLinkArrayOutput() LinkArrayOutput
	ToLinkArrayOutputWithContext(context.Context) LinkArrayOutput
}

type LinkArray []LinkInput

func (LinkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Link)(nil)).Elem()
}

func (i LinkArray) ToLinkArrayOutput() LinkArrayOutput {
	return i.ToLinkArrayOutputWithContext(context.Background())
}

func (i LinkArray) ToLinkArrayOutputWithContext(ctx context.Context) LinkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkArrayOutput)
}

// LinkMapInput is an input type that accepts LinkMap and LinkMapOutput values.
// You can construct a concrete instance of `LinkMapInput` via:
//
//	LinkMap{ "key": LinkArgs{...} }
type LinkMapInput interface {
	pulumi.Input

	ToLinkMapOutput() LinkMapOutput
	ToLinkMapOutputWithContext(context.Context) LinkMapOutput
}

type LinkMap map[string]LinkInput

func (LinkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Link)(nil)).Elem()
}

func (i LinkMap) ToLinkMapOutput() LinkMapOutput {
	return i.ToLinkMapOutputWithContext(context.Background())
}

func (i LinkMap) ToLinkMapOutputWithContext(ctx context.Context) LinkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkMapOutput)
}

type LinkOutput struct{ *pulumi.OutputState }

func (LinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Link)(nil)).Elem()
}

func (o LinkOutput) ToLinkOutput() LinkOutput {
	return o
}

func (o LinkOutput) ToLinkOutputWithContext(ctx context.Context) LinkOutput {
	return o
}

func (o LinkOutput) Cloudspace_id() pulumi.StringOutput {
	return o.ApplyT(func(v *Link) pulumi.StringOutput { return v.Cloudspace_id }).(pulumi.StringOutput)
}

func (o LinkOutput) CustomerID() pulumi.StringOutput {
	return o.ApplyT(func(v *Link) pulumi.StringOutput { return v.CustomerID }).(pulumi.StringOutput)
}

func (o LinkOutput) Objectspace_id() pulumi.StringOutput {
	return o.ApplyT(func(v *Link) pulumi.StringOutput { return v.Objectspace_id }).(pulumi.StringOutput)
}

func (o LinkOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *Link) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

func (o LinkOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *Link) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type LinkArrayOutput struct{ *pulumi.OutputState }

func (LinkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Link)(nil)).Elem()
}

func (o LinkArrayOutput) ToLinkArrayOutput() LinkArrayOutput {
	return o
}

func (o LinkArrayOutput) ToLinkArrayOutputWithContext(ctx context.Context) LinkArrayOutput {
	return o
}

func (o LinkArrayOutput) Index(i pulumi.IntInput) LinkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Link {
		return vs[0].([]*Link)[vs[1].(int)]
	}).(LinkOutput)
}

type LinkMapOutput struct{ *pulumi.OutputState }

func (LinkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Link)(nil)).Elem()
}

func (o LinkMapOutput) ToLinkMapOutput() LinkMapOutput {
	return o
}

func (o LinkMapOutput) ToLinkMapOutputWithContext(ctx context.Context) LinkMapOutput {
	return o
}

func (o LinkMapOutput) MapIndex(k pulumi.StringInput) LinkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Link {
		return vs[0].(map[string]*Link)[vs[1].(string)]
	}).(LinkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LinkInput)(nil)).Elem(), &Link{})
	pulumi.RegisterInputType(reflect.TypeOf((*LinkArrayInput)(nil)).Elem(), LinkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LinkMapInput)(nil)).Elem(), LinkMap{})
	pulumi.RegisterOutputType(LinkOutput{})
	pulumi.RegisterOutputType(LinkArrayOutput{})
	pulumi.RegisterOutputType(LinkMapOutput{})
}