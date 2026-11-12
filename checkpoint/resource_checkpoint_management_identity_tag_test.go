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

func TestAccCheckpointManagementIdentityTag_basic(t *testing.T) {

	var identityTagMap map[string]interface{}
	resourceName := "checkpoint_management_identity_tag.test"
	objName := "tfTestManagementIdentityTag_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementIdentityTagDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementIdentityTagConfig(objName, "some external identifier"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementIdentityTagExists(resourceName, &identityTagMap),
					testAccCheckCheckpointManagementIdentityTagAttributes(&identityTagMap, objName, "some external identifier"),
				),
			},
		},
	})
}

func testAccCheckpointManagementIdentityTagDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_identity_tag" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-identity-tag", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("IdentityTag object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementIdentityTagExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("IdentityTag ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-identity-tag", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementIdentityTagAttributes(identityTagMap *map[string]interface{}, name string, externalIdentifier string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		identityTagName := (*identityTagMap)["name"].(string)
		if !strings.EqualFold(identityTagName, name) {
			return fmt.Errorf("name is %s, expected %s", name, identityTagName)
		}
		identityTagExternalIdentifier := (*identityTagMap)["external-identifier"].(string)
		if !strings.EqualFold(identityTagExternalIdentifier, externalIdentifier) {
			return fmt.Errorf("externalIdentifier is %s, expected %s", externalIdentifier, identityTagExternalIdentifier)
		}
		return nil
	}
}

func testAccManagementIdentityTagConfig(name string, externalIdentifier string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_identity_tag" "test" {
        name = "%s"
        external_identifier = "%s"
}
`, name, externalIdentifier)
}
