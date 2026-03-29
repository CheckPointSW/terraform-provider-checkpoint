package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementVoipDomainH323Gateway_basic(t *testing.T) {
	resourceName := "checkpoint_management_voip_domain_h323_gateway.test"
	dataSourceName := "data.checkpoint_management_voip_domain_h323_gateway.data_test"
	objectName := "test-voip_domain_h323_gateway"

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
				Config: testAccDataSourceManagementVoipDomainH323GatewayConfig(objectName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "color", resourceName, "color"),
					resource.TestCheckResourceAttrPair(dataSourceName, "comments", resourceName, "comments"),
				),
			},
		},
	})
}

func testAccDataSourceManagementVoipDomainH323GatewayConfig(objectName string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_group" "ds_gateway_group" {
	name = "ds-gateway-group-%s"
}

resource "checkpoint_management_host" "ds_gateway_host" {
	name = "ds-gateway-host-%s"
	ipv4_address = "192.0.2.32"
	ignore_warnings = true
}

resource "checkpoint_management_voip_domain_h323_gateway" "test" {
	name = "%s"
	color = "blue"
	comments = "test-value"
	endpoints_domain = "${checkpoint_management_group.ds_gateway_group.name}"
	installed_at = "${checkpoint_management_host.ds_gateway_host.name}"
}

data "checkpoint_management_voip_domain_h323_gateway" "data_test" {
	name = "${checkpoint_management_voip_domain_h323_gateway.test.name}"
}
`, objectName, objectName, objectName)
}
