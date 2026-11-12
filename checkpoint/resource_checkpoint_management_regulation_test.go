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

func TestAccCheckpointManagementRegulation_basic(t *testing.T) {

	var regulationMap map[string]interface{}
	resourceName := "checkpoint_management_regulation.test"
	objName := "tfTestManagementRegulation_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementRegulationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementRegulationConfig(objName, "my new regulation", "my compliance regulation"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementRegulationExists(resourceName, &regulationMap),
					testAccCheckCheckpointManagementRegulationAttributes(&regulationMap, objName, "my new regulation", "my compliance regulation"),
				),
			},
		},
	})
}

func testAccCheckpointManagementRegulationDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_regulation" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-regulation", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("Regulation object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementRegulationExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Regulation ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-regulation", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementRegulationAttributes(regulationMap *map[string]interface{}, name string, fullName string, comments string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		regulationName := (*regulationMap)["name"].(string)
		if !strings.EqualFold(regulationName, name) {
			return fmt.Errorf("name is %s, expected %s", name, regulationName)
		}
		regulationFullName := (*regulationMap)["full-name"].(string)
		if !strings.EqualFold(regulationFullName, fullName) {
			return fmt.Errorf("fullName is %s, expected %s", fullName, regulationFullName)
		}
		regulationComments := (*regulationMap)["comments"].(string)
		if !strings.EqualFold(regulationComments, comments) {
			return fmt.Errorf("comments is %s, expected %s", comments, regulationComments)
		}
		return nil
	}
}

func testAccManagementRegulationConfig(name string, fullName string, comments string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_regulation" "test" {
        name = "%s"
        full_name = "%s"
        comments = "%s"
}
`, name, fullName, comments)
}
