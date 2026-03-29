package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementThreatLayer_basic(t *testing.T) {

	objName := "tfTestManagementDataThreatLayer_" + acctest.RandString(6)
	resourceName := "checkpoint_management_threat_layer.threat_layer"
	dataSourceName := "data.checkpoint_management_threat_layer.data_threat_layer"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementThreatLayerConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementThreatLayerConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_threat_layer" "threat_layer" {
    name = "%s"
	color = "blue"
}

data "checkpoint_management_threat_layer" "data_threat_layer" {
    name = "${checkpoint_management_threat_layer.threat_layer.name}"
}
`, name)
}
