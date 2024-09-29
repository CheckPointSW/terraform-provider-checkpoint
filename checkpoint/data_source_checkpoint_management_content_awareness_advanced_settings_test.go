package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementContentAwarenessAdvancedSettings_basic(t *testing.T) {

	dataSourceName := "data.checkpoint_management_content_awareness_advanced_settings.content"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementContentAwarenessAdvancedSettingsConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", dataSourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementContentAwarenessAdvancedSettingsConfig() string {
	return fmt.Sprintf(`
data "checkpoint_management_content_awareness_advanced_settings" "content" {

}
`)
}
