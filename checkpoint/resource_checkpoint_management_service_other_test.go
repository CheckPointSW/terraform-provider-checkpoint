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

func TestAccCheckpointManagementServiceOther_basic(t *testing.T) {

    var serviceOtherMap map[string]interface{}
    resourceName := "checkpoint_management_service_other.test"
    objName := "tfTestManagementServiceOther_" + acctest.RandString(6)

    context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
        CheckDestroy: testAccCheckpointManagementServiceOtherDestroy,
        Steps: []resource.TestStep{
            {
                Config: testAccManagementServiceOtherConfig(objName, false, 100, true, true, 51),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckCheckpointManagementServiceOtherExists(resourceName, &serviceOtherMap),
                    testAccCheckCheckpointManagementServiceOtherAttributes(&serviceOtherMap, objName, false, 100, true, true, 51),
                ),
            },
        },
    })
}

func testAccCheckpointManagementServiceOtherDestroy(s *terraform.State) error {

    client := testAccProvider.Meta().(*checkpoint.ApiClient)
    for _, rs := range s.RootModule().Resources {
        if rs.Type != "checkpoint_management_service_other" {
            continue
        }
        if rs.Primary.ID != "" {
            res, _ := client.ApiCall("show-service-other", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
            if res.Success {
                return fmt.Errorf("ServiceOther object (%s) still exists", rs.Primary.ID)
            }
        }
        return nil
    }
    return nil
}

func testAccCheckCheckpointManagementServiceOtherExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        rs, ok := s.RootModule().Resources[resourceTfName]
        if !ok {
            return fmt.Errorf("Resource not found: %s", resourceTfName)
        }

        if rs.Primary.ID == "" {
            return fmt.Errorf("ServiceOther ID is not set")
        }

        client := testAccProvider.Meta().(*checkpoint.ApiClient)

        response, err := client.ApiCall("show-service-other", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
        if !response.Success {
            return err
        }

        *res = response.GetData()

        return nil
    }
}

func testAccCheckCheckpointManagementServiceOtherAttributes(serviceOtherMap *map[string]interface{}, name string, keepConnectionsOpenAfterPolicyInstallation bool, sessionTimeout int, matchForAny bool, syncConnectionsOnCluster bool, ipProtocol int) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        serviceOtherName := (*serviceOtherMap)["name"].(string)
        if !strings.EqualFold(serviceOtherName, name) {
            return fmt.Errorf("name is %s, expected %s", name, serviceOtherName)
        }
        serviceOtherKeepConnectionsOpenAfterPolicyInstallation := (*serviceOtherMap)["keep-connections-open-after-policy-installation"].(bool)
        if serviceOtherKeepConnectionsOpenAfterPolicyInstallation != keepConnectionsOpenAfterPolicyInstallation {
            return fmt.Errorf("keepConnectionsOpenAfterPolicyInstallation is %t, expected %t", keepConnectionsOpenAfterPolicyInstallation, serviceOtherKeepConnectionsOpenAfterPolicyInstallation)
        }
        serviceOtherMatchForAny := (*serviceOtherMap)["match-for-any"].(bool)
        if serviceOtherMatchForAny != matchForAny {
            return fmt.Errorf("matchForAny is %t, expected %t", matchForAny, serviceOtherMatchForAny)
        }
        serviceOtherSyncConnectionsOnCluster := (*serviceOtherMap)["sync-connections-on-cluster"].(bool)
        if serviceOtherSyncConnectionsOnCluster != syncConnectionsOnCluster {
            return fmt.Errorf("syncConnectionsOnCluster is %t, expected %t", syncConnectionsOnCluster, serviceOtherSyncConnectionsOnCluster)
        }
        serviceOtherIpProtocol := int((*serviceOtherMap)["ip-protocol"].(float64))
        if serviceOtherIpProtocol != ipProtocol {
            return fmt.Errorf("ipProtocol is %d, expected %d", ipProtocol, serviceOtherIpProtocol)
        }
        return nil
    }
}

func testAccManagementServiceOtherConfig(name string, keepConnectionsOpenAfterPolicyInstallation bool, sessionTimeout int, matchForAny bool, syncConnectionsOnCluster bool, ipProtocol int) string {
    return fmt.Sprintf(`
resource "checkpoint_management_service_other" "test" {
        name = "%s"
        keep_connections_open_after_policy_installation = %t
        session_timeout = %d
        match_for_any = %t
        sync_connections_on_cluster = %t
        ip_protocol = %d
        aggressive_aging = {
            use_default_timeout = true
            enable = true
            default_timeout = 600
            timeout = 600
        }
}
`, name, keepConnectionsOpenAfterPolicyInstallation, sessionTimeout, matchForAny, syncConnectionsOnCluster, ipProtocol)
}

