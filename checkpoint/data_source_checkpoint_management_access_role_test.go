package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementAccessRole_basic(t *testing.T) {

	objName := "tfTestManagementDataAccessRole_" + acctest.RandString(6)
	resourceName := "checkpoint_management_access_role.access_role"
	dataSourceName := "data.checkpoint_management_data_access_role.data_access_role"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementAccessRoleConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementAccessRoleConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_access_role" "access_role" {
    name = "%s"
	machines {
    selection = ["any"]
    source = "any"
	}
	users {
	selection = ["any"]
	source = "any"
	}
}

data "checkpoint_management_data_access_role" "data_access_role" {
    name = "${checkpoint_management_access_role.access_role.name}"
}
`, name)
}
