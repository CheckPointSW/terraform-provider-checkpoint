package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementAccessSection_basic(t *testing.T) {

	objName := "tfTestManagementDataAccessSection_" + acctest.RandString(6)
	resourceName := "checkpoint_management_access_section.access_section"
	dataSourceName := "data.checkpoint_management_data_access_section.data_access_section"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementAccessSectionConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementAccessSectionConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_access_layer" "access_layer" {
        name = "myaccesslayer"
        detect_using_x_forward_for = false
        applications_and_url_filtering = true
}

resource "checkpoint_management_access_section" "access_section" {
    name = "%s"
	layer = "${checkpoint_management_access_layer.access_layer.name}"
	position = {top = "top"}
}

data "checkpoint_management_data_access_section" "data_access_section" {
    name = "${checkpoint_management_access_section.access_section.name}"
    layer = "${checkpoint_management_access_section.access_section.layer}"
}
`, name)
}
