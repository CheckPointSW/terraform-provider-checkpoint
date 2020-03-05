package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform/helper/resource"
    "github.com/hashicorp/terraform/terraform"
    "os"
    "strings"
    "testing"
    "github.com/hashicorp/terraform/helper/acctest"
)

func TestAccCheckpointManagementMulticastAddressRange_basic(t *testing.T) {

    var multicastAddressRangeMap map[string]interface{}
    resourceName := "checkpoint_management_multicast_address_range.test"
    objName := "tfTestManagementMulticastAddressRange_" + acctest.RandString(6)

    context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
        CheckDestroy: testAccCheckpointManagementMulticastAddressRangeDestroy,
        Steps: []resource.TestStep{
            {
                Config: testAccManagementMulticastAddressRangeConfig(objName, "224.0.0.1", "224.0.0.4"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckCheckpointManagementMulticastAddressRangeExists(resourceName, &multicastAddressRangeMap),
                    testAccCheckCheckpointManagementMulticastAddressRangeAttributes(&multicastAddressRangeMap, objName, "224.0.0.1", "224.0.0.4"),
                ),
            },
        },
    })
}

func testAccCheckpointManagementMulticastAddressRangeDestroy(s *terraform.State) error {

    client := testAccProvider.Meta().(*checkpoint.ApiClient)
    for _, rs := range s.RootModule().Resources {
        if rs.Type != "checkpoint_management_multicast_address_range" {
            continue
        }
        if rs.Primary.ID != "" {
            res, _ := client.ApiCall("show-multicast-address-range", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
            if res.Success {
                return fmt.Errorf("MulticastAddressRange object (%s) still exists", rs.Primary.ID)
            }
        }
        return nil
    }
    return nil
}

func testAccCheckCheckpointManagementMulticastAddressRangeExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        rs, ok := s.RootModule().Resources[resourceTfName]
        if !ok {
            return fmt.Errorf("Resource not found: %s", resourceTfName)
        }

        if rs.Primary.ID == "" {
            return fmt.Errorf("MulticastAddressRange ID is not set")
        }

        client := testAccProvider.Meta().(*checkpoint.ApiClient)

        response, err := client.ApiCall("show-multicast-address-range", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
        if !response.Success {
            return err
        }

        *res = response.GetData()

        return nil
    }
}

func testAccCheckCheckpointManagementMulticastAddressRangeAttributes(multicastAddressRangeMap *map[string]interface{}, name string, ipv4AddressFirst string, ipv4AddressLast string) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        multicastAddressRangeName := (*multicastAddressRangeMap)["name"].(string)
        if !strings.EqualFold(multicastAddressRangeName, name) {
            return fmt.Errorf("name is %s, expected %s", name, multicastAddressRangeName)
        }
        multicastAddressRangeIpv4AddressFirst := (*multicastAddressRangeMap)["ipv4-address-first"].(string)
        if !strings.EqualFold(multicastAddressRangeIpv4AddressFirst, ipv4AddressFirst) {
            return fmt.Errorf("ipv4AddressFirst is %s, expected %s", ipv4AddressFirst, multicastAddressRangeIpv4AddressFirst)
        }
        multicastAddressRangeIpv4AddressLast := (*multicastAddressRangeMap)["ipv4-address-last"].(string)
        if !strings.EqualFold(multicastAddressRangeIpv4AddressLast, ipv4AddressLast) {
            return fmt.Errorf("ipv4AddressLast is %s, expected %s", ipv4AddressLast, multicastAddressRangeIpv4AddressLast)
        }
        return nil
    }
}

func testAccManagementMulticastAddressRangeConfig(name string, ipv4AddressFirst string, ipv4AddressLast string) string {
    return fmt.Sprintf(`
resource "checkpoint_management_multicast_address_range" "test" {
        name = "%s"
        ipv4_address_first = "%s"
        ipv4_address_last = "%s"
}
`, name, ipv4AddressFirst, ipv4AddressLast)
}

