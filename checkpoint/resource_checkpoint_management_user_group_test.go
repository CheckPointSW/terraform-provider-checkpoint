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

func TestAccCheckpointManagementUserGroup_basic(t *testing.T) {

	var userGroupMap map[string]interface{}
	resourceName := "checkpoint_management_user_group.test"
	objName := "tfTestManagementUserGroup_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementUserGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementUserGroupConfig(objName, "myusergroup@email.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementUserGroupExists(resourceName, &userGroupMap),
					testAccCheckCheckpointManagementUserGroupAttributes(&userGroupMap, objName, "myusergroup@email.com"),
				),
			},
		},
	})
}

func testAccCheckpointManagementUserGroupDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_user_group" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-user-group", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("UserGroup object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementUserGroupExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("UserGroup ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-user-group", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementUserGroupAttributes(userGroupMap *map[string]interface{}, name string, email string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		userGroupName := (*userGroupMap)["name"].(string)
		if !strings.EqualFold(userGroupName, name) {
			return fmt.Errorf("name is %s, expected %s", name, userGroupName)
		}
		userGroupEmail := (*userGroupMap)["email"].(string)
		if !strings.EqualFold(userGroupEmail, email) {
			return fmt.Errorf("email is %s, expected %s", email, userGroupEmail)
		}
		return nil
	}
}

func testAccManagementUserGroupConfig(name string, email string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_user_group" "test" {
        name = "%s"
        email = "%s"
}
`, name, email)
}
