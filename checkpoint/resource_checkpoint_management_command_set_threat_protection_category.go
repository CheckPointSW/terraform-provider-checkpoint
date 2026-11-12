package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func resourceManagementSetThreatProtectionCategory() *schema.Resource {
	return &schema.Resource{
		Create: createManagementSetThreatProtectionCategory,
		Read:   readManagementSetThreatProtectionCategory,
		Delete: deleteManagementSetThreatProtectionCategory,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The Category name.",
			},
			"category_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The Category unique identifier.",
			},
			"blade": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The blade this category belongs to. Required when using 'name'.",
			},
			"show_profiles": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Indicates whether to calculate and show \"profiles\" field in reply.",
			},
			"all_profiles": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Apply action to all profiles. Default: true.",
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
						"uid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Profile UID.",
						},
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
							Description: "Override action applied for this profile.",
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
				},
			},
		},
	}
}

func createManagementSetThreatProtectionCategory(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
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

	if v, ok := d.GetOkExists("all_profiles"); ok {
		payload["all-profiles"] = v.(bool)
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

	SetThreatProtectionCategoryRes, err := client.ApiCall("set-threat-protection-category", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !SetThreatProtectionCategoryRes.Success {
		return fmt.Errorf(SetThreatProtectionCategoryRes.ErrorMsg)
	}

	d.SetId("set-threat-protection-category-" + acctest.RandString(10))
	if v := SetThreatProtectionCategoryRes.GetData()["category-id"]; v != nil {
		_ = d.Set("category_id", v)
	}
	if v := SetThreatProtectionCategoryRes.GetData()["name"]; v != nil {
		_ = d.Set("name", v)
	}
	if v := SetThreatProtectionCategoryRes.GetData()["blade"]; v != nil {
		_ = d.Set("blade", v)
	}
	if v := SetThreatProtectionCategoryRes.GetData()["engine"]; v != nil {
		_ = d.Set("engine", v)
	}
	if v := SetThreatProtectionCategoryRes.GetData()["known-today"]; v != nil {
		_ = d.Set("known_today", v)
	}
	if v := SetThreatProtectionCategoryRes.GetData()["last-update"]; v != nil {
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
	if v := SetThreatProtectionCategoryRes.GetData()["confidence-level"]; v != nil {
		if clMap, ok := v.(map[string]interface{}); ok {
			confidenceLevel := make(map[string]interface{})
			if val := clMap["high"]; val != nil {
				confidenceLevel["high"] = val
			}
			if val := clMap["low"]; val != nil {
				confidenceLevel["low"] = val
			}
			if val := clMap["medium"]; val != nil {
				confidenceLevel["medium"] = val
			}
			_ = d.Set("confidence_level", []interface{}{confidenceLevel})
		}
	}
	if v := SetThreatProtectionCategoryRes.GetData()["performance-impact"]; v != nil {
		if piMap, ok := v.(map[string]interface{}); ok {
			performanceImpact := make(map[string]interface{})
			if val := piMap["high"]; val != nil {
				performanceImpact["high"] = val
			}
			if val := piMap["low"]; val != nil {
				performanceImpact["low"] = val
			}
			if val := piMap["medium"]; val != nil {
				performanceImpact["medium"] = val
			}
			_ = d.Set("performance_impact", []interface{}{performanceImpact})
		}
	}
	if v := SetThreatProtectionCategoryRes.GetData()["description"]; v != nil {
		_ = d.Set("description", v)
	}
	if v := SetThreatProtectionCategoryRes.GetData()["profiles"]; v != nil {
		if profilesList, ok := v.([]interface{}); ok {
			var profilesToReturn []map[string]interface{}
			for _, profile := range profilesList {
				profileMap, ok := profile.(map[string]interface{})
				if !ok {
					continue
				}
				profileToAdd := make(map[string]interface{})
				if val := profileMap["uid"]; val != nil {
					profileToAdd["uid"] = val
				}
				if val := profileMap["name"]; val != nil {
					profileToAdd["name"] = val
				}
				if val := profileMap["default-action"]; val != nil {
					if actionMap, ok := val.(map[string]interface{}); ok {
						action := make(map[string]interface{})
						if n := actionMap["name"]; n != nil {
							action["name"] = n
						}
						if u := actionMap["uid"]; u != nil {
							action["uid"] = u
						}
						profileToAdd["default_action"] = []interface{}{action}
					}
				}
				if val := profileMap["override-action"]; val != nil {
					if actionMap, ok := val.(map[string]interface{}); ok {
						action := make(map[string]interface{})
						if n := actionMap["name"]; n != nil {
							action["name"] = n
						}
						if u := actionMap["uid"]; u != nil {
							action["uid"] = u
						}
						profileToAdd["override_action"] = []interface{}{action}
					}
				}
				profilesToReturn = append(profilesToReturn, profileToAdd)
			}
			_ = d.Set("profiles", profilesToReturn)
		}
	}
	if v := SetThreatProtectionCategoryRes.GetData()["domain"]; v != nil {
		_ = d.Set("domain", v)
	}
	return nil
}

func readManagementSetThreatProtectionCategory(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementSetThreatProtectionCategory(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
