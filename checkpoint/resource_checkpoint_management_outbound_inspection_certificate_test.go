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

func TestAccCheckpointManagementOutboundInspectionCertificate_basic(t *testing.T) {

	var outboundInspectionCertificateMap map[string]interface{}
	resourceName := "checkpoint_management_outbound_inspection_certificate.test"
	objName := "tfTestManagementOutboundInspectionCertificate_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementOutboundInspectionCertificateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementOutboundInspectionCertificateConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementOutboundInspectionCertificateExists(resourceName, &outboundInspectionCertificateMap),
					testAccCheckCheckpointManagementOutboundInspectionCertificateAttributes(&outboundInspectionCertificateMap, objName),
				),
			},
		},
	})
}

func testAccCheckpointManagementOutboundInspectionCertificateDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_outbound_inspection_certificate" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-outbound-inspection-certificate", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success {
				return fmt.Errorf("OutboundInspectionCertificate object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementOutboundInspectionCertificateExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("OutboundInspectionCertificate ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-outbound-inspection-certificate", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementOutboundInspectionCertificateAttributes(outboundInspectionCertificateMap *map[string]interface{}, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		outboundInspectionCertificateName := (*outboundInspectionCertificateMap)["name"].(string)
		if !strings.EqualFold(outboundInspectionCertificateName, name) {
			return fmt.Errorf("name is %s, expected %s", name, outboundInspectionCertificateName)
		}

		return nil
	}
}

func testAccManagementOutboundInspectionCertificateConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_outbound_inspection_certificate" "test" {
        name = "%s"
  issued_by       = "www.checkpoint.com"
  base64_password = "bXlfcGFzc3dvcmQ="
  valid_from      = "2021-04-17"
  valid_to        = "2028-04-17"
}
`, name)
}
