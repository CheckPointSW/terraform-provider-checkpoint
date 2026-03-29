package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"os"
	"strings"
	"testing"
)

func TestAccCheckpointManagementMds_basic(t *testing.T) {

	var mdsMap map[string]interface{}
	resourceName := "checkpoint_management_mds.test"
	objName := "tfTestManagementMds_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementMdsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementMdsConfig(objName, "multi-domain server", "Gaia", "Open server", "1.1.1.1", "2.2.2.2", "3.3.3.3"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementMdsExists(resourceName, &mdsMap),
					testAccCheckCheckpointManagementMdsAttributes(&mdsMap, objName, "multi-domain server", "Gaia", "Open server", "1.1.1.1", "2.2.2.2", "3.3.3.3"),
				),
			},
		},
	})
}

func testAccCheckpointManagementMdsDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_mds" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-mds", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("Mds object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementMdsExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Mds ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-mds", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementMdsAttributes(mdsMap *map[string]interface{}, name string, serverType string, os string, hardware string, ipv4Address string, ipPoolFirst string, ipPoolLast string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		mdsName := (*mdsMap)["name"].(string)
		if !strings.EqualFold(mdsName, name) {
			return fmt.Errorf("name is %s, expected %s", name, mdsName)
		}
		mdsServerType := (*mdsMap)["server-type"].(string)
		if !strings.EqualFold(mdsServerType, serverType) {
			return fmt.Errorf("serverType is %s, expected %s", serverType, mdsServerType)
		}
		mdsOs := (*mdsMap)["os"].(map[string]interface{})["name"].(string)
		if !strings.EqualFold(mdsOs, os) {
			return fmt.Errorf("os is %s, expected %s", os, mdsOs)
		}
		mdsHardware := (*mdsMap)["hardware"].(map[string]interface{})["name"].(string)
		if !strings.EqualFold(mdsHardware, hardware) {
			return fmt.Errorf("hardware is %s, expected %s", hardware, mdsHardware)
		}
		mdsIpv4Address := (*mdsMap)["ipv4-address"].(string)
		if !strings.EqualFold(mdsIpv4Address, ipv4Address) {
			return fmt.Errorf("ipv4Address is %s, expected %s", ipv4Address, mdsIpv4Address)
		}
		mdsIpPoolFirst := (*mdsMap)["ip-pool-first"].(string)
		if !strings.EqualFold(mdsIpPoolFirst, ipPoolFirst) {
			return fmt.Errorf("ipPoolFirst is %s, expected %s", ipPoolFirst, mdsIpPoolFirst)
		}
		mdsIpPoolLast := (*mdsMap)["ip-pool-last"].(string)
		if !strings.EqualFold(mdsIpPoolLast, ipPoolLast) {
			return fmt.Errorf("ipPoolLast is %s, expected %s", ipPoolLast, mdsIpPoolLast)
		}
		return nil
	}
}

func testAccManagementMdsConfig(name string, serverType string, os string, hardware string, ipv4Address string, ipPoolFirst string, ipPoolLast string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_mds" "test" {
        name = "%s"
        server_type = "%s"
        os = "%s"
        hardware = "%s"
        ipv4_address = "%s"
        ip_pool_first = "%s"
        ip_pool_last = "%s"
}
`, name, serverType, os, hardware, ipv4Address, ipPoolFirst, ipPoolLast)
}
