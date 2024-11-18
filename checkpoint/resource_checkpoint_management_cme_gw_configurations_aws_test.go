package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	"testing"
)

func TestAccCheckpointManagementCMEGWConfigurationsAWS_basic(t *testing.T) {
	var awsGWConfiguration map[string]interface{}
	resourceName := "checkpoint_management_cme_gw_configurations_aws.gw_configuration_test"
	accountName := "test-account"
	gwConfigurationName := "test-gw-configuration"
	gwConfigurationVersion := "R82"
	gwConfigurationBase64SIC := "MTIzNDU2Nzg="
	gwConfigurationPolicy := "Standard"
	gwConfigurationXForwardedFor := true
	gwConfigurationColor := "black"
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
		CheckDestroy: testAccCheckpointManagementCMEGWConfigurationsAWSDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementCMEGWConfigurationsAWSConfig(accountName, gwConfigurationName, gwConfigurationVersion,
					gwConfigurationBase64SIC, gwConfigurationPolicy, gwConfigurationXForwardedFor, gwConfigurationColor,
					gwConfigurationCommunicationWithServersBehindNAT),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementCMEGWConfigurationsAWSExists(resourceName, &awsGWConfiguration),
					testAccCheckCheckpointManagementCMEGWConfigurationsAWSAttributes(&awsGWConfiguration, gwConfigurationName, accountName, gwConfigurationVersion,
						gwConfigurationPolicy, true, true, true, gwConfigurationXForwardedFor,
						gwConfigurationColor, gwConfigurationCommunicationWithServersBehindNAT),
				),
			},
		},
	})
}
func testAccCheckpointManagementCMEGWConfigurationsAWSDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_cme_gw_configurations_aws" {
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
				return fmt.Errorf("AWS gw configuration (%s) still exists", rs.Primary.Attributes["name"])
			}
		}
		return nil
	}
	return nil
}

func testAccManagementCMEGWConfigurationsAWSConfig(accountName string, gwConfigurationName string, gwConfigurationVersion string,
	gwConfigurationBase64SIC string, gwConfigurationPolicy string, gwConfigurationXForwardedFor bool, gwConfigurationColor string,
	gwConfigurationCommunicationWithServersBehindNAT string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_cme_accounts_aws" "account_test" {
  name                  = "%s"
  regions               = ["us-east-1"]
  credentials_file      = "IAM"
}

resource "checkpoint_management_cme_gw_configurations_aws" "gw_configuration_test" {
  name                  = "%s"
  related_account = checkpoint_management_cme_accounts_aws.account_test.name
  version         = "%s"
  base64_sic_key  = "%s"
  policy          = "%s"
  x_forwarded_for = %t
  color           = "%s"
  communication_with_servers_behind_nat = "%s"
  blades {
    ips      = true
    anti_bot = true
	anti_virus = false
	https_inspection = false
	application_control = false
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
}
`, accountName, gwConfigurationName, gwConfigurationVersion, gwConfigurationBase64SIC, gwConfigurationPolicy, gwConfigurationXForwardedFor,
   gwConfigurationColor, gwConfigurationCommunicationWithServersBehindNAT)
}

func testAccCheckCheckpointManagementCMEGWConfigurationsAWSExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
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

func testAccCheckCheckpointManagementCMEGWConfigurationsAWSAttributes(awsGWConfiguration *map[string]interface{}, gwConfigurationName string,
	accountName string, gwConfigurationVersion string, gwConfigurationPolicyName string, ipsFlag bool,
	antiBotFlag bool, IDAFlag bool, gwConfigurationXForwardedFor bool, gwConfigurationColor string, gwConfigurationCommunicationWithServersBehindNAT string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		gwConfiguration := (*awsGWConfiguration)["result"].(map[string]interface{})
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
		ips := blades["ips"]
		antiBot := blades["anti-bot"]
		IDA := blades["identity-awareness"]
		if ips != ipsFlag {
			return fmt.Errorf("ips is %t, expected %t", ips, ipsFlag)
		}
		if antiBot != antiBotFlag {
			return fmt.Errorf("anti bot is %t, expected %t", antiBot, antiBotFlag)
		}
		if IDA != IDAFlag {
			return fmt.Errorf("identity awareness is %t, expected %t", IDA, IDAFlag)
		}
		IDASettings := gwConfiguration["identity-awareness-settings"].(map[string]interface{})
		enableCgController := IDASettings["enable-cloudguard-controller"]
		if enableCgController != IDAFlag{
			return fmt.Errorf("enable-cloudguard-controller identity source is %t, expected %t", enableCgController, IDAFlag)
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
