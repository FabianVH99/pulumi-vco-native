// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ingress

import (
	"context"
	"reflect"

	"github.com/fabianv-cloud/pulumi-vco-native/sdk/go/vco/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type BackEnd struct {
	Serverpool_id   string  `pulumi:"serverpool_id"`
	Serverpool_name *string `pulumi:"serverpool_name"`
	Target_port     int     `pulumi:"target_port"`
}

type BackEndOutput struct{ *pulumi.OutputState }

func (BackEndOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackEnd)(nil)).Elem()
}

func (o BackEndOutput) ToBackEndOutput() BackEndOutput {
	return o
}

func (o BackEndOutput) ToBackEndOutputWithContext(ctx context.Context) BackEndOutput {
	return o
}

func (o BackEndOutput) Serverpool_id() pulumi.StringOutput {
	return o.ApplyT(func(v BackEnd) string { return v.Serverpool_id }).(pulumi.StringOutput)
}

func (o BackEndOutput) Serverpool_name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackEnd) *string { return v.Serverpool_name }).(pulumi.StringPtrOutput)
}

func (o BackEndOutput) Target_port() pulumi.IntOutput {
	return o.ApplyT(func(v BackEnd) int { return v.Target_port }).(pulumi.IntOutput)
}

type BackEndPtrOutput struct{ *pulumi.OutputState }

func (BackEndPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackEnd)(nil)).Elem()
}

func (o BackEndPtrOutput) ToBackEndPtrOutput() BackEndPtrOutput {
	return o
}

func (o BackEndPtrOutput) ToBackEndPtrOutputWithContext(ctx context.Context) BackEndPtrOutput {
	return o
}

func (o BackEndPtrOutput) Elem() BackEndOutput {
	return o.ApplyT(func(v *BackEnd) BackEnd {
		if v != nil {
			return *v
		}
		var ret BackEnd
		return ret
	}).(BackEndOutput)
}

func (o BackEndPtrOutput) Serverpool_id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackEnd) *string {
		if v == nil {
			return nil
		}
		return &v.Serverpool_id
	}).(pulumi.StringPtrOutput)
}

func (o BackEndPtrOutput) Serverpool_name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackEnd) *string {
		if v == nil {
			return nil
		}
		return v.Serverpool_name
	}).(pulumi.StringPtrOutput)
}

func (o BackEndPtrOutput) Target_port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BackEnd) *int {
		if v == nil {
			return nil
		}
		return &v.Target_port
	}).(pulumi.IntPtrOutput)
}

type FrontEnd struct {
	Ip_address *string `pulumi:"ip_address"`
	Port       int     `pulumi:"port"`
	Tls        *TLS    `pulumi:"tls"`
}

type FrontEndOutput struct{ *pulumi.OutputState }

func (FrontEndOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontEnd)(nil)).Elem()
}

func (o FrontEndOutput) ToFrontEndOutput() FrontEndOutput {
	return o
}

func (o FrontEndOutput) ToFrontEndOutputWithContext(ctx context.Context) FrontEndOutput {
	return o
}

func (o FrontEndOutput) Ip_address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontEnd) *string { return v.Ip_address }).(pulumi.StringPtrOutput)
}

func (o FrontEndOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v FrontEnd) int { return v.Port }).(pulumi.IntOutput)
}

func (o FrontEndOutput) Tls() TLSPtrOutput {
	return o.ApplyT(func(v FrontEnd) *TLS { return v.Tls }).(TLSPtrOutput)
}

type FrontEndPtrOutput struct{ *pulumi.OutputState }

func (FrontEndPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FrontEnd)(nil)).Elem()
}

func (o FrontEndPtrOutput) ToFrontEndPtrOutput() FrontEndPtrOutput {
	return o
}

func (o FrontEndPtrOutput) ToFrontEndPtrOutputWithContext(ctx context.Context) FrontEndPtrOutput {
	return o
}

func (o FrontEndPtrOutput) Elem() FrontEndOutput {
	return o.ApplyT(func(v *FrontEnd) FrontEnd {
		if v != nil {
			return *v
		}
		var ret FrontEnd
		return ret
	}).(FrontEndOutput)
}

func (o FrontEndPtrOutput) Ip_address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FrontEnd) *string {
		if v == nil {
			return nil
		}
		return v.Ip_address
	}).(pulumi.StringPtrOutput)
}

func (o FrontEndPtrOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FrontEnd) *int {
		if v == nil {
			return nil
		}
		return &v.Port
	}).(pulumi.IntPtrOutput)
}

