package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementResourceFtp_basic(t *testing.T) {

	objName := "tfTestManagementDataResouceFtp_" + acctest.RandString(6)
	resourceName := "checkpoint_management_resource_ftp.test"
	dataSourceName := "data.checkpoint_management_resource_ftp.data"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementResourceFtpConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementResourceFtpConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_resource_ftp" "test" {
  name = "%s"
  resource_matching_method = "get_and_put"
  resources_path = "path"
}
data "checkpoint_management_resource_ftp" "data" {
  uid = "${checkpoint_management_resource_ftp.test.id}"
}
`, name)
}
