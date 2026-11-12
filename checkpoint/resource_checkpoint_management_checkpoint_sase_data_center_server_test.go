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

func TestAccCheckpointManagementCheckpointSaseDataCenterServer_basic(t *testing.T) {

	var checkpointSaseDataCenterServerMap map[string]interface{}
	resourceName := "checkpoint_management_checkpoint_sase_data_center_server.test"
	objName := "tfTestManagementCheckpointSaseDataCenterServer_" + acctest.RandString(6)
	connectTo := "connected-portal"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementCheckpointSaseDataCenterServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementCheckpointSaseDataCenterServerConfig(objName, connectTo),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementCheckpointSaseDataCenterServerExists(resourceName, &checkpointSaseDataCenterServerMap),
					testAccCheckCheckpointManagementCheckpointSaseDataCenterServerAttributes(&checkpointSaseDataCenterServerMap, objName),
				),
			},
		},
	})
}

func testAccCheckpointManagementCheckpointSaseDataCenterServerDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_checkpoint_sase_data_center_server" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-data-center-server", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("CheckpointSaseDataCenterServer object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementCheckpointSaseDataCenterServerExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("CheckpointSaseDataCenterServer ID is not set")
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

func testAccCheckCheckpointManagementCheckpointSaseDataCenterServerAttributes(checkpointSaseDataCenterServerMap *map[string]interface{}, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		checkpointSaseDataCenterServerName := (*checkpointSaseDataCenterServerMap)["name"].(string)
		if !strings.EqualFold(checkpointSaseDataCenterServerName, name) {
			return fmt.Errorf("name is %s, expected %s", name, checkpointSaseDataCenterServerName)
		}
		return nil
	}
}

func testAccManagementCheckpointSaseDataCenterServerConfig(name string, connectTo string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_checkpoint_sase_data_center_server" "test" {
    name = "%s"
	connect_to = "%s"
	ignore_warnings = true
}
`, name, connectTo)
}
