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

func TestAccCheckpointManagementWildcard_basic(t *testing.T) {

    var wildcardMap map[string]interface{}
    resourceName := "checkpoint_management_wildcard.test"
    objName := "tfTestManagementWildcard_" + acctest.RandString(6)

    context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
        CheckDestroy: testAccCheckpointManagementWildcardDestroy,
        Steps: []resource.TestStep{
            {
                Config: testAccManagementWildcardConfig(objName, "192.168.2.1", "0.0.0.128"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckCheckpointManagementWildcardExists(resourceName, &wildcardMap),
                    testAccCheckCheckpointManagementWildcardAttributes(&wildcardMap, objName, "192.168.2.1", "0.0.0.128"),
                ),
            },
        },
    })
}

func testAccCheckpointManagementWildcardDestroy(s *terraform.State) error {

    client := testAccProvider.Meta().(*checkpoint.ApiClient)
    for _, rs := range s.RootModule().Resources {
        if rs.Type != "checkpoint_management_wildcard" {
            continue
        }
        if rs.Primary.ID != "" {
            res, _ := client.ApiCall("show-wildcard", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
            if res.Success {
                return fmt.Errorf("Wildcard object (%s) still exists", rs.Primary.ID)
            }
        }
        return nil
    }
    return nil
}

func testAccCheckCheckpointManagementWildcardExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        rs, ok := s.RootModule().Resources[resourceTfName]
        if !ok {
            return fmt.Errorf("Resource not found: %s", resourceTfName)
        }

        if rs.Primary.ID == "" {
            return fmt.Errorf("Wildcard ID is not set")
        }

        client := testAccProvider.Meta().(*checkpoint.ApiClient)

        response, err := client.ApiCall("show-wildcard", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
        if !response.Success {
            return err
        }

        *res = response.GetData()

        return nil
    }
}

func testAccCheckCheckpointManagementWildcardAttributes(wildcardMap *map[string]interface{}, name string, ipv4Address string, ipv4MaskWildcard string) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        wildcardName := (*wildcardMap)["name"].(string)
        if !strings.EqualFold(wildcardName, name) {
            return fmt.Errorf("name is %s, expected %s", name, wildcardName)
        }
        wildcardIpv4Address := (*wildcardMap)["ipv4-address"].(string)
        if !strings.EqualFold(wildcardIpv4Address, ipv4Address) {
            return fmt.Errorf("ipv4Address is %s, expected %s", ipv4Address, wildcardIpv4Address)
        }
        wildcardIpv4MaskWildcard := (*wildcardMap)["ipv4-mask-wildcard"].(string)
        if !strings.EqualFold(wildcardIpv4MaskWildcard, ipv4MaskWildcard) {
            return fmt.Errorf("ipv4MaskWildcard is %s, expected %s", ipv4MaskWildcard, wildcardIpv4MaskWildcard)
        }
        return nil
    }
}

func testAccManagementWildcardConfig(name string, ipv4Address string, ipv4MaskWildcard string) string {
    return fmt.Sprintf(`
resource "checkpoint_management_wildcard" "test" {
        name = "%s"
        ipv4_address = "%s"
        ipv4_mask_wildcard = "%s"
}
`, name, ipv4Address, ipv4MaskWildcard)
}

