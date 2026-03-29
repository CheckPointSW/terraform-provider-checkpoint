package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementGenericDataCenterServer_basic(t *testing.T) {

	objName := "tfTestManagementDataGenericDataCenterServer_" + acctest.RandString(6)
	resourceName := "checkpoint_management_generic_data_center_server.generic_data_center_server"
	dataSourceName := "data.checkpoint_management_generic_data_center_server.generic_data_center_server"
	url := "MY_URL"
	interval := "60"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementGenericDataCenterServerConfig(objName, url, interval),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementGenericDataCenterServerConfig(name string, url string, interval string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_generic_data_center_server" "generic_data_center_server" {
    name = "%s"
	url = "%s"
	interval = "%s"
	ignore_warnings = true
}

data "checkpoint_management_generic_data_center_server" "generic_data_center_server" {
    name = "${checkpoint_management_generic_data_center_server.generic_data_center_server.name}"
}
`, name, url, interval)
}
