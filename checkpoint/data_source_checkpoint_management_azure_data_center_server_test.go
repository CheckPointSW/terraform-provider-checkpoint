package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementAzureDataCenterServer_basic(t *testing.T) {

	objName := "tfTestManagementDataAzureDataCenterServer_" + acctest.RandString(6)
	resourceName := "checkpoint_management_azure_data_center_server.azure_data_center_server"
	dataSourceName := "data.checkpoint_management_azure_data_center_server.azure_data_center_server"
	authenticationMethod := "user-authentication"
	username := "MY-KEY-ID"
	password := "MY-SECRET-KEY"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementAzureDataCenterServerConfig(objName, username, password, authenticationMethod),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementAzureDataCenterServerConfig(name string, username string, password string, authenticationMethod string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_azure_data_center_server" "azure_data_center_server" {
	name = "%s"
	username         = "%s"
	password     = "%s"
	authentication_method = "%s"
	ignore_warnings = true
}

data "checkpoint_management_azure_data_center_server" "azure_data_center_server" {
    name = "${checkpoint_management_azure_data_center_server.azure_data_center_server.name}"
}
`, name, username, password, authenticationMethod)
}
