package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func dataSourceManagementDefSetting() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementDefSettingRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object unique identifier.",
			},
			"assignments": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Assignments.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description for this setting.",
						},
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "If the setting is enabled.",
						},
						"from_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The gateway version this setting applies from.",
						},
						"model": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The gateway model this setting applies to.",
						},
						"position": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The position of the setting.",
						},
						"targets": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Collection of Gateways identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"to_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The gateway version this setting applies to.",
						},
						"value": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The value of the setting.",
						},
					},
				},
			},
			"custom": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether this is a user-created custom setting or a predefined setting.",
			},
			"data_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The data type of the setting.",
			},
			"global": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether this setting applies globally to all gateways.",
			},
			"color": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Color of the object. Should be one of existing colors.",
			},
			"comments": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Comments string.",
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
		},
	}
}

func dataSourceManagementDefSettingRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{}

	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
	} else if v, ok := d.GetOk("uid"); ok {
		payload["uid"] = v.(string)
	} else {
		return fmt.Errorf("Either name or uid must be specified")
	}

	showDefSettingRes, err := client.ApiCall("show-def-setting", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showDefSettingRes.Success {
		if objectNotFound(showDefSettingRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showDefSettingRes.ErrorMsg)
	}

	defSetting := showDefSettingRes.GetData()

	log.Println("Read DefSetting - Show JSON = ", defSetting)

	if v := defSetting["uid"]; v != nil {
		d.SetId(v.(string))
		_ = d.Set("uid", v)
	}

	if v := defSetting["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if defSetting["assignments"] != nil {

		assignmentsList := defSetting["assignments"].([]interface{})

		if len(assignmentsList) > 0 {

			var assignmentsListToReturn []map[string]interface{}

			for i := range assignmentsList {

				assignmentsMap := assignmentsList[i].(map[string]interface{})

				assignmentsMapToAdd := make(map[string]interface{})

				if v := assignmentsMap["description"]; v != nil {
					assignmentsMapToAdd["description"] = v
				}
				if v := assignmentsMap["enabled"]; v != nil {
					assignmentsMapToAdd["enabled"] = v
				}
				if v := assignmentsMap["from-version"]; v != nil {
					assignmentsMapToAdd["from_version"] = v
				}
				if v := assignmentsMap["model"]; v != nil {
					assignmentsMapToAdd["model"] = v
				}
				if v := assignmentsMap["position"]; v != nil {
					assignmentsMapToAdd["position"] = v
				}
				if v := assignmentsMap["targets"]; v != nil {

					targetsList, ok := v.([]interface{})

					if ok {

						targetsIds := make([]string, 0)

						if len(targetsList) > 0 {
							for _, target := range targetsList {
								target := target.(map[string]interface{})
								targetsIds = append(targetsIds, target["name"].(string))
							}
						}

						assignmentsMapToAdd["targets"] = targetsIds
					}
				}

				if v := assignmentsMap["to-version"]; v != nil {
					assignmentsMapToAdd["to_version"] = v
				}
				if v := assignmentsMap["value"]; v != nil {
					assignmentsMapToAdd["value"] = v
				}

				assignmentsListToReturn = append(assignmentsListToReturn, assignmentsMapToAdd)
			}

			_ = d.Set("assignments", assignmentsListToReturn)
		}
	} else {
		_ = d.Set("assignments", nil)
	}

	if v := defSetting["custom"]; v != nil {
		_ = d.Set("custom", v)
	}

	if v := defSetting["data-type"]; v != nil {
		_ = d.Set("data_type", v)
	}

	if v := defSetting["global"]; v != nil {
		_ = d.Set("global", v)
	}

	if v := defSetting["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := defSetting["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := defSetting["icon"]; v != nil {
		_ = d.Set("icon", v)
	}

	if defSetting["tags"] != nil {
		tagsJson, ok := defSetting["tags"].([]interface{})
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

	if v := defSetting["uid"]; v != nil {
		_ = d.Set("uid", v)
	}

	return nil

}
