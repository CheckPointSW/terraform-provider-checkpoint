package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform/helper/resource"
    "github.com/hashicorp/terraform/terraform"
    "os"
    "strings"
    "testing"
    "github.com/hashicorp/terraform/helper/acctest"
)

func TestAccCheckpointManagementVpnCommunityMeshed_basic(t *testing.T) {

    var vpnCommunityMeshedMap map[string]interface{}
    resourceName := "checkpoint_management_vpn_community_meshed.test"
    objName := "tfTestManagementVpnCommunityMeshed_" + acctest.RandString(6)

    context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
        CheckDestroy: testAccCheckpointManagementVpnCommunityMeshedDestroy,
        Steps: []resource.TestStep{
            {
                Config: testAccManagementVpnCommunityMeshedConfig(objName, "prefer ikev2 but support ikev1", "custom"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckCheckpointManagementVpnCommunityMeshedExists(resourceName, &vpnCommunityMeshedMap),
                    testAccCheckCheckpointManagementVpnCommunityMeshedAttributes(&vpnCommunityMeshedMap, objName, "prefer ikev2 but support ikev1", "custom"),
                ),
            },
        },
    })
}

func testAccCheckpointManagementVpnCommunityMeshedDestroy(s *terraform.State) error {

    client := testAccProvider.Meta().(*checkpoint.ApiClient)
    for _, rs := range s.RootModule().Resources {
        if rs.Type != "checkpoint_management_vpn_community_meshed" {
            continue
        }
        if rs.Primary.ID != "" {
            res, _ := client.ApiCall("show-vpn-community-meshed", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
            if res.Success {
                return fmt.Errorf("VpnCommunityMeshed object (%s) still exists", rs.Primary.ID)
            }
        }
        return nil
    }
    return nil
}

func testAccCheckCheckpointManagementVpnCommunityMeshedExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        rs, ok := s.RootModule().Resources[resourceTfName]
        if !ok {
            return fmt.Errorf("Resource not found: %s", resourceTfName)
        }

        if rs.Primary.ID == "" {
            return fmt.Errorf("VpnCommunityMeshed ID is not set")
        }

        client := testAccProvider.Meta().(*checkpoint.ApiClient)

        response, err := client.ApiCall("show-vpn-community-meshed", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
        if !response.Success {
            return err
        }

        *res = response.GetData()

        return nil
    }
}

func testAccCheckCheckpointManagementVpnCommunityMeshedAttributes(vpnCommunityMeshedMap *map[string]interface{}, name string, encryptionMethod string, encryptionSuite string) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        vpnCommunityMeshedName := (*vpnCommunityMeshedMap)["name"].(string)
        if !strings.EqualFold(vpnCommunityMeshedName, name) {
            return fmt.Errorf("name is %s, expected %s", name, vpnCommunityMeshedName)
        }
        vpnCommunityMeshedEncryptionMethod := (*vpnCommunityMeshedMap)["encryption-method"].(string)
        if !strings.EqualFold(vpnCommunityMeshedEncryptionMethod, encryptionMethod) {
            return fmt.Errorf("encryptionMethod is %s, expected %s", encryptionMethod, vpnCommunityMeshedEncryptionMethod)
        }
        vpnCommunityMeshedEncryptionSuite := (*vpnCommunityMeshedMap)["encryption-suite"].(string)
        if !strings.EqualFold(vpnCommunityMeshedEncryptionSuite, encryptionSuite) {
            return fmt.Errorf("encryptionSuite is %s, expected %s", encryptionSuite, vpnCommunityMeshedEncryptionSuite)
        }
        return nil
    }
}

func testAccManagementVpnCommunityMeshedConfig(name string, encryptionMethod string, encryptionSuite string) string {
    return fmt.Sprintf(`
resource "checkpoint_management_vpn_community_meshed" "test" {
        name = "%s"
        encryption_method = "%s"
        encryption_suite = "%s"
}
`, name, encryptionMethod, encryptionSuite)
}

