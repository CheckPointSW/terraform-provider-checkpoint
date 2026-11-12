package checkpoint

import (
	"github.com/CheckPointSW/terraform-provider-checkpoint/v3/upgraders"
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceManagementConnectCloudServices() *schema.Resource {
	return &schema.Resource{
		Create: createManagementConnectCloudServices,
		Read:   readManagementConnectCloudServices,
		Delete: deleteManagementConnectCloudServices,
		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    upgraders.ResourceManagementCommandConnectCloudServicesV0().CoreConfigSchema().ImpliedType(),
				Upgrade: upgraders.ResourceManagementCommandConnectCloudServicesStateUpgradeV0,
				Version: 0,
			},
		},
		Schema: map[string]*schema.Schema{
			"auth_token": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Sensitive:   true,
				Description: "Copy the authentication token from the Smart-1 cloud service hosted in the Infinity Portal.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status of the connection to the Infinity Portal.",
			},
			"connected_at": {
				Type:        schema.TypeList,
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
			"gateways_onboarding_settings": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				ForceNew:    true,
				Description: "Gateways on-boarding to Infinity Portal settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connection_method": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Indicate whether Gateways will be connected to Infinity Portal automatically or only after policy installation.",
						},
						"details_level": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation o...",
						},
						"enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable/Disable automatic connection of Security Gateways to Infinity Portal.",
						},
						"participant_gateways": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Which Gateways will be connected to Infinity Portal.",
						},
						"specific_gateways": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Selection of targets identified by the name or UID which will be on-boarded to the cloud. Configuration will be applied only when 'participant-gateway...",
						},
					},
				},
			},
		},
	}
}

func createManagementConnectCloudServices(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}

	if v, ok := d.GetOk("auth_token"); ok {
		payload["auth-token"] = v.(string)
	}

	if _, ok := d.GetOk("gateways_onboarding_settings"); ok {
		gatewaysOnboardingSettingsPayload := make(map[string]interface{})
		if v, ok := d.GetOk("gateways_onboarding_settings.0.connection_method"); ok {
			gatewaysOnboardingSettingsPayload["connection-method"] = v.(string)
		}
		if v, ok := d.GetOk("gateways_onboarding_settings.0.details_level"); ok {
			gatewaysOnboardingSettingsPayload["details-level"] = v.(string)
		}
		if v, ok := d.GetOkExists("gateways_onboarding_settings.0.enabled"); ok {
			gatewaysOnboardingSettingsPayload["enabled"] = v.(bool)
		}
		if v, ok := d.GetOk("gateways_onboarding_settings.0.participant_gateways"); ok {
			gatewaysOnboardingSettingsPayload["participant-gateways"] = v.(string)
		}
		if v, ok := d.GetOk("gateways_onboarding_settings.0.specific_gateways"); ok {
			gatewaysOnboardingSettingsPayload["specific-gateways"] = v.(string)
		}
		payload["gateways-onboarding-settings"] = gatewaysOnboardingSettingsPayload
	}

	ConnectCloudServicesRes, err := client.ApiCall("connect-cloud-services", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	if !ConnectCloudServicesRes.Success {
		return fmt.Errorf("%s", ConnectCloudServicesRes.ErrorMsg)
	}

	connectCloudServicesRes := ConnectCloudServicesRes.GetData()

	log.Println("Connect Cloud Services - JSON = ", connectCloudServicesRes)

	d.SetId("connect-cloud-services" + acctest.RandString(5))

	if v := connectCloudServicesRes["status"]; v != nil {
		_ = d.Set("status", v)
	} else {
		_ = d.Set("status", nil)
	}

	if v := connectCloudServicesRes["connected-at"]; v != nil {
		connectedAtShow := connectCloudServicesRes["connected-at"].(map[string]interface{})
		connectedAtState := make(map[string]interface{})
		if v := connectedAtShow["iso-8601"]; v != nil {
			connectedAtState["iso_8601"] = v
		}
		if v := connectedAtShow["posix"]; v != nil {
			connectedAtState["posix"] = v
		}
		_ = d.Set("connected_at", []interface{}{connectedAtState})
	} else {
		_ = d.Set("connected_at", nil)
	}

	if v := connectCloudServicesRes["management-url"]; v != nil {
		_ = d.Set("management_url", v)
	} else {
		_ = d.Set("management_url", nil)
	}

	return readManagementConnectCloudServices(d, m)
}

func readManagementConnectCloudServices(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementConnectCloudServices(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
