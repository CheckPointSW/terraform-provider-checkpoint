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

func TestAccCheckpointManagementDnsDomain_basic(t *testing.T) {

    var dnsDomainMap map[string]interface{}
    resourceName := "checkpoint_management_dns_domain.test"
    objName := ".tfTestManagementDnsDomain_" + acctest.RandString(6)

    context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
        CheckDestroy: testAccCheckpointManagementDnsDomainDestroy,
        Steps: []resource.TestStep{
            {
                Config: testAccManagementDnsDomainConfig(objName, false),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckCheckpointManagementDnsDomainExists(resourceName, &dnsDomainMap),
                    testAccCheckCheckpointManagementDnsDomainAttributes(&dnsDomainMap, objName, false),
                ),
            },
        },
    })
}

func testAccCheckpointManagementDnsDomainDestroy(s *terraform.State) error {

    client := testAccProvider.Meta().(*checkpoint.ApiClient)
    for _, rs := range s.RootModule().Resources {
        if rs.Type != "checkpoint_management_dns_domain" {
            continue
        }
        if rs.Primary.ID != "" {
            res, _ := client.ApiCall("show-dns-domain", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
            if res.Success {
                return fmt.Errorf("DnsDomain object (%s) still exists", rs.Primary.ID)
            }
        }
        return nil
    }
    return nil
}

func testAccCheckCheckpointManagementDnsDomainExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        rs, ok := s.RootModule().Resources[resourceTfName]
        if !ok {
            return fmt.Errorf("Resource not found: %s", resourceTfName)
        }

        if rs.Primary.ID == "" {
            return fmt.Errorf("DnsDomain ID is not set")
        }

        client := testAccProvider.Meta().(*checkpoint.ApiClient)

        response, err := client.ApiCall("show-dns-domain", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
        if !response.Success {
            return err
        }

        *res = response.GetData()

        return nil
    }
}

func testAccCheckCheckpointManagementDnsDomainAttributes(dnsDomainMap *map[string]interface{}, name string, isSubDomain bool) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        dnsDomainName := (*dnsDomainMap)["name"].(string)
        if !strings.EqualFold(dnsDomainName, name) {
            return fmt.Errorf("name is %s, expected %s", name, dnsDomainName)
        }
        dnsDomainIsSubDomain := (*dnsDomainMap)["is-sub-domain"].(bool)
        if dnsDomainIsSubDomain != isSubDomain {
            return fmt.Errorf("isSubDomain is %t, expected %t", isSubDomain, dnsDomainIsSubDomain)
        }
        return nil
    }
}

func testAccManagementDnsDomainConfig(name string, isSubDomain bool) string {
    return fmt.Sprintf(`
resource "checkpoint_management_dns_domain" "test" {
        name = "%s"
        is_sub_domain = %t
}
`, name, isSubDomain)
}

