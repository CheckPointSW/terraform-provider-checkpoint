package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"os"
	"testing"
)

// Resource host acceptance test:
// 1. Create resource
// 2. Check if resource exists
// 3. Validate resource attributes are the same as in configuration
// 4. Check resource destroy
func TestAccCheckpointManagementHost_basic(t *testing.T) {

	var hostMap map[string]interface{}
	resourceName := "checkpoint_management_host.test"
	objName := "tfTestManagementHost_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementHostDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementHostConfig(objName, "192.167.2.3", "blue"), //runs "terraform apply"
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementHostExists(resourceName, &hostMap),
					testAccCheckCheckpointManagementHostAttributes(&hostMap, objName, "192.167.2.3", "blue"),
				),
			},
		},
	})

}

// verifies Host resource has been destroyed
func testAccCheckpointManagementHostDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_host" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-host", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("host object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

// verifies Host resource exists by ID and init map with attributes
func testAccCheckCheckpointManagementHostExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Host ID is not set")
		}

		// retrieve the client from the test provider
		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-host", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}
		// init res - host map object for next step
		*res = response.GetData()

		return nil
	}
}

// verifies host resource attributes are same as in configure
func testAccCheckCheckpointManagementHostAttributes(hostMap *map[string]interface{}, name string, ipv4address string, color string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		hostName := (*hostMap)["name"].(string)
		if hostName != name {
			return fmt.Errorf("name is %s, expected %s", hostName, name)
		}
		hostIpv4address := (*hostMap)["ipv4-address"].(string)
		if hostIpv4address != ipv4address {
			return fmt.Errorf("ipv4address is %s, expected %s", hostIpv4address, ipv4address)
		}

		hostColor := (*hostMap)["color"].(string)
		if hostColor != color {
			return fmt.Errorf("color is %s, expected %s", hostColor, color)
		}

		return nil

	}
}

// return a string of host resource like define in a .tf file
func testAccManagementHostConfig(name string, ipv4address string, color string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_host" "test" {
    name = "%s"
    ipv4_address = "%s"
    color = "%s"
}
`, name, ipv4address, color)
}
