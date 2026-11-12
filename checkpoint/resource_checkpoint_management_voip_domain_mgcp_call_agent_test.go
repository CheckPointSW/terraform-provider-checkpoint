package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"os"
	"strings"
	"testing"
)

func TestAccCheckpointManagementVoipDomainMgcpCallAgent_basic(t *testing.T) {

	var voipDomainMgcpCallAgentMap map[string]interface{}
	resourceName := "checkpoint_management_voip_domain_mgcp_call_agent.test"
	objName := "tfTestManagementVoipDomainMgcpCallAgent_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementVoipDomainMgcpCallAgentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementVoipDomainMgcpCallAgentConfig(objName, "new_group13", "test_host13"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementVoipDomainMgcpCallAgentExists(resourceName, &voipDomainMgcpCallAgentMap),
					testAccCheckCheckpointManagementVoipDomainMgcpCallAgentAttributes(&voipDomainMgcpCallAgentMap, objName, "new_group13", "test_host13"),
				),
			},
		},
	})
}

func testAccCheckpointManagementVoipDomainMgcpCallAgentDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_voip_domain_mgcp_call_agent" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-voip-domain-mgcp-call-agent", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("VoipDomainMgcpCallAgent object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementVoipDomainMgcpCallAgentExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("VoipDomainMgcpCallAgent ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-voip-domain-mgcp-call-agent", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementVoipDomainMgcpCallAgentAttributes(voipDomainMgcpCallAgentMap *map[string]interface{}, name string, endpointsDomain string, installedAt string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		voipDomainMgcpCallAgentName := (*voipDomainMgcpCallAgentMap)["name"].(string)
		if !strings.EqualFold(voipDomainMgcpCallAgentName, name) {
			return fmt.Errorf("name is %s, expected %s", name, voipDomainMgcpCallAgentName)
		}
		voipDomainMgcpCallAgentEndpointsDomain := (*voipDomainMgcpCallAgentMap)["endpoints-domain"].(map[string]interface{})["name"].(string)
		if !strings.EqualFold(voipDomainMgcpCallAgentEndpointsDomain, endpointsDomain) {
			return fmt.Errorf("endpointsDomain is %s, expected %s", endpointsDomain, voipDomainMgcpCallAgentEndpointsDomain)
		}
		voipDomainMgcpCallAgentInstalledAt := (*voipDomainMgcpCallAgentMap)["installed-at"].(map[string]interface{})["name"].(string)
		if !strings.EqualFold(voipDomainMgcpCallAgentInstalledAt, installedAt) {
			return fmt.Errorf("installedAt is %s, expected %s", installedAt, voipDomainMgcpCallAgentInstalledAt)
		}
		return nil
	}
}

func testAccManagementVoipDomainMgcpCallAgentConfig(name string, endpointsDomain string, installedAt string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_group" "group1" {
  name = "%s"
}

resource "checkpoint_management_host" "host1" {
  name = "%s"
  ipv4_address = "192.0.2.7"
}

resource "checkpoint_management_voip_domain_mgcp_call_agent" "test" {
        name = "%s"
        endpoints_domain = "${checkpoint_management_group.group1.name}"
        installed_at = "${checkpoint_management_host.host1.name}"
}
`, endpointsDomain, installedAt, name)
}
