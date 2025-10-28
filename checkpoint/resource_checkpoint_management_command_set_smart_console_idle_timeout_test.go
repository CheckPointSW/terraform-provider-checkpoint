package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"os"
	"testing"
)

func TestAccCheckpointManagementSmartConsoleIdleTimeout_basic(t *testing.T) {

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
				Config: testAccManagementSmartConsoleIdleTimeoutConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(commandName, "timeout_duration", "30"),
					resource.TestCheckResourceAttr(commandName, "enabled", "true"),
				),
			},
		},
	})

}

func testAccManagementSmartConsoleIdleTimeoutConfig() string {
	return fmt.Sprintf(`
resource "checkpoint_management_command_set_smart_console_idle_timeout" "command_set_smart_console_idle_timeout" {
	enabled = true
	timeout_duration = 30
}
`)
}
