package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementServiceCompoundTcp_basic(t *testing.T) {

	objName := "tfTestManagementDataServiceCompoundTcp_" + acctest.RandString(6)
	resourceName := "checkpoint_management_service_compound_tcp.service_compound_tcp"
	dataSourceName := "data.checkpoint_management_service_compound_tcp.test"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementServiceCompoundTcpConfig(objName, "pointcast", true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementServiceCompoundTcpConfig(name string, compoundService string, keepConnectionsOpenAfterPolicyInstallation bool) string {
	return fmt.Sprintf(`
resource "checkpoint_management_service_compound_tcp" "service_compound_tcp" {
        name = "%s"
        compound_service = "%s"
        keep_connections_open_after_policy_installation = %t
}

data "checkpoint_management_service_compound_tcp" "test" {
    name = "${checkpoint_management_service_compound_tcp.service_compound_tcp.name}"
}
`, name, compoundService, keepConnectionsOpenAfterPolicyInstallation)
}
