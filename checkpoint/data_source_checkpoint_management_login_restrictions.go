package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceManagementLoginRestrictions() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementLoginRestrictionsRead,
		Schema: map[string]*schema.Schema{
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object unique identifier.",
			},
			"lockout_admin_account": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether to lockout administrator's account after specified number of failed authentication attempts.",
			},
			"failed_authentication_attempts": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Number of failed authentication attempts before lockout administrator account. <font color=\"red\">Required only when</font> lockout-admin-account is set to true.",
			},
			"unlock_admin_account": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether to unlock administrator account after specified number of minutes. <font color=\"red\">Required only when</font> lockout-admin-account is set to true.",
			},
			"lockout_duration": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Number of minutes of administrator account lockout. <font color=\"red\">Required only when</font> lockout-admin-account is set to true.",
			},
			"display_access_denied_message": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether to display informative message upon denying access. <font color=\"red\">Required only when</font> lockout-admin-account is set to true.",
			},
		},
	}
}

func dataSourceManagementLoginRestrictionsRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}

	loginRestrictionsRes, err := client.ApiCallSimple("show-login-restrictions", payload)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !loginRestrictionsRes.Success {
		return fmt.Errorf(loginRestrictionsRes.ErrorMsg)
	}
	loginRestrictionsData := loginRestrictionsRes.GetData()

	if v := loginRestrictionsData["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := loginRestrictionsData["lockout-admin-account"]; v != nil {
		_ = d.Set("lockout_admin_account", v)
	}

	if v := loginRestrictionsData["failed-authentication-attempts"]; v != nil {
		_ = d.Set("failed_authentication_attempts", v)
	}

	if v := loginRestrictionsData["unlock-admin-account"]; v != nil {
		_ = d.Set("unlock_admin_account", v)
	}

	if v := loginRestrictionsData["lockout-duration"]; v != nil {
		_ = d.Set("lockout_duration", v)
	}

	if v := loginRestrictionsData["display-access-denied-message"]; v != nil {
		_ = d.Set("display_access_denied_message", v)
	}

	return nil
}
