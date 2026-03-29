package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementServiceCitrixTcp_basic(t *testing.T) {

	objName := "tfTestManagementDataServiceCitrixTcp_" + acctest.RandString(6)
	resourceName := "checkpoint_management_service_citrix_tcp.service_citrix_tcp"
	dataSourceName := "data.checkpoint_management_service_citrix_tcp.test"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementServiceCitrixTcpConfig(objName, "my citrix application"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "application", resourceName, "application"),
				),
			},
		},
	})

}

func testAccDataSourceManagementServiceCitrixTcpConfig(name string, application string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_service_citrix_tcp" "service_citrix_tcp" {
     name = "%s"
     application = "%s"
}

data "checkpoint_management_service_citrix_tcp" "test" {
    name = "${checkpoint_management_service_citrix_tcp.service_citrix_tcp.name}"
}
`, name, application)
}
