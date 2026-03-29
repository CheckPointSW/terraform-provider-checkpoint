package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementGsnHandoverGroup_basic(t *testing.T) {
	objName := "tfTestManagementGsnHandoverGroup_" + acctest.RandString(6)
	resourceName := "checkpoint_management_gsn_handover_group.test"
	dataSourceName := "data.checkpoint_management_gsn_handover_group.data_test"

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
				Config: testAccDataSourceManagementGsnHandoverGroupConfig(objName, true, 2048, "All_Internet"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})
}

func testAccDataSourceManagementGsnHandoverGroupConfig(name string, enforceGtp bool, gtpRate int, members1 string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_gsn_handover_group" "test" {
        name = "%s"
        enforce_gtp = %t
        gtp_rate = %d
        members = ["%s"]
}

data "checkpoint_management_gsn_handover_group" "data_test" {
        name = "${checkpoint_management_gsn_handover_group.test.name}"
}
`, name, enforceGtp, gtpRate, members1)
}
