package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementMdPermissionsProfile_basic(t *testing.T) {

	objName := "tfTestManagementDataMdPermissionsProfile_" + acctest.RandString(6)
	resourceName := "checkpoint_management_md_permissions_profile.md_permissions_profile"
	dataSourceName := "data.checkpoint_management_md_permissions_profile.data_md_permissions_profile"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementMdPermissionsProfileConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementMdPermissionsProfileConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_md_permissions_profile" "md_permissions_profile" {
	name = "%s"
}

data "checkpoint_management_md_permissions_profile" "data_md_permissions_profile" {
    name = "${checkpoint_management_md_permissions_profile.md_permissions_profile.name}"
}
`, name)
}
