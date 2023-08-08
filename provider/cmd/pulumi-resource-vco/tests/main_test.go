package tests

import (
	"github.com/fabianv-cloud/pulumi-vco-native/provider/cmd/pulumi-resource-vco/resources"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var CloudSpaceState resources.CloudspaceState
var VirtualMachineState resources.VirtualMachineState
var ServerPoolState resources.ServerPoolState
var AntiAffinityGroupState resources.AntiAffinityGroupState
var DiskState resources.DiskState

func TestProvider(t *testing.T) {
	token := refreshJWT()
	customer := os.Getenv("LAB_CUSTOMER")
	url := os.Getenv("LAB_URL")

	t.Run("TestCloudspace", func(t *testing.T) {
		cloudspace := resources.Cloudspace{}

		_, state, err := cloudspace.Create(nil, "cs-test-pulumi", resources.CloudspaceArgs{
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
		antiAffinityGroup := resources.AntiAffinityGroup{}

		_, state, err := antiAffinityGroup.Create(nil, "ag-test-pulumi", resources.AntiAffinityGroupArgs{
			URL:          url,
			Token:        token,
			CustomerID:   customer,
			CloudSpaceID: CloudSpaceState.CloudSpaceID,
			GroupID:      "Pulumi_AF_Group",
			Spread:       3,
		}, false)
		assert.NoError(t, err)
		assert.Equal(t, "INCOMPLETE", state.Status)

		status, updatedState, err := antiAffinityGroup.Update(nil, state, resources.AntiAffinityGroupArgs{
			URL:          url,
			Token:        token,
			CustomerID:   customer,
			CloudSpaceID: CloudSpaceState.CloudSpaceID,
			GroupID:      state.GroupID,
			Spread:       2,
		})
		assert.NoError(t, err)
		assert.Equal(t, true, status)

		AntiAffinityGroupState = updatedState
	})

	t.Run("TestConnectedCloudspace", func(t *testing.T) {
		cloudspace := resources.Cloudspace{}

		_, csState, err := cloudspace.Create(nil, "cs-test-pulumi-2", resources.CloudspaceArgs{
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

		connectedCloudspace := resources.ConnectedCloudspace{}

		_, cCsState, err := connectedCloudspace.Create(nil, "cCs-test-pulumi", resources.ConnectedCloudspaceArgs{
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

		statusCode, err := connectedCloudspace.Delete(nil, resources.ConnectedCloudspaceArgs{
			URL:                   url,
			Token:                 token,
			CustomerID:            customer,
			CloudSpaceID:          csState.CloudSpaceID,
			ConnectedCloudSpaceID: CloudSpaceState.CloudSpaceID,
			RemoteCloudSpaceIP:    CloudSpaceState.ExternalNetworkIP,
			LocalCloudSpaceIP:     csState.ExternalNetworkIP,
		})
		assert.NoError(t, err)
		assert.Equal(t, 204, statusCode)

		csDeleteStatus, err := cloudspace.Delete(nil, csState, resources.CloudspaceArgs{
			URL:        url,
			Token:      token,
			CustomerID: customer,
		})
		assert.NoError(t, err)
		assert.Equal(t, true, csDeleteStatus)
	})

	t.Run("TestDisk", func(t *testing.T) {
		disk := resources.Disk{}

		_, state, err := disk.Create(nil, "disk-test-pulumi", resources.DiskArgs{
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

		exposedDisk := resources.ExposedDisk{}

		_, exposedDiskState, err := exposedDisk.Create(nil, "edisk-test-pulumi", resources.ExposedDiskArgs{
			URL:        url,
			Token:      token,
			CustomerID: customer,
			DiskID:     state.DiskID,
		}, false)
		assert.NoError(t, err)
		assert.Equal(t, state.DiskID, exposedDiskState.DiskID)
	})
	t.Run("TestExternalNetwork", func(t *testing.T) {
		externalNetwork := resources.ExternalNetwork{}

		_, state, err := externalNetwork.Create(nil, "ext-test-pulumi", resources.ExternalNetworkArgs{
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

		updateStatus, updatedState, err := externalNetwork.Update(nil, state, resources.ExternalNetworkArgs{
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

		deleteStatus, err := externalNetwork.Delete(nil, resources.ExternalNetworkArgs{
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
		serverPool := resources.ServerPool{}

		_, state, err := serverPool.Create(nil, "sv-test-pulumi", resources.ServerPoolArgs{
			URL:          url,
			Token:        token,
			CustomerID:   customer,
			CloudSpaceID: CloudSpaceState.CloudSpaceID,
			Name:         "Pulumi_server_pool",
			Description:  "test",
		}, false)
		assert.NoError(t, err)
		assert.NotEmpty(t, state.ServerPoolID)

		updateStatus, updatedState, err := serverPool.Update(nil, state, resources.ServerPoolArgs{
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
		host := resources.Host{}

		_, hostState, err := host.Create(nil, "host-test-pulumi", resources.HostArgs{
			URL:          url,
			Token:        token,
			CustomerID:   customer,
			CloudSpaceID: CloudSpaceState.CloudSpaceID,
			ServerPoolID: ServerPoolState.ServerPoolID,
			Address:      "192.168.10.10",
		}, false)
		assert.NoError(t, err)
		assert.NotEmpty(t, hostState.HostID)

		status, err := host.Delete(nil, hostState, resources.HostArgs{
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
		loadBalancer := resources.LoadBalancer{}

		_, state, err := loadBalancer.Create(nil, "lb-test-pulumi", resources.LoadBalancerArgs{
			URL:          url,
			Token:        token,
			CustomerID:   customer,
			CloudSpaceID: CloudSpaceState.CloudSpaceID,
			Name:         "Pulumi_lb_test",
			Type:         "tcp",
			FrontEnd: resources.FrontEnd{
				Port: 23,
			},
			BackEnd: resources.BackEnd{
				ServerpoolID: ServerPoolState.ServerPoolID,
				TargetPort:   23,
			},
		}, false)
		assert.NoError(t, err)
		assert.NotEmpty(t, state.LoadBalancerID)

		updateStatus, _, err := loadBalancer.Update(nil, state, resources.LoadBalancerArgs{
			URL:          url,
			Token:        token,
			CustomerID:   customer,
			CloudSpaceID: CloudSpaceState.CloudSpaceID,
			Name:         "Pulumi_lb_test",
			Type:         "udp",
			FrontEnd: resources.FrontEnd{
				Port: 23,
			},
			BackEnd: resources.BackEnd{
				ServerpoolID: ServerPoolState.ServerPoolID,
				TargetPort:   23,
			},
		})
		assert.NoError(t, err)
		assert.Equal(t, true, updateStatus)

		status, err := loadBalancer.Delete(nil, state, resources.LoadBalancerArgs{
			URL:          url,
			Token:        token,
			CustomerID:   customer,
			CloudSpaceID: CloudSpaceState.CloudSpaceID,
		})
		assert.NoError(t, err)
		assert.Equal(t, true, status)
	})
	t.Run("TestReverseProxy", func(t *testing.T) {
		reverseProxy := resources.ReverseProxy{}

		_, state, err := reverseProxy.Create(nil, "rp-test-pulumi", resources.ReverseProxyArgs{
			URL:          url,
			Token:        token,
			CustomerID:   customer,
			CloudSpaceID: CloudSpaceState.CloudSpaceID,
			Name:         "Reverse_Proxy_Pulumi",
			ReverseProxyFrontEnd: resources.ReverseProxyFrontEnd{
				Domain:   "Pulumi",
				Scheme:   "http",
				HTTPPort: intPtr(25),
				LetsEncrypt: resources.LetsEncrypt{
					Enabled: false,
				},
			},
			ReverseProxyBackend: resources.ReverseProxyBackend{
				Scheme:       "http",
				ServerpoolID: ServerPoolState.ServerPoolID,
				TargetPort:   25,
			},
		}, false)
		assert.NoError(t, err)
		assert.NotEmpty(t, state.ReverseProxyID)

		updateStatus, updatedState, err := reverseProxy.Update(nil, state, resources.ReverseProxyArgs{
			URL:          url,
			Token:        token,
			CustomerID:   customer,
			CloudSpaceID: CloudSpaceState.CloudSpaceID,
			Name:         "Reverse_Proxy_Pulumi",
			ReverseProxyFrontEnd: resources.ReverseProxyFrontEnd{
				Domain:   "Pulumi",
				Scheme:   "http",
				HTTPPort: intPtr(25),
				LetsEncrypt: resources.LetsEncrypt{
					Enabled: false,
				},
			},
			ReverseProxyBackend: resources.ReverseProxyBackend{
				Scheme:       "http",
				ServerpoolID: ServerPoolState.ServerPoolID,
				TargetPort:   26,
			},
		})
		assert.NoError(t, err)
		assert.Equal(t, true, updateStatus)

		deleteStatus, err := reverseProxy.Delete(nil, updatedState, resources.ReverseProxyArgs{
			URL:          url,
			Token:        token,
			CustomerID:   customer,
			CloudSpaceID: CloudSpaceState.CloudSpaceID,
		})
		assert.NoError(t, err)
		assert.Equal(t, true, deleteStatus)
	})
	t.Run("TestObjectSpace", func(t *testing.T) {
		objectSpace := resources.ObjectSpace{}

		_, state, err := objectSpace.Create(nil, "obj-test-pulumi", resources.ObjectSpaceArgs{
			URL:        url,
			Token:      token,
			CustomerID: customer,
			Location:   "nl-rmd-dc01-001",
			Name:       "Pulumi_obj_space",
		}, false)
		assert.NoError(t, err)
		assert.NotEmpty(t, state.ObjectSpaceID)

		objectSpaceLink := resources.ObjectSpaceLink{}

		_, linkState, err := objectSpaceLink.Create(nil, "objlink-test-pulumi", resources.ObjectSpaceLinkArgs{
			URL:           url,
			Token:         token,
			CustomerID:    customer,
			CloudSpaceID:  CloudSpaceState.CloudSpaceID,
			ObjectSpaceID: state.ObjectSpaceID,
		}, false)
		assert.NoError(t, err)
		assert.Equal(t, true, linkState.Status)

		status, err := objectSpaceLink.Delete(nil, linkState, resources.ObjectSpaceLinkArgs{
			URL:          url,
			Token:        token,
			CustomerID:   customer,
			CloudSpaceID: CloudSpaceState.CloudSpaceID,
		})
		assert.NoError(t, err)
		assert.Equal(t, true, status)

		deleteStatus, err := objectSpace.Delete(nil, state, resources.ObjectSpaceArgs{
			URL:        url,
			Token:      token,
			CustomerID: customer,
		})
		assert.NoError(t, err)
		assert.Equal(t, true, deleteStatus)
	})
	t.Run("TestVirtualMachine", func(t *testing.T) {
		virtualMachine := resources.VirtualMachine{}

		_, state, err := virtualMachine.Create(nil, "vm-test-pulumi", resources.VirtualMachineArgs{
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
		antiAffinityVM := resources.AntiAffinityGroupVM{}

		_, state, err := antiAffinityVM.Create(nil, "vm-test-pulumi", resources.AntiAffinityGroupVMArgs{
			URL:              url,
			Token:            token,
			CustomerID:       customer,
			CloudSpaceID:     CloudSpaceState.CloudSpaceID,
			GroupID:          AntiAffinityGroupState.GroupID,
			VirtualMachineID: VirtualMachineState.VirtualMachineID,
		}, false)
		assert.NoError(t, err)
		assert.Equal(t, "RUNNING", state.Status)

		deleteStatus, err := antiAffinityVM.Delete(nil, state, resources.AntiAffinityGroupVMArgs{
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
		portForward := resources.PortForward{}

		_, state, err := portForward.Create(nil, "pf-test-pulumi", resources.PortForwardArgs{
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

		updatedState, err := portForward.Update(nil, state, resources.PortForwardArgs{
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

		statusCode, err := portForward.Delete(nil, updatedState, resources.PortForwardArgs{
			URL:          url,
			Token:        token,
			CustomerID:   customer,
			CloudSpaceID: CloudSpaceState.CloudSpaceID,
		})
		assert.NoError(t, err)
		assert.Equal(t, 200, statusCode)
	})
	t.Run("TestVirtualMachineCDrom", func(t *testing.T) {
		virtualMachineCD := resources.VirtualMachineCD{}

		_, state, err := virtualMachineCD.Create(nil, "vmcd-test-pulumi", resources.VirtualMachineCDArgs{
			URL:              url,
			Token:            token,
			CustomerID:       customer,
			CloudSpaceID:     CloudSpaceState.CloudSpaceID,
			VirtualMachineID: VirtualMachineState.VirtualMachineID,
			CdRomID:          9000,
		}, false)
		assert.NoError(t, err)
		assert.Equal(t, "CREATED", state.Status)

		deleteStatus, err := virtualMachineCD.Delete(nil, state, resources.VirtualMachineCDArgs{
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
		virtualMachineDisk := resources.VirtualMachineDisk{}

		_, state, err := virtualMachineDisk.Create(nil, "vmdsk-test-pulumi", resources.VirtualMachineDiskArgs{
			URL:              url,
			Token:            token,
			CustomerID:       customer,
			CloudSpaceID:     CloudSpaceState.CloudSpaceID,
			VirtualMachineID: VirtualMachineState.VirtualMachineID,
			DiskID:           DiskState.DiskID,
		}, false)
		assert.NoError(t, err)
		assert.Equal(t, true, state.Success)

		deleteStatus, err := virtualMachineDisk.Delete(nil, state, resources.VirtualMachineDiskArgs{
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
		VirtualMachineNIC := resources.VirtualMachineNIC{}

		_, state, err := VirtualMachineNIC.Create(nil, "vmdsk-test-pulumi", resources.VirtualMachineNICArgs{
			URL:               url,
			Token:             token,
			CustomerID:        customer,
			CloudSpaceID:      CloudSpaceState.CloudSpaceID,
			VirtualMachineID:  VirtualMachineState.VirtualMachineID,
			ExternalNetworkID: 13,
		}, false)
		assert.NoError(t, err)
		assert.NotEmpty(t, state.DeviceName)

		deleteStatus, err := VirtualMachineNIC.Delete(nil, state, resources.VirtualMachineNICArgs{
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
		serverPool := resources.ServerPool{}

		svDeleteStatus, err := serverPool.Delete(nil, ServerPoolState, resources.ServerPoolArgs{
			URL:          url,
			Token:        token,
			CustomerID:   customer,
			CloudSpaceID: CloudSpaceState.CloudSpaceID,
		})
		assert.NoError(t, err)
		assert.Equal(t, true, svDeleteStatus)

		antiAffinityGroup := resources.AntiAffinityGroup{}

		_, err = antiAffinityGroup.Delete(nil, AntiAffinityGroupState, resources.AntiAffinityGroupArgs{
			URL:          url,
			Token:        token,
			CustomerID:   customer,
			CloudSpaceID: CloudSpaceState.CloudSpaceID,
		})
		assert.NoError(t, err)

		disk := resources.Disk{}

		dskDeleteStatus, err := disk.Delete(nil, DiskState, resources.DiskArgs{
			URL:        url,
			Token:      token,
			CustomerID: customer,
			Location:   CloudSpaceState.Location,
		})
		assert.NoError(t, err)
		assert.Equal(t, true, dskDeleteStatus)

		virtualMachine := resources.VirtualMachine{}

		vmDeleteStatus, err := virtualMachine.Delete(nil, VirtualMachineState, resources.VirtualMachineArgs{
			URL:          url,
			Token:        token,
			CustomerID:   customer,
			CloudSpaceID: CloudSpaceState.CloudSpaceID,
		})
		assert.NoError(t, err)
		assert.Equal(t, true, vmDeleteStatus)

		cloudspace := resources.Cloudspace{}

		csDeleteStatus, err := cloudspace.Delete(nil, CloudSpaceState, resources.CloudspaceArgs{
			URL:        url,
			Token:      token,
			CustomerID: customer,
		})
		assert.NoError(t, err)
		assert.Equal(t, true, csDeleteStatus)
	})
}
