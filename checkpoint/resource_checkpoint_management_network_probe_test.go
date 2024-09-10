package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	"strings"
	"testing"
)

func TestAccCheckpointManagementNetworkProbe_basic(t *testing.T) {

	var networkProbeMap map[string]interface{}
	resourceName := "checkpoint_management_network_probe.test"
	objName := "tfTestManagementNetworkProbe_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementNetworkProbeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementNetworkProbeConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementNetworkProbeExists(resourceName, &networkProbeMap),
					testAccCheckCheckpointManagementNetworkProbeAttributes(&networkProbeMap, objName),
				),
			},
		},
	})
}

func testAccCheckpointManagementNetworkProbeDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_network_probe" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-network-probe", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success {
				return fmt.Errorf("NetworkProbe object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementNetworkProbeExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("NetworkProbe ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-network-probe", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementNetworkProbeAttributes(networkProbeMap *map[string]interface{}, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		networkProbeName := (*networkProbeMap)["name"].(string)
		if !strings.EqualFold(networkProbeName, name) {
			return fmt.Errorf("name is %s, expected %s", name, networkProbeName)
		}
		return nil
	}
}

func testAccManagementNetworkProbeConfig(name string) string {
	return fmt.Sprintf(`
    resource "checkpoint_management_simple_gateway" "example" {
  name = "gw4"
  ipv4_address = "192.0.2.14"
vpn =true
}
     
resource "checkpoint_management_network_probe" "test" {
        name = "%s"
        install_on = ["${checkpoint_management_simple_gateway.example.name}"]
 icmp_options = {
    source = "10.10.10.10"
    destination = "25.20.20.20"
  }
  interval  = "20"
  protocol = "icmp"
}
`, name)
}
