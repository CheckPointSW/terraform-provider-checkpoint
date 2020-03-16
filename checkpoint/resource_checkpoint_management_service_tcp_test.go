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

func TestAccCheckpointManagementServiceTcp_basic(t *testing.T) {
	var serviceTcp map[string]interface{}
	resourceName := "checkpoint_management_service_tcp.test"
	objName := "tfTestManagementServiceTcp_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointServiceTcpDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementServiceTcpConfig(objName, "15214"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointServiceTcpExists(resourceName, &serviceTcp),
					testAccCheckCheckpointServiceTcpAttributes(&serviceTcp, objName, "15214"),
				),
			},
		},
	})
}

func testAccCheckpointServiceTcpDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_service_tcp" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-service-tcp", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success { // Resource still exists. failed to destroy.
				return fmt.Errorf("service tcp object (%s) still exists", rs.Primary.ID)
			}
		}
		break
	}
	return nil
}

func testAccCheckCheckpointServiceTcpExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("service tcp ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)
		response, _ := client.ApiCall("show-service-tcp", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return fmt.Errorf(response.ErrorMsg)
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointServiceTcpAttributes(serviceTcp *map[string]interface{}, name string, port string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		serviceTcp := *serviceTcp
		if serviceTcp == nil {
			return fmt.Errorf("service tcp is nil")
		}

		serviceTcpName := serviceTcp["name"].(string)
		if serviceTcpName != name {
			return fmt.Errorf("name is %s, expected %s", serviceTcpName, name)
		}
		serviceTcpPort := serviceTcp["port"]
		if serviceTcpPort != port {
			return fmt.Errorf("port is %s, expected %s", serviceTcpPort, port)
		}

		return nil
	}
}

func testAccManagementServiceTcpConfig(name string, port string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_service_tcp" "test" {
    name = "%s"
	port = "%s"
}
`, name, port)
}
