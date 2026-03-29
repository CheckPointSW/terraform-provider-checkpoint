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

func TestAccCheckpointManagementDataCenterQuery_basic(t *testing.T) {
	var dataCenterQueryMap map[string]interface{}
	resourceName := "checkpoint_management_data_center_query.test"
	objName := "tfTestManagementDataCenterQuery_" + acctest.RandString(6)
	firstVal := "value1"
	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementDataCenterQueryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementDataCenterQueryConfig(objName, firstVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementDataCenterQueryExists(resourceName, &dataCenterQueryMap),
					testAccCheckCheckpointManagementDataCenterQueryAttributes(&dataCenterQueryMap, objName),
				),
			},
		},
	})
}

func testAccCheckpointManagementDataCenterQueryDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_data_center_query" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-data-center-query", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("DataCenterQuery object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementDataCenterQueryExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("DataCenterQuery ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-data-center-query", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementDataCenterQueryAttributes(dataCenterQueryMap *map[string]interface{}, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		dataCenterQueryName := (*dataCenterQueryMap)["name"].(string)
		if !strings.EqualFold(dataCenterQueryName, name) {
			return fmt.Errorf("name is %s, expected %s", name, dataCenterQueryName)
		}
		return nil
	}
}

func testAccManagementDataCenterQueryConfig(name string, firstVal string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_data_center_query" "test" {
	name = "%s"
  	data_centers = ["All"]
  	query_rules {
    key_type = "predefined"
    key      = "name-in-data-center"
    values   = ["%s"]
  }
}
`, name, firstVal)
}
