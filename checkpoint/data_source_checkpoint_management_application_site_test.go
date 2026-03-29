package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementApplicationSite_basic(t *testing.T) {

	objName := "tfTestManagementDataApplicationSite_" + acctest.RandString(6)
	resourceName := "checkpoint_management_application_site.application_site"
	dataSourceName := "data.checkpoint_management_data_application_site.data_application_site"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementApplicationSiteConfig(objName, "Social Networking", "www.cnet.com"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "url_list", resourceName, "url_list"),
				),
			},
		},
	})

}

func testAccDataSourceManagementApplicationSiteConfig(name string, primaryCategory string, urlList1 string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_application_site" "application_site" {
        name = "%s"
        primary_category = "%s"
        url_list = ["%s"]
}

data "checkpoint_management_data_application_site" "data_application_site" {
    name = "${checkpoint_management_application_site.application_site.name}"
}
`, name, primaryCategory, urlList1)
}
