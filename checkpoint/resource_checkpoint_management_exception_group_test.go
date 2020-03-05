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

func TestAccCheckpointManagementExceptionGroup_basic(t *testing.T) {

    var exceptionGroupMap map[string]interface{}
    resourceName := "checkpoint_management_exception_group.test"
    objName := "tfTestManagementExceptionGroup_" + acctest.RandString(6)

    context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
        CheckDestroy: testAccCheckpointManagementExceptionGroupDestroy,
        Steps: []resource.TestStep{
            {
                Config: testAccManagementExceptionGroupConfig(objName, "manually-select-threat-rules"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckCheckpointManagementExceptionGroupExists(resourceName, &exceptionGroupMap),
                    testAccCheckCheckpointManagementExceptionGroupAttributes(&exceptionGroupMap, objName, "manually-select-threat-rules"),
                ),
            },
        },
    })
}

func testAccCheckpointManagementExceptionGroupDestroy(s *terraform.State) error {

    client := testAccProvider.Meta().(*checkpoint.ApiClient)
    for _, rs := range s.RootModule().Resources {
        if rs.Type != "checkpoint_management_exception_group" {
            continue
        }
        if rs.Primary.ID != "" {
            res, _ := client.ApiCall("show-exception-group", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
            if res.Success {
                return fmt.Errorf("ExceptionGroup object (%s) still exists", rs.Primary.ID)
            }
        }
        return nil
    }
    return nil
}

func testAccCheckCheckpointManagementExceptionGroupExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        rs, ok := s.RootModule().Resources[resourceTfName]
        if !ok {
            return fmt.Errorf("Resource not found: %s", resourceTfName)
        }

        if rs.Primary.ID == "" {
            return fmt.Errorf("ExceptionGroup ID is not set")
        }

        client := testAccProvider.Meta().(*checkpoint.ApiClient)

        response, err := client.ApiCall("show-exception-group", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
        if !response.Success {
            return err
        }

        *res = response.GetData()

        return nil
    }
}

func testAccCheckCheckpointManagementExceptionGroupAttributes(exceptionGroupMap *map[string]interface{}, name string, applyOn string) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        exceptionGroupName := (*exceptionGroupMap)["name"].(string)
        if !strings.EqualFold(exceptionGroupName, name) {
            return fmt.Errorf("name is %s, expected %s", name, exceptionGroupName)
        }
        exceptionGroupApplyOn := (*exceptionGroupMap)["apply-on"].(string)
        if !strings.EqualFold(exceptionGroupApplyOn, applyOn) {
            return fmt.Errorf("applyOn is %s, expected %s", applyOn, exceptionGroupApplyOn)
        }
        return nil
    }
}

func testAccManagementExceptionGroupConfig(name string, applyOn string) string {
    return fmt.Sprintf(`
resource "checkpoint_management_exception_group" "test" {
        name = "%s"
        apply_on = "%s"
}
`, name, applyOn)
}

