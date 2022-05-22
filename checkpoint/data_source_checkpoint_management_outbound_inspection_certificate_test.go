package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementOutboundInspectionCertificate_basic(t *testing.T) {

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementOutboundInspectionCertificateConfig(),
				Check: resource.ComposeTestCheckFunc(
				),
			},
		},
	})

}

func testAccDataSourceManagementOutboundInspectionCertificateConfig() string {
	return fmt.Sprintf(`
data "checkpoint_management_outbound_inspection_certificate" "data_outbound_inspection_certificate" {
}
`)
}
