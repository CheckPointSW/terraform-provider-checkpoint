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

func TestAccCheckpointManagementDataTypeGroup_basic(t *testing.T) {

	var dataTypeGroupMap map[string]interface{}
	resourceName := "checkpoint_management_data_type_group.test"
	objName := "tfTestManagementDataTypeGroup_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementDataTypeGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementDataTypeGroupConfig(objName, "add data type group object", "Archive File"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementDataTypeGroupExists(resourceName, &dataTypeGroupMap),
					testAccCheckCheckpointManagementDataTypeGroupAttributes(&dataTypeGroupMap, objName, "add data type group object", "Archive File"),
				),
			},
		},
	})
}

func testAccCheckpointManagementDataTypeGroupDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_data_type_group" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-data-type-group", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success {
				return fmt.Errorf("DataTypeGroup object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementDataTypeGroupExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("DataTypeGroup ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-data-type-group", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementDataTypeGroupAttributes(dataTypeGroupMap *map[string]interface{}, name string, description string, fileType1 string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		dataTypeGroupName := (*dataTypeGroupMap)["name"].(string)
		if !strings.EqualFold(dataTypeGroupName, name) {
			return fmt.Errorf("name is %s, expected %s", name, dataTypeGroupName)
		}
		dataTypeGroupDescription := (*dataTypeGroupMap)["description"].(string)
		if !strings.EqualFold(dataTypeGroupDescription, description) {
			return fmt.Errorf("description is %s, expected %s", description, dataTypeGroupDescription)
		}
		return nil
	}
}

func testAccManagementDataTypeGroupConfig(name string, description string, fileType1 string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_data_type_group" "test" {
        name = "%s"
        description = "%s"
        file_type = ["%s"]
        file_content = ["CSV File"]
}
`, name, description, fileType1)
}
