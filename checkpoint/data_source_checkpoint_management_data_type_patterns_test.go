package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAcDataSourcecCheckpointManagementDataTypePatterns_basic(t *testing.T) {

	resourceName := "checkpoint_management_data_type_patterns.test"
	dataSourceName := "data.checkpoint_management_data_type_patterns.data_test"

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
				Config: testAccDataSourceManagementDataTypePatternsConfig("objname", "keywords object"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})
}

func testAccDataSourceManagementDataTypePatternsConfig(name string, desc string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_data_type_patterns" "test" {
        name = "%s"
        description = "%s"
        patterns = [ "aaa*" , "r^" ]
}
data "checkpoint_management_data_type_patterns" "data_test" {
    name = "${checkpoint_management_data_type_patterns.test.name}"
}
`, name, desc)
}
