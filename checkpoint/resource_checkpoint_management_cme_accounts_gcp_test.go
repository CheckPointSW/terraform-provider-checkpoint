package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	"testing"
)

func TestAccCheckpointManagementCMEAccountsGCP_basic(t *testing.T) {
	var gcpAccount map[string]interface{}
	resourceName := "checkpoint_management_cme_accounts_gcp.test"
	accountName := "test-account"
	projectId := "my-project-1"
	credentialsFile := "LocalGWSetMap.json"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this test")
	} else if context != "web_api" {
		t.Skip("Skipping cme api test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementCMEAccountGCPDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementCMEAccountsGCPConfig(accountName, projectId, credentialsFile),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementCMEAccountsGCPExists(resourceName, &gcpAccount),
					testAccCheckCheckpointManagementCMEAccountsGCPAttributes(&gcpAccount, accountName, projectId, 0),
				),
			},
		},
	})
}
func testAccCheckpointManagementCMEAccountGCPDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_cme_accounts_gcp" {
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
				return fmt.Errorf("GCP account (%s) still exists", rs.Primary.Attributes["name"])
			}
		}
		return nil
	}
	return nil
}

func testAccManagementCMEAccountsGCPConfig(accountName string, projectId string, credentialsFile string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_cme_accounts_gcp" "test" {
  name           = "%s"
  project_id   = "%s"
  credentials_file = "%s"
}
`, accountName, projectId, credentialsFile)
}

func testAccCheckCheckpointManagementCMEAccountsGCPExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
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

func testAccCheckCheckpointManagementCMEAccountsGCPAttributes(gcpAccount *map[string]interface{}, name string,
	projectId string, expectedDeletionTolerance int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		account := (*gcpAccount)["result"].(map[string]interface{})
		if account["name"] != name {
			return fmt.Errorf("name is %s, expected %s", account["name"], name)
		}
		if account["project_id"] != projectId {
			return fmt.Errorf("project_id is %s, expected %s", account["project_id"], projectId)
		}
		deletionTolerance := int(account["deletion_tolerance"].(float64))
		if deletionTolerance != expectedDeletionTolerance {
			return fmt.Errorf("deletion_tolerance is %d, expected %d", deletionTolerance, expectedDeletionTolerance)
		}
		return nil
	}
}
