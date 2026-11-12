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

func TestAccCheckpointManagementThreatLayer_basic(t *testing.T) {

	var threatLayerMap map[string]interface{}
	resourceName := "checkpoint_management_threat_layer.test"
	objName := "tfTestManagementThreatLayer_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementThreatLayerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementThreatLayerConfig(objName, "blue"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementThreatLayerExists(resourceName, &threatLayerMap),
					testAccCheckCheckpointManagementThreatLayerAttributes(&threatLayerMap, objName, "blue"),
				),
			},
		},
	})
}

func testAccCheckpointManagementThreatLayerDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_threat_layer" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-threat-layer", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success {
				return fmt.Errorf("Threat Layer object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementThreatLayerExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Threat Layer ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-threat-layer", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementThreatLayerAttributes(threatLayerMap *map[string]interface{}, objName string, color string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		threatLayerName := (*threatLayerMap)["name"].(string)
		if !strings.EqualFold(threatLayerName, objName) {
			return fmt.Errorf("name is %s, expected %s", objName, threatLayerName)
		}
		threatLayerColor := (*threatLayerMap)["color"].(string)
		if threatLayerColor != color {
			return fmt.Errorf("color is %s, expected %s", threatLayerColor, color)
		}

		return nil
	}
}

func testAccManagementThreatLayerConfig(objName string, color string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_threat_layer" "test" {
        name = "%s"
        color = "%s"
}
`, objName, color)
}
