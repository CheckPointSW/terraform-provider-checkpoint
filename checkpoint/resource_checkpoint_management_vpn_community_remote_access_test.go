package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"os"
	"strings"
	"testing"
)

func TestAccCheckpointManagementVpnCommunityRemoteAccess_basic(t *testing.T) {

	var vpnCommunityRemoteAccessMap map[string]interface{}
	resourceName := "checkpoint_management_vpn_community_remote_access.test"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementVpnCommunityRemoteAccessDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementVpnCommunityRemoteAccessConfig("RemoteAccess"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementVpnCommunityRemoteAccessExists(resourceName, &vpnCommunityRemoteAccessMap),
					testAccCheckCheckpointManagementVpnCommunityRemoteAccessAttributes(&vpnCommunityRemoteAccessMap, "RemoteAccess", "myusergroup123"),
				),
			},
		},
	})
}

func testAccCheckpointManagementVpnCommunityRemoteAccessDestroy(s *terraform.State) error {
	return nil
}

func testAccCheckCheckpointManagementVpnCommunityRemoteAccessExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("VpnCommunityMeshed ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-vpn-community-remote-access", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementVpnCommunityRemoteAccessAttributes(vpnCommunityRemoteAccessMap *map[string]interface{}, name string, userGroup string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		vpnCommunityRemoteAccessName := (*vpnCommunityRemoteAccessMap)["name"].(string)
		if !strings.EqualFold(vpnCommunityRemoteAccessName, name) {
			return fmt.Errorf("name is %s, expected %s", vpnCommunityRemoteAccessName, name)
		}
		return nil
	}
}

func testAccManagementVpnCommunityRemoteAccessConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_vpn_community_remote_access" "test" {
        name = "%s"
		user_groups = ["All Users"]
}
`, name)
}
