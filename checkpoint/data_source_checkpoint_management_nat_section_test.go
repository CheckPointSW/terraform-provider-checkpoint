package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementNatSection_basic(t *testing.T) {
	objName := "tfTestManagementDataNatSection_" + acctest.RandString(6)
	packageName := "Standard"
	resourceName := "checkpoint_management_nat_section.test"
	dataSourceName := "data.checkpoint_management_nat_section.test_section"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementNatSectionConfig(objName, packageName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})
}

func testAccDataSourceManagementNatSectionConfig(name string, packageName string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_nat_section" "test" {
	name = "%s"
    package = "%s"
	position = {top = "top"}
}

data "checkpoint_management_nat_section" "test_section" {
	package = "${checkpoint_management_nat_section.test.package}"
    name = "${checkpoint_management_nat_section.test.name}"
}
`, name, packageName)
}
