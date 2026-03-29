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

func TestAccCheckpointManagementVoipDomainSccpCallManager_basic(t *testing.T) {

	var voipDomainSccpCallManagerMap map[string]interface{}
	resourceName := "checkpoint_management_voip_domain_sccp_call_manager.test"
	objName := "tfTestManagementVoipDomainSccpCallManager_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementVoipDomainSccpCallManagerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementVoipDomainSccpCallManagerConfig(objName, "new_group15", "test_host15"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementVoipDomainSccpCallManagerExists(resourceName, &voipDomainSccpCallManagerMap),
					testAccCheckCheckpointManagementVoipDomainSccpCallManagerAttributes(&voipDomainSccpCallManagerMap, objName, "new_group15", "test_host15"),
				),
			},
		},
	})
}

func testAccCheckpointManagementVoipDomainSccpCallManagerDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_voip_domain_sccp_call_manager" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-voip-domain-sccp-call-manager", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("VoipDomainSccpCallManager object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementVoipDomainSccpCallManagerExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("VoipDomainSccpCallManager ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-voip-domain-sccp-call-manager", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementVoipDomainSccpCallManagerAttributes(voipDomainSccpCallManagerMap *map[string]interface{}, name string, endpointsDomain string, installedAt string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		voipDomainSccpCallManagerName := (*voipDomainSccpCallManagerMap)["name"].(string)
		if !strings.EqualFold(voipDomainSccpCallManagerName, name) {
			return fmt.Errorf("name is %s, expected %s", name, voipDomainSccpCallManagerName)
		}
		voipDomainSccpCallManagerEndpointsDomain := (*voipDomainSccpCallManagerMap)["endpoints-domain"].(map[string]interface{})["name"].(string)
		if !strings.EqualFold(voipDomainSccpCallManagerEndpointsDomain, endpointsDomain) {
			return fmt.Errorf("endpointsDomain is %s, expected %s", endpointsDomain, voipDomainSccpCallManagerEndpointsDomain)
		}
		voipDomainSccpCallManagerInstalledAt := (*voipDomainSccpCallManagerMap)["installed-at"].(map[string]interface{})["name"].(string)
		if !strings.EqualFold(voipDomainSccpCallManagerInstalledAt, installedAt) {
			return fmt.Errorf("installedAt is %s, expected %s", installedAt, voipDomainSccpCallManagerInstalledAt)
		}
		return nil
	}
}

func testAccManagementVoipDomainSccpCallManagerConfig(name string, endpointsDomain string, installedAt string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_group" "group1" {
  name = "%s"
}

resource "checkpoint_management_host" "host1" {
  name = "%s"
  ipv4_address = "192.0.2.8"
}

resource "checkpoint_management_voip_domain_sccp_call_manager" "test" {
        name = "%s"
        endpoints_domain = "${checkpoint_management_group.group1.name}"
        installed_at = "${checkpoint_management_host.host1.name}"
}
`, endpointsDomain, installedAt, name)
}
