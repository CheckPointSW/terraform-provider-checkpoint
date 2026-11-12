package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementThreatProtectionSubCategory_basic(t *testing.T) {

	objName := "tfTestManagementDataThreatProtectionSubCategory_" + acctest.RandString(6)
	resourceName := "checkpoint_management_threat_protection_sub_category.threat_protection_sub_category"
	dataSourceName := "data.checkpoint_management_data_threat_protection_sub_category.data_threat_protection_sub_category"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementThreatProtectionSubCategoryConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementThreatProtectionSubCategoryConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_threat_protection_sub_category" "threat_protection_sub_category" {
    name = "%s"
}

data "checkpoint_management_data_threat_protection_sub_category" "data_threat_protection_sub_category" {
    name = "${checkpoint_management_threat_protection_sub_category.threat_protection_sub_category.name}"
}
`, name)
}
