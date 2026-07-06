package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementGuideline_basic(t *testing.T) {

	objName := "tfTestManagementDataGuideline_" + acctest.RandString(6)
	resourceName := "checkpoint_management_guideline.guideline"
	dataSourceName := "data.checkpoint_management_data_guideline.data_guideline"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementGuidelineConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementGuidelineConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_guideline" "guideline" {

    name = "%s"
	type = "type_test"
	default_action = "default_action_test"
}

data "checkpoint_management_data_guideline" "data_guideline" {
    name = "${checkpoint_management_guideline.guideline.name}"
}
`, name)
}
