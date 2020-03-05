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

func TestAccCheckpointManagementDynamicObject_basic(t *testing.T) {

    var dynamicObjectMap map[string]interface{}
    resourceName := "checkpoint_management_dynamic_object.test"
    objName := "tfTestManagementDynamicObject_" + acctest.RandString(6)

    context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
        CheckDestroy: testAccCheckpointManagementDynamicObjectDestroy,
        Steps: []resource.TestStep{
            {
                Config: testAccManagementDynamicObjectConfig(objName, "my dynamic object 1", "yellow"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckCheckpointManagementDynamicObjectExists(resourceName, &dynamicObjectMap),
                    testAccCheckCheckpointManagementDynamicObjectAttributes(&dynamicObjectMap, objName, "my dynamic object 1", "yellow"),
                ),
            },
        },
    })
}

func testAccCheckpointManagementDynamicObjectDestroy(s *terraform.State) error {

    client := testAccProvider.Meta().(*checkpoint.ApiClient)
    for _, rs := range s.RootModule().Resources {
        if rs.Type != "checkpoint_management_dynamic_object" {
            continue
        }
        if rs.Primary.ID != "" {
            res, _ := client.ApiCall("show-dynamic-object", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
            if res.Success {
                return fmt.Errorf("DynamicObject object (%s) still exists", rs.Primary.ID)
            }
        }
        return nil
    }
    return nil
}

func testAccCheckCheckpointManagementDynamicObjectExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        rs, ok := s.RootModule().Resources[resourceTfName]
        if !ok {
            return fmt.Errorf("Resource not found: %s", resourceTfName)
        }

        if rs.Primary.ID == "" {
            return fmt.Errorf("DynamicObject ID is not set")
        }

        client := testAccProvider.Meta().(*checkpoint.ApiClient)

        response, err := client.ApiCall("show-dynamic-object", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
        if !response.Success {
            return err
        }

        *res = response.GetData()

        return nil
    }
}

func testAccCheckCheckpointManagementDynamicObjectAttributes(dynamicObjectMap *map[string]interface{}, name string, comments string, color string) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        dynamicObjectName := (*dynamicObjectMap)["name"].(string)
        if !strings.EqualFold(dynamicObjectName, name) {
            return fmt.Errorf("name is %s, expected %s", name, dynamicObjectName)
        }
        dynamicObjectComments := (*dynamicObjectMap)["comments"].(string)
        if !strings.EqualFold(dynamicObjectComments, comments) {
            return fmt.Errorf("comments is %s, expected %s", comments, dynamicObjectComments)
        }
        dynamicObjectColor := (*dynamicObjectMap)["color"].(string)
        if !strings.EqualFold(dynamicObjectColor, color) {
            return fmt.Errorf("color is %s, expected %s", color, dynamicObjectColor)
        }
        return nil
    }
}

func testAccManagementDynamicObjectConfig(name string, comments string, color string) string {
    return fmt.Sprintf(`
resource "checkpoint_management_dynamic_object" "test" {
        name = "%s"
        comments = "%s"
        color = "%s"
}
`, name, comments, color)
}

