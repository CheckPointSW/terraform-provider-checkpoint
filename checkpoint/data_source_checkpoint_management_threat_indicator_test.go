package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementThreatIndicator_basic(t *testing.T) {

	objName := "tfTestManagementDataThreatIndicator_" + acctest.RandString(6)
	resourceName := "checkpoint_management_threat_indicator.threat_indicator"
	dataSourceName := "data.checkpoint_management_data_threat_indicator.data_threat_indicator"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementThreatIndicatorConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})
}

func testAccDataSourceManagementThreatIndicatorConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_threat_indicator" "threat_indicator" {
    name = "%s"
	observables {
    	name = "obs1"
    	ip_address = "5.4.7.1"
  	}
	ignore_warnings = true
}

data "checkpoint_management_data_threat_indicator" "data_threat_indicator" {
    name = "${checkpoint_management_threat_indicator.threat_indicator.name}"
}
`, name)
}
