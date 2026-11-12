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

func TestAccCheckpointManagementThreatIocFeed_basic(t *testing.T) {

	var threatIocFeedMap map[string]interface{}
	resourceName := "checkpoint_management_threat_ioc_feed.test"
	objName := "tfTestManagementThreatIocFeed_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementThreatIocFeedDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementThreatIocFeedConfig(objName, "https://www.feedsresource.com/resource", "Prevent"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementThreatIocFeedExists(resourceName, &threatIocFeedMap),
					testAccCheckCheckpointManagementThreatIocFeedAttributes(&threatIocFeedMap, objName, "https://www.feedsresource.com/resource", "Prevent"),
				),
			},
		},
	})
}

func testAccCheckpointManagementThreatIocFeedDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_threat_ioc_feed" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-threat-ioc-feed", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("ThreatIocFeed object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementThreatIocFeedExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ThreatIocFeed ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-threat-ioc-feed", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementThreatIocFeedAttributes(threatIocFeedMap *map[string]interface{}, name string, feedUrl string, action string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		threatIocFeedName := (*threatIocFeedMap)["name"].(string)
		if !strings.EqualFold(threatIocFeedName, name) {
			return fmt.Errorf("name is %s, expected %s", name, threatIocFeedName)
		}
		threatIocFeedFeedUrl := (*threatIocFeedMap)["feed-url"].(string)
		if !strings.EqualFold(threatIocFeedFeedUrl, feedUrl) {
			return fmt.Errorf("feedUrl is %s, expected %s", feedUrl, threatIocFeedFeedUrl)
		}
		threatIocFeedAction := (*threatIocFeedMap)["action"].(string)
		if !strings.EqualFold(threatIocFeedAction, action) {
			return fmt.Errorf("action is %s, expected %s", action, threatIocFeedAction)
		}
		return nil
	}
}

func testAccManagementThreatIocFeedConfig(name string, feedUrl string, action string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_threat_ioc_feed" "test" {
        name = "%s"
        feed_url = "%s"
        action = "%s"
}
`, name, feedUrl, action)
}
