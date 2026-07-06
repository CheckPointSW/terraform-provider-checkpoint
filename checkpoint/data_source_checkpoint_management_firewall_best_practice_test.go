package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementFirewallBestPractice_basic(t *testing.T) {

	objName := "tfTestManagementDataFirewallBestPractice_" + acctest.RandString(6)
	resourceName := "checkpoint_management_firewall_best_practice.firewall_best_practice"
	dataSourceName := "data.checkpoint_management_data_firewall_best_practice.data_firewall_best_practice"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementFirewallBestPracticeConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementFirewallBestPracticeConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_firewall_best_practice" "test" {
        name = "%s"
        action_item = "test"
        description = "test"
        enabled = true
}

data "checkpoint_management_data_firewall_best_practice" "data_firewall_best_practice" {
    name = "${checkpoint_management_firewall_best_practice.firewall_best_practice.name}"
}
`, name)
}
