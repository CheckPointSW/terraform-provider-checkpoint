package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementHttpsRule_basic(t *testing.T) {

	objName := "tfTestManagementDataHttpsRule_" + acctest.RandString(6)
	resourceName := "checkpoint_management_https_rule.https_rule"
	dataSourceName := "data.checkpoint_management_data_https_rule.data_https_rule"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementHttpsRuleConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementHttpsRuleConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_https_rule" "https_rule" {
	name = "%s"
	position = {top = "top"}
	layer = "Default Layer"
	blade = ["IPS"]
	destination = ["Internet"]
	enabled = true
	service = ["HTTPS default services"]
	source = ["DMZNet"]
	install_on = ["Policy HTTPS Targets"]
	site_category = ["Any"]
}

data "checkpoint_management_data_https_rule" "data_https_rule" {
    rule_number = "1"
    layer = "${checkpoint_management_https_rule.https_rule.layer}"
}
`, name)
}
