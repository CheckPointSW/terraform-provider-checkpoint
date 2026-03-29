package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementServiceTcp_basic(t *testing.T) {

	objName := "tfTestManagementDataServiceTcp_" + acctest.RandString(6)
	resourceName := "checkpoint_management_service_tcp.service_tcp"
	dataSourceName := "data.checkpoint_management_data_service_tcp.data_service_tcp"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementServiceTcpConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})
}

func testAccDataSourceManagementServiceTcpConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_service_tcp" "service_tcp" {
    name = "%s"
	port = "1122"
}

data "checkpoint_management_data_service_tcp" "data_service_tcp" {
    name = "${checkpoint_management_service_tcp.service_tcp.name}"
}
`, name)
}
