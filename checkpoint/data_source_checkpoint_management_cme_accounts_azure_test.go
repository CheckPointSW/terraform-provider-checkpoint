package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
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
					resource.TestCheckResourceAttrPair(dataSourceName, "application_id", resourceName, "application_id"),
					resource.TestCheckResourceAttrPair(dataSourceName, "environment", resourceName, "environment"),
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
  client_secret = "1234abcdefgh----"
  subscription = "46707d92-02f4-4817-8116-a4c3b23e6267"
  environment = "AzureCloud"
}

data "checkpoint_management_cme_accounts_azure" "data_test"{
	name = "${checkpoint_management_cme_accounts_azure.test.name}"
}
`, accountName)
}
