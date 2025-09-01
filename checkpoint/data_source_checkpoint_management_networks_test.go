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

func TestAccDataSourceCheckpointManagementShowNetworks_basic(t *testing.T) {
	var showNetworksQuery map[string]interface{}
	dataSourceShowNetworks := "data.checkpoint_management_networks.my_query"
	networkName := "network_test21"
	networkSubnet := "8.9.21.0"
	networkMask := 24

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementShowNetworksConfig(1, networkName, networkSubnet, networkMask),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointShowNetworks(dataSourceShowNetworks, &showNetworksQuery),
					testAccCheckCheckpointShowNetworksAttributes(&showNetworksQuery, networkName, networkSubnet, networkMask),
				),
			},
		},
	})
}

func testAccCheckCheckpointShowNetworks(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("show-networks data source not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("show-networks data source ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)
		response, _ := client.ApiCall("show-networks", map[string]interface{}{"filter": "network_test1", "limit": 1}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return fmt.Errorf(response.ErrorMsg)
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointShowNetworksAttributes(networksMap *map[string]interface{}, networkName string, networkSubnet string, networkMask int) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		networksMap := *networksMap
		if networksMap == nil {
			return fmt.Errorf("networksMap is nil")
		}

		// Create slice of obj names
		objectsJson := networksMap["objects"].([]interface{})
		var objectsIds = make([]string, 0)
		var objectsSubnets = make([]string, 0)
		var objectsMasks = make([]int, 0)
		if len(objectsJson) > 0 {
			for _, obj := range objectsJson {
				obj := obj.(map[string]interface{})
				objectsIds = append(objectsIds, obj["name"].(string))
				objectsSubnets = append(objectsSubnets, obj["subnet4"].(string))
				objectsMasks = append(objectsMasks, int(obj["mask-length4"].(float64)))
			}
		}

		if len(objectsIds) != 1 {
			return fmt.Errorf("show-networks returned wrong number of networks. exptected for 1, found %d", len(objectsIds))
		}

		if networkName != objectsIds[0] {
			return fmt.Errorf("show-networks returned wrong network. exptected for %s, found %s", networkName, objectsIds[0])
		}

		if networkSubnet != objectsSubnets[0] {
			return fmt.Errorf("show-networks returned wrong network subnet. expected %s, found %s", networkSubnet, objectsSubnets[0])
		}

		if networkMask != objectsMasks[0] {
			return fmt.Errorf("show-networks returned wrong network mask. expected %d, found %d", networkMask, objectsMasks[0])
		}

		return nil
	}
}

func testAccDataSourceManagementShowNetworksConfig(limit int, objName string, objSubnet string, objMask int) string {
	return fmt.Sprintf(`
resource "checkpoint_management_network" "example" {
  name = "%s"
  subnet4 = "%s"
  mask_length4 = %d
}

data "checkpoint_management_networks" "my_query" {
	filter = "${checkpoint_management_network.example.name}"
	limit = %d
	fetch_all = true
}

data "checkpoint_management_data_network" "data_service_tcp" {
    name = "${data.checkpoint_management_networks.my_query.objects[0].name}"
}
`, objName, objSubnet, objMask, limit)
}
