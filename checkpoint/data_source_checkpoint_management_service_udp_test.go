package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementServiceUdp_basic(t *testing.T) {

	objName := "tfTestManagementDataServiceUdp_" + acctest.RandString(6)
	resourceName := "checkpoint_management_service_udp.service_udp"
	dataSourceName := "data.checkpoint_management_data_service_udp.data_service_udp"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementServiceUdpConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})
}

func testAccDataSourceManagementServiceUdpConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_service_udp" "service_udp" {
    name = "%s"
	port = "1123"
}

data "checkpoint_management_data_service_udp" "data_service_udp" {
    name = "${checkpoint_management_service_udp.service_udp.name}"
}
`, name)
}
