package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementApplicationSiteCategory_basic(t *testing.T) {

	objName := "tfTestManagementDataApplicationSiteCategory_" + acctest.RandString(6)
	resourceName := "checkpoint_management_application_site_category.application_site_category"
	dataSourceName := "data.checkpoint_management_data_application_site_category.data_application_site_category"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementApplicationSiteCategoryConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementApplicationSiteCategoryConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_application_site_category" "application_site_category" {
    name = "%s"
}

data "checkpoint_management_data_application_site_category" "data_application_site_category" {
    name = "${checkpoint_management_application_site_category.application_site_category.name}"
}
`, name)
}
