package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementPolicySettings() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementPolicySettingsRead,
		Schema: map[string]*schema.Schema{
			"last_in_cell": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Added object after removing the last object in cell.",
			},
			"none_object_behavior": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "'None' object behavior. Rules with object 'None' will never be matched.",
			},
			"security_access_defaults": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Access Policy default values.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destination": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Destination default value identified by name.",
						},
						"service": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Service and Applications default value identified by name.",
						},
						"source": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Source default value identified by name.",
						},
					},
				},
			},
		},
	}
}

func dataSourceManagementPolicySettingsRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	payload := make(map[string]interface{})

	showPolicySettingsRes, err := client.ApiCall("show-policy-settings", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showPolicySettingsRes.Success {
		return fmt.Errorf(showPolicySettingsRes.ErrorMsg)
	}

	policySettings := showPolicySettingsRes.GetData()

	log.Println("Read Policy Settings - Show JSON = ", policySettings)

	d.SetId("show-policy-settings-" + acctest.RandString(10))

	if v := policySettings["last-in-cell"]; v != nil {
		_ = d.Set("last_in_cell", v)
	}

	if v := policySettings["none-object-behavior"]; v != nil {
		_ = d.Set("none_object_behavior", v)
	}

	if policySettings["security-access-defaults"] != nil {
		securityAccessDefaultsMap := policySettings["security-access-defaults"].(map[string]interface{})

		securityAccessDefaultsMapToReturn := make(map[string]interface{})

		if v, _ := securityAccessDefaultsMap["destination"]; v != nil {
			securityAccessDefaultsMapToReturn["destination"] = v
		}
		if v, _ := securityAccessDefaultsMap["service"]; v != nil {
			securityAccessDefaultsMapToReturn["service"] = v
		}
		if v, _ := securityAccessDefaultsMap["source"]; v != nil {
			securityAccessDefaultsMapToReturn["source"] = v
		}

		_ = d.Set("security_access_defaults", securityAccessDefaultsMapToReturn)
	} else {
		_ = d.Set("security_access_defaults", nil)
	}

	return nil
}
