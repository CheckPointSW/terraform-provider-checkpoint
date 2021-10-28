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

func TestAccCheckpointManagementThreatException_basic(t *testing.T) {
	var natRule map[string]interface{}
	resourceName := "checkpoint_management_threat_exception.test"
	objName := "tfTestManagementThreatException_" + acctest.RandString(6)
	layerName := "Standard Threat Prevention"
	threatRuleName := "threatRule"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointThreatExceptionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementThreatExceptionConfig(objName, layerName, threatRuleName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointThreatExceptionExists(resourceName, &natRule, threatRuleName, layerName),
					testAccCheckCheckpointThreatExceptionAttributes(&natRule, objName),
				),
			},
		},
	})
}

func testAccCheckpointThreatExceptionDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	layerName := "Standard Threat Prevention"
	threatRuleName := "threatRule"
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_threat_exception" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-threat-rule", map[string]interface{}{"uid": rs.Primary.ID, "layer": layerName, "rule-name": threatRuleName}, client.GetSessionID(), true, false)
			if res.Success { // Resource still exists. failed to destroy.
				return fmt.Errorf("threat rule object (%s) still exists", rs.Primary.ID)
			}
		}
		break
	}
	return nil
}

func testAccCheckCheckpointThreatExceptionExists(resourceTfName string, res *map[string]interface{}, ruleName string, layerName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("threat exception ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)
		response, _ := client.ApiCall("show-threat-exception", map[string]interface{}{"uid": rs.Primary.ID, "layer": layerName, "rule-name": ruleName}, client.GetSessionID(), true, false)
		if !response.Success {
			return fmt.Errorf(response.ErrorMsg)
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointThreatExceptionAttributes(exceptionRule *map[string]interface{}, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		exceptionRule := *exceptionRule
		if exceptionRule == nil {
			return fmt.Errorf("exception rule is nil")
		}

		exceptionRuleName := exceptionRule["name"].(string)
		if exceptionRuleName != name {
			return fmt.Errorf("name is %s, expected %s", exceptionRuleName, name)
		}

		return nil
	}
}

func testAccManagementThreatExceptionConfig(name string, layerName string, threatRuleName string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_threat_rule" "test" {
	name = "%s"
    layer = "%s"
	position = {top = "top"}
}

resource "checkpoint_management_threat_exception" "test" {
	name = "%s"
    layer = "%s"
	position = {top = "top"}
	rule_name = "${checkpoint_management_threat_rule.test.name}"
}
`, threatRuleName, layerName, name, layerName)
}
