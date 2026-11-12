package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func dataSourceManagementThreatProtectionCategory() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementThreatProtectionCategoryRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Category name.",
			},
			"category_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Category unique identifier.",
			},
			"blade": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The blade this category belongs to. Required when using 'name'.",
			},
			"show_profiles": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates whether to calculate and show \"profiles\" field in reply.",
			},
			"engine": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The engine that handles this category.",
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
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Confidence levels.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"high": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Count of protections classified with high level.",
						},
						"low": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Count of protections classified with low level.",
						},
						"medium": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Count of protections classified with medium level.",
						},
					},
				},
			},
			"performance_impact": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Performance impacts.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"high": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Count of protections classified with high level.",
						},
						"low": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Count of protections classified with low level.",
						},
						"medium": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Count of protections classified with medium level.",
						},
					},
				},
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description of the category.",
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

func dataSourceManagementThreatProtectionCategoryRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{}

	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
	} else if v, ok := d.GetOk("uid"); ok {
		payload["uid"] = v.(string)
	} else {
		return fmt.Errorf("Either name or uid must be specified")
	}

	if v, ok := d.GetOk("category_id"); ok {
		payload["category-id"] = v.(string)
	}

	if v, ok := d.GetOk("blade"); ok {
		payload["blade"] = v.(string)
	}

	if v, ok := d.GetOkExists("show_profiles"); ok {
		payload["show-profiles"] = v.(bool)
	}

	showThreatProtectionCategoryRes, err := client.ApiCall("show-threat-protection-category", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showThreatProtectionCategoryRes.Success {
		if objectNotFound(showThreatProtectionCategoryRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showThreatProtectionCategoryRes.ErrorMsg)
	}

	threatProtectionCategory := showThreatProtectionCategoryRes.GetData()

	log.Println("Read ThreatProtectionCategory - Show JSON = ", threatProtectionCategory)

	if v := threatProtectionCategory["category-id"]; v != nil {
		d.SetId(v.(string))
		_ = d.Set("category_id", v)
	}

	if v := threatProtectionCategory["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := threatProtectionCategory["blade"]; v != nil {
		_ = d.Set("blade", v)
	}

	if v := threatProtectionCategory["engine"]; v != nil {
		_ = d.Set("engine", v)
	}

	if v := threatProtectionCategory["known-today"]; v != nil {
		_ = d.Set("known_today", v)
	}

	if threatProtectionCategory["last-update"] != nil {

		lastUpdateMap := threatProtectionCategory["last-update"].(map[string]interface{})

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

	if threatProtectionCategory["confidence-level"] != nil {

		confidenceLevelMap := threatProtectionCategory["confidence-level"].(map[string]interface{})

		confidenceLevelMapToReturn := make(map[string]interface{})

		if v := confidenceLevelMap["high"]; v != nil {
			confidenceLevelMapToReturn["high"] = v
		}
		if v := confidenceLevelMap["low"]; v != nil {
			confidenceLevelMapToReturn["low"] = v
		}
		if v := confidenceLevelMap["medium"]; v != nil {
			confidenceLevelMapToReturn["medium"] = v
		}

		_ = d.Set("confidence_level", []interface{}{confidenceLevelMapToReturn})

	} else {
		_ = d.Set("confidence_level", nil)
	}

	if threatProtectionCategory["performance-impact"] != nil {

		performanceImpactMap := threatProtectionCategory["performance-impact"].(map[string]interface{})

		performanceImpactMapToReturn := make(map[string]interface{})

		if v := performanceImpactMap["high"]; v != nil {
			performanceImpactMapToReturn["high"] = v
		}
		if v := performanceImpactMap["low"]; v != nil {
			performanceImpactMapToReturn["low"] = v
		}
		if v := performanceImpactMap["medium"]; v != nil {
			performanceImpactMapToReturn["medium"] = v
		}

		_ = d.Set("performance_impact", []interface{}{performanceImpactMapToReturn})

	} else {
		_ = d.Set("performance_impact", nil)
	}

	if v := threatProtectionCategory["description"]; v != nil {
		_ = d.Set("description", v)
	}

	if threatProtectionCategory["profiles"] != nil {

		profilesList := threatProtectionCategory["profiles"].([]interface{})

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

	if v := threatProtectionCategory["icon"]; v != nil {
		_ = d.Set("icon", v)
	}

	return nil

}
