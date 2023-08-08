// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package resources

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"internal"
)

var _ = internal.GetEnvOrDefault

type BackEndState struct {
	Serverpool_id   string `pulumi:"serverpool_id"`
	Serverpool_name string `pulumi:"serverpool_name"`
	Target_port     int    `pulumi:"target_port"`
}

type BackEndStateOutput struct{ *pulumi.OutputState }

func (BackEndStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackEndState)(nil)).Elem()
}

func (o BackEndStateOutput) ToBackEndStateOutput() BackEndStateOutput {
	return o
}

func (o BackEndStateOutput) ToBackEndStateOutputWithContext(ctx context.Context) BackEndStateOutput {
	return o
}

func (o BackEndStateOutput) Serverpool_id() pulumi.StringOutput {
	return o.ApplyT(func(v BackEndState) string { return v.Serverpool_id }).(pulumi.StringOutput)
}

func (o BackEndStateOutput) Serverpool_name() pulumi.StringOutput {
	return o.ApplyT(func(v BackEndState) string { return v.Serverpool_name }).(pulumi.StringOutput)
}

func (o BackEndStateOutput) Target_port() pulumi.IntOutput {
	return o.ApplyT(func(v BackEndState) int { return v.Target_port }).(pulumi.IntOutput)
}

type CpuTopology struct {
	Cores   int `pulumi:"cores"`
	Sockets int `pulumi:"sockets"`
	Threads int `pulumi:"threads"`
}

type CpuTopologyOutput struct{ *pulumi.OutputState }

func (CpuTopologyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CpuTopology)(nil)).Elem()
}

func (o CpuTopologyOutput) ToCpuTopologyOutput() CpuTopologyOutput {
	return o
}

func (o CpuTopologyOutput) ToCpuTopologyOutputWithContext(ctx context.Context) CpuTopologyOutput {
	return o
}

func (o CpuTopologyOutput) Cores() pulumi.IntOutput {
	return o.ApplyT(func(v CpuTopology) int { return v.Cores }).(pulumi.IntOutput)
}

func (o CpuTopologyOutput) Sockets() pulumi.IntOutput {
	return o.ApplyT(func(v CpuTopology) int { return v.Sockets }).(pulumi.IntOutput)
}

func (o CpuTopologyOutput) Threads() pulumi.IntOutput {
	return o.ApplyT(func(v CpuTopology) int { return v.Threads }).(pulumi.IntOutput)
}

type Endpoint struct {
	Address         string `pulumi:"address"`
	Name            string `pulumi:"name"`
	Port            int    `pulumi:"port"`
	Private_address string `pulumi:"private_address"`
	Private_port    int    `pulumi:"private_port"`
	Psk             string `pulumi:"psk"`
	User            string `pulumi:"user"`
}

type EndpointOutput struct{ *pulumi.OutputState }

func (EndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Endpoint)(nil)).Elem()
}

func (o EndpointOutput) ToEndpointOutput() EndpointOutput {
	return o
}

func (o EndpointOutput) ToEndpointOutputWithContext(ctx context.Context) EndpointOutput {
	return o
}

func (o EndpointOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v Endpoint) string { return v.Address }).(pulumi.StringOutput)
}

func (o EndpointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Endpoint) string { return v.Name }).(pulumi.StringOutput)
}

func (o EndpointOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v Endpoint) int { return v.Port }).(pulumi.IntOutput)
}

func (o EndpointOutput) Private_address() pulumi.StringOutput {
	return o.ApplyT(func(v Endpoint) string { return v.Private_address }).(pulumi.StringOutput)
}

func (o EndpointOutput) Private_port() pulumi.IntOutput {
	return o.ApplyT(func(v Endpoint) int { return v.Private_port }).(pulumi.IntOutput)
}

func (o EndpointOutput) Psk() pulumi.StringOutput {
	return o.ApplyT(func(v Endpoint) string { return v.Psk }).(pulumi.StringOutput)
}

func (o EndpointOutput) User() pulumi.StringOutput {
	return o.ApplyT(func(v Endpoint) string { return v.User }).(pulumi.StringOutput)
}

