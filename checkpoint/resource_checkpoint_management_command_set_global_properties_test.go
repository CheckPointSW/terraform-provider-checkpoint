package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccCheckpointManagementSetGlobalProperties_basic(t *testing.T) {

	commandName := "checkpoint_management_command_set_global_properties.set_global"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementSetGlobalPropertiesConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(commandName, "id"),
				),
			},
		},
	})

}

func testAccManagementSetGlobalPropertiesConfig() string {
	return fmt.Sprintf(`
resource "checkpoint_management_command_set_global_properties" "set_global" {
  hit_count {
    enable_hit_count = true
  }
  data_access_control {
    auto_download_important_data = true
  }
}
`)
}
