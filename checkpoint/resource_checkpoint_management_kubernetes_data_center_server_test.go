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

func TestAccCheckpointManagementKubernetesDataCenterServer_basic(t *testing.T) {

	var kubernetesDataCenterServerMap map[string]interface{}
	resourceName := "checkpoint_management_kubernetes_data_center_server.test"
	objName := "tfTestManagementKubernetesDataCenterServer_" + acctest.RandString(6)
	hostname := "MY_HOSTNAME"
	token_file := "MY_TOKEN"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementKubernetesDataCenterServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementKubernetesDataCenterServerConfig(objName, hostname, token_file),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementKubernetesDataCenterServerExists(resourceName, &kubernetesDataCenterServerMap),
					testAccCheckCheckpointManagementKubernetesDataCenterServerAttributes(&kubernetesDataCenterServerMap, objName),
				),
			},
		},
	})
}

func testAccCheckpointManagementKubernetesDataCenterServerDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_kubernetes_data_center_server" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-data-center-server", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("KubernetesDataCenterServer object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementKubernetesDataCenterServerExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("KubernetesDataCenterServer ID is not set")
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

func testAccCheckCheckpointManagementKubernetesDataCenterServerAttributes(kubernetesDataCenterServerMap *map[string]interface{}, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		kubernetesDataCenterServerName := (*kubernetesDataCenterServerMap)["name"].(string)
		if !strings.EqualFold(kubernetesDataCenterServerName, name) {
			return fmt.Errorf("name is %s, expected %s", name, kubernetesDataCenterServerName)
		}
		return nil
	}
}

func testAccManagementKubernetesDataCenterServerConfig(name string, hostname string, token_file string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_kubernetes_data_center_server" "test" {
    name = "%s"
	hostname = "%s"
	token_file = "%s"
    unsafe_auto_accept = true
	ignore_warnings = true
}
`, name, hostname, token_file)
}
