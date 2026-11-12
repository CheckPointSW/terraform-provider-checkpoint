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

func TestAccCheckpointManagementPasscodeProfile_basic(t *testing.T) {

	var passcodeProfileMap map[string]interface{}
	resourceName := "checkpoint_management_passcode_profile.test"
	objName := "tfTestManagementPasscodeProfile_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementPasscodeProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementPasscodeProfileConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementPasscodeProfileExists(resourceName, &passcodeProfileMap),
					testAccCheckCheckpointManagementPasscodeProfileAttributes(&passcodeProfileMap, objName),
				),
			},
		},
	})
}

func testAccCheckpointManagementPasscodeProfileDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_passcode_profile" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-passcode-profile", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success {
				return fmt.Errorf("PasscodeProfile object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementPasscodeProfileExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("PasscodeProfile ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-passcode-profile", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementPasscodeProfileAttributes(passcodeProfileMap *map[string]interface{}, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		passcodeProfileName := (*passcodeProfileMap)["name"].(string)
		if !strings.EqualFold(passcodeProfileName, name) {
			return fmt.Errorf("name is %s, expected %s", name, passcodeProfileName)
		}
		return nil
	}
}

func testAccManagementPasscodeProfileConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_passcode_profile" "test" {
        name = "%s"
}
`, name)
}
