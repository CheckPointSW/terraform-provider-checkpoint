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

func TestAccCheckpointManagementGenericDataCenterServer_basic(t *testing.T) {

	var genericDataCenterServerMap map[string]interface{}
	resourceName := "checkpoint_management_generic_data_center_server.test"
	objName := "tfTestManagementGenericDataCenterServer_" + acctest.RandString(6)
	url := "MY_URL"
	interval := "60"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementGenericDataCenterServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementGenericDataCenterServerConfig(objName, url, interval),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementGenericDataCenterServerExists(resourceName, &genericDataCenterServerMap),
					testAccCheckCheckpointManagementGenericDataCenterServerAttributes(&genericDataCenterServerMap, objName),
				),
			},
		},
	})
}

func testAccCheckpointManagementGenericDataCenterServerDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_generic_data_center_server" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-data-center-server", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("GenericDataCenterServer object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementGenericDataCenterServerExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("GenericDataCenterServer ID is not set")
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

func testAccCheckCheckpointManagementGenericDataCenterServerAttributes(genericDataCenterServerMap *map[string]interface{}, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		genericDataCenterServerName := (*genericDataCenterServerMap)["name"].(string)
		if !strings.EqualFold(genericDataCenterServerName, name) {
			return fmt.Errorf("name is %s, expected %s", name, genericDataCenterServerName)
		}
		return nil
	}
}

func testAccManagementGenericDataCenterServerConfig(name string, url string, interval string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_generic_data_center_server" "test" {
        name = "%s"
		url = "%s"
		interval = "%s"
		ignore_warnings = true
}
`, name, url, interval)
}
