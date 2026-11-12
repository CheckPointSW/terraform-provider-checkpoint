package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementSecuridServer_basic(t *testing.T) {

	objName := "tfTestManagementDataSecuridServer_" + acctest.RandString(6)
	resourceName := "checkpoint_management_securid_server.test"
	dataSourceName := "data.checkpoint_management_securid_server.data_securid_server"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementSecuridServerConfig(objName, "configfile", "Q0xJRU5UX0lQPSAxLjEuMS4xMQ=="),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "config_file_name", resourceName, "config_file_name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "base64_config_file_content", resourceName, "base64_config_file_content"),
				),
			},
		},
	})

}

func testAccDataSourceManagementSecuridServerConfig(name string, configFileName string, base64ConfigFileContent string) string {
	return fmt.Sprintf(`

resource "checkpoint_management_securid_server" "test" {
        name = "%s"
        config_file_name = "%s"
        base64_config_file_content = "%s"
}

data "checkpoint_management_securid_server" "data_securid_server" {
  name = "${checkpoint_management_securid_server.test.name}"
}
`, name, configFileName, base64ConfigFileContent)
}
