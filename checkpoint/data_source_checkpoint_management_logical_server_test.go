package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementLogicalServer_basic(t *testing.T) {
	resourceName := "checkpoint_management_logical_server.test"
	dataSourceName := "data.checkpoint_management_logical_server.data_test"
	objectName := "tfTestManagementLogicalServer_" + acctest.RandString(6)
	groupName := "tfTestManagementGroup_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementLogicalServerConfig(groupName, objectName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "ipv4_address", resourceName, "ipv4_address"),
					resource.TestCheckResourceAttrPair(dataSourceName, "ipv6_address", resourceName, "ipv6_address"),
					resource.TestCheckResourceAttrPair(dataSourceName, "server_type", resourceName, "server_type"),
				),
			},
		},
	})
}

func testAccDataSourceManagementLogicalServerConfig(groupName string, objectName string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_group" "test" {
    name = "%s"
}

resource "checkpoint_management_logical_server" "test" {
	name = "%s"
	ipv4_address = "1.1.1.1"
	server_group = "${checkpoint_management_group.test.name}"
}

data "checkpoint_management_logical_server" "data_test" {
	name = "${checkpoint_management_logical_server.test.name}"
}
`, groupName, objectName)
}
