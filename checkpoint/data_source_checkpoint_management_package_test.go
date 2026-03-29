package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementPackage_basic(t *testing.T) {

	objName := "tfTestManagementDataPackage_" + acctest.RandString(3)
	resourceName := "checkpoint_management_package.package"
	dataSourceName := "data.checkpoint_management_data_package.data_package"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementPackageConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})
}

func testAccDataSourceManagementPackageConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_package" "package" {
    name = "%s"
}

data "checkpoint_management_data_package" "data_package" {
    name = "${checkpoint_management_package.package.name}"
}
`, name)
}
