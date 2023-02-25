package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementIpsUpdateSchedule_basic(t *testing.T) {

	dataSourceName := "data.checkpoint_management_ips_update_schedule.ips_update"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementIpsUpdateScheduleConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", dataSourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementIpsUpdateScheduleConfig() string {
	return fmt.Sprintf(`
data "checkpoint_management_ips_update_schedule" "ips_update" {

}
`)
}

