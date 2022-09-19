package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"os"
	"testing"
)

func TestAccCheckpointManagementExportSmartTask_basic(t *testing.T) {

	commandName := "checkpoint_management_command_export_management.export"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementExportSmartTaskConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(commandName, "name", commandName, "name"),
				),
			},
		},
	})

}

func testAccManagementExportSmartTaskConfig() string {
	return fmt.Sprintf(`
resource "checkpoint_management_smart_task" "smart" {
	name = "dummy"
	enabled = true
	trigger = "After Install Policy"
	action = {
		send_web_request = {
			url = "https://demo.example.com/policy-installation-reports"
			fingerprint = "3FDD902286DBF130EF4CEC7939EF81060AB0FEB6"
		} 	
	}
}

resource "checkpoint_management_command_export_smart_task" "export" {
	name = "dummy"
}
`)
}
