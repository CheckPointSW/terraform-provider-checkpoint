package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementGuardicoreDataCenterServer_basic(t *testing.T) {

	objName := "tfTestManagementDataGuardicoreDataCenterServer_" + acctest.RandString(6)
	resourceName := "checkpoint_management_guardicore_data_center_server.guardicore_data_center_server"
	dataSourceName := "data.checkpoint_management_guardicore_data_center_server.guardicore_data_center_server"
	hostname := "1.2.3.4"
	username := "example-username"
	password := "example-password"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementGuardicoreDataCenterServerConfig(objName, hostname, username, password),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementGuardicoreDataCenterServerConfig(name string, hostname string, username string, password string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_guardicore_data_center_server" "guardicore_data_center_server" {
    name = "%s"
	hostname = "%s"
	username = "%s"
	password = "%s"
	ignore_warnings = true
}

data "checkpoint_management_guardicore_data_center_server" "guardicore_data_center_server" {
    name = "${checkpoint_management_guardicore_data_center_server.guardicore_data_center_server.name}"
}
`, name, hostname, username, password)
}
