package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementServiceIcmp6_basic(t *testing.T) {

	objName := "tfTestManagementDataServiceIcmp6_" + acctest.RandString(6)
	resourceName := "checkpoint_management_service_icmp6.service_icmp6"
	dataSourceName := "data.checkpoint_management_data_service_icmp6.data_service_icmp6"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementServiceIcmp6Config(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementServiceIcmp6Config(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_service_icmp6" "service_icmp6" {
    name = "%s"
}

data "checkpoint_management_data_service_icmp6" "data_service_icmp6" {
    name = "${checkpoint_management_service_icmp6.service_icmp6.name}"
}
`, name)
}
