package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"os"
	"testing"
)

func TestAccCheckpointManagementSetInterTrustedCa_basic(t *testing.T) {

	commandName := "checkpoint_management_command_set_internal_trusted_ca.set_internal_ca"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementSetInternalTrustedCaConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(commandName, "inspect_archives", commandName, "inspect_archives"),
				),
			},
		},
	})

}

func testAccManagementSetInternalTrustedCaConfig() string {
	return fmt.Sprintf(`
resource "checkpoint_management_command_set_internal_trusted_ca" "set_internal_ca" {

  cache_crl = "false"
  crl_cache_timeout = 1200
  branches = [
    "branch1"
  ]
}`)
}
