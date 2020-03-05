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

func TestAccCheckpointManagementServiceDceRpc_basic(t *testing.T) {

    var serviceDceRpcMap map[string]interface{}
    resourceName := "checkpoint_management_service_dce_rpc.test"
    objName := "tfTestManagementServiceDceRpc_" + acctest.RandString(6)

    context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
        CheckDestroy: testAccCheckpointManagementServiceDceRpcDestroy,
        Steps: []resource.TestStep{
            {
                Config: testAccManagementServiceDceRpcConfig(objName, "97aeb460-9aea-11d5-bd16-0090272ccb30", false),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckCheckpointManagementServiceDceRpcExists(resourceName, &serviceDceRpcMap),
                    testAccCheckCheckpointManagementServiceDceRpcAttributes(&serviceDceRpcMap, objName, "97aeb460-9aea-11d5-bd16-0090272ccb30", false),
                ),
            },
        },
    })
}

func testAccCheckpointManagementServiceDceRpcDestroy(s *terraform.State) error {

    client := testAccProvider.Meta().(*checkpoint.ApiClient)
    for _, rs := range s.RootModule().Resources {
        if rs.Type != "checkpoint_management_service_dce_rpc" {
            continue
        }
        if rs.Primary.ID != "" {
            res, _ := client.ApiCall("show-service-dce-rpc", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
            if res.Success {
                return fmt.Errorf("ServiceDceRpc object (%s) still exists", rs.Primary.ID)
            }
        }
        return nil
    }
    return nil
}

func testAccCheckCheckpointManagementServiceDceRpcExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        rs, ok := s.RootModule().Resources[resourceTfName]
        if !ok {
            return fmt.Errorf("Resource not found: %s", resourceTfName)
        }

        if rs.Primary.ID == "" {
            return fmt.Errorf("ServiceDceRpc ID is not set")
        }

        client := testAccProvider.Meta().(*checkpoint.ApiClient)

        response, err := client.ApiCall("show-service-dce-rpc", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
        if !response.Success {
            return err
        }

        *res = response.GetData()

        return nil
    }
}

func testAccCheckCheckpointManagementServiceDceRpcAttributes(serviceDceRpcMap *map[string]interface{}, name string, interfaceUuid string, keepConnectionsOpenAfterPolicyInstallation bool) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        serviceDceRpcName := (*serviceDceRpcMap)["name"].(string)
        if !strings.EqualFold(serviceDceRpcName, name) {
            return fmt.Errorf("name is %s, expected %s", name, serviceDceRpcName)
        }
        serviceDceRpcInterfaceUuid := (*serviceDceRpcMap)["interface-uuid"].(string)
        if !strings.EqualFold(serviceDceRpcInterfaceUuid, interfaceUuid) {
            return fmt.Errorf("interfaceUuid is %s, expected %s", interfaceUuid, serviceDceRpcInterfaceUuid)
        }
        serviceDceRpcKeepConnectionsOpenAfterPolicyInstallation := (*serviceDceRpcMap)["keep-connections-open-after-policy-installation"].(bool)
        if serviceDceRpcKeepConnectionsOpenAfterPolicyInstallation != keepConnectionsOpenAfterPolicyInstallation {
            return fmt.Errorf("keepConnectionsOpenAfterPolicyInstallation is %t, expected %t", keepConnectionsOpenAfterPolicyInstallation, serviceDceRpcKeepConnectionsOpenAfterPolicyInstallation)
        }
        return nil
    }
}

func testAccManagementServiceDceRpcConfig(name string, interfaceUuid string, keepConnectionsOpenAfterPolicyInstallation bool) string {
    return fmt.Sprintf(`
resource "checkpoint_management_service_dce_rpc" "test" {
        name = "%s"
        interface_uuid = "%s"
        keep_connections_open_after_policy_installation = %t
}
`, name, interfaceUuid, keepConnectionsOpenAfterPolicyInstallation)
}

