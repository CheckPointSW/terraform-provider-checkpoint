package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementThreatEmulationFileType_basic(t *testing.T) {

	objName := "tfTestManagementDataThreatEmulationFileType_" + acctest.RandString(6)
	resourceName := "checkpoint_management_threat_emulation_file_type.threat_emulation_file_type"
	dataSourceName := "data.checkpoint_management_data_threat_emulation_file_type.data_threat_emulation_file_type"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementThreatEmulationFileTypeConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementThreatEmulationFileTypeConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_threat_emulation_file_type" "threat_emulation_file_type" {

    name = "%s"
	type = "type_test"
	description = "description_test"
}

data "checkpoint_management_data_threat_emulation_file_type" "data_threat_emulation_file_type" {
    name = "${checkpoint_management_threat_emulation_file_type.threat_emulation_file_type.name}"
}
`, name)
}
