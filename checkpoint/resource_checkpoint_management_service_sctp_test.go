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

func TestAccCheckpointManagementServiceSctp_basic(t *testing.T) {

    var serviceSctpMap map[string]interface{}
    resourceName := "checkpoint_management_service_sctp.test"
    objName := "tfTestManagementServiceSctp_" + acctest.RandString(6)

    context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
        CheckDestroy: testAccCheckpointManagementServiceSctpDestroy,
        Steps: []resource.TestStep{
            {
                Config: testAccManagementServiceSctpConfig(objName, "5669", false, 100, true, true),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckCheckpointManagementServiceSctpExists(resourceName, &serviceSctpMap),
                    testAccCheckCheckpointManagementServiceSctpAttributes(&serviceSctpMap, objName, "5669", false, 100, true, true),
                ),
            },
        },
    })
}

func testAccCheckpointManagementServiceSctpDestroy(s *terraform.State) error {

    client := testAccProvider.Meta().(*checkpoint.ApiClient)
    for _, rs := range s.RootModule().Resources {
        if rs.Type != "checkpoint_management_service_sctp" {
            continue
        }
        if rs.Primary.ID != "" {
            res, _ := client.ApiCall("show-service-sctp", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
            if res.Success {
                return fmt.Errorf("ServiceSctp object (%s) still exists", rs.Primary.ID)
            }
        }
        return nil
    }
    return nil
}

func testAccCheckCheckpointManagementServiceSctpExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        rs, ok := s.RootModule().Resources[resourceTfName]
        if !ok {
            return fmt.Errorf("Resource not found: %s", resourceTfName)
        }

        if rs.Primary.ID == "" {
            return fmt.Errorf("ServiceSctp ID is not set")
        }

        client := testAccProvider.Meta().(*checkpoint.ApiClient)

        response, err := client.ApiCall("show-service-sctp", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
        if !response.Success {
            return err
        }

        *res = response.GetData()

        return nil
    }
}

func testAccCheckCheckpointManagementServiceSctpAttributes(serviceSctpMap *map[string]interface{}, name string, port string, keepConnectionsOpenAfterPolicyInstallation bool, sessionTimeout int, matchForAny bool, syncConnectionsOnCluster bool) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        serviceSctpName := (*serviceSctpMap)["name"].(string)
        if !strings.EqualFold(serviceSctpName, name) {
            return fmt.Errorf("name is %s, expected %s", name, serviceSctpName)
        }
        serviceSctpPort := (*serviceSctpMap)["port"].(string)
        if !strings.EqualFold(serviceSctpPort, port) {
            return fmt.Errorf("port is %s, expected %s", port, serviceSctpPort)
        }
        serviceSctpKeepConnectionsOpenAfterPolicyInstallation := (*serviceSctpMap)["keep-connections-open-after-policy-installation"].(bool)
        if serviceSctpKeepConnectionsOpenAfterPolicyInstallation != keepConnectionsOpenAfterPolicyInstallation {
            return fmt.Errorf("keepConnectionsOpenAfterPolicyInstallation is %t, expected %t", keepConnectionsOpenAfterPolicyInstallation, serviceSctpKeepConnectionsOpenAfterPolicyInstallation)
        }
        serviceSctpMatchForAny := (*serviceSctpMap)["match-for-any"].(bool)
        if serviceSctpMatchForAny != matchForAny {
            return fmt.Errorf("matchForAny is %t, expected %t", matchForAny, serviceSctpMatchForAny)
        }
        serviceSctpSyncConnectionsOnCluster := (*serviceSctpMap)["sync-connections-on-cluster"].(bool)
        if serviceSctpSyncConnectionsOnCluster != syncConnectionsOnCluster {
            return fmt.Errorf("syncConnectionsOnCluster is %t, expected %t", syncConnectionsOnCluster, serviceSctpSyncConnectionsOnCluster)
        }
        return nil
    }
}

func testAccManagementServiceSctpConfig(name string, port string, keepConnectionsOpenAfterPolicyInstallation bool, sessionTimeout int, matchForAny bool, syncConnectionsOnCluster bool) string {
    return fmt.Sprintf(`
resource "checkpoint_management_service_sctp" "test" {
        name = "%s"
        port = "%s"
        keep_connections_open_after_policy_installation = %t
        session_timeout = %d
        match_for_any = %t
        sync_connections_on_cluster = %t
        aggressive_aging = {
            use_default_timeout = true
            enable = true
            default_timeout = 600
            timeout = 600
        }
}
`, name, port, keepConnectionsOpenAfterPolicyInstallation, sessionTimeout, matchForAny, syncConnectionsOnCluster)
}

