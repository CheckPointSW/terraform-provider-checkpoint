package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementVoipDomainH323Gateway_basic(t *testing.T) {
	resourceName := "checkpoint_management_voip_domain_h323_gateway.test"
	dataSourceName := "data.checkpoint_management_voip_domain_h323_gateway.data_test"
	objectName := "test-voip_domain_h323_gateway"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementVoipDomainH323GatewayConfig(objectName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "color", resourceName, "color"),
					resource.TestCheckResourceAttrPair(dataSourceName, "comments", resourceName, "comments"),
					resource.TestCheckResourceAttrPair(dataSourceName, "icon", resourceName, "icon"),
				),
			},
		},
	})
}

func testAccDataSourceManagementVoipDomainH323GatewayConfig(objectName string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_voip_domain_h323_gateway" "test" {
	name = "%s"
	color = "test-value"
	comments = "test-value"
}

data "checkpoint_management_voip_domain_h323_gateway" "data_test" {
	name = "${checkpoint_management_voip_domain_h323_gateway.test.name}"
}
`, objectName)
}
