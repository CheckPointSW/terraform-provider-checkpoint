package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementCMEAPIVersions_basic(t *testing.T) {
	dataSourceName := "data.checkpoint_management_cme_api_versions.test"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this test")
	} else if context != "web_api" {
		t.Skip("Skipping cme api test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementCMEAPIVersionsConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceName, "current_version"),
					resource.TestCheckResourceAttrSet(dataSourceName, "supported_versions.#"),
				),
			},
		},
	})
}

func testAccDataSourceManagementCMEAPIVersionsConfig() string {
	return fmt.Sprintf(`
data "checkpoint_management_cme_api_versions" "test"{
}
`)
}
