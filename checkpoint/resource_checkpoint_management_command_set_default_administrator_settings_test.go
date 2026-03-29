package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccCheckpointManagementSetDefaultAdministratorSettings_basic(t *testing.T) {

	commandName := "checkpoint_management_set_default_administrator_settings.command_set_default_administrator_settings"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementSetDefaultAdministratorSettingsConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(commandName, "expiration_type", commandName, "expiration_type"),
					resource.TestCheckResourceAttrPair(commandName, "expiration_date", commandName, "expiration_date"),
					resource.TestCheckResourceAttrPair(commandName, "indicate_expiration_in_admin_view", commandName, "indicate_expiration_in_admin_view"),
					resource.TestCheckResourceAttrPair(commandName, "notify_expiration_to_admin", commandName, "notify_expiration_to_admin"),
					resource.TestCheckResourceAttrPair(commandName, "days_to_notify_expiration_to_admin", commandName, "days_to_notify_expiration_to_admin"),
				),
			},
		},
	})

}

func testAccManagementSetDefaultAdministratorSettingsConfig() string {
	return fmt.Sprintf(`
resource "checkpoint_management_set_default_administrator_settings" "command_set_default_administrator_settings" {
  expiration_type = "expiration date"
  expiration_date = "2025-06-23"
  indicate_expiration_in_admin_view = false
  notify_expiration_to_admin = true
  days_to_notify_expiration_to_admin = 5
}
`)
}
