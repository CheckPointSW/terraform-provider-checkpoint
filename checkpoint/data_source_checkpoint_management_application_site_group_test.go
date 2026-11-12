package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementApplicationSiteGroup_basic(t *testing.T) {

	objName := "tfTestManagementDataHost_" + acctest.RandString(6)
	resourceName := "checkpoint_management_application_site_group.application_site_group"
	dataSourceName := "data.checkpoint_management_data_application_site_group.data_application_site_group"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementApplicationSiteGroupConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})
}

func testAccDataSourceManagementApplicationSiteGroupConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_application_site_group" "application_site_group" {
    name = "%s"
}

data "checkpoint_management_data_application_site_group" "data_application_site_group" {
    name = "${checkpoint_management_application_site_group.application_site_group.name}"
}
`, name)
}
