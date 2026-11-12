package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementVpnCommunityRemoteAccess_basic(t *testing.T) {

	objName := "RemoteAccess"
	resourceName := "checkpoint_management_vpn_community_remote_access.vpn_community_remote_access"
	dataSourceName := "data.checkpoint_management_vpn_community_remote_access.data_vpn_community_remote_access"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementVpnCommunityRemoteAccessConfig(objName, "All Users"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "user_groups", resourceName, "user_groups"),
				),
			},
		},
	})
}

func testAccDataSourceManagementVpnCommunityRemoteAccessConfig(name string, userGroups string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_vpn_community_remote_access" "vpn_community_remote_access" {
    name = "%s"
	user_groups = ["%s"]
}

data "checkpoint_management_vpn_community_remote_access" "data_vpn_community_remote_access" {
    name = "${checkpoint_management_vpn_community_remote_access.vpn_community_remote_access.name}"
}
`, name, userGroups)
}
