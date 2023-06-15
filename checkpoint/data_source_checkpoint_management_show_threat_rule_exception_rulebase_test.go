package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"log"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementShowThreatRuleExceptionRuleBase_basic(t *testing.T) {

	objName := "tfTestManagementDataThreatRuleExceptionRulebase" + acctest.RandString(6)
	var showObjectsQuery map[string]interface{}
	dataSourceShowObjects := "data.checkpoint_management_threat_rule_exception_rulebase.data"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementShowThreatRuleExceptionRuleBaseConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointThreatRuleExceptionRulebase(dataSourceShowObjects, &showObjectsQuery, objName),
					testAccCheckCheckpointThreatRuleExceptionRulebaseAttributes(&showObjectsQuery, objName),
				),
			},
		},
	})
}

func testAccCheckCheckpointThreatRuleExceptionRulebaseAttributes(showThreatExceptionRulebaseMap *map[string]interface{}, objName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		showAccessRulebaseMap := *showThreatExceptionRulebaseMap
		if showAccessRulebaseMap == nil {
			return fmt.Errorf("showThreatRuleExcpetionRulebaseMap is nil")
		}

		name := showAccessRulebaseMap["rulebase"].([]interface{})[0].(map[string]interface{})["rulebase"].([]interface{})[0].(map[string]interface{})["name"]

		if name != objName {
			return fmt.Errorf("rule name is %s. while expected name is %s\n", name, objName)
		}
		log.Println("rule name match.")
		return nil
	}
}

func testAccCheckCheckpointThreatRuleExceptionRulebase(resourceTfName string, res *map[string]interface{}, filter string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("show-threat-rule-exception-rulebase data source not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("show-threat-rule-exception-rulebase data source ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)
		response, _ := client.ApiCall("show-threat-rule-exception-rulebase", map[string]interface{}{"name": "Standard Threat Prevention", "rule-name": "rule1", "use-object-dictionary": "false"}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return fmt.Errorf(response.ErrorMsg)
		}

		*res = response.GetData()

		return nil
	}
}

func testAccDataSourceManagementShowThreatRuleExceptionRuleBaseConfig(objName string) string {
	return fmt.Sprintf(`

 resource "checkpoint_management_threat_exception" "threat_exception" {
  name = "%s"
  position = {top = "top"}
  exception_group_name = "Global Exceptions"
  track = "Log"
  service = ["AH", "AOL"]
 
}
data "checkpoint_management_threat_rule_exception_rulebase" "data" {
  name = "Standard Threat Prevention"
  rule_number = 1
 
}
`, objName)
}
