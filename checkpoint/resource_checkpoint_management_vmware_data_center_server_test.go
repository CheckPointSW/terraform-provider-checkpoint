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

func TestAccCheckpointManagementVMwareDataCenterServer_basic(t *testing.T) {

	var vmwareDataCenterServerMap map[string]interface{}
	resourceName := "checkpoint_management_vmware_data_center_server.test"
	objName := "tfTestManagementVMwareDataCenterServer_" + acctest.RandString(6)
	vmType := "vcenter"
	username := "USERNAME"
	password := "PASSWORD"
	hostname := "HOSTNAME"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementVMwareDataCenterServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementVMwareDataCenterServerConfig(objName, vmType, username, password, hostname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementVMwareDataCenterServerExists(resourceName, &vmwareDataCenterServerMap),
					testAccCheckCheckpointManagementVMwareDataCenterServerAttributes(&vmwareDataCenterServerMap, objName),
				),
			},
		},
	})
}

func testAccCheckpointManagementVMwareDataCenterServerDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_vmware_data_center_server" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-data-center-server", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("VMwareDataCenterServer object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementVMwareDataCenterServerExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("VMwareDataCenterServer ID is not set")
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

func testAccCheckCheckpointManagementVMwareDataCenterServerAttributes(vmwareDataCenterServerMap *map[string]interface{}, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		vmwareDataCenterServerName := (*vmwareDataCenterServerMap)["name"].(string)
		if !strings.EqualFold(vmwareDataCenterServerName, name) {
			return fmt.Errorf("name is %s, expected %s", name, vmwareDataCenterServerName)
		}
		return nil
	}
}

func testAccManagementVMwareDataCenterServerConfig(name string, vmType string, username string, password string, hostname string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_vmware_data_center_server" "test" {
    name = "%s"
	type = "%s"
	username = "%s"
	password = "%s"
	hostname = "%s"
    unsafe_auto_accept = true
	ignore_warnings = true
}
`, name, vmType, username, password, hostname)
}
