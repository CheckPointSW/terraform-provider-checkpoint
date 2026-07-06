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

func TestAccCheckpointManagementDefSetting_basic(t *testing.T) {

	var defSettingMap map[string]interface{}
	resourceName := "checkpoint_management_def_setting.test"
	objName := "tfTestManagementDefSetting_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementDefSettingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementDefSettingConfig(objName, "boolean"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementDefSettingExists(resourceName, &defSettingMap),
					testAccCheckCheckpointManagementDefSettingAttributes(&defSettingMap, objName, "boolean"),
				),
			},
		},
	})
}

func testAccCheckpointManagementDefSettingDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_def_setting" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-def-setting", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("DefSetting object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementDefSettingExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("DefSetting ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-def-setting", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementDefSettingAttributes(defSettingMap *map[string]interface{}, name string, dataType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		defSettingName := (*defSettingMap)["name"].(string)
		if !strings.EqualFold(defSettingName, name) {
			return fmt.Errorf("name is %s, expected %s", name, defSettingName)
		}
		defSettingDataType := (*defSettingMap)["data-type"].(string)
		if !strings.EqualFold(defSettingDataType, dataType) {
			return fmt.Errorf("dataType is %s, expected %s", dataType, defSettingDataType)
		}
		return nil
	}
}

func testAccManagementDefSettingConfig(name string, dataType string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_def_setting" "example" {
  name      = "%s"
  data_type = "%s"
  assignments {
    value       = "true"
    description = "Default for Quantum gateways"
    model       = "quantum"
  }
  assignments {
    value       = "false"
    description = "Default for Spark gateways"
    model       = "spark"
  }
}
`, name, dataType)
}
