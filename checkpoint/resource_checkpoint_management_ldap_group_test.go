package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"os"
	"strings"
	"testing"
)

func TestAccCheckpointManagementLdapGroup_basic(t *testing.T) {

	var ldapGroupMap map[string]interface{}
	resourceName := "checkpoint_management_ldap_group.test"
	objName := "tfTestManagementLdapGroup_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementLdapGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementLdapGroupConfig(objName, "testldapaccountunit", "only_sub_tree", "poo = poo"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementLdapGroupExists(resourceName, &ldapGroupMap),
					testAccCheckCheckpointManagementLdapGroupAttributes(&ldapGroupMap, objName, "testldapaccountunit", "only_sub_tree", "poo = poo"),
				),
			},
		},
	})
}

func testAccCheckpointManagementLdapGroupDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_ldap_group" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-ldap-group", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success {
				return fmt.Errorf("LdapGroup object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementLdapGroupExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("LdapGroup ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-ldap-group", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementLdapGroupAttributes(ldapGroupMap *map[string]interface{}, name string, accountUnit string, scope string, accountUnitBranch string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		ldapGroupName := (*ldapGroupMap)["name"].(string)
		if !strings.EqualFold(ldapGroupName, name) {
			return fmt.Errorf("name is %s, expected %s", name, ldapGroupName)
		}
		ldapGroupAccountUnit := (*ldapGroupMap)["account-unit"].(map[string]interface{})["name"].(string)
		if !strings.EqualFold(ldapGroupAccountUnit, accountUnit) {
			return fmt.Errorf("accountUnit is %s, expected %s", accountUnit, ldapGroupAccountUnit)
		}
		ldapGroupScope := (*ldapGroupMap)["scope"].(string)
		if !strings.EqualFold(ldapGroupScope, scope) {
			return fmt.Errorf("scope is %s, expected %s", scope, ldapGroupScope)
		}

		ldapGroupAccountUnitBranch := (*ldapGroupMap)["account-unit-branch"].(string)
		if !strings.EqualFold(ldapGroupAccountUnitBranch, accountUnitBranch) {
			return fmt.Errorf("accountUnitBranch is %s, expected %s", accountUnitBranch, ldapGroupAccountUnitBranch)
		}

		return nil
	}
}

func testAccManagementLdapGroupConfig(name string, accountUnit string, scope string, accountUnitBranch string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_ldap_group" "test" {
    name = "%s"
    account_unit = "%s"
    scope = "%s"
    account_unit_branch = "%s"
}
`, name, accountUnit, scope, accountUnitBranch)
}
