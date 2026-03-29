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

func TestAccCheckpointManagementTrustedClient_basic(t *testing.T) {

	var trustedClientMap map[string]interface{}
	resourceName := "checkpoint_management_trusted_client.test"
	objName := "tfTestManagementTrustedClient_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementTrustedClientDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementTrustedClientConfig(objName, "192.168.2.1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementTrustedClientExists(resourceName, &trustedClientMap),
					testAccCheckCheckpointManagementTrustedClientAttributes(&trustedClientMap, objName, "192.168.2.1"),
				),
			},
		},
	})
}

func testAccCheckpointManagementTrustedClientDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_trustedClient" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-trusted-client", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("TrustedClient object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementTrustedClientExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("TrustedClient ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-trusted-client", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementTrustedClientAttributes(trustedClientMap *map[string]interface{}, name string, ipv4Address string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		trustedClientName := (*trustedClientMap)["name"].(string)
		if !strings.EqualFold(trustedClientName, name) {
			return fmt.Errorf("name is %s, expected %s", name, trustedClientName)
		}
		trustedClientIpv4Address := (*trustedClientMap)["ipv4-address"].(string)
		if !strings.EqualFold(trustedClientIpv4Address, ipv4Address) {
			return fmt.Errorf("ipv4Address is %s, expected %s", ipv4Address, trustedClientIpv4Address)
		}
		return nil
	}
}

func testAccManagementTrustedClientConfig(name string, ipv4Address string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_trusted_client" "test" {
        name = "%s"
        ipv4_address = "%s"
}
`, name, ipv4Address)
}
