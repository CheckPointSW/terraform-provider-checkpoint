package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementAccessRule_basic(t *testing.T) {

	objName := "tfTestManagementDataAccessRule_" + acctest.RandString(6)
	resourceName := "checkpoint_management_access_rule.access_rule"
	dataSourceName := "data.checkpoint_management_data_access_rule.data_access_rule"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementAccessRuleConfig(objName, "Network"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementAccessRuleConfig(name string, layer string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_access_rule" "access_rule" {
    name = "%s"
	layer = "%s"
	position = {top = "top"}
	source = ["Any"]
	destination = ["Any"]
	service = ["Any"]
	track = {
    accounting = false
    alert = "none"
    enable_firewall_session = false
    per_connection = false
    per_session = false
    type = "None"
  }
}

data "checkpoint_management_data_access_rule" "data_access_rule" {
    name = "${checkpoint_management_access_rule.access_rule.name}"
    layer = "${checkpoint_management_access_rule.access_rule.layer}"
}
`, name, layer)
}
