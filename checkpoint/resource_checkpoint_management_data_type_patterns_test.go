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

func TestAccCheckpointManagementDataTypePatterns_basic(t *testing.T) {

	var dataTypePatternsMap map[string]interface{}
	resourceName := "checkpoint_management_data_type_patterns.test"
	objName := "tfTestManagementDataTypePatterns_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementDataTypePatternsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementDataTypePatternsConfig(objName, "data type pattern object", "a*b", "^d"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementDataTypePatternsExists(resourceName, &dataTypePatternsMap),
					testAccCheckCheckpointManagementDataTypePatternsAttributes(&dataTypePatternsMap, objName, "data type pattern object", "a*b", "^d"),
				),
			},
		},
	})
}

func testAccCheckpointManagementDataTypePatternsDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_data_type_patterns" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-data-type-patterns", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success {
				return fmt.Errorf("DataTypePatterns object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementDataTypePatternsExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("DataTypePatterns ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-data-type-patterns", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementDataTypePatternsAttributes(dataTypePatternsMap *map[string]interface{}, name string, description string, patterns1 string, patterns2 string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		dataTypePatternsName := (*dataTypePatternsMap)["name"].(string)
		if !strings.EqualFold(dataTypePatternsName, name) {
			return fmt.Errorf("name is %s, expected %s", name, dataTypePatternsName)
		}
		dataTypePatternsDescription := (*dataTypePatternsMap)["description"].(string)
		if !strings.EqualFold(dataTypePatternsDescription, description) {
			return fmt.Errorf("description is %s, expected %s", description, dataTypePatternsDescription)
		}

		return nil
	}
}

func testAccManagementDataTypePatternsConfig(name string, description string, patterns1 string, patterns2 string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_data_type_patterns" "test" {
        name = "%s"
        description = "%s"
       patterns = [ "%s" , "%s" ]
}
`, name, description, patterns1, patterns2)
}
