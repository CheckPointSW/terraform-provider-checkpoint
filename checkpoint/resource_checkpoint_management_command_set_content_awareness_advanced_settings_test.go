package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"os"
	"testing"
)

func TestAccCheckpointManagementSetContentAwarenessSettings_basic(t *testing.T) {

	commandName := "checkpoint_management_content_awareness_advanced_settings.set_content_awarness"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementSetContentAwarenessSettingsConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(commandName, "inspect_archives", commandName, "inspect_archives"),
				),
			},
		},
	})

}

func testAccManagementSetContentAwarenessSettingsConfig() string {
	return fmt.Sprintf(`
resource "checkpoint_management_content_awareness_advanced_settings" "set_content_awarness" {
  inspect_archives = "false"
}
`)
}
