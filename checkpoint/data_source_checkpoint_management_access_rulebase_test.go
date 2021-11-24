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

func TestAccDataSourceCheckpointManagementAccessRulebase_basic(t *testing.T) {
	var showObjectsQuery map[string]interface{}
	dataSourceShowObjects := "data.checkpoint_management_access_rulebase.test"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementAccessRulebaseConfig("Network", 1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointAccessRulebase(dataSourceShowObjects, &showObjectsQuery),
					testAccCheckCheckpointAccessRulebaseAttributes(&showObjectsQuery),
				),
			},
		},
	})
}

func testAccCheckCheckpointAccessRulebase(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("show-access-rulebase data source not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("show-access-rulebase data source ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)
		response, _ := client.ApiCall("show-access-rulebase", map[string]interface{}{"name": "Network", "limit": 1}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return fmt.Errorf(response.ErrorMsg)
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointAccessRulebaseAttributes(showAccessRulebaseMap *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		showAccessRulebaseMap := *showAccessRulebaseMap
		if showAccessRulebaseMap == nil {
			return fmt.Errorf("showAccessRulebaseMap is nil")
		}

		rulebase := showAccessRulebaseMap["rulebase"].([]interface{})

		if len(rulebase) != 1 {
			return fmt.Errorf("show-access-rulebase returned wrong number of rulebase objects. exptected for 1, found %d", len(rulebase))
		}

		return nil
	}
}

func testAccDataSourceManagementAccessRulebaseConfig(name string, limit int) string {
	return fmt.Sprintf(`
data "checkpoint_management_access_rulebase" "test" {
	name = "%s"
	limit = %d
}
`, name, limit)
}
