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

func TestAccCheckpointManagementResourceTcp_basic(t *testing.T) {

	var resourceTcpMap map[string]interface{}
	resourceName := "checkpoint_management_resource_tcp.test"
	objName := "tfTestManagementResourceTcp_" + acctest.RandString(6)
	hostName := "tfTestManagementHost_" + acctest.RandString(6)
	opsecAppName := "tfTestManagementOpsecApp_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementResourceTcpDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementResourceTcpConfig(hostName, opsecAppName, objName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementResourceTcpExists(resourceName, &resourceTcpMap),
					testAccCheckCheckpointManagementResourceTcpAttributes(&resourceTcpMap, objName, opsecAppName),
				),
			},
		},
	})
}

func testAccCheckpointManagementResourceTcpDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_resource_tcp" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-resource-tcp", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success {
				return fmt.Errorf("ResourceTcp object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementResourceTcpExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ResourceTcp ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-resource-tcp", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementResourceTcpAttributes(resourceTcpMap *map[string]interface{}, name string, opsecAppName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		resourceTcpName := (*resourceTcpMap)["name"].(string)
		if !strings.EqualFold(resourceTcpName, name) {
			return fmt.Errorf("name is %s, expected %s", name, resourceTcpName)
		}

		ufpSettingsServer := (*resourceTcpMap)["ufp-settings"].(map[string]interface{})["server"].(map[string]interface{})["name"].(string)
		if !strings.EqualFold(ufpSettingsServer, opsecAppName) {
			return fmt.Errorf("server is %s, expected %s", ufpSettingsServer, opsecAppName)
		}
		return nil
	}
}

func testAccManagementResourceTcpConfig(hostName string, opsecAppName string, objName string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_host" "test_host" {
    name = "%s"
    ipv4_address = "88.88.88.88"
}

resource "checkpoint_management_opsec_application" "test_opsec_app" {
    name = "%s"
    host = "${checkpoint_management_host.test_host.name}"
    one_time_password = "SomePassword"
}

resource "checkpoint_management_resource_tcp" "test" {
    name = "%s"
    ufp_settings {
        server = "${checkpoint_management_opsec_application.test_opsec_app.id}"
    }
}
`, hostName, opsecAppName, objName)
}
