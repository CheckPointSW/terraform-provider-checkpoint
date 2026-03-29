package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementTag_basic(t *testing.T) {

	objName := "tfTestManagementDataAccessLayer_" + acctest.RandString(6)
	resourceName := "checkpoint_management_tag.tag"
	dataSourceName := "data.checkpoint_management_tag.data_tag"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementTagConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementTagConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_tag" "tag" {
    name = "%s"
	tags = ["tag1", "tag2"]
}

data "checkpoint_management_tag" "data_tag" {
    name = "${checkpoint_management_tag.tag.name}"
}
`, name)
}
