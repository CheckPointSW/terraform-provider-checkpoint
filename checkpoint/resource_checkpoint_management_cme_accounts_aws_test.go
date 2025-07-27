package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	"reflect"
	"testing"
)

func TestAccCheckpointManagementCMEAccountsAWS_basic(t *testing.T) {
	var awsAccount map[string]interface{}
	resourceName := "checkpoint_management_cme_accounts_aws.test"
	accountName := "test-account"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this test")
	} else if context != "web_api" {
		t.Skip("Skipping cme api test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementCMEAccountAWSDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementCMEAccountsAWSConfig(accountName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementCMEAccountsAWSExists(resourceName, &awsAccount),
					testAccCheckCheckpointManagementCMEAccountsAWSAttributes(&awsAccount, accountName, []interface{}{"us-east-1"},
						"IAM", true, true, true, true,
						[]map[string]interface{}{{"name": "sub_account_a", "access_key": "ABCDEAAAAAAAAAAAHAA", "secret_key": "1",
							"sts_role": "arn:aws:iam::123456789012:role/role-name", "sts_external_id": "xyzx"}}, 0),
				),
			},
		},
	})
}
func testAccCheckpointManagementCMEAccountAWSDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_cme_accounts_aws" {
			continue
		}
		if rs.Primary.ID != "" {
			url := CmeApiPath + "/accounts/" + rs.Primary.Attributes["name"]
			response, err := client.ApiCall(url, nil, client.GetSessionID(), true, client.IsProxyUsed(), "GET")
			if err != nil {
				return err
			}
			res := response.GetData()
			if !checkIfRequestFailed(res) {
				return fmt.Errorf("AWS account (%s) still exists", rs.Primary.Attributes["name"])
			}
		}
		return nil
	}
	return nil
}

func testAccManagementCMEAccountsAWSConfig(accountName string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_cme_accounts_aws" "test" {
  name                  = "%s"
  regions               = ["us-east-1"]
  credentials_file      = "IAM"
  scan_vpn              = true
  scan_load_balancers   = true
  scan_subnets			= true
  scan_subnets_6		= true
  sub_accounts {
    name       = "sub_account_a"
    access_key = "ABCDEAAAAAAAAAAAHAA"
    secret_key = "aaaaaaaaaaaaaaaaee1"
	sts_role = "arn:aws:iam::123456789012:role/role-name"
	sts_external_id = "xyzx"
  }
}
`, accountName)
}

func testAccCheckCheckpointManagementCMEAccountsAWSExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)
		url := CmeApiPath + "/accounts/" + rs.Primary.Attributes["name"]
		response, err := client.ApiCall(url, nil, client.GetSessionID(), true, client.IsProxyUsed(), "GET")
		if err != nil {
			return err
		}

		*res = response.GetData()
		if checkIfRequestFailed(*res) {
			errMessage := buildErrorMessage(*res)
			return fmt.Errorf(errMessage)
		}
		return nil
	}
}

func testAccCheckCheckpointManagementCMEAccountsAWSAttributes(awsAccount *map[string]interface{}, name string, regions []interface{},
	credFile string, scanVpn bool, scanLoadBalancers bool, scanSubnets bool, scanSubnets6 bool, subAccounts []map[string]interface{}, expectedDeletionTolerance int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		account := (*awsAccount)["result"].(map[string]interface{})
		if account["name"] != name {
			return fmt.Errorf("name is %s, expected %s", account["name"], name)
		}
		if !reflect.DeepEqual(account["regions"], regions) {
			return fmt.Errorf("regions are %v, expected %v", account["regions"], regions)
		}
		if account["credentials_file"] != credFile {
			return fmt.Errorf("credentials_file is %s, expected %s", account["credentials_file"], credFile)
		}
		deletionTolerance := int(account["deletion_tolerance"].(float64))
		if deletionTolerance != expectedDeletionTolerance {
			return fmt.Errorf("deletion_tolerance is %d, expected %d", deletionTolerance, expectedDeletionTolerance)
		}
		vpnFlag := account["sync"].(map[string]interface{})["scan_vpn"]
		if vpnFlag != scanVpn {
			return fmt.Errorf("scan_vpn is %t, expected %t", vpnFlag, scanVpn)
		}
		lbFlag := account["sync"].(map[string]interface{})["scan_load_balancers"]
		if lbFlag != scanLoadBalancers {
			return fmt.Errorf("scan_load_balancers is %t, expected %t", lbFlag, scanLoadBalancers)
		}
		if scanSubnets != account["sync"].(map[string]interface{})["scan_subnets"] {
			return fmt.Errorf("scan_subnets is %t, expected %t", account["sync"].(map[string]interface{})["scan-subnets"], scanSubnets)
		}
		if scanSubnets6 != account["sync"].(map[string]interface{})["scan_subnets_6"] {
			return fmt.Errorf("scan_subnets_6 is %t, expected %t", account["sync"].(map[string]interface{})["scan-subnets"], scanSubnets6)
		}
		subAccountsMap := account["sub_accounts"].(map[string]interface{})
		if len(subAccountsMap) != len(subAccounts) {
			return fmt.Errorf("sub accounts list length is %d, expected %d", len(subAccountsMap), len(subAccounts))
		}
		for key, value := range subAccountsMap {
			subAccountMap := value.(map[string]interface{})
			subAccountInput := subAccounts[0]
			if key != subAccountInput["name"] {
				return fmt.Errorf("sub account name is %s, expected %s", key, subAccountInput["name"])
			}
			if subAccountMap["access_key"] != subAccountInput["access_key"] {
				return fmt.Errorf("sub account access key is %s, expected %s", subAccountMap["access_key"], subAccountInput["access_key"])
			}
			if subAccountMap["sts_role"] != subAccountInput["sts_role"] {
				return fmt.Errorf("sub account sts role is %s, expected %s", subAccountMap["sts_role"], subAccountInput["sts_role"])
			}
			if subAccountMap["sts_external_id"] != subAccountInput["sts_external_id"] {
				return fmt.Errorf("sub account sts external id is %s, expected %s", subAccountMap["sts_external_id"], subAccountInput["sts_external_id"])
			}
		}
		return nil
	}
}
