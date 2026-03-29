package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementLogExporter_basic(t *testing.T) {

	objName := "tfTestManagementDataLogExporter_" + acctest.RandString(6)
	resourceName := "checkpoint_management_log_exporter.test"
	dataSourceName := "data.checkpoint_management_log_exporter.data_log_exporter"
	targetServer := "1.2.3.4"
	targetPort := 1234
	protocol := "tcp"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementLogExporterConfig(objName, targetServer, targetPort, protocol),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "target_server", resourceName, "target_server"),
					resource.TestCheckResourceAttrPair(dataSourceName, "target_port", resourceName, "target_port"),
					resource.TestCheckResourceAttrPair(dataSourceName, "protocol", resourceName, "protocol"),
					resource.TestCheckResourceAttrPair(dataSourceName, "enabled", resourceName, "enabled"),
				),
			},
		},
	})

}

func testAccDataSourceManagementLogExporterConfig(name string, targetServer string, targetPort int, protocol string) string {
	return fmt.Sprintf(`

resource "checkpoint_management_log_exporter" "test" {
  	name = "%s"
	target_server = "%s"
	target_port = %d
	protocol = "%s"
	enabled = true
	attachments {
		add_link_to_log_attachment = true
		add_link_to_log_details = false
		add_log_attachment_id = true
	}
	
	data_manipulation {
		aggregate_log_updates = true
		format = "splunk"
	}
}

data "checkpoint_management_log_exporter" "data_log_exporter" {
	name = "${checkpoint_management_log_exporter.test.name}"
}
`, name, targetServer, targetPort, protocol)
}
