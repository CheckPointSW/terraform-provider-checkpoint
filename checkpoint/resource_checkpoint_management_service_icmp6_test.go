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

func TestAccCheckpointManagementServiceIcmp6_basic(t *testing.T) {

    var serviceIcmp6Map map[string]interface{}
    resourceName := "checkpoint_management_service_icmp6.test"
    objName := "tfTestManagementServiceIcmp6_" + acctest.RandString(6)

    context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
        CheckDestroy: testAccCheckpointManagementServiceIcmp6Destroy,
        Steps: []resource.TestStep{
            {
                Config: testAccManagementServiceIcmp6Config(objName, 5, 7),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckCheckpointManagementServiceIcmp6Exists(resourceName, &serviceIcmp6Map),
                    testAccCheckCheckpointManagementServiceIcmp6Attributes(&serviceIcmp6Map, objName, 5, 7),
                ),
            },
        },
    })
}

func testAccCheckpointManagementServiceIcmp6Destroy(s *terraform.State) error {

    client := testAccProvider.Meta().(*checkpoint.ApiClient)
    for _, rs := range s.RootModule().Resources {
        if rs.Type != "checkpoint_management_service_icmp6" {
            continue
        }
        if rs.Primary.ID != "" {
            res, _ := client.ApiCall("show-service-icmp6", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
            if res.Success {
                return fmt.Errorf("ServiceIcmp6 object (%s) still exists", rs.Primary.ID)
            }
        }
        return nil
    }
    return nil
}

func testAccCheckCheckpointManagementServiceIcmp6Exists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        rs, ok := s.RootModule().Resources[resourceTfName]
        if !ok {
            return fmt.Errorf("Resource not found: %s", resourceTfName)
        }

        if rs.Primary.ID == "" {
            return fmt.Errorf("ServiceIcmp6 ID is not set")
        }

        client := testAccProvider.Meta().(*checkpoint.ApiClient)

        response, err := client.ApiCall("show-service-icmp6", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
        if !response.Success {
            return err
        }

        *res = response.GetData()

        return nil
    }
}

func testAccCheckCheckpointManagementServiceIcmp6Attributes(serviceIcmp6Map *map[string]interface{}, name string, icmpType int, icmpCode int) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        serviceIcmp6Name := (*serviceIcmp6Map)["name"].(string)
        if !strings.EqualFold(serviceIcmp6Name, name) {
            return fmt.Errorf("name is %s, expected %s", name, serviceIcmp6Name)
        }
        serviceIcmp6IcmpType := int((*serviceIcmp6Map)["icmp-type"].(float64))
        if serviceIcmp6IcmpType != icmpType {
            return fmt.Errorf("icmpType is %d, expected %d", icmpType, serviceIcmp6IcmpType)
        }
        serviceIcmp6IcmpCode := int((*serviceIcmp6Map)["icmp-code"].(float64))
        if serviceIcmp6IcmpCode != icmpCode {
            return fmt.Errorf("icmpCode is %d, expected %d", icmpCode, serviceIcmp6IcmpCode)
        }
        return nil
    }
}

func testAccManagementServiceIcmp6Config(name string, icmpType int, icmpCode int) string {
    return fmt.Sprintf(`
resource "checkpoint_management_service_icmp6" "test" {
        name = "%s"
        icmp_type = %d
        icmp_code = %d
}
`, name, icmpType, icmpCode)
}

