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

func TestAccDataSourceCheckpointManagementNatRulebase_basic(t *testing.T) {
	var showObjectsQuery map[string]interface{}
	dataSourceShowObjects := "data.checkpoint_management_nat_rulebase.test"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementNatRulebaseConfig("Standard", 1, "Hide NAT"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointNatRulebase(dataSourceShowObjects, &showObjectsQuery),
					testAccCheckCheckpointNatRulebaseAttributes(&showObjectsQuery),
				),
			},
		},
	})
}

func testAccCheckCheckpointNatRulebase(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("show-nat-rulebase data source not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("show-nat-rulebase data source ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)
		response, _ := client.ApiCall("show-nat-rulebase", map[string]interface{}{"package": "Standard", "filter": "Hide NAT", "limit": 1}, client.GetSessionID(), true, false)
		if !response.Success {
			return fmt.Errorf(response.ErrorMsg)
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointNatRulebaseAttributes(showNatRulebaseMap *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		showNatRulebaseMap := *showNatRulebaseMap
		if showNatRulebaseMap == nil {
			return fmt.Errorf("showNatRulebaseMap is nil")
		}

		rulebase := showNatRulebaseMap["rulebase"].([]interface{})

		if len(rulebase) != 1 {
			return fmt.Errorf("show-nat-rulebase returned wrong number of rulebase objects. exptected for 1, found %d", len(rulebase))
		}

		return nil
	}
}

func testAccDataSourceManagementNatRulebaseConfig(packageName string, limit int, filter string) string {
	return fmt.Sprintf(`
data "checkpoint_management_nat_rulebase" "test" {
	package = "%s"
	filter= "%s"
	limit = %d
}
`, packageName, filter, limit)
}
