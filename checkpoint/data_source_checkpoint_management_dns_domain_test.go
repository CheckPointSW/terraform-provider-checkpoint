package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementDnsDomain_basic(t *testing.T) {

	objName := ".tfTestManagementDataDnsDomain_" + acctest.RandString(6)
	resourceName := "checkpoint_management_dns_domain.dns_domain"
	dataSourceName := "data.checkpoint_management_data_dns_domain.data_dns_domain"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementDnsDomainConfig(objName, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementDnsDomainConfig(name string, isSubDomain bool) string {
	return fmt.Sprintf(`
resource "checkpoint_management_dns_domain" "dns_domain" {
        name = "%s"
		is_sub_domain = %t
}

data "checkpoint_management_data_dns_domain" "data_dns_domain" {
    name = "${checkpoint_management_dns_domain.dns_domain.name}"
}
`, name, isSubDomain)
}
