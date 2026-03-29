package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementWildcard_basic(t *testing.T) {

	objName := "tfTestManagementDataWildcard_" + acctest.RandString(6)
	resourceName := "checkpoint_management_wildcard.wildcard"
	dataSourceName := "data.checkpoint_management_data_wildcard.data_wildcard"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementWildcardConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})
}

func testAccDataSourceManagementWildcardConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_wildcard" "wildcard" {
    name = "%s"
	ipv4_address = "192.168.2.1"
 	ipv4_mask_wildcard = "0.0.0.128"
}

data "checkpoint_management_data_wildcard" "data_wildcard" {
    name = "${checkpoint_management_wildcard.wildcard.name}"
}
`, name)
}
