package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementRegulation_basic(t *testing.T) {

	objName := "tfTestManagementDataRegulation_" + acctest.RandString(6)
	resourceName := "checkpoint_management_regulation.regulation"
	dataSourceName := "data.checkpoint_management_data_regulation.data_regulation"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementRegulationConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementRegulationConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_regulation" "regulation" {

    name = "%s"
	type = "type_test"
	enabled = false
}

data "checkpoint_management_data_regulation" "data_regulation" {
    name = "${checkpoint_management_regulation.regulation.name}"
}
`, name)
}
