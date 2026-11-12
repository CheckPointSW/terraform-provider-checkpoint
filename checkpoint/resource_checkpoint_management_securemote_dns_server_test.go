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

func TestAccCheckpointManagementSecuremoteDnsServer_basic(t *testing.T) {

	var securemoteDnsServerMap map[string]interface{}
	resourceName := "checkpoint_management_securemote_dns_server.test"
	objName := "tfTestManagementSecuremoteDnsServer_" + acctest.RandString(6)
	hostName := "tfTestManagementHost_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementSecuremoteDnsServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementSecuremoteDnsServerConfig(objName, hostName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementSecuremoteDnsServerExists(resourceName, &securemoteDnsServerMap),
					testAccCheckCheckpointManagementSecuremoteDnsServerAttributes(&securemoteDnsServerMap, objName, hostName),
				),
			},
		},
	})
}

func testAccCheckpointManagementSecuremoteDnsServerDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_securemote_dns_server" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-securemote-dns-server", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success {
				return fmt.Errorf("SecuremoteDnsServer object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementSecuremoteDnsServerExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("SecuremoteDnsServer ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-securemote-dns-server", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementSecuremoteDnsServerAttributes(securemoteDnsServerMap *map[string]interface{}, name string, host string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		securemoteDnsServerName := (*securemoteDnsServerMap)["name"].(string)
		if !strings.EqualFold(securemoteDnsServerName, name) {
			return fmt.Errorf("name is %s, expected %s", name, securemoteDnsServerName)
		}
		securemoteDnsServerHost := (*securemoteDnsServerMap)["host"].(map[string]interface{})["name"].(string)
		if !strings.EqualFold(securemoteDnsServerHost, host) {
			return fmt.Errorf("host is %s, expected %s", host, securemoteDnsServerHost)
		}
		return nil
	}
}

func testAccManagementSecuremoteDnsServerConfig(name string, host string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_host" "test_host" {
    name = "%s"
    ipv4_address = "1.1.14.151"
}

resource "checkpoint_management_securemote_dns_server" "test" {
	name = "%s"
	host = "${checkpoint_management_host.test_host.name}"
	domains {
		domain_suffix = ".com"
		maximum_prefix_label_count = 2
    }

  	domains {
    	domain_suffix = ".local"
    	maximum_prefix_label_count = 3
  	}
}
`, host, name)
}
