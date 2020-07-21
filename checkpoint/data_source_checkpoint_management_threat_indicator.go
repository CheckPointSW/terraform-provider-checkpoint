package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementThreatIndicator() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementThreatIndicatorRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name. Should be unique in the domain.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object unique identifier.",
			},
			"action": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The indicator's action.",
			},
			"profile_overrides": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Profiles in which to override the indicator's default action.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The indicator's action in this profile.",
						},
						"profile": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The profile in which to override the indicator's action.",
						},
					},
				},
			},
			"color": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Color of the object. Should be one of existing colors.",
			},
			"comments": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Comments string.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceManagementThreatIndicatorRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showThreatIndicatorRes, err := client.ApiCall("show-threat-indicator", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showThreatIndicatorRes.Success {
		return fmt.Errorf(showThreatIndicatorRes.ErrorMsg)
	}

	threatIndicator := showThreatIndicatorRes.GetData()

	log.Println("Read Threat Indicator - Show JSON = ", threatIndicator)

	if v := threatIndicator["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := threatIndicator["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := threatIndicator["action"]; v != nil {
		_ = d.Set("action", v)
	}

	if threatIndicator["profile-overrides"] != nil {

		profileOverridesList := threatIndicator["profile-overrides"].([]interface{})

		if len(profileOverridesList) > 0 {

			var profileOverridesListToReturn []map[string]interface{}

			for i := range profileOverridesList {

				profileOverridesMap := profileOverridesList[i].(map[string]interface{})

				profileOverridesMapToAdd := make(map[string]interface{})

				if v, _ := profileOverridesMap["action"]; v != nil {
					profileOverridesMapToAdd["action"] = v
				}
				if v, _ := profileOverridesMap["profile"]; v != nil {
					profileOverridesMapToAdd["profile"] = v
				}

				profileOverridesListToReturn = append(profileOverridesListToReturn, profileOverridesMapToAdd)
			}
			_ = d.Set("profile_overrides", profileOverridesListToReturn)
		} else {
			_ = d.Set("interfaces", profileOverridesList)
		}
	} else {
		_ = d.Set("profile_overrides", nil)
	}

	if v := threatIndicator["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := threatIndicator["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if threatIndicator["tags"] != nil {
		tagsJson := threatIndicator["tags"].([]interface{})
		var tagsIds = make([]string, 0)
		if len(tagsJson) > 0 {
			// Create slice of tag names
			for _, tag := range tagsJson {
				tag := tag.(map[string]interface{})
				tagsIds = append(tagsIds, tag["name"].(string))
			}
		}
		_ = d.Set("tags", tagsIds)
	} else {
		_ = d.Set("tags", nil)
	}

	return nil
}
