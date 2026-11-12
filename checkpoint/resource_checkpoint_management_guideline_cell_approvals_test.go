package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"os"
	"strings"
	"testing"
)

func TestAccCheckpointManagementGuidelineCellApprovals_basic(t *testing.T) {

	var guidelineCellApprovalsMap map[string]interface{}
	resourceName := "checkpoint_management_guideline_cell_approvals.test"
	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementGuidelineCellApprovalsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementGuidelineCellApprovalsConfig("Corporate policy", "This is approved for all segments", "any", "any"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementGuidelineCellApprovalsExists(resourceName, &guidelineCellApprovalsMap),
					testAccCheckCheckpointManagementGuidelineCellApprovalsAttributes(&guidelineCellApprovalsMap, "Corporate policy", "This is approved for all segments", "any", "any"),
				),
			},
		},
	})
}

func testAccCheckpointManagementGuidelineCellApprovalsDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_guideline_cell_approvals" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-guideline-cell-approvals", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("GuidelineCellApprovals object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementGuidelineCellApprovalsExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("GuidelineCellApprovals ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-guideline-cell-approvals", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementGuidelineCellApprovalsAttributes(guidelineCellApprovalsMap *map[string]interface{}, guideline string, comment string, from string, to string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		guidelineCellApprovalsGuideline := (*guidelineCellApprovalsMap)["guideline"].(string)
		if !strings.EqualFold(guidelineCellApprovalsGuideline, guideline) {
			return fmt.Errorf("guideline is %s, expected %s", guideline, guidelineCellApprovalsGuideline)
		}
		guidelineCellApprovalsComment := (*guidelineCellApprovalsMap)["comment"].(string)
		if !strings.EqualFold(guidelineCellApprovalsComment, comment) {
			return fmt.Errorf("comment is %s, expected %s", comment, guidelineCellApprovalsComment)
		}
		guidelineCellApprovalsFrom := (*guidelineCellApprovalsMap)["from"].(string)
		if !strings.EqualFold(guidelineCellApprovalsFrom, from) {
			return fmt.Errorf("from is %s, expected %s", from, guidelineCellApprovalsFrom)
		}
		guidelineCellApprovalsTo := (*guidelineCellApprovalsMap)["to"].(string)
		if !strings.EqualFold(guidelineCellApprovalsTo, to) {
			return fmt.Errorf("to is %s, expected %s", to, guidelineCellApprovalsTo)
		}
		return nil
	}
}

func testAccManagementGuidelineCellApprovalsConfig(guideline string, comment string, from string, to string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_access_rule" "rule1" {
  name  = "test-rule-for-approval"
  layer = "Network"
  position {
    top = "top"
  }
}

resource "checkpoint_management_group" "example" {
  name = "approvalTestGroup"
}

resource "checkpoint_management_guideline" "guideline1" {
  name = "%s"
  access_layers {
    access_layer = "Network"
  }
  guideline_groups {
    name = checkpoint_management_group.example.name
    position {
      top = "top"
    }
  }
}

resource "checkpoint_management_guideline_cell_approvals" "test" {
  guideline = checkpoint_management_guideline.guideline1.name
  comment   = "%s"
  from      = "%s"
  to        = "%s"
  approvals {
    rules {
      layer = "Network"
      rule  = checkpoint_management_access_rule.rule1.name
    }
  }
}
`, guideline, comment, from, to)
}
