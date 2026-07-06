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

func TestAccCheckpointManagementDlpNextDataType_basic(t *testing.T) {

	var dlpNextDataTypeMap map[string]interface{}
	resourceName := "checkpoint_management_dlp_next_data_type.test"
	objName := "tfTestManagementDlpNextDataType_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementDlpNextDataTypeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementDlpNextDataTypeConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementDlpNextDataTypeExists(resourceName, &dlpNextDataTypeMap),
					testAccCheckCheckpointManagementDlpNextDataTypeAttributes(&dlpNextDataTypeMap, objName),
				),
			},
		},
	})
}

func testAccCheckpointManagementDlpNextDataTypeDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_dlp_next_data_type" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-dlp-next-data-type", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("DlpNextDataType object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementDlpNextDataTypeExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("DlpNextDataType ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-dlp-next-data-type", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementDlpNextDataTypeAttributes(dlpNextDataTypeMap *map[string]interface{}, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		dlpNextDataTypeName := (*dlpNextDataTypeMap)["name"].(string)
		if !strings.EqualFold(dlpNextDataTypeName, name) {
			return fmt.Errorf("name is %s, expected %s", name, dlpNextDataTypeName)
		}
		return nil
	}
}

func testAccManagementDlpNextDataTypeConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_dlp_next_data_type" "test" {
        name = "%s"
}
`, name)
}