type FirewallCustom struct {
	Cdrom_id  int    `pulumi:"cdrom_id"`
	Disk_size int    `pulumi:"disk_size"`
	Image_id  int    `pulumi:"image_id"`
	Memory    int    `pulumi:"memory"`
	Type      string `pulumi:"type"`
	Vcpus     int    `pulumi:"vcpus"`
}

// FirewallCustomInput is an input type that accepts FirewallCustomArgs and FirewallCustomOutput values.
// You can construct a concrete instance of `FirewallCustomInput` via:
//
//	FirewallCustomArgs{...}
type FirewallCustomInput interface {
	pulumi.Input

	ToFirewallCustomOutput() FirewallCustomOutput
	ToFirewallCustomOutputWithContext(context.Context) FirewallCustomOutput
}

type FirewallCustomArgs struct {
	Cdrom_id  pulumi.IntInput    `pulumi:"cdrom_id"`
	Disk_size pulumi.IntInput    `pulumi:"disk_size"`
	Image_id  pulumi.IntInput    `pulumi:"image_id"`
	Memory    pulumi.IntInput    `pulumi:"memory"`
	Type      pulumi.StringInput `pulumi:"type"`
	Vcpus     pulumi.IntInput    `pulumi:"vcpus"`
}

func (FirewallCustomArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallCustom)(nil)).Elem()
}

func (i FirewallCustomArgs) ToFirewallCustomOutput() FirewallCustomOutput {
	return i.ToFirewallCustomOutputWithContext(context.Background())
}

func (i FirewallCustomArgs) ToFirewallCustomOutputWithContext(ctx context.Context) FirewallCustomOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallCustomOutput)
}

type FirewallCustomOutput struct{ *pulumi.OutputState }

func (FirewallCustomOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallCustom)(nil)).Elem()
}

func (o FirewallCustomOutput) ToFirewallCustomOutput() FirewallCustomOutput {
	return o
}

func (o FirewallCustomOutput) ToFirewallCustomOutputWithContext(ctx context.Context) FirewallCustomOutput {
	return o
}

func (o FirewallCustomOutput) Cdrom_id() pulumi.IntOutput {
	return o.ApplyT(func(v FirewallCustom) int { return v.Cdrom_id }).(pulumi.IntOutput)
}

func (o FirewallCustomOutput) Disk_size() pulumi.IntOutput {
	return o.ApplyT(func(v FirewallCustom) int { return v.Disk_size }).(pulumi.IntOutput)
}

func (o FirewallCustomOutput) Image_id() pulumi.IntOutput {
	return o.ApplyT(func(v FirewallCustom) int { return v.Image_id }).(pulumi.IntOutput)
}

func (o FirewallCustomOutput) Memory() pulumi.IntOutput {
	return o.ApplyT(func(v FirewallCustom) int { return v.Memory }).(pulumi.IntOutput)
}

func (o FirewallCustomOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v FirewallCustom) string { return v.Type }).(pulumi.StringOutput)
}

func (o FirewallCustomOutput) Vcpus() pulumi.IntOutput {
	return o.ApplyT(func(v FirewallCustom) int { return v.Vcpus }).(pulumi.IntOutput)
}

type FrontEnd struct {
	Ip_address string `pulumi:"ip_address"`
	Port       int    `pulumi:"port"`
	Tls        TLS    `pulumi:"tls"`
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

func (o FrontEndOutput) Ip_address() pulumi.StringOutput {
	return o.ApplyT(func(v FrontEnd) string { return v.Ip_address }).(pulumi.StringOutput)
}

func (o FrontEndOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v FrontEnd) int { return v.Port }).(pulumi.IntOutput)
}

func (o FrontEndOutput) Tls() TLSOutput {
	return o.ApplyT(func(v FrontEnd) TLS { return v.Tls }).(TLSOutput)
}

type HealthCheck struct {
	Interval int    `pulumi:"interval"`
	Path     string `pulumi:"path"`
	Port     int    `pulumi:"port"`
	Scheme   string `pulumi:"scheme"`
	Timeout  int    `pulumi:"timeout"`
}

