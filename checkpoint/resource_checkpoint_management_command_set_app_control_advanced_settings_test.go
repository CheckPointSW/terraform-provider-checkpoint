package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"os"
	"testing"
)

func TestAccCheckpointManagementSetAppControlSettings_basic(t *testing.T) {

	commandName := "checkpoint_management_app_control_advanced_settings.set_app_control"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementSetAppControlSettingsConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(commandName, "inspect_archives", commandName, "inspect_archives"),
				),
			},
		},
	})

}

func testAccManagementSetAppControlSettingsConfig() string {
	return fmt.Sprintf(`
resource "checkpoint_management_app_control_advanced_settings" "set_app_control" {
  url_filtering_settings = {
    categorize_cached_and_translated_pages = "true"
    categorize_https_websites = "false"
    enforce_safe_search ="true"
  }
  custom_categorization_settings = {
    social_network_widgets_mode = "hold"
    url_filtering_mode = "background"
  }
  web_browsing_services = ["https","AH"]
  match_application_on_any_port = "false"
}
`)
}
