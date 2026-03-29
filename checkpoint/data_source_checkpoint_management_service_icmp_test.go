package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementServiceIcmp_basic(t *testing.T) {

	objName := "tfTestManagementDataServiceIcmp_" + acctest.RandString(6)
	resourceName := "checkpoint_management_service_icmp.service_icmp"
	dataSourceName := "data.checkpoint_management_data_service_icmp.data_service_icmp"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementServiceIcmpConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementServiceIcmpConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_service_icmp" "service_icmp" {
    name = "%s"
}

data "checkpoint_management_data_service_icmp" "data_service_icmp" {
    name = "${checkpoint_management_service_icmp.service_icmp.name}"
}
`, name)
}
