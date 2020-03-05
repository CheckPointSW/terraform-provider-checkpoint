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

func TestAccCheckpointManagementAccessSection_basic(t *testing.T) {

    var accessSectionMap map[string]interface{}
    resourceName := "checkpoint_management_access_section.test"
    objName := "tfTestManagementAccessSection_" + acctest.RandString(6)

    context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
        CheckDestroy: testAccCheckpointManagementAccessSectionDestroy,
        Steps: []resource.TestStep{
            {
                Config: testAccManagementAccessSectionConfig(objName, ),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckCheckpointManagementAccessSectionExists(resourceName, &accessSectionMap),
                    testAccCheckCheckpointManagementAccessSectionAttributes(&accessSectionMap, objName, ),
                ),
            },
        },
    })
}

func testAccCheckpointManagementAccessSectionDestroy(s *terraform.State) error {

    client := testAccProvider.Meta().(*checkpoint.ApiClient)
    for _, rs := range s.RootModule().Resources {
        if rs.Type != "checkpoint_management_access_section" {
            continue
        }
        if rs.Primary.ID != "" {
            res, _ := client.ApiCall("show-access-section", map[string]interface{}{"uid": rs.Primary.ID, "layer": "network"}, client.GetSessionID(), true, false)
            if res.Success {
                return fmt.Errorf("AccessSection object (%s) still exists", rs.Primary.ID)
            }
        }
        return nil
    }
    return nil
}

func testAccCheckCheckpointManagementAccessSectionExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        rs, ok := s.RootModule().Resources[resourceTfName]
        if !ok {
            return fmt.Errorf("Resource not found: %s", resourceTfName)
        }

        if rs.Primary.ID == "" {
            return fmt.Errorf("AccessSection ID is not set")
        }

        client := testAccProvider.Meta().(*checkpoint.ApiClient)

        response, err := client.ApiCall("show-access-section", map[string]interface{}{"uid": rs.Primary.ID, "layer": "network"}, client.GetSessionID(), true, false)
        if !response.Success {
            return err
        }

        *res = response.GetData()

        return nil
    }
}

func testAccCheckCheckpointManagementAccessSectionAttributes(accessSectionMap *map[string]interface{}, name string) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        accessSectionName := (*accessSectionMap)["name"].(string)
        if !strings.EqualFold(accessSectionName, name) {
            return fmt.Errorf("name is %s, expected %s", name, accessSectionName)
        }
        return nil
    }
}

func testAccManagementAccessSectionConfig(name string) string {
    return fmt.Sprintf(`
resource "checkpoint_management_access_section" "test" {
        name = "%s"
        position = {top = "top"}
        layer = "network"
}
`, name)
}

