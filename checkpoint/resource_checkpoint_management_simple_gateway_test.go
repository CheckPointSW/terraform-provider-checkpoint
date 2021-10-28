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

func TestAccCheckpointManagementSimpleGateway_basic(t *testing.T) {
	var simpleGatewayMap map[string]interface{}
	resourceName := "checkpoint_management_simple_gateway.test"
	objName := "tfTestManagementGateway_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementSimpleGatewayDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementSimpleGatewayConfig(objName, "1.2.3.4"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementSimpleGatewayExists(resourceName, &simpleGatewayMap),
					testAccCheckCheckpointManagementSimpleGatewayAttributes(&simpleGatewayMap, objName, "1.2.3.4"),
				),
			},
		},
	})
}

func testAccCheckpointManagementSimpleGatewayDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_simple_gateway" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-simple-gateway", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success {
				return fmt.Errorf("SimpleGateway object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementSimpleGatewayExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("SimpleGateway ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-simple-gateway", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementSimpleGatewayAttributes(simpleGatewayJson *map[string]interface{}, name string, ipv4 string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		simpleGatewayName := (*simpleGatewayJson)["name"].(string)
		if !strings.EqualFold(simpleGatewayName, name) {
			return fmt.Errorf("name is %s, expected %s", simpleGatewayName, name)
		}

		simpleGatewayIpv4 := (*simpleGatewayJson)["ipv4-address"].(string)
		if !strings.EqualFold(simpleGatewayIpv4, ipv4) {
			return fmt.Errorf("ipv4 is %s, expected %s", simpleGatewayIpv4, ipv4)
		}

		return nil
	}
}

func testAccManagementSimpleGatewayConfig(name string, ipv4 string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_checkpoint_host" "checkpoint_host" {
	name = "mycheckpointhost"
	ipv4_address = "5.6.9.8"
	management_blades = {
		network_policy_management = true
		logging_and_status = true
	}
}

resource "checkpoint_management_simple_gateway" "test" {
	name = "%s"
	ipv4_address = "%s"
	version = "R81"
	send_logs_to_server = ["${checkpoint_management_checkpoint_host.checkpoint_host.name}"]
}
`, name, ipv4)
}
