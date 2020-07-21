package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementHttpsRule() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementHttpsRuleRead,
		Schema: map[string]*schema.Schema{
			"layer": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Layer that holds the Object. Identified by the Name or UID.",
			},
			"rule_number": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "HTTPS rule number.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object unique identifier.",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object name.",
			},
			"destination": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of Network objects identified by Name or UID that represents connection destination.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"service": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of Network objects identified by Name or UID that represents connection service.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"source": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of Network objects identified by Name or UID that represents connection source.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"action": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Rule inspect level. \"Bypass\" or \"Inspect\".",
			},
			"blade": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Blades for HTTPS Inspection. Identified by Name or UID to enable the inspection for. \"Anti Bot\",\"Anti Virus\",\"Application Control\",\"Data Awareness\",\"DLP\",\"IPS\",\"Threat Emulation\",\"Url Filtering\".",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"certificate": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Internal Server Certificate identified by Name or UID, otherwise, \"Outbound Certificate\" is a default value.",
			},
			"destination_negate": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "TRUE if \"negate\" value is set for Destination.",
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
			"service_negate": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "TRUE if \"negate\" value is set for Service.",
			},
			"site_category": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of Site Categories objects identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"site_category_negate": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "TRUE if \"negate\" value is set for Site Category.",
			},
			"source_negate": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "TRUE if \"negate\" value is set for Source.",
			},
			"track": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "\"None\",\"Log\",\"Alert\",\"Mail\",\"SNMP trap\",\"Mail\",\"User Alert\", \"User Alert 2\", \"User Alert 3\".",
			},
			"comments": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Comments string.",
			},
		},
	}
}

func dataSourceManagementHttpsRuleRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	ruleNumber := d.Get("rule_number").(string)
	uid := d.Get("uid").(string)

	payload := map[string]interface{}{
		"layer": d.Get("layer"),
	}

	if ruleNumber != "" {
		payload["rule-number"] = ruleNumber
	} else if uid != "" {
		payload["uid"] = uid
	}

	showHttpsRuleRes, err := client.ApiCall("show-https-rule", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showHttpsRuleRes.Success {
		return fmt.Errorf(showHttpsRuleRes.ErrorMsg)
	}

	httpsRule := showHttpsRuleRes.GetData()

	log.Println("Read HttpsRule - Show JSON = ", httpsRule)

	if v := httpsRule["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := httpsRule["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if httpsRule["destination"] != nil {
		destinationJson, ok := httpsRule["destination"].([]interface{})
		if ok {
			destinationIds := make([]string, 0)
			if len(destinationJson) > 0 {
				for _, destination := range destinationJson {
					destination := destination.(map[string]interface{})
					destinationIds = append(destinationIds, destination["name"].(string))
				}
			}
			_ = d.Set("destination", destinationIds)
		}
	} else {
		_ = d.Set("destination", nil)
	}

	if httpsRule["service"] != nil {
		serviceJson, ok := httpsRule["service"].([]interface{})
		if ok {
			serviceIds := make([]string, 0)
			if len(serviceJson) > 0 {
				for _, service := range serviceJson {
					service := service.(map[string]interface{})
					serviceIds = append(serviceIds, service["name"].(string))
				}
			}
			_ = d.Set("service", serviceIds)
		}
	} else {
		_ = d.Set("service", nil)
	}

	if httpsRule["source"] != nil {
		sourceJson, ok := httpsRule["source"].([]interface{})
		if ok {
			sourceIds := make([]string, 0)
			if len(sourceJson) > 0 {
				for _, source := range sourceJson {
					source := source.(map[string]interface{})
					sourceIds = append(sourceIds, source["name"].(string))
				}
			}
			_ = d.Set("source", sourceIds)
		}
	} else {
		_ = d.Set("source", nil)
	}

	if v := httpsRule["action"]; v != nil {
		_ = d.Set("action", v)
	}

	if httpsRule["blade"] != nil {
		bladeJson, ok := httpsRule["blade"].([]interface{})
		if ok {
			bladeIds := make([]string, 0)
			if len(bladeJson) > 0 {
				for _, blade := range bladeJson {
					blade := blade.(map[string]interface{})
					bladeIds = append(bladeIds, blade["name"].(string))
				}
			}
			_ = d.Set("blade", bladeIds)
		}
	} else {
		_ = d.Set("blade", nil)
	}

	if v := httpsRule["certificate"]; v != nil {
		_ = d.Set("certificate", v)
	}

	if v := httpsRule["destination-negate"]; v != nil {
		_ = d.Set("destination_negate", v)
	}

	if v := httpsRule["enabled"]; v != nil {
		_ = d.Set("enabled", v)
	}

	if httpsRule["install_on"] != nil {
		installOnJson, ok := httpsRule["install_on"].([]interface{})
		if ok {
			installOnIds := make([]string, 0)
			if len(installOnJson) > 0 {
				for _, install_on := range installOnJson {
					install_on := install_on.(map[string]interface{})
					installOnIds = append(installOnIds, install_on["name"].(string))
				}
			}
			_ = d.Set("install_on", installOnIds)
		}
	} else {
		_ = d.Set("install_on", nil)
	}

	if v := httpsRule["service-negate"]; v != nil {
		_ = d.Set("service_negate", v)
	}

	if httpsRule["site_category"] != nil {
		siteCategoryJson, ok := httpsRule["site_category"].([]interface{})
		if ok {
			siteCategoryIds := make([]string, 0)
			if len(siteCategoryJson) > 0 {
				for _, site_category := range siteCategoryJson {
					site_category := site_category.(map[string]interface{})
					siteCategoryIds = append(siteCategoryIds, site_category["name"].(string))
				}
			}
			_ = d.Set("site_category", siteCategoryIds)
		}
	} else {
		_ = d.Set("site_category", nil)
	}

	if v := httpsRule["site-category-negate"]; v != nil {
		_ = d.Set("site_category_negate", v)
	}

	if v := httpsRule["source-negate"]; v != nil {
		_ = d.Set("source_negate", v)
	}

	if v := httpsRule["track"]; v != nil {
		_ = d.Set("track", v)
	}

	if v := httpsRule["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil
}
