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

func TestAccCheckpointManagementExternalTrustedCa_basic(t *testing.T) {

	var externalTrustedCaMap map[string]interface{}
	resourceName := "checkpoint_management_external_trusted_ca.test"
	objName := "tfTestManagementExternalTrustedCa_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementExternalTrustedCaDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementExternalTrustedCaConfig(objName, "MIICujCCAaKgAwIBAgIIP1+IHWHbl0EwDQYJKoZIhvcNAQELBQAwFDESMBAGA1UEAxMJd3d3LnouY29tMB4XDTIzMTEyOTEyMzAwMFoXDTI0MTEyMDE2MDAwMFowFDESMBAGA1UEAxMJd3d3LnouY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAoBreRGuq8u43GBog+ZaAnaR8ZF8cT2ppvtd3JoFmzTOQivLIt9sNtFYqEgHCtnNkKn9TRrxN14YscHgKIxfDSVlC9Rh0rrBvWgFqcm715Whr99Ogx6JbYFkusFWJarSejIFx4n6MM48MJxLdtCP6Hy1G2cj1BCiCHj4i3VIVaDE/aMkSqJbYEvf+vFqUWxY8/uEuKI/HGhI7mhUPW4NSGL0Oafz5eEFVsxqV5NA19/JJZ9NajSkyANnaNL5raxGV0oeqaE3JB3lSEZfWbH6mQsToUxxwIQfsZiIBozajDdTgP3Kn4SMY0b+I/WAWgfigMSDTAIR8J1sdzGXy2w2kqQIDAQABoxAwDjAMBgNVHRMEBTADAQH/MA0GCSqGSIb3DQEBCwUAA4IBAQBUgrHztHwC1E0mU5c4reMrHg+z+YRHrgJNHVIYQbL5I2TJHk9S3UZsynoMa1CO86rReOtR5xoGv4PCkyyOW+PNlWUtXF3tNgqWj/21+XzG4RBHPw89TaTxRCdo+MHX58fi07SIzKjmxfdkEi+7+HQEQluDZGViolrGBAw2rXq/SZ3q/11mNqlb5ZyqyOa2u1sBF1ApvG5a/FBRTaO8gaiNelRf0PGYkuV+1HhF2XyP8Qk565d+uxUH5M7eHF2PNyVk/r/36T+x+UMql9y9iizA0ekuAjXLok1xYl3Vw4S5zXCXYtNZLOVrs+plJb7IrlElyTOAbDFuPugh0medz7uZ"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementExternalTrustedCaExists(resourceName, &externalTrustedCaMap),
					//	testAccCheckCheckpointManagementExternalTrustedCaAttributes(&externalTrustedCaMap, objName, "miicujccaakgawibagiip1+ihwhbl0ewdqyjkozihvcnaqelbqawfdesmbaga1ueaxmjd3d3lnouy29tmb4xdtizmteyoteymzawmfoxdti0mteymde2mdawmfowfdesmbaga1ueaxmjd3d3lnouy29tmiibijanbgkqhkig9w0baqefaaocaq8amiibcgkcaqeaobrerguq8u43gbog+zaanar8zf8ct2ppvtd3jofmztoqivlit9sntfyqeghctnnkkn9trrxn14yschgkixfdsvlc9rh0rrbvwgfqcm715whr99ogx6jbyfkusfwjarsejifx4n6mm48mjxldtcp6hy1g2cj1bcichj4i3vivade/amksqjbyevf+vfquwxy8/ueuki/hghi7mhupw4nsgl0oafz5eefvsxqv5na19/jjz9najskyannanl5raxgv0oeqae3jb3lsezfwbh6mqstouxxwiqfsziibozajddtgp3kn4smy0b+i/wawgfigmsdtair8j1sdzgxy2w2kqqidaqaboxawdjambgnvhrmebtadaqh/ma0gcsqgsib3dqebcwuaa4ibaqbugrhzthwc1e0mu5c4remrhg+z+yrhrgjnhviyqbl5i2tjhk9s3uzsynoma1co86rreotr5xogv4pckyyow+pnlwutxf3tngqwj/21+xzg4rbhpw89tatxrcdo+mhx58fi07sizkjmxfdkei+7+hqeqludzgviolrgbaw2rxq/sz3q/11mnqlb5zyqyoa2u1sbf1apvg5a/fbrtao8gainelrf0pgykuv+1hhf2xyp8qk565d+uxuh5m7ehf2pnyvk/r/36t+x+umql9y9iiza0ekuajxlok1xyl3vw4s5zxcxytnzlovrs+pljb7irlelytoabdfupugh0medz7uz"),
				),
			},
		},
	})
}

func testAccCheckpointManagementExternalTrustedCaDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_external_trusted_ca" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-external-trusted-ca", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success {
				return fmt.Errorf("ExternalTrustedCa object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementExternalTrustedCaExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ExternalTrustedCa ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-external-trusted-ca", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementExternalTrustedCaAttributes(externalTrustedCaMap *map[string]interface{}, name string, base64Certificate string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		externalTrustedCaName := (*externalTrustedCaMap)["name"].(string)
		if !strings.EqualFold(externalTrustedCaName, name) {
			return fmt.Errorf("name is %s, expected %s", name, externalTrustedCaName)
		}
		externalTrustedCaBase64Certificate := (*externalTrustedCaMap)["base64-certificate"].(string)
		if !strings.EqualFold(externalTrustedCaBase64Certificate, base64Certificate) {
			return fmt.Errorf("base64Certificate is %s, expected %s", base64Certificate, externalTrustedCaBase64Certificate)
		}
		return nil
	}
}

func testAccManagementExternalTrustedCaConfig(name string, base64Certificate string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_external_trusted_ca" "test" {
        name = "%s"
        base64_certificate = "%s"
}
`, name, base64Certificate)
}
