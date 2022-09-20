package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementSmartTaskTrigger_basic(t *testing.T) {
	dataSourceName := "data.checkpoint_management_smart_task_trigger.smart"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementSmartTaskTriggerConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", dataSourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementSmartTaskTriggerConfig() string {
	return fmt.Sprintf(`
data "checkpoint_management_smart_task_trigger" "smart" {
    name = "After Approve"
	uid = "73cd9767-2b02-4752-9f0c-8e51859c3fe7"
}
`)
}
