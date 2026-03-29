package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAcDataSourcecCheckpointManagementDataTypeWeightedKeywords_basic(t *testing.T) {

	resourceName := "checkpoint_management_data_type_weighted_keywords.test"
	dataSourceName := "data.checkpoint_management_data_type_weighted_keywords.data_test"

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
				Config: testAccDataSourceManagementDataTypeWweightedKeywordsConfig("objname", "weighted keywords object", "word1", 12),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})
}

func testAccDataSourceManagementDataTypeWweightedKeywordsConfig(name string, desc string, keyword1 string, weight int) string {
	return fmt.Sprintf(`
resource "checkpoint_management_data_type_weighted_keywords" "test" {
        name = "%s"
        description = "%s"
        weighted_keywords {
      			keyword = "%s"
                weight =  "%d"
       }
}
data "checkpoint_management_data_type_weighted_keywords" "data_test" {
    name = "${checkpoint_management_data_type_weighted_keywords.test.name}"
}
`, name, desc, keyword1, weight)
}
