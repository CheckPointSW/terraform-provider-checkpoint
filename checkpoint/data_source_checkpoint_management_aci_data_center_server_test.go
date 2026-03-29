package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementAciDataCenterServer_basic(t *testing.T) {

	objName := "tfTestManagementDataAciDataCenterServer_" + acctest.RandString(6)
	resourceName := "checkpoint_management_aci_data_center_server.aci_data_center_server"
	dataSourceName := "data.checkpoint_management_aci_data_center_server.aci_data_center_server"
	username := "USERNAME"
	password := "PASSWORD"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementAciDataCenterServerConfig(objName, username, password),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementAciDataCenterServerConfig(name string, username string, password string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_aci_data_center_server" "aci_data_center_server" {
    name = "%s"
	username = "%s"
	password = "%s"
	urls = ["url1", "url2"]
    unsafe_auto_accept = true
	ignore_warnings = true
}

data "checkpoint_management_aci_data_center_server" "aci_data_center_server" {
    name = "${checkpoint_management_aci_data_center_server.aci_data_center_server.name}"
}
`, name, username, password)
}
