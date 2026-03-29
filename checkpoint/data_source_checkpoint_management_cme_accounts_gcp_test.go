package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementCMEAccountsGCP_basic(t *testing.T) {
	resourceName := "checkpoint_management_cme_accounts_gcp.test"
	dataSourceName := "data.checkpoint_management_cme_accounts_gcp.data_test"
	accountName := "test-account"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this test")
	} else if context != "web_api" {
		t.Skip("Skipping cme api test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementCMEAccountsGCPConfig(accountName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "project_id", resourceName, "project_id"),
				),
			},
		},
	})
}

func testAccDataSourceManagementCMEAccountsGCPConfig(accountName string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_cme_accounts_gcp" "test" {
  name                  = "%s"
  project_id       = "my-project-1" 
  credentials_file = "LocalGWSetMap.json"
}

data "checkpoint_management_cme_accounts_gcp" "data_test"{
	name = "${checkpoint_management_cme_accounts_gcp.test.name}"
}
`, accountName)
}
