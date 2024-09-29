package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	"strings"
	"testing"
)

func TestAccCheckpointManagementMobileAccessProfileRule_basic(t *testing.T) {

	var mobileAccessProfileRuleMap map[string]interface{}
	resourceName := "checkpoint_management_mobile_access_profile_rule.test"
	objName := "tfTestManagementMobileAccessProfileRule_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementMobileAccessProfileRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementMobileAccessProfileRuleConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementMobileAccessProfileRuleExists(resourceName, &mobileAccessProfileRuleMap),
					testAccCheckCheckpointManagementMobileAccessProfileRuleAttributes(&mobileAccessProfileRuleMap, objName),
				),
			},
		},
	})
}

func testAccCheckpointManagementMobileAccessProfileRuleDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_mobile_access_profile_rule" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-mobile-access-profile-rule", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success {
				return fmt.Errorf("MobileAccessProfileRule object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementMobileAccessProfileRuleExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("MobileAccessProfileRule ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-mobile-access-profile-rule", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementMobileAccessProfileRuleAttributes(mobileAccessProfileRuleMap *map[string]interface{}, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		mobileAccessProfileRuleName := (*mobileAccessProfileRuleMap)["name"].(string)
		if !strings.EqualFold(mobileAccessProfileRuleName, name) {
			return fmt.Errorf("name is %s, expected %s", name, mobileAccessProfileRuleName)
		}

		userGroupsJson := (*mobileAccessProfileRuleMap)["user-groups"].([]interface{})
		var userGroupsIds = make([]string, 0)
		if len(userGroupsJson) > 0 {
			for _, userGroups := range userGroupsJson {
				userGroupsTry1, ok := userGroups.(map[string]interface{})
				if ok {
					userGroupsIds = append([]string{userGroupsTry1["name"].(string)}, userGroupsIds...)
				} else {
					userGroupsTry2 := userGroups.(string)
					userGroupsIds = append([]string{userGroupsTry2}, userGroupsIds...)
				}
			}
		}

		return nil
	}
}

func testAccManagementMobileAccessProfileRuleConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_mobile_access_profile_rule" "test" {
        name = "%s"
        position = {top = "top"}
       
        
}
`, name)
}
