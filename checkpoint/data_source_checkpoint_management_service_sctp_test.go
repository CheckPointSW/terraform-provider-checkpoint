package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementServiceSctp_basic(t *testing.T) {

	objName := "tfTestManagementDataServiceSctp_" + acctest.RandString(6)
	resourceName := "checkpoint_management_service_sctp.service_sctp"
	dataSourceName := "data.checkpoint_management_data_service_sctp.data_service_sctp"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementServiceSctpConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementServiceSctpConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_service_sctp" "service_sctp" {
    name = "%s"
    port = "1234"
    session_timeout = "3600"
    sync_connections_on_cluster = true
}

data "checkpoint_management_data_service_sctp" "data_service_sctp" {
    name = "${checkpoint_management_service_sctp.service_sctp.name}"
}
`, name)
}
