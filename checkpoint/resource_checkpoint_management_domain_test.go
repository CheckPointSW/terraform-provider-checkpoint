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

func TestAccCheckpointManagementDomain_basic(t *testing.T) {

	var domainMap map[string]interface{}
	resourceName := "checkpoint_management_domain.test"
	objName := "tfTestManagementDomain_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementDomainConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementDomainExists(resourceName, &domainMap),
					testAccCheckCheckpointManagementDomainAttributes(&domainMap, objName),
				),
			},
		},
	})
}

func testAccCheckpointManagementDomainDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_domain" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-domain", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("Domain object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementDomainExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Domain ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-domain", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementDomainAttributes(domainMap *map[string]interface{}, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		domainName := (*domainMap)["name"].(string)
		if !strings.EqualFold(domainName, name) {
			return fmt.Errorf("name is %s, expected %s", name, domainName)
		}
		return nil
	}
}

func testAccManagementDomainConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_domain" "test" {
        name = "%s"
        servers {
		name = "serv"
		ipv4_address = "1.2.3.4"
		multi_domain_server = "5.5.5.5"
		}
	
}
`, name)
}
