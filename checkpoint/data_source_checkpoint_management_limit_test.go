package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementLimit_basic(t *testing.T) {

	objName := "tfTestManagementLimit_" + acctest.RandString(6)
	resourceName := "checkpoint_management_limit.test"
	dataSourceName := "data.checkpoint_management_limit.data"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementLimitDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementLimitConfig(objName, true, "gbps", 4),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})
}

func testAccDataSourceManagementLimitConfig(name string, enableDownload bool, downloadUnit string, downloadRate int) string {
	return fmt.Sprintf(`
resource "checkpoint_management_limit" "test" {
        name = "%s"
        enable_download = %t
        download_unit = "%s"
        download_rate = %d
}
data "checkpoint_management_limit" "data" {
     name = "${checkpoint_management_limit.test.name}"
}
`, name, enableDownload, downloadUnit, downloadRate)
}
