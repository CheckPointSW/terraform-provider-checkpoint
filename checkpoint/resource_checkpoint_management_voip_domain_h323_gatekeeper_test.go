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

func TestAccCheckpointManagementVoipDomainH323Gatekeeper_basic(t *testing.T) {

	var voipDomainH323GatekeeperMap map[string]interface{}
	resourceName := "checkpoint_management_voip_domain_h323_gatekeeper.test"
	objName := "tfTestManagementVoipDomainH323Gatekeeper_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementVoipDomainH323GatekeeperDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementVoipDomainH323GatekeeperConfig(objName, "new_group5", "test_host5"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementVoipDomainH323GatekeeperExists(resourceName, &voipDomainH323GatekeeperMap),
					testAccCheckCheckpointManagementVoipDomainH323GatekeeperAttributes(&voipDomainH323GatekeeperMap, objName, "new_group5", "test_host5"),
				),
			},
		},
	})
}

func testAccCheckpointManagementVoipDomainH323GatekeeperDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_voip_domain_h323_gatekeeper" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-voip-domain-h323-gatekeeper", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("VoipDomainH323Gatekeeper object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementVoipDomainH323GatekeeperExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("VoipDomainH323Gatekeeper ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-voip-domain-h323-gatekeeper", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementVoipDomainH323GatekeeperAttributes(voipDomainH323GatekeeperMap *map[string]interface{}, name string, endpointsDomain string, installedAt string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		voipDomainH323GatekeeperName := (*voipDomainH323GatekeeperMap)["name"].(string)
		if !strings.EqualFold(voipDomainH323GatekeeperName, name) {
			return fmt.Errorf("name is %s, expected %s", name, voipDomainH323GatekeeperName)
		}
		voipDomainH323GatekeeperEndpointsDomain := (*voipDomainH323GatekeeperMap)["endpoints-domain"].(map[string]interface{})["name"].(string)
		if !strings.EqualFold(voipDomainH323GatekeeperEndpointsDomain, endpointsDomain) {
			return fmt.Errorf("endpointsDomain is %s, expected %s", endpointsDomain, voipDomainH323GatekeeperEndpointsDomain)
		}
		voipDomainH323GatekeeperInstalledAt := (*voipDomainH323GatekeeperMap)["installed-at"].(map[string]interface{})["name"].(string)
		if !strings.EqualFold(voipDomainH323GatekeeperInstalledAt, installedAt) {
			return fmt.Errorf("installedAt is %s, expected %s", installedAt, voipDomainH323GatekeeperInstalledAt)
		}
		return nil
	}
}

func testAccManagementVoipDomainH323GatekeeperConfig(name string, endpointsDomain string, installedAt string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_group" "group1" {
  name = "%s"
}

resource "checkpoint_management_host" "host1" {
  name = "%s"
  ipv4_address = "192.0.2.6"
}

resource "checkpoint_management_voip_domain_h323_gatekeeper" "test" {
        name = "%s"
        endpoints_domain = "${checkpoint_management_group.group1.name}"
        installed_at = "${checkpoint_management_host.host1.name}"
        routing_mode {
          direct = false
          call_setup = false
          call_setup_and_call_control = false
        }
}
`, endpointsDomain, installedAt, name)
}
