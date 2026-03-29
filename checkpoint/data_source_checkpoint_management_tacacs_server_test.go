package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementTacacsServer_basic(t *testing.T) {

	objName := "tfTestManagementDataTacacsServer_" + acctest.RandString(6)
	resourceName := "checkpoint_management_tacacs_server.tacacs_server"
	dataSourceName := "data.checkpoint_management_tacacs_server.data_tacacs_server"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementTacacsServerConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementTacacsServerConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_tacacs_server" "tacacs_server" {
    name = "%s"
	server = "yoni"
}

data "checkpoint_management_tacacs_server" "data_tacacs_server" {
    name = "${checkpoint_management_tacacs_server.tacacs_server.name}"
}
`, name)
}
