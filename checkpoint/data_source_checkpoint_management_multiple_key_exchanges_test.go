package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementMultipleKeyExchanges_basic(t *testing.T) {

	objName := "tfTestManagementDataMulticastAddressRange_" + acctest.RandString(6)
	resourceName := "checkpoint_management_multiple_key_exchanges.test"
	dataSourceName := "data.checkpoint_management_multiple_key_exchanges.data"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementMultipleKeyExchangesConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})

}

func testAccDataSourceManagementMultipleKeyExchangesConfig(name string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_multiple_key_exchanges" "test" {
        name = "%s"
        key_exchange_methods = ["group-2"] 
        additional_key_exchange_1_methods =  ["kyber-1024"]
}

data "checkpoint_management_multiple_key_exchanges" "data" {
  name = "${checkpoint_management_multiple_key_exchanges.test.name}"
}
`, name)
}
