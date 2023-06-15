package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementServiceGtp_basic(t *testing.T) {

	objName := "tfTestManagementDataServiceIcmp6_" + acctest.RandString(6)
	resourceName := "checkpoint_management_service_gtp.service_gtp"
	dataSourceName := "data.checkpoint_management_service_gtp.data_service_gtp"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementServiceGtpConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})
}

func testAccDataSourceManagementServiceGtpConfig(name string) string {
	return fmt.Sprintf(`
   resource "checkpoint_management_service_gtp" "service_gtp" {
        name = "%s"
        version = "v2"
        reverse_service = true
        trace_management = true
 imsi_prefix = {
    enable = true
    prefix = "123"
  }
  interface_profile = {
    profile = "Custom"
    custom_message_types = "32-35"
  }
  selection_mode {
    enable = true
    mode = 1
  }
   ms_isdn= {
    enable =  true
    ms_isdn = "312"
  }
  access_point_name ={
    enable = true
    apn = "AccP2"
  }
  apply_access_policy_on_user_traffic ={
    enable = true
    add_imsi_field_to_log = true
  }
  radio_access_technology {
    other_types_range {
      enable = true
      types = "11-50"
    }
  }
 ldap_group = {
    enable = true
    group = "ldap_group_1"
    according_to = "MS-ISDN"
  }
}
data "checkpoint_management_service_gtp" "data_service_gtp" {
  name = "${checkpoint_management_service_gtp.service_gtp.name}"
}
`, name)

}
