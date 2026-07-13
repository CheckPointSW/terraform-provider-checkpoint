package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementScheduledEvent_basic(t *testing.T) {

	objName := "tfTestManagementDataScheduledEvent_" + acctest.RandString(6)
	resourceName := "checkpoint_management_scheduled_event.scheduled_event"
	dataSourceName := "data.checkpoint_management_data_scheduled_event.data_scheduled_event"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementScheduledEventConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementScheduledEventConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_scheduled_event" "scheduled_event" {

    name = "%s"
	type = "type_test"
	read_only = false
}

data "checkpoint_management_data_scheduled_event" "data_scheduled_event" {
    name = "${checkpoint_management_scheduled_event.scheduled_event.name}"
}
`, name)
}
