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

func TestAccCheckpointManagementAddressRange_basic(t *testing.T) {
	var addressRangeMap map[string]interface{}
	resourceName := "checkpoint_management_address_range.test"
	objName := "tfTestManagementAddressRange_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementAddressRangeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementAddressRangeConfig(objName, "10.123.174.32", "10.123.174.35"), //runs "terraform apply"
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementAddressRangeExists(resourceName, &addressRangeMap),
					testAccCheckCheckpointManagementAddressRangeAttributes(&addressRangeMap, objName, "10.123.174.32", "10.123.174.35"),
				),
			},
		},
	})

}

func testAccCheckpointManagementAddressRangeDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_address_range" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-address-range", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("address-range object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementAddressRangeExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("Address Range ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-address-range", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementAddressRangeAttributes(addressRangeMap *map[string]interface{}, name string, ipv4AddressFirst string, ipv4AddressLast string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		addressRangeName := (*addressRangeMap)["name"].(string)
		if addressRangeName != name {
			return fmt.Errorf("name is %s, expected %s", addressRangeName, name)
		}

		addressRangeIpv4addressFirst := (*addressRangeMap)["ipv4-address-first"].(string)
		if addressRangeIpv4addressFirst != ipv4AddressFirst {
			return fmt.Errorf("ipv4-address-first is %s, expected %s", addressRangeIpv4addressFirst, ipv4AddressFirst)
		}

		addressRangeIpv4addressLast := (*addressRangeMap)["ipv4-address-last"].(string)
		if addressRangeIpv4addressLast != ipv4AddressLast {
			return fmt.Errorf("ipv4-address-last is %s, expected %s", addressRangeIpv4addressLast, ipv4AddressLast)
		}

		return nil

	}
}

func testAccManagementAddressRangeConfig(name string, ipv4AddressFirst string, ipv4AddressLast string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_address_range" "test" {
    name = "%s"
    ipv4_address_first = "%s"
    ipv4_address_last = "%s"
}
`, name, ipv4AddressFirst, ipv4AddressLast)
}
