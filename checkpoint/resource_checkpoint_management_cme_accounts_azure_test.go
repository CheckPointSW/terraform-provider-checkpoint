package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	"testing"
)

func TestAccCheckpointManagementCMEAccountsAzure_basic(t *testing.T) {
	var azureAccount map[string]interface{}
	resourceName := "checkpoint_management_cme_accounts_azure.test"
	accountName := "test-account"
	directoryId := "46707d92-02f4-4817-8116-a4c3b23e6266"
	applicationId := "46707d92-02f4-4817-8116-a4c3b23e6266"
	clientSecret := "1234abcdefgh----"
	subscription := "46707d92-02f4-4817-8116-a4c3b23e6267"
	environment := "AzureCloud"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this test")
	} else if context != "web_api" {
		t.Skip("Skipping cme api test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementCMEAccountAzureDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementCMEAccountsAzureConfig(accountName, directoryId, applicationId, clientSecret, subscription, environment),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementCMEAccountsAzureExists(resourceName, &azureAccount),
					testAccCheckCheckpointManagementCMEAccountsAzureAttributes(&azureAccount, accountName, directoryId, applicationId,
						subscription, 3, environment),
				),
			},
		},
	})
}
func testAccCheckpointManagementCMEAccountAzureDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_cme_accounts_azure" {
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
				return fmt.Errorf("Azure account (%s) still exists", rs.Primary.Attributes["name"])
			}
		}
		return nil
	}
	return nil
}

func testAccManagementCMEAccountsAzureConfig(accountName string, directoryId string, applicationId string, clientSecret string, subscription string, environment string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_cme_accounts_azure" "test" {
  name           = "%s"
  directory_id   = "%s"
  application_id = "%s"
  client_secret  = "%s"
  subscription   = "%s"
  environment    = "%s"
}
`, accountName, directoryId, applicationId, clientSecret, subscription, environment)
}

func testAccCheckCheckpointManagementCMEAccountsAzureExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
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

func testAccCheckCheckpointManagementCMEAccountsAzureAttributes(azureAccount *map[string]interface{}, name string,
	directoryId string, applicationId string, subscription string, expectedDeletionTolerance int, environment string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		account := (*azureAccount)["result"].(map[string]interface{})
		if account["name"] != name {
			return fmt.Errorf("name is %s, expected %s", account["name"], name)
		}
		if account["directory_id"] != directoryId {
			return fmt.Errorf("directory_id is %s, expected %s", account["directory_id"], directoryId)
		}
		if account["application_id"] != applicationId {
			return fmt.Errorf("application_id is %s, expected %s", account["application_id"], applicationId)
		}
		if account["subscription"] != subscription {
			return fmt.Errorf("subscription is %s, expected %s", account["subscription"], subscription)
		}
		deletionTolerance := int(account["deletion_tolerance"].(float64))
		if deletionTolerance != expectedDeletionTolerance {
			return fmt.Errorf("deletion_tolerance is %d, expected %d", deletionTolerance, expectedDeletionTolerance)
		}
		if account["environment"] != environment {
			return fmt.Errorf("environment is %s, expected %s", account["environment"], environment)
		}
		return nil
	}
}
