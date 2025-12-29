package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"os"
	"testing"
)

func TestAccCheckpointManagementSetSyncWithUserCenter_basic(t *testing.T) {

	commandName := "checkpoint_management_set_sync_with_user_center.set_sync_with_user_center"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementSetSyncWithUserCenterConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(commandName, "name", commandName, "name"),
				),
			},
		},
	})

}

func testAccManagementSetSyncWithUserCenterConfig() string {
	return fmt.Sprintf(`
resource "checkpoint_management_set_sync_with_user_center" "set_sync_with_user_center" {
  enabled = true
}
`)
}
