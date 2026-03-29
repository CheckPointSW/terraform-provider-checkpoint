package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccCheckpointManagementImportManagement_basic(t *testing.T) {

	commandName := "checkpoint_management_command_import_management.import"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementImportManagementConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(commandName, "file_path", commandName, "file_path"),
				),
			},
		},
	})

}

func testAccManagementImportManagementConfig() string {
	return fmt.Sprintf(`
resource "checkpoint_management_command_export_management" "export" {
  file_path = "/var/log/domian1_backup.tgz"
}

resource "checkpoint_management_command_import_management" "import" {
	file_path = "${checkpoint_management_command_export_management.export.file_path}"
}
`)
}
