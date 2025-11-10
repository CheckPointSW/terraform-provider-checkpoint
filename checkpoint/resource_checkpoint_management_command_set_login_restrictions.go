package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementSetLoginRestrictions() *schema.Resource {
	return &schema.Resource{
		Create: createManagementSetLoginRestrictions,
		Read:   readManagementSetLoginRestrictions,
		Delete: deleteManagementSetLoginRestrictions,
		Schema: map[string]*schema.Schema{
			"lockout_admin_account": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Indicates whether to lockout administrator's account after specified number of failed authentication attempts.",
			},
			"failed_authentication_attempts": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "Number of failed authentication attempts before lockout administrator account. <font color=\"red\">Required only when</font> lockout-admin-account is set to true.",
			},
			"unlock_admin_account": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Indicates whether to unlock administrator account after specified number of minutes. <font color=\"red\">Required only when</font> lockout-admin-account is set to true.",
			},
			"lockout_duration": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "Number of minutes of administrator account lockout. <font color=\"red\">Required only when</font> lockout-admin-account is set to true.",
			},
			"display_access_denied_message": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Indicates whether to display informative message upon denying access. <font color=\"red\">Required only when</font> lockout-admin-account is set to true.",
			},
		},
	}
}

func createManagementSetLoginRestrictions(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOkExists("lockout_admin_account"); ok {
		payload["lockout-admin-account"] = v.(bool)
	}

	if v, ok := d.GetOk("failed_authentication_attempts"); ok {
		payload["failed-authentication-attempts"] = v.(int)
	}

	if v, ok := d.GetOkExists("unlock_admin_account"); ok {
		payload["unlock-admin-account"] = v.(bool)
	}

	if v, ok := d.GetOk("lockout_duration"); ok {
		payload["lockout-duration"] = v.(int)
	}

	if v, ok := d.GetOkExists("display_access_denied_message"); ok {
		payload["display-access-denied-message"] = v.(bool)
	}

	SetLoginRestrictionsRes, err := client.ApiCall("set-login-restrictions", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !SetLoginRestrictionsRes.Success {
		return fmt.Errorf(SetLoginRestrictionsRes.ErrorMsg)
	}

	d.SetId("set-login-restrictions-" + acctest.RandString(10))
	return nil
}

func readManagementSetLoginRestrictions(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementSetLoginRestrictions(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
