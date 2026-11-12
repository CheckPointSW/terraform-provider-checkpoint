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

func TestAccCheckpointManagementInteroperableDevice_basic(t *testing.T) {

	var interoperableDeviceMap map[string]interface{}
	resourceName := "checkpoint_management_interoperable_device.test"
	objName := "tfTestManagementInteroperableDevice_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementInteroperableDeviceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementInteroperableDeviceConfig(objName, "192.168.1.6"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementInteroperableDeviceExists(resourceName, &interoperableDeviceMap),
					testAccCheckCheckpointManagementInteroperableDeviceAttributes(&interoperableDeviceMap, objName, "192.168.1.6"),
				),
			},
		},
	})
}

func testAccCheckpointManagementInteroperableDeviceDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_interoperable_device" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-interoperable-device", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("InteroperableDevice object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementInteroperableDeviceExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("InteroperableDevice ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-interoperable-device", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementInteroperableDeviceAttributes(interoperableDeviceMap *map[string]interface{}, name string, ipv4Address string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		interoperableDeviceName := (*interoperableDeviceMap)["name"].(string)
		if !strings.EqualFold(interoperableDeviceName, name) {
			return fmt.Errorf("name is %s, expected %s", name, interoperableDeviceName)
		}
		interoperableDeviceIpv4Address := (*interoperableDeviceMap)["ipv4-address"].(string)
		if !strings.EqualFold(interoperableDeviceIpv4Address, ipv4Address) {
			return fmt.Errorf("ipv4Address is %s, expected %s", ipv4Address, interoperableDeviceIpv4Address)
		}
		return nil
	}
}

func testAccManagementInteroperableDeviceConfig(name string, ipv4Address string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_interoperable_device" "test" {
        name = "%s"
        ipv4_address = "%s"
}
`, name, ipv4Address)
}
