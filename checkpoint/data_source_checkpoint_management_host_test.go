package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementHost_basic(t *testing.T) {

	objName := "tfTestManagementDataHost_" + acctest.RandString(6)
	resourceName := "checkpoint_management_host.host"
	dataSourceName := "data.checkpoint_management_data_host.data_host"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementHostConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "ipv4_address", resourceName, "ipv4_address"),
				),
			},
		},
	})
}

func testAccDataSourceManagementHostConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_host" "host" {
    name = "%s"
    ipv4_address = "1.2.3.4"
}

data "checkpoint_management_data_host" "data_host" {
    name = "${checkpoint_management_host.host.name}"
}
`, name)
}
