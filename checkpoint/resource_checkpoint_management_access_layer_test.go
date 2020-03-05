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

func TestAccCheckpointManagementAccessLayer_basic(t *testing.T) {

    var accessLayerMap map[string]interface{}
    resourceName := "checkpoint_management_access_layer.test"
    objName := "tfTestManagementAccessLayer_" + acctest.RandString(6)

    context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
        CheckDestroy: testAccCheckpointManagementAccessLayerDestroy,
        Steps: []resource.TestStep{
            {
                Config: testAccManagementAccessLayerConfig(objName, ),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckCheckpointManagementAccessLayerExists(resourceName, &accessLayerMap),
                    testAccCheckCheckpointManagementAccessLayerAttributes(&accessLayerMap, objName, ),
                ),
            },
        },
    })
}

func testAccCheckpointManagementAccessLayerDestroy(s *terraform.State) error {

    client := testAccProvider.Meta().(*checkpoint.ApiClient)
    for _, rs := range s.RootModule().Resources {
        if rs.Type != "checkpoint_management_access_layer" {
            continue
        }
        if rs.Primary.ID != "" {
            res, _ := client.ApiCall("show-access-layer", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
            if res.Success {
                return fmt.Errorf("AccessLayer object (%s) still exists", rs.Primary.ID)
            }
        }
        return nil
    }
    return nil
}

func testAccCheckCheckpointManagementAccessLayerExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        rs, ok := s.RootModule().Resources[resourceTfName]
        if !ok {
            return fmt.Errorf("Resource not found: %s", resourceTfName)
        }

        if rs.Primary.ID == "" {
            return fmt.Errorf("AccessLayer ID is not set")
        }

        client := testAccProvider.Meta().(*checkpoint.ApiClient)

        response, err := client.ApiCall("show-access-layer", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
        if !response.Success {
            return err
        }

        *res = response.GetData()

        return nil
    }
}

func testAccCheckCheckpointManagementAccessLayerAttributes(accessLayerMap *map[string]interface{}, name string) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        accessLayerName := (*accessLayerMap)["name"].(string)
        if !strings.EqualFold(accessLayerName, name) {
            return fmt.Errorf("name is %s, expected %s", name, accessLayerName)
        }
        return nil
    }
}

func testAccManagementAccessLayerConfig(name string) string {
    return fmt.Sprintf(`
resource "checkpoint_management_access_layer" "test" {
        name = "%s"
        detect_using_x_forward_for = false
        applications_and_url_filtering = true
}
`, name)
}