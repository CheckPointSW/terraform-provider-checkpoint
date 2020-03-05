package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform/helper/resource"
    "github.com/hashicorp/terraform/terraform"
    "os"
    "strings"
    "testing"
    "github.com/hashicorp/terraform/helper/acctest"
)

func TestAccCheckpointManagementHttpsRule_basic(t *testing.T) {

    var httpsRuleMap map[string]interface{}
    resourceName := "checkpoint_management_https_rule.test"
    objName := "tfTestManagementHttpsRule_" + acctest.RandString(6)

    context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
        CheckDestroy: testAccCheckpointManagementHttpsRuleDestroy,
        Steps: []resource.TestStep{
            {
                Config: testAccManagementHttpsRuleConfig(objName, ),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckCheckpointManagementHttpsRuleExists(resourceName, &httpsRuleMap),
                    testAccCheckCheckpointManagementHttpsRuleAttributes(&httpsRuleMap, objName, ),
                ),
            },
        },
    })
}

func testAccCheckpointManagementHttpsRuleDestroy(s *terraform.State) error {

    client := testAccProvider.Meta().(*checkpoint.ApiClient)
    for _, rs := range s.RootModule().Resources {
        if rs.Type != "checkpoint_management_https_rule" {
            continue
        }
        if rs.Primary.ID != "" {
            res, _ := client.ApiCall("show-https-rule", map[string]interface{}{"uid": rs.Primary.ID, "layer": "New Layer 2"}, client.GetSessionID(), true, false)
            if res.Success {
                return fmt.Errorf("HttpsRule object (%s) still exists", rs.Primary.ID)
            }
        }
        return nil
    }
    return nil
}

func testAccCheckCheckpointManagementHttpsRuleExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        rs, ok := s.RootModule().Resources[resourceTfName]
        if !ok {
            return fmt.Errorf("Resource not found: %s", resourceTfName)
        }

        if rs.Primary.ID == "" {
            return fmt.Errorf("HttpsRule ID is not set")
        }

        client := testAccProvider.Meta().(*checkpoint.ApiClient)

        response, err := client.ApiCall("show-https-rule", map[string]interface{}{"uid": rs.Primary.ID, "layer": "New Layer 2"}, client.GetSessionID(), true, false)
        if !response.Success {
            return err
        }

        *res = response.GetData()

        return nil
    }
}

func testAccCheckCheckpointManagementHttpsRuleAttributes(httpsRuleMap *map[string]interface{}, name string) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        httpsRuleName := (*httpsRuleMap)["name"].(string)
        if !strings.EqualFold(httpsRuleName, name) {
            return fmt.Errorf("name is %s, expected %s", name, httpsRuleName)
        }
        return nil
    }
}

func testAccManagementHttpsRuleConfig(name string) string {
    return fmt.Sprintf(`
resource "checkpoint_management_https_rule" "test" {
        name = "%s"
        position = {top = "top"}
        layer = "New Layer 2"
        blade = ["IPS"]
        destination = ["Internet"]
        enabled = true
        service = ["HTTPS default services"]
        source = ["DMZNet"]
}
`, name)
}