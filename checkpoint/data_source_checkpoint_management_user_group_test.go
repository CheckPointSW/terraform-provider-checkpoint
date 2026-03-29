package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementUserGroup_basic(t *testing.T) {
	objName := "UserGroup" + acctest.RandString(2)
	resourceName := "checkpoint_management_user_group.user_group"
	dataSourceName := "data.checkpoint_management_user_group.test_user_group"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementUserGroupConfig(objName, "myuser@email.com"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "email", resourceName, "email"),
				),
			},
		},
	})
}

func testAccDataSourceManagementUserGroupConfig(name string, email string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_user_group" "user_group" {
        name = "%s"
		email = "%s"
}

data "checkpoint_management_user_group" "test_user_group" {
    name = "${checkpoint_management_user_group.user_group.name}"
}
`, name, email)
}
