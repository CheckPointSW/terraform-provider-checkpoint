package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementOpenStackDataCenterServer_basic(t *testing.T) {

	objName := "tfTestManagementDataOpenStackDataCenterServer_" + acctest.RandString(6)
	resourceName := "checkpoint_management_openstack_data_center_server.openstack_data_center_server"
	dataSourceName := "data.checkpoint_management_openstack_data_center_server.openstack_data_center_server"
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
				Config: testAccDataSourceManagementOpenStackDataCenterServerConfig(objName, username, password, hostname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementOpenStackDataCenterServerConfig(name string, username string, password string, hostname string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_openstack_data_center_server" "openstack_data_center_server" {
    name = "%s"
	username = "%s"
	password = "%s"
	hostname = "%s"
    unsafe_auto_accept = true
	ignore_warnings = true
}

data "checkpoint_management_openstack_data_center_server" "openstack_data_center_server" {
    name = "${checkpoint_management_openstack_data_center_server.openstack_data_center_server.name}"
}
`, name, username, password, hostname)
}
