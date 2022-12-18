package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementCloudServices() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementCloudServicesRead,
		Schema: map[string]*schema.Schema{
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status of the connection to the Infinity Portal.",
			},
			"connected_at": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "The time of the connection between the Management Server and the Infinity Portal.",
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
			"management_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Management Server's public URL.",
			},
			"tenant_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Tenant ID of Infinity Portal.",
			},
			"gateways_onboarding_settings": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Computed:    true,
				Description: "Gateways on-boarding to Infinity Portal settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connection_method": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Indicate whether Gateways will be connected to Infinity Portal automatically or only after policy installation.",
						},
						"participant_gateways": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Which Gateways will be connected to Infinity Portal.",
						},
						"specific_gateways": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Collection of targets identified by Name or UID.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceManagementCloudServicesRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	showCloudServices, err := client.ApiCall("show-cloud-services", make(map[string]interface{}), client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showCloudServices.Success {
		return fmt.Errorf(showCloudServices.ErrorMsg)
	}

	showCloudServicesRes := showCloudServices.GetData()

	log.Println("Show Cloud Services - JSON = ", showCloudServicesRes)

	if v := showCloudServicesRes["status"]; v != nil {
		_ = d.Set("status", v)
	} else {
		_ = d.Set("status", nil)
	}

	if v := showCloudServicesRes["connected-at"]; v != nil {
		if connectedAtShow, ok := showCloudServicesRes["connected-at"].(map[string]interface{}); ok {
			connectedAtState := make(map[string]interface{})
			if v := connectedAtShow["iso-8601"]; v != nil {
				connectedAtState["iso_8601"] = v
			}
			if v := connectedAtShow["posix"]; v != nil {
				connectedAtState["posix"] = v
			}
			_ = d.Set("connected_at", connectedAtState)
		}
	} else {
		_ = d.Set("connected_at", nil)
	}

	if v := showCloudServicesRes["management-url"]; v != nil {
		_ = d.Set("management_url", v)
	} else {
		_ = d.Set("management_url", nil)
	}

	if v := showCloudServicesRes["tenant-id"]; v != nil {
		_ = d.Set("tenant_id", v)
	} else {
		_ = d.Set("tenant_id", nil)
	}

	if v := showCloudServicesRes["gateways-onboarding-settings"]; v != nil {
		gatewaysOnboardingSettingsMap := v.(map[string]interface{})
		gatewaysOnboardingSettings := make(map[string]interface{})

		if v := gatewaysOnboardingSettingsMap["connection-method"]; v != nil {
			gatewaysOnboardingSettings["connection_method"] = v.(string)
		}

		if v := gatewaysOnboardingSettingsMap["participant-gateways"]; v != nil {
			gatewaysOnboardingSettings["participant_gateways"] = v.(string)
		}

		if v := gatewaysOnboardingSettingsMap["specific-gateways"]; v != nil {
			specificGatewaysJson, _ := v.([]interface{})
			specificGatewaysRes := make([]string, 0)
			if len(specificGatewaysJson) > 0 {
				for _, gw := range specificGatewaysJson {
					gw := gw.(map[string]interface{})
					specificGatewaysRes = append(specificGatewaysRes, gw["name"].(string))
				}
			}
			gatewaysOnboardingSettings["specific_gateways"] = specificGatewaysRes
		}
		_ = d.Set("gateways_onboarding_settings", []interface{}{gatewaysOnboardingSettings})
	} else {
		_ = d.Set("gateways_onboarding_settings", nil)
	}

	d.SetId("show-cloud-services-" + acctest.RandString(5))

	return nil
}
