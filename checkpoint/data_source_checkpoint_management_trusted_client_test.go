package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementTrustedClient_basic(t *testing.T) {

	objName := "tfTestManagementDataTrustedClient_" + acctest.RandString(6)
	resourceName := "checkpoint_management_trusted_client.trustedClient"
	dataSourceName := "data.checkpoint_management_trusted_client.data_trustedClient"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementTrustedClientConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})
}

func testAccDataSourceManagementTrustedClientConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_trusted_client" "trustedClient" {
    name = "%s"
	ipv4_address = "192.168.2.1"
}

data "checkpoint_management_trusted_client" "data_trusted_client" {
    name = "${checkpoint_management_trusted_client.trustedClient.name}"
}
`, name)
}
