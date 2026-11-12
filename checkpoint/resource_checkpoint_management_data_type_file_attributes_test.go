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

func TestAccCheckpointManagementDataTypeFileAttributes_basic(t *testing.T) {

	var dataTypeFileAttributesMap map[string]interface{}
	resourceName := "checkpoint_management_data_type_file_attributes.test"
	objName := "tfTestManagementDataTypeFileAttributes_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementDataTypeFileAttributesDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementDataTypeFileAttributesConfig(objName, true, "expression", true, 14),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementDataTypeFileAttributesExists(resourceName, &dataTypeFileAttributesMap),
					testAccCheckCheckpointManagementDataTypeFileAttributesAttributes(&dataTypeFileAttributesMap, objName, true, "expression", true, 14),
				),
			},
		},
	})
}

func testAccCheckpointManagementDataTypeFileAttributesDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_data_type_file_attributes" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-data-type-file-attributes", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success {
				return fmt.Errorf("DataTypeFileAttributes object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementDataTypeFileAttributesExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("DataTypeFileAttributes ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-data-type-file-attributes", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementDataTypeFileAttributesAttributes(dataTypeFileAttributesMap *map[string]interface{}, name string, matchByFileName bool, fileNameContains string, matchByFileSize bool, fileSize int) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		dataTypeFileAttributesName := (*dataTypeFileAttributesMap)["name"].(string)
		if !strings.EqualFold(dataTypeFileAttributesName, name) {
			return fmt.Errorf("name is %s, expected %s", name, dataTypeFileAttributesName)
		}

		dataTypeFileAttributesMatchByFileName := (*dataTypeFileAttributesMap)["match-by-file-name"].(bool)
		if dataTypeFileAttributesMatchByFileName != matchByFileName {
			return fmt.Errorf("matchByFileName is %t, expected %t", matchByFileName, dataTypeFileAttributesMatchByFileName)
		}
		dataTypeFileAttributesFileNameContains := (*dataTypeFileAttributesMap)["file-name-contains"].(string)
		if !strings.EqualFold(dataTypeFileAttributesFileNameContains, fileNameContains) {
			return fmt.Errorf("fileNameContains is %s, expected %s", fileNameContains, dataTypeFileAttributesFileNameContains)
		}
		dataTypeFileAttributesMatchByFileSize := (*dataTypeFileAttributesMap)["match-by-file-size"].(bool)
		if dataTypeFileAttributesMatchByFileSize != matchByFileSize {
			return fmt.Errorf("matchByFileSize is %t, expected %t", matchByFileSize, dataTypeFileAttributesMatchByFileSize)
		}
		dataTypeFileAttributesFileSize := int((*dataTypeFileAttributesMap)["file-size"].(float64))
		if dataTypeFileAttributesFileSize != fileSize {
			return fmt.Errorf("fileSize is %d, expected %d", fileSize, dataTypeFileAttributesFileSize)
		}
		return nil
	}
}

func testAccManagementDataTypeFileAttributesConfig(name string, matchByFileName bool, fileNameContains string, matchByFileSize bool, fileSize int) string {
	return fmt.Sprintf(`
resource "checkpoint_management_data_type_file_attributes" "test" {
        name = "%s"
       match_by_file_name = "true"
        file_name_contains = "%s"
        match_by_file_size = "true"
        file_size = "%d"
}
`, name, fileNameContains, fileSize)
}
