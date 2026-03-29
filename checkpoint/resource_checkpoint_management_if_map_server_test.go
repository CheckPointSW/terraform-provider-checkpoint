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

func TestAccCheckpointManagementIfMapServer_basic(t *testing.T) {

	var ifMapServerMap map[string]interface{}
	resourceName := "checkpoint_management_if_map_server.test"
	objName := "tfTestManagementIfMapServer_" + acctest.RandString(6)
	hostName := "tfTestManagementHost_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementIfMapServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementIfMapServerConfig(objName, hostName, "2.0", 1, "path"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementIfMapServerExists(resourceName, &ifMapServerMap),
					testAccCheckCheckpointManagementIfMapServerAttributes(&ifMapServerMap, objName, hostName, "2.0", 1, "path"),
				),
			},
		},
	})
}

func testAccCheckpointManagementIfMapServerDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_if_map_server" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-if-map-server", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success {
				return fmt.Errorf("IfMapServer object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementIfMapServerExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("IfMapServer ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-if-map-server", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementIfMapServerAttributes(ifMapServerMap *map[string]interface{}, name string, host string, version string, port int, path string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		ifMapServerName := (*ifMapServerMap)["name"].(string)
		if !strings.EqualFold(ifMapServerName, name) {
			return fmt.Errorf("name is %s, expected %s", name, ifMapServerName)
		}
		ifMapServerHost := (*ifMapServerMap)["host"].(map[string]interface{})["name"].(string)
		if !strings.EqualFold(ifMapServerHost, host) {
			return fmt.Errorf("host is %s, expected %s", host, ifMapServerHost)
		}
		ifMapServerVersion := (*ifMapServerMap)["version"].(string)
		if !strings.EqualFold(ifMapServerVersion, version) {
			return fmt.Errorf("version is %s, expected %s", version, ifMapServerVersion)
		}
		ifMapServerPort := int((*ifMapServerMap)["port"].(float64))
		if ifMapServerPort != port {
			return fmt.Errorf("port is %d, expected %d", port, ifMapServerPort)
		}
		ifMapServerPath := (*ifMapServerMap)["path"].(string)
		if !strings.EqualFold(ifMapServerPath, path) {
			return fmt.Errorf("path is %s, expected %s", path, ifMapServerPath)
		}
		return nil
	}
}

func testAccManagementIfMapServerConfig(name string, host string, version string, port int, path string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_host" "test_host" {
    name = "%s"
    ipv4_address = "9.9.9.9"
}

resource "checkpoint_management_if_map_server" "test" {
    name = "%s"
    host = "${checkpoint_management_host.test_host.name}"
    monitored_ips {
        first_ip = "0.0.0.0"
        last_ip = "0.0.0.0"
    }
    version = "%s"
    port = %d
    path = "%s"
}
`, host, name, version, port, path)
}