// HealthCheckInput is an input type that accepts HealthCheckArgs and HealthCheckOutput values.
// You can construct a concrete instance of `HealthCheckInput` via:
//
//	HealthCheckArgs{...}
type HealthCheckInput interface {
	pulumi.Input

	ToHealthCheckOutput() HealthCheckOutput
	ToHealthCheckOutputWithContext(context.Context) HealthCheckOutput
}

type HealthCheckArgs struct {
	Interval pulumi.IntInput    `pulumi:"interval"`
	Path     pulumi.StringInput `pulumi:"path"`
	Port     pulumi.IntInput    `pulumi:"port"`
	Scheme   pulumi.StringInput `pulumi:"scheme"`
	Timeout  pulumi.IntInput    `pulumi:"timeout"`
}

func (HealthCheckArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthCheck)(nil)).Elem()
}

func (i HealthCheckArgs) ToHealthCheckOutput() HealthCheckOutput {
	return i.ToHealthCheckOutputWithContext(context.Background())
}

func (i HealthCheckArgs) ToHealthCheckOutputWithContext(ctx context.Context) HealthCheckOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthCheckOutput)
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

func (o HealthCheckOutput) Interval() pulumi.IntOutput {
	return o.ApplyT(func(v HealthCheck) int { return v.Interval }).(pulumi.IntOutput)
}

func (o HealthCheckOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v HealthCheck) string { return v.Path }).(pulumi.StringOutput)
}

func (o HealthCheckOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v HealthCheck) int { return v.Port }).(pulumi.IntOutput)
}

func (o HealthCheckOutput) Scheme() pulumi.StringOutput {
	return o.ApplyT(func(v HealthCheck) string { return v.Scheme }).(pulumi.StringOutput)
}

func (o HealthCheckOutput) Timeout() pulumi.IntOutput {
	return o.ApplyT(func(v HealthCheck) int { return v.Timeout }).(pulumi.IntOutput)
}

type IOTune struct {
	Iops float64 `pulumi:"iops"`
}

type IOTuneOutput struct{ *pulumi.OutputState }

func (IOTuneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IOTune)(nil)).Elem()
}

func (o IOTuneOutput) ToIOTuneOutput() IOTuneOutput {
	return o
}

func (o IOTuneOutput) ToIOTuneOutputWithContext(ctx context.Context) IOTuneOutput {
	return o
}

func (o IOTuneOutput) Iops() pulumi.Float64Output {
	return o.ApplyT(func(v IOTune) float64 { return v.Iops }).(pulumi.Float64Output)
}

type LetsEncrypt struct {
	Email   string `pulumi:"email"`
	Enabled bool   `pulumi:"enabled"`
}

// LetsEncryptInput is an input type that accepts LetsEncryptArgs and LetsEncryptOutput values.
// You can construct a concrete instance of `LetsEncryptInput` via:
//
//	LetsEncryptArgs{...}
type LetsEncryptInput interface {
	pulumi.Input

	ToLetsEncryptOutput() LetsEncryptOutput
	ToLetsEncryptOutputWithContext(context.Context) LetsEncryptOutput
}

type LetsEncryptArgs struct {
	Email   pulumi.StringInput `pulumi:"email"`
	Enabled pulumi.BoolInput   `pulumi:"enabled"`
}

func (LetsEncryptArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LetsEncrypt)(nil)).Elem()
}

func (i LetsEncryptArgs) ToLetsEncryptOutput() LetsEncryptOutput {
	return i.ToLetsEncryptOutputWithContext(context.Background())
}

func (i LetsEncryptArgs) ToLetsEncryptOutputWithContext(ctx context.Context) LetsEncryptOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LetsEncryptOutput)
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

func (o LetsEncryptOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v LetsEncrypt) string { return v.Email }).(pulumi.StringOutput)
}

func (o LetsEncryptOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LetsEncrypt) bool { return v.Enabled }).(pulumi.BoolOutput)
}

