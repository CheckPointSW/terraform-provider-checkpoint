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

func TestAccCheckpointManagementUserTemplate_basic(t *testing.T) {

	var userTemplateMap map[string]interface{}
	resourceName := "checkpoint_management_user_template.test"
	objName := "tfTestManagementUserTemplate_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementUserTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementUserTemplateConfig(objName, "2030-05-30", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementUserTemplateExists(resourceName, &userTemplateMap),
					testAccCheckCheckpointManagementUserTemplateAttributes(&userTemplateMap, objName, "2030-05-30", false),
				),
			},
		},
	})
}

func testAccCheckpointManagementUserTemplateDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_user_template" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-user-template", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("UserTemplate object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementUserTemplateExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("UserTemplate ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-user-template", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementUserTemplateAttributes(userTemplateMap *map[string]interface{}, name string, expirationDate string, expirationByGlobalProperties bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		userTemplateName := (*userTemplateMap)["name"].(string)
		if !strings.EqualFold(userTemplateName, name) {
			return fmt.Errorf("name is %s, expected %s", name, userTemplateName)
		}
		userTemplateExpirationDate := (*userTemplateMap)["expiration-date"].(map[string]interface{})["iso-8601"].(string)
		date := strings.Split(userTemplateExpirationDate, "T")[0]
		if !strings.EqualFold(date, expirationDate) {
			return fmt.Errorf("expirationDate is %s, expected %s", expirationDate, userTemplateExpirationDate)
		}
		userTemplateExpirationByGlobalProperties := (*userTemplateMap)["expiration-by-global-properties"].(bool)
		if userTemplateExpirationByGlobalProperties != expirationByGlobalProperties {
			return fmt.Errorf("expirationByGlobalProperties is %t, expected %t", expirationByGlobalProperties, userTemplateExpirationByGlobalProperties)
		}
		return nil
	}
}

func testAccManagementUserTemplateConfig(name string, expirationDate string, expirationByGlobalProperties bool) string {
	return fmt.Sprintf(`
resource "checkpoint_management_user_template" "test" {
        name = "%s"
        expiration_date = "%s"
        expiration_by_global_properties = %t
}
`, name, expirationDate, expirationByGlobalProperties)
}
