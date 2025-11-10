package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceManagementDefaultAdministratorSettings() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementDefaultAdministratorSettingsRead,
		Schema: map[string]*schema.Schema{
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object unique identifier.",
			},
			"authentication_method": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Authentication method for new administrator.",
			},
			"expiration_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Expiration type for new administrator.",
			},
			"expiration_date": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Expiration date for new administrator in YYYY-MM-DD format. <font color=\"red\">Required only when</font> 'expiration-type' is set to 'expiration date'.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"iso_8601": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Date and time represented in international ISO 8601 format.",
						},
						"posix": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.",
						},
					},
				},
			},
			"expiration_period": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Expiration period for new administrator. <font color=\"red\">Required only when</font> 'expiration-type' is set to 'expiration period'.",
			},
			"expiration_period_time_units": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Expiration period time units for new administrator. <font color=\"red\">Required only when</font> 'expiration-type' is set to 'expiration period'.",
			},
			"indicate_expiration_in_admin_view": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether to notify administrator about expiration.",
			},
			"notify_expiration_to_admin": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether to show 'about to expire' indication in administrator view.",
			},
			"days_to_indicate_expiration_in_admin_view": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Number of days in advanced to show 'about to expire' indication in administrator view.",
			},
			"days_to_notify_expiration_to_admin": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Number of days in advanced to notify administrator about expiration.",
			},
		},
	}
}

func dataSourceManagementDefaultAdministratorSettingsRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}

	defaultAdministratorSettingsRes, err := client.ApiCallSimple("show-default-administrator-settings", payload)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !defaultAdministratorSettingsRes.Success {
		return fmt.Errorf(defaultAdministratorSettingsRes.ErrorMsg)
	}
	defaultAdministratorSettingsData := defaultAdministratorSettingsRes.GetData()

	if v := defaultAdministratorSettingsData["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := defaultAdministratorSettingsData["authentication-method"]; v != nil {
		_ = d.Set("authentication_method", v)
	}

	if v := defaultAdministratorSettingsData["expiration-type"]; v != nil {
		_ = d.Set("expiration_type", v)
	}

	if defaultAdministratorSettingsData["expiration-date"] != nil {
		expirationDateMap := defaultAdministratorSettingsData["expiration-date"].(map[string]interface{})
		expirationDateList := make([]interface{}, 0)

		schedulingMapToReturn := make(map[string]interface{})

		if v, _ := expirationDateMap["iso-8601"]; v != nil {
			schedulingMapToReturn["iso_8601"] = v
		}
		if v, _ := expirationDateMap["posix"]; v != nil {
			schedulingMapToReturn["posix"] = v
		}

		expirationDateList = append(expirationDateList, schedulingMapToReturn)
		_ = d.Set("expiration_date", expirationDateList)
	} else {
		_ = d.Set("expiration_date", nil)
	}

	if v := defaultAdministratorSettingsData["expiration-period"]; v != nil {
		_ = d.Set("expiration_period", v)
	}

	if v := defaultAdministratorSettingsData["expiration-period-time-units"]; v != nil {
		_ = d.Set("expiration_period_time_units", v)
	}

	if v := defaultAdministratorSettingsData["indicate-expiration-in-admin-view"]; v != nil {
		_ = d.Set("indicate_expiration_in_admin_view", v)
	}

	if v := defaultAdministratorSettingsData["notify-expiration-to-admin"]; v != nil {
		_ = d.Set("notify_expiration_to_admin", v)
	}

	if v := defaultAdministratorSettingsData["days-to-indicate-expiration-in-admin-view"]; v != nil {
		_ = d.Set("days_to_indicate_expiration_in_admin_view", v)
	}

	if v := defaultAdministratorSettingsData["days-to-notify-expiration-to-admin"]; v != nil {
		_ = d.Set("days_to_notify_expiration_to_admin", v)
	}

	return nil
}
