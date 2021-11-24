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

func TestAccDataSourceCheckpointManagementHttpsRulebase_basic(t *testing.T) {
	var showObjectsQuery map[string]interface{}
	dataSourceShowObjects := "data.checkpoint_management_https_rulebase.test"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementHttpsRulebaseConfig("Default Layer", 1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointHttpsRulebase(dataSourceShowObjects, &showObjectsQuery),
					testAccCheckCheckpointHttpsRulebaseAttributes(&showObjectsQuery),
				),
			},
		},
	})
}

func testAccCheckCheckpointHttpsRulebase(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("show-https-rulebase data source not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("show-https-rulebase data source ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)
		response, _ := client.ApiCall("show-https-rulebase", map[string]interface{}{"name": "Default Layer", "limit": 1}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return fmt.Errorf(response.ErrorMsg)
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointHttpsRulebaseAttributes(showHttpsRulebaseMap *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		showHttpsRulebaseMap := *showHttpsRulebaseMap
		if showHttpsRulebaseMap == nil {
			return fmt.Errorf("showHttpsRulebaseMap is nil")
		}

		rulebase := showHttpsRulebaseMap["rulebase"].([]interface{})

		if len(rulebase) != 1 {
			return fmt.Errorf("show-https-rulebase returned wrong number of rulebase objects. exptected for 1, found %d", len(rulebase))
		}

		return nil
	}
}

func testAccDataSourceManagementHttpsRulebaseConfig(name string, limit int) string {
	return fmt.Sprintf(`
data "checkpoint_management_https_rulebase" "test" {
	name = "%s"
	limit = %d
}
`, name, limit)
}
