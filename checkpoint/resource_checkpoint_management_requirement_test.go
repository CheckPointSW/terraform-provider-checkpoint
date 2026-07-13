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

func TestAccCheckpointManagementRequirement_basic(t *testing.T) {

	var requirementMap map[string]interface{}
	resourceName := "checkpoint_management_requirement.test"
	objName := "tfTestManagementRequirement_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementRequirementDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementRequirementConfig(objName, "myreg", "my new requirement"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementRequirementExists(resourceName, &requirementMap),
					testAccCheckCheckpointManagementRequirementAttributes(&requirementMap, objName, "myreg", "my new requirement"),
				),
			},
		},
	})
}

func testAccCheckpointManagementRequirementDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_requirement" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-requirement", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("Requirement object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementRequirementExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Requirement ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-requirement", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementRequirementAttributes(requirementMap *map[string]interface{}, name string, regulationName string, comments string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		requirementName := (*requirementMap)["name"].(string)
		if !strings.EqualFold(requirementName, name) {
			return fmt.Errorf("name is %s, expected %s", name, requirementName)
		}
		requirementRegulationName := (*requirementMap)["regulation"].(map[string]interface{})["name"].(string)
		if !strings.EqualFold(requirementRegulationName, regulationName) {
			return fmt.Errorf("regulationName is %s, expected %s", regulationName, requirementRegulationName)
		}
		requirementComments := (*requirementMap)["comments"].(string)
		if !strings.EqualFold(requirementComments, comments) {
			return fmt.Errorf("comments is %s, expected %s", comments, requirementComments)
		}
		return nil
	}
}

func testAccManagementRequirementConfig(name string, regulationName string, comments string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_requirement" "test" {
        name = "%s"
        regulation = "%s"
        comments = "%s"
}
`, name, regulationName, comments)
}
