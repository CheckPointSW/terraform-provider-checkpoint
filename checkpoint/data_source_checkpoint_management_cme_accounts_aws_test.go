package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementCMEAccountsAWS_basic(t *testing.T) {
	resourceName := "checkpoint_management_cme_accounts_aws.test"
	dataSourceName := "data.checkpoint_management_cme_accounts_aws.data_test"
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
				Config: testAccDataSourceManagementCMEAccountsAWSConfig(accountName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "regions", resourceName, "regions"),
					resource.TestCheckResourceAttrPair(dataSourceName, "credentials_file", resourceName, "credentials_file"),
					resource.TestCheckResourceAttrPair(dataSourceName, "scan_subnets", resourceName, "scan_subnets"),
					resource.TestCheckResourceAttrPair(dataSourceName, "scan_subnets_6", resourceName, "scan_subnets_6"),
				),
			},
		},
	})
}

func testAccDataSourceManagementCMEAccountsAWSConfig(accountName string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_cme_accounts_aws" "test" {
  name                  = "%s"
  regions               = ["us-east-1"]
  credentials_file      = "IAM"
  scan_subnets			= true
  scan_subnets_6		= true

}

data "checkpoint_management_cme_accounts_aws" "data_test"{
	name = "${checkpoint_management_cme_accounts_aws.test.name}"
}
`, accountName)
}
