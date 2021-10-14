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

func TestAccCheckpointManagementNatRule_basic(t *testing.T) {
	var natRule map[string]interface{}
	resourceName := "checkpoint_management_nat_rule.test"
	objName := "tfTestManagementNatRule_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointNatRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementNatRuleConfig(objName, "Standard"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointNatRuleExists(resourceName, &natRule),
					testAccCheckCheckpointNatRuleAttributes(&natRule, objName),
				),
			},
		},
	})
}

func testAccCheckpointNatRuleDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_nat_rule" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-nat-rule", map[string]interface{}{"uid": rs.Primary.ID, "package": "Standard"}, client.GetSessionID(), true, false)
			if res.Success { // Resource still exists. failed to destroy.
				return fmt.Errorf("nat rule object (%s) still exists", rs.Primary.ID)
			}
		}
		break
	}
	return nil
}

func testAccCheckCheckpointNatRuleExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("access rule ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)
		response, _ := client.ApiCall("show-nat-rule", map[string]interface{}{"uid": rs.Primary.ID, "package": "Standard"}, client.GetSessionID(), true, false)
		if !response.Success {
			return fmt.Errorf(response.ErrorMsg)
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointNatRuleAttributes(natRule *map[string]interface{}, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		natRule := *natRule
		if natRule == nil {
			return fmt.Errorf("nat rule is nil")
		}

		natRuleName := natRule["name"].(string)
		if natRuleName != name {
			return fmt.Errorf("name is %s, expected %s", natRuleName, name)
		}

		return nil
	}
}

func testAccManagementNatRuleConfig(name string, packageName string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_nat_rule" "test" {
	name = "%s"
    package = "%s"
	position = {top = "top"}
}
`, name, packageName)
}
