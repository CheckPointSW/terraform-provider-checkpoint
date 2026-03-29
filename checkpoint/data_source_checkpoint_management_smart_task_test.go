package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementSmartTask_basic(t *testing.T) {

	objName := "tfTestManagementDataServiceIcmp6_" + acctest.RandString(6)
	resourceName := "checkpoint_management_smart_task.test"
	dataSourceName := "data.checkpoint_management_smart_task.data_test"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementSmartTaskConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "action", resourceName, "action"),
					resource.TestCheckResourceAttrPair(dataSourceName, "trigger", resourceName, "trigger"),
				),
			},
		},
	})
}

func testAccDataSourceManagementSmartTaskConfig(name string) string {
	return fmt.Sprintf(`
   resource "checkpoint_management_smart_task" "test" {

  name = "%s"
  trigger = "Before Publish"
  description = "my smart task"
  action {

    send_web_request {
      url            = "https://demo.example.com/policy-installation-reports"
      fingerprint    = "8023a5652ba2c8f5b0902363a5314cd2b4fdbc5c"
      override_proxy = true
      proxy_url      = "https://demo.example.com/policy-installation-reports"
      time_out       = 200
      shared_secret  = " secret"
    }
  }
  enabled = true
}
data "checkpoint_management_smart_task" "data_test" {
 name = "${checkpoint_management_smart_task.test.name}"
}
`, name)
}
