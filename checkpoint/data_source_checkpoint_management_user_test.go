package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceCheckpointManagementUser_basic(t *testing.T) {

	objName := "User" + acctest.RandString(2)
	resourceName := "checkpoint_management_user.user"
	dataSourceName := "data.checkpoint_management_user.test_user"

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManagementUserConfig(objName, "myuser@email.com", "2030-05-30", "0501112233", "securid", true, "08:00", "17:00"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
				),
			},
		},
	})
}

func testAccDataSourceManagementUserConfig(name string, email string, expirationDate string, phoneNumber string, authenticationMethod string, connectDaily bool, fromHour string, toHour string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_user" "user" {
        name = "%s"
        email = "%s"
        expiration_date = "%s"
        phone_number = "%s"
        authentication_method = "%s"
        connect_daily = %t
        from_hour = "%s"
        to_hour = "%s"
}

data "checkpoint_management_user" "test_user" {
    name = "${checkpoint_management_user.user.name}"
}
`, name, email, expirationDate, phoneNumber, authenticationMethod, connectDaily, fromHour, toHour)
}
