package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementIdpAdministratorGroup_basic(t *testing.T) {

	objName := "tfTestManagementDataIdpAdministratorGroup_" + acctest.RandString(6)
	resourceName := "checkpoint_management_idp_administrator_group.idp_administrator_group"
	dataSourceName := "data.checkpoint_management_idp_administrator_group.data_idp_administrator_group"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementIdpAdministratorGroupConfig(objName, "it-team", "domain super user"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "group_id", resourceName, "group_id"),
				),
			},
		},
	})

}

func testAccDataSourceManagementIdpAdministratorGroupConfig(name string, groupId string, multiDomainProfile string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_idp_administrator_group" "idp_administrator_group" {
	name = "%s"
	group_id = "%s"
	multi_domain_profile = "%s"
}

data "checkpoint_management_idp_administrator_group" "data_idp_administrator_group" {
    name = "${checkpoint_management_idp_administrator_group.idp_administrator_group.name}"
}
`, name, groupId, multiDomainProfile)
}
