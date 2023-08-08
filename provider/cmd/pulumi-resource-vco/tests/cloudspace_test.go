package tests

import (
	"fmt"
	"github.com/fabianv-cloud/pulumi-vco-native/provider/cmd/pulumi-resource-vco/resources"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

var CsState resources.CloudspaceState
var csMachine resources.VirtualMachineState

func TestCloudspace(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		return
	}
	customer := os.Getenv("CUSTOMER")
	token := os.Getenv("TOKEN")

	t.Run("Create", func(t *testing.T) {
		cloudspace := resources.Cloudspace{}

		_, state, err := cloudspace.Create(nil, "cs-test-pulumi", resources.CloudspaceArgs{
			Token:             token,
			CustomerID:        customer,
			Name:              "Pulumi",
			Location:          "nl-rmd-dc01-001",
			PrivateNetwork:    "192.168.10.0/24",
			ExternalNetworkID: 13,
			Private:           false,
			VcpuQuota:         intPtr(-1),
			VdiskSpaceQuota:   intPtr(-1),
			MemoryQuota:       intPtr(-1),
			PublicIPQuota:     intPtr(-1),
			FirewallCustom: &resources.FirewallCustom{
				ImageID:  intPtr(343),
				Type:     stringPtr("BSD"),
				DiskSize: intPtr(250),
				Memory:   intPtr(2048),
				Vcpus:    intPtr(4),
			},
		}, false)
		assert.NoError(t, err)
		assert.Equal(t, "DEPLOYED", state.Status)
	})

	t.Run("Delete", func(t *testing.T) {
		url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/vms", customer, CsState.CloudSpaceID)

		client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		defer resp.Body.Close()

		var result map[string]interface{}

		var vmID int
		if vms, ok := result["result"].([]interface{}); ok && len(vms) > 0 {
			if vm, ok := vms[0].(map[string]interface{}); ok {
				if id, ok := vm["vm_id"].(float64); ok {
					vmID = int(id)
				}
			}
		}

		csMachine.VirtualMachineID = vmID

		virtualMachine := resources.VirtualMachine{}

		vmDeleteStatus, err := virtualMachine.Delete(nil, csMachine, resources.VirtualMachineArgs{
			Token:        token,
			CustomerID:   customer,
			CloudSpaceID: CsState.CloudSpaceID,
		})
		assert.NoError(t, err)
		assert.Equal(t, true, vmDeleteStatus)

		cloudspace := resources.Cloudspace{}

		csDeleteStatus, err := cloudspace.Delete(nil, CsState, resources.CloudspaceArgs{
			Token:      token,
			CustomerID: customer,
		})
		assert.NoError(t, err)
		assert.Equal(t, true, csDeleteStatus)
	})
}
