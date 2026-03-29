package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementVpnCommunityStar_basic(t *testing.T) {

	objName := "tfTestManagementDataVpnCommunityStar_" + acctest.RandString(6)
	resourceName := "checkpoint_management_vpn_community_star.vpn_community_star"
	dataSourceName := "data.checkpoint_management_data_vpn_community_star.data_vpn_community_star"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementVpnCommunityStarConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})
}

func testAccDataSourceManagementVpnCommunityStarConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_vpn_community_star" "vpn_community_star" {
    name = "%s"
	encryption_method = "ikev1 for ipv4 and ikev2 for ipv6 only"
	encryption_suite = "custom"
}

data "checkpoint_management_data_vpn_community_star" "data_vpn_community_star" {
    name = "${checkpoint_management_vpn_community_star.vpn_community_star.name}"
}
`, name)
}
