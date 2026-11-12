package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccCheckpointManagementLoginToDomain_basic(t *testing.T) {

	commandName := "checkpoint_management_command_login_to_domain.login_to_domain"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementLoginToDomainConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(commandName, "", commandName, ""),
				),
			},
		},
	})

}

func testAccManagementLoginToDomainConfig() string {
	return fmt.Sprintf(`
resource "checkpoint_management_command_login_to_domain" "login_to_domain" {
  domain = "System Data"
}
`)
}
