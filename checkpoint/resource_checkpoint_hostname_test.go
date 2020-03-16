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

// Resource hostname acceptance test:
// 1. Create resource
// 2. Check if resource exists
// 3. Check resource attributes are the same as in configure
// 4. Check resource destroy
func TestAccCheckpointHostname_basic(t *testing.T) {
	var hostname map[string]interface{}
	resourceName := "checkpoint_hostname.test"
	objName := "terratest"
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
				Config: testAccHostnameConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointHostnameExists(resourceName, &hostname),
					testAccCheckCheckpointHostnameAttributes(&hostname, objName),
				),
			},
		},
	})
}

// verifies resource exists by ID and init res with response data
func testAccCheckCheckpointHostnameExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		log.Println("Enter testAccCheckCheckpointHostnameExists...")
		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}

		// retrieve the client from test provider. client is after providerConfigure()
		client := testAccProvider.Meta().(*checkpoint.ApiClient)
		response, _ := client.ApiCall("show-hostname", map[string]interface{}{}, client.GetSessionID(), true, false)
		if !response.Success {
			return fmt.Errorf(response.ErrorMsg)
		}
		// init res with response data for next step (CheckAttributes)
		*res = response.GetData()
		log.Println("Exit testAccCheckCheckpointHostnameExists...")
		return nil
	}
}

// verifies resource attributes are same as in configure
func testAccCheckCheckpointHostnameAttributes(hostname *map[string]interface{}, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		log.Println("Enter testAccCheckCheckpointHostnameAttributes")
		hostnameMap := *hostname
		if hostnameMap == nil {
			return fmt.Errorf("hostnameMap is nil")
		}

		hostname := hostnameMap["name"].(string)
		if hostname != name {
			return fmt.Errorf("name is %s, expected %s", hostname, name)
		}
		log.Println("Exit testAccCheckCheckpointHostnameAttributes")
		return nil
	}
}

// return a string of the resource like define in a .tf file
func testAccHostnameConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_hostname" "test" {
    name = "%s"
}
`, name)
}
