package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementSyslogServer_basic(t *testing.T) {

	objName := "tfTestManagementDataSyslogServer_" + acctest.RandString(6)
	hostName := "tfTestManagementDataHost_" + acctest.RandString(6)
	resourceName := "checkpoint_management_syslog_server.test"
	dataSourceName := "data.checkpoint_management_syslog_server.data_syslog_server"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementSyslogServerConfig(hostName, objName, 25),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "host", resourceName, "host"),
					resource.TestCheckResourceAttrPair(dataSourceName, "port", resourceName, "port"),
				),
			},
		},
	})

}

func testAccDataSourceManagementSyslogServerConfig(host string, name string, port int) string {
	return fmt.Sprintf(`

resource "checkpoint_management_host" "test_host" {
    name = "%s"
    ipv4_address = "4.4.4.4"
}

resource "checkpoint_management_syslog_server" "test" {
    name = "%s"
    host = "${checkpoint_management_host.test_host.name}"
    port = %d
}

data "checkpoint_management_syslog_server" "data_syslog_server" {
    name = "${checkpoint_management_syslog_server.test.name}"
}
`, host, name, port)
}
