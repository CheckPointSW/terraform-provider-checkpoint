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

func TestAccCheckpointManagementDataTypeCompoundGroup_basic(t *testing.T) {

	var dataTypeCompoundGroupMap map[string]interface{}
	resourceName := "checkpoint_management_data_type_compound_group.test"
	objName := "tfTestManagementDataTypeCompoundGroup_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementDataTypeCompoundGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementDataTypeCompoundGroupConfig(objName, "compound group object", "Source Code", "Large File"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementDataTypeCompoundGroupExists(resourceName, &dataTypeCompoundGroupMap),
					testAccCheckCheckpointManagementDataTypeCompoundGroupAttributes(&dataTypeCompoundGroupMap, objName, "compound group object", "Source Code", "Large File"),
				),
			},
		},
	})
}

func testAccCheckpointManagementDataTypeCompoundGroupDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_data_type_compound_group" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-data-type-compound-group", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success {
				return fmt.Errorf("DataTypeCompoundGroup object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementDataTypeCompoundGroupExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("DataTypeCompoundGroup ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-data-type-compound-group", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementDataTypeCompoundGroupAttributes(dataTypeCompoundGroupMap *map[string]interface{}, name string, description string, matchedGroups1 string, unmatchedGroups1 string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		dataTypeCompoundGroupName := (*dataTypeCompoundGroupMap)["name"].(string)
		if !strings.EqualFold(dataTypeCompoundGroupName, name) {
			return fmt.Errorf("name is %s, expected %s", name, dataTypeCompoundGroupName)
		}
		dataTypeCompoundGroupDescription := (*dataTypeCompoundGroupMap)["description"].(string)
		if !strings.EqualFold(dataTypeCompoundGroupDescription, description) {
			return fmt.Errorf("description is %s, expected %s", description, dataTypeCompoundGroupDescription)
		}
		return nil
	}
}

func testAccManagementDataTypeCompoundGroupConfig(name string, description string, matchedGroups1 string, unmatchedGroups1 string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_data_type_compound_group" "test" {
        name = "%s"
        description = "%s"
        matched_groups =  ["%s"]
        unmatched_groups = ["%s"]
}
`, name, description, matchedGroups1, unmatchedGroups1)
}
