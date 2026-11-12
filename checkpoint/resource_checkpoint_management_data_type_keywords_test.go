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

func TestAccCheckpointManagementDataTypeKeywords_basic(t *testing.T) {

	var dataTypeKeywordsMap map[string]interface{}
	resourceName := "checkpoint_management_data_type_keywords.test"
	objName := "tfTestManagementDataTypeKeywords_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementDataTypeKeywordsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementDataTypeKeywordsConfig(objName, "keywords object", "word1", "word2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementDataTypeKeywordsExists(resourceName, &dataTypeKeywordsMap),
					testAccCheckCheckpointManagementDataTypeKeywordsAttributes(&dataTypeKeywordsMap, objName, "keywords object", "word1", "word2"),
				),
			},
		},
	})
}

func testAccCheckpointManagementDataTypeKeywordsDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_data_type_keywords" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-data-type-keywords", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success {
				return fmt.Errorf("DataTypeKeywords object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementDataTypeKeywordsExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("DataTypeKeywords ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-data-type-keywords", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementDataTypeKeywordsAttributes(dataTypeKeywordsMap *map[string]interface{}, name string, description string, keywords1 string, keywords2 string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		dataTypeKeywordsName := (*dataTypeKeywordsMap)["name"].(string)
		if !strings.EqualFold(dataTypeKeywordsName, name) {
			return fmt.Errorf("name is %s, expected %s", name, dataTypeKeywordsName)
		}
		dataTypeKeywordsDescription := (*dataTypeKeywordsMap)["description"].(string)
		if !strings.EqualFold(dataTypeKeywordsDescription, description) {
			return fmt.Errorf("description is %s, expected %s", description, dataTypeKeywordsDescription)
		}

		return nil
	}
}

func testAccManagementDataTypeKeywordsConfig(name string, desc string, keyword1 string, keyword2 string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_data_type_keywords" "test" {
        name = "%s"
        description = "%s"
        keywords = ["%s","%s"]
}
`, name, desc, keyword1, keyword2)
}
