package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementNatRule() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementNatRuleRead,
		Schema: map[string]*schema.Schema{
			"package": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the package.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Rule name.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Rule UID.",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable/Disable the rule.",
			},
			"install_on": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Which Gateways identified by the name or UID to install the policy on.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"method": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Nat method.",
			},
			"original_destination": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Original destination.",
			},
			"original_service": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Original service.",
			},
			"original_source": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Original source.",
			},
			"translated_destination": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Translated destination.",
			},
			"translated_service": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Translated service.",
			},
			"translated_source": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Translated source.",
			},
			"auto_generated": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Auto generated.",
			},
			"comments": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Comments string.",
			},
		},
	}
}

func dataSourceManagementNatRuleRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := map[string]interface{}{
		"package": d.Get("package"),
	}

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showNatRuleRes, err := client.ApiCall("show-nat-rule", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showNatRuleRes.Success {
		return fmt.Errorf(showNatRuleRes.ErrorMsg)
	}

	natRule := showNatRuleRes.GetData()

	log.Println("Read NAT Rule - Show JSON = ", natRule)

	if v := natRule["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := natRule["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := natRule["auto-generated"]; v != nil {
		_ = d.Set("auto_generated", v)
	}

	if v := natRule["enabled"]; v != nil {
		_ = d.Set("enabled", v)
	}

	if v := natRule["method"]; v != nil {
		_ = d.Set("method", v)
	}

	if natRule["install-on"] != nil {
		installOnJson := natRule["install-on"].([]interface{})
		installOnJsonIds := make([]string, 0)
		if len(installOnJson) > 0 {
			for _, installOn := range installOnJson {
				installOn := installOn.(map[string]interface{})
				installOnJsonIds = append(installOnJsonIds, installOn["name"].(string))
			}
		}
		_, installOnInConf := d.GetOk("install_on")
		if installOnJsonIds[0] == "Policy Targets" && !installOnInConf {
			_ = d.Set("install_on", []interface{}{})
		} else {
			_ = d.Set("install_on", installOnJsonIds)
		}
	}

	if v := natRule["original-destination"]; v != nil {
		_ = d.Set("original_destination", v.(map[string]interface{})["name"])
	}

	if v := natRule["original-service"]; v != nil {
		_ = d.Set("original_service", v.(map[string]interface{})["name"])
	}

	if v := natRule["original-source"]; v != nil {
		_ = d.Set("original_source", v.(map[string]interface{})["name"])
	}

	if v := natRule["translated-destination"]; v != nil {
		_ = d.Set("translated_destination", v.(map[string]interface{})["name"])
	}

	if v := natRule["translated-service"]; v != nil {
		_ = d.Set("translated_service", v.(map[string]interface{})["name"])
	}

	if v := natRule["translated-source"]; v != nil {
		_ = d.Set("translated_source", v.(map[string]interface{})["name"])
	}

	if v := natRule["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil
}