func (o FrontEndPtrOutput) Tls() TLSPtrOutput {
	return o.ApplyT(func(v *FrontEnd) *TLS {
		if v == nil {
			return nil
		}
		return v.Tls
	}).(TLSPtrOutput)
}

type HealthCheck struct {
	Interval *int    `pulumi:"interval"`
	Path     *string `pulumi:"path"`
	Port     *int    `pulumi:"port"`
	Scheme   *string `pulumi:"scheme"`
	Timeout  *int    `pulumi:"timeout"`
}

type HealthCheckOutput struct{ *pulumi.OutputState }

func (HealthCheckOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthCheck)(nil)).Elem()
}

func (o HealthCheckOutput) ToHealthCheckOutput() HealthCheckOutput {
	return o
}

func (o HealthCheckOutput) ToHealthCheckOutputWithContext(ctx context.Context) HealthCheckOutput {
	return o
}

func (o HealthCheckOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HealthCheck) *int { return v.Interval }).(pulumi.IntPtrOutput)
}

func (o HealthCheckOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthCheck) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o HealthCheckOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HealthCheck) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o HealthCheckOutput) Scheme() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthCheck) *string { return v.Scheme }).(pulumi.StringPtrOutput)
}

func (o HealthCheckOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HealthCheck) *int { return v.Timeout }).(pulumi.IntPtrOutput)
}

type HealthCheckPtrOutput struct{ *pulumi.OutputState }

func (HealthCheckPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HealthCheck)(nil)).Elem()
}

func (o HealthCheckPtrOutput) ToHealthCheckPtrOutput() HealthCheckPtrOutput {
	return o
}

func (o HealthCheckPtrOutput) ToHealthCheckPtrOutputWithContext(ctx context.Context) HealthCheckPtrOutput {
	return o
}

func (o HealthCheckPtrOutput) Elem() HealthCheckOutput {
	return o.ApplyT(func(v *HealthCheck) HealthCheck {
		if v != nil {
			return *v
		}
		var ret HealthCheck
		return ret
	}).(HealthCheckOutput)
}

func (o HealthCheckPtrOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HealthCheck) *int {
		if v == nil {
			return nil
		}
		return v.Interval
	}).(pulumi.IntPtrOutput)
}

func (o HealthCheckPtrOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HealthCheck) *string {
		if v == nil {
			return nil
		}
		return v.Path
	}).(pulumi.StringPtrOutput)
}

func (o HealthCheckPtrOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HealthCheck) *int {
		if v == nil {
			return nil
		}
		return v.Port
	}).(pulumi.IntPtrOutput)
}

func (o HealthCheckPtrOutput) Scheme() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HealthCheck) *string {
		if v == nil {
			return nil
		}
		return v.Scheme
	}).(pulumi.StringPtrOutput)
}

func (o HealthCheckPtrOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HealthCheck) *int {
		if v == nil {
			return nil
		}
		return v.Timeout
	}).(pulumi.IntPtrOutput)
}

type LetsEncrypt struct {
	Email   *string `pulumi:"email"`
	Enabled *bool   `pulumi:"enabled"`
}

type LetsEncryptOutput struct{ *pulumi.OutputState }

func (LetsEncryptOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LetsEncrypt)(nil)).Elem()
}

func (o LetsEncryptOutput) ToLetsEncryptOutput() LetsEncryptOutput {
	return o
}

func (o LetsEncryptOutput) ToLetsEncryptOutputWithContext(ctx context.Context) LetsEncryptOutput {
	return o
}

func (o LetsEncryptOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LetsEncrypt) *string { return v.Email }).(pulumi.StringPtrOutput)
}

func (o LetsEncryptOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LetsEncrypt) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

type LetsEncryptPtrOutput struct{ *pulumi.OutputState }

func (LetsEncryptPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LetsEncrypt)(nil)).Elem()
}

func (o LetsEncryptPtrOutput) ToLetsEncryptPtrOutput() LetsEncryptPtrOutput {
	return o
}

func (o LetsEncryptPtrOutput) ToLetsEncryptPtrOutputWithContext(ctx context.Context) LetsEncryptPtrOutput {
	return o
}

func (o LetsEncryptPtrOutput) Elem() LetsEncryptOutput {
	return o.ApplyT(func(v *LetsEncrypt) LetsEncrypt {
		if v != nil {
			return *v
		}
		var ret LetsEncrypt
		return ret
	}).(LetsEncryptOutput)
}

