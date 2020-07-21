package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	"testing"
)

func TestAccCheckpointManagementGroup_basic(t *testing.T) {
	var group map[string]interface{}
	resourceName := "checkpoint_management_group.test"
	objName := "tfTestManagementGroup_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementGroupConfig(objName, "CP_default_Office_Mode_addresses_pool"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointGroupExists(resourceName, &group),
					testAccCheckCheckpointGroupAttributes(&group, objName, "CP_default_Office_Mode_addresses_pool"),
				),
			},
		},
	})
}

func testAccCheckpointGroupDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_group" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-group", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success { // Resource still exists. failed to destroy.
				return fmt.Errorf("group object (%s) still exists", rs.Primary.ID)
			}
		}
		break
	}
	return nil
}

func testAccCheckCheckpointGroupExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("group ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)
		response, _ := client.ApiCall("show-group", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return fmt.Errorf(response.ErrorMsg)
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointGroupAttributes(group *map[string]interface{}, name string, member1 string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		group := *group
		if group == nil {
			return fmt.Errorf("group is nil")
		}

		membersJson := group["members"].([]interface{})
		var membersIds = make([]string, 0)
		if len(membersJson) > 0 {
			// Create slice of members names
			for _, member := range membersJson {
				member := member.(map[string]interface{})
				membersIds = append(membersIds, member["name"].(string))
			}
		}

		groupName := group["name"].(string)
		if groupName != name {
			return fmt.Errorf("name is %s, expected %s", groupName, name)
		}
		groupMember1 := membersIds[0]
		if groupMember1 != member1 {
			return fmt.Errorf("member1 is %s, expected %s", groupMember1, member1)
		}

		return nil
	}
}

func testAccManagementGroupConfig(name string, member1 string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_group" "test" {
    name = "%s"
	members = ["%s"]
}
`, name, member1)
}
