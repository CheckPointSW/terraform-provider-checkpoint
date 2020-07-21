package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementHttpsSection_basic(t *testing.T) {

	objName := "tfTestManagementDataHttpsSection_" + acctest.RandString(6)
	resourceName := "checkpoint_management_https_section.https_section"
	dataSourceName := "data.checkpoint_management_data_https_section.data_https_section"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementHttpsSectionConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementHttpsSectionConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_https_section" "https_section" {
        name = "%s"
		layer = "Default Layer"
        position = {top = "top"}
}

data "checkpoint_management_data_https_section" "data_https_section" {
    name = "${checkpoint_management_https_section.https_section.name}"
    layer = "${checkpoint_management_https_section.https_section.layer}"
}
`, name)
}