func (o LetsEncryptPtrOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LetsEncrypt) *string {
		if v == nil {
			return nil
		}
		return v.Email
	}).(pulumi.StringPtrOutput)
}

func (o LetsEncryptPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LetsEncrypt) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

type Options struct {
	Health_check          *HealthCheck         `pulumi:"health_check"`
	Sticky_session_cookie *StickySessionCookie `pulumi:"sticky_session_cookie"`
}

type OptionsOutput struct{ *pulumi.OutputState }

func (OptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Options)(nil)).Elem()
}

func (o OptionsOutput) ToOptionsOutput() OptionsOutput {
	return o
}

func (o OptionsOutput) ToOptionsOutputWithContext(ctx context.Context) OptionsOutput {
	return o
}

func (o OptionsOutput) Health_check() HealthCheckPtrOutput {
	return o.ApplyT(func(v Options) *HealthCheck { return v.Health_check }).(HealthCheckPtrOutput)
}

func (o OptionsOutput) Sticky_session_cookie() StickySessionCookiePtrOutput {
	return o.ApplyT(func(v Options) *StickySessionCookie { return v.Sticky_session_cookie }).(StickySessionCookiePtrOutput)
}

type OptionsPtrOutput struct{ *pulumi.OutputState }

func (OptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Options)(nil)).Elem()
}

func (o OptionsPtrOutput) ToOptionsPtrOutput() OptionsPtrOutput {
	return o
}

func (o OptionsPtrOutput) ToOptionsPtrOutputWithContext(ctx context.Context) OptionsPtrOutput {
	return o
}

func (o OptionsPtrOutput) Elem() OptionsOutput {
	return o.ApplyT(func(v *Options) Options {
		if v != nil {
			return *v
		}
		var ret Options
		return ret
	}).(OptionsOutput)
}

func (o OptionsPtrOutput) Health_check() HealthCheckPtrOutput {
	return o.ApplyT(func(v *Options) *HealthCheck {
		if v == nil {
			return nil
		}
		return v.Health_check
	}).(HealthCheckPtrOutput)
}

func (o OptionsPtrOutput) Sticky_session_cookie() StickySessionCookiePtrOutput {
	return o.ApplyT(func(v *Options) *StickySessionCookie {
		if v == nil {
			return nil
		}
		return v.Sticky_session_cookie
	}).(StickySessionCookiePtrOutput)
}

type ReverseProxyBackend struct {
	Options       *Options `pulumi:"options"`
	Scheme        string   `pulumi:"scheme"`
	Serverpool_id string   `pulumi:"serverpool_id"`
	Target_port   int      `pulumi:"target_port"`
}

type ReverseProxyBackendOutput struct{ *pulumi.OutputState }

func (ReverseProxyBackendOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReverseProxyBackend)(nil)).Elem()
}

func (o ReverseProxyBackendOutput) ToReverseProxyBackendOutput() ReverseProxyBackendOutput {
	return o
}

func (o ReverseProxyBackendOutput) ToReverseProxyBackendOutputWithContext(ctx context.Context) ReverseProxyBackendOutput {
	return o
}

func (o ReverseProxyBackendOutput) Options() OptionsPtrOutput {
	return o.ApplyT(func(v ReverseProxyBackend) *Options { return v.Options }).(OptionsPtrOutput)
}

func (o ReverseProxyBackendOutput) Scheme() pulumi.StringOutput {
	return o.ApplyT(func(v ReverseProxyBackend) string { return v.Scheme }).(pulumi.StringOutput)
}

func (o ReverseProxyBackendOutput) Serverpool_id() pulumi.StringOutput {
	return o.ApplyT(func(v ReverseProxyBackend) string { return v.Serverpool_id }).(pulumi.StringOutput)
}

func (o ReverseProxyBackendOutput) Target_port() pulumi.IntOutput {
	return o.ApplyT(func(v ReverseProxyBackend) int { return v.Target_port }).(pulumi.IntOutput)
}

type ReverseProxyBackendPtrOutput struct{ *pulumi.OutputState }

func (ReverseProxyBackendPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReverseProxyBackend)(nil)).Elem()
}

func (o ReverseProxyBackendPtrOutput) ToReverseProxyBackendPtrOutput() ReverseProxyBackendPtrOutput {
	return o
}

func (o ReverseProxyBackendPtrOutput) ToReverseProxyBackendPtrOutputWithContext(ctx context.Context) ReverseProxyBackendPtrOutput {
	return o
}

