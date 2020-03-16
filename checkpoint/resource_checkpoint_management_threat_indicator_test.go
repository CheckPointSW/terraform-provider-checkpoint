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

func TestAccCheckpointManagementThreatIndicator_basic(t *testing.T) {
	var threatIndicator map[string]interface{}
	resourceName := "checkpoint_management_threat_indicator.test"
	objName := "tfTestManagementThreatIndicator_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointThreatIndicatorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementThreatIndicatorConfig(objName, "observable1", "1.2.3.4"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointThreatIndicatorExists(resourceName, &threatIndicator),
					testAccCheckCheckpointThreatIndicatorAttributes(&threatIndicator, objName),
				),
			},
		},
	})
}

func testAccCheckpointThreatIndicatorDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_threat_indicator" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-threat-indicator", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success { // Resource still exists. failed to destroy.
				return fmt.Errorf("threat indicator object (%s) still exists", rs.Primary.ID)
			}
		}
		break
	}
	return nil
}

func testAccCheckCheckpointThreatIndicatorExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("threat indicator ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)
		response, _ := client.ApiCall("show-threat-indicator", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return fmt.Errorf(response.ErrorMsg)
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointThreatIndicatorAttributes(threatIndicator *map[string]interface{}, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		threatIndicator := *threatIndicator
		if threatIndicator == nil {
			return fmt.Errorf("threat indicator is nil")
		}

		threatIndicatorName := threatIndicator["name"].(string)
		if threatIndicatorName != name {
			return fmt.Errorf("name is %s, expected %s", threatIndicatorName, name)
		}

		return nil
	}
}

func testAccManagementThreatIndicatorConfig(name string, observableName string, ipAddress string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_threat_indicator" "test" {
	name = "%s"
    observables {
    name = "%s"
    ip_address = "%s"
  	}
	ignore_warnings = true
}
`, name, observableName, ipAddress)
}
