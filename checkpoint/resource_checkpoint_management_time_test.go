package checkpoint

import (
	"fmt"
	"os"
	"strings"
	"testing"

	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccCheckpointManagementTime_basic(t *testing.T) {

	var timeMap map[string]interface{}
	resourceName := "checkpoint_management_time.test"
	objName := "time_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementTimeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementTimeConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementTimeExists(resourceName, &timeMap),
					testAccCheckCheckpointManagementTimeAttributes(&timeMap, objName),
				),
			},
		},
	})
}

func testAccCheckpointManagementTimeDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_time" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-time", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success {
				return fmt.Errorf("Time object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementTimeExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Time ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-time", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementTimeAttributes(timeMap *map[string]interface{}, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		timeName := (*timeMap)["name"].(string)
		if !strings.EqualFold(timeName, name) {
			return fmt.Errorf("name is %s, expected %s", name, timeName)
		}
		return nil
	}
}

func testAccManagementTimeConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_time" "test" {
        name = "%s"
}
`, name)
}
