package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementNetwork_basic(t *testing.T) {

	objName := "tfTestManagementDataNetwork_" + acctest.RandString(6)
	resourceName := "checkpoint_management_network.network"
	dataSourceName := "data.checkpoint_management_data_network.data_network"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementNetworkConfig(objName, "10.0.0.0", 24),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "subnet4", resourceName, "subnet4"),
					resource.TestCheckResourceAttrPair(dataSourceName, "mask_length4", resourceName, "mask_length4"),
				),
			},
		},
	})

}

func testAccDataSourceManagementNetworkConfig(name string, subnet4 string, masklen4 int) string {
	return fmt.Sprintf(`
resource "checkpoint_management_network" "network" {
    name = "%s"
	subnet4 = "%s"
	mask_length4 = "%d"
}

data "checkpoint_management_data_network" "data_network" {
    name = "${checkpoint_management_network.network.name}"
}
`, name, subnet4, masklen4)
}
