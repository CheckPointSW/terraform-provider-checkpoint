package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	"testing"
)

func TestAccCheckpointManagementCMEGWConfigurationsAzure_basic(t *testing.T) {
	var azureGWConfiguration map[string]interface{}
	resourceName := "checkpoint_management_cme_gw_configurations_azure.gw_configuration_test"
	accountName := "test-account"
	gwConfigurationName := "test-gw-configuration"
	gwConfigurationVersion := "R81.10"
	gwConfigurationBase64SIC := "MTIzNDU2Nzg="
	gwConfigurationPolicy := "Standard"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this test")
	} else if context != "web_api" {
		t.Skip("Skipping cme api test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementCMEGWConfigurationsAzureDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementCMEGWConfigurationsAzureConfig(accountName, gwConfigurationName, gwConfigurationVersion,
					gwConfigurationBase64SIC, gwConfigurationPolicy),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementCMEGWConfigurationsAzureExists(resourceName, &azureGWConfiguration),
					testAccCheckCheckpointManagementCMEGWConfigurationsAzureAttributes(&azureGWConfiguration, gwConfigurationName, accountName, gwConfigurationVersion,
						gwConfigurationPolicy, true, true),
				),
			},
		},
	})
}
func testAccCheckpointManagementCMEGWConfigurationsAzureDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_cme_gw_configurations_azure" {
			continue
		}
		if rs.Primary.ID != "" {
			url := CmeApiPath + "/gwConfigurations/" + rs.Primary.Attributes["name"]
			response, err := client.ApiCall(url, nil, client.GetSessionID(), true, client.IsProxyUsed(), "GET")
			if err != nil {
				return err
			}
			res := response.GetData()
			if !checkIfRequestFailed(res) {
				return fmt.Errorf("Azure gw configuration (%s) still exists", rs.Primary.Attributes["name"])
			}
		}
		return nil
	}
	return nil
}

func testAccManagementCMEGWConfigurationsAzureConfig(accountName string, gwConfigurationName string, gwConfigurationVersion string,
	gwConfigurationBase64SIC string, gwConfigurationPolicy string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_cme_accounts_azure" "account_test" {
  name           = "%s"
  directory_id   = "46707d92-02f4-4817-8116-a4c3b23e6266"
  application_id = "46707d92-02f4-4817-8116-a4c3b23e6266"
  client_secret  = "mySecret"
  subscription   = "46707d92-02f4-4817-8116-a4c3b23e6266"
}

resource "checkpoint_management_cme_gw_configurations_azure" "gw_configuration_test" {
  name            = "%s"
  related_account = checkpoint_management_cme_accounts_azure.account_test.name
  version         = "%s"
  base64_sic_key  = "%s"
  policy          = "%s"
  blades {
    https_inspection = true
    application_control = true
    ips      = false
    anti_bot = false
	anti_virus = false
	autonomous_threat_prevention = false
	content_awareness = false
	identity_awareness = false
	ipsec_vpn = false
	threat_emulation = false
	url_filtering = false
	vpn = false
  }
}
`, accountName, gwConfigurationName, gwConfigurationVersion, gwConfigurationBase64SIC, gwConfigurationPolicy)
}

func testAccCheckCheckpointManagementCMEGWConfigurationsAzureExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)
		url := CmeApiPath + "/gwConfigurations/" + rs.Primary.Attributes["name"]
		response, err := client.ApiCall(url, nil, client.GetSessionID(), true, client.IsProxyUsed(), "GET")
		if err != nil {
			return err
		}

		*res = response.GetData()
		if checkIfRequestFailed(*res) {
			errMessage := buildErrorMessage(*res)
			return fmt.Errorf(errMessage)
		}
		return nil
	}
}

func testAccCheckCheckpointManagementCMEGWConfigurationsAzureAttributes(azureGWConfiguration *map[string]interface{}, gwConfigurationName string,
	accountName string, gwConfigurationVersion string, gwConfigurationPolicyName string, httpsInspectionFlag bool,
	applicationControlFlag bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		gwConfiguration := (*azureGWConfiguration)["result"].(map[string]interface{})
		if gwConfiguration["name"] != gwConfigurationName {
			return fmt.Errorf("name is %s, expected %s", gwConfiguration["name"], gwConfigurationName)
		}
		if gwConfiguration["related_account"] != accountName {
			return fmt.Errorf("related account name is %s, expected %s", gwConfiguration["related_account"], accountName)
		}
		if gwConfiguration["version"] != gwConfigurationVersion {
			return fmt.Errorf("version is %s, expected %s", gwConfiguration["version"], gwConfigurationVersion)
		}
		if gwConfiguration["policy"] != gwConfigurationPolicyName {
			return fmt.Errorf("policy is %s, expected %s", gwConfiguration["policy"], gwConfigurationPolicyName)
		}
		blades := gwConfiguration["blades"].(map[string]interface{})
		httpsInspection := blades["https-inspection"]
		applicationControl := blades["application-control"]
		if httpsInspection != httpsInspectionFlag {
			return fmt.Errorf("https inspection is %t, expected %t", httpsInspection, httpsInspectionFlag)
		}
		if applicationControl != applicationControlFlag {
			return fmt.Errorf("application control is %t, expected %t", applicationControl, applicationControlFlag)
		}
		return nil
	}
}
