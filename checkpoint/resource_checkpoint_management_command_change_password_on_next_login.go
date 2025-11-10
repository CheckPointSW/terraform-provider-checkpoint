package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementChangePasswordOnNextLogin() *schema.Resource {
	return &schema.Resource{
		Create: createManagementChangePasswordOnNextLogin,
		Read:   readManagementChangePasswordOnNextLogin,
		Delete: deleteManagementChangePasswordOnNextLogin,
		Schema: map[string]*schema.Schema{
			"message": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Operation status.",
			},
		},
	}
}

func createManagementChangePasswordOnNextLogin(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	ChangePasswordOnNextLoginRes, err := client.ApiCall("change-password-on-next-login", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !ChangePasswordOnNextLoginRes.Success {
		return fmt.Errorf(ChangePasswordOnNextLoginRes.ErrorMsg)
	}

	changePasswordOnNextLogin := ChangePasswordOnNextLoginRes.GetData()

	if v := changePasswordOnNextLogin["message"]; v != nil {
		_ = d.Set("message", v)
	}

	d.SetId("change-password-on-next-login-" + acctest.RandString(10))
	return nil
}

func readManagementChangePasswordOnNextLogin(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementChangePasswordOnNextLogin(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
