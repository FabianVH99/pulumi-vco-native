package provider

import (
	"github.com/blang/semver"
	"github.com/fabianv-cloud/pulumi-vco-native/provider/pkg/provider/anti_affinity_group"
	"github.com/fabianv-cloud/pulumi-vco-native/provider/pkg/provider/base"
	"github.com/fabianv-cloud/pulumi-vco-native/provider/pkg/provider/cloudspace"
	"github.com/fabianv-cloud/pulumi-vco-native/provider/pkg/provider/disk"
	"github.com/fabianv-cloud/pulumi-vco-native/provider/pkg/provider/ingress"
	"github.com/fabianv-cloud/pulumi-vco-native/provider/pkg/provider/objectspace"
	"github.com/fabianv-cloud/pulumi-vco-native/provider/pkg/provider/virtual_machine"
	"github.com/fabianv-cloud/pulumi-vco-native/provider/pkg/version"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi-go-provider/integration"
	"github.com/pulumi/pulumi-go-provider/middleware/schema"
	"strings"
)

func NewProvider() p.Provider {
	return infer.Provider(infer.Options{
		Metadata: schema.Metadata{
			DisplayName: "vco",
			Description: "The Whitesky.Cloud Vco Provider for Pulumi enables you to manipulate resources in the vco portal.\n" +
				"Please read the documentation or consult the API docs in your vco portal in order to properly understand how the resources work.\n" +
				"Before using this package you need to install the resource-vco plugin. The latest release of which can be found in the repository.\n" +
				"You then have to install the plugin on your system by storing the contents of the zipped file in: \n" +
				"Windows: %USERPROFILE%\\\\.pulumi\\\\plugins\\resource-vco-" + version.Version + ".\n" +
				"Linux and MacOS: ~/.pulumi/plugins/resource-vco-" + version.Version + ".\n" +
				"When you have moved the plugin to the appropriate directory you can in stall it using: pulumi plugin install resource vco" + version.Version + ".\n",
			Keywords: []string{
				"whitesky.cloud",
				"pulumi",
				"vco",
				"category/utility",
				"kind/native",
			},
			Homepage:   "https://whitesky.cloud/",
			Repository: "https://github.com/fabianv-cloud/pulumi-vco-native",
			Publisher:  "Whitesky.cloud",
			LanguageMap: map[string]any{
				"csharp": map[string]any{
					"packageReferences": map[string]string{
						"Pulumi": "3.*",
					},
				},
				"go": map[string]any{
					"generateResourceContainerTypes": true,
					"importBasePath":                 "github.com/fabianv-cloud/pulumi-vco-native/sdk/go/vco",
				},
				"nodejs": map[string]any{
					"dependencies": map[string]string{
						"@pulumi/pulumi": "^3.0.0",
					},
				},
				"python": map[string]any{
					"requires": map[string]string{
						"pulumi": ">=3.0.0,<4.0.0",
					},
				},
			},
		},
		Resources: []infer.InferredResource{
			infer.Resource[base.Cloudspace, base.CloudspaceArgs, base.CloudspaceState](),
			infer.Resource[base.Disk, base.DiskArgs, base.DiskState](),
			infer.Resource[base.ObjectSpace, base.ObjectSpaceArgs, base.ObjectSpaceState](),
			infer.Resource[cloudspace.AntiAffinityGroup, cloudspace.AntiAffinityGroupArgs, cloudspace.AntiAffinityGroupState](),
			infer.Resource[cloudspace.VirtualMachine, cloudspace.VirtualMachineArgs, cloudspace.VirtualMachineState](),
			infer.Resource[cloudspace.ConnectedCloudspace, cloudspace.ConnectedCloudspaceArgs, cloudspace.ConnectedCloudspaceState](),
			infer.Resource[cloudspace.ExternalNetwork, cloudspace.ExternalNetworkArgs, cloudspace.ExternalNetworkState](),
			infer.Resource[cloudspace.PortForward, cloudspace.PortForwardArgs, cloudspace.PortForwardState](),
			infer.Resource[anti_affinity_group.AntiAffinityGroupVM, anti_affinity_group.AntiAffinityGroupVMArgs, anti_affinity_group.AntiAffinityGroupVMState](),
			infer.Resource[disk.ExposedDisk, disk.ExposedDiskArgs, disk.ExposedDiskState](),
			infer.Resource[ingress.Host, ingress.HostArgs, ingress.HostState](),
			infer.Resource[ingress.LoadBalancer, ingress.LoadBalancerArgs, ingress.LoadBalancerState](),
			infer.Resource[ingress.ReverseProxy, ingress.ReverseProxyArgs, ingress.ReverseProxyState](),
			infer.Resource[ingress.ServerPool, ingress.ServerPoolArgs, ingress.ServerPoolState](),
			infer.Resource[objectspace.Link, objectspace.LinkArgs, objectspace.LinkState](),
			infer.Resource[virtual_machine.VirtualMachineCD, virtual_machine.VirtualMachineCDArgs, virtual_machine.VirtualMachineCDState](),
			infer.Resource[virtual_machine.VirtualMachineDisk, virtual_machine.VirtualMachineDiskArgs, virtual_machine.VirtualMachineDiskState](),
			infer.Resource[virtual_machine.VirtualMachineNIC, virtual_machine.VirtualMachineNICArgs, virtual_machine.VirtualMachineNICState](),
		},
	})
}

func Schema(version string) (string, error) {
	if strings.HasPrefix(version, "v") {
		version = version[1:]
	}
	s, err := integration.NewServer("vco", semver.MustParse(version), NewProvider()).
		GetSchema(p.GetSchemaRequest{})
	return s.Schema, err
}
