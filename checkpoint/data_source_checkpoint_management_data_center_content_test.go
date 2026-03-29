package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementDataCenterContent_basic(t *testing.T) {

	objName := "myApic"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementDataCenterContentConfig(objName),
			},
		},
	})

}

func testAccDataSourceManagementDataCenterContentConfig(name string) string {
	return fmt.Sprintf(`
data "checkpoint_management_data_center_content" "data_center_content" {
    data_center_name = "%s"
}
`, name)
}
