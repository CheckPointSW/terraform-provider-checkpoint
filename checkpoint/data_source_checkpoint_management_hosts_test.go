package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	_ "github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	_ "strings"
	"testing"
)

func TestAccDataSourceCheckpointManagementShowHosts_basic(t *testing.T) {
	var showHostsQuery map[string]interface{}
	dataSourceShowHosts := "data.checkpoint_management_hosts.my_query"
	hostName := "tfTestManagementHost_" + acctest.RandString(6)
	ipv4Address := "2.3.4.115"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementShowHostsConfig(1, hostName, ipv4Address),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointShowHosts(dataSourceShowHosts, &showHostsQuery),
					testAccCheckCheckpointShowHostsAttributes(&showHostsQuery, hostName, ipv4Address),
				),
			},
		},
	})
}

func testAccCheckCheckpointShowHosts(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("show-hosts data source not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("show-hosts data source ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)
		response, _ := client.ApiCall("show-hosts", map[string]interface{}{"filter": "host1", "limit": 1}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return fmt.Errorf(response.ErrorMsg)
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointShowHostsAttributes(hostsMap *map[string]interface{}, hostName string, ipv4Address string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		hostsMap := *hostsMap
		if hostsMap == nil {
			return fmt.Errorf("hostsMap is nil")
		}

		// Create slice of obj names
		objectsJson := hostsMap["objects"].([]interface{})
		var objectsIds = make([]string, 0)
		var objectsIps = make([]string, 0)
		if len(objectsJson) > 0 {
			for _, obj := range objectsJson {
				obj := obj.(map[string]interface{})
				objectsIds = append(objectsIds, obj["name"].(string))
				objectsIps = append(objectsIps, obj["ipv4-address"].(string))
			}
		}

		if len(objectsIds) != 1 {
			return fmt.Errorf("show-hosts returned wrong number of hosts. exptected for 1, found %d", len(objectsIds))
		}

		if hostName != objectsIds[0] {
			return fmt.Errorf("show-hosts returned wrong host. exptected for %s, found %s", hostName, objectsIds[0])
		}

		if ipv4Address != objectsIps[0] {
			return fmt.Errorf("show-hosts returned wrong host IP. expected for %s, found %s", ipv4Address, objectsIps[0])
		}

		return nil
	}
}

func testAccDataSourceManagementShowHostsConfig(limit int, objName string, ipv4Address string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_host" "example" {
  name = "%s"
  ipv4_address = "%s"
}

data "checkpoint_management_hosts" "my_query" {
	filter = "${checkpoint_management_host.example.name}"
	limit = %d
	fetch_all = false
}
`, objName, ipv4Address, limit)
}
