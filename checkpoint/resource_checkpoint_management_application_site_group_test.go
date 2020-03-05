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

func TestAccCheckpointManagementApplicationSiteGroup_basic(t *testing.T) {

    var applicationSiteGroupMap map[string]interface{}
    resourceName := "checkpoint_management_application_site_group.test"
    objName := "tfTestManagementApplicationSiteGroup_" + acctest.RandString(6)

    context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
        CheckDestroy: testAccCheckpointManagementApplicationSiteGroupDestroy,
        Steps: []resource.TestStep{
            {
                Config: testAccManagementApplicationSiteGroupConfig(objName, "New Application Site 1", "New Application Site Category 1", "Social Networking", "facebook"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckCheckpointManagementApplicationSiteGroupExists(resourceName, &applicationSiteGroupMap),
                    testAccCheckCheckpointManagementApplicationSiteGroupAttributes(&applicationSiteGroupMap, objName, "New Application Site 1", "New Application Site Category 1", "Social Networking", "facebook"),
                ),
            },
        },
    })
}

func testAccCheckpointManagementApplicationSiteGroupDestroy(s *terraform.State) error {

    client := testAccProvider.Meta().(*checkpoint.ApiClient)
    for _, rs := range s.RootModule().Resources {
        if rs.Type != "checkpoint_management_application_site_group" {
            continue
        }
        if rs.Primary.ID != "" {
            res, _ := client.ApiCall("show-application-site-group", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
            if res.Success {
                return fmt.Errorf("ApplicationSiteGroup object (%s) still exists", rs.Primary.ID)
            }
        }
        return nil
    }
    return nil
}

func testAccCheckCheckpointManagementApplicationSiteGroupExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        rs, ok := s.RootModule().Resources[resourceTfName]
        if !ok {
            return fmt.Errorf("Resource not found: %s", resourceTfName)
        }

        if rs.Primary.ID == "" {
            return fmt.Errorf("ApplicationSiteGroup ID is not set")
        }

        client := testAccProvider.Meta().(*checkpoint.ApiClient)

        response, err := client.ApiCall("show-application-site-group", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
        if !response.Success {
            return err
        }

        *res = response.GetData()

        return nil
    }
}

func testAccCheckCheckpointManagementApplicationSiteGroupAttributes(applicationSiteGroupMap *map[string]interface{}, name string, members1 string, members2 string, members3 string, members4 string) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        applicationSiteGroupName := (*applicationSiteGroupMap)["name"].(string)
        if !strings.EqualFold(applicationSiteGroupName, name) {
            return fmt.Errorf("name is %s, expected %s", name, applicationSiteGroupName)
        }
        membersJson := (*applicationSiteGroupMap)["members"].([]interface{})
        var membersIds = make([]string, 0)
        if len(membersJson) > 0 {
            for _, members := range membersJson {
                membersTry1, ok := members.(map[string]interface{})
                if ok {
                    membersIds = append([]string{membersTry1["name"].(string)}, membersIds...)
                } else {
                    membersTry2:= members.(string)
                    membersIds = append([]string{membersTry2}, membersIds...)
                }
            }
        }

        ApplicationSiteGroupmembers1 := membersIds[0]
        if ApplicationSiteGroupmembers1 != members1 {
            return fmt.Errorf("members1 is %s, expected %s", members1, ApplicationSiteGroupmembers1)
        }
        ApplicationSiteGroupmembers2 := membersIds[1]
        if ApplicationSiteGroupmembers2 != members2 {
            return fmt.Errorf("members2 is %s, expected %s", members2, ApplicationSiteGroupmembers2)
        }
        ApplicationSiteGroupmembers3 := membersIds[2]
        if ApplicationSiteGroupmembers3 != members3 {
            return fmt.Errorf("members3 is %s, expected %s", members3, ApplicationSiteGroupmembers3)
        }
        ApplicationSiteGroupmembers4 := membersIds[3]
        if ApplicationSiteGroupmembers4 != members4 {
            return fmt.Errorf("members4 is %s, expected %s", members4, ApplicationSiteGroupmembers4)
        }
        return nil
    }
}

func testAccManagementApplicationSiteGroupConfig(name string, members1 string, members2 string, members3 string, members4 string) string {
    return fmt.Sprintf(`
resource "checkpoint_management_application_site_group" "test" {
        name = "%s"
        members = ["%s","%s","%s","%s"]
}
`, name, members1, members2, members3, members4)
}

