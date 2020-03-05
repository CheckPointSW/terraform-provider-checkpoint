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

func TestAccCheckpointManagementOpsecApplication_basic(t *testing.T) {

    var opsecApplicationMap map[string]interface{}
    resourceName := "checkpoint_management_opsec_application.test"
    objName := "tfTestManagementOpsecApplication_" + acctest.RandString(6)

    context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
        CheckDestroy: testAccCheckpointManagementOpsecApplicationDestroy,
        Steps: []resource.TestStep{
            {
                Config: testAccManagementOpsecApplicationConfig(objName, "somehost", "somepassword"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckCheckpointManagementOpsecApplicationExists(resourceName, &opsecApplicationMap),
                    testAccCheckCheckpointManagementOpsecApplicationAttributes(&opsecApplicationMap, objName, "somehost", "somepassword"),
                ),
            },
        },
    })
}

func testAccCheckpointManagementOpsecApplicationDestroy(s *terraform.State) error {

    client := testAccProvider.Meta().(*checkpoint.ApiClient)
    for _, rs := range s.RootModule().Resources {
        if rs.Type != "checkpoint_management_opsec_application" {
            continue
        }
        if rs.Primary.ID != "" {
            res, _ := client.ApiCall("show-opsec-application", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
            if res.Success {
                return fmt.Errorf("OpsecApplication object (%s) still exists", rs.Primary.ID)
            }
        }
        return nil
    }
    return nil
}

func testAccCheckCheckpointManagementOpsecApplicationExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        rs, ok := s.RootModule().Resources[resourceTfName]
        if !ok {
            return fmt.Errorf("Resource not found: %s", resourceTfName)
        }

        if rs.Primary.ID == "" {
            return fmt.Errorf("OpsecApplication ID is not set")
        }

        client := testAccProvider.Meta().(*checkpoint.ApiClient)

        response, err := client.ApiCall("show-opsec-application", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
        if !response.Success {
            return err
        }

        *res = response.GetData()

        return nil
    }
}

func testAccCheckCheckpointManagementOpsecApplicationAttributes(opsecApplicationMap *map[string]interface{}, name string, host string, oneTimePassword string) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        opsecApplicationName := (*opsecApplicationMap)["name"].(string)
        if !strings.EqualFold(opsecApplicationName, name) {
            return fmt.Errorf("name is %s, expected %s", name, opsecApplicationName)
        }
        opsecApplicationHost := (*opsecApplicationMap)["host"].(string)
        if !strings.EqualFold(opsecApplicationHost, host) {
            return fmt.Errorf("host is %s, expected %s", host, opsecApplicationHost)
        }
        return nil
    }
}

func testAccManagementOpsecApplicationConfig(name string, host string, oneTimePassword string) string {
    return fmt.Sprintf(`
resource "checkpoint_management_opsec_application" "test" {
        name = "%s"
        host = "%s"
        one_time_password = "%s"
      cpmi = {
        enabled = true
        administrator_profile = "read only all"
        use_administrator_credentials = false
      }
      lea = {
        enabled = true
        access_permissions = "show all"
      }
}
`, name, host, oneTimePassword)
}