package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementServiceGroup_basic(t *testing.T) {

	objName := "tfTestManagementDataServiceGroup_" + acctest.RandString(6)
	resourceName := "checkpoint_management_service_group.service_group"
	dataSourceName := "data.checkpoint_management_data_service_group.data_service_group"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementServiceGroupConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementServiceGroupConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_service_group" "service_group" {
    name = "%s"
}

data "checkpoint_management_data_service_group" "data_service_group" {
    name = "${checkpoint_management_service_group.service_group.name}"
}
`, name)
}
