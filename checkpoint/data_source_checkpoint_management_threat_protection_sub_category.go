package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func dataSourceManagementThreatProtectionSubCategory() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementThreatProtectionSubCategoryRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The sub-category's name.",
			},
			"category_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The sub-category's unique identifier.",
			},
			"show_profiles": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates whether to calculate and show \"profiles\" field in reply.",
			},
			"category": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Parent category reference.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Parent category unique identifier.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Parent category name.",
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
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Protection settings per profile.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Profile name.",
						},
						"default_action": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Default action applied for this profile.",
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
						"override_action": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Override action applied for this profile if explicitly set.",
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
						"uid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Profile UID.",
						},
					},
				},
			},
			"icon": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object icon.",
			},
		},
	}
}

func dataSourceManagementThreatProtectionSubCategoryRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{}

	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
	} else if v, ok := d.GetOk("uid"); ok {
		payload["uid"] = v.(string)
	}

	if v, ok := d.GetOk("category_id"); ok {
		payload["category-id"] = v.(string)
	}

	if v, ok := d.GetOkExists("show_profiles"); ok {
		payload["show-profiles"] = v.(bool)
	}

	showThreatProtectionSubCategoryRes, err := client.ApiCall("show-threat-protection-sub-category", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showThreatProtectionSubCategoryRes.Success {
		if objectNotFound(showThreatProtectionSubCategoryRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showThreatProtectionSubCategoryRes.ErrorMsg)
	}

	threatProtectionSubCategory := showThreatProtectionSubCategoryRes.GetData()

	log.Println("Read ThreatProtectionSubCategory - Show JSON = ", threatProtectionSubCategory)

	if v := threatProtectionSubCategory["category-id"]; v != nil {
		d.SetId(v.(string))
		_ = d.Set("category_id", v)
	}

	if v := threatProtectionSubCategory["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if threatProtectionSubCategory["category"] != nil {

		categoryMap := threatProtectionSubCategory["category"].(map[string]interface{})

		categoryMapToReturn := make(map[string]interface{})

		if v := categoryMap["id"]; v != nil {
			categoryMapToReturn["id"] = v
		}
		if v := categoryMap["name"]; v != nil {
			categoryMapToReturn["name"] = v
		}

		_ = d.Set("category", []interface{}{categoryMapToReturn})

	} else {
		_ = d.Set("category", nil)
	}

	if v := threatProtectionSubCategory["blade"]; v != nil {
		_ = d.Set("blade", v)
	}

	if v := threatProtectionSubCategory["engine"]; v != nil {
		_ = d.Set("engine", v)
	}

	if v := threatProtectionSubCategory["known-today"]; v != nil {
		_ = d.Set("known_today", v)
	}

	if threatProtectionSubCategory["last-update"] != nil {

		lastUpdateMap := threatProtectionSubCategory["last-update"].(map[string]interface{})

		lastUpdateMapToReturn := make(map[string]interface{})

		if v := lastUpdateMap["iso-8601"]; v != nil {
			lastUpdateMapToReturn["iso_8601"] = v
		}
		if v := lastUpdateMap["posix"]; v != nil {
			lastUpdateMapToReturn["posix"] = v
		}

		_ = d.Set("last_update", []interface{}{lastUpdateMapToReturn})

	} else {
		_ = d.Set("last_update", nil)
	}

	if v := threatProtectionSubCategory["confidence-level"]; v != nil {
		_ = d.Set("confidence_level", v)
	}

	if v := threatProtectionSubCategory["performance-impact"]; v != nil {
		_ = d.Set("performance_impact", v)
	}

	if v := threatProtectionSubCategory["description"]; v != nil {
		_ = d.Set("description", v)
	}

	if threatProtectionSubCategory["profiles"] != nil {

		profilesList := threatProtectionSubCategory["profiles"].([]interface{})

		if len(profilesList) > 0 {

			var profilesListToReturn []map[string]interface{}

			for i := range profilesList {

				profilesMap := profilesList[i].(map[string]interface{})

				profilesMapToAdd := make(map[string]interface{})

				if v := profilesMap["name"]; v != nil {
					profilesMapToAdd["name"] = v
				}
				if v := profilesMap["default-action"]; v != nil {

					defaultActionMap := v.(map[string]interface{})

					defaultActionMapToReturn := make(map[string]interface{})

					if v := defaultActionMap["name"]; v != nil {
						defaultActionMapToReturn["name"] = v
					}
					if v := defaultActionMap["uid"]; v != nil {
						defaultActionMapToReturn["uid"] = v
					}

					profilesMapToAdd["default_action"] = []interface{}{defaultActionMapToReturn}
				}

				if v := profilesMap["override-action"]; v != nil {

					overrideActionMap := v.(map[string]interface{})

					overrideActionMapToReturn := make(map[string]interface{})

					if v := overrideActionMap["name"]; v != nil {
						overrideActionMapToReturn["name"] = v
					}
					if v := overrideActionMap["uid"]; v != nil {
						overrideActionMapToReturn["uid"] = v
					}

					profilesMapToAdd["override_action"] = []interface{}{overrideActionMapToReturn}
				}

				if v := profilesMap["uid"]; v != nil {
					profilesMapToAdd["uid"] = v
				}

				profilesListToReturn = append(profilesListToReturn, profilesMapToAdd)
			}

			_ = d.Set("profiles", profilesListToReturn)
		}
	} else {
		_ = d.Set("profiles", nil)
	}

	if v := threatProtectionSubCategory["icon"]; v != nil {
		_ = d.Set("icon", v)
	}

	return nil

}
