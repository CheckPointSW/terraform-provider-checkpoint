package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	"testing"
)

func TestAccCheckpointManagementCMEManagement_basic(t *testing.T) {
	var managementObject map[string]interface{}
	DefaultManagementName:=  "MGMT"
	resourceName := "checkpoint_management_cme_management.test"
	ManagementName := "test-management"

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
				Config: testAccManagementCMEManagementConfig(ManagementName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementCMEManagementExists(resourceName, &managementObject),
					testAccCheckCheckpointManagementCMEManagementAttributes(&managementObject, ManagementName, "localhost"),
				),
			},
			{
				Config: testAccManagementCMEManagementConfig(DefaultManagementName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementCMEManagementExists(resourceName, &managementObject),
					testAccCheckCheckpointManagementCMEManagementAttributes(&managementObject, DefaultManagementName,"localhost"),
				),
			},
		},
	})
}

func testAccManagementCMEManagementConfig(ManagementName string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_cme_management" "test" {
  name           = "%s"
}
`, ManagementName)
}

func testAccCheckCheckpointManagementCMEManagementExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)
		url := CmeApiPath + "/management"
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

func testAccCheckCheckpointManagementCMEManagementAttributes(managementObject *map[string]interface{}, ManagementName string, host string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		managementData := (*managementObject)["result"].(map[string]interface{})
		if managementData["name"] != ManagementName {
			return fmt.Errorf("management name is %s, expected %s", managementData["name"], ManagementName)
		}
		if managementData["host"] != host {
			return fmt.Errorf("host is %s, expected %s", managementData["host"], host)
		}
		return nil
	}
}
