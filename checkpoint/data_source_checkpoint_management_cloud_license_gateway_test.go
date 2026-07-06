package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementCloudLicenseGateway_basic(t *testing.T) {

	objName := "tfTestManagementDataCloudLicenseGateway_" + acctest.RandString(6)
	resourceName := "checkpoint_management_cloud_license_gateway.cloud_license_gateway"
	dataSourceName := "data.checkpoint_management_data_cloud_license_gateway.data_cloud_license_gateway"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementCloudLicenseGatewayConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementCloudLicenseGatewayConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_cloud_license_gateway" "cloud_license_gateway" {

    name = "%s"
	enable_auto_distribution = false
	scheduled_auto_distribution = "scheduled_auto_distribution_test"
}

data "checkpoint_management_data_cloud_license_gateway" "data_cloud_license_gateway" {
    name = "${checkpoint_management_cloud_license_gateway.cloud_license_gateway.name}"
}
`, name)
}
