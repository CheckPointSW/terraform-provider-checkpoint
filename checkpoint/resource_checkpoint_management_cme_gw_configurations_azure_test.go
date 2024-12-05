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
	gwConfigurationVersion := "R82"
	gwConfigurationBase64SIC := "MTIzNDU2Nzg="
	gwConfigurationPolicy := "Standard"
	gwConfigurationIpv6 := true
	gwConfigurationColor := "black"
	gwConfigurationXForwardedFor := true
	gwConfigurationCommunicationWithServersBehindNAT := "translated-ip-only"



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
					gwConfigurationBase64SIC, gwConfigurationPolicy, gwConfigurationIpv6, gwConfigurationXForwardedFor,
					gwConfigurationColor, gwConfigurationCommunicationWithServersBehindNAT),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementCMEGWConfigurationsAzureExists(resourceName, &azureGWConfiguration),
					testAccCheckCheckpointManagementCMEGWConfigurationsAzureAttributes(&azureGWConfiguration, gwConfigurationName, accountName, gwConfigurationVersion,
						gwConfigurationPolicy, true, true, true, gwConfigurationIpv6, gwConfigurationXForwardedFor,
						gwConfigurationColor, gwConfigurationCommunicationWithServersBehindNAT),
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
	gwConfigurationBase64SIC string, gwConfigurationPolicy string, gwConfigurationIpv6 bool, gwConfigurationXForwardedFor bool,
	gwConfigurationColor string, gwConfigurationCommunicationWithServersBehindNAT string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_cme_accounts_azure" "account_test" {
  name           = "%s"
  directory_id   = "46707d92-02f4-4817-8116-a4c3b23e6266"
  application_id = "46707d92-02f4-4817-8116-a4c3b23e6266"
  client_secret  = "abcdef-123456"
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
	identity_awareness = true
	ipsec_vpn = false
	threat_emulation = false
	url_filtering = false
	vpn = false
  }
  identity_awareness_settings {
    enable_cloudguard_controller = true
  }
  ipv6 = %t
  x_forwarded_for = %t
  color           = "%s"
  communication_with_servers_behind_nat = "%s"
}
`, accountName, gwConfigurationName, gwConfigurationVersion, gwConfigurationBase64SIC, gwConfigurationPolicy, gwConfigurationIpv6, gwConfigurationXForwardedFor,
  gwConfigurationColor, gwConfigurationCommunicationWithServersBehindNAT)
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
	applicationControlFlag bool, IDAFlag bool, gwConfigurationIpv6 bool, gwConfigurationXForwardedFor bool,
    gwConfigurationColor string, gwConfigurationCommunicationWithServersBehindNAT string) resource.TestCheckFunc {
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
		IDA := blades["identity-awareness"]
		if httpsInspection != httpsInspectionFlag {
			return fmt.Errorf("https inspection is %t, expected %t", httpsInspection, httpsInspectionFlag)
		}
		if applicationControl != applicationControlFlag {
			return fmt.Errorf("application control is %t, expected %t", applicationControl, applicationControlFlag)
		}
		if IDA != IDAFlag {
			return fmt.Errorf("identity awareness is %t, expected %t", IDA, IDAFlag)
		}
		IDASettings := gwConfiguration["identity-awareness-settings"].(map[string]interface{})
		enableCgController := IDASettings["enable-cloudguard-controller"]
		if enableCgController != IDAFlag{
			return fmt.Errorf("enable-cloudguard-controller identity source is %t, expected %t", enableCgController, IDAFlag)
		}
		if gwConfiguration["ipv6"] != gwConfigurationIpv6 {
			return fmt.Errorf("ipv6 is %t, expected %t", gwConfiguration["ipv6"], gwConfigurationIpv6)
		}
		if gwConfiguration["x_forwarded_for"] != gwConfigurationXForwardedFor {
			return fmt.Errorf("x_forwarded_for is %t, expected %t", gwConfiguration["x_forwarded_for"], gwConfigurationXForwardedFor)
		}
		if gwConfiguration["color"] != gwConfigurationColor {
			return fmt.Errorf("color is %s, expected %s", gwConfiguration["color"], gwConfigurationColor)
		}
		if gwConfiguration["communication-with-servers-behind-nat"] != gwConfigurationCommunicationWithServersBehindNAT {
			return fmt.Errorf("communication_with_servers_behind_nat is %s, expected %s", gwConfiguration["communication_with_servers_behind_nat"], gwConfigurationCommunicationWithServersBehindNAT)
		}
		return nil
	}
}
