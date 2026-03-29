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

func TestAccCheckpointManagementAwsDataCenterServer_basic(t *testing.T) {

	var awsDataCenterServerMap map[string]interface{}
	resourceName := "checkpoint_management_aws_data_center_server.test"
	objName := "tfTestManagementAwsDataCenterServer_" + acctest.RandString(6)
	authenticationMethod := "user-authentication"
	accessKeyId := "MY-KEY-ID"
	secretAccessKey := "MY-SECRET-KEY"
	region := "us-east-1"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementAwsDataCenterServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementAwsDataCenterServerConfig(objName, authenticationMethod, accessKeyId, secretAccessKey, region),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementAwsDataCenterServerExists(resourceName, &awsDataCenterServerMap),
					testAccCheckCheckpointManagementAwsDataCenterServerAttributes(&awsDataCenterServerMap, objName),
				),
			},
		},
	})
}

func testAccCheckpointManagementAwsDataCenterServerDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_aws_data_center_server" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-data-center-server", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("AwsDataCenterServer object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementAwsDataCenterServerExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("AwsDataCenterServer ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-data-center-server", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementAwsDataCenterServerAttributes(awsDataCenterServerMap *map[string]interface{}, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		awsDataCenterServerName := (*awsDataCenterServerMap)["name"].(string)
		if !strings.EqualFold(awsDataCenterServerName, name) {
			return fmt.Errorf("name is %s, expected %s", name, awsDataCenterServerName)
		}
		return nil
	}
}

func testAccManagementAwsDataCenterServerConfig(name string, authenticationMethod string, accessKeyId string, secretAccessKey string, region string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_aws_data_center_server" "test" {
    name = "%s"
	authentication_method = "%s"
	access_key_id = "%s"
	secret_access_key = "%s"
    region = "%s"
	ignore_warnings = true
}
`, name, authenticationMethod, accessKeyId, secretAccessKey, region)
}
