package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementNetworkFeed_basic(t *testing.T) {

	objName := "tfTestManagementDataNetworkFeed_" + acctest.RandString(6)
	resourceName := "checkpoint_management_network_feed.network_feed"
	dataSourceName := "data.checkpoint_management_network_feed.data_network_feed"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementNetworkFeedConfig(objName, "https://www.feedsresource.com/resource"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "feed_url", resourceName, "feed_url"),
				),
			},
		},
	})

}

func testAccDataSourceManagementNetworkFeedConfig(name string, feedUrl string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_network_feed" "network_feed" {
    name = "%s"
	feed_url = "%s"
}

data "checkpoint_management_network_feed" "data_network_feed" {
    name = "${checkpoint_management_network_feed.network_feed.name}"
}
`, name, feedUrl)
}
