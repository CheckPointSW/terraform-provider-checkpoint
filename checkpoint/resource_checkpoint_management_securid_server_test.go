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

func TestAccCheckpointManagementSecuridServer_basic(t *testing.T) {

	var securidServerMap map[string]interface{}
	resourceName := "checkpoint_management_securid_server.test"
	objName := "tfTestManagementSecuridServer_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementSecuridServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementSecuridServerConfig(objName, "configfile", "Q0xJRU5UX0lQPSAxLjEuMS4xMQ=="),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementSecuridServerExists(resourceName, &securidServerMap),
					testAccCheckCheckpointManagementSecuridServerAttributes(&securidServerMap, objName, "configfile", "q0xjru5ux0lqpsaxljeums4xmq=="),
				),
			},
		},
	})
}

func testAccCheckpointManagementSecuridServerDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_securid_server" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-securid-server", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success {
				return fmt.Errorf("SecuridServer object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementSecuridServerExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("SecuridServer ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-securid-server", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementSecuridServerAttributes(securidServerMap *map[string]interface{}, name string, configFileName string, base64ConfigFileContent string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		securidServerName := (*securidServerMap)["name"].(string)
		if !strings.EqualFold(securidServerName, name) {
			return fmt.Errorf("name is %s, expected %s", name, securidServerName)
		}
		securidServerConfigFileName := (*securidServerMap)["config-file-name"].(string)
		if !strings.EqualFold(securidServerConfigFileName, configFileName) {
			return fmt.Errorf("configFileName is %s, expected %s", configFileName, securidServerConfigFileName)
		}
		securidServerBase64ConfigFileContent := (*securidServerMap)["base64-config-file-content"].(string)
		if !strings.EqualFold(securidServerBase64ConfigFileContent, base64ConfigFileContent) {
			return fmt.Errorf("base64ConfigFileContent is %s, expected %s", base64ConfigFileContent, securidServerBase64ConfigFileContent)
		}
		return nil
	}
}

func testAccManagementSecuridServerConfig(name string, configFileName string, base64ConfigFileContent string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_securid_server" "test" {
        name = "%s"
        config_file_name = "%s"
        base64_config_file_content = "%s"
}
`, name, configFileName, base64ConfigFileContent)
}
