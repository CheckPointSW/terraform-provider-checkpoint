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

func TestAccCheckpointManagementTacacsServer_basic(t *testing.T) {

	var tacacsServerMap map[string]interface{}
	resourceName := "checkpoint_management_tacacs_server.test"
	objName := "tfTestManagementTacacsServer_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementTacacsServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementTacacsServerConfig(objName, "yoni"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementTacacsServerExists(resourceName, &tacacsServerMap),
					testAccCheckCheckpointManagementTacacsServerAttributes(&tacacsServerMap, objName, "yoni"),
				),
			},
		},
	})
}

func testAccCheckpointManagementTacacsServerDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_tacacs_server" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-tacacs-server", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("TacacsServer object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementTacacsServerExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("TacacsServer ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-tacacs-server", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementTacacsServerAttributes(tacacsServerMap *map[string]interface{}, name string, server string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		tacacsServerName := (*tacacsServerMap)["name"].(string)
		if !strings.EqualFold(tacacsServerName, name) {
			return fmt.Errorf("name is %s, expected %s", name, tacacsServerName)
		}
		tacacsServerServerMap := (*tacacsServerMap)["server"].(map[string]interface{})
		tacacsServerServer := tacacsServerServerMap["name"].(string)
		if !strings.EqualFold(tacacsServerServer, server) {
			return fmt.Errorf("server is %s, expected %s", server, tacacsServerServer)
		}
		return nil
	}
}

func testAccManagementTacacsServerConfig(name string, server string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_tacacs_server" "test" {
        name = "%s"
        server = "%s"
}
`, name, server)
}
