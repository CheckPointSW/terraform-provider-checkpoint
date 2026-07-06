package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementDlpNextDataType_basic(t *testing.T) {

	objName := "tfTestManagementDataDlpNextDataType_" + acctest.RandString(6)
	resourceName := "checkpoint_management_dlp_next_data_type.dlp_next_data_type"
	dataSourceName := "data.checkpoint_management_data_dlp_next_data_type.data_dlp_next_data_type"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementDlpNextDataTypeConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementDlpNextDataTypeConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_dlp_next_data_type" "dlp_next_data_type" {
    name = "%s"
	description = "description_test"
}

data "checkpoint_management_data_dlp_next_data_type" "data_dlp_next_data_type" {
    name = "${checkpoint_management_dlp_next_data_type.dlp_next_data_type.name}"
}
`, name)
}
