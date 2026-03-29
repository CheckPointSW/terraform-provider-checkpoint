package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementTcp_basic(t *testing.T) {

	hostName := "tfTestManagementHost_" + acctest.RandString(6)
	opsecName := "tfTestManagementOpsec_" + acctest.RandString(6)
	objName := "tfTestManagementDataTcp_" + acctest.RandString(6)
	resourceName := "checkpoint_management_resource_tcp.resource_tcp"
	dataSourceName := "data.checkpoint_management_resource_tcp.data_tcp"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementTcpConfig(hostName, opsecName, objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "resource_type", resourceName, "resource_type"),
					resource.TestCheckResourceAttrPair(dataSourceName, "exception_track", resourceName, "exception_track"),
				),
			},
		},
	})

}

func testAccDataSourceManagementTcpConfig(host string, opsec string, name string) string {
	return fmt.Sprintf(`

resource "checkpoint_management_host" "test_host" {
    name = "%s"
    ipv4_address = "1.1.1.1"
}

resource "checkpoint_management_opsec_application" "test_opsec_app" {
    name = "%s"
    host = "${checkpoint_management_host.test_host.name}"
    one_time_password = "SomePassword"
}

resource "checkpoint_management_resource_tcp" "resource_tcp" {
    name = "%s"
    ufp_settings {
        server = "${checkpoint_management_opsec_application.test_opsec_app.id}"
    }
}

data "checkpoint_management_resource_tcp" "data_tcp" {
  name = "${checkpoint_management_resource_tcp.resource_tcp.name}"
}
`, host, opsec, name)
}
