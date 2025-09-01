package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"os"
	"testing"
)

func TestAccCheckpointManagementSetAppControlUpdateSchedule_basic(t *testing.T) {

	commandName := "checkpoint_management_set_app_control_update_schedule.set_app_control"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementSetAppControlUpdateScheduleConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(commandName, "name", commandName, "name"),
				),
			},
		},
	})

}

func testAccManagementSetAppControlUpdateScheduleConfig() string {
	return fmt.Sprintf(`
resource "checkpoint_management_set_app_control_update_schedule" "set_app_control" {
	schedule_gateway_update {
        schedule {
            recurrence {
                pattern = "interval"
                interval_hours = 4
                interval_minutes = 30
                interval_seconds = 10
            }
        }
    }
}
`)
}
