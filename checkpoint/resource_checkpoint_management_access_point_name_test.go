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

func TestAccCheckpointManagementAccessPointName_basic(t *testing.T) {

	var accessPointNameMap map[string]interface{}
	resourceName := "checkpoint_management_access_point_name.test"
	objName := "tfTestManagementAccessPointName_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementAccessPointNameDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementAccessPointNameConfig(objName, "apnname", true, "All_Internet"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementAccessPointNameExists(resourceName, &accessPointNameMap),
					testAccCheckCheckpointManagementAccessPointNameAttributes(&accessPointNameMap, objName, "apnname", true, "All_Internet"),
				),
			},
		},
	})
}

func testAccCheckpointManagementAccessPointNameDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_access_point_name" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-access-point-name", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("AccessPointName object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementAccessPointNameExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("AccessPointName ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-access-point-name", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementAccessPointNameAttributes(accessPointNameMap *map[string]interface{}, name string, apn string, enforceEndUserDomain bool, endUserDomain string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		accessPointNameName := (*accessPointNameMap)["name"].(string)
		if !strings.EqualFold(accessPointNameName, name) {
			return fmt.Errorf("name is %s, expected %s", name, accessPointNameName)
		}
		accessPointNameApn := (*accessPointNameMap)["apn"].(string)
		if !strings.EqualFold(accessPointNameApn, apn) {
			return fmt.Errorf("apn is %s, expected %s", apn, accessPointNameApn)
		}
		accessPointNameEnforceEndUserDomain := (*accessPointNameMap)["enforce-end-user-domain"].(bool)
		if accessPointNameEnforceEndUserDomain != enforceEndUserDomain {
			return fmt.Errorf("enforceEndUserDomain is %t, expected %t", enforceEndUserDomain, accessPointNameEnforceEndUserDomain)
		}
		accessPointNameEndUserDomain := (*accessPointNameMap)["end-user-domain"].(map[string]interface{})["name"].(string)
		if !strings.EqualFold(accessPointNameEndUserDomain, endUserDomain) {
			return fmt.Errorf("endUserDomain is %s, expected %s", endUserDomain, accessPointNameEndUserDomain)
		}
		return nil
	}
}

func testAccManagementAccessPointNameConfig(name string, apn string, enforceEndUserDomain bool, endUserDomain string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_access_point_name" "test" {
        name = "%s"
        apn = "%s"
        enforce_end_user_domain = %t
        end_user_domain = "%s"
}
`, name, apn, enforceEndUserDomain, endUserDomain)
}
