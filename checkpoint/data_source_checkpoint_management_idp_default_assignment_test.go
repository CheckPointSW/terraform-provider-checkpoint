package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementIdpDefaultAssignment_basic(t *testing.T) {

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementIdpDefaultAssignmentConfig(),
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
	})

}

func testAccDataSourceManagementIdpDefaultAssignmentConfig() string {
	return fmt.Sprintf(`
data "checkpoint_management_idp_default_assignment" "data_idp_default_assignment" {
}
`)
}
