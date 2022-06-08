package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	"strings"
	"testing"
)

func TestAccCheckpointManagementHttpsLayer_basic(t *testing.T) {

	var httpsLayerMap map[string]interface{}
	resourceName := "checkpoint_management_https_layer.test"
	objName := "tfTestManagementHttpsLayer_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementHttpsLayerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementHttpsLayerConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementHttpsLayerExists(resourceName, &httpsLayerMap),
					testAccCheckCheckpointManagementHttpsLayerAttributes(&httpsLayerMap, objName),
				),
			},
		},
	})
}

func testAccCheckpointManagementHttpsLayerDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_https_layer" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-https-layer", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("HttpsLayer object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementHttpsLayerExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("HttpsLayer ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-https-layer", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementHttpsLayerAttributes(httpsLayerMap *map[string]interface{}, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		httpsLayerName := (*httpsLayerMap)["name"].(string)
		if !strings.EqualFold(httpsLayerName, name) {
			return fmt.Errorf("name is %s, expected %s", name, httpsLayerName)
		}
		return nil
	}
}

func testAccManagementHttpsLayerConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_https_layer" "test" {
        name = "%s"
}
`, name)
}
