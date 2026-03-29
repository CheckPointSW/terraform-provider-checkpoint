package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementPasscodeProfile_basic(t *testing.T) {

	resourceName := "checkpoint_management_passcode_profile.test"
	dataSourceName := "data.checkpoint_management_passcode_profile.data"
	objName := "tfTestManagementPasscodeProfile_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementPasscodeProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementPasscodeProfileConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name")),
			},
		},
	})
}

func testAccDataSourceManagementPasscodeProfileConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_passcode_profile" "test" {
        name = "%s"
}
data "checkpoint_management_passcode_profile" "data" {
  uid = "${checkpoint_management_passcode_profile.test.id}"
}
`, name)
}