func (o ReverseProxyBackendPtrOutput) Elem() ReverseProxyBackendOutput {
	return o.ApplyT(func(v *ReverseProxyBackend) ReverseProxyBackend {
		if v != nil {
			return *v
		}
		var ret ReverseProxyBackend
		return ret
	}).(ReverseProxyBackendOutput)
}

func (o ReverseProxyBackendPtrOutput) Options() OptionsPtrOutput {
	return o.ApplyT(func(v *ReverseProxyBackend) *Options {
		if v == nil {
			return nil
		}
		return v.Options
	}).(OptionsPtrOutput)
}

func (o ReverseProxyBackendPtrOutput) Scheme() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReverseProxyBackend) *string {
		if v == nil {
			return nil
		}
		return &v.Scheme
	}).(pulumi.StringPtrOutput)
}

func (o ReverseProxyBackendPtrOutput) Serverpool_id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReverseProxyBackend) *string {
		if v == nil {
			return nil
		}
		return &v.Serverpool_id
	}).(pulumi.StringPtrOutput)
}

func (o ReverseProxyBackendPtrOutput) Target_port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ReverseProxyBackend) *int {
		if v == nil {
			return nil
		}
		return &v.Target_port
	}).(pulumi.IntPtrOutput)
}

type ReverseProxyFrontEnd struct {
	Domain      *string      `pulumi:"domain"`
	Http_port   *int         `pulumi:"http_port"`
	Https_port  *int         `pulumi:"https_port"`
	Ip_address  *string      `pulumi:"ip_address"`
	Letsencrypt *LetsEncrypt `pulumi:"letsencrypt"`
	Scheme      string       `pulumi:"scheme"`
}

type ReverseProxyFrontEndOutput struct{ *pulumi.OutputState }

func (ReverseProxyFrontEndOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReverseProxyFrontEnd)(nil)).Elem()
}

func (o ReverseProxyFrontEndOutput) ToReverseProxyFrontEndOutput() ReverseProxyFrontEndOutput {
	return o
}

func (o ReverseProxyFrontEndOutput) ToReverseProxyFrontEndOutputWithContext(ctx context.Context) ReverseProxyFrontEndOutput {
	return o
}

func (o ReverseProxyFrontEndOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReverseProxyFrontEnd) *string { return v.Domain }).(pulumi.StringPtrOutput)
}

func (o ReverseProxyFrontEndOutput) Http_port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ReverseProxyFrontEnd) *int { return v.Http_port }).(pulumi.IntPtrOutput)
}

func (o ReverseProxyFrontEndOutput) Https_port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ReverseProxyFrontEnd) *int { return v.Https_port }).(pulumi.IntPtrOutput)
}

func (o ReverseProxyFrontEndOutput) Ip_address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReverseProxyFrontEnd) *string { return v.Ip_address }).(pulumi.StringPtrOutput)
}

func (o ReverseProxyFrontEndOutput) Letsencrypt() LetsEncryptPtrOutput {
	return o.ApplyT(func(v ReverseProxyFrontEnd) *LetsEncrypt { return v.Letsencrypt }).(LetsEncryptPtrOutput)
}

func (o ReverseProxyFrontEndOutput) Scheme() pulumi.StringOutput {
	return o.ApplyT(func(v ReverseProxyFrontEnd) string { return v.Scheme }).(pulumi.StringOutput)
}

type ReverseProxyFrontEndPtrOutput struct{ *pulumi.OutputState }

func (ReverseProxyFrontEndPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReverseProxyFrontEnd)(nil)).Elem()
}

func (o ReverseProxyFrontEndPtrOutput) ToReverseProxyFrontEndPtrOutput() ReverseProxyFrontEndPtrOutput {
	return o
}

func (o ReverseProxyFrontEndPtrOutput) ToReverseProxyFrontEndPtrOutputWithContext(ctx context.Context) ReverseProxyFrontEndPtrOutput {
	return o
}

func (o ReverseProxyFrontEndPtrOutput) Elem() ReverseProxyFrontEndOutput {
	return o.ApplyT(func(v *ReverseProxyFrontEnd) ReverseProxyFrontEnd {
		if v != nil {
			return *v
		}
		var ret ReverseProxyFrontEnd
		return ret
	}).(ReverseProxyFrontEndOutput)
}

func (o ReverseProxyFrontEndPtrOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReverseProxyFrontEnd) *string {
		if v == nil {
			return nil
		}
		return v.Domain
	}).(pulumi.StringPtrOutput)
}

