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

func TestAccCheckpointManagementNatSection_basic(t *testing.T) {

	var natSectionMap map[string]interface{}
	resourceName := "checkpoint_management_nat_section.nat_section"
	objName := "tfTestManagementNatSection_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementNatSectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementNatSectionConfig(objName, "Standard", "top"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementNatSectionExists(resourceName, &natSectionMap),
					testAccCheckCheckpointManagementNatSectionAttributes(&natSectionMap, objName),
				),
			},
		},
	})
}

func testAccCheckpointManagementNatSectionDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_nat_section" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-nat-section", map[string]interface{}{"uid": rs.Primary.ID, "package": "Standard"}, client.GetSessionID(), true, false)
			if res.Success {
				return fmt.Errorf("NAT section object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementNatSectionExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("NAT section ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-nat-section", map[string]interface{}{"uid": rs.Primary.ID, "package": "Standard"}, client.GetSessionID(), true, false)
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementNatSectionAttributes(natSectionMap *map[string]interface{}, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		natSectionName := (*natSectionMap)["name"].(string)
		if !strings.EqualFold(natSectionName, name) {
			return fmt.Errorf("name is %s, expected %s", name, natSectionName)
		}

		return nil
	}
}

func testAccManagementNatSectionConfig(name string, packageName string, pos string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_nat_section" "nat_section" {
        name = "%s"
        package = "%s"
		position = { "top": "%s" }
}
`, name, packageName, pos)
}
