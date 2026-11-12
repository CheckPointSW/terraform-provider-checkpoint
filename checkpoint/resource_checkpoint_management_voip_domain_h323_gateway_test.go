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

func TestAccCheckpointManagementVoipDomainH323Gateway_basic(t *testing.T) {

	var voipDomainH323GatewayMap map[string]interface{}
	resourceName := "checkpoint_management_voip_domain_h323_gateway.test"
	objName := "tfTestManagementVoipDomainH323Gateway_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementVoipDomainH323GatewayDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementVoipDomainH323GatewayConfig(objName, "new_group10", "test_host10"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementVoipDomainH323GatewayExists(resourceName, &voipDomainH323GatewayMap),
					testAccCheckCheckpointManagementVoipDomainH323GatewayAttributes(&voipDomainH323GatewayMap, objName, "new_group10", "test_host10"),
				),
			},
		},
	})
}

func testAccCheckpointManagementVoipDomainH323GatewayDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_voip_domain_h323_gateway" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-voip-domain-h323-gateway", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("VoipDomainH323Gateway object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementVoipDomainH323GatewayExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("VoipDomainH323Gateway ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-voip-domain-h323-gateway", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementVoipDomainH323GatewayAttributes(voipDomainH323GatewayMap *map[string]interface{}, name string, endpointsDomain string, installedAt string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		voipDomainH323GatewayName := (*voipDomainH323GatewayMap)["name"].(string)
		if !strings.EqualFold(voipDomainH323GatewayName, name) {
			return fmt.Errorf("name is %s, expected %s", name, voipDomainH323GatewayName)
		}
		voipDomainH323GatewayEndpointsDomain := (*voipDomainH323GatewayMap)["endpoints-domain"].(map[string]interface{})["name"].(string)
		if !strings.EqualFold(voipDomainH323GatewayEndpointsDomain, endpointsDomain) {
			return fmt.Errorf("endpointsDomain is %s, expected %s", endpointsDomain, voipDomainH323GatewayEndpointsDomain)
		}
		voipDomainH323GatewayInstalledAt := (*voipDomainH323GatewayMap)["installed-at"].(map[string]interface{})["name"].(string)
		if !strings.EqualFold(voipDomainH323GatewayInstalledAt, installedAt) {
			return fmt.Errorf("installedAt is %s, expected %s", installedAt, voipDomainH323GatewayInstalledAt)
		}
		return nil
	}
}

func testAccManagementVoipDomainH323GatewayConfig(name string, endpointsDomain string, installedAt string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_group" "group1" {
  name = "%s"
}

resource "checkpoint_management_host" "host1" {
  name = "%s"
  ipv4_address = "192.0.2.6"
}

resource "checkpoint_management_voip_domain_h323_gateway" "test" {
        name = "%s"
        endpoints_domain = "${checkpoint_management_group.group1.name}"
        installed_at = "${checkpoint_management_host.host1.name}"
        routing_mode {
          call_setup = false
          call_setup_and_call_control = false
        }
}
`, endpointsDomain, installedAt, name)
}
