package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementSimpleGateway_basic(t *testing.T) {
	objName := "mysimplegw_" + acctest.RandString(4)
	resourceName := "checkpoint_management_simple_gateway.test"
	dataSourceName := "data.checkpoint_management_simple_gateway.simple_gateway"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementSimpleGatewayConfig(objName, "1.2.3.4"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "ipv4_address", resourceName, "ipv4_address"),
				),
			},
		},
	})
}

func testAccDataSourceManagementSimpleGatewayConfig(name string, ipv4 string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_simple_gateway" "test" {
	name = "%s"
	ipv4_address = "%s"
	version = "R81"
	send_logs_to_server = ["ice-main-take-392"]
}

data "checkpoint_management_simple_gateway" "simple_gateway" {
	name = "${checkpoint_management_simple_gateway.test.name}"
}
`, name, ipv4)
}
