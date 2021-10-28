package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementThreatException_basic(t *testing.T) {
	objName := "tfTestManagementDataThreatException_" + acctest.RandString(6)
	layerName := "Standard Threat Prevention"
	threatRuleName := "threatRule"
	resourceName := "checkpoint_management_threat_exception.threat_exception"
	dataSourceName := "data.checkpoint_management_threat_exception.data_threat_exception"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementThreatExceptionConfig(objName, layerName, threatRuleName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "enabled", resourceName, "enabled"),
				),
			},
		},
	})
}

func testAccDataSourceManagementThreatExceptionConfig(name string, layerName string, threatRuleName string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_threat_rule" "threat_rule" {
	name = "%s"
    layer = "%s"
	position = {top = "top"}
}

resource "checkpoint_management_threat_exception" "threat_exception" {
	name = "%s"
    layer = "%s"
	position = {top = "top"}
	rule_name = "${checkpoint_management_threat_rule.threat_rule.name}"
}

data "checkpoint_management_threat_exception" "data_threat_exception" {
	name = "${checkpoint_management_threat_exception.threat_exception.name}"
    layer = "${checkpoint_management_threat_exception.threat_exception.layer}"
	rule_name = "${checkpoint_management_threat_exception.threat_exception.rule_name}"
}
`, threatRuleName, layerName, name, layerName)
}
