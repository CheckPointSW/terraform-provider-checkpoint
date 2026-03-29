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

func TestAccCheckpointManagementTag_basic(t *testing.T) {

	var tagMap map[string]interface{}
	resourceName := "checkpoint_management_tag.test"
	objName := "tfTestManagementTag_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementTagDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementTagConfig(objName, "tag1", "tag2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementTagExists(resourceName, &tagMap),
					testAccCheckCheckpointManagementTagAttributes(&tagMap, objName, "tag1", "tag2"),
				),
			},
		},
	})
}

func testAccCheckpointManagementTagDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_tag" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-tag", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success {
				return fmt.Errorf("Tag object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementTagExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Tag ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-tag", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementTagAttributes(tagMap *map[string]interface{}, objName string, tags1 string, tags2 string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		tagName := (*tagMap)["name"].(string)
		if !strings.EqualFold(tagName, objName) {
			return fmt.Errorf("name is %s, expected %s", objName, tagName)
		}
		tagsJson := (*tagMap)["tags"].([]interface{})
		var tagsIds = make([]string, 0)
		if len(tagsJson) > 0 {
			for _, tags := range tagsJson {
				tagsTry1, ok := tags.(map[string]interface{})
				if ok {
					tagsIds = append([]string{tagsTry1["name"].(string)}, tagsIds...)
				} else {
					tagsTry2 := tags.(string)
					tagsIds = append([]string{tagsTry2}, tagsIds...)
				}
			}
		}

		Tagtags1 := tagsIds[0]
		if Tagtags1 != tags1 {
			return fmt.Errorf("tags1 is %s, expected %s", tags1, Tagtags1)
		}
		Tagtags2 := tagsIds[1]
		if Tagtags2 != tags2 {
			return fmt.Errorf("tags2 is %s, expected %s", tags2, Tagtags2)
		}
		return nil
	}
}

func testAccManagementTagConfig(objName string, tags1 string, tags2 string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_tag" "test" {
        name = "%s"
        tags = ["%s","%s"]
}
`, objName, tags1, tags2)
}
