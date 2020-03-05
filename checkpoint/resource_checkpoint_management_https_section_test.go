package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform/helper/resource"
    "github.com/hashicorp/terraform/terraform"
    "os"
    "strings"
    "testing"
    "github.com/hashicorp/terraform/helper/acctest"
)

func TestAccCheckpointManagementHttpsSection_basic(t *testing.T) {

    var httpsSectionMap map[string]interface{}
    resourceName := "checkpoint_management_https_section.test"
    objName := "tfTestManagementHttpsSection_" + acctest.RandString(6)

    context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
        CheckDestroy: testAccCheckpointManagementHttpsSectionDestroy,
        Steps: []resource.TestStep{
            {
                Config: testAccManagementHttpsSectionConfig(objName, ),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckCheckpointManagementHttpsSectionExists(resourceName, &httpsSectionMap),
                    testAccCheckCheckpointManagementHttpsSectionAttributes(&httpsSectionMap, objName, ),
                ),
            },
        },
    })
}

func testAccCheckpointManagementHttpsSectionDestroy(s *terraform.State) error {

    client := testAccProvider.Meta().(*checkpoint.ApiClient)
    for _, rs := range s.RootModule().Resources {
        if rs.Type != "checkpoint_management_https_section" {
            continue
        }
        if rs.Primary.ID != "" {
            res, _ := client.ApiCall("show-https-section", map[string]interface{}{"uid": rs.Primary.ID, "layer": "New Layer 2"}, client.GetSessionID(), true, false)
            if res.Success {
                return fmt.Errorf("HttpsSection object (%s) still exists", rs.Primary.ID)
            }
        }
        return nil
    }
    return nil
}

func testAccCheckCheckpointManagementHttpsSectionExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        rs, ok := s.RootModule().Resources[resourceTfName]
        if !ok {
            return fmt.Errorf("Resource not found: %s", resourceTfName)
        }

        if rs.Primary.ID == "" {
            return fmt.Errorf("HttpsSection ID is not set")
        }

        client := testAccProvider.Meta().(*checkpoint.ApiClient)

        response, err := client.ApiCall("show-https-section", map[string]interface{}{"uid": rs.Primary.ID, "layer": "New Layer 2"}, client.GetSessionID(), true, false)
        if !response.Success {
            return err
        }

        *res = response.GetData()

        return nil
    }
}

func testAccCheckCheckpointManagementHttpsSectionAttributes(httpsSectionMap *map[string]interface{}, name string) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        httpsSectionName := (*httpsSectionMap)["name"].(string)
        if !strings.EqualFold(httpsSectionName, name) {
            return fmt.Errorf("name is %s, expected %s", name, httpsSectionName)
        }
        return nil
    }
}

func testAccManagementHttpsSectionConfig(name string) string {
    return fmt.Sprintf(`
resource "checkpoint_management_https_section" "test" {
        name = "%s"
        position = {top = "top"}
        layer = "New Layer 2"
}
`, name)
}