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

type ReverseProxy struct {
	pulumi.CustomResourceState

	Back_end            ReverseProxyBackendPtrOutput  `pulumi:"back_end"`
	Cloudspace_id       pulumi.StringOutput           `pulumi:"cloudspace_id"`
	CustomerID          pulumi.StringOutput           `pulumi:"customerID"`
	Description         pulumi.StringPtrOutput        `pulumi:"description"`
	Domain              pulumi.StringOutput           `pulumi:"domain"`
	Email               pulumi.StringPtrOutput        `pulumi:"email"`
	Enabled             pulumi.BoolOutput             `pulumi:"enabled"`
	Front_end           ReverseProxyFrontEndPtrOutput `pulumi:"front_end"`
	Health_check_scheme pulumi.StringPtrOutput        `pulumi:"health_check_scheme"`
	Http_only           pulumi.BoolPtrOutput          `pulumi:"http_only"`
	Http_port           pulumi.IntPtrOutput           `pulumi:"http_port"`
	Https_port          pulumi.IntPtrOutput           `pulumi:"https_port"`
	Interval            pulumi.IntPtrOutput           `pulumi:"interval"`
	Ip_address          pulumi.StringPtrOutput        `pulumi:"ip_address"`
	Name                pulumi.StringOutput           `pulumi:"name"`
	Path                pulumi.StringPtrOutput        `pulumi:"path"`
	Port                pulumi.IntPtrOutput           `pulumi:"port"`
	Reverseproxy_id     pulumi.StringOutput           `pulumi:"reverseproxy_id"`
	Same_site           pulumi.StringPtrOutput        `pulumi:"same_site"`
	Scheme              pulumi.StringOutput           `pulumi:"scheme"`
	Secure              pulumi.BoolPtrOutput          `pulumi:"secure"`
	Serverpool_id       pulumi.StringOutput           `pulumi:"serverpool_id"`
	StickySession_name  pulumi.StringPtrOutput        `pulumi:"stickySession_name"`
	Target_port         pulumi.IntOutput              `pulumi:"target_port"`
	Timeout             pulumi.IntPtrOutput           `pulumi:"timeout"`
	Token               pulumi.StringOutput           `pulumi:"token"`
	Type                pulumi.StringOutput           `pulumi:"type"`
	Url                 pulumi.StringOutput           `pulumi:"url"`
}

