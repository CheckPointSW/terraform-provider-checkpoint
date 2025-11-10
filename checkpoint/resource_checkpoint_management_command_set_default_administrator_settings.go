package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementSetDefaultAdministratorSettings() *schema.Resource {
	return &schema.Resource{
		Create: createManagementSetDefaultAdministratorSettings,
		Read:   readManagementSetDefaultAdministratorSettings,
		Delete: deleteManagementSetDefaultAdministratorSettings,
		Schema: map[string]*schema.Schema{
			"authentication_method": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Authentication method for new administrator.",
			},
			"expiration_type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Expiration type for new administrator.",
			},
			"expiration_date": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Expiration date for new administrator in YYYY-MM-DD format. <font color=\"red\">Required only when</font> 'expiration-type' is set to 'expiration date'.",
			},
			"expiration_period": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "Expiration period for new administrator. <font color=\"red\">Required only when</font> 'expiration-type' is set to 'expiration period'.",
			},
			"expiration_period_time_units": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Expiration period time units for new administrator. <font color=\"red\">Required only when</font> 'expiration-type' is set to 'expiration period'.",
			},
			"indicate_expiration_in_admin_view": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Indicates whether to notify administrator about expiration.",
			},
			"notify_expiration_to_admin": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Indicates whether to show 'about to expire' indication in administrator view.",
			},
			"days_to_indicate_expiration_in_admin_view": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "Number of days in advanced to show 'about to expire' indication in administrator view.",
			},
			"days_to_notify_expiration_to_admin": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "Number of days in advanced to notify administrator about expiration.",
			},
		},
	}
}

func createManagementSetDefaultAdministratorSettings(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("authentication_method"); ok {
		payload["authentication-method"] = v.(string)
	}

	if v, ok := d.GetOk("expiration_type"); ok {
		payload["expiration-type"] = v.(string)
	}

	if v, ok := d.GetOk("expiration_date"); ok {
		payload["expiration-date"] = v.(string)
	}

	if v, ok := d.GetOk("expiration_period"); ok {
		payload["expiration-period"] = v.(int)
	}

	if v, ok := d.GetOk("expiration_period_time_units"); ok {
		payload["expiration-period-time-units"] = v.(string)
	}

	if v, ok := d.GetOkExists("indicate_expiration_in_admin_view"); ok {
		payload["indicate-expiration-in-admin-view"] = v.(bool)
	}

	if v, ok := d.GetOkExists("notify_expiration_to_admin"); ok {
		payload["notify-expiration-to-admin"] = v.(bool)
	}

	if v, ok := d.GetOk("days_to_indicate_expiration_in_admin_view"); ok {
		payload["days-to-indicate-expiration-in-admin-view"] = v.(int)
	}

	if v, ok := d.GetOk("days_to_notify_expiration_to_admin"); ok {
		payload["days-to-notify-expiration-to-admin"] = v.(int)
	}

	SetDefaultAdministratorSettingsRes, err := client.ApiCall("set-default-administrator-settings", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !SetDefaultAdministratorSettingsRes.Success {
		return fmt.Errorf(SetDefaultAdministratorSettingsRes.ErrorMsg)
	}

	d.SetId("set-default-administrator-settings-" + acctest.RandString(10))
	return nil
}

func readManagementSetDefaultAdministratorSettings(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementSetDefaultAdministratorSettings(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
