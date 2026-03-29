package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementRepositoryScript_basic(t *testing.T) {

	objName := "tfTestManagementDataRepositoryScript_" + acctest.RandString(6)
	resourceName := "checkpoint_management_repository_script.repository_script"
	dataSourceName := "data.checkpoint_management_repository_script.data_repository_script"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementRepositoryScriptConfig(objName, "bHMgLWwgLw=="),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "color", resourceName, "color"),
					resource.TestCheckResourceAttrPair(dataSourceName, "comments", resourceName, "comments"),
					resource.TestCheckResourceAttrPair(dataSourceName, "script_body_base64", resourceName, "script_body"),
				),
			},
		},
	})

}

func testAccDataSourceManagementRepositoryScriptConfig(name string, scriptBody string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_repository_script" "repository_script" {
	name = "%s"
	script_body_base64 = "%s"
}

data "checkpoint_management_repository_script" "data_repository_script" {
    name = "${checkpoint_management_repository_script.repository_script.name}"
}
`, name, scriptBody)
}