// NewReverseProxy registers a new resource with the given unique name, arguments, and options.
func NewReverseProxy(ctx *pulumi.Context,
	name string, args *ReverseProxyArgs, opts ...pulumi.ResourceOption) (*ReverseProxy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Cloudspace_id == nil {
		return nil, errors.New("invalid value for required argument 'Cloudspace_id'")
	}
	if args.CustomerID == nil {
		return nil, errors.New("invalid value for required argument 'CustomerID'")
	}
	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.Scheme == nil {
		return nil, errors.New("invalid value for required argument 'Scheme'")
	}
	if args.Serverpool_id == nil {
		return nil, errors.New("invalid value for required argument 'Serverpool_id'")
	}
	if args.Target_port == nil {
		return nil, errors.New("invalid value for required argument 'Target_port'")
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
	var resource ReverseProxy
	err := ctx.RegisterResource("vco:ingress:ReverseProxy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReverseProxy gets an existing ReverseProxy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReverseProxy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReverseProxyState, opts ...pulumi.ResourceOption) (*ReverseProxy, error) {
	var resource ReverseProxy
	err := ctx.ReadResource("vco:ingress:ReverseProxy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReverseProxy resources.
type reverseProxyState struct {
}

type ReverseProxyState struct {
}

func (ReverseProxyState) ElementType() reflect.Type {
	return reflect.TypeOf((*reverseProxyState)(nil)).Elem()
}

type reverseProxyArgs struct {
	Cloudspace_id       string  `pulumi:"cloudspace_id"`
	CustomerID          string  `pulumi:"customerID"`
	Description         *string `pulumi:"description"`
	Domain              string  `pulumi:"domain"`
	Email               *string `pulumi:"email"`
	Enabled             bool    `pulumi:"enabled"`
	Health_check_scheme *string `pulumi:"health_check_scheme"`
	Http_only           *bool   `pulumi:"http_only"`
	Http_port           *int    `pulumi:"http_port"`
	Https_port          *int    `pulumi:"https_port"`
	Interval            *int    `pulumi:"interval"`
	Ip_address          *string `pulumi:"ip_address"`
	Name                string  `pulumi:"name"`
	Path                *string `pulumi:"path"`
	Port                *int    `pulumi:"port"`
	Same_site           *string `pulumi:"same_site"`
	Scheme              string  `pulumi:"scheme"`
	Secure              *bool   `pulumi:"secure"`
	Serverpool_id       string  `pulumi:"serverpool_id"`
	StickySession_name  *string `pulumi:"stickySession_name"`
	Target_port         int     `pulumi:"target_port"`
	Timeout             *int    `pulumi:"timeout"`
	Token               string  `pulumi:"token"`
	Url                 string  `pulumi:"url"`
}

// The set of arguments for constructing a ReverseProxy resource.
type ReverseProxyArgs struct {
	Cloudspace_id       pulumi.StringInput
	CustomerID          pulumi.StringInput
	Description         pulumi.StringPtrInput
	Domain              pulumi.StringInput
	Email               pulumi.StringPtrInput
	Enabled             pulumi.BoolInput
	Health_check_scheme pulumi.StringPtrInput
	Http_only           pulumi.BoolPtrInput
	Http_port           pulumi.IntPtrInput
	Https_port          pulumi.IntPtrInput
	Interval            pulumi.IntPtrInput
	Ip_address          pulumi.StringPtrInput
	Name                pulumi.StringInput
	Path                pulumi.StringPtrInput
	Port                pulumi.IntPtrInput
	Same_site           pulumi.StringPtrInput
	Scheme              pulumi.StringInput
	Secure              pulumi.BoolPtrInput
	Serverpool_id       pulumi.StringInput
	StickySession_name  pulumi.StringPtrInput
	Target_port         pulumi.IntInput
	Timeout             pulumi.IntPtrInput
	Token               pulumi.StringInput
	Url                 pulumi.StringInput
}

func (ReverseProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*reverseProxyArgs)(nil)).Elem()
}

type ReverseProxyInput interface {
	pulumi.Input

	ToReverseProxyOutput() ReverseProxyOutput
	ToReverseProxyOutputWithContext(ctx context.Context) ReverseProxyOutput
}

func (*ReverseProxy) ElementType() reflect.Type {
	return reflect.TypeOf((**ReverseProxy)(nil)).Elem()
}

func (i *ReverseProxy) ToReverseProxyOutput() ReverseProxyOutput {
	return i.ToReverseProxyOutputWithContext(context.Background())
}

func (i *ReverseProxy) ToReverseProxyOutputWithContext(ctx context.Context) ReverseProxyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReverseProxyOutput)
}

// ReverseProxyArrayInput is an input type that accepts ReverseProxyArray and ReverseProxyArrayOutput values.
// You can construct a concrete instance of `ReverseProxyArrayInput` via:
//
//	ReverseProxyArray{ ReverseProxyArgs{...} }
type ReverseProxyArrayInput interface {
	pulumi.Input

	ToReverseProxyArrayOutput() ReverseProxyArrayOutput
	ToReverseProxyArrayOutputWithContext(context.Context) ReverseProxyArrayOutput
}

type ReverseProxyArray []ReverseProxyInput

func (ReverseProxyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReverseProxy)(nil)).Elem()
}

func (i ReverseProxyArray) ToReverseProxyArrayOutput() ReverseProxyArrayOutput {
	return i.ToReverseProxyArrayOutputWithContext(context.Background())
}

