package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementDynamicObject_basic(t *testing.T) {

	objName := "tfTestManagementDataDynamicObject_" + acctest.RandString(6)
	resourceName := "checkpoint_management_dynamic_object.dynamic_object"
	dataSourceName := "data.checkpoint_management_data_dynamic_object.data_dynamic_object"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementDynamicObjectConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})
}

func testAccDataSourceManagementDynamicObjectConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_dynamic_object" "dynamic_object" {
    name = "%s"
}

data "checkpoint_management_data_dynamic_object" "data_dynamic_object" {
    name = "${checkpoint_management_dynamic_object.dynamic_object.name}"
}
`, name)
}
