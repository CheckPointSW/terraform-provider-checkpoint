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

func TestAccCheckpointManagementLogicalServer_basic(t *testing.T) {

	var logicalServerMap map[string]interface{}
	resourceName := "checkpoint_management_logical_server.test"
	objName := "tfTestManagementLogicalServer_" + acctest.RandString(6)
	groupName := "tfTestManagementGroup_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementLogicalServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementLogicalServerConfig(objName, groupName, "other", true, "by_server", "domain", "1.1.1.19"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementLogicalServerExists(resourceName, &logicalServerMap),
					testAccCheckCheckpointManagementLogicalServerAttributes(&logicalServerMap, objName, groupName, "other", true, "by_server", "domain", "1.1.1.19"),
				),
			},
		},
	})
}

func testAccCheckpointManagementLogicalServerDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_logical_server" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-logical-server", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("LogicalServer object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementLogicalServerExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("LogicalServer ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-logical-server", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementLogicalServerAttributes(logicalServerMap *map[string]interface{}, name string, serverGroup string, serverType string, persistenceMode bool, persistencyType string, balanceMethod string, ipv4Address string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		logicalServerName := (*logicalServerMap)["name"].(string)
		if !strings.EqualFold(logicalServerName, name) {
			return fmt.Errorf("name is %s, expected %s", name, logicalServerName)
		}
		logicalServerServerGroup := (*logicalServerMap)["server-group"].(map[string]interface{})["name"].(string)
		if !strings.EqualFold(logicalServerServerGroup, serverGroup) {
			return fmt.Errorf("serverGroup is %s, expected %s", serverGroup, logicalServerServerGroup)
		}
		logicalServerServerType := (*logicalServerMap)["server-type"].(string)
		if !strings.EqualFold(logicalServerServerType, serverType) {
			return fmt.Errorf("serverType is %s, expected %s", serverType, logicalServerServerType)
		}
		logicalServerPersistenceMode := (*logicalServerMap)["persistence-mode"].(bool)
		if logicalServerPersistenceMode != persistenceMode {
			return fmt.Errorf("persistenceMode is %t, expected %t", persistenceMode, logicalServerPersistenceMode)
		}
		logicalServerPersistencyType := (*logicalServerMap)["persistency-type"].(string)
		if !strings.EqualFold(logicalServerPersistencyType, persistencyType) {
			return fmt.Errorf("persistencyType is %s, expected %s", persistencyType, logicalServerPersistencyType)
		}
		logicalServerBalanceMethod := (*logicalServerMap)["balance-method"].(string)
		if !strings.EqualFold(logicalServerBalanceMethod, balanceMethod) {
			return fmt.Errorf("balanceMethod is %s, expected %s", balanceMethod, logicalServerBalanceMethod)
		}
		logicalServerIpv4Address := (*logicalServerMap)["ipv4-address"].(string)
		if !strings.EqualFold(logicalServerIpv4Address, ipv4Address) {
			return fmt.Errorf("ipv4Address is %s, expected %s", ipv4Address, logicalServerIpv4Address)
		}
		return nil
	}
}

func testAccManagementLogicalServerConfig(name string, serverGroup string, serverType string, persistenceMode bool, persistencyType string, balanceMethod string, ipv4Address string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_group" "test" {
    name = "%s"
}

resource "checkpoint_management_logical_server" "test" {
        name = "%s"
        server_group = "${checkpoint_management_group.test.name}"
        server_type = "%s"
        persistence_mode = %t
        persistency_type = "%s"
        balance_method = "%s"
        ipv4_address = "%s"
}
`, serverGroup, name, serverType, persistenceMode, persistencyType, balanceMethod, ipv4Address)
}
