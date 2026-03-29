package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementGroup_basic(t *testing.T) {

	objName := "tfTestManagementDataGroup_" + acctest.RandString(6)
	resourceName := "checkpoint_management_group.group"
	dataSourceName := "data.checkpoint_management_data_group.data_group"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementGroupConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementGroupConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_group" "group" {
    name = "%s"
}

data "checkpoint_management_data_group" "data_group" {
    name = "${checkpoint_management_group.group.name}"
}
`, name)
}
