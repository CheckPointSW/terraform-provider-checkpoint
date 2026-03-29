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

func TestAccCheckpointManagementMobileProfile_basic(t *testing.T) {

	//	var mobileProfileMap map[string]interface{}
	//	resourceName := "checkpoint_management_mobile_profile.test"
	objName := "tfTestManagementMobileProfile_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementMobileProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementMobileProfileConfig(objName),
				Check:  resource.ComposeTestCheckFunc(
				//	testAccCheckCheckpointManagementMobileProfileExists(resourceName, &mobileProfileMap),
				//	testAccCheckCheckpointManagementMobileProfileAttributes(&mobileProfileMap, objName),
				),
			},
		},
	})
}

func testAccCheckpointManagementMobileProfileDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_mobile_profile" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-mobile-profile", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success {
				return fmt.Errorf("MobileProfile object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementMobileProfileExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("MobileProfile ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-mobile-profile", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementMobileProfileAttributes(mobileProfileMap *map[string]interface{}, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		mobileProfileName := (*mobileProfileMap)["name"].(string)
		if !strings.EqualFold(mobileProfileName, name) {
			return fmt.Errorf("name is %s, expected %s", name, mobileProfileName)
		}
		return nil
	}
}

func testAccManagementMobileProfileConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_mobile_profile" "obj1" {

  name = "%s"
  applications  {
    enable_print_mails = true
    calendar_from_the_last_unit = "weeks"
  }
  harmony_mobile {
    enable_harmony_mobile_sdk = true
    os_integrity_compromised =  "ignore"
  }
 client_customization {

     allow_calendar        = true
     allow_contacts        = true
     allow_mail            = true
      allow_notes_sync      = true
    allow_saved_file_apps = true
     allow_secure_chat     = true
     allow_tasks           = true
     app_theme_color_dark  = "fc037b"
     app_theme_color_light = "fc037b"

  }
security {
      session_timeout = 2
      session_timeout_unit = "weeks"
      activate_passcode_lock = false
      allow_store_credentials = true
      hide_ssl_connect_anyway_button = true
      block_jailbroken = "block"

  }
  data_leak_prevention {
    accept_protected_file_extensions   = [
     "any file",
  ]
   accept_unprotected_file_extensions = [
         "any file",
]
   allow_copy_paste                   = true
 allow_import_from_gallery          = true
 allow_taking_photos_and_videos     = true
 block_forward_attachments          = false
 block_screenshot                   = false
 offer_capsule_as_viewer            = true  
open_extension_with_external_app   = [
 "any file",
]
 share_protected_extension          = [
 "any file",
]
 share_unprotected_extension        = []
 }
}
`, name)
}
