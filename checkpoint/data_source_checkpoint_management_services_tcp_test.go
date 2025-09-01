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

func TestAccDataSourceCheckpointManagementShowServicesTcp_basic(t *testing.T) {
	var showServicesTcpQuery map[string]interface{}
	dataSourceShowServicesTcp := "data.checkpoint_management_data_services_tcp.my_query"
	serviceTcpName := "serviceTcp_test"
	serviceTcpPort := "5683"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementShowServicesTcpConfig(1, serviceTcpName, serviceTcpPort),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointShowServicesTcp(dataSourceShowServicesTcp, &showServicesTcpQuery),
					testAccCheckCheckpointShowServicesTcpAttributes(&showServicesTcpQuery, serviceTcpName, serviceTcpPort),
				),
			},
		},
	})
}

func testAccCheckCheckpointShowServicesTcp(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("show-services-tcp data source not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("show-services-tcp data source ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)
		response, _ := client.ApiCall("show-services-tcp", map[string]interface{}{"filter": "serviceTcp_test", "limit": 1}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return fmt.Errorf(response.ErrorMsg)
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointShowServicesTcpAttributes(servicesTcpMap *map[string]interface{}, serviceTcpName string, serviceTcpPort string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		servicesTcpMap := *servicesTcpMap
		if servicesTcpMap == nil {
			return fmt.Errorf("servicesTcpMap is nil")
		}

		// Create slice of obj names
		objectsJson := servicesTcpMap["objects"].([]interface{})
		var objectsIds = make([]string, 0)
		var objectsPorts = make([]string, 0)
		if len(objectsJson) > 0 {
			for _, obj := range objectsJson {
				obj := obj.(map[string]interface{})
				objectsIds = append(objectsIds, obj["name"].(string))
				objectsPorts = append(objectsPorts, obj["port"].(string))
			}
		}

		if len(objectsIds) != 1 {
			return fmt.Errorf("show-services-tcp returned wrong number of services tcp. exptected for 1, found %d", len(objectsIds))
		}

		if serviceTcpName != objectsIds[0] {
			return fmt.Errorf("show-services-tcp returned wrong service tcp. exptected for %s, found %s", serviceTcpName, objectsIds[0])
		}

		if serviceTcpPort != objectsPorts[0] {
			return fmt.Errorf("show-services-tcp returned wrong port. exptected for %s, found %s", serviceTcpPort, objectsPorts[0])
		}

		return nil
	}
}

func testAccDataSourceManagementShowServicesTcpConfig(limit int, objName string, objPort string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_service_tcp" "example" {
  name = "%s"
  port = "%s"
}

data "checkpoint_management_data_services_tcp" "my_query" {
	filter = "${checkpoint_management_service_tcp.example.name}"
	limit = %d
}

data "checkpoint_management_data_service_tcp" "data_service_tcp" {
    name = "${data.checkpoint_management_data_services_tcp.my_query.objects[0].name}"
}
`, objName, objPort, limit)
}
