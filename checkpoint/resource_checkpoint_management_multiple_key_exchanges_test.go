package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"os"
	"strings"
	"testing"
)

func TestAccCheckpointManagementMultipleKeyExchanges_basic(t *testing.T) {

	var multipleKeyExchangesMap map[string]interface{}
	resourceName := "checkpoint_management_multiple_key_exchanges.test"
	objName := "tfTestManagementMultipleKeyExchanges_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementMultipleKeyExchangesDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementMultipleKeyExchangesConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementMultipleKeyExchangesExists(resourceName, &multipleKeyExchangesMap),
					testAccCheckCheckpointManagementMultipleKeyExchangesAttributes(&multipleKeyExchangesMap, objName),
				),
			},
		},
	})
}

func testAccCheckpointManagementMultipleKeyExchangesDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_multiple_key_exchanges" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-multiple-key-exchanges", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success {
				return fmt.Errorf("MultipleKeyExchanges object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementMultipleKeyExchangesExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("MultipleKeyExchanges ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-multiple-key-exchanges", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementMultipleKeyExchangesAttributes(multipleKeyExchangesMap *map[string]interface{}, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		multipleKeyExchangesName := (*multipleKeyExchangesMap)["name"].(string)
		if !strings.EqualFold(multipleKeyExchangesName, name) {
			return fmt.Errorf("name is %s, expected %s", name, multipleKeyExchangesName)
		}
		return nil
	}
}

func testAccManagementMultipleKeyExchangesConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_multiple_key_exchanges" "test" {
        name = "%s"
        key_exchange_methods = ["group-2"] 
        additional_key_exchange_1_methods =  ["kyber-1024"]
}
`, name)
}
