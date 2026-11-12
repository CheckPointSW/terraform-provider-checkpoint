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

func TestAccCheckpointManagementResourceMms_basic(t *testing.T) {

	var resourceMmsMap map[string]interface{}
	resourceName := "checkpoint_management_resource_mms.test"
	objName := "tfTestManagementResourceMms_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementResourceMmsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementResourceMmsConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementResourceMmsExists(resourceName, &resourceMmsMap),
					testAccCheckCheckpointManagementResourceMmsAttributes(&resourceMmsMap, objName),
				),
			},
		},
	})
}

func testAccCheckpointManagementResourceMmsDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_resource_mms" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-resource-mms", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success {
				return fmt.Errorf("ResourceMms object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementResourceMmsExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ResourceMms ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-resource-mms", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementResourceMmsAttributes(resourceMmsMap *map[string]interface{}, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		resourceMmsName := (*resourceMmsMap)["name"].(string)
		if !strings.EqualFold(resourceMmsName, name) {
			return fmt.Errorf("name is %s, expected %s", name, resourceMmsName)
		}
		return nil
	}
}

func testAccManagementResourceMmsConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_resource_mms" "test" {
        name = "%s"
}
`, name)
}
