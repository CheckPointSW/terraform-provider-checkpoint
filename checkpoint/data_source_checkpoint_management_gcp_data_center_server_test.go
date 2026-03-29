package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementGcpDataCenterServer_basic(t *testing.T) {

	objName := "tfTestManagementDataGcpDataCenterServer_" + acctest.RandString(6)
	resourceName := "checkpoint_management_gcp_data_center_server.gcp_data_center_server"
	dataSourceName := "data.checkpoint_management_gcp_data_center_server.gcp_data_center_server"
	authenticationMethod := "key-authentication"
	privateKey := "MYKEY.json"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementGcpDataCenterServerConfig(objName, authenticationMethod, privateKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementGcpDataCenterServerConfig(name string, authenticationMethod string, privateKey string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_gcp_data_center_server" "gcp_data_center_server" {
	name = "%s"
	authentication_method = "%s"
	private_key = "%s"
	ignore_warnings = true
}

data "checkpoint_management_gcp_data_center_server" "gcp_data_center_server" {
    name = "${checkpoint_management_gcp_data_center_server.gcp_data_center_server.name}"
}
`, name, authenticationMethod, privateKey)
}
