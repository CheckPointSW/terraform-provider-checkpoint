package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	"testing"
)

func TestAccCheckpointManagementPackage_basic(t *testing.T) {
	var packageMap map[string]interface{}
	resourceName := "checkpoint_management_package.test"
	objName := "tfTestManagementPackage_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementPackageDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementPackageConfig(objName), //runs "terraform apply"
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementPackageExists(resourceName, &packageMap),
					testAccCheckCheckpointManagementPackageAttributes(&packageMap, objName),
				),
			},
		},
	})

}

func testAccCheckpointManagementPackageDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_package" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-package", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("package object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementPackageExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("Package ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-package", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementPackageAttributes(packageMap *map[string]interface{}, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		packageName := (*packageMap)["name"].(string)
		if packageName != name {
			return fmt.Errorf("name is %s, expected %s", packageName, name)
		}

		return nil

	}
}

func testAccManagementPackageConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_package" "test" {
    name = "%s"
}
`, name)
}
