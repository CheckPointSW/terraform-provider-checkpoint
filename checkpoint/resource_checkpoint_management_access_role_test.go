package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	"strings"
	"testing"
)

func TestAccCheckpointManagementAccessRole_basic(t *testing.T) {

	var accessRoleMap map[string]interface{}
	resourceName := "checkpoint_management_access_role.test"
	objName := "tfTestManagementAccessRole_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementAccessRoleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementAccessRoleConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementAccessRoleExists(resourceName, &accessRoleMap),
					testAccCheckCheckpointManagementAccessRoleAttributes(&accessRoleMap, objName),
				),
			},
		},
	})
}

func testAccCheckpointManagementAccessRoleDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_access_role" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-access-role", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success {
				return fmt.Errorf("AccessRole object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementAccessRoleExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("AccessRole ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-access-role", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementAccessRoleAttributes(accessRoleMap *map[string]interface{}, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		accessRoleName := (*accessRoleMap)["name"].(string)
		if !strings.EqualFold(accessRoleName, name) {
			return fmt.Errorf("name is %s, expected %s", name, accessRoleName)
		}
		return nil
	}
}

func testAccManagementAccessRoleConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_access_role" "test" {
	name = "%s"
	machines {
    selection = ["any"]
    source = "any"
	}
	users {
	selection = ["any"]
	source = "any"
	}
}
`, name)
}
