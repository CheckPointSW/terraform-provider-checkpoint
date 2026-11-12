package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementDefaultAdministratorSettings_basic(t *testing.T) {
	resourceName := "checkpoint_management_set_default_administrator_settings.test"
	dataSourceName := "data.checkpoint_management_default_administrator_settings.data_test"

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
				Config: testAccDataSourceManagementDefaultAdministratorSettingsConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "expiration_type", resourceName, "expiration_type"),
					resource.TestCheckResourceAttrPair(dataSourceName, "expiration_date", resourceName, "expiration_date"),
					resource.TestCheckResourceAttrPair(dataSourceName, "indicate_expiration_in_admin_view", resourceName, "indicate_expiration_in_admin_view"),
					resource.TestCheckResourceAttrPair(dataSourceName, "notify_expiration_to_admin", resourceName, "notify_expiration_to_admin"),
					resource.TestCheckResourceAttrPair(dataSourceName, "days_to_notify_expiration_to_admin", resourceName, "days_to_notify_expiration_to_admin"),
				),
			},
		},
	})
}

func testAccDataSourceManagementDefaultAdministratorSettingsConfig() string {
	return fmt.Sprintf(`
resource "checkpoint_management_set_default_administrator_settings" "test" {
  expiration_type = "expiration date"
  expiration_date = "2025-06-23"
  indicate_expiration_in_admin_view = false
  notify_expiration_to_admin = true
  days_to_notify_expiration_to_admin = 5
}

data "checkpoint_management_default_administrator_settings" "data_test" {
}
`)
}