type NetworkInterface struct {
	Device_name            string `pulumi:"device_name"`
	External_cloudspace_id string `pulumi:"external_cloudspace_id"`
	Ip_address             string `pulumi:"ip_address"`
	Mac_address            string `pulumi:"mac_address"`
	Model                  string `pulumi:"model"`
	Network_id             int    `pulumi:"network_id"`
	Nic_type               string `pulumi:"nic_type"`
}

type NetworkInterfaceOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterface)(nil)).Elem()
}

func (o NetworkInterfaceOutput) ToNetworkInterfaceOutput() NetworkInterfaceOutput {
	return o
}

func (o NetworkInterfaceOutput) ToNetworkInterfaceOutputWithContext(ctx context.Context) NetworkInterfaceOutput {
	return o
}

func (o NetworkInterfaceOutput) Device_name() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkInterface) string { return v.Device_name }).(pulumi.StringOutput)
}

func (o NetworkInterfaceOutput) External_cloudspace_id() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkInterface) string { return v.External_cloudspace_id }).(pulumi.StringOutput)
}

func (o NetworkInterfaceOutput) Ip_address() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkInterface) string { return v.Ip_address }).(pulumi.StringOutput)
}

func (o NetworkInterfaceOutput) Mac_address() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkInterface) string { return v.Mac_address }).(pulumi.StringOutput)
}

func (o NetworkInterfaceOutput) Model() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkInterface) string { return v.Model }).(pulumi.StringOutput)
}

func (o NetworkInterfaceOutput) Network_id() pulumi.IntOutput {
	return o.ApplyT(func(v NetworkInterface) int { return v.Network_id }).(pulumi.IntOutput)
}

func (o NetworkInterfaceOutput) Nic_type() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkInterface) string { return v.Nic_type }).(pulumi.StringOutput)
}

type NetworkInterfaceArrayOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterface)(nil)).Elem()
}

func (o NetworkInterfaceArrayOutput) ToNetworkInterfaceArrayOutput() NetworkInterfaceArrayOutput {
	return o
}

func (o NetworkInterfaceArrayOutput) ToNetworkInterfaceArrayOutputWithContext(ctx context.Context) NetworkInterfaceArrayOutput {
	return o
}

func (o NetworkInterfaceArrayOutput) Index(i pulumi.IntInput) NetworkInterfaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkInterface {
		return vs[0].([]NetworkInterface)[vs[1].(int)]
	}).(NetworkInterfaceOutput)
}

type Options struct {
	Health_check          HealthCheck         `pulumi:"health_check"`
	Sticky_session_cookie StickySessionCookie `pulumi:"sticky_session_cookie"`
}

// OptionsInput is an input type that accepts OptionsArgs and OptionsOutput values.
// You can construct a concrete instance of `OptionsInput` via:
//
//	OptionsArgs{...}
type OptionsInput interface {
	pulumi.Input

	ToOptionsOutput() OptionsOutput
	ToOptionsOutputWithContext(context.Context) OptionsOutput
}

type OptionsArgs struct {
	Health_check          HealthCheckInput         `pulumi:"health_check"`
	Sticky_session_cookie StickySessionCookieInput `pulumi:"sticky_session_cookie"`
}

func (OptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Options)(nil)).Elem()
}

func (i OptionsArgs) ToOptionsOutput() OptionsOutput {
	return i.ToOptionsOutputWithContext(context.Background())
}

func (i OptionsArgs) ToOptionsOutputWithContext(ctx context.Context) OptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OptionsOutput)
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

func (o OptionsOutput) Health_check() HealthCheckOutput {
	return o.ApplyT(func(v Options) HealthCheck { return v.Health_check }).(HealthCheckOutput)
}

func (o OptionsOutput) Sticky_session_cookie() StickySessionCookieOutput {
	return o.ApplyT(func(v Options) StickySessionCookie { return v.Sticky_session_cookie }).(StickySessionCookieOutput)
}

type OsAccount struct {
	Login    string `pulumi:"login"`
	Password string `pulumi:"password"`
}

type OsAccountOutput struct{ *pulumi.OutputState }

func (OsAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OsAccount)(nil)).Elem()
}

func (o OsAccountOutput) ToOsAccountOutput() OsAccountOutput {
	return o
}

