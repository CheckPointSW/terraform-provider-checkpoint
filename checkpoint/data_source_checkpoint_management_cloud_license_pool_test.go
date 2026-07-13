package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementCloudLicensePool_basic(t *testing.T) {

	objName := "tfTestManagementDataCloudLicensePool_" + acctest.RandString(6)
	resourceName := "checkpoint_management_cloud_license_pool.cloud_license_pool"
	dataSourceName := "data.checkpoint_management_data_cloud_license_pool.data_cloud_license_pool"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementCloudLicensePoolConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementCloudLicensePoolConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_cloud_license_pool" "cloud_license_pool" {

    name = "%s"
	pool = "pool_test"
	available_quota = "available_quota_test"
}

data "checkpoint_management_data_cloud_license_pool" "data_cloud_license_pool" {
    name = "${checkpoint_management_cloud_license_pool.cloud_license_pool.name}"
}
`, name)
}
