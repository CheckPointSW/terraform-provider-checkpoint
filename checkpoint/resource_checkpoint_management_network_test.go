package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	"testing"
)

// Resource network acceptance test:
// 1. Create resource
// 2. Check if resource exists
// 3. Check resource attributes are the same as in configure
// 4. Check resource destroy
func TestAccCheckpointManagementNetwork_basic(t *testing.T) {
	var network map[string]interface{}
	resourceName := "checkpoint_management_network.test"
	objName := "tfTestManagementNetwork_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointNetworkDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementNetworkConfig(objName, "10.20.0.0", 24),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointNetworkExists(resourceName, &network),
					testAccCheckCheckpointNetworkAttributes(&network, objName, "10.20.0.0", 24),
				),
			},
		},
	})
}

// verifies Network resource has been destroyed
func testAccCheckpointNetworkDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_network" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-network", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success { // Resource still exists. failed to destroy.
				return fmt.Errorf("network object (%s) still exists", rs.Primary.ID)
			}
		}
		break
	}

	return nil
}

func testAccCheckCheckpointNetworkExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("network ID is not set")
		}

		// retrieve the client from test provider. client is after providerConfigure()
		client := testAccProvider.Meta().(*checkpoint.ApiClient)
		response, _ := client.ApiCall("show-network", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return fmt.Errorf(response.ErrorMsg)
		}
		// init res with response data for next step (CheckAttributes)
		*res = response.GetData()

		return nil
	}
}

// verifies Network resource attributes are same as in configure
func testAccCheckCheckpointNetworkAttributes(network *map[string]interface{}, name string, subnet4 string, masklen4 int) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		network := *network
		if network == nil {
			return fmt.Errorf("network is nil")
		}

		networkName := network["name"].(string)
		if networkName != name {
			return fmt.Errorf("name is %s, expected %s", networkName, name)
		}
		networkSb4 := network["subnet4"].(string)
		if networkSb4 != subnet4 {
			return fmt.Errorf("subnet4 is %s, expected %s", networkSb4, subnet4)
		}

		networkMl4 := int(network["mask-length4"].(float64))
		if networkMl4 != masklen4 {
			return fmt.Errorf("mask-length4 is %d, expected %d", networkMl4, masklen4)
		}

		return nil
	}
}

// return a string of Network resource like define in a .tf file
func testAccManagementNetworkConfig(name string, subnet4 string, masklen4 int) string {
	return fmt.Sprintf(`
resource "checkpoint_management_network" "test" {
    name = "%s"
	subnet4 = "%s"
	mask_length4 = "%d"
}
`, name, subnet4, masklen4)
}
