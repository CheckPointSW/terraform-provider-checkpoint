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

func TestAccCheckpointManagementUser_basic(t *testing.T) {

	var userMap map[string]interface{}
	resourceName := "checkpoint_management_user.test"
	objName := "tfTestManagementUser_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementUserConfig(objName, "myuser@email.com", "2030-05-30", "0501112233", "securid", true, "08:00", "17:00"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementUserExists(resourceName, &userMap),
					testAccCheckCheckpointManagementUserAttributes(&userMap, objName, "myuser@email.com", "2030-05-30", "0501112233", "securid", true, "08:00", "17:00"),
				),
			},
		},
	})
}

func testAccCheckpointManagementUserDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_user" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-user", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("User object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementUserExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("User ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-user", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementUserAttributes(userMap *map[string]interface{}, name string, email string, expirationDate string, phoneNumber string, authenticationMethod string, connectDaily bool, fromHour string, toHour string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		userName := (*userMap)["name"].(string)
		if !strings.EqualFold(userName, name) {
			return fmt.Errorf("name is %s, expected %s", name, userName)
		}
		userEmail := (*userMap)["email"].(string)
		if !strings.EqualFold(userEmail, email) {
			return fmt.Errorf("email is %s, expected %s", email, userEmail)
		}
		userDate := (*userMap)["expiration-date"].(map[string]interface{})["iso-8601"].(string)
		userExpirationDate := strings.Split(userDate, "T")[0]
		if !strings.EqualFold(userExpirationDate, expirationDate) {
			return fmt.Errorf("expirationDate is %s, expected %s", expirationDate, userExpirationDate)
		}
		userPhoneNumber := (*userMap)["phone-number"].(string)
		if !strings.EqualFold(userPhoneNumber, phoneNumber) {
			return fmt.Errorf("phoneNumber is %s, expected %s", phoneNumber, userPhoneNumber)
		}
		userAuthenticationMethod := (*userMap)["authentication-method"].(string)
		if !strings.EqualFold(userAuthenticationMethod, authenticationMethod) {
			return fmt.Errorf("authenticationMethod is %s, expected %s", authenticationMethod, userAuthenticationMethod)
		}
		userConnectDaily := (*userMap)["connect-daily"].(bool)
		if userConnectDaily != connectDaily {
			return fmt.Errorf("connectDaily is %t, expected %t", connectDaily, userConnectDaily)
		}
		userFromHour := (*userMap)["from-hour"].(string)
		if !strings.EqualFold(userFromHour, fromHour) {
			return fmt.Errorf("fromHour is %s, expected %s", fromHour, userFromHour)
		}
		userToHour := (*userMap)["to-hour"].(string)
		if !strings.EqualFold(userToHour, toHour) {
			return fmt.Errorf("toHour is %s, expected %s", toHour, userToHour)
		}
		return nil
	}
}

func testAccManagementUserConfig(name string, email string, expirationDate string, phoneNumber string, authenticationMethod string, connectDaily bool, fromHour string, toHour string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_user" "test" {
        name = "%s"
        email = "%s"
        expiration_date = "%s"
        phone_number = "%s"
        authentication_method = "%s"
        connect_daily = %t
        from_hour = "%s"
        to_hour = "%s"
}
`, name, email, expirationDate, phoneNumber, authenticationMethod, connectDaily, fromHour, toHour)
}
