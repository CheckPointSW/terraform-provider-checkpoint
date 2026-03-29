package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementServiceDceRpc_basic(t *testing.T) {

	objName := "tfTestManagementDataServiceDceRpc_" + acctest.RandString(6)
	resourceName := "checkpoint_management_service_dce_rpc.service_dce_rpc"
	dataSourceName := "data.checkpoint_management_data_service_dce_rpc.data_service_dce_rpc"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementServiceDceRpcConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementServiceDceRpcConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_service_dce_rpc" "service_dce_rpc" {
    name = "%s"
    interface_uuid = "97aeb460-9aea-11d5-bd16-0090272ccb30"
}

data "checkpoint_management_data_service_dce_rpc" "data_service_dce_rpc" {
    name = "${checkpoint_management_service_dce_rpc.service_dce_rpc.name}"
}
`, name)
}
