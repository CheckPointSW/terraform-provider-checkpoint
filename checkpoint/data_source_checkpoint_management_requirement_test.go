package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementRequirement_basic(t *testing.T) {

	objName := "tfTestManagementDataRequirement_" + acctest.RandString(6)
	resourceName := "checkpoint_management_requirement.requirement"
	dataSourceName := "data.checkpoint_management_data_requirement.data_requirement"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementRequirementConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementRequirementConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_requirement" "requirement" {

    name = "%s"
	type = "type_test"
	score_level = "score_level_test"
}

data "checkpoint_management_data_requirement" "data_requirement" {
    name = "${checkpoint_management_requirement.requirement.name}"
}
`, name)
}
