package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"os"
	"strings"
	"testing"
)

func TestAccCheckpointManagementFirewallBestPractice_basic(t *testing.T) {

	var firewallBestPracticeMap map[string]interface{}
	resourceName := "checkpoint_management_firewall_best_practice.test"
	objName := "tfTestManagementFirewallBestPractice_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementFirewallBestPracticeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementFirewallBestPracticeConfig(objName, "define a clean-up rule at the end of the policy.", "checks that the rule base ends with a clean-up rule.", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementFirewallBestPracticeExists(resourceName, &firewallBestPracticeMap),
					testAccCheckCheckpointManagementFirewallBestPracticeAttributes(&firewallBestPracticeMap, objName, "define a clean-up rule at the end of the policy.", "checks that the rule base ends with a clean-up rule.", true),
				),
			},
		},
	})
}

func testAccCheckpointManagementFirewallBestPracticeDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_firewall_best_practice" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-firewall-best-practice", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("FirewallBestPractice object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementFirewallBestPracticeExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("FirewallBestPractice ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-firewall-best-practice", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementFirewallBestPracticeAttributes(firewallBestPracticeMap *map[string]interface{}, name string, actionItem string, description string, enabled bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		firewallBestPracticeName := (*firewallBestPracticeMap)["name"].(string)
		if !strings.EqualFold(firewallBestPracticeName, name) {
			return fmt.Errorf("name is %s, expected %s", name, firewallBestPracticeName)
		}
		firewallBestPracticeActionItem := (*firewallBestPracticeMap)["action-item"].(string)
		if !strings.EqualFold(firewallBestPracticeActionItem, actionItem) {
			return fmt.Errorf("actionItem is %s, expected %s", actionItem, firewallBestPracticeActionItem)
		}
		firewallBestPracticeDescription := (*firewallBestPracticeMap)["description"].(string)
		if !strings.EqualFold(firewallBestPracticeDescription, description) {
			return fmt.Errorf("description is %s, expected %s", description, firewallBestPracticeDescription)
		}
		firewallBestPracticeEnabled := (*firewallBestPracticeMap)["enabled"].(bool)
		if firewallBestPracticeEnabled != enabled {
			return fmt.Errorf("enabled is %t, expected %t", enabled, firewallBestPracticeEnabled)
		}
		return nil
	}
}

func testAccManagementFirewallBestPracticeConfig(name string, actionItem string, description string, enabled bool) string {
	return fmt.Sprintf(`
resource "checkpoint_management_firewall_best_practice" "test" {
        name = "%s"
        action_item = "%s"
        description = "%s"
        enabled = %t
}
`, name, actionItem, description, enabled)
}
