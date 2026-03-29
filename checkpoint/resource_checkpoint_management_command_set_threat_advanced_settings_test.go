package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccCheckpointManagementSetThreatAdvancedSettings_basic(t *testing.T) {

	commandName := "checkpoint_management_command_set_threat_advanced_settings.set"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementSetThreatAdvancedSettingsConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(commandName, "name", commandName, "name"),
				),
			},
		},
	})

}

func testAccManagementSetThreatAdvancedSettingsConfig() string {
	return fmt.Sprintf(`
resource "checkpoint_management_command_set_threat_advanced_settings" "set" {
  httpi_non_standard_ports = false
  resource_classification {
    mode = "background"
    custom_settings  {
      anti_bot = "hold"
      anti_virus = "background"
      zero_phishing = "hold"
    }
    web_service_fail_mode = "allow"
  }
  feed_retrieving_interval = "00:05"
  internal_error_fail_mode = "allow connections"
  log_unification_timeout = 600
}
`)
}
