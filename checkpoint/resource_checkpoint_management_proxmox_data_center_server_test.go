package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	"strings"
	"testing"
)

func TestAccCheckpointManagementProxmoxDataCenterServer_basic(t *testing.T) {

	var proxmoxDataCenterServerMap map[string]interface{}
	resourceName := "checkpoint_management_proxmox_data_center_server.test"
	objName := "tfTestManagementProxmoxDataCenterServer_" + acctest.RandString(6)
	token_id := "USER@PAM!TOKEN_ID"
	secret := "SECRET"
	hostname := "HOSTNAME"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementProxmoxDataCenterServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementProxmoxDataCenterServerConfig(objName, hostname, token_id, secret),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementProxmoxDataCenterServerExists(resourceName, &proxmoxDataCenterServerMap),
					testAccCheckCheckpointManagementProxmoxDataCenterServerAttributes(&proxmoxDataCenterServerMap, objName),
				),
			},
		},
	})
}

func testAccCheckpointManagementProxmoxDataCenterServerDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_proxmox_data_center_server" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-data-center-server", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("ProxmoxDataCenterServer object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementProxmoxDataCenterServerExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ProxmoxDataCenterServer ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-data-center-server", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementProxmoxDataCenterServerAttributes(proxmoxDataCenterServerMap *map[string]interface{}, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		proxmoxDataCenterServerName := (*proxmoxDataCenterServerMap)["name"].(string)
		if !strings.EqualFold(proxmoxDataCenterServerName, name) {
			return fmt.Errorf("name is %s, expected %s", name, proxmoxDataCenterServerName)
		}
		return nil
	}
}

func testAccManagementProxmoxDataCenterServerConfig(name string, hostname string, token_id string, secret string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_proxmox_data_center_server" "test" {
    name = "%s"
	token_id = "%s"
	secret = "%s"
	hostname = "%s"
    unsafe_auto_accept = true
	ignore_warnings = true
}
`, name, token_id, secret, hostname)
}
