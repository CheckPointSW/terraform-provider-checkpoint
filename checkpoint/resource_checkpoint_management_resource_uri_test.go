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

func TestAccCheckpointManagementResourceUri_basic(t *testing.T) {

	var resourceUriMap map[string]interface{}
	resourceName := "checkpoint_management_resource_uri.test"
	objName := "tfTestManagementResourceUri_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementResourceUriDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementResourceUriConfig(objName, "optimize_url_logging", "wildcards"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementResourceUriExists(resourceName, &resourceUriMap),
					testAccCheckCheckpointManagementResourceUriAttributes(&resourceUriMap, objName, "optimize_url_logging", "wildcards"),
				),
			},
		},
	})
}

func testAccCheckpointManagementResourceUriDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_resource_uri" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-resource-uri", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success {
				return fmt.Errorf("ResourceUri object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementResourceUriExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ResourceUri ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-resource-uri", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementResourceUriAttributes(resourceUriMap *map[string]interface{}, name string, useThisResourceTo string, uriMatchSpecificationType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		resourceUriName := (*resourceUriMap)["name"].(string)
		if !strings.EqualFold(resourceUriName, name) {
			return fmt.Errorf("name is %s, expected %s", name, resourceUriName)
		}
		resourceUriUseThisResourceTo := (*resourceUriMap)["use-this-resource-to"].(string)
		if !strings.EqualFold(resourceUriUseThisResourceTo, useThisResourceTo) {
			return fmt.Errorf("useThisResourceTo is %s, expected %s", useThisResourceTo, resourceUriUseThisResourceTo)
		}
		resourceUriUriMatchSpecificationType := (*resourceUriMap)["uri-match-specification-type"].(string)
		if !strings.EqualFold(resourceUriUriMatchSpecificationType, uriMatchSpecificationType) {
			return fmt.Errorf("uriMatchSpecificationType is %s, expected %s", uriMatchSpecificationType, resourceUriUriMatchSpecificationType)
		}
		return nil
	}
}

func testAccManagementResourceUriConfig(name string, useThisResourceTo string, uriMatchSpecificationType string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_resource_uri" "test" {
        name = "%s"
        use_this_resource_to = "%s"
        uri_match_specification_type = "%s"
}
`, name, useThisResourceTo, uriMatchSpecificationType)
}
