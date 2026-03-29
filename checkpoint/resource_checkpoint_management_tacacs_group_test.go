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

func TestAccCheckpointManagementTacacsGroup_basic(t *testing.T) {

	var tacacsGroupMap map[string]interface{}
	resourceName := "checkpoint_management_tacacs_group.test"
	objName := "tfTestManagementTacacsGroup_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementTacacsGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementTacacsGroupConfig(objName, "my_t"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementTacacsGroupExists(resourceName, &tacacsGroupMap),
					testAccCheckCheckpointManagementTacacsGroupAttributes(&tacacsGroupMap, objName, "my_t"),
				),
			},
		},
	})
}

func testAccCheckpointManagementTacacsGroupDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_tacacs_group" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-tacacs-group", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success {
				return fmt.Errorf("TacacsGroup object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementTacacsGroupExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("TacacsGroup ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-tacacs-group", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementTacacsGroupAttributes(tacacsGroupMap *map[string]interface{}, name string, members1 string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		tacacsGroupName := (*tacacsGroupMap)["name"].(string)
		if !strings.EqualFold(tacacsGroupName, name) {
			return fmt.Errorf("name is %s, expected %s", name, tacacsGroupName)
		}
		membersJson := (*tacacsGroupMap)["members"].([]interface{})
		var membersIds = make([]string, 0)
		if len(membersJson) > 0 {
			for _, members := range membersJson {
				membersTry1, ok := members.(map[string]interface{})
				if ok {
					membersIds = append([]string{membersTry1["name"].(string)}, membersIds...)
				} else {
					membersTry2 := members.(string)
					membersIds = append([]string{membersTry2}, membersIds...)
				}
			}
		}

		TacacsGroupmembers1 := membersIds[0]
		if TacacsGroupmembers1 != members1 {
			return fmt.Errorf("members1 is %s, expected %s", members1, TacacsGroupmembers1)
		}

		return nil
	}
}

func testAccManagementTacacsGroupConfig(name string, members1 string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_tacacs_group" "test" {
        name = "%s"
        members = ["%s"]
}
`, name, members1)
}
