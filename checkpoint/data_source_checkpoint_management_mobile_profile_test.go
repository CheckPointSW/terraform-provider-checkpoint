package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointMobileProfile_basic(t *testing.T) {
	objName := "tfTestManagementMds_" + acctest.RandString(6)
	resourceName := "checkpoint_management_mobile_profile.test"
	dataSourceName := "data.checkpoint_management_mobile_profile.data"

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
				Config: testAccDataSourceManagementMobileProfileConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})
}

func testAccDataSourceManagementMobileProfileConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_mobile_profile" "test" {

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
data "checkpoint_management_mobile_profile" "data" {
  uid = "${checkpoint_management_mobile_profile.test.id}"
}
`, name)
}
