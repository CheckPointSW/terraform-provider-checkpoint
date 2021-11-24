package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	_ "github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	_ "strings"
	"testing"
)

func TestAccDataSourceCheckpointManagementThreatRulebase_basic(t *testing.T) {
	var showObjectsQuery map[string]interface{}
	dataSourceShowObjects := "data.checkpoint_management_threat_rulebase.test"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementThreatRulebaseConfig("Standard Threat Prevention", 1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointThreatRulebase(dataSourceShowObjects, &showObjectsQuery),
					testAccCheckCheckpointThreatRulebaseAttributes(&showObjectsQuery),
				),
			},
		},
	})
}

func testAccCheckCheckpointThreatRulebase(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("show-threat-rulebase data source not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("show-threat-rulebase data source ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)
		response, _ := client.ApiCall("show-threat-rulebase", map[string]interface{}{"name": "Standard Threat Prevention", "limit": 1}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return fmt.Errorf(response.ErrorMsg)
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointThreatRulebaseAttributes(showThreatRulebaseMap *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		showThreatRulebaseMap := *showThreatRulebaseMap
		if showThreatRulebaseMap == nil {
			return fmt.Errorf("showThreatRulebaseMap is nil")
		}

		rulebase := showThreatRulebaseMap["rulebase"].([]interface{})

		if len(rulebase) != 1 {
			return fmt.Errorf("show-threat-rulebase returned wrong number of rulebase objects. exptected for 1, found %d", len(rulebase))
		}

		return nil
	}
}

func testAccDataSourceManagementThreatRulebaseConfig(name string, limit int) string {
	return fmt.Sprintf(`
data "checkpoint_management_threat_rulebase" "test" {
	name = "%s"
	limit = %d
}
`, name, limit)
}
