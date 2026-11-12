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

func TestAccCheckpointManagementClarotyCtdDataCenterServer_basic(t *testing.T) {

	var clarotyCtdDataCenterServerMap map[string]interface{}
	resourceName := "checkpoint_management_claroty_ctd_data_center_server.test"
	objName := "tfTestManagementClarotyCtdDataCenterServer_" + acctest.RandString(6)
	hostname := "1.2.3.4"
	username := "example-username"
	password := "example-password"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementClarotyCtdDataCenterServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementClarotyCtdDataCenterServerConfig(objName, hostname, username, password),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementClarotyCtdDataCenterServerExists(resourceName, &clarotyCtdDataCenterServerMap),
					testAccCheckCheckpointManagementClarotyCtdDataCenterServerAttributes(&clarotyCtdDataCenterServerMap, objName),
				),
			},
		},
	})
}

func testAccCheckpointManagementClarotyCtdDataCenterServerDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_claroty_ctd_data_center_server" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-data-center-server", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("ClarotyCtdDataCenterServer object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementClarotyCtdDataCenterServerExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ClarotyCtdDataCenterServer ID is not set")
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

func testAccCheckCheckpointManagementClarotyCtdDataCenterServerAttributes(clarotyCtdDataCenterServerMap *map[string]interface{}, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		clarotyCtdDataCenterServerName := (*clarotyCtdDataCenterServerMap)["name"].(string)
		if !strings.EqualFold(clarotyCtdDataCenterServerName, name) {
			return fmt.Errorf("name is %s, expected %s", name, clarotyCtdDataCenterServerName)
		}
		return nil
	}
}

func testAccManagementClarotyCtdDataCenterServerConfig(name string, hostname string, username string, password string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_claroty_ctd_data_center_server" "test" {
    name = "%s"
	hostname = "%s"
	username = "%s"
	password = "%s"
	ignore_warnings = true
}
`, name, hostname, username, password)
}