func (o OsAccountOutput) ToOsAccountOutputWithContext(ctx context.Context) OsAccountOutput {
	return o
}

func (o OsAccountOutput) Login() pulumi.StringOutput {
	return o.ApplyT(func(v OsAccount) string { return v.Login }).(pulumi.StringOutput)
}

func (o OsAccountOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v OsAccount) string { return v.Password }).(pulumi.StringOutput)
}

type OsAccountArrayOutput struct{ *pulumi.OutputState }

func (OsAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OsAccount)(nil)).Elem()
}

func (o OsAccountArrayOutput) ToOsAccountArrayOutput() OsAccountArrayOutput {
	return o
}

func (o OsAccountArrayOutput) ToOsAccountArrayOutputWithContext(ctx context.Context) OsAccountArrayOutput {
	return o
}

func (o OsAccountArrayOutput) Index(i pulumi.IntInput) OsAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OsAccount {
		return vs[0].([]OsAccount)[vs[1].(int)]
	}).(OsAccountOutput)
}

type ResourceLimits struct {
	External_network_quota float64 `pulumi:"external_network_quota"`
	Memory_quota           float64 `pulumi:"memory_quota"`
	Public_ip_quota        float64 `pulumi:"public_ip_quota"`
	Vcpu_quota             float64 `pulumi:"vcpu_quota"`
	Vdisk_space_quota      float64 `pulumi:"vdisk_space_quota"`
}

type ResourceLimitsOutput struct{ *pulumi.OutputState }

func (ResourceLimitsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceLimits)(nil)).Elem()
}

func (o ResourceLimitsOutput) ToResourceLimitsOutput() ResourceLimitsOutput {
	return o
}

func (o ResourceLimitsOutput) ToResourceLimitsOutputWithContext(ctx context.Context) ResourceLimitsOutput {
	return o
}

func (o ResourceLimitsOutput) External_network_quota() pulumi.Float64Output {
	return o.ApplyT(func(v ResourceLimits) float64 { return v.External_network_quota }).(pulumi.Float64Output)
}

func (o ResourceLimitsOutput) Memory_quota() pulumi.Float64Output {
	return o.ApplyT(func(v ResourceLimits) float64 { return v.Memory_quota }).(pulumi.Float64Output)
}

func (o ResourceLimitsOutput) Public_ip_quota() pulumi.Float64Output {
	return o.ApplyT(func(v ResourceLimits) float64 { return v.Public_ip_quota }).(pulumi.Float64Output)
}

func (o ResourceLimitsOutput) Vcpu_quota() pulumi.Float64Output {
	return o.ApplyT(func(v ResourceLimits) float64 { return v.Vcpu_quota }).(pulumi.Float64Output)
}

func (o ResourceLimitsOutput) Vdisk_space_quota() pulumi.Float64Output {
	return o.ApplyT(func(v ResourceLimits) float64 { return v.Vdisk_space_quota }).(pulumi.Float64Output)
}

type ReverseProxyBackend struct {
	Options       Options `pulumi:"options"`
	Scheme        string  `pulumi:"scheme"`
	Serverpool_id string  `pulumi:"serverpool_id"`
	Target_port   int     `pulumi:"target_port"`
}

// ReverseProxyBackendInput is an input type that accepts ReverseProxyBackendArgs and ReverseProxyBackendOutput values.
// You can construct a concrete instance of `ReverseProxyBackendInput` via:
//
//	ReverseProxyBackendArgs{...}
type ReverseProxyBackendInput interface {
	pulumi.Input

	ToReverseProxyBackendOutput() ReverseProxyBackendOutput
	ToReverseProxyBackendOutputWithContext(context.Context) ReverseProxyBackendOutput
}

type ReverseProxyBackendArgs struct {
	Options       OptionsInput       `pulumi:"options"`
	Scheme        pulumi.StringInput `pulumi:"scheme"`
	Serverpool_id pulumi.StringInput `pulumi:"serverpool_id"`
	Target_port   pulumi.IntInput    `pulumi:"target_port"`
}

func (ReverseProxyBackendArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReverseProxyBackend)(nil)).Elem()
}

