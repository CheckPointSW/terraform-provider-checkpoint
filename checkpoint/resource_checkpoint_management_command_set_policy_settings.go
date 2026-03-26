package checkpoint

import (
	"github.com/CheckPointSW/terraform-provider-checkpoint/upgraders"
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceManagementSetPolicySettings() *schema.Resource {
	return &schema.Resource{
		Create: createManagementSetPolicySettings,
		Read:   readManagementSetPolicySettings,
		Delete: deleteManagementSetPolicySettings,
		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    upgraders.ResourceManagementCommandSetPolicySettingsV0().CoreConfigSchema().ImpliedType(),
				Upgrade: upgraders.ResourceManagementCommandSetPolicySettingsStateUpgradeV0,
				Version: 0,
			},
		},
		Schema: map[string]*schema.Schema{
			"last_in_cell": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Added object after removing the last object in cell.",
			},
			"none_object_behavior": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "'None' object behavior. Rules with object 'None' will never be matched.",
			},
			"security_access_defaults": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Access Policy default values.",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destination": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Destination default value for new rule creation. Any or None.",
						},
						"service": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Service and Applications default value for new rule creation. Any or None.",
						},
						"source": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Source default value for new rule creation. Any or None.",
						},
					},
				},
			},
		},
	}
}

func createManagementSetPolicySettings(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("last_in_cell"); ok {
		payload["last-in-cell"] = v.(string)
	}

	if v, ok := d.GetOk("none_object_behavior"); ok {
		payload["none-object-behavior"] = v.(string)
	}

	if v, ok := d.GetOk("security_access_defaults"); ok {

		securityAccessDefaultsList := v.([]interface{})

		if len(securityAccessDefaultsList) > 0 {

			securityAccessDefaultsPayload := make(map[string]interface{})

			if v, ok := d.GetOk("security_access_defaults.0.destination"); ok {
				securityAccessDefaultsPayload["destination"] = v.(string)
			}
			if v, ok := d.GetOk("security_access_defaults.0.service"); ok {
				securityAccessDefaultsPayload["service"] = v.(string)
			}
			if v, ok := d.GetOk("security_access_defaults.0.source"); ok {
				securityAccessDefaultsPayload["source"] = v.(string)
			}
			if v, ok := d.GetOk("security_access_defaults.0.track"); ok {
				securityAccessDefaultsPayload["track"] = v.(string)
			}
			payload["security-access-defaults"] = securityAccessDefaultsPayload
		}
	}

	SetPolicySettingsRes, _ := client.ApiCall("set-policy-settings", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if !SetPolicySettingsRes.Success {
		return fmt.Errorf(SetPolicySettingsRes.ErrorMsg)
	}

	d.SetId("set-policy-settings-" + acctest.RandString(10))
	return readManagementSetPolicySettings(d, m)
}

func readManagementSetPolicySettings(d *schema.ResourceData, m interface{}) error {

	return nil
}

func deleteManagementSetPolicySettings(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
