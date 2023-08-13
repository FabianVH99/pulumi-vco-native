package tests

import (
	"github.com/fabianv-cloud/pulumi-vco-native/provider/pkg/provider/anti_affinity_group"
	"github.com/fabianv-cloud/pulumi-vco-native/provider/pkg/provider/base"
	"github.com/fabianv-cloud/pulumi-vco-native/provider/pkg/provider/cloudspace"
	"github.com/fabianv-cloud/pulumi-vco-native/provider/pkg/provider/disk"
	"github.com/fabianv-cloud/pulumi-vco-native/provider/pkg/provider/ingress"
	"github.com/fabianv-cloud/pulumi-vco-native/provider/pkg/provider/objectspace"
	"github.com/fabianv-cloud/pulumi-vco-native/provider/pkg/provider/virtual_machine"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var CloudSpaceState base.CloudspaceState
var VirtualMachineState cloudspace.VirtualMachineState
var ServerPoolState ingress.ServerPoolState
var AntiAffinityGroupState cloudspace.AntiAffinityGroupState
var DiskState base.DiskState

func TestProvider(t *testing.T) {
	token := refreshJWT()
	customer := os.Getenv("LAB_CUSTOMER")
	url := os.Getenv("LAB_URL")

	t.Run("TestCloudspace", func(t *testing.T) {
		cloudspace := base.Cloudspace{}

		_, state, err := cloudspace.Create(nil, "cs-test-pulumi", base.CloudspaceArgs{
			URL:               url,
			Token:             token,
			CustomerID:        customer,
			Name:              "Pulumi",
			Location:          "nl-rmd-dc01-001",
			PrivateNetwork:    "192.168.10.0/24",
			ExternalNetworkID: 13,
			Private:           false,
		}, false)
		assert.NoError(t, err)
		assert.Equal(t, "DEPLOYED", state.Status)

		CloudSpaceState = state
	})

	t.Run("TestAntiAffinityGroup", func(t *testing.T) {
		antiAffinityGroup := cloudspace.AntiAffinityGroup{}

		_, state, err := antiAffinityGroup.Create(nil, "ag-test-pulumi", cloudspace.AntiAffinityGroupArgs{
			URL:          url,
			Token:        token,
			CustomerID:   customer,
			CloudSpaceID: CloudSpaceState.CloudSpaceID,
			GroupID:      "Pulumi_AF_Group",
			Spread:       3,
		}, false)
		assert.NoError(t, err)
		assert.Equal(t, "INCOMPLETE", state.Status)

		updatedState, err := antiAffinityGroup.Update(nil, "ag-test-pulumi", state, cloudspace.AntiAffinityGroupArgs{
			URL:          url,
			Token:        token,
			CustomerID:   customer,
			CloudSpaceID: CloudSpaceState.CloudSpaceID,
			GroupID:      state.GroupID,
			Spread:       2,
		}, false)
		assert.NoError(t, err)
		assert.Equal(t, 2, state.Spread)
		AntiAffinityGroupState = updatedState
	})

	t.Run("TestConnectedCloudspace", func(t *testing.T) {
		newCloudspace := base.Cloudspace{}

		_, csState, err := newCloudspace.Create(nil, "cs-test-pulumi-2", base.CloudspaceArgs{
			URL:               url,
			Token:             token,
			CustomerID:        customer,
			Name:              "Pulumi-2",
			Location:          "nl-rmd-dc01-001",
			PrivateNetwork:    "192.168.11.0/24",
			ExternalNetworkID: 13,
			Private:           false,
		}, false)
		assert.NoError(t, err)
		assert.Equal(t, "DEPLOYED", csState.Status)

		connectedCloudspace := cloudspace.ConnectedCloudspace{}

		_, cCsState, err := connectedCloudspace.Create(nil, "cCs-test-pulumi", cloudspace.ConnectedCloudspaceArgs{
			URL:                   url,
			Token:                 token,
			CustomerID:            customer,
			CloudSpaceID:          csState.CloudSpaceID,
			ConnectedCloudSpaceID: CloudSpaceState.CloudSpaceID,
			RemoteCloudSpaceIP:    CloudSpaceState.ExternalNetworkIP,
			LocalCloudSpaceIP:     csState.ExternalNetworkIP,
		}, false)
		assert.NoError(t, err)
		assert.Equal(t, CloudSpaceState.CloudSpaceID, cCsState.ConnectedCloudSpaceID)

		err = connectedCloudspace.Delete(nil, "cCs-test-pulumi", cCsState)
		assert.NoError(t, err)

		err = newCloudspace.Delete(nil, "cs-test-pulumi-2", csState)
		assert.NoError(t, err)
	})

	t.Run("TestDisk", func(t *testing.T) {
		newDisk := base.Disk{}

		_, state, err := newDisk.Create(nil, "disk-test-pulumi", base.DiskArgs{
			URL:             url,
			Token:           token,
			CustomerID:      customer,
			Location:        CloudSpaceState.Location,
			DiskName:        "pulumi_disk",
			DiskDescription: "test",
			DiskSize:        100,
		}, false)
		assert.NoError(t, err)
		assert.NotEmpty(t, state.DiskID)

		DiskState = state

		exposedDisk := disk.ExposedDisk{}

		_, exposedDiskState, err := exposedDisk.Create(nil, "edisk-test-pulumi", disk.ExposedDiskArgs{
			URL:        url,
			Token:      token,
			CustomerID: customer,
			DiskID:     state.DiskID,
		}, false)
		assert.NoError(t, err)
		assert.Equal(t, state.DiskID, exposedDiskState.DiskID)
	})
	t.Run("TestExternalNetwork", func(t *testing.T) {
		externalNetwork := cloudspace.ExternalNetwork{}

		_, state, err := externalNetwork.Create(nil, "ext-test-pulumi", cloudspace.ExternalNetworkArgs{
			URL:                 url,
			Token:               token,
			CustomerID:          customer,
			CloudSpaceID:        CloudSpaceState.CloudSpaceID,
			ExternalNetworkID:   CloudSpaceState.ExternalNetworkID,
			ExternalNetworkType: "external",
			ExternalNetworkIP:   "185.151.60.222",
			Metric:              503,
		}, false)
		assert.NoError(t, err)
		assert.Equal(t, true, state.Status)

		updateStatus, updatedState, err := externalNetwork.Update(nil, state, cloudspace.ExternalNetworkArgs{
			URL:               url,
			Token:             token,
			CustomerID:        customer,
			CloudSpaceID:      CloudSpaceState.CloudSpaceID,
			ExternalNetworkID: CloudSpaceState.ExternalNetworkID,
			ExternalNetworkIP: "185.151.60.222/24",
			Metric:            502,
		})
		assert.NoError(t, err)
		assert.Equal(t, true, updateStatus)

		deleteStatus, err := externalNetwork.Delete(nil, cloudspace.ExternalNetworkArgs{
			URL:               url,
			Token:             token,
			CustomerID:        customer,
			CloudSpaceID:      CloudSpaceState.CloudSpaceID,
			ExternalNetworkID: CloudSpaceState.ExternalNetworkID,
			ExternalNetworkIP: updatedState.ExternalNetworkIP,
		})
		assert.NoError(t, err)
		assert.Equal(t, true, deleteStatus)
	})
	t.Run("TestServerPool", func(t *testing.T) {
		serverPool := ingress.ServerPool{}

		_, state, err := serverPool.Create(nil, "sv-test-pulumi", ingress.ServerPoolArgs{
			URL:          url,
			Token:        token,
			CustomerID:   customer,
			CloudSpaceID: CloudSpaceState.CloudSpaceID,
			Name:         "Pulumi_server_pool",
			Description:  "test",
		}, false)
		assert.NoError(t, err)
		assert.NotEmpty(t, state.ServerPoolID)

		updateStatus, updatedState, err := serverPool.Update(nil, state, ingress.ServerPoolArgs{
			URL:          url,
			Token:        token,
			CustomerID:   customer,
			CloudSpaceID: CloudSpaceState.CloudSpaceID,
			Name:         "Pulumi_server_pool_update",
			Description:  "Update",
		})
		assert.NoError(t, err)
		assert.Equal(t, true, updateStatus)

		ServerPoolState = updatedState
	})
	t.Run("TestHost", func(t *testing.T) {
		host := ingress.Host{}

		_, hostState, err := host.Create(nil, "host-test-pulumi", ingress.HostArgs{
			URL:          url,
			Token:        token,
			CustomerID:   customer,
			CloudSpaceID: CloudSpaceState.CloudSpaceID,
			ServerPoolID: ServerPoolState.ServerPoolID,
			Address:      "192.168.10.10",
		}, false)
		assert.NoError(t, err)
		assert.NotEmpty(t, hostState.HostID)

		status, err := host.Delete(nil, hostState, ingress.HostArgs{
			URL:          url,
			Token:        token,
			CustomerID:   customer,
			CloudSpaceID: CloudSpaceState.CloudSpaceID,
			ServerPoolID: ServerPoolState.ServerPoolID,
			Address:      "192.168.10.10",
		})
		assert.NoError(t, err)
		assert.Equal(t, true, status)
	})
	t.Run("TestLoadBalancer", func(t *testing.T) {
		loadBalancer := ingress.LoadBalancer{}

		_, state, err := loadBalancer.Create(nil, "lb-test-pulumi", ingress.LoadBalancerArgs{
			URL:          url,
			Token:        token,
			CustomerID:   customer,
			CloudSpaceID: CloudSpaceState.CloudSpaceID,
			Name:         "Pulumi_lb_test",
			Type:         "tcp",
			FrontEnd: ingress.FrontEnd{
				Port: 23,
			},
			BackEnd: ingress.BackEnd{
				ServerpoolID: ServerPoolState.ServerPoolID,
				TargetPort:   23,
			},
		}, false)
		assert.NoError(t, err)
		assert.NotEmpty(t, state.LoadBalancerID)

		updateStatus, _, err := loadBalancer.Update(nil, state, ingress.LoadBalancerArgs{
			URL:          url,
			Token:        token,
			CustomerID:   customer,
			CloudSpaceID: CloudSpaceState.CloudSpaceID,
			Name:         "Pulumi_lb_test",
			Type:         "udp",
			FrontEnd: ingress.FrontEnd{
				Port: 23,
			},
			BackEnd: ingress.BackEnd{
				ServerpoolID: ServerPoolState.ServerPoolID,
				TargetPort:   23,
			},
		})
		assert.NoError(t, err)
		assert.Equal(t, true, updateStatus)

		status, err := loadBalancer.Delete(nil, state, ingress.LoadBalancerArgs{
			URL:          url,
			Token:        token,
			CustomerID:   customer,
			CloudSpaceID: CloudSpaceState.CloudSpaceID,
		})
		assert.NoError(t, err)
		assert.Equal(t, true, status)
	})
	t.Run("TestReverseProxy", func(t *testing.T) {
		reverseProxy := ingress.ReverseProxy{}

		_, state, err := reverseProxy.Create(nil, "rp-test-pulumi", ingress.ReverseProxyArgs{
			URL:          url,
			Token:        token,
			CustomerID:   customer,
			CloudSpaceID: CloudSpaceState.CloudSpaceID,
			Name:         "Reverse_Proxy_Pulumi",
			ReverseProxyFrontEnd: ingress.ReverseProxyFrontEnd{
				Domain:   "Pulumi",
				Scheme:   "http",
				HTTPPort: intPtr(25),
				LetsEncrypt: ingress.LetsEncrypt{
					Enabled: false,
				},
			},
			ReverseProxyBackend: ingress.ReverseProxyBackend{
				Scheme:       "http",
				ServerpoolID: ServerPoolState.ServerPoolID,
				TargetPort:   25,
			},
		}, false)
		assert.NoError(t, err)
		assert.NotEmpty(t, state.ReverseProxyID)

		updateStatus, updatedState, err := reverseProxy.Update(nil, state, ingress.ReverseProxyArgs{
			URL:          url,
			Token:        token,
			CustomerID:   customer,
			CloudSpaceID: CloudSpaceState.CloudSpaceID,
			Name:         "Reverse_Proxy_Pulumi",
			ReverseProxyFrontEnd: ingress.ReverseProxyFrontEnd{
				Domain:   "Pulumi",
				Scheme:   "http",
				HTTPPort: intPtr(25),
				LetsEncrypt: ingress.LetsEncrypt{
					Enabled: false,
				},
			},
			ReverseProxyBackend: ingress.ReverseProxyBackend{
				Scheme:       "http",
				ServerpoolID: ServerPoolState.ServerPoolID,
				TargetPort:   26,
			},
		})
		assert.NoError(t, err)
		assert.Equal(t, true, updateStatus)

		deleteStatus, err := reverseProxy.Delete(nil, updatedState, ingress.ReverseProxyArgs{
			URL:          url,
			Token:        token,
			CustomerID:   customer,
			CloudSpaceID: CloudSpaceState.CloudSpaceID,
		})
		assert.NoError(t, err)
		assert.Equal(t, true, deleteStatus)
	})
	t.Run("TestObjectSpace", func(t *testing.T) {
		objectSpace := base.ObjectSpace{}

		_, state, err := objectSpace.Create(nil, "obj-test-pulumi", base.ObjectSpaceArgs{
			URL:        url,
			Token:      token,
			CustomerID: customer,
			Location:   "nl-rmd-dc01-001",
			Name:       "Pulumi_obj_space",
		}, false)
		assert.NoError(t, err)
		assert.NotEmpty(t, state.ObjectSpaceID)

		objectSpaceLink := objectspace.Link{}

		_, linkState, err := objectSpaceLink.Create(nil, "objlink-test-pulumi", objectspace.LinkArgs{
			URL:           url,
			Token:         token,
			CustomerID:    customer,
			CloudSpaceID:  CloudSpaceState.CloudSpaceID,
			ObjectSpaceID: state.ObjectSpaceID,
		}, false)
		assert.NoError(t, err)
		assert.Equal(t, true, linkState.Status)

		status, err := objectSpaceLink.Delete(nil, linkState, objectspace.LinkArgs{
			URL:          url,
			Token:        token,
			CustomerID:   customer,
			CloudSpaceID: CloudSpaceState.CloudSpaceID,
		})
		assert.NoError(t, err)
		assert.Equal(t, true, status)

		deleteStatus, err := objectSpace.Delete(nil, state, base.ObjectSpaceArgs{
			URL:        url,
			Token:      token,
			CustomerID: customer,
		})
		assert.NoError(t, err)
		assert.Equal(t, true, deleteStatus)
	})
	t.Run("TestVirtualMachine", func(t *testing.T) {
		virtualMachine := cloudspace.VirtualMachine{}

		_, state, err := virtualMachine.Create(nil, "vm-test-pulumi", cloudspace.VirtualMachineArgs{
			URL:          url,
			Token:        token,
			CustomerID:   customer,
			CloudSpaceID: CloudSpaceState.CloudSpaceID,
			Name:         "Pulumi-VM",
			Description:  "Machine_for_testing_purposes",
			Vcpus:        4,
			Memory:       2048,
			ImageID:      intPtr(258),
			DiskSize:     intPtr(250),
			BootType:     stringPtr("bios"),
		}, false)
		assert.NoError(t, err)
		assert.NotEmpty(t, state.VirtualMachineID)

		VirtualMachineState = state
	})
	t.Run("TestVirtualMachineAntiAffinityGroup", func(t *testing.T) {
		antiAffinityVM := anti_affinity_group.AntiAffinityGroupVM{}

		_, state, err := antiAffinityVM.Create(nil, "vm-test-pulumi", anti_affinity_group.AntiAffinityGroupVMArgs{
			URL:              url,
			Token:            token,
			CustomerID:       customer,
			CloudSpaceID:     CloudSpaceState.CloudSpaceID,
			GroupID:          AntiAffinityGroupState.GroupID,
			VirtualMachineID: VirtualMachineState.VirtualMachineID,
		}, false)
		assert.NoError(t, err)
		assert.Equal(t, "RUNNING", state.Status)

		deleteStatus, err := antiAffinityVM.Delete(nil, state, anti_affinity_group.AntiAffinityGroupVMArgs{
			URL:          url,
			Token:        token,
			CustomerID:   customer,
			CloudSpaceID: CloudSpaceState.CloudSpaceID,
			GroupID:      AntiAffinityGroupState.GroupID,
		})
		assert.NoError(t, err)
		assert.Equal(t, true, deleteStatus)
	})

	t.Run("TestPortForward", func(t *testing.T) {
		portForward := cloudspace.PortForward{}

		_, state, err := portForward.Create(nil, "pf-test-pulumi", cloudspace.PortForwardArgs{
			URL:              url,
			Token:            token,
			CustomerID:       customer,
			CloudSpaceID:     CloudSpaceState.CloudSpaceID,
			LocalPort:        22,
			PublicPort:       22,
			Protocol:         "tcp",
			PublicIP:         CloudSpaceState.ExternalNetworkIP,
			VirtualMachineID: VirtualMachineState.VirtualMachineID,
		}, false)
		assert.NoError(t, err)
		assert.NotEmpty(t, state.PortForwardID)

		updatedState, err := portForward.Update(nil, state, cloudspace.PortForwardArgs{
			URL:              url,
			Token:            token,
			CustomerID:       customer,
			CloudSpaceID:     CloudSpaceState.CloudSpaceID,
			Protocol:         "udp",
			PublicPort:       22,
			LocalPort:        22,
			VirtualMachineID: VirtualMachineState.VirtualMachineID,
		})
		assert.NoError(t, err)
		assert.Equal(t, "udp", updatedState.Protocol)

		statusCode, err := portForward.Delete(nil, updatedState, cloudspace.PortForwardArgs{
			URL:          url,
			Token:        token,
			CustomerID:   customer,
			CloudSpaceID: CloudSpaceState.CloudSpaceID,
		})
		assert.NoError(t, err)
		assert.Equal(t, 200, statusCode)
	})
	t.Run("TestVirtualMachineCDrom", func(t *testing.T) {
		virtualMachineCD := virtual_machine.VirtualMachineCD{}

		_, state, err := virtualMachineCD.Create(nil, "vmcd-test-pulumi", virtual_machine.VirtualMachineCDArgs{
			URL:              url,
			Token:            token,
			CustomerID:       customer,
			CloudSpaceID:     CloudSpaceState.CloudSpaceID,
			VirtualMachineID: VirtualMachineState.VirtualMachineID,
			CdRomID:          9000,
		}, false)
		assert.NoError(t, err)
		assert.Equal(t, "CREATED", state.Status)

		deleteStatus, err := virtualMachineCD.Delete(nil, state, virtual_machine.VirtualMachineCDArgs{
			URL:              url,
			Token:            token,
			CustomerID:       customer,
			CloudSpaceID:     CloudSpaceState.CloudSpaceID,
			VirtualMachineID: VirtualMachineState.VirtualMachineID,
		})
		assert.NoError(t, err)
		assert.Equal(t, true, deleteStatus)
	})
	t.Run("TestVirtualMachineDisk", func(t *testing.T) {
		virtualMachineDisk := virtual_machine.VirtualMachineDisk{}

		_, state, err := virtualMachineDisk.Create(nil, "vmdsk-test-pulumi", virtual_machine.VirtualMachineDiskArgs{
			URL:              url,
			Token:            token,
			CustomerID:       customer,
			CloudSpaceID:     CloudSpaceState.CloudSpaceID,
			VirtualMachineID: VirtualMachineState.VirtualMachineID,
			DiskID:           DiskState.DiskID,
		}, false)
		assert.NoError(t, err)
		assert.Equal(t, true, state.Success)

		deleteStatus, err := virtualMachineDisk.Delete(nil, state, virtual_machine.VirtualMachineDiskArgs{
			URL:              url,
			Token:            token,
			CustomerID:       customer,
			CloudSpaceID:     CloudSpaceState.CloudSpaceID,
			VirtualMachineID: VirtualMachineState.VirtualMachineID,
			DiskID:           DiskState.DiskID,
		})
		assert.NoError(t, err)
		assert.Equal(t, true, deleteStatus)
	})
	t.Run("TestVirtualMachineNIC", func(t *testing.T) {
		VirtualMachineNIC := virtual_machine.VirtualMachineNIC{}

		_, state, err := VirtualMachineNIC.Create(nil, "vmdsk-test-pulumi", virtual_machine.VirtualMachineNICArgs{
			URL:               url,
			Token:             token,
			CustomerID:        customer,
			CloudSpaceID:      CloudSpaceState.CloudSpaceID,
			VirtualMachineID:  VirtualMachineState.VirtualMachineID,
			ExternalNetworkID: 13,
		}, false)
		assert.NoError(t, err)
		assert.NotEmpty(t, state.DeviceName)

		deleteStatus, err := VirtualMachineNIC.Delete(nil, state, virtual_machine.VirtualMachineNICArgs{
			URL:              url,
			Token:            token,
			CustomerID:       customer,
			CloudSpaceID:     CloudSpaceState.CloudSpaceID,
			VirtualMachineID: VirtualMachineState.VirtualMachineID,
		})
		assert.NoError(t, err)
		assert.Equal(t, true, deleteStatus)
	})
	t.Run("CleanUp", func(t *testing.T) {
		serverPool := ingress.ServerPool{}

		svDeleteStatus, err := serverPool.Delete(nil, ServerPoolState, ingress.ServerPoolArgs{
			URL:          url,
			Token:        token,
			CustomerID:   customer,
			CloudSpaceID: CloudSpaceState.CloudSpaceID,
		})
		assert.NoError(t, err)
		assert.Equal(t, true, svDeleteStatus)

		antiAffinityGroup := cloudspace.AntiAffinityGroup{}

		_, err = antiAffinityGroup.Delete(nil, AntiAffinityGroupState, cloudspace.AntiAffinityGroupArgs{
			URL:          url,
			Token:        token,
			CustomerID:   customer,
			CloudSpaceID: CloudSpaceState.CloudSpaceID,
		})
		assert.NoError(t, err)

		disk := base.Disk{}

		dskDeleteStatus, err := disk.Delete(nil, DiskState, base.DiskArgs{
			URL:        url,
			Token:      token,
			CustomerID: customer,
			Location:   CloudSpaceState.Location,
		})
		assert.NoError(t, err)
		assert.Equal(t, true, dskDeleteStatus)

		virtualMachine := cloudspace.VirtualMachine{}

		vmDeleteStatus, err := virtualMachine.Delete(nil, VirtualMachineState, cloudspace.VirtualMachineArgs{
			URL:          url,
			Token:        token,
			CustomerID:   customer,
			CloudSpaceID: CloudSpaceState.CloudSpaceID,
		})
		assert.NoError(t, err)
		assert.Equal(t, true, vmDeleteStatus)

		cloudspace := base.Cloudspace{}

		csDeleteStatus, err := cloudspace.Delete(nil, CloudSpaceState, base.CloudspaceArgs{
			URL:        url,
			Token:      token,
			CustomerID: customer,
		})
		assert.NoError(t, err)
		assert.Equal(t, true, csDeleteStatus)
	})
}
