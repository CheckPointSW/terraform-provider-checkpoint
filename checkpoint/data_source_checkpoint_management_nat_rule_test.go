package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementNatRule_basic(t *testing.T) {
	objName := "tfTestManagementDataNatRule_" + acctest.RandString(6)
	packageName := "Standard"
	resourceName := "checkpoint_management_nat_rule.test"
	dataSourceName := "data.checkpoint_management_nat_rule.test_rule"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementNatRuleConfig(objName, packageName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "enabled", resourceName, "enabled"),
				),
			},
		},
	})
}

func testAccDataSourceManagementNatRuleConfig(name string, packageName string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_nat_rule" "test" {
	name = "%s"
    package = "%s"
	position = {top = "top"}
}

data "checkpoint_management_nat_rule" "test_rule" {
	package = "${checkpoint_management_nat_rule.test.package}"
    name = "${checkpoint_management_nat_rule.test.name}"
}
`, name, packageName)
}
