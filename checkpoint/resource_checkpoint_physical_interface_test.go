package checkpoint

import (
	"fmt"
	checkpoint "github.com/Checkpoint/api_go_sdk/APIFiles"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"log"
	"os"
	"strconv"
	"testing"
)

// Resource physical interface acceptance test:
// 1. Create resource
// 2. Check if resource exists
// 3. Check resource attributes are the same as in configure
// 4. Check resource destroy
func TestAccCheckpointPhysicalInterface_basic(t *testing.T){
	var physical_inter map[string]interface{}
	resourceName := "checkpoint_physical_interface.test"
	objName := "eth1"
	objPhysicalInterface := "20.30.1.2"
	objMaskLen := 24
	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "gaia_api" {
		t.Skip("Skipping Gaia test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{

			{
				Config: testAccPhysicalInterfaceConfig(objName, objPhysicalInterface, objMaskLen),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointPhysicalInterfaceExists(resourceName,&physical_inter),
					testAccCheckCheckpointPhysicalInterfaceAttributes(&physical_inter,objName, objPhysicalInterface, objMaskLen),
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

		response, _ := client.ApiCall("show-physical-interface",payload,client.GetSessionID(),true,false)
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
func testAccCheckCheckpointPhysicalInterfaceAttributes(piRes *map[string]interface{},name string, address string, maskLen int) resource.TestCheckFunc {
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

		inter_address := PIMap["ipv4-address"].(string)
		if inter_address != address {
			return fmt.Errorf("name is %s, expected %s", inter_address, address)
		}

		inter_mask_len, _ := strconv.Atoi(PIMap["ipv4-mask-length"].(string))
		if inter_mask_len != maskLen {
			return fmt.Errorf("name is %d, expected %d", inter_mask_len, maskLen)
		}

		log.Println("Exit testAccCheckCheckpointPhysicalInterfaceAttributes")
		return nil
	}
}

// return a string of the resource like define in a .tf file
func testAccPhysicalInterfaceConfig(name string, address string,masklen int) string {
	return fmt.Sprintf(`
resource "checkpoint_physical_interface" "test" {
      name = "%s"
      ipv4_address = "%s"
      ipv4_mask_length = %d
}
`,name, address, masklen)
}
