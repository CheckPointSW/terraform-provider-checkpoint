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

func TestAccCheckpointManagementSecurityZone_basic(t *testing.T) {

    var securityZoneMap map[string]interface{}
    resourceName := "checkpoint_management_security_zone.test"
    objName := "tfTestManagementSecurityZone_" + acctest.RandString(6)

    context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
        CheckDestroy: testAccCheckpointManagementSecurityZoneDestroy,
        Steps: []resource.TestStep{
            {
                Config: testAccManagementSecurityZoneConfig(objName, "my security zone 1", "yellow"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckCheckpointManagementSecurityZoneExists(resourceName, &securityZoneMap),
                    testAccCheckCheckpointManagementSecurityZoneAttributes(&securityZoneMap, objName, "my security zone 1", "yellow"),
                ),
            },
        },
    })
}

func testAccCheckpointManagementSecurityZoneDestroy(s *terraform.State) error {

    client := testAccProvider.Meta().(*checkpoint.ApiClient)
    for _, rs := range s.RootModule().Resources {
        if rs.Type != "checkpoint_management_security_zone" {
            continue
        }
        if rs.Primary.ID != "" {
            res, _ := client.ApiCall("show-security-zone", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
            if res.Success {
                return fmt.Errorf("SecurityZone object (%s) still exists", rs.Primary.ID)
            }
        }
        return nil
    }
    return nil
}

func testAccCheckCheckpointManagementSecurityZoneExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        rs, ok := s.RootModule().Resources[resourceTfName]
        if !ok {
            return fmt.Errorf("Resource not found: %s", resourceTfName)
        }

        if rs.Primary.ID == "" {
            return fmt.Errorf("SecurityZone ID is not set")
        }

        client := testAccProvider.Meta().(*checkpoint.ApiClient)

        response, err := client.ApiCall("show-security-zone", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
        if !response.Success {
            return err
        }

        *res = response.GetData()

        return nil
    }
}

func testAccCheckCheckpointManagementSecurityZoneAttributes(securityZoneMap *map[string]interface{}, name string, comments string, color string) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        securityZoneName := (*securityZoneMap)["name"].(string)
        if !strings.EqualFold(securityZoneName, name) {
            return fmt.Errorf("name is %s, expected %s", name, securityZoneName)
        }
        securityZoneComments := (*securityZoneMap)["comments"].(string)
        if !strings.EqualFold(securityZoneComments, comments) {
            return fmt.Errorf("comments is %s, expected %s", comments, securityZoneComments)
        }
        securityZoneColor := (*securityZoneMap)["color"].(string)
        if !strings.EqualFold(securityZoneColor, color) {
            return fmt.Errorf("color is %s, expected %s", color, securityZoneColor)
        }
        return nil
    }
}

func testAccManagementSecurityZoneConfig(name string, comments string, color string) string {
    return fmt.Sprintf(`
resource "checkpoint_management_security_zone" "test" {
        name = "%s"
        comments = "%s"
        color = "%s"
}
`, name, comments, color)
}

