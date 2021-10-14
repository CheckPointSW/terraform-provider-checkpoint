package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	_ "github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	_ "strings"
	"testing"
)

func TestAccDataSourceCheckpointManagementShowObjects_basic(t *testing.T) {
	var showObjectsQuery map[string]interface{}
	dataSourceShowObjects := "data.checkpoint_management_show_objects.my_query"
	objName := "daytime-tcp"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementShowObjectsConfig("service-tcp", 1, objName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointShowObjects(dataSourceShowObjects, &showObjectsQuery),
					testAccCheckCheckpointShowObjectsAttributes(&showObjectsQuery, objName),
				),
			},
		},
	})
}

func testAccCheckCheckpointShowObjects(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("show-objects data source not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("show-objects data source ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)
		response, _ := client.ApiCall("show-objects", map[string]interface{}{"type": "service-tcp", "filter": "daytime-tcp", "limit": 1}, client.GetSessionID(), true, false)
		if !response.Success {
			return fmt.Errorf(response.ErrorMsg)
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointShowObjectsAttributes(showObjectsMap *map[string]interface{}, objName string) resource.TestCheckFunc {
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
				objectsIds = append(objectsIds, obj["name"].(string))
			}
		}

		if len(objectsIds) != 1 {
			return fmt.Errorf("show-objects returned wrong number of objects. exptected for 1, found %d", len(objectsIds))
		}

		if objName != objectsIds[0] {
			return fmt.Errorf("show-objects returned wrong object. exptected for %s, found %s", objName, objectsIds[0])
		}

		return nil
	}
}

func testAccDataSourceManagementShowObjectsConfig(objType string, limit int, objName string) string {
	return fmt.Sprintf(`
data "checkpoint_management_show_objects" "my_query" {
    type = "%s"
	filter = "%s"
	limit = %d
}

data "checkpoint_management_data_service_tcp" "data_service_tcp" {
    name = "${data.checkpoint_management_show_objects.my_query.objects[0].name}"
}
`, objType, objName, limit)
}
