package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAcDataSourcecCheckpointManagementDataTypeFileAttributes_basic(t *testing.T) {

	resourceName := "checkpoint_management_data_type_file_attributes.test"
	dataSourceName := "data.checkpoint_management_data_type_file_attributes.data_test"

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
				Config: testAccDataSourceManagementDataTypeFileAttributesConfig("objname", "keywords object"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})
}

func testAccDataSourceManagementDataTypeFileAttributesConfig(name string, desc string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_data_type_file_attributes" "test" {
        name = "%s"
        description = "%s"
         match_by_file_name = "true"
        file_name_contains = "r^"
        match_by_file_size = "true"
        file_size = "12"
}
data "checkpoint_management_data_type_file_attributes" "data_test" {
    name = "${checkpoint_management_data_type_file_attributes.test.name}"
}
`, name, desc)
}
