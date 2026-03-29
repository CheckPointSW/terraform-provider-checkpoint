package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementUserTemplate_basic(t *testing.T) {

	objName := "UserTemplate" + acctest.RandString(2)
	resourceName := "checkpoint_management_user_template.user_template"
	dataSourceName := "data.checkpoint_management_user_template.test_user_template"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementUserTemplateConfig(objName, "2030-05-30", false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})
}

func testAccDataSourceManagementUserTemplateConfig(name string, expirationDate string, expirationByGlobalProperties bool) string {
	return fmt.Sprintf(`
resource "checkpoint_management_user_template" "user_template" {
      name = "%s"
      expiration_date = "%s"
      expiration_by_global_properties = %t
}

data "checkpoint_management_user_template" "test_user_template" {
    name = "${checkpoint_management_user_template.user_template.name}"
}
`, name, expirationDate, expirationByGlobalProperties)
}
