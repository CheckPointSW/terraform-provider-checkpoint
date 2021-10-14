package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	_ "github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	"strings"
	_ "strings"
	"testing"
)

func TestAccDataSourceCheckpointManagementShowUpdatableObjectsRepositoryContent_basic(t *testing.T) {
	var showObjectsQuery map[string]interface{}
	dataSourceShowObjects := "data.checkpoint_management_show_updatable_objects_repository_content.test"
	objName := "API Gateway AF South"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementShowUpdatableObjectsRepositoryContentConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointShowUpdatableObjectsRepositoryContent(dataSourceShowObjects, &showObjectsQuery, objName),
					testAccCheckCheckpointShowUpdatableObjectsRepositoryContentAttributes(&showObjectsQuery, objName),
				),
			},
		},
	})
}

func testAccCheckCheckpointShowUpdatableObjectsRepositoryContent(resourceTfName string, res *map[string]interface{}, objName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("show-updatable-objects-repository-content data source not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("show-updatable-objects-repository-content data source ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)
		payload := make(map[string]interface{})
		payload["filter"] = map[string]interface{}{"text": objName}
		response, _ := client.ApiCall("show-updatable-objects-repository-content", payload, client.GetSessionID(), true, false)
		if !response.Success {
			return fmt.Errorf(response.ErrorMsg)
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointShowUpdatableObjectsRepositoryContentAttributes(showObjectsMap *map[string]interface{}, objName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		showObjectsMap := *showObjectsMap
		if showObjectsMap == nil {
			return fmt.Errorf("showObjectsMap is nil")
		}

		// Create slice of obj names
		objectsJson := showObjectsMap["objects"].([]interface{})
		var objectsIds = make([]string, 0)
		if len(objectsJson) > 0 {
			for _, obj := range objectsJson {
				obj := obj.(map[string]interface{})
				objectsIds = append(objectsIds, obj["name-in-updatable-objects-repository"].(string))
			}
		}

		if len(objectsIds) != 1 {
			return fmt.Errorf("show-updatable-objects-repository-content returned wrong number of objects. exptected for 1, found %d", len(objectsIds))
		}

		if !strings.Contains(objectsIds[0], objName) {
			return fmt.Errorf("show-updatable-objects-repository-content returned wrong object. exptected for %s, found %s", objName, objectsIds[0])
		}

		return nil
	}
}

func testAccDataSourceManagementShowUpdatableObjectsRepositoryContentConfig(objName string) string {
	return fmt.Sprintf(`
data "checkpoint_management_show_updatable_objects_repository_content" "test" {
    filter = {
		text = "%s"
	}
}
`, objName)
}
