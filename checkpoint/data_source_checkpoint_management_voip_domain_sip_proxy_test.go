package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementVoipDomainSipProxy_basic(t *testing.T) {
	resourceName := "checkpoint_management_voip_domain_sip_proxy.test"
	dataSourceName := "data.checkpoint_management_voip_domain_sip_proxy.data_test"
	objectName := "test-voip_domain_sip_proxy"

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
				Config: testAccDataSourceManagementVoipDomainSipProxyConfig(objectName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "color", resourceName, "color"),
					resource.TestCheckResourceAttrPair(dataSourceName, "comments", resourceName, "comments"),
				),
			},
		},
	})
}

func testAccDataSourceManagementVoipDomainSipProxyConfig(objectName string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_group" "ds_sip_group" {
	name = "ds-sip-group-%s"
}

resource "checkpoint_management_host" "ds_sip_host" {
	name = "ds-sip-host-%s"
	ipv4_address = "192.0.2.30"
	ignore_warnings = true
}

resource "checkpoint_management_voip_domain_sip_proxy" "test" {
	name = "%s"
	color = "blue"
	comments = "test-value"
	endpoints_domain = "${checkpoint_management_group.ds_sip_group.name}"
	installed_at = "${checkpoint_management_host.ds_sip_host.name}"
}

data "checkpoint_management_voip_domain_sip_proxy" "data_test" {
	name = "${checkpoint_management_voip_domain_sip_proxy.test.name}"
}
`, objectName, objectName, objectName)
}
