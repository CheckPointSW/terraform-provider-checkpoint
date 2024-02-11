package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementCMEGWConfigurationsGCP_basic(t *testing.T) {
	resourceName := "checkpoint_management_cme_gw_configurations_gcp.test"
	dataSourceName := "data.checkpoint_management_cme_gw_configurations_gcp.data_test"
	gwConfigurationName := "test-gw-configuration"
	accountName := "test-account"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this test")
	} else if context != "web_api" {
		t.Skip("Skipping cme api test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementCMEGWConfigurationsGCPConfig(accountName, gwConfigurationName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "related_account", resourceName, "related_account"),
					resource.TestCheckResourceAttrPair(dataSourceName, "version", resourceName, "version"),
				),
			},
		},
	})
}

func testAccDataSourceManagementCMEGWConfigurationsGCPConfig(accountName string, gwConfigurationName string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_cme_accounts_gcp" "gcp_account" {
  name                  = "%s"
  project_id       = "my-project-1" 
  credentials_file = "LocalGWSetMap.json"
}

resource "checkpoint_management_cme_gw_configurations_gcp" "test" {
  name            = "%s"
  related_account = "${checkpoint_management_cme_accounts_gcp.gcp_account.name}"
  version         = "R81"
  base64_sic_key  = "MTIzNDU2Nzg="
  policy          = "Standard"
  blades {
	ips                          = false
	anti_bot                     = false
	anti_virus                   = false
	https_inspection             = false
	application_control          = false
	autonomous_threat_prevention = false
	content_awareness            = false
	identity_awareness           = false
	ipsec_vpn                    = false
	threat_emulation             = false
	url_filtering                = false
	vpn                          = false
  }
}

data "checkpoint_management_cme_gw_configurations_gcp" "data_test" {
  name = "${checkpoint_management_cme_gw_configurations_gcp.test.name}"
}
`, accountName, gwConfigurationName)
}
