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

func TestAccCheckpointManagementGsnHandoverGroup_basic(t *testing.T) {

	var gsnHandoverGroupMap map[string]interface{}
	resourceName := "checkpoint_management_gsn_handover_group.test"
	objName := "tfTestManagementGsnHandoverGroup_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementGsnHandoverGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementGsnHandoverGroupConfig(objName, true, 2048, "All_Internet"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementGsnHandoverGroupExists(resourceName, &gsnHandoverGroupMap),
					testAccCheckCheckpointManagementGsnHandoverGroupAttributes(&gsnHandoverGroupMap, objName, true, 2048, "All_Internet"),
				),
			},
		},
	})
}

func testAccCheckpointManagementGsnHandoverGroupDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_gsn_handover_group" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-gsn-handover-group", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("GsnHandoverGroup object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementGsnHandoverGroupExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("GsnHandoverGroup ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-gsn-handover-group", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementGsnHandoverGroupAttributes(gsnHandoverGroupMap *map[string]interface{}, name string, enforceGtp bool, gtpRate int, members1 string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		gsnHandoverGroupName := (*gsnHandoverGroupMap)["name"].(string)
		if !strings.EqualFold(gsnHandoverGroupName, name) {
			return fmt.Errorf("name is %s, expected %s", name, gsnHandoverGroupName)
		}
		gsnHandoverGroupEnforceGtp := (*gsnHandoverGroupMap)["enforce-gtp"].(bool)
		if gsnHandoverGroupEnforceGtp != enforceGtp {
			return fmt.Errorf("enforceGtp is %t, expected %t", enforceGtp, gsnHandoverGroupEnforceGtp)
		}
		gsnHandoverGroupGtpRate := int((*gsnHandoverGroupMap)["gtp-rate"].(float64))
		if gsnHandoverGroupGtpRate != gtpRate {
			return fmt.Errorf("gtpRate is %d, expected %d", gtpRate, gsnHandoverGroupGtpRate)
		}
		membersJson := (*gsnHandoverGroupMap)["members"].([]interface{})
		var membersIds = make([]string, 0)
		if len(membersJson) > 0 {
			for _, members := range membersJson {
				membersTry1, ok := members.(map[string]interface{})
				if ok {
					membersIds = append([]string{membersTry1["name"].(string)}, membersIds...)
				} else {
					membersTry2 := members.(string)
					membersIds = append([]string{membersTry2}, membersIds...)
				}
			}
		}

		GsnHandoverGroupmembers1 := membersIds[0]
		if GsnHandoverGroupmembers1 != members1 {
			return fmt.Errorf("members1 is %s, expected %s", members1, GsnHandoverGroupmembers1)
		}
		return nil
	}
}

func testAccManagementGsnHandoverGroupConfig(name string, enforceGtp bool, gtpRate int, members1 string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_gsn_handover_group" "test" {
        name = "%s"
        enforce_gtp = %t
        gtp_rate = %d
        members = ["%s"]
}
`, name, enforceGtp, gtpRate, members1)
}
