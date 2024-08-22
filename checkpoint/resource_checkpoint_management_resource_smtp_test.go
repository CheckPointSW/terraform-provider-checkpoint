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

func TestAccCheckpointManagementResourceSmtp_basic(t *testing.T) {

	var resourceSmtpMap map[string]interface{}
	resourceName := "checkpoint_management_resource_smtp.test"
	objName := "tfTestManagementResourceSmtp_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementResourceSmtpDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementResourceSmtpConfig(objName, "deliverserver", "exception log", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementResourceSmtpExists(resourceName, &resourceSmtpMap),
					testAccCheckCheckpointManagementResourceSmtpAttributes(&resourceSmtpMap, objName, "deliverserver", "exception log", true),
				),
			},
		},
	})
}

func testAccCheckpointManagementResourceSmtpDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_resource_smtp" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-resource-smtp", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success {
				return fmt.Errorf("ResourceSmtp object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementResourceSmtpExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ResourceSmtp ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-resource-smtp", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementResourceSmtpAttributes(resourceSmtpMap *map[string]interface{}, name string, mailDeliveryServer string, exceptionTrack string, deliverMessagesUsingDnsMxRecords bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		resourceSmtpName := (*resourceSmtpMap)["name"].(string)
		if !strings.EqualFold(resourceSmtpName, name) {
			return fmt.Errorf("name is %s, expected %s", name, resourceSmtpName)
		}
		resourceSmtpMailDeliveryServer := (*resourceSmtpMap)["mail-delivery-server"].(string)
		if !strings.EqualFold(resourceSmtpMailDeliveryServer, mailDeliveryServer) {
			return fmt.Errorf("mailDeliveryServer is %s, expected %s", mailDeliveryServer, resourceSmtpMailDeliveryServer)
		}
		resourceSmtpExceptionTrack := (*resourceSmtpMap)["exception-track"].(string)
		if !strings.EqualFold(resourceSmtpExceptionTrack, exceptionTrack) {
			return fmt.Errorf("exceptionTrack is %s, expected %s", exceptionTrack, resourceSmtpExceptionTrack)
		}
		resourceSmtpDeliverMessagesUsingDnsMxRecords := (*resourceSmtpMap)["deliver-messages-using-dns-mx-records"].(bool)
		if resourceSmtpDeliverMessagesUsingDnsMxRecords != deliverMessagesUsingDnsMxRecords {
			return fmt.Errorf("deliverMessagesUsingDnsMxRecords is %t, expected %t", deliverMessagesUsingDnsMxRecords, resourceSmtpDeliverMessagesUsingDnsMxRecords)
		}
		return nil
	}
}

func testAccManagementResourceSmtpConfig(name string, mailDeliveryServer string, exceptionTrack string, deliverMessagesUsingDnsMxRecords bool) string {
	return fmt.Sprintf(`
resource "checkpoint_management_resource_smtp" "test" {
        name = "%s"
        mail_delivery_server = "%s"
        exception_track = "%s"
        deliver_messages_using_dns_mx_records = %t
}
`, name, mailDeliveryServer, exceptionTrack, deliverMessagesUsingDnsMxRecords)
}
