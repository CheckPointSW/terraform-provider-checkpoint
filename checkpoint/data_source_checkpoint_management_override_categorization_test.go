package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementOverrideCategorization_basic(t *testing.T) {

	//resourceName := "checkpoint_management_override_categorization.test"
	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementOverrideCategorizationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementOverrideCategorizationConfig("ramdomUrl", "Botnets", "low"),
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
	})
}

func testAccDataSourceManagementOverrideCategorizationConfig(url string, newPrimaryCategory string, risk string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_override_categorization" "test" {
        url = "%s"
        new_primary_category = "%s"
        risk = "%s"
}

data "checkpoint_management_override_categorization" "data" {
  url = "${checkpoint_management_override_categorization.test.url}"
}
`, url, newPrimaryCategory, risk)
}
