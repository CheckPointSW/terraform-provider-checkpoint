package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementInterface_basic(t *testing.T) {

	objName := "tfTestManagementDataInteroperableDevice_" + acctest.RandString(6)
	resourceName := "checkpoint_management_interface.test"
	dataSourceName := "data.checkpoint_management_interface.data"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementInterfaceConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementInterfaceConfig(name string) string {
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
data "checkpoint_management_interface" "data" {
  name = "${checkpoint_management_interface.test.name}"
  gateway_uid = "${checkpoint_management_interface.test.gateway_uid}"
}
`, name)
}