func (i ReverseProxyArray) ToReverseProxyArrayOutputWithContext(ctx context.Context) ReverseProxyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReverseProxyArrayOutput)
}

// ReverseProxyMapInput is an input type that accepts ReverseProxyMap and ReverseProxyMapOutput values.
// You can construct a concrete instance of `ReverseProxyMapInput` via:
//
//	ReverseProxyMap{ "key": ReverseProxyArgs{...} }
type ReverseProxyMapInput interface {
	pulumi.Input

	ToReverseProxyMapOutput() ReverseProxyMapOutput
	ToReverseProxyMapOutputWithContext(context.Context) ReverseProxyMapOutput
}

type ReverseProxyMap map[string]ReverseProxyInput

func (ReverseProxyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReverseProxy)(nil)).Elem()
}

func (i ReverseProxyMap) ToReverseProxyMapOutput() ReverseProxyMapOutput {
	return i.ToReverseProxyMapOutputWithContext(context.Background())
}

func (i ReverseProxyMap) ToReverseProxyMapOutputWithContext(ctx context.Context) ReverseProxyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReverseProxyMapOutput)
}

type ReverseProxyOutput struct{ *pulumi.OutputState }

func (ReverseProxyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReverseProxy)(nil)).Elem()
}

func (o ReverseProxyOutput) ToReverseProxyOutput() ReverseProxyOutput {
	return o
}

func (o ReverseProxyOutput) ToReverseProxyOutputWithContext(ctx context.Context) ReverseProxyOutput {
	return o
}

func (o ReverseProxyOutput) Back_end() ReverseProxyBackendPtrOutput {
	return o.ApplyT(func(v *ReverseProxy) ReverseProxyBackendPtrOutput { return v.Back_end }).(ReverseProxyBackendPtrOutput)
}

func (o ReverseProxyOutput) Cloudspace_id() pulumi.StringOutput {
	return o.ApplyT(func(v *ReverseProxy) pulumi.StringOutput { return v.Cloudspace_id }).(pulumi.StringOutput)
}

func (o ReverseProxyOutput) CustomerID() pulumi.StringOutput {
	return o.ApplyT(func(v *ReverseProxy) pulumi.StringOutput { return v.CustomerID }).(pulumi.StringOutput)
}

func (o ReverseProxyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReverseProxy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ReverseProxyOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *ReverseProxy) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

func (o ReverseProxyOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReverseProxy) pulumi.StringPtrOutput { return v.Email }).(pulumi.StringPtrOutput)
}

func (o ReverseProxyOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ReverseProxy) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

func (o ReverseProxyOutput) Front_end() ReverseProxyFrontEndPtrOutput {
	return o.ApplyT(func(v *ReverseProxy) ReverseProxyFrontEndPtrOutput { return v.Front_end }).(ReverseProxyFrontEndPtrOutput)
}

func (o ReverseProxyOutput) Health_check_scheme() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReverseProxy) pulumi.StringPtrOutput { return v.Health_check_scheme }).(pulumi.StringPtrOutput)
}

func (o ReverseProxyOutput) Http_only() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ReverseProxy) pulumi.BoolPtrOutput { return v.Http_only }).(pulumi.BoolPtrOutput)
}

func (o ReverseProxyOutput) Http_port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ReverseProxy) pulumi.IntPtrOutput { return v.Http_port }).(pulumi.IntPtrOutput)
}

func (o ReverseProxyOutput) Https_port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ReverseProxy) pulumi.IntPtrOutput { return v.Https_port }).(pulumi.IntPtrOutput)
}

func (o ReverseProxyOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ReverseProxy) pulumi.IntPtrOutput { return v.Interval }).(pulumi.IntPtrOutput)
}

func (o ReverseProxyOutput) Ip_address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReverseProxy) pulumi.StringPtrOutput { return v.Ip_address }).(pulumi.StringPtrOutput)
}

