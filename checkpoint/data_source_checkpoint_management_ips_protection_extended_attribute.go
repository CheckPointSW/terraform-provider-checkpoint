package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementIpsProtectionExtendedAttribute() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementIpsProtectionExtendedAttributeRead,
		Schema: map[string]*schema.Schema{
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object unique identifier.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"values": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "N/A",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object name. Must be unique in the domain.",
						},
						"uid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object unique identifier.",
						},
					},
				},
			},
		},
	}
}

func dataSourceManagementIpsProtectionExtendedAttributeRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showIpsProtectionExtendedAttributeRes, err := client.ApiCall("show-ips-protection-extended-attribute", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showIpsProtectionExtendedAttributeRes.Success {
		return fmt.Errorf(showIpsProtectionExtendedAttributeRes.ErrorMsg)
	}

	ipsProtectionExtendedAttribute := showIpsProtectionExtendedAttributeRes.GetData()

	log.Println("Read Ips Protection Extended Attribute - Show JSON = ", ipsProtectionExtendedAttribute)

	if ipsProtectionExtendedAttribute["object"] != nil {
		objectMap := ipsProtectionExtendedAttribute["object"].(map[string]interface{})

		if v, _ := objectMap["name"]; v != nil {
			_ = d.Set("name", v)
		}
		if v, _ := objectMap["uid"]; v != nil {
			_ = d.Set("uid", v)
			d.SetId(v.(string))
		}

		var valuesListToReturn []map[string]interface{}

		if objectMap["values"] != nil {
			valuesList := objectMap["values"].([]interface{})

			if len(valuesList) > 0 {

				for i := range valuesList {

					valuesMap := valuesList[i].(map[string]interface{})

					valuesMapToAdd := make(map[string]interface{})

					if v, _ := valuesMap["name"]; v != nil {
						valuesMapToAdd["name"] = v
					}
					if v, _ := valuesMap["uid"]; v != nil {
						valuesMapToAdd["uid"] = v
					}
					valuesListToReturn = append(valuesListToReturn, valuesMapToAdd)
				}
			}
		}
		_ = d.Set("values", valuesListToReturn)
	} else {
		_ = d.Set("values", nil)
	}

	return nil
}
