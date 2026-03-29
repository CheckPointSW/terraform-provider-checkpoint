package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementDataCenterQuery_basic(t *testing.T) {

	objName := "tfTestManagementDataDataCenterQuery_" + acctest.RandString(6)
	resourceName := "checkpoint_management_data_center_query.data_center_query"
	dataSourceName := "data.checkpoint_management_data_center_query.data_center_query"
	firstVal := "value1"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementDataCenterQueryConfig(objName, firstVal),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementDataCenterQueryConfig(name string, firstVal string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_data_center_query" "data_center_query" {
	name = "%s"
  	data_centers = ["All"]
  	query_rules {
    key_type = "predefined"
    key      = "name-in-data-center"
    values   = ["%s"]
	}
}

data "checkpoint_management_data_center_query" "data_center_query" {
    name = "${checkpoint_management_data_center_query.data_center_query.name}"
}
`, name, firstVal)
}
