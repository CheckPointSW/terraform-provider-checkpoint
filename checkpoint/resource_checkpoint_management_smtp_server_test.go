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

func TestAccCheckpointManagementSmtpServer_basic(t *testing.T) {

	var smtpServerMap map[string]interface{}
	resourceName := "checkpoint_management_smtp_server.test"
	objName := "tfTestManagementSmtpServer_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementSmtpServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementSmtpServerConfig(objName, "smtp.example.com", 25, "none"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementSmtpServerExists(resourceName, &smtpServerMap),
					testAccCheckCheckpointManagementSmtpServerAttributes(&smtpServerMap, objName, "smtp.example.com", 25, "none"),
				),
			},
		},
	})
}

func testAccCheckpointManagementSmtpServerDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_smtp_server" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-smtp-server", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("SmtpServer object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementSmtpServerExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("SmtpServer ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-smtp-server", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementSmtpServerAttributes(smtpServerMap *map[string]interface{}, name string, server string, port int, encryption string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		smtpServerName := (*smtpServerMap)["name"].(string)
		if !strings.EqualFold(smtpServerName, name) {
			return fmt.Errorf("name is %s, expected %s", name, smtpServerName)
		}
		smtpServerServer := (*smtpServerMap)["server"].(string)
		if !strings.EqualFold(smtpServerServer, server) {
			return fmt.Errorf("server is %s, expected %s", server, smtpServerServer)
		}
		smtpServerPort := int((*smtpServerMap)["port"].(float64))
		if smtpServerPort != port {
			return fmt.Errorf("port is %d, expected %d", port, smtpServerPort)
		}
		smtpServerEncryption := (*smtpServerMap)["encryption"].(string)
		if !strings.EqualFold(smtpServerEncryption, encryption) {
			return fmt.Errorf("encryption is %s, expected %s", encryption, smtpServerEncryption)
		}
		return nil
	}
}

func testAccManagementSmtpServerConfig(name string, server string, port int, encryption string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_smtp_server" "test" {
        name = "%s"
        server = "%s"
        port = %d
        encryption = "%s"
}
`, name, server, port, encryption)
}
