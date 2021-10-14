package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	"strings"
	"testing"
)

func TestAccCheckpointManagementSimpleCluster_basic(t *testing.T) {
	var simpleClusterMap map[string]interface{}
	resourceName := "checkpoint_management_simple_cluster.test"
	objName := "tfTestManagementCluster_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementSimpleClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementSimpleClusterConfig(objName, "1.2.3.4"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementSimpleClusterExists(resourceName, &simpleClusterMap),
					testAccCheckCheckpointManagementSimpleClusterAttributes(&simpleClusterMap, objName, "1.2.3.4"),
				),
			},
		},
	})
}

func testAccCheckpointManagementSimpleClusterDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_simple_cluster" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-simple-cluster", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success {
				return fmt.Errorf("simple cluster object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementSimpleClusterExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("simple cluster ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-simple-cluster", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementSimpleClusterAttributes(simpleClusterJson *map[string]interface{}, name string, ipv4 string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		simpleClusterName := (*simpleClusterJson)["name"].(string)
		if !strings.EqualFold(simpleClusterName, name) {
			return fmt.Errorf("name is %s, expected %s", simpleClusterName, name)
		}

		simpleClusterIpv4 := (*simpleClusterJson)["ipv4-address"].(string)
		if !strings.EqualFold(simpleClusterIpv4, ipv4) {
			return fmt.Errorf("ipv4 is %s, expected %s", simpleClusterIpv4, ipv4)
		}

		return nil
	}
}

func testAccManagementSimpleClusterConfig(name string, ipv4 string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_checkpoint_host" "checkpoint_host" {
	name = "mycheckpointhost1"
	ipv4_address = "5.6.9.4"
	management_blades = {
		network_policy_management = true
		logging_and_status = true
	}
}

resource "checkpoint_management_simple_cluster" "test" {
	name = "%s"
	ipv4_address = "%s"
	version = "R81"
	hardware = "Open server"
	send_logs_to_server = ["${checkpoint_management_checkpoint_host.checkpoint_host.name}"]
	firewall = true
}
`, name, ipv4)
}
