package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementVMwareDataCenterServer_basic(t *testing.T) {

	objName := "tfTestManagementDataVMwareDataCenterServer_" + acctest.RandString(6)
	resourceName := "checkpoint_management_vmware_data_center_server.vmware_data_center_server"
	dataSourceName := "data.checkpoint_management_vmware_data_center_server.vmware_data_center_server"
	vmType := "vcenter"
	username := "USERNAME"
	password := "PASSWORD"
	hostname := "HOSTNAME"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementVMwareDataCenterServerConfig(objName, vmType, username, password, hostname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementVMwareDataCenterServerConfig(name string, vmType string, username string, password string, hostname string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_vmware_data_center_server" "vmware_data_center_server" {
    name = "%s"
	type = "%s"
	username = "%s"
	password = "%s"
	hostname = "%s"
    unsafe_auto_accept = true
	ignore_warnings = true
}

data "checkpoint_management_vmware_data_center_server" "vmware_data_center_server" {
    name = "${checkpoint_management_vmware_data_center_server.vmware_data_center_server.name}"
}
`, name, vmType, username, password, hostname)
}
