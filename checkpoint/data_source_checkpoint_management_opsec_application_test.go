package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementOpsecApplication_basic(t *testing.T) {

	objName := "tfTestManagementDataOpsecApplication_" + acctest.RandString(6)
	resourceName := "checkpoint_management_opsec_application.opsec_application"
	dataSourceName := "data.checkpoint_management_data_opsec_application.data_opsec_application"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementOpsecApplicationConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementOpsecApplicationConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_host" "myhost1" {
    name = "myhost1"
    ipv4_address = "1.2.3.5" 
}

resource "checkpoint_management_opsec_application" "opsec_application" {
    name = "%s"
    host = "${checkpoint_management_host.myhost1.name}"
    cpmi = {
        enabled = true
        administrator_profile = "read only all"
        use_administrator_credentials = false
    }
    lea = {
        enabled = true
        access_permissions = "show all"
    }
}

data "checkpoint_management_data_opsec_application" "data_opsec_application" {
    name = "${checkpoint_management_opsec_application.opsec_application.name}"
}
`, name)
}
