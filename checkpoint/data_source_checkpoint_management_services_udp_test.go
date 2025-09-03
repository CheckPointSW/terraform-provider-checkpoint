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

func TestAccDataSourceCheckpointManagementShowServicesUdp_basic(t *testing.T) {
	var showServicesUdpQuery map[string]interface{}
	dataSourceShowServicesUdp := "data.checkpoint_management_data_services_udp.my_query"
	serviceUdpName := "serviceUdp_test"
	serviceUdpPort := "5683"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementShowServicesUdpConfig(1, serviceUdpName, serviceUdpPort),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointShowServicesUdp(dataSourceShowServicesUdp, &showServicesUdpQuery),
					testAccCheckCheckpointShowServicesUdpAttributes(&showServicesUdpQuery, serviceUdpName, serviceUdpPort),
				),
			},
		},
	})
}

func testAccCheckCheckpointShowServicesUdp(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("show-services-udp data source not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("show-services-udp data source ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)
		response, _ := client.ApiCall("show-services-udp", map[string]interface{}{"filter": "serviceUdp_test", "limit": 1}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return fmt.Errorf(response.ErrorMsg)
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointShowServicesUdpAttributes(servicesUdpMap *map[string]interface{}, serviceUdpName string, serviceUdpPort string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		servicesUdpMap := *servicesUdpMap
		if servicesUdpMap == nil {
			return fmt.Errorf("servicesUdpMap is nil")
		}

		// Create slice of obj names
		objectsJson := servicesUdpMap["objects"].([]interface{})
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
			return fmt.Errorf("show-services-udp returned wrong number of services udp. exptected for 1, found %d", len(objectsIds))
		}

		if serviceUdpName != objectsIds[0] {
			return fmt.Errorf("show-services-udp returned wrong service udp. exptected for %s, found %s", serviceUdpName, objectsIds[0])
		}

		if serviceUdpPort != objectsPorts[0] {
			return fmt.Errorf("show-services-udp returned wrong port. exptected for %s, found %s", serviceUdpPort, objectsPorts[0])
		}

		return nil
	}
}

func testAccDataSourceManagementShowServicesUdpConfig(limit int, objName string, objPort string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_service_udp" "example" {
  name = "%s"
  port = "%s"
}

data "checkpoint_management_data_services_udp" "my_query" {
	filter = "${checkpoint_management_service_udp.example.name}"
	limit = %d
}

data "checkpoint_management_data_service_udp" "data_service_udp" {
    name = "${data.checkpoint_management_data_services_udp.my_query.objects[0].name}"
}
`, objName, objPort, limit)
}
