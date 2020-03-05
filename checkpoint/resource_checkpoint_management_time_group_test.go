package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform/helper/resource"
    "github.com/hashicorp/terraform/terraform"
    "os"
    "strings"
    "testing"
    "github.com/hashicorp/terraform/helper/acctest"
)

func TestAccCheckpointManagementTimeGroup_basic(t *testing.T) {

    var timeGroupMap map[string]interface{}
    resourceName := "checkpoint_management_time_group.test"
    objName := acctest.RandString(6)

    context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
        CheckDestroy: testAccCheckpointManagementTimeGroupDestroy,
        Steps: []resource.TestStep{
            {
                Config: testAccManagementTimeGroupConfig(objName, ),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckCheckpointManagementTimeGroupExists(resourceName, &timeGroupMap),
                    testAccCheckCheckpointManagementTimeGroupAttributes(&timeGroupMap, objName, ),
                ),
            },
        },
    })
}

func testAccCheckpointManagementTimeGroupDestroy(s *terraform.State) error {

    client := testAccProvider.Meta().(*checkpoint.ApiClient)
    for _, rs := range s.RootModule().Resources {
        if rs.Type != "checkpoint_management_time_group" {
            continue
        }
        if rs.Primary.ID != "" {
            res, _ := client.ApiCall("show-time-group", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
            if res.Success {
                return fmt.Errorf("TimeGroup object (%s) still exists", rs.Primary.ID)
            }
        }
        return nil
    }
    return nil
}

func testAccCheckCheckpointManagementTimeGroupExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        rs, ok := s.RootModule().Resources[resourceTfName]
        if !ok {
            return fmt.Errorf("Resource not found: %s", resourceTfName)
        }

        if rs.Primary.ID == "" {
            return fmt.Errorf("TimeGroup ID is not set")
        }

        client := testAccProvider.Meta().(*checkpoint.ApiClient)

        response, err := client.ApiCall("show-time-group", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
        if !response.Success {
            return err
        }

        *res = response.GetData()

        return nil
    }
}

func testAccCheckCheckpointManagementTimeGroupAttributes(timeGroupMap *map[string]interface{}, name string) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        timeGroupName := (*timeGroupMap)["name"].(string)
        if !strings.EqualFold(timeGroupName, name) {
            return fmt.Errorf("name is %s, expected %s", name, timeGroupName)
        }
        return nil
    }
}

func testAccManagementTimeGroupConfig(name string) string {
    return fmt.Sprintf(`
resource "checkpoint_management_time_group" "test" {
        name = "%s"
}
`, name)
}

