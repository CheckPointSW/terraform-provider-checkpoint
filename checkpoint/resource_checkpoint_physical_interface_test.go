package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"log"
	"os"
	"testing"
)

// Resource physical interface acceptance test:
// 1. Create resource
// 2. Check if resource exists
// 3. Check resource attributes are the same as in configure
// 4. Check resource destroy
func TestAccCheckpointPhysicalInterface_basic(t *testing.T) {
	var physical_inter map[string]interface{}
	resourceName := "checkpoint_physical_interface.test"
	objName := "eth0"
	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "gaia_api" {
		t.Skip("Skipping Gaia test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{

			{
				Config: testAccPhysicalInterfaceConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointPhysicalInterfaceExists(resourceName, &physical_inter),
					testAccCheckCheckpointPhysicalInterfaceAttributes(&physical_inter, objName),
				),
			},
		},
	})
}

// verifies resource exists by ID and init res with response data
func testAccCheckCheckpointPhysicalInterfaceExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		log.Println("Enter testAccCheckCheckpointPhysicalInterfaceExists...")
		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}

		// retrieve the client from test provider. client is after providerConfigure()
		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		payload := make(map[string]interface{})

		payload["name"] = rs.Primary.Attributes["name"]

		response, _ := client.ApiCall("show-physical-interface", payload, client.GetSessionID(), true, false)
		if !response.Success {
			return fmt.Errorf(response.ErrorMsg)
		}
		// init res with response data for next step (CheckAttributes)
		*res = response.GetData()
		log.Println("Exit testAccCheckCheckpointPhysicalInterfaceExists...")
		return nil
	}
}

// verifies resource attributes are same as in configure
func testAccCheckCheckpointPhysicalInterfaceAttributes(piRes *map[string]interface{}, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		log.Println("Enter testAccCheckCheckpointPhysicalInterfaceAttributes")
		PIMap := *piRes
		if PIMap == nil {
			return fmt.Errorf("PIMap is nil")
		}

		inter_name := PIMap["name"].(string)
		if inter_name != name {
			return fmt.Errorf("name is %s, expected %s", inter_name, name)
		}

		enabled := PIMap["enabled"].(bool)
		if enabled != true {
			return fmt.Errorf("enabled is %t, expected true", enabled)
		}

		log.Println("Exit testAccCheckCheckpointPhysicalInterfaceAttributes")
		return nil
	}
}

// return a string of the resource like define in a .tf file
func testAccPhysicalInterfaceConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_physical_interface" "test" {
      name = "%s"
      enabled = "true"
}
`, name)
}
