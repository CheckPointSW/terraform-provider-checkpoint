package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	"testing"
)

func TestAccCheckpointManagementCMEGWConfigurationsGCP_basic(t *testing.T) {
	var gcpGWConfiguration map[string]interface{}
	resourceName := "checkpoint_management_cme_gw_configurations_gcp.gw_configuration_test"
	accountName := "test-account"
	gwConfigurationName := "test-gw-configuration"
	gwConfigurationVersion := "R81.20"
	gwConfigurationBase64SIC := "MTIzNDU2Nzg="
	gwConfigurationPolicy := "Standard"
	gwConfigurationColor := "blue"
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
		CheckDestroy: testAccCheckpointManagementCMEGWConfigurationsGCPDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementCMEGWConfigurationsGCPConfig(accountName, gwConfigurationName, gwConfigurationVersion,
					gwConfigurationBase64SIC, gwConfigurationPolicy, gwConfigurationXForwardedFor,
					gwConfigurationColor, gwConfigurationCommunicationWithServersBehindNAT),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementCMEGWConfigurationsGCPExists(resourceName, &gcpGWConfiguration),
					testAccCheckCheckpointManagementCMEGWConfigurationsGCPAttributes(&gcpGWConfiguration, gwConfigurationName, accountName, gwConfigurationVersion,
						gwConfigurationPolicy, true, true, gwConfigurationXForwardedFor,
						gwConfigurationColor, gwConfigurationCommunicationWithServersBehindNAT),
				),
			},
		},
	})
}
func testAccCheckpointManagementCMEGWConfigurationsGCPDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_cme_gw_configurations_gcp" {
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
				return fmt.Errorf("GCP gw configuration (%s) still exists", rs.Primary.Attributes["name"])
			}
		}
		return nil
	}
	return nil
}

func testAccManagementCMEGWConfigurationsGCPConfig(accountName string, gwConfigurationName string, gwConfigurationVersion string,
	gwConfigurationBase64SIC string, gwConfigurationPolicy string,  gwConfigurationXForwardedFor bool,
	gwConfigurationColor string, gwConfigurationCommunicationWithServersBehindNAT string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_cme_accounts_gcp" "account_test" {
  name           = "%s"
  project_id   = "my-project-1"
  credentials_file = "LocalGWSetMap.json"
}

resource "checkpoint_management_cme_gw_configurations_gcp" "gw_configuration_test" {
  name            = "%s"
  related_account = checkpoint_management_cme_accounts_gcp.account_test.name
  version         = "%s"
  base64_sic_key  = "%s"
  policy          = "%s"
  x_forwarded_for =  %t
  color           = "%s"
  communication_with_servers_behind_nat = "%s"
  blades {
    content_awareness = true
    identity_awareness = true
	https_inspection = false
    application_control = false
    ips      = false
    anti_bot = false
	anti_virus = false
	autonomous_threat_prevention = false
	ipsec_vpn = false
	threat_emulation = false
	url_filtering = false
	vpn = false
  }
}
`, accountName, gwConfigurationName, gwConfigurationVersion, gwConfigurationBase64SIC, gwConfigurationPolicy,  gwConfigurationXForwardedFor,
   gwConfigurationColor, gwConfigurationCommunicationWithServersBehindNAT)
}

func testAccCheckCheckpointManagementCMEGWConfigurationsGCPExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
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

func testAccCheckCheckpointManagementCMEGWConfigurationsGCPAttributes(gcpGWConfiguration *map[string]interface{}, gwConfigurationName string,
	accountName string, gwConfigurationVersion string, gwConfigurationPolicyName string, contentAwarenessFlag bool,
	identityAwarenessFlag bool, gwConfigurationXForwardedFor bool, gwConfigurationColor string,
    gwConfigurationCommunicationWithServersBehindNAT string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		gwConfiguration := (*gcpGWConfiguration)["result"].(map[string]interface{})
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
		contentAwareness := blades["content-awareness"]
		identityAwareness := blades["identity-awareness"]
		if contentAwareness != contentAwarenessFlag {
			return fmt.Errorf("content awareness is %t, expected %t", contentAwareness, contentAwarenessFlag)
		}
		if identityAwareness != identityAwarenessFlag {
			return fmt.Errorf("identity awareness is %t, expected %t", identityAwareness, identityAwarenessFlag)
		}
		if gwConfiguration["x_forwarded_for"] != gwConfigurationXForwardedFor {
			return fmt.Errorf("x_forwarded_for is %t, expected %t", gwConfiguration["x_forwarded_for"], gwConfigurationXForwardedFor)
		}
		if gwConfiguration["color"] != gwConfigurationColor {
			return fmt.Errorf("color is %s, expected %s", gwConfiguration["color"], gwConfigurationColor)
		}
		if gwConfiguration["communication_with_servers_behind_nat"] != gwConfigurationCommunicationWithServersBehindNAT {
			return fmt.Errorf("communication_with_servers_behind_nat is %s, expected %s", gwConfiguration["communication_with_servers_behind_nat"], gwConfigurationCommunicationWithServersBehindNAT)
		}
		return nil
	}
}