func (i ReverseProxyBackendArgs) ToReverseProxyBackendOutput() ReverseProxyBackendOutput {
	return i.ToReverseProxyBackendOutputWithContext(context.Background())
}

func (i ReverseProxyBackendArgs) ToReverseProxyBackendOutputWithContext(ctx context.Context) ReverseProxyBackendOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReverseProxyBackendOutput)
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

func (o ReverseProxyBackendOutput) Options() OptionsOutput {
	return o.ApplyT(func(v ReverseProxyBackend) Options { return v.Options }).(OptionsOutput)
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

type ReverseProxyFrontEnd struct {
	Domain      string      `pulumi:"domain"`
	Http_port   int         `pulumi:"http_port"`
	Https_port  int         `pulumi:"https_port"`
	Ip_address  string      `pulumi:"ip_address"`
	Letsencrypt LetsEncrypt `pulumi:"letsencrypt"`
	Scheme      string      `pulumi:"scheme"`
}

// ReverseProxyFrontEndInput is an input type that accepts ReverseProxyFrontEndArgs and ReverseProxyFrontEndOutput values.
// You can construct a concrete instance of `ReverseProxyFrontEndInput` via:
//
//	ReverseProxyFrontEndArgs{...}
type ReverseProxyFrontEndInput interface {
	pulumi.Input

	ToReverseProxyFrontEndOutput() ReverseProxyFrontEndOutput
	ToReverseProxyFrontEndOutputWithContext(context.Context) ReverseProxyFrontEndOutput
}

type ReverseProxyFrontEndArgs struct {
	Domain      pulumi.StringInput `pulumi:"domain"`
	Http_port   pulumi.IntInput    `pulumi:"http_port"`
	Https_port  pulumi.IntInput    `pulumi:"https_port"`
	Ip_address  pulumi.StringInput `pulumi:"ip_address"`
	Letsencrypt LetsEncryptInput   `pulumi:"letsencrypt"`
	Scheme      pulumi.StringInput `pulumi:"scheme"`
}

func (ReverseProxyFrontEndArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReverseProxyFrontEnd)(nil)).Elem()
}

func (i ReverseProxyFrontEndArgs) ToReverseProxyFrontEndOutput() ReverseProxyFrontEndOutput {
	return i.ToReverseProxyFrontEndOutputWithContext(context.Background())
}

func (i ReverseProxyFrontEndArgs) ToReverseProxyFrontEndOutputWithContext(ctx context.Context) ReverseProxyFrontEndOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReverseProxyFrontEndOutput)
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

func (o ReverseProxyFrontEndOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v ReverseProxyFrontEnd) string { return v.Domain }).(pulumi.StringOutput)
}

func (o ReverseProxyFrontEndOutput) Http_port() pulumi.IntOutput {
	return o.ApplyT(func(v ReverseProxyFrontEnd) int { return v.Http_port }).(pulumi.IntOutput)
}

func (o ReverseProxyFrontEndOutput) Https_port() pulumi.IntOutput {
	return o.ApplyT(func(v ReverseProxyFrontEnd) int { return v.Https_port }).(pulumi.IntOutput)
}

func (o ReverseProxyFrontEndOutput) Ip_address() pulumi.StringOutput {
	return o.ApplyT(func(v ReverseProxyFrontEnd) string { return v.Ip_address }).(pulumi.StringOutput)
}

func (o ReverseProxyFrontEndOutput) Letsencrypt() LetsEncryptOutput {
	return o.ApplyT(func(v ReverseProxyFrontEnd) LetsEncrypt { return v.Letsencrypt }).(LetsEncryptOutput)
}

func (o ReverseProxyFrontEndOutput) Scheme() pulumi.StringOutput {
	return o.ApplyT(func(v ReverseProxyFrontEnd) string { return v.Scheme }).(pulumi.StringOutput)
}

type ServerPoolHost struct {
	Address string `pulumi:"address"`
	Host_id string `pulumi:"host_id"`
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

func (o ServerPoolHostOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v ServerPoolHost) string { return v.Address }).(pulumi.StringOutput)
}

