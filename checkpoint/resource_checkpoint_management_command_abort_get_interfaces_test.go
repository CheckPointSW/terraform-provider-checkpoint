package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccCheckpointManagementAbortGetInterfaces_basic(t *testing.T) {

	commandName := "checkpoint_management_command_abort_get_interfaces.abort"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementAbortGetInterfacesConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(commandName, "name", commandName, "name"),
				),
			},
		},
	})

}

func testAccManagementAbortGetInterfacesConfig() string {
	return fmt.Sprintf(`
resource "checkpoint_management_command_get_interfaces" "get" {
    target_name = "my_gateway"
}

resource "checkpoint_management_command_abort_get_interfaces" "abort" {
	task_id = "${checkpoint_management_command_get_interfaces.get.task_id}"
}
`)
}
