package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementNozomiDataCenterServer_basic(t *testing.T) {

	objName := "tfTestManagementDataNozomiDataCenterServer_" + acctest.RandString(6)
	resourceName := "checkpoint_management_nozomi_data_center_server.nozomi_data_center_server"
	dataSourceName := "data.checkpoint_management_nozomi_data_center_server.nozomi_data_center_server"
	hostname := "1.2.3.4"
	keyName := "example-key-name"
	keyToken := "example-key-token"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementNozomiDataCenterServerConfig(objName, hostname, keyName, keyToken),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementNozomiDataCenterServerConfig(name string, hostname string, keyName string, keyToken string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_nozomi_data_center_server" "nozomi_data_center_server" {
    name = "%s"
	hostname = "%s"
	key_name = "%s"
	key_token = "%s"
	ignore_warnings = true
}

data "checkpoint_management_nozomi_data_center_server" "nozomi_data_center_server" {
    name = "${checkpoint_management_nozomi_data_center_server.nozomi_data_center_server.name}"
}
`, name, hostname, keyName, keyToken)
}
