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

func TestAccCheckpointManagementIllumioDataCenterServer_basic(t *testing.T) {

	var illumioDataCenterServerMap map[string]interface{}
	resourceName := "checkpoint_management_illumio_data_center_server.test"
	objName := "tfTestManagementIllumioDataCenterServer_" + acctest.RandString(6)
	hostname := "hostname.illum.io"
	orgId := 1234567
	authUsername := "api_6e8d5249c27d64185"
	secret := "e7f5b0e8-4f5e-a17d-8f9e-0d2e5bdsdfe1"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementIllumioDataCenterServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementIllumioDataCenterServerConfig(objName, hostname, orgId, authUsername, secret),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementIllumioDataCenterServerExists(resourceName, &illumioDataCenterServerMap),
					testAccCheckCheckpointManagementIllumioDataCenterServerAttributes(&illumioDataCenterServerMap, objName),
				),
			},
		},
	})
}

func testAccCheckpointManagementIllumioDataCenterServerDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_illumio_data_center_server" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-data-center-server", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("IllumioDataCenterServer object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementIllumioDataCenterServerExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("IllumioDataCenterServer ID is not set")
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

func testAccCheckCheckpointManagementIllumioDataCenterServerAttributes(illumioDataCenterServerMap *map[string]interface{}, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		illumioDataCenterServerName := (*illumioDataCenterServerMap)["name"].(string)
		if !strings.EqualFold(illumioDataCenterServerName, name) {
			return fmt.Errorf("name is %s, expected %s", name, illumioDataCenterServerName)
		}
		return nil
	}
}

func testAccManagementIllumioDataCenterServerConfig(name string, hostname string, orgId int, authUsername string, secret string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_illumio_data_center_server" "test" {
    name = "%s"
	hostname = "%s"
	org_id = %d
	auth_username = "%s"
	secret = "%s"
	ignore_warnings = true
}
`, name, hostname, orgId, authUsername, secret)
}
