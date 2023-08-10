package main

import (
	"github.com/fabianv-cloud/pulumi-vco-native/provider/cmd/pulumi-resource-vco/resources"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

var Version = "0.1.0"

func main() {
	p.RunProvider("vco", Version,
		infer.Provider(infer.Options{
			Resources: []infer.InferredResource{
				infer.Resource[resources.AntiAffinityGroup, resources.AntiAffinityGroupArgs, resources.AntiAffinityGroupState](),
				infer.Resource[resources.AntiAffinityGroupVM, resources.AntiAffinityGroupVMArgs, resources.AntiAffinityGroupVMState](),
				infer.Resource[resources.Cloudspace, resources.CloudspaceArgs, resources.CloudspaceState](),
				infer.Resource[resources.ConnectedCloudspace, resources.ConnectedCloudspaceArgs, resources.ConnectedCloudspaceState](),
				infer.Resource[resources.Disk, resources.DiskArgs, resources.DiskState](),
				infer.Resource[resources.ExposedDisk, resources.ExposedDiskArgs, resources.ExposedDiskState](),
				infer.Resource[resources.ExternalNetwork, resources.ExternalNetworkArgs, resources.ExternalNetworkState](),
				infer.Resource[resources.Host, resources.HostArgs, resources.HostState](),
				infer.Resource[resources.LoadBalancer, resources.LoadBalancerArgs, resources.LoadBalancerState](),
				infer.Resource[resources.ObjectSpaceLink, resources.ObjectSpaceLinkArgs, resources.ObjectSpaceLinkState](),
				infer.Resource[resources.PortForward, resources.PortForwardArgs, resources.PortForwardState](),
				infer.Resource[resources.ReverseProxy, resources.ReverseProxyArgs, resources.ReverseProxyState](),
				infer.Resource[resources.ServerPool, resources.ServerPoolArgs, resources.ServerPoolState](),
				infer.Resource[resources.VirtualMachine, resources.VirtualMachineArgs, resources.VirtualMachineState](),
				infer.Resource[resources.VirtualMachineCD, resources.VirtualMachineCDArgs, resources.VirtualMachineCDState](),
				infer.Resource[resources.VirtualMachineDisk, resources.VirtualMachineDiskArgs, resources.VirtualMachineDiskState](),
				infer.Resource[resources.VirtualMachineNIC, resources.VirtualMachineNICArgs, resources.VirtualMachineNICState](),
			},
		}))
}
