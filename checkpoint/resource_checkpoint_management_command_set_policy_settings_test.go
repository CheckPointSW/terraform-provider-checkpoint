package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"os"
	"testing"
)

func TestAccCheckpointManagementSetPolicySettings_basic(t *testing.T) {

	commandName := "checkpoint_management_command_set_policy_settings.command_set_policy_settings"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementSetPolicySettingsConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(commandName, "name", commandName, "name"),
				),
			},
		},
	})

}

func testAccManagementSetPolicySettingsConfig() string {
	return fmt.Sprintf(`
resource "checkpoint_management_command_set_policy_settings" "command_set_policy_settings" {
    last_in_cell = "none"
	none_object_behavior = "none"
	security_access_defaults = {
		source = "none"
		destination = "none"
		service = "none"
	}
}
`)
}
