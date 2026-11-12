package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementGroupWithExclusion_basic(t *testing.T) {

	objName := "tfTestManagementDataGroupWithExclusion_" + acctest.RandString(6)
	resourceName := "checkpoint_management_group_with_exclusion.group_with_exclusion"
	dataSourceName := "data.checkpoint_management_data_group_with_exclusion.data_group_with_exclusion"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementGroupWithExclusionConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementGroupWithExclusionConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_group" "group1" {
    name = "group1"
}

resource "checkpoint_management_group" "group2" {
    name = "group2"
}

resource "checkpoint_management_group_with_exclusion" "group_with_exclusion" {
    name = "%s"
    include = "${checkpoint_management_group.group1.name}"
    except = "${checkpoint_management_group.group2.name}"
}

data "checkpoint_management_data_group_with_exclusion" "data_group_with_exclusion" {
    name = "${checkpoint_management_group_with_exclusion.group_with_exclusion.name}"
}
`, name)
}
