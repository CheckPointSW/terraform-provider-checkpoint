package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func dataSourceManagementRegulation() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementRegulationRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Regulation name.",
			},
			"full_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Regulation full name.",
			},
			"show_requirements": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Show the requirements of the regulation.",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Shows if the regulation is enabled.",
			},
			"score": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The regulation score.",
			},
			"user_defined": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Shows if the regulation is user defined.",
			},
			"comments": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Compliance Regulation comments.",
			},
			"color": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Color of the object. Should be one of existing colors.",
			},
			"icon": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object icon.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of tag objects identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object unique identifier.",
			},
			"requirements": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The requirements of the regulation, identified by name. Appears only when the value of the 'show-requirements' parameter is set to 'true'.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceManagementRegulationRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{}

	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
	} else if v, ok := d.GetOk("uid"); ok {
		payload["uid"] = v.(string)
	} else {
		return fmt.Errorf("Either name or uid must be specified")
	}

	if v, ok := d.GetOk("full_name"); ok {
		payload["full-name"] = v.(string)
	}

	if v, ok := d.GetOkExists("show_requirements"); ok {
		payload["show-requirements"] = v.(bool)
	}

	showRegulationRes, err := client.ApiCall("show-regulation", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showRegulationRes.Success {
		if objectNotFound(showRegulationRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showRegulationRes.ErrorMsg)
	}

	regulation := showRegulationRes.GetData()

	log.Println("Read Regulation - Show JSON = ", regulation)

	if v := regulation["uid"]; v != nil {
		d.SetId(v.(string))
	}

	if v := regulation["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := regulation["enabled"]; v != nil {
		_ = d.Set("enabled", v)
	}

	if v := regulation["full-name"]; v != nil {
		_ = d.Set("full_name", v)
	}

	if regulation["requirements"] != nil {
		requirementsJson, ok := regulation["requirements"].([]interface{})
		if ok {
			requirementsIds := make([]string, 0)
			if len(requirementsJson) > 0 {
				for _, requirement := range requirementsJson {
					if requirementMap, ok := requirement.(map[string]interface{}); ok {
						requirementsIds = append(requirementsIds, requirementMap["name"].(string))
					} else {
						requirementsIds = append(requirementsIds, requirement.(string))
					}
				}
			}
			_ = d.Set("requirements", requirementsIds)
		}
	} else {
		_ = d.Set("requirements", nil)
	}

	if v := regulation["score"]; v != nil {
		_ = d.Set("score", v)
	}

	if v := regulation["user-defined"]; v != nil {
		_ = d.Set("user_defined", v)
	}

	if v := regulation["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := regulation["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := regulation["icon"]; v != nil {
		_ = d.Set("icon", v)
	}

	if regulation["tags"] != nil {
		tagsJson, ok := regulation["tags"].([]interface{})
		if ok {
			tagsIds := make([]string, 0)
			if len(tagsJson) > 0 {
				for _, tags := range tagsJson {
					tags := tags.(map[string]interface{})
					tagsIds = append(tagsIds, tags["name"].(string))
				}
			}
			_ = d.Set("tags", tagsIds)
		}
	} else {
		_ = d.Set("tags", nil)
	}

	if v := regulation["uid"]; v != nil {
		_ = d.Set("uid", v)
	}

	return nil

}
