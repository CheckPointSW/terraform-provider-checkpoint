package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementInteroperableDevice_basic(t *testing.T) {

	objName := "tfTestManagementDataInteroperableDevice_" + acctest.RandString(6)
	resourceName := "checkpoint_management_interoperable_device.interoperable_device"
	dataSourceName := "data.checkpoint_management_interoperable_device.data_interoperable_device"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementInteroperableDeviceConfig(objName, "1.1.1.1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "ipv4_address", resourceName, "ipv4_address"),
				),
			},
		},
	})

}

func testAccDataSourceManagementInteroperableDeviceConfig(name string, ipAddress string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_interoperable_device" "interoperable_device" {
    name = "%s"
	ipv4_address = "%s"
}

data "checkpoint_management_interoperable_device" "data_interoperable_device" {
    name = "${checkpoint_management_interoperable_device.interoperable_device.name}"
}
`, name, ipAddress)
}
