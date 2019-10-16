package checkpoint

import (
	"fmt"
	chkp "github.com/Checkpoint/api_go_sdk/APIFiles"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"log"
	"testing"
)

// Resource hostname acceptance test:
// 1. Create resource
// 2. Check if resource exists
// 3. Check resource attributes are the same as in configure
// 4. Check resource destroy
func TestAccChkpHostname_basic(t *testing.T){
	var hostname map[string]interface{}
	resourceName := "chkp_hostname.test"
	objName := "terratest"
	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{

			{
				Config: testAccHostnameConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckChkpHostnameExists(resourceName,&hostname),
					testAccCheckChkpHostnameAttributes(&hostname,objName),
				),
			},
		},
	})
}

// verifies resource exists by ID and init res with response data
func testAccCheckChkpHostnameExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		log.Println("Enter testAccCheckChkpHostnameExists...")
		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}

		// retrieve the client from test provider. client is after providerConfigure()
		client := testAccProvider.Meta().(*chkp.ApiClient)
		response, _ := client.ApiCall("show-hostname",map[string]interface{}{},client.GetSessionID(),true,false)
		if !response.Success {
			return fmt.Errorf(response.ErrorMsg)
		}
		// init res with response data for next step (CheckAttributes)
		*res = response.GetData()
		log.Println("Exit testAccCheckChkpHostnameExists...")
		return nil
	}
}

// verifies resource attributes are same as in configure
func testAccCheckChkpHostnameAttributes(hostname *map[string]interface{},name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		log.Println("Enter testAccCheckChkpHostnameAttributes")
		hostnameMap := *hostname
		if hostnameMap == nil {
			return fmt.Errorf("hostnameMap is nil")
		}

		hostname := hostnameMap["name"].(string)
		if hostname != name {
			return fmt.Errorf("name is %s, expected %s", hostname, name)
		}
		log.Println("Exit testAccCheckChkpHostnameAttributes")
		return nil
	}
}

// return a string of the resource like define in a .tf file
func testAccHostnameConfig(name string) string {
	return fmt.Sprintf(`
resource "chkp_hostname" "test" {
    name = "%s"
}
`,name)
}
