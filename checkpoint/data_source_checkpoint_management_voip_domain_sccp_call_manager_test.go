package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementVoipDomainSccpCallManager_basic(t *testing.T) {
	resourceName := "checkpoint_management_voip_domain_sccp_call_manager.test"
	dataSourceName := "data.checkpoint_management_voip_domain_sccp_call_manager.data_test"
	objectName := "test-voip_domain_sccp_call_manager"

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
				Config: testAccDataSourceManagementVoipDomainSccpCallManagerConfig(objectName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "color", resourceName, "color"),
					resource.TestCheckResourceAttrPair(dataSourceName, "comments", resourceName, "comments"),
				),
			},
		},
	})
}

func testAccDataSourceManagementVoipDomainSccpCallManagerConfig(objectName string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_group" "ds_sccp_group" {
	name = "ds-sccp-group-%s"
}

resource "checkpoint_management_host" "ds_sccp_host" {
	name = "ds-sccp-host-%s"
	ipv4_address = "192.0.2.34"
	ignore_warnings = true
}

resource "checkpoint_management_voip_domain_sccp_call_manager" "test" {
	name = "%s"
	color = "blue"
	comments = "test-value"
	endpoints_domain = "${checkpoint_management_group.ds_sccp_group.name}"
	installed_at = "${checkpoint_management_host.ds_sccp_host.name}"
}

data "checkpoint_management_voip_domain_sccp_call_manager" "data_test" {
	name = "${checkpoint_management_voip_domain_sccp_call_manager.test.name}"
}
`, objectName, objectName, objectName)
}
