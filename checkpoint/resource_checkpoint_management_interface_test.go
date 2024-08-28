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

func TestAccCheckpointManagementInterface_basic(t *testing.T) {

	var interfaceMap map[string]interface{}
	resourceName := "checkpoint_management_interface.test"
	objName := "tfTestManagementInterface_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementInterfaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementInterfaceConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementInterfaceExists(resourceName, &interfaceMap),
					//	testAccCheckCheckpointManagementInterfaceAttributes(&interfaceMap, objName),
				),
			},
		},
	})
}

func testAccCheckpointManagementInterfaceDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_interface" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-interface", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success {
				return fmt.Errorf("Interface object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementInterfaceExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Interface ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-interface", map[string]interface{}{"uid": rs.Primary}, client.GetSessionID(), true, false)
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementInterfaceAttributes(interfaceMap *map[string]interface{}, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		interfaceName := (*interfaceMap)["name"].(string)
		if !strings.EqualFold(interfaceName, name) {
			return fmt.Errorf("name is %s, expected %s", name, interfaceName)
		}

		return nil
	}
}

func testAccManagementInterfaceConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_simple_gateway" "fw" {
  name = "gw3"
  ipv4_address = "192.0.2.1"
}


resource "checkpoint_management_interface" "test" {
  name = "%s"
  gateway_uid = "${checkpoint_management_simple_gateway.fw.id}"
  ipv4_address = "1.1.1.114"
  ipv4_mask_length = 24
  anti_spoofing_settings {
    action = "detect"
    exclude_packets = false
    spoof_tracking = "log"

  }

  security_zone_settings {
    auto_calculated = false
    specific_zone = "InternalZone"

  }
  topology_settings {
    interface_leads_to_dmz = false
    ip_address_behind_this_interface = "network defined by routing"

  }
  topology = "internal"
  network_interface_type = "ethernet"
  cluster_members {
      name = "eth4"
      member_name = "member1"
    ipv4_address = "192.168.1.2"
  }

}
`, name)
}
