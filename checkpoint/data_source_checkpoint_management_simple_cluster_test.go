package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementSimpleCluster_basic(t *testing.T) {
	objName := "mysimplecluster_" + acctest.RandString(4)
	resourceName := "checkpoint_management_simple_cluster.test"
	dataSourceName := "data.checkpoint_management_simple_cluster.simple_cluster"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementSimpleClusterConfig(objName, "1.2.3.4"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "ipv4_address", resourceName, "ipv4_address"),
				),
			},
		},
	})
}

func testAccDataSourceManagementSimpleClusterConfig(name string, ipv4 string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_simple_cluster" "test" {
	name = "%s"
	ipv4_address = "%s"
	version = "R81"
	hardware = "Open server"
	send_logs_to_server = ["ice-main-take-392"]
	firewall = true
}

data "checkpoint_management_simple_cluster" "simple_cluster" {
	name = "${checkpoint_management_simple_cluster.test.name}"
}
`, name, ipv4)
}
