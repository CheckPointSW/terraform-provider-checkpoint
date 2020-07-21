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

func TestAccCheckpointManagementServiceGroup_basic(t *testing.T) {
	var serviceGroup map[string]interface{}
	resourceName := "checkpoint_management_service_group.test"
	objName := "tfTestManagementServiceGroup_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointServiceGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementServiceGroupConfig(objName, "domain-tcp", "domain-udp"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointServiceGroupExists(resourceName, &serviceGroup),
					testAccCheckCheckpointServiceGroupAttributes(&serviceGroup, objName, "domain-udp", "domain-tcp"),
				),
			},
		},
	})
}

func testAccCheckpointServiceGroupDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_service_group" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-service-group", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success { // Resource still exists. failed to destroy.
				return fmt.Errorf("service group object (%s) still exists", rs.Primary.ID)
			}
		}
		break
	}
	return nil
}

func testAccCheckCheckpointServiceGroupExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("service group ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)
		response, _ := client.ApiCall("show-service-group", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return fmt.Errorf(response.ErrorMsg)
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointServiceGroupAttributes(serviceGroup *map[string]interface{}, name string, member1 string, member2 string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		serviceGroup := *serviceGroup
		if serviceGroup == nil {
			return fmt.Errorf("service group is nil")
		}

		membersJson := serviceGroup["members"].([]interface{})
		var membersIds = make([]string, 0)
		if len(membersJson) > 0 {
			// Create slice of members names
			for _, member := range membersJson {
				member := member.(map[string]interface{})
				membersIds = append(membersIds, member["name"].(string))
			}
		}

		serviceGroupName := serviceGroup["name"].(string)
		if serviceGroupName != name {
			return fmt.Errorf("name is %s, expected %s", serviceGroupName, name)
		}
		serviceGroupMember1 := membersIds[0]
		if serviceGroupMember1 != member1 {
			return fmt.Errorf("member1 is %s, expected %s", serviceGroupMember1, member1)
		}
		serviceGroupMember2 := membersIds[1]
		if serviceGroupMember2 != member2 {
			return fmt.Errorf("member2 is %s, expected %s", serviceGroupMember2, member2)
		}

		return nil
	}
}

func testAccManagementServiceGroupConfig(name string, member1 string, member2 string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_service_group" "test" {
    name = "%s"
	members = ["%s","%s"]
}
`, name, member1, member2)
}
