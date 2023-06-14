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

func TestAccCheckpointManagementServiceGtp_basic(t *testing.T) {

	var serviceGtpMap map[string]interface{}
	resourceName := "checkpoint_management_service_gtp.test"
	objName := "tfTestManagementServiceGtp" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementServiceGtpDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementServiceGtpConfig(objName, "v2", false, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementServiceGtpExists(resourceName, &serviceGtpMap),
					testAccCheckCheckpointManagementServiceGtpAttributes(&serviceGtpMap, objName, "v2", false, true),
				),
			},
		},
	})
}

func testAccCheckpointManagementServiceGtpDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_service_gtp" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-service-gtp", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success {
				return fmt.Errorf("ServiceGtp object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementServiceGtpExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ServiceGtp ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-service-gtp", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementServiceGtpAttributes(serviceGtpMap *map[string]interface{}, name string, version string, reverseService bool, traceManagement bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		serviceGtpName := (*serviceGtpMap)["name"].(string)
		if !strings.EqualFold(serviceGtpName, name) {
			return fmt.Errorf("name is %s, expected %s", name, serviceGtpName)
		}
		serviceGtpVersion := (*serviceGtpMap)["version"].(string)
		if !strings.EqualFold(serviceGtpVersion, version) {
			return fmt.Errorf("version is %s, expected %s", version, serviceGtpVersion)
		}
		serviceGtpReverseService := (*serviceGtpMap)["reverse-service"].(bool)
		if serviceGtpReverseService != reverseService {
			return fmt.Errorf("reverseService is %t, expected %t", reverseService, serviceGtpReverseService)
		}
		serviceGtpTraceManagement := (*serviceGtpMap)["trace-management"].(bool)
		if serviceGtpTraceManagement != traceManagement {
			return fmt.Errorf("traceManagement is %t, expected %t", traceManagement, serviceGtpTraceManagement)
		}
		return nil
	}
}

func testAccManagementServiceGtpConfig(name string, version string, reverseService bool, traceManagement bool) string {
	return fmt.Sprintf(`
resource "checkpoint_management_service_gtp" "test" {
        name = "%s"
        version = "%s"
        reverse_service = %t
        trace_management = %t
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
`, name, version, reverseService, traceManagement)
}
