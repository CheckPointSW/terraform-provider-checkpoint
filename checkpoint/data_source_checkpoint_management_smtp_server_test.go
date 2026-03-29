package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementSmtpServer_basic(t *testing.T) {

	objName := "tfTestManagementDataSmtpServer_" + acctest.RandString(6)
	resourceName := "checkpoint_management_smtp_server.smtp_server"
	dataSourceName := "data.checkpoint_management_smtp_server.data_smtp_server"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementSmtpServerConfig(objName, "smtp.example.com", "25"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "server", resourceName, "server"),
					resource.TestCheckResourceAttrPair(dataSourceName, "port", resourceName, "port"),
				),
			},
		},
	})

}

func testAccDataSourceManagementSmtpServerConfig(name string, server string, port string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_smtp_server" "smtp_server" {
    name = "%s"
	server = "%s"
	port = "%s"
}

data "checkpoint_management_smtp_server" "data_smtp_server" {
    name = "${checkpoint_management_smtp_server.smtp_server.name}"
}
`, name, server, port)
}
