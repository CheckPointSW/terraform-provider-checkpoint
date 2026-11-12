package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementCpPasswordRequirements_basic(t *testing.T) {
	resourceName := "checkpoint_management_set_cp_password_requirements.test"
	dataSourceName := "data.checkpoint_management_cp_password_requirements.data_test"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementCpPasswordRequirementsConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "min_password_length", resourceName, "min_password_length"),
				),
			},
		},
	})
}

func testAccDataSourceManagementCpPasswordRequirementsConfig() string {
	return fmt.Sprintf(`
resource "checkpoint_management_set_cp_password_requirements" "test" {
	min_password_length = 7
}

data "checkpoint_management_cp_password_requirements" "data_test" {
}
`)
}
