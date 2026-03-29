package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementSecuremoteDnsServer_basic(t *testing.T) {

	hostName := "tfTestManagementDataHost_" + acctest.RandString(6)
	objName := "tfTestManagementDataSecuremoteDnsServer_" + acctest.RandString(6)
	resourceName := "checkpoint_management_securemote_dns_server.test"
	dataSourceName := "data.checkpoint_management_securemote_dns_server.data_securemote_dns_server"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementSecuremoteDnsServerConfig(hostName, objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "host", resourceName, "host"),
				),
			},
		},
	})

}

func testAccDataSourceManagementSecuremoteDnsServerConfig(host string, name string) string {
	return fmt.Sprintf(`

resource "checkpoint_management_host" "test_host" {
    name = "%s"
    ipv4_address = "4.4.4.4"
}

resource "checkpoint_management_securemote_dns_server" "test" {
	name = "%s"
	host = "${checkpoint_management_host.test_host.name}"
}

data "checkpoint_management_securemote_dns_server" "data_securemote_dns_server" {
  name = "${checkpoint_management_securemote_dns_server.test.name}"
}
`, host, name)
}
