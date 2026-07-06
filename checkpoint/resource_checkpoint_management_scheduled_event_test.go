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

func TestAccCheckpointManagementScheduledEvent_basic(t *testing.T) {

	var scheduledEventMap map[string]interface{}
	resourceName := "checkpoint_management_scheduled_event.test"
	objName := "tfTestManagementScheduledEvent_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementScheduledEventDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementScheduledEventConfig(objName, "scheduled event for daily backup operations"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementScheduledEventExists(resourceName, &scheduledEventMap),
					testAccCheckCheckpointManagementScheduledEventAttributes(&scheduledEventMap, objName, "scheduled event for daily backup operations"),
				),
			},
		},
	})
}

func testAccCheckpointManagementScheduledEventDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_scheduled_event" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-scheduled-event", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("ScheduledEvent object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementScheduledEventExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ScheduledEvent ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-scheduled-event", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementScheduledEventAttributes(scheduledEventMap *map[string]interface{}, name string, comments string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		scheduledEventName := (*scheduledEventMap)["name"].(string)
		if !strings.EqualFold(scheduledEventName, name) {
			return fmt.Errorf("name is %s, expected %s", name, scheduledEventName)
		}
		scheduledEventComments := (*scheduledEventMap)["comments"].(string)
		if !strings.EqualFold(scheduledEventComments, comments) {
			return fmt.Errorf("comments is %s, expected %s", comments, scheduledEventComments)
		}
		return nil
	}
}

func testAccManagementScheduledEventConfig(name string, comments string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_scheduled_event" "test" {
        name = "%s"
        comments = "%s"
}
`, name, comments)
}
