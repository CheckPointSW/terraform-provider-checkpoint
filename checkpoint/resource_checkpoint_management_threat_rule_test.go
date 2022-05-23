package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	"testing"
)

func TestAccCheckpointManagementThreatRule_basic(t *testing.T) {
	var natRule map[string]interface{}
	resourceName := "checkpoint_management_threat_rule.test"
	objName := "tfTestManagementThreatRule_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointThreatRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementThreatRuleConfig(objName, "Standard Threat Prevention"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointThreatRuleExists(resourceName, &natRule),
					testAccCheckCheckpointThreatRuleAttributes(&natRule, objName),
				),
			},
		},
	})
}

func testAccCheckpointThreatRuleDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_threat_rule" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-threat-rule", map[string]interface{}{"uid": rs.Primary.ID, "package": "Standard"}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success { // Resource still exists. failed to destroy.
				return fmt.Errorf("threat rule object (%s) still exists", rs.Primary.ID)
			}
		}
		break
	}
	return nil
}

func testAccCheckCheckpointThreatRuleExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("threat rule ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)
		response, _ := client.ApiCall("show-threat-rule", map[string]interface{}{"uid": rs.Primary.ID, "layer": "Standard Threat Prevention"}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return fmt.Errorf(response.ErrorMsg)
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointThreatRuleAttributes(threatRule *map[string]interface{}, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		threatRule := *threatRule
		if threatRule == nil {
			return fmt.Errorf("threat rule is nil")
		}

		threatRuleName := threatRule["name"].(string)
		if threatRuleName != name {
			return fmt.Errorf("name is %s, expected %s", threatRuleName, name)
		}

		return nil
	}
}

func testAccManagementThreatRuleConfig(name string, layerName string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_threat_rule" "test" {
	name = "%s"
    layer = "%s"
	position = {top = "top"}
}
`, name, layerName)
}
