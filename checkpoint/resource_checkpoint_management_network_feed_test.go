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

func TestAccCheckpointManagementNetworkFeed_basic(t *testing.T) {

	var networkFeedMap map[string]interface{}
	resourceName := "checkpoint_management_network_feed.test"
	objName := "tfTestManagementNetworkFeed_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementNetworkFeedDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementNetworkFeedConfig(objName, "https://www.feedsresource.com/resource", "feed_username", "feed_password", "Flat List", "IP Address", 60, 1, false, "	", "!"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementNetworkFeedExists(resourceName, &networkFeedMap),
					testAccCheckCheckpointManagementNetworkFeedAttributes(&networkFeedMap, objName, "https://www.feedsresource.com/resource", "feed_username", "feed_password", "Flat List", "IP Address", 60, 1, false, "	", "!"),
				),
			},
		},
	})
}

func testAccCheckpointManagementNetworkFeedDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_network_feed" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-network-feed", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("NetworkFeed object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementNetworkFeedExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("NetworkFeed ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-network-feed", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementNetworkFeedAttributes(networkFeedMap *map[string]interface{}, name string, feedUrl string, username string, password string, feedFormat string, feedType string, updateInterval int, dataColumn int, useGatewayProxy bool, fieldsDelimiter string, ignoreLinesThatStartWith string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		networkFeedName := (*networkFeedMap)["name"].(string)
		if !strings.EqualFold(networkFeedName, name) {
			return fmt.Errorf("name is %s, expected %s", name, networkFeedName)
		}
		networkFeedFeedUrl := (*networkFeedMap)["feed-url"].(string)
		if !strings.EqualFold(networkFeedFeedUrl, feedUrl) {
			return fmt.Errorf("feedUrl is %s, expected %s", feedUrl, networkFeedFeedUrl)
		}
		networkFeedUsername := (*networkFeedMap)["username"].(string)
		if !strings.EqualFold(networkFeedUsername, username) {
			return fmt.Errorf("username is %s, expected %s", username, networkFeedUsername)
		}
		networkFeedFeedFormat := (*networkFeedMap)["feed-format"].(string)
		if !strings.EqualFold(networkFeedFeedFormat, feedFormat) {
			return fmt.Errorf("feedFormat is %s, expected %s", feedFormat, networkFeedFeedFormat)
		}
		networkFeedFeedType := (*networkFeedMap)["feed-type"].(string)
		if !strings.EqualFold(networkFeedFeedType, feedType) {
			return fmt.Errorf("feedType is %s, expected %s", feedType, networkFeedFeedType)
		}
		networkFeedUpdateInterval := int((*networkFeedMap)["update-interval"].(float64))
		if networkFeedUpdateInterval != updateInterval {
			return fmt.Errorf("updateInterval is %d, expected %d", updateInterval, networkFeedUpdateInterval)
		}
		networkFeedDataColumn := int((*networkFeedMap)["data-column"].(float64))
		if networkFeedDataColumn != dataColumn {
			return fmt.Errorf("dataColumn is %d, expected %d", dataColumn, networkFeedDataColumn)
		}
		networkFeedUseGatewayProxy := (*networkFeedMap)["use-gateway-proxy"].(bool)
		if networkFeedUseGatewayProxy != useGatewayProxy {
			return fmt.Errorf("useGatewayProxy is %t, expected %t", useGatewayProxy, networkFeedUseGatewayProxy)
		}
		networkFeedFieldsDelimiter := (*networkFeedMap)["fields-delimiter"].(string)
		if !strings.EqualFold(networkFeedFieldsDelimiter, fieldsDelimiter) {
			return fmt.Errorf("fieldsDelimiter is %s, expected %s", fieldsDelimiter, networkFeedFieldsDelimiter)
		}
		networkFeedIgnoreLinesThatStartWith := (*networkFeedMap)["ignore-lines-that-start-with"].(string)
		if !strings.EqualFold(networkFeedIgnoreLinesThatStartWith, ignoreLinesThatStartWith) {
			return fmt.Errorf("ignoreLinesThatStartWith is %s, expected %s", ignoreLinesThatStartWith, networkFeedIgnoreLinesThatStartWith)
		}
		return nil
	}
}

func testAccManagementNetworkFeedConfig(name string, feedUrl string, username string, password string, feedFormat string, feedType string, updateInterval int, dataColumn int, useGatewayProxy bool, fieldsDelimiter string, ignoreLinesThatStartWith string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_network_feed" "test" {
        name = "%s"
        feed_url = "%s"
        username = "%s"
        password = "%s"
        feed_format = "%s"
        feed_type = "%s"
        update_interval = %d
        data_column = %d
        use_gateway_proxy = %t
        fields_delimiter = "%s"
        ignore_lines_that_start_with = "%s"
}
`, name, feedUrl, username, password, feedFormat, feedType, updateInterval, dataColumn, useGatewayProxy, fieldsDelimiter, ignoreLinesThatStartWith)
}
