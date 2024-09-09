package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	"strings"
	"testing"
)

func TestAccCheckpointManagementMobileAccessSection_basic(t *testing.T) {

	var mobileAccessSectionMap map[string]interface{}
	resourceName := "checkpoint_management_mobile_access_section.test"
	objName := "tfTestManagementMobileAccessSection_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementMobileAccessSectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementMobileAccessSectionConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementMobileAccessSectionExists(resourceName, &mobileAccessSectionMap),
					testAccCheckCheckpointManagementMobileAccessSectionAttributes(&mobileAccessSectionMap, objName),
				),
			},
		},
	})
}

func testAccCheckpointManagementMobileAccessSectionDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_mobile_access_section" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-mobile-access-section", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success {
				return fmt.Errorf("MobileAccessSection object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementMobileAccessSectionExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("MobileAccessSection ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-mobile-access-section", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementMobileAccessSectionAttributes(mobileAccessSectionMap *map[string]interface{}, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		mobileAccessSectionName := (*mobileAccessSectionMap)["name"].(string)
		if !strings.EqualFold(mobileAccessSectionName, name) {
			return fmt.Errorf("name is %s, expected %s", name, mobileAccessSectionName)
		}
		return nil
	}
}

func testAccManagementMobileAccessSectionConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_mobile_access_section" "test" {
        name = "%s"
        position = {top = "top"}
}
`, name)
}
