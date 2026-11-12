package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func dataSourceManagementRequirement() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementRequirementRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Requirement name.",
			},
			"regulation": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The relevant regulation of the requirement.",
			},
			"best_practices": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "The list of the best practices that make up the requirement. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"score": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The score of the requirement.",
			},
			"score_level": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The score level of the requirement.",
			},
			"user_defined": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Shows if the requirement is user defined.",
			},
			"color": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Color of the object. Should be one of existing colors.",
			},
			"comments": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The requirement comments.",
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
		},
	}
}

func dataSourceManagementRequirementRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{}

	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
	} else if v, ok := d.GetOk("uid"); ok {
		payload["uid"] = v.(string)
	} else {
		return fmt.Errorf("Either name or uid must be specified")
	}

	showRequirementRes, err := client.ApiCall("show-requirement", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showRequirementRes.Success {
		if objectNotFound(showRequirementRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showRequirementRes.ErrorMsg)
	}

	requirement := showRequirementRes.GetData()

	log.Println("Read Requirement - Show JSON = ", requirement)

	if v := requirement["uid"]; v != nil {
		d.SetId(v.(string))
	}

	if v := requirement["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if requirement["best-practices"] != nil {
		bestPracticesJson, ok := requirement["best-practices"].([]interface{})
		if ok {
			bestPracticesIds := make([]string, 0)
			if len(bestPracticesJson) > 0 {
				for _, bestPractices := range bestPracticesJson {
					bestPractices := bestPractices.(map[string]interface{})
					bestPracticesIds = append(bestPracticesIds, bestPractices["name"].(string))
				}
			}
			_ = d.Set("best_practices", bestPracticesIds)
		}
	} else {
		_ = d.Set("best_practices", nil)
	}

	if v := requirement["regulation"]; v != nil {
		if regulationMap, ok := v.(map[string]interface{}); ok {
			_ = d.Set("regulation", regulationMap["name"])
		} else {
			_ = d.Set("regulation", v)
		}
	}

	if v := requirement["score"]; v != nil {
		_ = d.Set("score", v)
	}

	if v := requirement["score-level"]; v != nil {
		_ = d.Set("score_level", v)
	}

	if v := requirement["user-defined"]; v != nil {
		_ = d.Set("user_defined", v)
	}

	if v := requirement["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := requirement["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := requirement["icon"]; v != nil {
		_ = d.Set("icon", v)
	}

	if requirement["tags"] != nil {
		tagsJson, ok := requirement["tags"].([]interface{})
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

	if v := requirement["uid"]; v != nil {
		_ = d.Set("uid", v)
	}

	return nil

}
