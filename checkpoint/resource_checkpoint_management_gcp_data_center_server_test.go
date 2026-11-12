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

func TestAccCheckpointManagementGcpDataCenterServer_basic(t *testing.T) {

	var gcpDataCenterServerMap map[string]interface{}
	resourceName := "checkpoint_management_gcp_data_center_server.test"
	objName := "tfTestManagementGcpDataCenterServer_" + acctest.RandString(6)
	authenticationMethod := "key-authentication"
	privateKey := "MYKEY.json"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementGcpDataCenterServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementGcpDataCenterServerConfig(objName, authenticationMethod, privateKey),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementGcpDataCenterServerExists(resourceName, &gcpDataCenterServerMap),
					testAccCheckCheckpointManagementGcpDataCenterServerAttributes(&gcpDataCenterServerMap, objName),
				),
			},
		},
	})
}

func testAccCheckpointManagementGcpDataCenterServerDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_gcp_data_center_server" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-data-center-server", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("GcpDataCenterServer object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementGcpDataCenterServerExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("GcpDataCenterServer ID is not set")
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

func testAccCheckCheckpointManagementGcpDataCenterServerAttributes(gcpDataCenterServerMap *map[string]interface{}, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		gcpDataCenterServerName := (*gcpDataCenterServerMap)["name"].(string)
		if !strings.EqualFold(gcpDataCenterServerName, name) {
			return fmt.Errorf("name is %s, expected %s", name, gcpDataCenterServerName)
		}
		return nil
	}
}

func testAccManagementGcpDataCenterServerConfig(name string, authenticationMethod string, privateKey string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_gcp_data_center_server" "test" {
	name = "%s"
	authentication_method = "%s"
	private_key = "%s"
	ignore_warnings = true
}
`, name, authenticationMethod, privateKey)
}
