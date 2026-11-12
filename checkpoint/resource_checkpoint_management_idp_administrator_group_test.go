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

func TestAccCheckpointManagementIdpAdministratorGroup_basic(t *testing.T) {

	var idpAdministratorGroupMap map[string]interface{}
	resourceName := "checkpoint_management_idp_administrator_group.test"
	objName := "tfTestManagementIdpAdministratorGroup_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementIdpAdministratorGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementIdpAdministratorGroupConfig(objName, "it-team", "domain super user"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementIdpAdministratorGroupExists(resourceName, &idpAdministratorGroupMap),
					testAccCheckCheckpointManagementIdpAdministratorGroupAttributes(&idpAdministratorGroupMap, objName, "it-team", "domain super user"),
				),
			},
		},
	})
}

func testAccCheckpointManagementIdpAdministratorGroupDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_idp_administrator_group" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-idp-administrator-group", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("idpAdministratorGroup object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementIdpAdministratorGroupExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("idpAdministratorGroup ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-idp-administrator-group", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementIdpAdministratorGroupAttributes(idpAdministratorGroupMap *map[string]interface{}, name string, groupId string, multiDomainProfile string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		idpAdministratorGroupName := (*idpAdministratorGroupMap)["name"].(string)
		if !strings.EqualFold(idpAdministratorGroupName, name) {
			return fmt.Errorf("name is %s, expected %s", name, idpAdministratorGroupName)
		}
		idpAdministratorGroupGroupId := (*idpAdministratorGroupMap)["group-id"].(string)
		if !strings.EqualFold(idpAdministratorGroupGroupId, groupId) {
			return fmt.Errorf("groupId is %s, expected %s", groupId, idpAdministratorGroupGroupId)
		}
		idpAdministratorGroupMultiDomainProfile := (*idpAdministratorGroupMap)["multi-domain-profile"].(string)
		if !strings.EqualFold(idpAdministratorGroupMultiDomainProfile, multiDomainProfile) {
			return fmt.Errorf("multi_domain_profile is %s, expected %s", multiDomainProfile, idpAdministratorGroupGroupId)
		}
		return nil
	}
}

func testAccManagementIdpAdministratorGroupConfig(name string, groupId string, multiDomainProfile string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_idp_administrator_group" "test" {
        name = "%s"
        group_id = "%s"
		multi_domain_profile = "%s"
}
`, name, groupId, multiDomainProfile)
}
