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

func TestAccCheckpointManagementServiceCompoundTcp_basic(t *testing.T) {

	var serviceCompoundTcpMap map[string]interface{}
	resourceName := "checkpoint_management_service_compound_tcp.test"
	objName := "tfTestManagementServiceCompoundTcp_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementServiceCompoundTcpDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementServiceCompoundTcpConfig(objName, "pointcast", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementServiceCompoundTcpExists(resourceName, &serviceCompoundTcpMap),
					testAccCheckCheckpointManagementServiceCompoundTcpAttributes(&serviceCompoundTcpMap, objName, "pointcast", true),
				),
			},
		},
	})
}

func testAccCheckpointManagementServiceCompoundTcpDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_service_compound_tcp" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-service-compound-tcp", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("ServiceCompoundTcp object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementServiceCompoundTcpExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ServiceCompoundTcp ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-service-compound-tcp", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementServiceCompoundTcpAttributes(serviceCompoundTcpMap *map[string]interface{}, name string, compoundService string, keepConnectionsOpenAfterPolicyInstallation bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		serviceCompoundTcpName := (*serviceCompoundTcpMap)["name"].(string)
		if !strings.EqualFold(serviceCompoundTcpName, name) {
			return fmt.Errorf("name is %s, expected %s", name, serviceCompoundTcpName)
		}
		serviceCompoundTcpCompoundService := (*serviceCompoundTcpMap)["compound-service"].(string)
		if !strings.EqualFold(serviceCompoundTcpCompoundService, compoundService) {
			return fmt.Errorf("compoundService is %s, expected %s", compoundService, serviceCompoundTcpCompoundService)
		}
		serviceCompoundTcpKeepConnectionsOpenAfterPolicyInstallation := (*serviceCompoundTcpMap)["keep-connections-open-after-policy-installation"].(bool)
		if serviceCompoundTcpKeepConnectionsOpenAfterPolicyInstallation != keepConnectionsOpenAfterPolicyInstallation {
			return fmt.Errorf("keepConnectionsOpenAfterPolicyInstallation is %t, expected %t", keepConnectionsOpenAfterPolicyInstallation, serviceCompoundTcpKeepConnectionsOpenAfterPolicyInstallation)
		}
		return nil
	}
}

func testAccManagementServiceCompoundTcpConfig(name string, compoundService string, keepConnectionsOpenAfterPolicyInstallation bool) string {
	return fmt.Sprintf(`
resource "checkpoint_management_service_compound_tcp" "test" {
        name = "%s"
        compound_service = "%s"
        keep_connections_open_after_policy_installation = %t
}
`, name, compoundService, keepConnectionsOpenAfterPolicyInstallation)
}
