package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementCMEAccountsAzure_basic(t *testing.T) {
	resourceName := "checkpoint_management_cme_accounts_azure.test"
	dataSourceName := "data.checkpoint_management_cme_accounts_azure.data_test"
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
				Config: testAccDataSourceManagementCMEAccountsAzureConfig(accountName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "directory_id", resourceName, "directory_id"),
					resource.TestCheckResourceAttrPair(dataSourceName, "subscription", resourceName, "subscription"),
				),
			},
		},
	})
}

func testAccDataSourceManagementCMEAccountsAzureConfig(accountName string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_cme_accounts_azure" "test" {
  name                  = "%s"
  directory_id = "46707d92-02f4-4817-8116-a4c3b23e6266"
  application_id = "46707d92-02f4-4817-8116-a4c3b23e6266"
  client_secret = "mySecret"
  subscription = "46707d92-02f4-4817-8116-a4c3b23e6267"
}

data "checkpoint_management_cme_accounts_azure" "data_test"{
	name = "${checkpoint_management_cme_accounts_azure.test.name}"
}
`, accountName)
}
