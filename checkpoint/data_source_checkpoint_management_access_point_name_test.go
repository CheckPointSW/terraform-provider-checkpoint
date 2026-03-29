package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementAccessPointName_basic(t *testing.T) {
	objName := "tfTestManagementDataAccessPointName_" + acctest.RandString(6)
	resourceName := "checkpoint_management_access_point_name.access_point_name"
	dataSourceName := "data.checkpoint_management_access_point_name.data_access_point_name"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementAccessPointNameConfig(objName, "apnname", true, "All_Internet"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementAccessPointNameConfig(name string, apn string, enforceEndUserDomain bool, endUserDomain string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_access_point_name" "access_point_name" {
        name = "%s"
        apn = "%s"
        enforce_end_user_domain = %t
        end_user_domain = "%s"
}

data "checkpoint_management_access_point_name" "data_access_point_name" {
    name = "${checkpoint_management_access_point_name.access_point_name.name}"
}
`, name, apn, enforceEndUserDomain, endUserDomain)
}