func (o ReverseProxyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ReverseProxy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ReverseProxyOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReverseProxy) pulumi.StringPtrOutput { return v.Path }).(pulumi.StringPtrOutput)
}

func (o ReverseProxyOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ReverseProxy) pulumi.IntPtrOutput { return v.Port }).(pulumi.IntPtrOutput)
}

func (o ReverseProxyOutput) Reverseproxy_id() pulumi.StringOutput {
	return o.ApplyT(func(v *ReverseProxy) pulumi.StringOutput { return v.Reverseproxy_id }).(pulumi.StringOutput)
}

func (o ReverseProxyOutput) Same_site() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReverseProxy) pulumi.StringPtrOutput { return v.Same_site }).(pulumi.StringPtrOutput)
}

func (o ReverseProxyOutput) Scheme() pulumi.StringOutput {
	return o.ApplyT(func(v *ReverseProxy) pulumi.StringOutput { return v.Scheme }).(pulumi.StringOutput)
}

func (o ReverseProxyOutput) Secure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ReverseProxy) pulumi.BoolPtrOutput { return v.Secure }).(pulumi.BoolPtrOutput)
}

func (o ReverseProxyOutput) Serverpool_id() pulumi.StringOutput {
	return o.ApplyT(func(v *ReverseProxy) pulumi.StringOutput { return v.Serverpool_id }).(pulumi.StringOutput)
}

func (o ReverseProxyOutput) StickySession_name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReverseProxy) pulumi.StringPtrOutput { return v.StickySession_name }).(pulumi.StringPtrOutput)
}

func (o ReverseProxyOutput) Target_port() pulumi.IntOutput {
	return o.ApplyT(func(v *ReverseProxy) pulumi.IntOutput { return v.Target_port }).(pulumi.IntOutput)
}

func (o ReverseProxyOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ReverseProxy) pulumi.IntPtrOutput { return v.Timeout }).(pulumi.IntPtrOutput)
}

func (o ReverseProxyOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *ReverseProxy) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

func (o ReverseProxyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ReverseProxy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ReverseProxyOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *ReverseProxy) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type ReverseProxyArrayOutput struct{ *pulumi.OutputState }

func (ReverseProxyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReverseProxy)(nil)).Elem()
}

func (o ReverseProxyArrayOutput) ToReverseProxyArrayOutput() ReverseProxyArrayOutput {
	return o
}

func (o ReverseProxyArrayOutput) ToReverseProxyArrayOutputWithContext(ctx context.Context) ReverseProxyArrayOutput {
	return o
}

func (o ReverseProxyArrayOutput) Index(i pulumi.IntInput) ReverseProxyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ReverseProxy {
		return vs[0].([]*ReverseProxy)[vs[1].(int)]
	}).(ReverseProxyOutput)
}

type ReverseProxyMapOutput struct{ *pulumi.OutputState }

func (ReverseProxyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReverseProxy)(nil)).Elem()
}

func (o ReverseProxyMapOutput) ToReverseProxyMapOutput() ReverseProxyMapOutput {
	return o
}

func (o ReverseProxyMapOutput) ToReverseProxyMapOutputWithContext(ctx context.Context) ReverseProxyMapOutput {
	return o
}

func (o ReverseProxyMapOutput) MapIndex(k pulumi.StringInput) ReverseProxyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ReverseProxy {
		return vs[0].(map[string]*ReverseProxy)[vs[1].(string)]
	}).(ReverseProxyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReverseProxyInput)(nil)).Elem(), &ReverseProxy{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReverseProxyArrayInput)(nil)).Elem(), ReverseProxyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReverseProxyMapInput)(nil)).Elem(), ReverseProxyMap{})
	pulumi.RegisterOutputType(ReverseProxyOutput{})
	pulumi.RegisterOutputType(ReverseProxyArrayOutput{})
	pulumi.RegisterOutputType(ReverseProxyMapOutput{})
}