func (o ReverseProxyFrontEndPtrOutput) Http_port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ReverseProxyFrontEnd) *int {
		if v == nil {
			return nil
		}
		return v.Http_port
	}).(pulumi.IntPtrOutput)
}

func (o ReverseProxyFrontEndPtrOutput) Https_port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ReverseProxyFrontEnd) *int {
		if v == nil {
			return nil
		}
		return v.Https_port
	}).(pulumi.IntPtrOutput)
}

func (o ReverseProxyFrontEndPtrOutput) Ip_address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReverseProxyFrontEnd) *string {
		if v == nil {
			return nil
		}
		return v.Ip_address
	}).(pulumi.StringPtrOutput)
}

func (o ReverseProxyFrontEndPtrOutput) Letsencrypt() LetsEncryptPtrOutput {
	return o.ApplyT(func(v *ReverseProxyFrontEnd) *LetsEncrypt {
		if v == nil {
			return nil
		}
		return v.Letsencrypt
	}).(LetsEncryptPtrOutput)
}

func (o ReverseProxyFrontEndPtrOutput) Scheme() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReverseProxyFrontEnd) *string {
		if v == nil {
			return nil
		}
		return &v.Scheme
	}).(pulumi.StringPtrOutput)
}

type ServerPoolHost struct {
	Address *string `pulumi:"address"`
	Host_id *string `pulumi:"host_id"`
}

type ServerPoolHostOutput struct{ *pulumi.OutputState }

func (ServerPoolHostOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerPoolHost)(nil)).Elem()
}

func (o ServerPoolHostOutput) ToServerPoolHostOutput() ServerPoolHostOutput {
	return o
}

func (o ServerPoolHostOutput) ToServerPoolHostOutputWithContext(ctx context.Context) ServerPoolHostOutput {
	return o
}

func (o ServerPoolHostOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerPoolHost) *string { return v.Address }).(pulumi.StringPtrOutput)
}

func (o ServerPoolHostOutput) Host_id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerPoolHost) *string { return v.Host_id }).(pulumi.StringPtrOutput)
}

type ServerPoolHostArrayOutput struct{ *pulumi.OutputState }

func (ServerPoolHostArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServerPoolHost)(nil)).Elem()
}

func (o ServerPoolHostArrayOutput) ToServerPoolHostArrayOutput() ServerPoolHostArrayOutput {
	return o
}

func (o ServerPoolHostArrayOutput) ToServerPoolHostArrayOutputWithContext(ctx context.Context) ServerPoolHostArrayOutput {
	return o
}

func (o ServerPoolHostArrayOutput) Index(i pulumi.IntInput) ServerPoolHostOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServerPoolHost {
		return vs[0].([]ServerPoolHost)[vs[1].(int)]
	}).(ServerPoolHostOutput)
}

type StickySessionCookie struct {
	Http_only *bool   `pulumi:"http_only"`
	Name      *string `pulumi:"name"`
	Same_site *string `pulumi:"same_site"`
	Secure    *bool   `pulumi:"secure"`
}

type StickySessionCookieOutput struct{ *pulumi.OutputState }

func (StickySessionCookieOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StickySessionCookie)(nil)).Elem()
}

func (o StickySessionCookieOutput) ToStickySessionCookieOutput() StickySessionCookieOutput {
	return o
}

func (o StickySessionCookieOutput) ToStickySessionCookieOutputWithContext(ctx context.Context) StickySessionCookieOutput {
	return o
}

func (o StickySessionCookieOutput) Http_only() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v StickySessionCookie) *bool { return v.Http_only }).(pulumi.BoolPtrOutput)
}

func (o StickySessionCookieOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StickySessionCookie) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o StickySessionCookieOutput) Same_site() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StickySessionCookie) *string { return v.Same_site }).(pulumi.StringPtrOutput)
}

func (o StickySessionCookieOutput) Secure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v StickySessionCookie) *bool { return v.Secure }).(pulumi.BoolPtrOutput)
}

type StickySessionCookiePtrOutput struct{ *pulumi.OutputState }

func (StickySessionCookiePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StickySessionCookie)(nil)).Elem()
}

func (o StickySessionCookiePtrOutput) ToStickySessionCookiePtrOutput() StickySessionCookiePtrOutput {
	return o
}

func (o StickySessionCookiePtrOutput) ToStickySessionCookiePtrOutputWithContext(ctx context.Context) StickySessionCookiePtrOutput {
	return o
}

