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

func TestAccCheckpointManagementDataTypeWeightedKeywords_basic(t *testing.T) {

	var dataTypeWeightedKeywordsMap map[string]interface{}
	resourceName := "checkpoint_management_data_type_weighted_keywords.test"
	objName := "tfTestManagementDataTypeWeightedKeywords_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementDataTypeWeightedKeywordsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementDataTypeWeightedKeywordsConfig(objName, "word1", 3, 4),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementDataTypeWeightedKeywordsExists(resourceName, &dataTypeWeightedKeywordsMap),
					testAccCheckCheckpointManagementDataTypeWeightedKeywordsAttributes(&dataTypeWeightedKeywordsMap, objName, "word1", 3, 4),
				),
			},
		},
	})
}

func testAccCheckpointManagementDataTypeWeightedKeywordsDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_data_type_weighted_keywords" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-data-type-weighted-keywords", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success {
				return fmt.Errorf("DataTypeWeightedKeywords object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementDataTypeWeightedKeywordsExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("DataTypeWeightedKeywords ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-data-type-weighted-keywords", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementDataTypeWeightedKeywordsAttributes(dataTypeWeightedKeywordsMap *map[string]interface{}, name string, Keyword1 string, Weight1 int, MaxWeight1 int) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		dataTypeWeightedKeywordsName := (*dataTypeWeightedKeywordsMap)["name"].(string)
		if !strings.EqualFold(dataTypeWeightedKeywordsName, name) {
			return fmt.Errorf("name is %s, expected %s", name, dataTypeWeightedKeywordsName)
		}

		return nil
	}
}

func testAccManagementDataTypeWeightedKeywordsConfig(name string, Keyword1 string, Weight1 int, MaxWeight1 int) string {
	return fmt.Sprintf(`
resource "checkpoint_management_data_type_weighted_keywords" "test" {
        name = "%s"
        weighted_keywords {
           keyword = "%s"
           weight = "%d"
           max_weight = "%d"
         
        }
}
`, name, Keyword1, Weight1, MaxWeight1)
}
