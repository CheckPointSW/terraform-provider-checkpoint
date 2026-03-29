package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccCheckpointManagementLoginRestrictions_basic(t *testing.T) {

	commandName := "checkpoint_management_command_set_smart_console_idle_timeout.command_set_smart_console_idle_timeout"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementLoginRestrictionsConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(commandName, "lockout_admin_account", "true"),
					resource.TestCheckResourceAttr(commandName, "failed_authentication_attempts", "10"),
					resource.TestCheckResourceAttr(commandName, "unlock_admin_account", "false"),
					resource.TestCheckResourceAttr(commandName, "lockout_duration", "30"),
					resource.TestCheckResourceAttr(commandName, "display_access_denied_message", "false"),
				),
			},
		},
	})

}

func testAccManagementLoginRestrictionsConfig() string {
	return fmt.Sprintf(`
resource "checkpoint_management_set_login_restrictions" "test" {
  lockout_admin_account = true
  failed_authentication_attempts = 10
  unlock_admin_account = false
  lockout_duration = 30
  display_access_denied_message = false
}
`)
}
