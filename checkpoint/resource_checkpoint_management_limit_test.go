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

func TestAccCheckpointManagementLimit_basic(t *testing.T) {

	var limitMap map[string]interface{}
	resourceName := "checkpoint_management_limit.test"
	objName := "tfTestManagementLimit_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementLimitDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementLimitConfig(objName, true, "gbps", 4),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementLimitExists(resourceName, &limitMap),
					testAccCheckCheckpointManagementLimitAttributes(&limitMap, objName, true, "gbps", 4),
				),
			},
		},
	})
}

func testAccCheckpointManagementLimitDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_limit" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-limit", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
			if res.Success {
				return fmt.Errorf("Limit object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementLimitExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Limit ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-limit", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementLimitAttributes(limitMap *map[string]interface{}, name string, enableDownload bool, downloadUnit string, downloadRate int) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		limitName := (*limitMap)["name"].(string)
		if !strings.EqualFold(limitName, name) {
			return fmt.Errorf("name is %s, expected %s", name, limitName)
		}
		limitEnableDownload := (*limitMap)["enable-download"].(bool)
		if limitEnableDownload != enableDownload {
			return fmt.Errorf("enableDownload is %t, expected %t", enableDownload, limitEnableDownload)
		}
		limitDownloadUnit := (*limitMap)["download-unit"].(string)
		if !strings.EqualFold(limitDownloadUnit, downloadUnit) {
			return fmt.Errorf("downloadUnit is %s, expected %s", downloadUnit, limitDownloadUnit)
		}
		limitDownloadRate := int((*limitMap)["download-rate"].(float64))
		if limitDownloadRate != downloadRate {
			return fmt.Errorf("downloadRate is %d, expected %d", downloadRate, limitDownloadRate)
		}
		return nil
	}
}

func testAccManagementLimitConfig(name string, enableDownload bool, downloadUnit string, downloadRate int) string {
	return fmt.Sprintf(`
resource "checkpoint_management_limit" "test" {
        name = "%s"
        enable_download = %t
        download_unit = "%s"
        download_rate = %d
}
`, name, enableDownload, downloadUnit, downloadRate)
}
