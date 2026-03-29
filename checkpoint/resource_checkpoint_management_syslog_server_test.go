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

func TestAccCheckpointManagementSyslogServer_basic(t *testing.T) {

	var syslogServerMap map[string]interface{}
	resourceName := "checkpoint_management_syslog_server.test"
	objName := "tfTestManagementSyslogServer_" + acctest.RandString(6)
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
		CheckDestroy: testAccCheckpointManagementSyslogServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementSyslogServerConfig(objName, hostName, 18889),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementSyslogServerExists(resourceName, &syslogServerMap),
					testAccCheckCheckpointManagementSyslogServerAttributes(&syslogServerMap, objName, hostName, 18889),
				),
			},
		},
	})
}

func testAccCheckpointManagementSyslogServerDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_syslog_server" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-syslog-server", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success {
				return fmt.Errorf("SyslogServer object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementSyslogServerExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("SyslogServer ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-syslog-server", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementSyslogServerAttributes(syslogServerMap *map[string]interface{}, name string, host string, port int) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		syslogServerName := (*syslogServerMap)["name"].(string)
		if !strings.EqualFold(syslogServerName, name) {
			return fmt.Errorf("name is %s, expected %s", name, syslogServerName)
		}
		syslogServerHost := (*syslogServerMap)["host"].(map[string]interface{})["name"].(string)
		if !strings.EqualFold(syslogServerHost, host) {
			return fmt.Errorf("host is %s, expected %s", host, syslogServerHost)
		}
		syslogServerPort := int((*syslogServerMap)["port"].(float64))
		if syslogServerPort != port {
			return fmt.Errorf("port is %d, expected %d", port, syslogServerPort)
		}
		return nil
	}
}

func testAccManagementSyslogServerConfig(name string, host string, port int) string {
	return fmt.Sprintf(`
resource "checkpoint_management_host" "test_host" {
    name = "%s"
    ipv4_address = "1.1.14.143"
}

resource "checkpoint_management_syslog_server" "test" {
    name = "%s"
    host = "${checkpoint_management_host.test_host.name}"
    port = %d
}
`, host, name, port)
}
