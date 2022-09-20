package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementAutomaticPurge_basic(t *testing.T) {

	dataSourceName := "data.checkpoint_management_automatic_purge.auto"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementAutomaticPurgeConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", dataSourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementAutomaticPurgeConfig() string {
	return fmt.Sprintf(`
data "checkpoint_management_automatic_purge" "auto" {

}
`)
}
