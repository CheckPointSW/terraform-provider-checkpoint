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

func TestAccCheckpointManagementServiceCitrixTcp_basic(t *testing.T) {

	var serviceCitrixTcpMap map[string]interface{}
	resourceName := "checkpoint_management_service_citrix_tcp.test"
	objName := "tfTestManagementServiceCitrixTcp_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementServiceCitrixTcpDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementServiceCitrixTcpConfig(objName, "my citrix application"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementServiceCitrixTcpExists(resourceName, &serviceCitrixTcpMap),
					testAccCheckCheckpointManagementServiceCitrixTcpAttributes(&serviceCitrixTcpMap, objName, "my citrix application"),
				),
			},
		},
	})
}

func testAccCheckpointManagementServiceCitrixTcpDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_service_citrix_tcp" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-service-citrix-tcp", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("ServiceCitrixTcp object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementServiceCitrixTcpExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ServiceCitrixTcp ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-service-citrix-tcp", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementServiceCitrixTcpAttributes(serviceCitrixTcpMap *map[string]interface{}, name string, application string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		serviceCitrixTcpName := (*serviceCitrixTcpMap)["name"].(string)
		if !strings.EqualFold(serviceCitrixTcpName, name) {
			return fmt.Errorf("name is %s, expected %s", name, serviceCitrixTcpName)
		}
		serviceCitrixTcpApplication := (*serviceCitrixTcpMap)["application"].(string)
		if !strings.EqualFold(serviceCitrixTcpApplication, application) {
			return fmt.Errorf("application is %s, expected %s", application, serviceCitrixTcpApplication)
		}
		return nil
	}
}

func testAccManagementServiceCitrixTcpConfig(name string, application string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_service_citrix_tcp" "test" {
        name = "%s"
        application = "%s"
}
`, name, application)
}
