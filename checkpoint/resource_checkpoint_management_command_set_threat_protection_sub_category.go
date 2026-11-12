package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func resourceManagementSetThreatProtectionSubCategory() *schema.Resource {
	return &schema.Resource{
		Create: createManagementSetThreatProtectionSubCategory,
		Read:   readManagementSetThreatProtectionSubCategory,
		Delete: deleteManagementSetThreatProtectionSubCategory,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The sub-category's name.",
			},
			"category_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The sub-category's unique identifier.",
			},
			"all_profiles": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Apply action to all profiles. Default: true.",
			},
			"show_profiles": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Indicates whether to calculate and show \"profiles\" field in reply.",
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Action to apply to all profiles. Required when all-profiles is true.",
			},
			"overrides": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Overrides per profile for this protection. Required when all-profiles is false.",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Action to apply for the specified profile.",
						},
						"profile": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Profile name or UID.",
						},
					},
				},
			},
			"category": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Parent category reference.",
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
			"blade": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The blade this protection belongs to.",
			},
			"engine": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The engine that handles this protection.",
			},
			"known_today": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The current number of protection items available in the latest update.",
			},
			"last_update": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The date in which the protection was updated by Check Point.",
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
			"confidence_level": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Confidence level of the protection.",
			},
			"performance_impact": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Performance impact of the protection.",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Detailed description.",
			},
			"profiles": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed:    true,
				Description: "Protection settings per profile.",
			},
			"domain": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed:    true,
				Description: "Information about the domain that holds the Object.",
			},
		},
	}
}

func createManagementSetThreatProtectionSubCategory(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
	}

	if v, ok := d.GetOk("category_id"); ok {
		payload["category-id"] = v.(string)
	}

	if v, ok := d.GetOkExists("all_profiles"); ok {
		payload["all-profiles"] = v.(bool)
	}

	if v, ok := d.GetOkExists("show_profiles"); ok {
		payload["show-profiles"] = v.(bool)
	}

	if v, ok := d.GetOk("action"); ok {
		payload["action"] = v.(string)
	}

	if v, ok := d.GetOk("overrides"); ok {

		overridesList := v.([]interface{})

		if len(overridesList) > 0 {

			var overridesPayload []map[string]interface{}

			for i := range overridesList {

				Payload := make(map[string]interface{})

				if v, ok := d.GetOk("overrides." + strconv.Itoa(i) + ".action"); ok {
					Payload["action"] = v.(string)
				}
				if v, ok := d.GetOk("overrides." + strconv.Itoa(i) + ".profile"); ok {
					Payload["profile"] = v.(string)
				}
				overridesPayload = append(overridesPayload, Payload)
			}
			payload["overrides"] = overridesPayload
		}
	}

	SetThreatProtectionSubCategoryRes, err := client.ApiCall("set-threat-protection-sub-category", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !SetThreatProtectionSubCategoryRes.Success {
		return fmt.Errorf(SetThreatProtectionSubCategoryRes.ErrorMsg)
	}

	d.SetId("set-threat-protection-sub-category-" + acctest.RandString(10))
	if v := SetThreatProtectionSubCategoryRes.GetData()["category-id"]; v != nil {
		_ = d.Set("category_id", v)
	}
	if v := SetThreatProtectionSubCategoryRes.GetData()["name"]; v != nil {
		_ = d.Set("name", v)
	}
	if v := SetThreatProtectionSubCategoryRes.GetData()["category"]; v != nil {
		if categoryMap, ok := v.(map[string]interface{}); ok {
			category := make(map[string]interface{})
			if n := categoryMap["name"]; n != nil {
				category["name"] = n
			}
			if u := categoryMap["uid"]; u != nil {
				category["uid"] = u
			}
			_ = d.Set("category", []interface{}{category})
		}
	}
	if v := SetThreatProtectionSubCategoryRes.GetData()["blade"]; v != nil {
		_ = d.Set("blade", v)
	}
	if v := SetThreatProtectionSubCategoryRes.GetData()["engine"]; v != nil {
		_ = d.Set("engine", v)
	}
	if v := SetThreatProtectionSubCategoryRes.GetData()["known-today"]; v != nil {
		_ = d.Set("known_today", v)
	}
	if v := SetThreatProtectionSubCategoryRes.GetData()["last-update"]; v != nil {
		if lastUpdateMap, ok := v.(map[string]interface{}); ok {
			lastUpdate := make(map[string]interface{})
			if val := lastUpdateMap["iso-8601"]; val != nil {
				lastUpdate["iso_8601"] = val
			}
			if val := lastUpdateMap["posix"]; val != nil {
				lastUpdate["posix"] = val
			}
			_ = d.Set("last_update", []interface{}{lastUpdate})
		}
	}
	if v := SetThreatProtectionSubCategoryRes.GetData()["confidence-level"]; v != nil {
		_ = d.Set("confidence_level", v)
	}
	if v := SetThreatProtectionSubCategoryRes.GetData()["performance-impact"]; v != nil {
		_ = d.Set("performance_impact", v)
	}
	if v := SetThreatProtectionSubCategoryRes.GetData()["description"]; v != nil {
		_ = d.Set("description", v)
	}
	if v := SetThreatProtectionSubCategoryRes.GetData()["profiles"]; v != nil {
		_ = d.Set("profiles", v)
	}
	if v := SetThreatProtectionSubCategoryRes.GetData()["domain"]; v != nil {
		_ = d.Set("domain", v)
	}
	return nil
}

func readManagementSetThreatProtectionSubCategory(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementSetThreatProtectionSubCategory(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