func (o ServerPoolHostOutput) Host_id() pulumi.StringOutput {
	return o.ApplyT(func(v ServerPoolHost) string { return v.Host_id }).(pulumi.StringOutput)
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
	Http_only bool   `pulumi:"http_only"`
	Name      string `pulumi:"name"`
	Same_site string `pulumi:"same_site"`
	Secure    bool   `pulumi:"secure"`
}

// StickySessionCookieInput is an input type that accepts StickySessionCookieArgs and StickySessionCookieOutput values.
// You can construct a concrete instance of `StickySessionCookieInput` via:
//
//	StickySessionCookieArgs{...}
type StickySessionCookieInput interface {
	pulumi.Input

	ToStickySessionCookieOutput() StickySessionCookieOutput
	ToStickySessionCookieOutputWithContext(context.Context) StickySessionCookieOutput
}

type StickySessionCookieArgs struct {
	Http_only pulumi.BoolInput   `pulumi:"http_only"`
	Name      pulumi.StringInput `pulumi:"name"`
	Same_site pulumi.StringInput `pulumi:"same_site"`
	Secure    pulumi.BoolInput   `pulumi:"secure"`
}

func (StickySessionCookieArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StickySessionCookie)(nil)).Elem()
}

func (i StickySessionCookieArgs) ToStickySessionCookieOutput() StickySessionCookieOutput {
	return i.ToStickySessionCookieOutputWithContext(context.Background())
}

func (i StickySessionCookieArgs) ToStickySessionCookieOutputWithContext(ctx context.Context) StickySessionCookieOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StickySessionCookieOutput)
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

func (o StickySessionCookieOutput) Http_only() pulumi.BoolOutput {
	return o.ApplyT(func(v StickySessionCookie) bool { return v.Http_only }).(pulumi.BoolOutput)
}

func (o StickySessionCookieOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v StickySessionCookie) string { return v.Name }).(pulumi.StringOutput)
}

func (o StickySessionCookieOutput) Same_site() pulumi.StringOutput {
	return o.ApplyT(func(v StickySessionCookie) string { return v.Same_site }).(pulumi.StringOutput)
}

func (o StickySessionCookieOutput) Secure() pulumi.BoolOutput {
	return o.ApplyT(func(v StickySessionCookie) bool { return v.Secure }).(pulumi.BoolOutput)
}

type TLS struct {
	Domain          string `pulumi:"domain"`
	Is_enabled      bool   `pulumi:"is_enabled"`
	Tls_termination bool   `pulumi:"tls_termination"`
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

func (o TLSOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v TLS) string { return v.Domain }).(pulumi.StringOutput)
}

func (o TLSOutput) Is_enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v TLS) bool { return v.Is_enabled }).(pulumi.BoolOutput)
}

func (o TLSOutput) Tls_termination() pulumi.BoolOutput {
	return o.ApplyT(func(v TLS) bool { return v.Tls_termination }).(pulumi.BoolOutput)
}

type VmDisk struct {
	Description string `pulumi:"description"`
	Disk_id     int    `pulumi:"disk_id"`
	Disk_name   string `pulumi:"disk_name"`
	Disk_size   int    `pulumi:"disk_size"`
	Disk_type   string `pulumi:"disk_type"`
	Exposed     bool   `pulumi:"exposed"`
	Order       string `pulumi:"order"`
	Pci_bus     int    `pulumi:"pci_bus"`
	Pci_slot    int    `pulumi:"pci_slot"`
	Status      string `pulumi:"status"`
}

type VmDiskOutput struct{ *pulumi.OutputState }

func (VmDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VmDisk)(nil)).Elem()
}

func (o VmDiskOutput) ToVmDiskOutput() VmDiskOutput {
	return o
}

func (o VmDiskOutput) ToVmDiskOutputWithContext(ctx context.Context) VmDiskOutput {
	return o
}

func (o VmDiskOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v VmDisk) string { return v.Description }).(pulumi.StringOutput)
}

func (o VmDiskOutput) Disk_id() pulumi.IntOutput {
	return o.ApplyT(func(v VmDisk) int { return v.Disk_id }).(pulumi.IntOutput)
}

