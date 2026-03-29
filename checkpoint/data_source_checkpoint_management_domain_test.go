package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementDomain_basic(t *testing.T) {
	objName := "tfTestManagementDomain_" + acctest.RandString(6)
	resourceName := "checkpoint_management_domain.test"
	dataSourceName := "data.checkpoint_management_domain.data_test"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementDomainConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})
}

func testAccDataSourceManagementDomainConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_domain" "test" {
        name = "%s"
        servers {
		name = "serv"
		ipv4_address = "1.2.3.4"
		multi_domain_server = "5.5.5.5"
		}
	
}

data "checkpoint_management_domain" "data_test" {
        name = "${checkpoint_management_domain.test.name}"
}
`, name)
}
