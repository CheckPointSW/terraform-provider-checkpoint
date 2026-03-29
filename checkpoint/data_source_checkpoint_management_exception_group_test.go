package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementExceptionGroup_basic(t *testing.T) {

	objName := "tfTestManagementDataHost_" + acctest.RandString(6)
	resourceName := "checkpoint_management_exception_group.exception_group"
	dataSourceName := "data.checkpoint_management_data_exception_group.data_exception_group"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementExceptionGroupConfig(objName, "manually-select-threat-rules"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "apply_on", resourceName, "apply_on"),
				),
			},
		},
	})
}

func testAccDataSourceManagementExceptionGroupConfig(name string, applyOn string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_exception_group" "exception_group" {
    name = "%s"
	apply_on = "%s"
}

data "checkpoint_management_data_exception_group" "data_exception_group" {
    name = "${checkpoint_management_exception_group.exception_group.name}"
}
`, name, applyOn)
}
