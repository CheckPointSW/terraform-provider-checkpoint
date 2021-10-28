package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementIdentityTag_basic(t *testing.T) {
	objName := "tfTestManagementIdentityTag_" + acctest.RandString(6)
	resourceName := "checkpoint_management_identity_tag.test"
	dataSourceName := "data.checkpoint_management_identity_tag.data_test"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementIdentityTagConfig(objName, "some external identifier"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "external-identifier", resourceName, "external-identifier"),
				),
			},
		},
	})
}

func testAccDataSourceManagementIdentityTagConfig(name string, externalIdentifier string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_identity_tag" "test" {
        name = "%s"
        external_identifier = "%s"
}

data "checkpoint_management_identity_tag" "data_test" {
        name = "${checkpoint_management_identity_tag.test.name}"
}
`, name, externalIdentifier)
}
