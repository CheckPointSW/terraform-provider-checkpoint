package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementCheckpointHost_basic(t *testing.T) {
	objName := "tfTestManagementDataCheckpointHost_" + acctest.RandString(6)
	resourceName := "checkpoint_management_checkpoint_host.checkpoint_host"
	dataSourceName := "data.checkpoint_management_checkpoint_host.data_checkpoint_host"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementCheckpointHostConfig(objName, "5.5.5.5"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "ipv4_address", resourceName, "ipv4_address"),
				),
			},
		},
	})

}

func testAccDataSourceManagementCheckpointHostConfig(name string, ipv4Address string) string {
	return fmt.Sprintf(
		`resource "checkpoint_management_checkpoint_host" "checkpoint_host" {
	     name = "%s"
	    ipv4_address = "%s"
       }

data "checkpoint_management_checkpoint_host" "data_checkpoint_host" {
	name = "${checkpoint_management_checkpoint_host.checkpoint_host.name}"
}
`, name, ipv4Address)
}
