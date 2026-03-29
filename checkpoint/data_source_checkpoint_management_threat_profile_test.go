package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementThreatProfile_basic(t *testing.T) {

	objName := "ThreatProfile" + acctest.RandString(2)
	resourceName := "checkpoint_management_threat_profile.threat_profile"
	dataSourceName := "data.checkpoint_management_threat_profile.test_threat_profile"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementThreatProfileConfig(objName, "high", "Critical"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "active_protections_performance_impact", resourceName, "active_protections_performance_impact"),
					resource.TestCheckResourceAttrPair(dataSourceName, "active_protections_severity", resourceName, "active_protections_severity"),
				),
			},
		},
	})
}

func testAccDataSourceManagementThreatProfileConfig(name string, performanceImpact string, protectionsSeverity string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_threat_profile" "threat_profile" {
	name = "%s"
	active_protections_performance_impact = "%s"
	active_protections_severity	 = "%s"
}

data "checkpoint_management_threat_profile" "test_threat_profile" {
    name = "${checkpoint_management_threat_profile.threat_profile.name}"
}
`, name, performanceImpact, protectionsSeverity)
}
