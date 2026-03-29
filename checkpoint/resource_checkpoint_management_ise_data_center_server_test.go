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

func TestAccCheckpointManagementIseDataCenterServer_basic(t *testing.T) {

	var iseDataCenterServerMap map[string]interface{}
	resourceName := "checkpoint_management_ise_data_center_server.test"
	objName := "tfTestManagementIseDataCenterServer_" + acctest.RandString(6)
	username := "USERNAME"
	password := "PASSWORD"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementIseDataCenterServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementIseDataCenterServerConfig(objName, username, password),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementIseDataCenterServerExists(resourceName, &iseDataCenterServerMap),
					testAccCheckCheckpointManagementIseDataCenterServerAttributes(&iseDataCenterServerMap, objName),
				),
			},
		},
	})
}

func testAccCheckpointManagementIseDataCenterServerDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_ise_data_center_server" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-data-center-server", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("IseDataCenterServer object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementIseDataCenterServerExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("IseDataCenterServer ID is not set")
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

func testAccCheckCheckpointManagementIseDataCenterServerAttributes(iseDataCenterServerMap *map[string]interface{}, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		iseDataCenterServerName := (*iseDataCenterServerMap)["name"].(string)
		if !strings.EqualFold(iseDataCenterServerName, name) {
			return fmt.Errorf("name is %s, expected %s", name, iseDataCenterServerName)
		}
		return nil
	}
}

func testAccManagementIseDataCenterServerConfig(name string, username string, password string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_ise_data_center_server" "test" {
    name = "%s"
	username = "%s"
	password = "%s"
	hostnames = ["host1", "host2"]
    unsafe_auto_accept = true
	ignore_warnings = true
}
`, name, username, password)
}
