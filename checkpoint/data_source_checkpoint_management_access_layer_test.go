package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementAccessLayer_basic(t *testing.T) {

	objName := "tfTestManagementDataAccessLayer_" + acctest.RandString(6)
	resourceName := "checkpoint_management_access_layer.access_layer"
	dataSourceName := "data.checkpoint_management_data_access_layer.data_access_layer"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementAccessLayerConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementAccessLayerConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_access_layer" "access_layer" {
    name = "%s"
	detect_using_x_forward_for = false
	applications_and_url_filtering = true
}

data "checkpoint_management_data_access_layer" "data_access_layer" {
    name = "${checkpoint_management_access_layer.access_layer.name}"
}
`, name)
}
