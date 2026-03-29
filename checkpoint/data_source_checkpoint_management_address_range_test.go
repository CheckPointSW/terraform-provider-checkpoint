package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementAddressRange_basic(t *testing.T) {

	objName := "tfTestManagementDataAddressRange_" + acctest.RandString(6)
	resourceName := "checkpoint_management_address_range.address_range"
	dataSourceName := "data.checkpoint_management_data_address_range.data_address_range"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementAddressRangeConfig(objName, "1.1.1.1", "2.2.2.2"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "ipv4_address_first", resourceName, "ipv4_address_first"),
					resource.TestCheckResourceAttrPair(dataSourceName, "ipv4_address_last", resourceName, "ipv4_address_last"),
				),
			},
		},
	})

}

func testAccDataSourceManagementAddressRangeConfig(name string, ipv4First string, ipv4Last string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_address_range" "address_range" {
    name = "%s"
    ipv4_address_first = "%s"
    ipv4_address_last = "%s"
}

data "checkpoint_management_data_address_range" "data_address_range" {
    name = "${checkpoint_management_address_range.address_range.name}"
}
`, name, ipv4First, ipv4Last)
}
