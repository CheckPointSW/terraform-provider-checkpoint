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

func TestAccCheckpointManagementResourceUriForQos_basic(t *testing.T) {

	var resourceUriForQosMap map[string]interface{}
	resourceName := "checkpoint_management_resource_uri_for_qos.test"
	objName := "tfTestManagementResourceUriForQos_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementResourceUriForQosDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementResourceUriForQosConfig(objName, "www.checkpoint.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementResourceUriForQosExists(resourceName, &resourceUriForQosMap),
					testAccCheckCheckpointManagementResourceUriForQosAttributes(&resourceUriForQosMap, objName, "www.checkpoint.com"),
				),
			},
		},
	})
}

func testAccCheckpointManagementResourceUriForQosDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_resource_uri_for_qos" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-resource-uri-for-qos", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success {
				return fmt.Errorf("ResourceUriForQos object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementResourceUriForQosExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ResourceUriForQos ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-resource-uri-for-qos", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementResourceUriForQosAttributes(resourceUriForQosMap *map[string]interface{}, name string, searchForUrl string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		resourceUriForQosName := (*resourceUriForQosMap)["name"].(string)
		if !strings.EqualFold(resourceUriForQosName, name) {
			return fmt.Errorf("name is %s, expected %s", name, resourceUriForQosName)
		}
		resourceUriForQosSearchForUrl := (*resourceUriForQosMap)["search-for-url"].(string)
		if !strings.EqualFold(resourceUriForQosSearchForUrl, searchForUrl) {
			return fmt.Errorf("searchForUrl is %s, expected %s", searchForUrl, resourceUriForQosSearchForUrl)
		}
		return nil
	}
}

func testAccManagementResourceUriForQosConfig(name string, searchForUrl string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_resource_uri_for_qos" "test" {
        name = "%s"
        search_for_url = "%s"
}
`, name, searchForUrl)
}
