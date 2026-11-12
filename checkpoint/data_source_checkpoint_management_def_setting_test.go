package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementDefSetting_basic(t *testing.T) {

	objName := "tfTestManagementDataDefSetting_" + acctest.RandString(6)
	resourceName := "checkpoint_management_def_setting.def_setting"
	dataSourceName := "data.checkpoint_management_data_def_setting.data_def_setting"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementDefSettingConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementDefSettingConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_def_setting" "example" {
  name      = "%s"
  data_type = "boolean"
  assignments {
    value       = "true"
    description = "Default for Quantum gateways"
    model       = "quantum"
  }
  assignments {
    value       = "false"
    description = "Default for Spark gateways"
    model       = "spark"
  }
}

data "checkpoint_management_data_def_setting" "data_def_setting" {
    name = "${checkpoint_management_def_setting.def_setting.name}"
}
`, name)
}
