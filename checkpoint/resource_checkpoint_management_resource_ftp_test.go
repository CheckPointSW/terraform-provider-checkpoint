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

func TestAccCheckpointManagementResourceFtp_basic(t *testing.T) {

	var resourceFtpMap map[string]interface{}
	resourceName := "checkpoint_management_resource_ftp.test"
	objName := "tfTestManagementResourceFtp_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementResourceFtpDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementResourceFtpConfig(objName, "get_and_put", "Exception Log", "path"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementResourceFtpExists(resourceName, &resourceFtpMap),
					testAccCheckCheckpointManagementResourceFtpAttributes(&resourceFtpMap, objName, "get_and_put", "exception log", "path"),
				),
			},
		},
	})
}

func testAccCheckpointManagementResourceFtpDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_resource_ftp" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-resource-ftp", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success {
				return fmt.Errorf("ResourceFtp object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementResourceFtpExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ResourceFtp ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-resource-ftp", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementResourceFtpAttributes(resourceFtpMap *map[string]interface{}, name string, resourceMatchingMethod string, exceptionTrack string, resourcesPath string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		resourceFtpName := (*resourceFtpMap)["name"].(string)
		if !strings.EqualFold(resourceFtpName, name) {
			return fmt.Errorf("name is %s, expected %s", name, resourceFtpName)
		}
		resourceFtpResourceMatchingMethod := (*resourceFtpMap)["resource-matching-method"].(string)
		if !strings.EqualFold(resourceFtpResourceMatchingMethod, resourceMatchingMethod) {
			return fmt.Errorf("resourceMatchingMethod is %s, expected %s", resourceMatchingMethod, resourceFtpResourceMatchingMethod)
		}
		return nil
	}
}

func testAccManagementResourceFtpConfig(name string, resourceMatchingMethod string, exceptionTrack string, resourcesPath string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_resource_ftp" "test" {
        name = "%s"
        resource_matching_method = "%s"
        exception_track = "%s"
        resources_path = "%s"
cvp {
    allowed_to_modify_content = true
    enable_cvp =  false
    reply_order = "return_data_before_content_is_approved"

  }
}
`, name, resourceMatchingMethod, exceptionTrack, resourcesPath)
}
