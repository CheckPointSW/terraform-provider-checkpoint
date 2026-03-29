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

func TestAccCheckpointManagementResourceCifs_basic(t *testing.T) {

	var resourceCifsMap map[string]interface{}
	resourceName := "checkpoint_management_resource_cifs.test"
	objName := "tfTestManagementResourceCifs_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementResourceCifsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementResourceCifsConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementResourceCifsExists(resourceName, &resourceCifsMap),
					testAccCheckCheckpointManagementResourceCifsAttributes(&resourceCifsMap, objName),
				),
			},
		},
	})
}

func testAccCheckpointManagementResourceCifsDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_resource_cifs" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-resource-cifs", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success {
				return fmt.Errorf("ResourceCifs object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementResourceCifsExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ResourceCifs ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-resource-cifs", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementResourceCifsAttributes(resourceCifsMap *map[string]interface{}, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		resourceCifsName := (*resourceCifsMap)["name"].(string)
		if !strings.EqualFold(resourceCifsName, name) {
			return fmt.Errorf("name is %s, expected %s", name, resourceCifsName)
		}
		return nil
	}
}

func testAccManagementResourceCifsConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_resource_cifs" "test" {

   name = "%s"
   allowed_disk_and_print_shares {
     server_name = "server1"
     share_name = "share12"
   }
}
`, name)
}
