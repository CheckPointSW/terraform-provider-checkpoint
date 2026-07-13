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

func TestAccCheckpointManagementGuideline_basic(t *testing.T) {

	var guidelineMap map[string]interface{}
	resourceName := "checkpoint_management_guideline.test"
	objName := "tfTestManagementGuideline_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementGuidelineDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementGuidelineConfig(objName, "Network"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementGuidelineExists(resourceName, &guidelineMap),
					testAccCheckCheckpointManagementGuidelineAttributes(&guidelineMap, objName, "Network"),
				),
			},
		},
	})
}

func testAccCheckpointManagementGuidelineDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_guideline" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-guideline", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("Guideline object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementGuidelineExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Guideline ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-guideline", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementGuidelineAttributes(guidelineMap *map[string]interface{}, name string, accessLayers1 string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		guidelineName := (*guidelineMap)["name"].(string)
		if !strings.EqualFold(guidelineName, name) {
			return fmt.Errorf("name is %s, expected %s", name, guidelineName)
		}
		accessLayersJson := (*guidelineMap)["access-layers"].([]interface{})
		var accessLayersIds = make([]string, 0)
		if len(accessLayersJson) > 0 {
			for _, accessLayers := range accessLayersJson {
				accessLayersTry1, ok := accessLayers.(map[string]interface{})
				if ok {
					accessLayersIds = append([]string{accessLayersTry1["name"].(string)}, accessLayersIds...)
				} else {
					accessLayersTry2 := accessLayers.(string)
					accessLayersIds = append([]string{accessLayersTry2}, accessLayersIds...)
				}
			}
		}

		GuidelineaccessLayers1 := accessLayersIds[0]
		if GuidelineaccessLayers1 != accessLayers1 {
			return fmt.Errorf("accessLayers1 is %s, expected %s", accessLayers1, GuidelineaccessLayers1)
		}
		return nil
	}
}

func testAccManagementGuidelineConfig(name string, accessLayers1 string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_group" "example" {
  name = "testGroup"
}
resource "checkpoint_management_guideline" "test" {
  name = "%s"
  access_layers {
    access_layer = "%s"
  }
  guideline_groups {
    name = checkpoint_management_group.example.name
    position {
      top = "top"
    }
  }
}
`, name, accessLayers1)
}
