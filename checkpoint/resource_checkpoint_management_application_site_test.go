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

func TestAccCheckpointManagementApplicationSite_basic(t *testing.T) {

    var applicationSiteMap map[string]interface{}
    resourceName := "checkpoint_management_application_site.test"
    objName := "tfTestManagementApplicationSite_" + acctest.RandString(6)

    context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
        CheckDestroy: testAccCheckpointManagementApplicationSiteDestroy,
        Steps: []resource.TestStep{
            {
                Config: testAccManagementApplicationSiteConfig(objName, "social networking", "my application site", "Instant Chat", "New Application Site Category 1", "Supports Streaming", "www.cnet.com", "www.stackoverflow.com", false),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckCheckpointManagementApplicationSiteExists(resourceName, &applicationSiteMap),
                    testAccCheckCheckpointManagementApplicationSiteAttributes(&applicationSiteMap, objName, "social networking", "my application site", "Instant Chat", "New Application Site Category 1", "Supports Streaming", "www.cnet.com", "www.stackoverflow.com", false),
                ),
            },
        },
    })
}

func testAccCheckpointManagementApplicationSiteDestroy(s *terraform.State) error {

    client := testAccProvider.Meta().(*checkpoint.ApiClient)
    for _, rs := range s.RootModule().Resources {
        if rs.Type != "checkpoint_management_application_site" {
            continue
        }
        if rs.Primary.ID != "" {
            res, _ := client.ApiCall("show-application-site", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
            if res.Success {
                return fmt.Errorf("ApplicationSite object (%s) still exists", rs.Primary.ID)
            }
        }
        return nil
    }
    return nil
}

func testAccCheckCheckpointManagementApplicationSiteExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        rs, ok := s.RootModule().Resources[resourceTfName]
        if !ok {
            return fmt.Errorf("Resource not found: %s", resourceTfName)
        }

        if rs.Primary.ID == "" {
            return fmt.Errorf("ApplicationSite ID is not set")
        }

        client := testAccProvider.Meta().(*checkpoint.ApiClient)

        response, err := client.ApiCall("show-application-site", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, false)
        if !response.Success {
            return err
        }

        *res = response.GetData()

        return nil
    }
}

func testAccCheckCheckpointManagementApplicationSiteAttributes(applicationSiteMap *map[string]interface{}, name string, primaryCategory string, description string, additionalCategories1 string, additionalCategories2 string, additionalCategories3 string, urlList1 string, urlList2 string, urlsDefinedAsRegularExpression bool) resource.TestCheckFunc {
    return func(s *terraform.State) error {

        applicationSiteName := (*applicationSiteMap)["name"].(string)
        if !strings.EqualFold(applicationSiteName, name) {
            return fmt.Errorf("name is %s, expected %s", name, applicationSiteName)
        }
        applicationSitePrimaryCategory := (*applicationSiteMap)["primary-category"].(string)
        if !strings.EqualFold(applicationSitePrimaryCategory, primaryCategory) {
            return fmt.Errorf("primaryCategory is %s, expected %s", primaryCategory, applicationSitePrimaryCategory)
        }
        applicationSiteDescription := (*applicationSiteMap)["description"].(string)
        if !strings.EqualFold(applicationSiteDescription, description) {
            return fmt.Errorf("description is %s, expected %s", description, applicationSiteDescription)
        }
        additionalCategoriesJson := (*applicationSiteMap)["additional-categories"].([]interface{})
        var additionalCategoriesIds = make([]string, 0)
        if len(additionalCategoriesJson) > 0 {
            for _, additionalCategories := range additionalCategoriesJson {
                additionalCategoriesTry1, ok := additionalCategories.(map[string]interface{})
                if ok {
                    additionalCategoriesIds = append([]string{additionalCategoriesTry1["name"].(string)}, additionalCategoriesIds...)
                } else {
                    additionalCategoriesTry2:= additionalCategories.(string)
                    additionalCategoriesIds = append([]string{additionalCategoriesTry2}, additionalCategoriesIds...)
                }
            }
        }

        ApplicationSiteadditionalCategories1 := additionalCategoriesIds[0]
        if ApplicationSiteadditionalCategories1 != additionalCategories1 {
            return fmt.Errorf("additionalCategories1 is %s, expected %s", additionalCategories1, ApplicationSiteadditionalCategories1)
        }
        ApplicationSiteadditionalCategories2 := additionalCategoriesIds[1]
        if ApplicationSiteadditionalCategories2 != additionalCategories2 {
            return fmt.Errorf("additionalCategories2 is %s, expected %s", additionalCategories2, ApplicationSiteadditionalCategories2)
        }
        ApplicationSiteadditionalCategories3 := additionalCategoriesIds[2]
        if ApplicationSiteadditionalCategories3 != additionalCategories3 {
            return fmt.Errorf("additionalCategories3 is %s, expected %s", additionalCategories3, ApplicationSiteadditionalCategories3)
        }
        urlListJson := (*applicationSiteMap)["url-list"].([]interface{})
        var urlListIds = make([]string, 0)
        if len(urlListJson) > 0 {
            for _, urlList := range urlListJson {
                urlListTry1, ok := urlList.(map[string]interface{})
                if ok {
                    urlListIds = append([]string{urlListTry1["name"].(string)}, urlListIds...)
                } else {
                    urlListTry2:= urlList.(string)
                    urlListIds = append([]string{urlListTry2}, urlListIds...)
                }
            }
        }

        ApplicationSiteurlList1 := urlListIds[0]
        if ApplicationSiteurlList1 != urlList1 {
            return fmt.Errorf("urlList1 is %s, expected %s", urlList1, ApplicationSiteurlList1)
        }
        ApplicationSiteurlList2 := urlListIds[1]
        if ApplicationSiteurlList2 != urlList2 {
            return fmt.Errorf("urlList2 is %s, expected %s", urlList2, ApplicationSiteurlList2)
        }
        applicationSiteUrlsDefinedAsRegularExpression := (*applicationSiteMap)["urls-defined-as-regular-expression"].(bool)
        if applicationSiteUrlsDefinedAsRegularExpression != urlsDefinedAsRegularExpression {
            return fmt.Errorf("urlsDefinedAsRegularExpression is %t, expected %t", urlsDefinedAsRegularExpression, applicationSiteUrlsDefinedAsRegularExpression)
        }
        return nil
    }
}

func testAccManagementApplicationSiteConfig(name string, primaryCategory string, description string, additionalCategories1 string, additionalCategories2 string, additionalCategories3 string, urlList1 string, urlList2 string, urlsDefinedAsRegularExpression bool) string {
    return fmt.Sprintf(`
resource "checkpoint_management_application_site" "test" {
        name = "%s"
        primary_category = "%s"
        description = "%s"
        additional_categories = ["%s","%s","%s"]
        url_list = ["%s","%s"]
        urls_defined_as_regular_expression = %t
}
`, name, primaryCategory, description, additionalCategories1, additionalCategories2, additionalCategories3, urlList1, urlList2, urlsDefinedAsRegularExpression)
}

