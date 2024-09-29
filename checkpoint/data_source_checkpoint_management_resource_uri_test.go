package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementResourceUri_basic(t *testing.T) {

	objName := "tfTestManagementDataResouceUri_" + acctest.RandString(6)
	resourceName := "checkpoint_management_resource_uri.test"
	dataSourceName := "data.checkpoint_management_resource_uri.data"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementResourceUriConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementResourceUriConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_resource_uri" "test" {

  name = "%s"
  use_this_resource_to = "optimize_url_logging"
  connection_methods = {
    transparent = "false"
    tunneling = "true"
    proxy = false
  }
  uri_match_specification_type = "wildcards"
  match_wildcards {
    host = "hostName"
    path = "pathName"
    schemes {
      gopher = true
      other = "string2"
    }
    methods {
      get = true
      post = true
      head = true
      put = true
      other = "done7"
    }
  }
  action {
    strip_activex_tags =  true
    strip_applet_tags = true
    strip_ftp_links = true
    strip_port_strings = true
    strip_script_tags = true

  }
  soap = {
    inspection = "allow_all_soap_requests"
    file_id = "scheme1"
    track_connections = "mail_alert"
  }
}

data "checkpoint_management_resource_uri" "data" {
  uid = "${checkpoint_management_resource_uri.test.id}"
}
`, name)
}