func (o StickySessionCookiePtrOutput) Elem() StickySessionCookieOutput {
	return o.ApplyT(func(v *StickySessionCookie) StickySessionCookie {
		if v != nil {
			return *v
		}
		var ret StickySessionCookie
		return ret
	}).(StickySessionCookieOutput)
}

func (o StickySessionCookiePtrOutput) Http_only() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StickySessionCookie) *bool {
		if v == nil {
			return nil
		}
		return v.Http_only
	}).(pulumi.BoolPtrOutput)
}

func (o StickySessionCookiePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StickySessionCookie) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o StickySessionCookiePtrOutput) Same_site() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StickySessionCookie) *string {
		if v == nil {
			return nil
		}
		return v.Same_site
	}).(pulumi.StringPtrOutput)
}

func (o StickySessionCookiePtrOutput) Secure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StickySessionCookie) *bool {
		if v == nil {
			return nil
		}
		return v.Secure
	}).(pulumi.BoolPtrOutput)
}

type TLS struct {
	Domain          *string `pulumi:"domain"`
	Is_enabled      *bool   `pulumi:"is_enabled"`
	Tls_termination *bool   `pulumi:"tls_termination"`
}

type TLSOutput struct{ *pulumi.OutputState }

func (TLSOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TLS)(nil)).Elem()
}

func (o TLSOutput) ToTLSOutput() TLSOutput {
	return o
}

func (o TLSOutput) ToTLSOutputWithContext(ctx context.Context) TLSOutput {
	return o
}

func (o TLSOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TLS) *string { return v.Domain }).(pulumi.StringPtrOutput)
}

func (o TLSOutput) Is_enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v TLS) *bool { return v.Is_enabled }).(pulumi.BoolPtrOutput)
}

func (o TLSOutput) Tls_termination() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v TLS) *bool { return v.Tls_termination }).(pulumi.BoolPtrOutput)
}

type TLSPtrOutput struct{ *pulumi.OutputState }

func (TLSPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TLS)(nil)).Elem()
}

func (o TLSPtrOutput) ToTLSPtrOutput() TLSPtrOutput {
	return o
}

func (o TLSPtrOutput) ToTLSPtrOutputWithContext(ctx context.Context) TLSPtrOutput {
	return o
}

func (o TLSPtrOutput) Elem() TLSOutput {
	return o.ApplyT(func(v *TLS) TLS {
		if v != nil {
			return *v
		}
		var ret TLS
		return ret
	}).(TLSOutput)
}

func (o TLSPtrOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TLS) *string {
		if v == nil {
			return nil
		}
		return v.Domain
	}).(pulumi.StringPtrOutput)
}

func (o TLSPtrOutput) Is_enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TLS) *bool {
		if v == nil {
			return nil
		}
		return v.Is_enabled
	}).(pulumi.BoolPtrOutput)
}

func (o TLSPtrOutput) Tls_termination() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TLS) *bool {
		if v == nil {
			return nil
		}
		return v.Tls_termination
	}).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(BackEndOutput{})
	pulumi.RegisterOutputType(BackEndPtrOutput{})
	pulumi.RegisterOutputType(FrontEndOutput{})
	pulumi.RegisterOutputType(FrontEndPtrOutput{})
	pulumi.RegisterOutputType(HealthCheckOutput{})
	pulumi.RegisterOutputType(HealthCheckPtrOutput{})
	pulumi.RegisterOutputType(LetsEncryptOutput{})
	pulumi.RegisterOutputType(LetsEncryptPtrOutput{})
	pulumi.RegisterOutputType(OptionsOutput{})
	pulumi.RegisterOutputType(OptionsPtrOutput{})
	pulumi.RegisterOutputType(ReverseProxyBackendOutput{})
	pulumi.RegisterOutputType(ReverseProxyBackendPtrOutput{})
	pulumi.RegisterOutputType(ReverseProxyFrontEndOutput{})
	pulumi.RegisterOutputType(ReverseProxyFrontEndPtrOutput{})
	pulumi.RegisterOutputType(ServerPoolHostOutput{})
	pulumi.RegisterOutputType(ServerPoolHostArrayOutput{})
	pulumi.RegisterOutputType(StickySessionCookieOutput{})
	pulumi.RegisterOutputType(StickySessionCookiePtrOutput{})
	pulumi.RegisterOutputType(TLSOutput{})
	pulumi.RegisterOutputType(TLSPtrOutput{})
}
