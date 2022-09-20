package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementIpsProtectionExtendedAttribute_basic(t *testing.T) {

	dataSourceName := "data.checkpoint_management_ips_protection_extended_attribute.ips"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementIpsProtectionExtendedAttributeConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", dataSourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementIpsProtectionExtendedAttributeConfig() string {
	return fmt.Sprintf(`
data "checkpoint_management_ips_protection_extended_attribute" "ips" {
    name = "File Type"
	uid = "1fc86799-672b-453e-876b-d47651cbc70e"
}
`)
}
