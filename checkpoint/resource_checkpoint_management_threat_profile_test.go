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

func TestAccCheckpointManagementThreatProfile_basic(t *testing.T) {
	var threatProfile map[string]interface{}
	resourceName := "checkpoint_management_threat_profile.test"
	objName := "tfTestManagementThreatProfile_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointThreatProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementThreatProfileConfig(objName, "high", "Critical"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointThreatProfileExists(resourceName, &threatProfile),
					testAccCheckCheckpointThreatProfileAttributes(&threatProfile, objName, "high", "Critical"),
				),
			},
		},
	})
}

func testAccCheckpointThreatProfileDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_threat_profile" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-threat-profile", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success { // Resource still exists. failed to destroy.
				return fmt.Errorf("threat profile object (%s) still exists", rs.Primary.ID)
			}
		}
		break
	}
	return nil
}

func testAccCheckCheckpointThreatProfileExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("threat profile ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)
		response, _ := client.ApiCall("show-threat-profile", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return fmt.Errorf(response.ErrorMsg)
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointThreatProfileAttributes(threatProfile *map[string]interface{}, name string, performanceImpact string, protectionsSeverity string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		threatProfile := *threatProfile
		if threatProfile == nil {
			return fmt.Errorf("threat profile is nil")
		}

		threatProfileName := threatProfile["name"].(string)
		if threatProfileName != name {
			return fmt.Errorf("name is %s, expected %s", threatProfileName, name)
		}

		performanceImpactValue := threatProfile["active-protections-performance-impact"].(string)
		if performanceImpactValue != performanceImpact {
			return fmt.Errorf("performance impact is %s, expected %s", performanceImpactValue, performanceImpact)
		}

		protectionsSeverityValue := threatProfile["active-protections-severity"].(string)
		if protectionsSeverityValue != protectionsSeverity {
			return fmt.Errorf("protections severity is %s, expected %s", protectionsSeverityValue, protectionsSeverity)
		}
		return nil
	}
}

func testAccManagementThreatProfileConfig(name string, performanceImpact string, protectionsSeverity string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_threat_profile" "test" {
	name = "%s"
	active_protections_performance_impact = "%s"
	active_protections_severity	 = "%s"
}
`, name, performanceImpact, protectionsSeverity)
}
