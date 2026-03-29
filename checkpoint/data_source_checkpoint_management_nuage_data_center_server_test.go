package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementNuageDataCenterServer_basic(t *testing.T) {

	objName := "tfTestManagementDataNuageDataCenterServer_" + acctest.RandString(6)
	resourceName := "checkpoint_management_nuage_data_center_server.nuage_data_center_server"
	dataSourceName := "data.checkpoint_management_nuage_data_center_server.nuage_data_center_server"
	username := "USERNAME"
	password := "PASSWORD"
	hostname := "MY_HOSTNAME"
	organization := "MY_ORG"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementNuageDataCenterServerConfig(objName, username, password, hostname, organization),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementNuageDataCenterServerConfig(name string, username string, password string, hostname string, organization string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_nuage_data_center_server" "nuage_data_center_server" {
    name = "%s"
	username = "%s"
	password = "%s"
	hostname = "%s"
	organization = "%s"
    unsafe_auto_accept = true
	ignore_warnings = true
}

data "checkpoint_management_nuage_data_center_server" "nuage_data_center_server" {
    name = "${checkpoint_management_nuage_data_center_server.nuage_data_center_server.name}"
}
`, name, username, password, hostname, organization)
}
