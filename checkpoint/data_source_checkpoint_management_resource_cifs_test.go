package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementResourceCifs_basic(t *testing.T) {

	objName := "tfTestManagementDataResouceCifs_" + acctest.RandString(6)
	resourceName := "checkpoint_management_resource_cifs.test"
	dataSourceName := "data.checkpoint_management_resource_cifs.data"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementResourceCifsConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementResourceCifsConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_resource_cifs" "test" {

   name = "%s"
   allowed_disk_and_print_shares {
     server_name = "server1"
     share_name = "share12"
   }

  allowed_disk_and_print_shares {
    server_name = "server3"
    share_name = "share4"
  }
   log_mapped_shares = true
   log_access_violation = true
   block_remote_registry_access = false

}
data "checkpoint_management_resource_cifs" "data" {
  uid = "${checkpoint_management_resource_cifs.test.id}"
}
`, name)
}
