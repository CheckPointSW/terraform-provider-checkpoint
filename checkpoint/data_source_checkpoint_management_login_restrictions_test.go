package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementLoginRestrictions_basic(t *testing.T) {

	dataSourceName := "data.checkpoint_management_login_restrictions.data"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementLoginRestrictionsConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceName, "uid"),
				),
			},
		},
	})

}

func testAccDataSourceManagementLoginRestrictionsConfig() string {
	return fmt.Sprintf(`
data "checkpoint_management_login_restrictions" "data" {

}
`)
}
