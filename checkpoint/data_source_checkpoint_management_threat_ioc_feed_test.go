package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementThreatIocFeed_basic(t *testing.T) {

	objName := "tfTestManagementDataThreatIocFeed_" + acctest.RandString(6)
	resourceName := "checkpoint_management_threat_ioc_feed.threat_ioc_feed"
	dataSourceName := "data.checkpoint_management_threat_ioc_feed.data_threat_ioc_feed"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementThreatIocFeedConfig(objName, "https://www.feedsresource.com/resource"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "feed_url", resourceName, "feed_url"),
				),
			},
		},
	})

}

func testAccDataSourceManagementThreatIocFeedConfig(name string, feedUrl string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_threat_ioc_feed" "threat_ioc_feed" {
    name = "%s"
	feed_url = "%s"
}

data "checkpoint_management_threat_ioc_feed" "data_threat_ioc_feed" {
    name = "${checkpoint_management_threat_ioc_feed.threat_ioc_feed.name}"
}
`, name, feedUrl)
}
