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

func TestAccCheckpointManagementSmartTask_basic(t *testing.T) {

	var smartTaskMap map[string]interface{}
	resourceName := "checkpoint_management_smart_task.smart_task"
	objName := "tfTestManagementSmartTask_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementSmartTaskDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementSmartTaskConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementSmartTaskExists(resourceName, &smartTaskMap),
					testAccCheckCheckpointManagementSmartTaskAttributes(&smartTaskMap, objName),
				),
			},
		},
	})
}

func testAccCheckpointManagementSmartTaskDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_smart_task" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-smart-task", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success {
				return fmt.Errorf("SmartTask object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementSmartTaskExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("SmartTask ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-smart-task", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementSmartTaskAttributes(smartTaskMap *map[string]interface{}, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		smartTaskName := (*smartTaskMap)["name"].(string)
		if !strings.EqualFold(smartTaskName, name) {
			return fmt.Errorf("name is %s, expected %s", name, smartTaskName)
		}
		return nil
	}
}

func testAccManagementSmartTaskConfig(name string) string {
	return fmt.Sprintf(`
  resource "checkpoint_management_smart_task" "smart_task" {

  name = "%s"
  trigger = "Before Publish"
  description = "my smart task"
  action {

    send_web_request {
      url            = "https://demo.example.com/policy-installation-reports"
      fingerprint    = "8023a5652ba2c8f5b0902363a5314cd2b4fdbc5c"
      override_proxy = true
      proxy_url      = "https://demo.example.com/policy-installation-reports"
      time_out       = 200
      shared_secret  = " secret"
    }
  }
  enabled = true
}
`, name)
}
