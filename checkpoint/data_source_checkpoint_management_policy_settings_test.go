package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementPolicySettings_basic(t *testing.T) {

	dataSourceName := "data.checkpoint_management_policy_settings.pl"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementPolicySettingsConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceName, "last_in_cell"),
				),
			},
		},
	})

}

func testAccDataSourceManagementPolicySettingsConfig() string {
	return fmt.Sprintf(`
data "checkpoint_management_policy_settings" "pl" {

}
`)
}
