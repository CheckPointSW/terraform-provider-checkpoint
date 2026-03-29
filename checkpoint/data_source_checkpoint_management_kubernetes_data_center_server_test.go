package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementKubernetesDataCenterServer_basic(t *testing.T) {

	objName := "tfTestManagementDataKubernetesDataCenterServer_" + acctest.RandString(6)
	resourceName := "checkpoint_management_kubernetes_data_center_server.kubernetes_data_center_server"
	dataSourceName := "data.checkpoint_management_kubernetes_data_center_server.kubernetes_data_center_server"
	hostname := "MY_HOSTNAME"
	token_file := "MY_TOKEN"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementKubernetesDataCenterServerConfig(objName, hostname, token_file),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementKubernetesDataCenterServerConfig(name string, hostname string, token_file string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_kubernetes_data_center_server" "kubernetes_data_center_server" {
    name = "%s"
	hostname = "%s"
	token_file = "%s"
    unsafe_auto_accept = true
	ignore_warnings = true
}

data "checkpoint_management_kubernetes_data_center_server" "kubernetes_data_center_server" {
    name = "${checkpoint_management_kubernetes_data_center_server.kubernetes_data_center_server.name}"
}
`, name, hostname, token_file)
}
