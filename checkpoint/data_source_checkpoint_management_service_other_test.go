package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementServiceOther_basic(t *testing.T) {

	objName := "tfTestManagementDataServiceOther_" + acctest.RandString(6)
	resourceName := "checkpoint_management_service_other.service_other"
	dataSourceName := "data.checkpoint_management_data_service_other.data_service_other"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementServiceOtherConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementServiceOtherConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_service_other" "service_other" {
    name = "%s"
    keep_connections_open_after_policy_installation = false
	session_timeout = 100
	match_for_any = true
	sync_connections_on_cluster = true
	ip_protocol = 51
	aggressive_aging = {
		use_default_timeout = true
		enable = true
		default_timeout = 600
		timeout = 600
	}
}

data "checkpoint_management_data_service_other" "data_service_other" {
    name = "${checkpoint_management_service_other.service_other.name}"
}
`, name)
}
