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

func TestAccCheckpointManagementAccessRule_basic(t *testing.T) {
	var accessRule map[string]interface{}
	resourceName := "checkpoint_management_access_rule.test"
	objName := "tfTestManagementAccessRule_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointAccessRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementAccessRuleConfig(objName, "Network"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointAccessRuleExists(resourceName, &accessRule),
					testAccCheckCheckpointAccessRuleAttributes(&accessRule, objName, "Network"),
				),
			},
		},
	})
}

func testAccCheckpointAccessRuleDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_access_rule" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-access-rule", map[string]interface{}{"uid": rs.Primary.ID, "layer": "Network"}, client.GetSessionID(), true, false)
			if res.Success { // Resource still exists. failed to destroy.
				return fmt.Errorf("access rule object (%s) still exists", rs.Primary.ID)
			}
		}
		break
	}
	return nil
}

func testAccCheckCheckpointAccessRuleExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("access rule ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)
		response, _ := client.ApiCall("show-access-rule", map[string]interface{}{"uid": rs.Primary.ID, "layer": "Network"}, client.GetSessionID(), true, false)
		if !response.Success {
			return fmt.Errorf(response.ErrorMsg)
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointAccessRuleAttributes(accessRule *map[string]interface{}, name string, layer string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		accessRule := *accessRule
		if accessRule == nil {
			return fmt.Errorf("access rule is nil")
		}

		accessRuleLayerUid := accessRule["layer"].(string)

		client := testAccProvider.Meta().(*checkpoint.ApiClient)
		response, _ := client.ApiCall("show-access-layer", map[string]interface{}{"uid": accessRuleLayerUid}, client.GetSessionID(), true, false)
		if !response.Success {
			return fmt.Errorf(response.ErrorMsg)
		}

		accessRuleLayerName := response.GetData()["name"]
		if accessRuleLayerName != layer {
			return fmt.Errorf("layer is %s, expected %s", accessRuleLayerName, layer)
		}
		accessRuleName := accessRule["name"].(string)
		if accessRuleName != name {
			return fmt.Errorf("name is %s, expected %s", accessRuleName, name)
		}

		return nil
	}
}

func testAccManagementAccessRuleConfig(name string, layer string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_access_rule" "test" {
	name = "%s"
    layer = "%s"
	position = {top = "top"}
}
`, name, layer)
}
