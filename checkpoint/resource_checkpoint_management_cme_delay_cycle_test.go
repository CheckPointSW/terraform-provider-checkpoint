package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	"testing"
)

func TestAccCheckpointManagementCMEDelayCycle_basic(t *testing.T) {
	var delayCycleObject map[string]interface{}
	DefaultDelayCycle:=  30
	resourceName := "checkpoint_management_cme_delay_cycle.test"
	delayCycleVal := 20

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
				Config: testAccManagementCMEDelayCycleConfig(delayCycleVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementCMEDelayCycleExists(resourceName, &delayCycleObject),
					testAccCheckCheckpointManagementCMEDelayCycleAttributes(&delayCycleObject, delayCycleVal),
				),
			},
			{
				Config: testAccManagementCMEDelayCycleConfig(DefaultDelayCycle),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementCMEDelayCycleExists(resourceName, &delayCycleObject),
					testAccCheckCheckpointManagementCMEDelayCycleAttributes(&delayCycleObject, DefaultDelayCycle),
				),
			},
		},
	})
}

func testAccManagementCMEDelayCycleConfig(delayCycleVal int) string {
	return fmt.Sprintf(`
resource "checkpoint_management_cme_delay_cycle" "test" {
  delay_cycle           = "%v"
}
`, delayCycleVal)
}

func testAccCheckCheckpointManagementCMEDelayCycleExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)
		url := CmeApiPath + "/generalConfiguration/delayCycle"
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

func testAccCheckCheckpointManagementCMEDelayCycleAttributes(delayCycleObject *map[string]interface{}, delayCycleVal int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		DelayCycleData := (*delayCycleObject)["result"].(map[string]interface{})
		if int(DelayCycleData["delay_cycle"].(float64)) != delayCycleVal {
			return fmt.Errorf("delay cycle value is %v, expected %v", DelayCycleData["delay_cycle"], delayCycleVal)
		}
		return nil
	}
}
