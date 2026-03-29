package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementTime_basic(t *testing.T) {

	objName := "time_" + acctest.RandString(6)
	resourceName := "checkpoint_management_time.time"
	dataSourceName := "data.checkpoint_management_time.data_time"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementTimeConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementTimeConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_time" "time" {
        name = "%s"
}

data "checkpoint_management_time" "data_time" {
    name = "${checkpoint_management_time.time.name}"
}
`, name)
}
