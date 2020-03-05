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

func TestAccCheckpointManagementGroupWithExclusion_basic(t *testing.T) {

    var groupWithExclusionMap map[string]interface{}
    resourceName := "checkpoint_management_group_with_exclusion.test"
    objName := "tfTestManagementGroupWithExclusion_" + acctest.RandString(6)

    context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
        CheckDestroy: testAccCheckpointManagementGroupWithExclusionDestroy,
        Steps: []resource.TestStep{
            {
                Config: testAccManagementGroupWithExclusionConfig(objName, "new group 1", "new group 2"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckCheckpointManagementGroupWithExclusionExists(resourceName, &groupWithExclusionMap),
                    testAccCheckCheckpointManagementGroupWithExclusionAttributes(&groupWithExclusionMap, objName, "new group 1", "new group 2"),
                ),
            },
        },
    })
}

func testAccCheckpointManagementGroupWithExclusionDestroy(s *terraform.State) error {

    client := testAccProvider.Meta().(*checkpoint.ApiClient)
    for _, rs := range s.RootModule().Resources {
        if rs.Type != "checkpoint_management_group_with_exclusion" {
            continue
        }
        if rs.Primary.ID != "" {
            res, _ := client.ApiCall("show-group-with-exclusion", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
            if res.Success {
                return fmt.Errorf("GroupWithExclusion object (%s) still exists", rs.Primary.ID)
            }
        }
        return nil
    }
    return nil
}

func testAccCheckCheckpointManagementGroupWithExclusionExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        rs, ok := s.RootModule().Resources[resourceTfName]
        if !ok {
            return fmt.Errorf("Resource not found: %s", resourceTfName)
        }

        if rs.Primary.ID == "" {
            return fmt.Errorf("GroupWithExclusion ID is not set")
        }

        client := testAccProvider.Meta().(*checkpoint.ApiClient)

        response, err := client.ApiCall("show-group-with-exclusion", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
        if !response.Success {
            return err
        }

        *res = response.GetData()

        return nil
    }
}

func testAccCheckCheckpointManagementGroupWithExclusionAttributes(groupWithExclusionMap *map[string]interface{}, name string, include string, except string) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        groupWithExclusionName := (*groupWithExclusionMap)["name"].(string)
        if !strings.EqualFold(groupWithExclusionName, name) {
            return fmt.Errorf("name is %s, expected %s", name, groupWithExclusionName)
        }
        groupWithExclusionInclude := (*groupWithExclusionMap)["include"].(map[string]interface{})
        if groupWithExclusionInclude["name"] != include {
            return fmt.Errorf("include is %s, expected %s", include, groupWithExclusionInclude)
        }
        groupWithExclusionExcept := (*groupWithExclusionMap)["except"].(map[string]interface{})
        if groupWithExclusionExcept["name"] != except {
            return fmt.Errorf("except is %s, expected %s", except, groupWithExclusionExcept)
        }
        return nil
    }
}

func testAccManagementGroupWithExclusionConfig(name string, include string, except string) string {
    return fmt.Sprintf(`
resource "checkpoint_management_group_with_exclusion" "test" {
        name = "%s"
        include = "%s"
        except = "%s"
}
`, name, include, except)
}

