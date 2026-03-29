package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementTacacsGroup_basic(t *testing.T) {

	objName := "tfTestManagementDataTacacsGroup_" + acctest.RandString(6)
	resourceName := "checkpoint_management_tacacs_group.tacacs_group"
	dataSourceName := "data.checkpoint_management_tacacs_group.data_tacacs_group"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementTacacsGroupConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementTacacsGroupConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_host" "t_host" {
	name = "tacacs_host"
	ipv4_address = "212.122.122.212"
}

resource "checkpoint_management_tacacs_server" "tacacs_server" {
	name = "tacacs_example"
	server = "${checkpoint_management_host.t_host.name}"
}

resource "checkpoint_management_tacacs_group" "tacacs_group" {
    name = "%s"
	members = ["${checkpoint_management_tacacs_server.tacacs_server.name}"]
}

data "checkpoint_management_tacacs_group" "data_tacacs_group" {
    name = "${checkpoint_management_tacacs_group.tacacs_group.name}"
}
`, name)
}
