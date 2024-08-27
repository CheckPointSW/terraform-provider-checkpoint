package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementResourceSmtp_basic(t *testing.T) {

	objName := "tfTestManagementDataResouceSmtp_" + acctest.RandString(6)
	resourceName := "checkpoint_management_resource_smtp.test"
	dataSourceName := "data.checkpoint_management_resource_smtp.data"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementResourceSmtpConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementResourceSmtpConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_resource_smtp" "test" {

  name = "%s"
  mail_delivery_server = "deliverServer"
  exception_track = "exception log"
  match = {
    sender = "bob"
    recipient = "liza"
  }
  action_1 {
    sender {
      original = "one"
      rewritten = "two"
    }
    recipient {
      original = "three"
      rewritten = "four"
    }
    custom_field{
      field = "field"
      original = "five"
      rewritten = "six"
    }
  }
}
data "checkpoint_management_resource_smtp" "data" {
  uid = "${checkpoint_management_resource_smtp.test.id}"
}
`, name)
}
