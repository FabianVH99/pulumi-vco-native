package testdata

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
	"strconv"
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
		vcoCloudspace := base.Cloudspace{}

		_, state, err := vcoCloudspace.Create(nil, "cs-test-pulumi", base.CloudspaceArgs{
			URL:               url,
			Token:             token,
			CustomerID:        customer,
			Name:              "Pulumi_cloudspace",
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
		assert.Equal(t, 2, updatedState.Spread)
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
			URL:          url,
			Token:        token,
			CustomerID:   customer,
			CloudSpaceID: CloudSpaceState.CloudSpaceID,
			DiskID:       state.DiskID,
		}, false)
		assert.NoError(t, err)
		assert.Equal(t, state.DiskID, exposedDiskState.DiskID)

		err = exposedDisk.Delete(nil, "edisk-test-pulumi", exposedDiskState)
		assert.NoError(t, err)
	})

	t.Run("TestExternalNetwork", func(t *testing.T) {
		externalNetwork := cloudspace.ExternalNetwork{}

		_, state, err := externalNetwork.Create(nil, "ext-test-pulumi", cloudspace.ExternalNetworkArgs{
			URL:                 url,
			Token:               token,
			CustomerID:          customer,
			CloudSpaceID:        CloudSpaceState.CloudSpaceID,
			ExternalNetworkID:   strconv.Itoa(CloudSpaceState.ExternalNetworkID),
			ExternalNetworkType: "external",
			ExternalNetworkIP:   "",
			Metric:              503,
		}, false)
		assert.NoError(t, err)
		assert.Equal(t, strconv.Itoa(CloudSpaceState.ExternalNetworkID), state.ExternalNetworkID)

		updatedState, err := externalNetwork.Update(nil, "ext-test-pulumi", state, cloudspace.ExternalNetworkArgs{
			URL:               url,
			Token:             token,
			CustomerID:        customer,
			CloudSpaceID:      CloudSpaceState.CloudSpaceID,
			ExternalNetworkID: strconv.Itoa(CloudSpaceState.ExternalNetworkID),
			ExternalNetworkIP: "185.151.60.222/24",
			Metric:            502,
		}, false)
		assert.NoError(t, err)
		assert.Equal(t, 502, updatedState.Metric)

		err = externalNetwork.Delete(nil, "ext-test-pulumi", updatedState)
		assert.NoError(t, err)
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

		updatedState, err := serverPool.Update(nil, "sv-test-pulumi", state, ingress.ServerPoolArgs{
			URL:          url,
			Token:        token,
			CustomerID:   customer,
			CloudSpaceID: CloudSpaceState.CloudSpaceID,
			Name:         "Pulumi_server_pool_update",
			Description:  "Update",
		}, false)
		assert.NoError(t, err)
		assert.Equal(t, "Update", updatedState.Description)

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

		err = host.Delete(nil, "host-test-pulumi", hostState)
		assert.NoError(t, err)
	})

	t.Run("TestLoadBalancer", func(t *testing.T) {
		loadBalancer := ingress.LoadBalancer{}

		_, state, err := loadBalancer.Create(nil, "lb-test-pulumi", ingress.LoadBalancerArgs{
			URL:          url,
			Token:        token,
			CustomerID:   url,
			CloudSpaceID: CloudSpaceState.CloudSpaceID,
			Name:         "Pulumi_lb_test",
			Type:         "tcp",
			Port:         23,
			ServerpoolID: ServerPoolState.ServerPoolID,
			TargetPort:   23,
		}, false)
		assert.NoError(t, err)
		assert.NotEmpty(t, state.LoadBalancerID)

		updatedState, err := loadBalancer.Update(nil, "lb-test-pulumi", state, ingress.LoadBalancerArgs{
			URL:          url,
			Token:        token,
			CustomerID:   url,
			CloudSpaceID: CloudSpaceState.CloudSpaceID,
			Name:         "Pulumi_lb_test",
			Type:         "udp",
			Port:         23,
			ServerpoolID: ServerPoolState.ServerPoolID,
			TargetPort:   23,
		}, false)
		assert.NoError(t, err)
		assert.Equal(t, "udp", updatedState.Type)

		err = loadBalancer.Delete(nil, "lb-test-pulumi", state)
		assert.NoError(t, err)
	})

	t.Run("TestReverseProxy", func(t *testing.T) {
		reverseProxy := ingress.ReverseProxy{}

		_, state, err := reverseProxy.Create(nil, "rp-test-pulumi", ingress.ReverseProxyArgs{
			URL:            url,
			Token:          token,
			CustomerID:     url,
			CloudSpaceID:   CloudSpaceState.CloudSpaceID,
			Name:           "Reverse_Proxy_Pulumi",
			Domain:         "Pulumi",
			FrontEndScheme: "http",
			HTTPPort:       intPtr(25),
			Enabled:        false,
			BackendScheme:  "http",
			ServerpoolID:   ServerPoolState.ServerPoolID,
			TargetPort:     25,
		}, false)
		assert.NoError(t, err)
		assert.NotEmpty(t, state.ReverseProxyID)

		updatedState, err := reverseProxy.Update(nil, "rp-test-pulumi", state, ingress.ReverseProxyArgs{
			URL:            url,
			Token:          token,
			CustomerID:     url,
			CloudSpaceID:   CloudSpaceState.CloudSpaceID,
			Name:           "Reverse_Proxy_Pulumi",
			Domain:         "Pulumi",
			FrontEndScheme: "http",
			HTTPPort:       intPtr(25),
			Enabled:        false,
			BackendScheme:  "http",
			ServerpoolID:   ServerPoolState.ServerPoolID,
			TargetPort:     26,
		}, false)
		assert.NoError(t, err)
		assert.Equal(t, 26, updatedState.BackEnd.TargetPort)

		err = reverseProxy.Delete(nil, "rp-test-pulumi", updatedState)
		assert.NoError(t, err)
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
		assert.Equal(t, state.ObjectSpaceID, linkState.ObjectSpaceID)

		err = objectSpaceLink.Delete(nil, "objlink-test-pulumi", linkState)
		assert.NoError(t, err)

		err = objectSpace.Delete(nil, "obj-test-pulumi", state)
		assert.NoError(t, err)
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

		err = antiAffinityVM.Delete(nil, "vm-test-pulumi", state)
		assert.NoError(t, err)
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

		updatedState, err := portForward.Update(nil, "pf-test-pulumi", state, cloudspace.PortForwardArgs{
			URL:              url,
			Token:            token,
			CustomerID:       customer,
			CloudSpaceID:     CloudSpaceState.CloudSpaceID,
			Protocol:         "udp",
			PublicPort:       22,
			LocalPort:        22,
			VirtualMachineID: VirtualMachineState.VirtualMachineID,
		}, false)
		assert.NoError(t, err)
		assert.Equal(t, "udp", updatedState.Protocol)

		err = portForward.Delete(nil, "pf-test-pulumi", updatedState)
		assert.NoError(t, err)
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

		err = virtualMachineCD.Delete(nil, "vmcd-test-pulumi", state)
		assert.NoError(t, err)
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
		assert.Equal(t, DiskState.DiskID, state.DiskID)

		err = virtualMachineDisk.Delete(nil, "vmdsk-test-pulumi", state)
		assert.NoError(t, err)
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

		err = VirtualMachineNIC.Delete(nil, "vmdsk-test-pulumi", state)
		assert.NoError(t, err)
	})
	t.Run("CleanUp", func(t *testing.T) {
		serverPool := ingress.ServerPool{}

		err := serverPool.Delete(nil, "sv-test-pulumi", ServerPoolState)
		assert.NoError(t, err)

		antiAffinityGroup := cloudspace.AntiAffinityGroup{}

		err = antiAffinityGroup.Delete(nil, "ag-test-pulumi", AntiAffinityGroupState)
		assert.NoError(t, err)

		vcoDisk := base.Disk{}

		err = vcoDisk.Delete(nil, "disk-test-pulumi", DiskState)
		assert.NoError(t, err)

		virtualMachine := cloudspace.VirtualMachine{}

		err = virtualMachine.Delete(nil, "vm-test-pulumi", VirtualMachineState)
		assert.NoError(t, err)
		assert.NoError(t, err)

		vcoCloudspace := base.Cloudspace{}

		err = vcoCloudspace.Delete(nil, "cs-test-pulumi", CloudSpaceState)
		assert.NoError(t, err)
	})
}
