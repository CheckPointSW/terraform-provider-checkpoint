package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementProxmoxDataCenterServer_basic(t *testing.T) {

	objName := "tfTestManagementDataProxmoxDataCenterServer_" + acctest.RandString(6)
	resourceName := "checkpoint_management_proxmox_data_center_server.proxmox_data_center_server"
	dataSourceName := "data.checkpoint_management_proxmox_data_center_server.proxmox_data_center_server"
	token_id := "USER@PAM!TOKEN_ID"
	secret := "SECRET"
	hostname := "HOSTNAME"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementProxmoxDataCenterServerConfig(objName, hostname, token_id, secret),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementProxmoxDataCenterServerConfig(name string, hostname string, token_id string, secret string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_proxmox_data_center_server" "proxmox_data_center_server" {
    name = "%s"
	token_id = "%s"
	secret = "%s"
	hostname = "%s"
    unsafe_auto_accept = true
	ignore_warnings = true
}

data "checkpoint_management_proxmox_data_center_server" "proxmox_data_center_server" {
    name = "${checkpoint_management_proxmox_data_center_server.proxmox_data_center_server.name}"
}
`, name, token_id, secret, hostname)
}
