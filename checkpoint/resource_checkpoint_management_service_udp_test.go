package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	"testing"
)

func TestAccCheckpointManagementServiceUdp_basic(t *testing.T) {
	var serviceUdp map[string]interface{}
	resourceName := "checkpoint_management_service_udp.test"
	objName := "tfTestManagementServiceUdp_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointServiceUdpDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementServiceUdpConfig(objName, "15114"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointServiceUdpExists(resourceName, &serviceUdp),
					testAccCheckCheckpointServiceUdpAttributes(&serviceUdp, objName, "15114"),
				),
			},
		},
	})
}

func testAccCheckpointServiceUdpDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_service_udp" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-service-udp", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success { // Resource still exists. failed to destroy.
				return fmt.Errorf("service udp object (%s) still exists", rs.Primary.ID)
			}
		}
		break
	}
	return nil
}

func testAccCheckCheckpointServiceUdpExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("service udp ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)
		response, _ := client.ApiCall("show-service-udp", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return fmt.Errorf(response.ErrorMsg)
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointServiceUdpAttributes(serviceUdp *map[string]interface{}, name string, port string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		serviceUdp := *serviceUdp
		if serviceUdp == nil {
			return fmt.Errorf("service udp is nil")
		}

		serviceUdpName := serviceUdp["name"].(string)
		if serviceUdpName != name {
			return fmt.Errorf("name is %s, expected %s", serviceUdpName, name)
		}
		serviceUdpPort := serviceUdp["port"]
		if serviceUdpPort != port {
			return fmt.Errorf("port is %s, expected %s", serviceUdpPort, port)
		}

		return nil
	}
}

func testAccManagementServiceUdpConfig(name string, port string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_service_udp" "test" {
    name = "%s"
	port = "%s"
}
`, name, port)
}
