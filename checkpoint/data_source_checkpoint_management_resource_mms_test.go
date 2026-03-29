package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementMms_basic(t *testing.T) {

	objName := "tfTestManagementDataMms_" + acctest.RandString(6)
	resourceName := "checkpoint_management_resource_mms.test"
	dataSourceName := "data.checkpoint_management_resource_mms.data_mms"
	track := "log"
	action := "drop"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementMmsConfig(objName, track, action),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "track", resourceName, "track"),
					resource.TestCheckResourceAttrPair(dataSourceName, "action", resourceName, "action"),
				),
			},
		},
	})

}

func testAccDataSourceManagementMmsConfig(name string, track string, action string) string {
	return fmt.Sprintf(`

resource "checkpoint_management_resource_mms" "test" {
	name = "%s"
	track = "%s"
	action = "%s"
}

data "checkpoint_management_resource_mms" "data_mms" {
	name = "${checkpoint_management_resource_mms.test.name}"
}
`, name, track, action)
}
