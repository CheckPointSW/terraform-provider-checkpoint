package checkpoint

import (
	"fmt"
	chkp "github.com/Checkpoint/api_go_sdk/APIFiles"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"log"
	"testing"
)

// Resource host acceptance test:
// 1. Create resource
// 2. Check if resource exists
// 3. Validate resource attributes are the same as in configuration
// 4. Check resource destroy
func TestAccChkpHost_basic(t *testing.T){
	t.Log("Enter TestAccChkpHost_basic")
	log.Println("Enter TestAccChkpHost_basic")
	var hostMap map[string]interface{}
	resourceName := "chkp_host.test"
	objName := "tfTestHost_" + acctest.RandString(6)
	resource.Test(t, resource.TestCase{
			PreCheck: func() { testAccPreCheck(t) },
			Providers: testAccProviders,
			CheckDestroy: testAccChkpHostDestroy,
			Steps: []resource.TestStep{
				{
					Config: testAccHostConfig(objName,"192.167.2.3","blue"), //runs "terraform apply"
					Check: resource.ComposeTestCheckFunc(
						testAccCheckChkpHostExists(resourceName, &hostMap),
						testAccCheckChkpHostAttributes(&hostMap, objName,"192.167.2.3","blue"),
					),
				},
			},
	})
	t.Log("Exit TestAccChkpHost_basic...")
	log.Println("Exit TestAccChkpHost_basic...")

}
// verifies Host resource has been destroyed
func testAccChkpHostDestroy(s *terraform.State) error {
	log.Println("Enter testAccChkpHostDestroy")

	client := testAccProvider.Meta().(chkp.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "chkp_host" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-host", map[string]interface{}{"uid": rs.Primary.ID,}, client.GetSessionID(),true,false)
			if res.Success {
				return fmt.Errorf("host object (%s) still exists", rs.Primary.ID)
			}
			_, _ = client.ApiCall("publish", map[string]interface{}{}, client.GetSessionID(), true, false)
		}
		log.Println("Exit testAccChkpHostDestroy (for)")
		return nil
	}
	log.Println("Exit testAccChkpHostDestroy")
	return nil
}

// verifies Host resource exists by ID and init map with attributes
func testAccCheckChkpHostExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Retrieve the resource from state. resourceTfName is the resource name in .tf file:
		// For: resource "chkp_host" "test" {...}
		// resourceTfName = "chkp_host.test"
		log.Println("Enter testAccCheckHostExists...")
		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Host ID is not set")
		}

		// retrieve the client from the test provider
		client := testAccProvider.Meta().(chkp.ApiClient)

		response, err := client.ApiCall("show-host", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(),true,false)
		if !response.Success {
			return err
		}
		// init res - host map object for next step
		*res = response.GetData()
		log.Println("Exit testAccCheckChkpHostExists...")

		_, _ = client.ApiCall("publish", map[string]interface{}{}, client.GetSessionID(), true, false)

		return nil
	}
}
// verifies host resource attributes are same as in configure
func testAccCheckChkpHostAttributes(hostMap *map[string]interface{}, name string, ipv4address string, color string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		log.Println("Enter testAccCheckChkpHostAttributes")
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
		log.Println("Exit testAccCheckChkpHostAttributes")
		return nil

	}
}
// return a string of host resource like define in a .tf file
func testAccHostConfig(name string, ipv4address string, color string) string {
	return fmt.Sprintf(`
resource "chkp_host" "test" {
    name = "%s"
    ipv4_address = "%s"
    color = "%s"
}
`,name,ipv4address,color)
}
