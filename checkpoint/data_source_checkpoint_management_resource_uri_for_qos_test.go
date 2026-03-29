package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementUriForQos_basic(t *testing.T) {

	objName := "tfTestManagementDataUriForQos_" + acctest.RandString(6)
	resourceName := "checkpoint_management_resource_uri_for_qos.test"
	dataSourceName := "data.checkpoint_management_resource_uri_for_qos.data_uri_for_qos"
	searchForUrl := "www.checkpoint.com"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementUriForQosConfig(objName, searchForUrl),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "search_for_url", resourceName, "search_for_url"),
				),
			},
		},
	})

}

func testAccDataSourceManagementUriForQosConfig(name string, searchForUrl string) string {
	return fmt.Sprintf(`

resource "checkpoint_management_resource_uri_for_qos" "test" {
	name = "%s"
	search_for_url = "%s"
}

data "checkpoint_management_resource_uri_for_qos" "data_uri_for_qos" {
  name = "${checkpoint_management_resource_uri_for_qos.test.name}"
}
`, name, searchForUrl)
}
