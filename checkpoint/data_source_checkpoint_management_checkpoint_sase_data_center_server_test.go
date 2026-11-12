package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementCheckpointSaseDataCenterServer_basic(t *testing.T) {

	objName := "tfTestManagementDataCheckpointSaseDataCenterServer_" + acctest.RandString(6)
	resourceName := "checkpoint_management_checkpoint_sase_data_center_server.checkpoint_sase_data_center_server"
	dataSourceName := "data.checkpoint_management_checkpoint_sase_data_center_server.checkpoint_sase_data_center_server"
	connectTo := "connected-portal"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementCheckpointSaseDataCenterServerConfig(objName, connectTo),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementCheckpointSaseDataCenterServerConfig(name string, connectTo string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_checkpoint_sase_data_center_server" "checkpoint_sase_data_center_server" {
    name = "%s"
	connect_to = "%s"
	ignore_warnings = true
}

data "checkpoint_management_checkpoint_sase_data_center_server" "checkpoint_sase_data_center_server" {
    name = "${checkpoint_management_checkpoint_sase_data_center_server.checkpoint_sase_data_center_server.name}"
}
`, name, connectTo)
}
