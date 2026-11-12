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

func TestAccCheckpointManagementDataTypeTraditionalGroup_basic(t *testing.T) {

	var dataTypeTraditionalGroupMap map[string]interface{}
	resourceName := "checkpoint_management_data_type_traditional_group.test"
	objName := "tfTestManagementDataTypeTraditionalGroup_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementDataTypeTraditionalGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementDataTypeTraditionalGroupConfig(objName, "traditional group object", "CSV File", "SSH Private Key"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementDataTypeTraditionalGroupExists(resourceName, &dataTypeTraditionalGroupMap),
					testAccCheckCheckpointManagementDataTypeTraditionalGroupAttributes(&dataTypeTraditionalGroupMap, objName, "traditional group object", "CSV File", "SSH Private Key"),
				),
			},
		},
	})
}

func testAccCheckpointManagementDataTypeTraditionalGroupDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_data_type_traditional_group" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-data-type-traditional-group", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success {
				return fmt.Errorf("DataTypeTraditionalGroup object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementDataTypeTraditionalGroupExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("DataTypeTraditionalGroup ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-data-type-traditional-group", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementDataTypeTraditionalGroupAttributes(dataTypeTraditionalGroupMap *map[string]interface{}, name string, description string, dataTypes1 string, dataTypes2 string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		dataTypeTraditionalGroupName := (*dataTypeTraditionalGroupMap)["name"].(string)
		if !strings.EqualFold(dataTypeTraditionalGroupName, name) {
			return fmt.Errorf("name is %s, expected %s", name, dataTypeTraditionalGroupName)
		}
		dataTypeTraditionalGroupDescription := (*dataTypeTraditionalGroupMap)["description"].(string)
		if !strings.EqualFold(dataTypeTraditionalGroupDescription, description) {
			return fmt.Errorf("description is %s, expected %s", description, dataTypeTraditionalGroupDescription)
		}
		return nil
	}
}

func testAccManagementDataTypeTraditionalGroupConfig(name string, description string, dataTypes1 string, dataTypes2 string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_data_type_traditional_group" "test" {
        name = "%s"
        description = "%s"
        data_types = [ "%s" , "%s"]
}
`, name, description, dataTypes1, dataTypes2)
}
