package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementThreatRule_basic(t *testing.T) {
	objName := "tfTestManagementDataThreatRule_" + acctest.RandString(6)
	layerName := "Standard Threat Prevention"
	resourceName := "checkpoint_management_threat_rule.test"
	dataSourceName := "data.checkpoint_management_threat_rule.test_rule"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementThreatRuleConfig(objName, layerName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "enabled", resourceName, "enabled"),
				),
			},
		},
	})
}

func testAccDataSourceManagementThreatRuleConfig(name string, layerName string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_threat_rule" "test" {
	name = "%s"
    layer = "%s"
	position = {top = "top"}
}

data "checkpoint_management_threat_rule" "test_rule" {
	layer = "${checkpoint_management_threat_rule.test.layer}"
    name = "${checkpoint_management_threat_rule.test.name}"
}
`, name, layerName)
}
