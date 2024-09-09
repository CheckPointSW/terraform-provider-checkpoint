package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementMobileAccessSection_basic(t *testing.T) {
	objName := "tfTestManagementMobileAccessSection_" + acctest.RandString(6)
	resourceName := "checkpoint_management_mobile_access_section.test"
	dataSourceName := "data.checkpoint_management_mobile_access_section.data"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementMobileAccessSectionConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})
}

func testAccDataSourceManagementMobileAccessSectionConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_mobile_access_section" "test" {
  name = "%s"
  position = {top = "top"}

}
data "checkpoint_management_mobile_access_section" "data" {
  uid = "${checkpoint_management_mobile_access_section.test.id}"
}

`, name)
}
