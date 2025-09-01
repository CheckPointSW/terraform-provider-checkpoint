package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementIfMapServer_basic(t *testing.T) {

	objName := "tfTestManagementDataIfMapServer_" + acctest.RandString(6)
	resourceName := "checkpoint_management_if_map_server.test"
	dataSourceName := "data.checkpoint_management_if_map_server.data_if_map_server"
	hostName := "tfTestManagementHost_" + acctest.RandString(6)
	version := "2.0"
	port := 1
	path := "path"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementIfMapServerConfig(hostName, objName, version, port, path),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "monitored_ips.#", resourceName, "monitored_ips.#"),
					resource.TestCheckResourceAttrPair(dataSourceName, "monitored_ips.0.first_ip", resourceName, "monitored_ips.0.first_ip"),
					resource.TestCheckResourceAttrPair(dataSourceName, "monitored_ips.0.last_ip", resourceName, "monitored_ips.0.last_ip"),
					resource.TestCheckResourceAttrPair(dataSourceName, "version", resourceName, "version"),
					resource.TestCheckResourceAttrPair(dataSourceName, "port", resourceName, "port"),
					resource.TestCheckResourceAttrPair(dataSourceName, "path", resourceName, "path"),
				),
			},
		},
	})

}

func testAccDataSourceManagementIfMapServerConfig(host string, name string, version string, port int, path string) string {
	return fmt.Sprintf(`

resource "checkpoint_management_host" "test_host" {
    name = "%s"
    ipv4_address = "6.6.6.6"
}

resource "checkpoint_management_if_map_server" "test1" {
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

data "checkpoint_management_if_map_server" "data_if_map_server" {
	name = "${checkpoint_management_if_map_server.test.name}"
}
`, host, name, version, port, path)
}
