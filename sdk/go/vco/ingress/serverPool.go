// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ingress

import (
	"context"
	"reflect"

	"errors"
	"github.com/fabianv-cloud/pulumi-vco-native/sdk/go/vco/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ServerPool struct {
	pulumi.CustomResourceState

	Cloudspace_id pulumi.StringOutput       `pulumi:"cloudspace_id"`
	CustomerID    pulumi.StringOutput       `pulumi:"customerID"`
	Description   pulumi.StringOutput       `pulumi:"description"`
	Hosts         ServerPoolHostArrayOutput `pulumi:"hosts"`
	Name          pulumi.StringOutput       `pulumi:"name"`
	Serverpool_id pulumi.StringOutput       `pulumi:"serverpool_id"`
	Token         pulumi.StringOutput       `pulumi:"token"`
	Url           pulumi.StringOutput       `pulumi:"url"`
}

// NewServerPool registers a new resource with the given unique name, arguments, and options.
func NewServerPool(ctx *pulumi.Context,
	name string, args *ServerPoolArgs, opts ...pulumi.ResourceOption) (*ServerPool, error) {
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
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
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
	var resource ServerPool
	err := ctx.RegisterResource("vco:ingress:ServerPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerPool gets an existing ServerPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerPoolState, opts ...pulumi.ResourceOption) (*ServerPool, error) {
	var resource ServerPool
	err := ctx.ReadResource("vco:ingress:ServerPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerPool resources.
type serverPoolState struct {
}

type ServerPoolState struct {
}

func (ServerPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverPoolState)(nil)).Elem()
}

type serverPoolArgs struct {
	Cloudspace_id string `pulumi:"cloudspace_id"`
	CustomerID    string `pulumi:"customerID"`
	Description   string `pulumi:"description"`
	Name          string `pulumi:"name"`
	Token         string `pulumi:"token"`
	Url           string `pulumi:"url"`
}

// The set of arguments for constructing a ServerPool resource.
type ServerPoolArgs struct {
	Cloudspace_id pulumi.StringInput
	CustomerID    pulumi.StringInput
	Description   pulumi.StringInput
	Name          pulumi.StringInput
	Token         pulumi.StringInput
	Url           pulumi.StringInput
}

func (ServerPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverPoolArgs)(nil)).Elem()
}

type ServerPoolInput interface {
	pulumi.Input

	ToServerPoolOutput() ServerPoolOutput
	ToServerPoolOutputWithContext(ctx context.Context) ServerPoolOutput
}

func (*ServerPool) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerPool)(nil)).Elem()
}

func (i *ServerPool) ToServerPoolOutput() ServerPoolOutput {
	return i.ToServerPoolOutputWithContext(context.Background())
}

func (i *ServerPool) ToServerPoolOutputWithContext(ctx context.Context) ServerPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerPoolOutput)
}

// ServerPoolArrayInput is an input type that accepts ServerPoolArray and ServerPoolArrayOutput values.
// You can construct a concrete instance of `ServerPoolArrayInput` via:
//
//	ServerPoolArray{ ServerPoolArgs{...} }
type ServerPoolArrayInput interface {
	pulumi.Input

	ToServerPoolArrayOutput() ServerPoolArrayOutput
	ToServerPoolArrayOutputWithContext(context.Context) ServerPoolArrayOutput
}

type ServerPoolArray []ServerPoolInput

func (ServerPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerPool)(nil)).Elem()
}

func (i ServerPoolArray) ToServerPoolArrayOutput() ServerPoolArrayOutput {
	return i.ToServerPoolArrayOutputWithContext(context.Background())
}

func (i ServerPoolArray) ToServerPoolArrayOutputWithContext(ctx context.Context) ServerPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerPoolArrayOutput)
}

// ServerPoolMapInput is an input type that accepts ServerPoolMap and ServerPoolMapOutput values.
// You can construct a concrete instance of `ServerPoolMapInput` via:
//
//	ServerPoolMap{ "key": ServerPoolArgs{...} }
type ServerPoolMapInput interface {
	pulumi.Input

	ToServerPoolMapOutput() ServerPoolMapOutput
	ToServerPoolMapOutputWithContext(context.Context) ServerPoolMapOutput
}

type ServerPoolMap map[string]ServerPoolInput

func (ServerPoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerPool)(nil)).Elem()
}

func (i ServerPoolMap) ToServerPoolMapOutput() ServerPoolMapOutput {
	return i.ToServerPoolMapOutputWithContext(context.Background())
}

func (i ServerPoolMap) ToServerPoolMapOutputWithContext(ctx context.Context) ServerPoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerPoolMapOutput)
}

type ServerPoolOutput struct{ *pulumi.OutputState }

func (ServerPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerPool)(nil)).Elem()
}

func (o ServerPoolOutput) ToServerPoolOutput() ServerPoolOutput {
	return o
}

func (o ServerPoolOutput) ToServerPoolOutputWithContext(ctx context.Context) ServerPoolOutput {
	return o
}

func (o ServerPoolOutput) Cloudspace_id() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerPool) pulumi.StringOutput { return v.Cloudspace_id }).(pulumi.StringOutput)
}

func (o ServerPoolOutput) CustomerID() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerPool) pulumi.StringOutput { return v.CustomerID }).(pulumi.StringOutput)
}

func (o ServerPoolOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerPool) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o ServerPoolOutput) Hosts() ServerPoolHostArrayOutput {
	return o.ApplyT(func(v *ServerPool) ServerPoolHostArrayOutput { return v.Hosts }).(ServerPoolHostArrayOutput)
}

func (o ServerPoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerPool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ServerPoolOutput) Serverpool_id() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerPool) pulumi.StringOutput { return v.Serverpool_id }).(pulumi.StringOutput)
}

func (o ServerPoolOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerPool) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

func (o ServerPoolOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerPool) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type ServerPoolArrayOutput struct{ *pulumi.OutputState }

func (ServerPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerPool)(nil)).Elem()
}

func (o ServerPoolArrayOutput) ToServerPoolArrayOutput() ServerPoolArrayOutput {
	return o
}

func (o ServerPoolArrayOutput) ToServerPoolArrayOutputWithContext(ctx context.Context) ServerPoolArrayOutput {
	return o
}

func (o ServerPoolArrayOutput) Index(i pulumi.IntInput) ServerPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServerPool {
		return vs[0].([]*ServerPool)[vs[1].(int)]
	}).(ServerPoolOutput)
}

type ServerPoolMapOutput struct{ *pulumi.OutputState }

func (ServerPoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerPool)(nil)).Elem()
}

func (o ServerPoolMapOutput) ToServerPoolMapOutput() ServerPoolMapOutput {
	return o
}

func (o ServerPoolMapOutput) ToServerPoolMapOutputWithContext(ctx context.Context) ServerPoolMapOutput {
	return o
}

func (o ServerPoolMapOutput) MapIndex(k pulumi.StringInput) ServerPoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServerPool {
		return vs[0].(map[string]*ServerPool)[vs[1].(string)]
	}).(ServerPoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServerPoolInput)(nil)).Elem(), &ServerPool{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerPoolArrayInput)(nil)).Elem(), ServerPoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerPoolMapInput)(nil)).Elem(), ServerPoolMap{})
	pulumi.RegisterOutputType(ServerPoolOutput{})
	pulumi.RegisterOutputType(ServerPoolArrayOutput{})
	pulumi.RegisterOutputType(ServerPoolMapOutput{})
}
