package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementSetTrust() *schema.Resource {
	return &schema.Resource{
		Create: createManagementSetTrust,
		Read:   readManagementSetTrust,
		Delete: deleteManagementSetTrust,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Object name.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Object unique identifier.",
			},
			"ipv4_address": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "IP address of the object, for establishing trust with dynamic gateways.",
			},
			"one_time_password": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Shared password to establish SIC between the Security Management and the Security Gateway.",
			},
			"trust_method": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Establish the trust communication method.",
			},
			"trust_settings": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Settings for the trusted communication establishment.",
				ForceNew:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gateway_mac_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Use the Security Gateway MAC address, relevant for the gateway_mac_address identification-method.",
						},
						"identification_method": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "How to identify the gateway (relevant for Spark DAIP gateways only).",
						},
						"initiation_phase": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Push the certificate to the Security Gateway immediately, or wait for the Security Gateway to pull the certificate. Default value for Spark Gateway is 'when_gateway_connects'.",
						},
					},
				},
			},
		},
	}
}

func createManagementSetTrust(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}

	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
	}

	if v, ok := d.GetOk("uid"); ok {
		payload["uid"] = v.(string)
	}

	if v, ok := d.GetOk("ipv4_address"); ok {
		payload["ipv4-address"] = v.(string)
	}

	if v, ok := d.GetOk("one_time_password"); ok {
		payload["one-time-password"] = v.(string)
	}

	if v, ok := d.GetOk("trust_method"); ok {
		payload["trust-method"] = v.(string)
	}

	if v, ok := d.GetOk("trust_settings"); ok {

		trustSettingsList := v.([]interface{})

		if len(trustSettingsList) > 0 {

			trustSettingsPayload := make(map[string]interface{})

			if v, ok := d.GetOk("trust_settings.0.gateway_mac_address"); ok {
				trustSettingsPayload["gateway-mac-address"] = v.(string)
			}
			if v, ok := d.GetOk("trust_settings.0.identification_method"); ok {
				trustSettingsPayload["identification-method"] = v.(string)
			}
			if v, ok := d.GetOk("trust_settings.0.initiation_phase"); ok {
				trustSettingsPayload["initiation-phase"] = v.(string)
			}
			payload["trust-settings"] = trustSettingsPayload
		}
	}

	SetTrustRes, err := client.ApiCall("set-trust", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !SetTrustRes.Success {
		return fmt.Errorf(SetTrustRes.ErrorMsg)
	}

	d.SetId("set-trust-" + acctest.RandString(10))
	return nil
}

func readManagementSetTrust(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementSetTrust(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
