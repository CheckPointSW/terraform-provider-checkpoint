package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementLdapGroup_basic(t *testing.T) {

	objName := "tfTestManagementDataLdapGroup_" + acctest.RandString(6)
	resourceName := "checkpoint_management_ldap_group.test"
	dataSourceName := "data.checkpoint_management_ldap_group.data_ldap_group"
	accountUnit := "testldapaccountunit"
	scope := "only_sub_tree"
	accountUnitBranch := "poo = poo"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementLdapGroupConfig(objName, accountUnit, scope, accountUnitBranch),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "account_unit", resourceName, "account_unit"),
					resource.TestCheckResourceAttrPair(dataSourceName, "scope", resourceName, "scope"),
					resource.TestCheckResourceAttrPair(dataSourceName, "account_unit_branch", resourceName, "account_unit_branch"),
				),
			},
		},
	})

}

func testAccDataSourceManagementLdapGroupConfig(name string, accountUnit string, scope string, accountUnitBranch string) string {
	return fmt.Sprintf(`

resource "checkpoint_management_ldap_group" "test" {
    name = "%s"
    account_unit = "%s"
    scope = "%s"
    account_unit_branch = "%s"
}

data "checkpoint_management_ldap_group" "data_ldap_group" {
	name = "${checkpoint_management_ldap_group.test.name}"
}
`, name, accountUnit, scope, accountUnitBranch)
}
