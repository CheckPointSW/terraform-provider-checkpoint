package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementMds_basic(t *testing.T) {
	objName := "tfTestManagementMds_" + acctest.RandString(6)
	resourceName := "checkpoint_management_mds.mds"
	dataSourceName := "data.checkpoint_management_mds.data_mds"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementMdsConfig(objName, "5.6.7.8"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "ipv4_address", resourceName, "ipv4_address"),
				),
			},
		},
	})
}

func testAccDataSourceManagementMdsConfig(name string, ipv4 string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_mds" "mds" {
        name = "%s"
        ipv4_address = "%s"
}

data "checkpoint_management_mds" "data_mds" {
        name = "${checkpoint_management_mds.mds.name}"
}
`, name, ipv4)
}
