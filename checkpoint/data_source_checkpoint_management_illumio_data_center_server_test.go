package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementIllumioDataCenterServer_basic(t *testing.T) {

	objName := "tfTestManagementDataIllumioDataCenterServer_" + acctest.RandString(6)
	resourceName := "checkpoint_management_illumio_data_center_server.illumio_data_center_server"
	dataSourceName := "data.checkpoint_management_illumio_data_center_server.illumio_data_center_server"
	hostname := "hostname.illum.io"
	orgId := 1234567
	authUsername := "api_6e8d5249c27d64185"
	secret := "e7f5b0e8-4f5e-a17d-8f9e-0d2e5bdsdfe1"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementIllumioDataCenterServerConfig(objName, hostname, orgId, authUsername, secret),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementIllumioDataCenterServerConfig(name string, hostname string, orgId int, authUsername string, secret string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_illumio_data_center_server" "illumio_data_center_server" {
    name = "%s"
	hostname = "%s"
	org_id = %d
	auth_username = "%s"
	secret = "%s"
	ignore_warnings = true
}

data "checkpoint_management_illumio_data_center_server" "illumio_data_center_server" {
    name = "${checkpoint_management_illumio_data_center_server.illumio_data_center_server.name}"
}
`, name, hostname, orgId, authUsername, secret)
}
