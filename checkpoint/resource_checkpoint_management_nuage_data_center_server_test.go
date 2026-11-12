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

func TestAccCheckpointManagementNuageDataCenterServer_basic(t *testing.T) {

	var nuageDataCenterServerMap map[string]interface{}
	resourceName := "checkpoint_management_nuage_data_center_server.test"
	objName := "tfTestManagementNuageDataCenterServer_" + acctest.RandString(6)
	username := "USERNAME"
	password := "PASSWORD"
	hostname := "MY_HOSTNAME"
	organization := "MY_ORG"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementNuageDataCenterServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementNuageDataCenterServerConfig(objName, username, password, hostname, organization),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementNuageDataCenterServerExists(resourceName, &nuageDataCenterServerMap),
					testAccCheckCheckpointManagementNuageDataCenterServerAttributes(&nuageDataCenterServerMap, objName),
				),
			},
		},
	})
}

func testAccCheckpointManagementNuageDataCenterServerDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_nuage_data_center_server" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-data-center-server", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("NuageDataCenterServer object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementNuageDataCenterServerExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("NuageDataCenterServer ID is not set")
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

func testAccCheckCheckpointManagementNuageDataCenterServerAttributes(nuageDataCenterServerMap *map[string]interface{}, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		nuageDataCenterServerName := (*nuageDataCenterServerMap)["name"].(string)
		if !strings.EqualFold(nuageDataCenterServerName, name) {
			return fmt.Errorf("name is %s, expected %s", name, nuageDataCenterServerName)
		}
		return nil
	}
}

func testAccManagementNuageDataCenterServerConfig(name string, username string, password string, hostname string, organization string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_nuage_data_center_server" "test" {
    name = "%s"
	username = "%s"
	password = "%s"
	hostname = "%s"
	organization = "%s"
    unsafe_auto_accept = true
	ignore_warnings = true
}
`, name, username, password, hostname, organization)
}
