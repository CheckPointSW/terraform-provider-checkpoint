package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"os"
	"testing"
)

func TestAccCommandCheckpointManagementLockObject_basic(t *testing.T) {

	commandName := "checkpoint_management_command_lock_object.lock"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCommandCheckpointManagementLockObjectConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(commandName, "name", commandName, "name"),
				),
			},
		},
	})

}

func testAccCommandCheckpointManagementLockObjectConfig() string {
	return fmt.Sprintf(`
resource "checkpoint_management_command_lock_object" "lock" {
    name = "MyIntranet"
	uid = "324bf311-8cb6-4f45-a966-c50304dd2445"
	type = "vpn-community-meshed"
}
`)
}
