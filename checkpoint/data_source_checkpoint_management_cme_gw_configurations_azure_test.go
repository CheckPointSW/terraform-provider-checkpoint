package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementCMEGWConfigurationsAzure_basic(t *testing.T) {
	resourceName := "checkpoint_management_cme_gw_configurations_azure.test"
	dataSourceName := "data.checkpoint_management_cme_gw_configurations_azure.data_test"
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
				Config: testAccDataSourceManagementCMEGWConfigurationsAzureConfig(accountName, gwConfigurationName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "related_account", resourceName, "related_account"),
					resource.TestCheckResourceAttrPair(dataSourceName, "version", resourceName, "version"),
				),
			},
		},
	})
}

func testAccDataSourceManagementCMEGWConfigurationsAzureConfig(accountName string, gwConfigurationName string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_cme_accounts_azure" "azure_account" {
  name                  = "%s"
  directory_id = "46707d92-02f4-4817-8116-a4c3b23e6266"
  application_id = "46707d92-02f4-4817-8116-a4c3b23e6266"
  client_secret = "mySecret"
  subscription = "46707d92-02f4-4817-8116-a4c3b23e6267"
}

resource "checkpoint_management_cme_gw_configurations_azure" "test" {
  name            = "%s"
  related_account = "${checkpoint_management_cme_accounts_azure.azure_account.name}"
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

data "checkpoint_management_cme_gw_configurations_azure" "data_test" {
  name = "${checkpoint_management_cme_gw_configurations_azure.test.name}"
}
`, accountName, gwConfigurationName)
}
