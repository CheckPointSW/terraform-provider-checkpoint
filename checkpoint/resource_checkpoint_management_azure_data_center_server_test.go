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

func TestAccCheckpointManagementAzureDataCenterServer_basic(t *testing.T) {

	var azureDataCenterServerMap map[string]interface{}
	resourceName := "checkpoint_management_azure_data_center_server.test"
	objName := "tfTestManagementAzureDataCenterServer_" + acctest.RandString(6)
	authenticationMethod := "user-authentication"
	username := "MY-KEY-ID"
	password := "MY-SECRET-KEY"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementAzureDataCenterServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementAzureDataCenterServerConfig(objName, username, password, authenticationMethod),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementAzureDataCenterServerExists(resourceName, &azureDataCenterServerMap),
					testAccCheckCheckpointManagementAzureDataCenterServerAttributes(&azureDataCenterServerMap, objName),
				),
			},
		},
	})
}

func testAccCheckpointManagementAzureDataCenterServerDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_azure_data_center_server" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-data-center-server", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("AzureDataCenterServer object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementAzureDataCenterServerExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("AzureDataCenterServer ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-data-center-server", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementAzureDataCenterServerAttributes(azureDataCenterServerMap *map[string]interface{}, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		azureDataCenterServerName := (*azureDataCenterServerMap)["name"].(string)
		if !strings.EqualFold(azureDataCenterServerName, name) {
			return fmt.Errorf("name is %s, expected %s", name, azureDataCenterServerName)
		}
		return nil
	}
}

func testAccManagementAzureDataCenterServerConfig(name string, username string, password string, authenticationMethod string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_azure_data_center_server" "test" {
	name = "%s"
	username         = "%s"
	password     = "%s"
	authentication_method = "%s"
	ignore_warnings = true
}
`, name, username, password, authenticationMethod)
}
