package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"os"
	"strings"
	"testing"
)

func TestAccCheckpointManagementOverrideCategorization_basic(t *testing.T) {

	var overrideCategorizationMap map[string]interface{}
	resourceName := "checkpoint_management_override_categorization.test"
	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementOverrideCategorizationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementOverrideCategorizationConfig("adam3", "Botnets", "low"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementOverrideCategorizationExists(resourceName, &overrideCategorizationMap),
					testAccCheckCheckpointManagementOverrideCategorizationAttributes(&overrideCategorizationMap, "adam3", "Botnets", "low"),
				),
			},
		},
	})
}

func testAccCheckpointManagementOverrideCategorizationDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_override_categorization" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-override-categorization", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success {
				return fmt.Errorf("OverrideCategorization object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementOverrideCategorizationExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("OverrideCategorization ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-override-categorization", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementOverrideCategorizationAttributes(overrideCategorizationMap *map[string]interface{}, url string, newPrimaryCategory string, risk string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		overrideCategorizationUrl := (*overrideCategorizationMap)["url"].(string)
		if !strings.EqualFold(overrideCategorizationUrl, url) {
			return fmt.Errorf("url is %s, expected %s", url, overrideCategorizationUrl)
		}
		/**overrideCategorizationNewPrimaryCategory := (*overrideCategorizationMap)["new-primary-category"].(string)
		if !strings.EqualFold(overrideCategorizationNewPrimaryCategory, newPrimaryCategory) {
			return fmt.Errorf("newPrimaryCategory is %s, expected %s", newPrimaryCategory, overrideCategorizationNewPrimaryCategory)
		}*/
		overrideCategorizationRisk := (*overrideCategorizationMap)["risk"].(string)
		if !strings.EqualFold(overrideCategorizationRisk, risk) {
			return fmt.Errorf("risk is %s, expected %s", risk, overrideCategorizationRisk)
		}
		return nil
	}
}

func testAccManagementOverrideCategorizationConfig(url string, newPrimaryCategory string, risk string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_override_categorization" "test" {
        url = "%s"
        new_primary_category = "%s"
        risk = "%s"
}
`, url, newPrimaryCategory, risk)
}
