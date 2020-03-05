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

func TestAccCheckpointManagementServiceIcmp_basic(t *testing.T) {

    var serviceIcmpMap map[string]interface{}
    resourceName := "checkpoint_management_service_icmp.test"
    objName := "tfTestManagementServiceIcmp_" + acctest.RandString(6)

    context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
        CheckDestroy: testAccCheckpointManagementServiceIcmpDestroy,
        Steps: []resource.TestStep{
            {
                Config: testAccManagementServiceIcmpConfig(objName, 5, 7),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckCheckpointManagementServiceIcmpExists(resourceName, &serviceIcmpMap),
                    testAccCheckCheckpointManagementServiceIcmpAttributes(&serviceIcmpMap, objName, 5, 7),
                ),
            },
        },
    })
}

func testAccCheckpointManagementServiceIcmpDestroy(s *terraform.State) error {

    client := testAccProvider.Meta().(*checkpoint.ApiClient)
    for _, rs := range s.RootModule().Resources {
        if rs.Type != "checkpoint_management_service_icmp" {
            continue
        }
        if rs.Primary.ID != "" {
            res, _ := client.ApiCall("show-service-icmp", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
            if res.Success {
                return fmt.Errorf("ServiceIcmp object (%s) still exists", rs.Primary.ID)
            }
        }
        return nil
    }
    return nil
}

func testAccCheckCheckpointManagementServiceIcmpExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        rs, ok := s.RootModule().Resources[resourceTfName]
        if !ok {
            return fmt.Errorf("Resource not found: %s", resourceTfName)
        }

        if rs.Primary.ID == "" {
            return fmt.Errorf("ServiceIcmp ID is not set")
        }

        client := testAccProvider.Meta().(*checkpoint.ApiClient)

        response, err := client.ApiCall("show-service-icmp", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
        if !response.Success {
            return err
        }

        *res = response.GetData()

        return nil
    }
}

func testAccCheckCheckpointManagementServiceIcmpAttributes(serviceIcmpMap *map[string]interface{}, name string, icmpType int, icmpCode int) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        serviceIcmpName := (*serviceIcmpMap)["name"].(string)
        if !strings.EqualFold(serviceIcmpName, name) {
            return fmt.Errorf("name is %s, expected %s", name, serviceIcmpName)
        }
        serviceIcmpIcmpType := int((*serviceIcmpMap)["icmp-type"].(float64))
        if serviceIcmpIcmpType != icmpType {
            return fmt.Errorf("icmpType is %d, expected %d", icmpType, serviceIcmpIcmpType)
        }
        serviceIcmpIcmpCode := int((*serviceIcmpMap)["icmp-code"].(float64))
        if serviceIcmpIcmpCode != icmpCode {
            return fmt.Errorf("icmpCode is %d, expected %d", icmpCode, serviceIcmpIcmpCode)
        }
        return nil
    }
}

func testAccManagementServiceIcmpConfig(name string, icmpType int, icmpCode int) string {
    return fmt.Sprintf(`
resource "checkpoint_management_service_icmp" "test" {
        name = "%s"
        icmp_type = %d
        icmp_code = %d
}
`, name, icmpType, icmpCode)
}

