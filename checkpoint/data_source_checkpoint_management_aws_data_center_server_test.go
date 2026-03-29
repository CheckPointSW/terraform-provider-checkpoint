package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementAwsDataCenterServer_basic(t *testing.T) {

	objName := "tfTestManagementDataAwsDataCenterServer_" + acctest.RandString(6)
	resourceName := "checkpoint_management_aws_data_center_server.aws_data_center_server"
	dataSourceName := "data.checkpoint_management_aws_data_center_server.aws_data_center_server"
	authenticationMethod := "user-authentication"
	accessKeyId := "MY-KEY-ID"
	secretAccessKey := "MY-SECRET-KEY"
	region := "us-east-1"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementAwsDataCenterServerConfig(objName, authenticationMethod, accessKeyId, secretAccessKey, region),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementAwsDataCenterServerConfig(name string, authenticationMethod string, accessKeyId string, secretAccessKey string, region string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_aws_data_center_server" "aws_data_center_server" {
    name = "%s"
	authentication_method = "%s"
	access_key_id = "%s"
	secret_access_key = "%s"
    region = "%s"
	ignore_warnings = true
}

data "checkpoint_management_aws_data_center_server" "aws_data_center_server" {
    name = "${checkpoint_management_aws_data_center_server.aws_data_center_server.name}"
}
`, name, authenticationMethod, accessKeyId, secretAccessKey, region)
}
