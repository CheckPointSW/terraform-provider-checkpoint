package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementTimeGroup_basic(t *testing.T) {

	objName := "TimeGroup" + acctest.RandString(2)
	resourceName := "checkpoint_management_time_group.time_group"
	dataSourceName := "data.checkpoint_management_data_time_group.data_time_group"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementTimeGroupConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})
}

func testAccDataSourceManagementTimeGroupConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_time_group" "time_group" {
    name = "%s"
}

data "checkpoint_management_data_time_group" "data_time_group" {
    name = "${checkpoint_management_time_group.time_group.name}"
}
`, name)
}
