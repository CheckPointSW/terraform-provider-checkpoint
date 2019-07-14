package main

import (
	chkp "api_go_sdk/APIFiles"
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"log"
	"testing"
)

// Resource network acceptance test:
// 1. Create resource
// 2. Check if resource exists
// 3. Check resource attributes are the same as in configure
// 4. Check resource destroy
func TestAccChkpNetwork_basic(t *testing.T){
	var network map[string]interface{}
	resourceName := "chkp_network.test"
	objName := "tfTestNetwork_" + acctest.RandString(6)
	resource.Test(t, resource.TestCase{
			PreCheck: func() { testAccPreCheck(t) },
			Providers: testAccProviders,
			CheckDestroy: testAccChkpNetworkDestroy,
			Steps: []resource.TestStep{
				{
					Config: testAccNetworkConfig(objName,"10.20.0.0",24),
					Check: resource.ComposeTestCheckFunc(
						testAccCheckChkpNetworkExists(resourceName,&network),
						testAccCheckChkpNetworkAttributes(&network,objName,"10.20.0.0",24),
					),
				},
			},
	})
}

// verifies Network resource has been destroyed
func testAccChkpNetworkDestroy(s *terraform.State) error {
	log.Println("Enter testAccChkpNetworkDestroy")
	client := testAccProvider.Meta().(*chkp.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "chkp_network" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-network", map[string]interface{}{"uid": rs.Primary.ID,},client.GetSessionID(),true,false)
			if res.Success { // Resource still exists. failed to destroy.
				return fmt.Errorf("network object (%s) still exists", rs.Primary.ID)
			}
		}
		break
	}
	log.Println("Exit testAccChkpNetworkDestroy")
	return nil
}
// verifies Network resource exists by ID and init res with response data
func testAccCheckChkpNetworkExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Retrieve the resource from state. resourceTfName is the terraform resource name in .tf file:
		// For: resource "chkp_network" "test" {...}
		// resourceTfName = "chkp_network.test"
		log.Println("Enter testAccCheckChkpNetworkExists...")
		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("network ID is not set")
		}

		// retrieve the client from test provider. client is after providerConfigure()
		client := testAccProvider.Meta().(*chkp.ApiClient)
		response, _ := client.ApiCall("show-network",map[string]interface{}{"uid": rs.Primary.ID},client.GetSessionID(),true,false)
		if !response.Success {
			return fmt.Errorf(response.ErrorMsg)
		}
		// init res with response data for next step (CheckAttributes)
		*res = response.GetData()
		log.Println("Exit testAccCheckChkpNetworkExists...")
		return nil
	}
}
// verifies Network resource attributes are same as in configure
func testAccCheckChkpNetworkAttributes(network *map[string]interface{},name string,subnet4 string,masklen4 int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		log.Println("Enter testAccCheckChkpNetworkAttributes")
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
		log.Println("Exit testAccCheckChkpNetworkAttributes")
		return nil
	}
}

// return a string of Network resource like define in a .tf file
func testAccNetworkConfig(name string, subnet4 string,masklen4 int) string {
	return fmt.Sprintf(`
resource "chkp_network" "test" {
    name = "%s"
	subnet4 = "%s"
	mask_length4 = "%d"
}
`,name,subnet4,masklen4)
}
