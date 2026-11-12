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

func TestAccCheckpointManagementCheckpointHost_basic(t *testing.T) {

	var checkpointHostMap map[string]interface{}
	resourceName := "checkpoint_management_checkpoint_host.test"
	objName := "tfTestManagementCheckpointHost_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementCheckpointHostDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementCheckpointHostConfig(objName, "5.5.5.5"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementCheckpointHostExists(resourceName, &checkpointHostMap),
					testAccCheckCheckpointManagementCheckpointHostAttributes(&checkpointHostMap, objName, "5.5.5.5"),
				),
			},
		},
	})
}

func testAccCheckpointManagementCheckpointHostDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_checkpoint_host" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-checkpoint-host", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("CheckpointHost object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementCheckpointHostExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("CheckpointHost ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-checkpoint-host", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementCheckpointHostAttributes(checkpointHostMap *map[string]interface{}, name string, ipv4Address string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		checkpointHostName := (*checkpointHostMap)["name"].(string)
		if !strings.EqualFold(checkpointHostName, name) {
			return fmt.Errorf("name is %s, expected %s", name, checkpointHostName)
		}
		checkpointHostIpv4Address := (*checkpointHostMap)["ipv4-address"].(string)
		if !strings.EqualFold(checkpointHostIpv4Address, ipv4Address) {
			return fmt.Errorf("ipv4Address is %s, expected %s", ipv4Address, checkpointHostIpv4Address)
		}
		return nil
	}
}

func testAccManagementCheckpointHostConfig(name string, ipv4Address string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_checkpoint_host" "test" {
        name = "%s"
        ipv4_address = "%s"
}
`, name, ipv4Address)
}
