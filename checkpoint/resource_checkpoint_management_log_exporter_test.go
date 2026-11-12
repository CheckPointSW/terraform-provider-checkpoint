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

func TestAccCheckpointManagementLogExporter_basic(t *testing.T) {

	var logExporterMap map[string]interface{}
	resourceName := "checkpoint_management_log_exporter.test"
	objName := "tfTestManagementLogExporter_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementLogExporterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementLogExporterConfig(objName, "1.2.3.4", 1234, "tcp"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementLogExporterExists(resourceName, &logExporterMap),
					testAccCheckCheckpointManagementLogExporterAttributes(&logExporterMap, objName, "1.2.3.4", 1234, "tcp"),
				),
			},
		},
	})
}

func testAccCheckpointManagementLogExporterDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_log_exporter" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-log-exporter", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success {
				return fmt.Errorf("LogExporter object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementLogExporterExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("LogExporter ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-log-exporter", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementLogExporterAttributes(logExporterMap *map[string]interface{}, name string, targetServer string, targetPort int, protocol string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		logExporterName := (*logExporterMap)["name"].(string)
		if !strings.EqualFold(logExporterName, name) {
			return fmt.Errorf("name is %s, expected %s", name, logExporterName)
		}
		logExporterTargetServer := (*logExporterMap)["target-server"].(string)
		if !strings.EqualFold(logExporterTargetServer, targetServer) {
			return fmt.Errorf("targetServer is %s, expected %s", targetServer, logExporterTargetServer)
		}
		logExporterTargetPort := int((*logExporterMap)["target-port"].(float64))
		if logExporterTargetPort != targetPort {
			return fmt.Errorf("targetPort is %d, expected %d", targetPort, logExporterTargetPort)
		}
		logExporterProtocol := (*logExporterMap)["protocol"].(string)
		if !strings.EqualFold(logExporterProtocol, protocol) {
			return fmt.Errorf("protocol is %s, expected %s", protocol, logExporterProtocol)
		}
		return nil
	}
}

func testAccManagementLogExporterConfig(name string, targetServer string, targetPort int, protocol string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_log_exporter" "test" {
        name = "%s"
        target_server = "%s"
        target_port = %d
        protocol = "%s"
}
`, name, targetServer, targetPort, protocol)
}
