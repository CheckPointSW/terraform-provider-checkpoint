package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementVoipDomainH323Gatekeeper_basic(t *testing.T) {
	resourceName := "checkpoint_management_voip_domain_h323_gatekeeper.test"
	dataSourceName := "data.checkpoint_management_voip_domain_h323_gatekeeper.data_test"
	objectName := "test-voip_domain_h323_gatekeeper"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementVoipDomainH323GatekeeperConfig(objectName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "color", resourceName, "color"),
					resource.TestCheckResourceAttrPair(dataSourceName, "comments", resourceName, "comments"),
				),
			},
		},
	})
}

func testAccDataSourceManagementVoipDomainH323GatekeeperConfig(objectName string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_group" "ds_gatekeeper_group" {
	name = "ds-gatekeeper-group-%s"
}

resource "checkpoint_management_host" "ds_gatekeeper_host" {
	name = "ds-gatekeeper-host-%s"
	ipv4_address = "192.0.2.31"
	ignore_warnings = true
}

resource "checkpoint_management_voip_domain_h323_gatekeeper" "test" {
	name = "%s"
	color = "blue"
	comments = "test-value"
	endpoints_domain = "${checkpoint_management_group.ds_gatekeeper_group.name}"
	installed_at = "${checkpoint_management_host.ds_gatekeeper_host.name}"
}

data "checkpoint_management_voip_domain_h323_gatekeeper" "data_test" {
	name = "${checkpoint_management_voip_domain_h323_gatekeeper.test.name}"
}
`, objectName, objectName, objectName)
}
