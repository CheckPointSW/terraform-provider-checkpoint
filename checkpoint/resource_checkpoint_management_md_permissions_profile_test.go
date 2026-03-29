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

func TestAccCheckpointManagementMdPermissionsProfile_basic(t *testing.T) {

	var mdPermissionsProfileMap map[string]interface{}
	resourceName := "checkpoint_management_md_permissions_profile.test"
	objName := "tfTestManagementMdPermissionsProfile_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementMdPermissionsProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementMdPermissionsProfileConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementMdPermissionsProfileExists(resourceName, &mdPermissionsProfileMap),
					testAccCheckCheckpointManagementMdPermissionsProfileAttributes(&mdPermissionsProfileMap, objName),
				),
			},
		},
	})
}

func testAccCheckpointManagementMdPermissionsProfileDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_md_permissions_profile" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-md-permissions-profile", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("MdPermissionsProfile object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementMdPermissionsProfileExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("MdPermissionsProfile ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-md-permissions-profile", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementMdPermissionsProfileAttributes(mdPermissionsProfileMap *map[string]interface{}, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		mdPermissionsProfileName := (*mdPermissionsProfileMap)["name"].(string)
		if !strings.EqualFold(mdPermissionsProfileName, name) {
			return fmt.Errorf("name is %s, expected %s", name, mdPermissionsProfileName)
		}
		return nil
	}
}

func testAccManagementMdPermissionsProfileConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_md_permissions_profile" "test" {
        name = "%s"
}
`, name)
}
