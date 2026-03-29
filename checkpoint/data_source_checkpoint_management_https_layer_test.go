package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementHttpsLayer_basic(t *testing.T) {

	objName := "tfTestManagementDataHttpsLayer_" + acctest.RandString(6)
	resourceName := "checkpoint_management_https_layer.https_layer"
	dataSourceName := "data.checkpoint_management_data_https_layer.data_https_layer"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementHttpsLayerConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementHttpsLayerConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_https_layer" "https_layer" {
    name = "%s"
}

data "checkpoint_management_data_https_layer" "data_https_layer" {
    name = "${checkpoint_management_https_layer.https_layer.name}"
}
`, name)
}
