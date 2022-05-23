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

func TestAccCheckpointManagementApplicationSiteCategory_basic(t *testing.T) {

	var applicationSiteCategoryMap map[string]interface{}
	resourceName := "checkpoint_management_application_site_category.test"
	objName := "tfTestManagementApplicationSiteCategory_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementApplicationSiteCategoryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementApplicationSiteCategoryConfig(objName, "my application site category"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementApplicationSiteCategoryExists(resourceName, &applicationSiteCategoryMap),
					testAccCheckCheckpointManagementApplicationSiteCategoryAttributes(&applicationSiteCategoryMap, objName, "my application site category"),
				),
			},
		},
	})
}

func testAccCheckpointManagementApplicationSiteCategoryDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_application_site_category" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-application-site-category", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("ApplicationSiteCategory object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementApplicationSiteCategoryExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ApplicationSiteCategory ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-application-site-category", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementApplicationSiteCategoryAttributes(applicationSiteCategoryMap *map[string]interface{}, name string, description string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		applicationSiteCategoryName := (*applicationSiteCategoryMap)["name"].(string)
		if !strings.EqualFold(applicationSiteCategoryName, name) {
			return fmt.Errorf("name is %s, expected %s", name, applicationSiteCategoryName)
		}
		applicationSiteCategoryDescription := (*applicationSiteCategoryMap)["description"].(string)
		if !strings.EqualFold(applicationSiteCategoryDescription, description) {
			return fmt.Errorf("description is %s, expected %s", description, applicationSiteCategoryDescription)
		}
		return nil
	}
}

func testAccManagementApplicationSiteCategoryConfig(name string, description string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_application_site_category" "test" {
        name = "%s"
        description = "%s"
}
`, name, description)
}
