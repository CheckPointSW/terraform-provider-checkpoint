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

func TestAccCheckpointManagementServiceRpc_basic(t *testing.T) {

    var serviceRpcMap map[string]interface{}
    resourceName := "checkpoint_management_service_rpc.test"
    objName := "tfTestManagementServiceRpc_" + acctest.RandString(6)

    context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
        CheckDestroy: testAccCheckpointManagementServiceRpcDestroy,
        Steps: []resource.TestStep{
            {
                Config: testAccManagementServiceRpcConfig(objName, 5669, false),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckCheckpointManagementServiceRpcExists(resourceName, &serviceRpcMap),
                    testAccCheckCheckpointManagementServiceRpcAttributes(&serviceRpcMap, objName, 5669, false),
                ),
            },
        },
    })
}

func testAccCheckpointManagementServiceRpcDestroy(s *terraform.State) error {

    client := testAccProvider.Meta().(*checkpoint.ApiClient)
    for _, rs := range s.RootModule().Resources {
        if rs.Type != "checkpoint_management_service_rpc" {
            continue
        }
        if rs.Primary.ID != "" {
            res, _ := client.ApiCall("show-service-rpc", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
            if res.Success {
                return fmt.Errorf("ServiceRpc object (%s) still exists", rs.Primary.ID)
            }
        }
        return nil
    }
    return nil
}

func testAccCheckCheckpointManagementServiceRpcExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        rs, ok := s.RootModule().Resources[resourceTfName]
        if !ok {
            return fmt.Errorf("Resource not found: %s", resourceTfName)
        }

        if rs.Primary.ID == "" {
            return fmt.Errorf("ServiceRpc ID is not set")
        }

        client := testAccProvider.Meta().(*checkpoint.ApiClient)

        response, err := client.ApiCall("show-service-rpc", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
        if !response.Success {
            return err
        }

        *res = response.GetData()

        return nil
    }
}

func testAccCheckCheckpointManagementServiceRpcAttributes(serviceRpcMap *map[string]interface{}, name string, programNumber int, keepConnectionsOpenAfterPolicyInstallation bool) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        serviceRpcName := (*serviceRpcMap)["name"].(string)
        if !strings.EqualFold(serviceRpcName, name) {
            return fmt.Errorf("name is %s, expected %s", name, serviceRpcName)
        }
        serviceRpcProgramNumber := int((*serviceRpcMap)["program-number"].(float64))
        if serviceRpcProgramNumber != programNumber {
            return fmt.Errorf("programNumber is %d, expected %d", programNumber, serviceRpcProgramNumber)
        }
        serviceRpcKeepConnectionsOpenAfterPolicyInstallation := (*serviceRpcMap)["keep-connections-open-after-policy-installation"].(bool)
        if serviceRpcKeepConnectionsOpenAfterPolicyInstallation != keepConnectionsOpenAfterPolicyInstallation {
            return fmt.Errorf("keepConnectionsOpenAfterPolicyInstallation is %t, expected %t", keepConnectionsOpenAfterPolicyInstallation, serviceRpcKeepConnectionsOpenAfterPolicyInstallation)
        }
        return nil
    }
}

func testAccManagementServiceRpcConfig(name string, programNumber int, keepConnectionsOpenAfterPolicyInstallation bool) string {
    return fmt.Sprintf(`
resource "checkpoint_management_service_rpc" "test" {
        name = "%s"
        program_number = %d
        keep_connections_open_after_policy_installation = %t
}
`, name, programNumber, keepConnectionsOpenAfterPolicyInstallation)
}

