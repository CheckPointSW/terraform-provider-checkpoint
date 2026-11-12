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

func TestAccCheckpointManagementRepositoryScript_basic(t *testing.T) {

	var repositoryScriptMap map[string]interface{}
	resourceName := "checkpoint_management_repository_script.test"
	objName := "tfTestManagementRepositoryScript_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementRepositoryScriptDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementRepositoryScriptConfig(objName, "bHMgLWwgLw=="),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementRepositoryScriptExists(resourceName, &repositoryScriptMap),
					testAccCheckCheckpointManagementRepositoryScriptAttributes(&repositoryScriptMap, objName, "bHMgLWwgLw=="),
				),
			},
		},
	})
}

func testAccCheckpointManagementRepositoryScriptDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_repository_script" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-repository-script", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("RepositoryScript object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementRepositoryScriptExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("RepositoryScript ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-repository-script", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}

		*res = response.GetData()
		return nil
	}
}

func testAccCheckCheckpointManagementRepositoryScriptAttributes(repositoryScriptMap *map[string]interface{}, name string, scriptBody string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		repositoryScriptName := (*repositoryScriptMap)["name"].(string)
		if !strings.EqualFold(repositoryScriptName, name) {
			return fmt.Errorf("name is %s, expected %s", name, repositoryScriptName)
		}
		repositoryScriptScriptBody := (*repositoryScriptMap)["script-body"].(string)
		if !strings.EqualFold(repositoryScriptScriptBody, scriptBody) {
			return fmt.Errorf("scriptBody is %s, expected %s", scriptBody, repositoryScriptScriptBody)
		}
		return nil
	}
}

func testAccManagementRepositoryScriptConfig(name string, scriptBody string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_repository_script" "test" {
        name = "%s"
        script_body_base64 = "%s"
}
`, name, scriptBody)
}