func (o VmDiskOutput) Disk_name() pulumi.StringOutput {
	return o.ApplyT(func(v VmDisk) string { return v.Disk_name }).(pulumi.StringOutput)
}

func (o VmDiskOutput) Disk_size() pulumi.IntOutput {
	return o.ApplyT(func(v VmDisk) int { return v.Disk_size }).(pulumi.IntOutput)
}

func (o VmDiskOutput) Disk_type() pulumi.StringOutput {
	return o.ApplyT(func(v VmDisk) string { return v.Disk_type }).(pulumi.StringOutput)
}

func (o VmDiskOutput) Exposed() pulumi.BoolOutput {
	return o.ApplyT(func(v VmDisk) bool { return v.Exposed }).(pulumi.BoolOutput)
}

func (o VmDiskOutput) Order() pulumi.StringOutput {
	return o.ApplyT(func(v VmDisk) string { return v.Order }).(pulumi.StringOutput)
}

func (o VmDiskOutput) Pci_bus() pulumi.IntOutput {
	return o.ApplyT(func(v VmDisk) int { return v.Pci_bus }).(pulumi.IntOutput)
}

func (o VmDiskOutput) Pci_slot() pulumi.IntOutput {
	return o.ApplyT(func(v VmDisk) int { return v.Pci_slot }).(pulumi.IntOutput)
}

func (o VmDiskOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v VmDisk) string { return v.Status }).(pulumi.StringOutput)
}

type VmDiskArrayOutput struct{ *pulumi.OutputState }

func (VmDiskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VmDisk)(nil)).Elem()
}

func (o VmDiskArrayOutput) ToVmDiskArrayOutput() VmDiskArrayOutput {
	return o
}

func (o VmDiskArrayOutput) ToVmDiskArrayOutputWithContext(ctx context.Context) VmDiskArrayOutput {
	return o
}

func (o VmDiskArrayOutput) Index(i pulumi.IntInput) VmDiskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VmDisk {
		return vs[0].([]VmDisk)[vs[1].(int)]
	}).(VmDiskOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallCustomInput)(nil)).Elem(), FirewallCustomArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HealthCheckInput)(nil)).Elem(), HealthCheckArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LetsEncryptInput)(nil)).Elem(), LetsEncryptArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OptionsInput)(nil)).Elem(), OptionsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReverseProxyBackendInput)(nil)).Elem(), ReverseProxyBackendArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReverseProxyFrontEndInput)(nil)).Elem(), ReverseProxyFrontEndArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StickySessionCookieInput)(nil)).Elem(), StickySessionCookieArgs{})
	pulumi.RegisterOutputType(BackEndStateOutput{})
	pulumi.RegisterOutputType(CpuTopologyOutput{})
	pulumi.RegisterOutputType(EndpointOutput{})
	pulumi.RegisterOutputType(FirewallCustomOutput{})
	pulumi.RegisterOutputType(FrontEndOutput{})
	pulumi.RegisterOutputType(HealthCheckOutput{})
	pulumi.RegisterOutputType(IOTuneOutput{})
	pulumi.RegisterOutputType(LetsEncryptOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceArrayOutput{})
	pulumi.RegisterOutputType(OptionsOutput{})
	pulumi.RegisterOutputType(OsAccountOutput{})
	pulumi.RegisterOutputType(OsAccountArrayOutput{})
	pulumi.RegisterOutputType(ResourceLimitsOutput{})
	pulumi.RegisterOutputType(ReverseProxyBackendOutput{})
	pulumi.RegisterOutputType(ReverseProxyFrontEndOutput{})
	pulumi.RegisterOutputType(ServerPoolHostOutput{})
	pulumi.RegisterOutputType(ServerPoolHostArrayOutput{})
	pulumi.RegisterOutputType(StickySessionCookieOutput{})
	pulumi.RegisterOutputType(TLSOutput{})
	pulumi.RegisterOutputType(VmDiskOutput{})
	pulumi.RegisterOutputType(VmDiskArrayOutput{})
}