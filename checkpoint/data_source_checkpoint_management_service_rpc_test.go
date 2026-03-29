package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementServiceRpc_basic(t *testing.T) {

	objName := "tfTestManagementDataServiceRpc_" + acctest.RandString(6)
	resourceName := "checkpoint_management_service_rpc.service_rpc"
	dataSourceName := "data.checkpoint_management_data_service_rpc.data_service_rpc"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementServiceRpcConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementServiceRpcConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_service_rpc" "service_rpc" {
    name = "%s"
}

data "checkpoint_management_data_service_rpc" "data_service_rpc" {
    name = "${checkpoint_management_service_rpc.service_rpc.name}"
}
`, name)
}
