package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementSyncWIthUserCenter_basic(t *testing.T) {

	dataSourceName := "data.checkpoint_management_sync_with_user_center.data"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementSyncWIthUserCenterConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", dataSourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementSyncWIthUserCenterConfig() string {
	return fmt.Sprintf(`
resource "checkpoint_management_set_sync_with_user_center" "set_sync_with_user_center" {
  enabled = true
}

data "checkpoint_management_sync_with_user_center" "data" {
}
`)
}
