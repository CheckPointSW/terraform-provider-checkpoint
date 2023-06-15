package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	"strings"
	"testing"
)

func TestAccCheckpointManagementServerCertificate_basic(t *testing.T) {

	var serverCertificateMap map[string]interface{}

	resourceName := "checkpoint_management_server_certificate.test"
	objName := "tfTestManagementServerCertificate" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementServerCertificateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementServerCertificateConfig(objName, "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", "bXlfcGFzc3dvcmQ=", "this is a comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementServerCertificateExists(resourceName, &serverCertificateMap),
					testAccCheckCheckpointManagementServerCertificateAttributes(&serverCertificateMap, objName, "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", "bXlfcGFzc3dvcmQ=", "this is a comment"),
				),
			},
		},
	})
}

func testAccCheckpointManagementServerCertificateDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_server_certificate" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-server-certificate", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success {
				return fmt.Errorf("ServerCertificate object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementServerCertificateExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ServerCertificate ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-server-certificate", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementServerCertificateAttributes(serverCertificateMap *map[string]interface{}, name string, base64Certificate string, base64Password string, comments string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		serverCertificateName := (*serverCertificateMap)["name"].(string)
		if !strings.EqualFold(serverCertificateName, name) {
			return fmt.Errorf("name is %s, expected %s", name, serverCertificateName)
		}

		return nil
	}
}

func testAccManagementServerCertificateConfig(name string, base64Certificate string, base64Password string, comments string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_server_certificate" "test" {
        name = "%s"
        base64_certificate = "%s"
        base64_password = "%s"
        comments = "%s"
}
`, name, base64Certificate, base64Password, comments)
}
