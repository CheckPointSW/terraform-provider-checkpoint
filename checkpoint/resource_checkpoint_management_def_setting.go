package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"

	"strconv"
)

func resourceManagementDefSetting() *schema.Resource {
	return &schema.Resource{
		Create: createManagementDefSetting,
		Read:   readManagementDefSetting,
		Update: updateManagementDefSetting,
		Delete: deleteManagementDefSetting,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"data_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The data type of the setting.",
			},
			"assignments": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "Assignments.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The value of the setting.",
						},
						"description": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The description for this setting.",
						},
						"enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "If the setting is enabled.",
							Default:     true,
						},
						"from_version": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The gateway version this setting applies from.",
							Default:     "earliest",
						},
						"model": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The gateway model this setting applies to.",
							Default:     "all",
						},
						"position": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Position in the rulebase.",
						},
						"targets": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "The Gateways or Clusters the assignment is applied to, identified by name or UID.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"to_version": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The gateway version this setting applies to.",
							Default:     "latest",
						},
					},
				},
			},
			"color": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Color of the object. Should be one of existing colors.",
				Default:     "black",
			},
			"comments": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comments string.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ignore_warnings": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Apply changes ignoring warnings.",
				Default:     false,
			},
			"ignore_errors": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.",
				Default:     false,
			},
		},
	}
}

func createManagementDefSetting(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	defSetting := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		defSetting["name"] = v.(string)
	}

	if v, ok := d.GetOk("data_type"); ok {
		defSetting["data-type"] = v.(string)
	}

	if v, ok := d.GetOk("assignments"); ok {

		assignmentsList := v.([]interface{})

		if len(assignmentsList) > 0 {

			var assignmentsPayload []map[string]interface{}

			for i := range assignmentsList {

				Payload := make(map[string]interface{})

				if v, ok := d.GetOk("assignments." + strconv.Itoa(i) + ".value"); ok {
					Payload["value"] = v.(string)
				}
				if v, ok := d.GetOk("assignments." + strconv.Itoa(i) + ".description"); ok {
					Payload["description"] = v.(string)
				}
				if v, ok := d.GetOkExists("assignments." + strconv.Itoa(i) + ".enabled"); ok {
					Payload["enabled"] = v.(bool)
				}
				if v, ok := d.GetOk("assignments." + strconv.Itoa(i) + ".from_version"); ok {
					Payload["from-version"] = v.(string)
				}
				if v, ok := d.GetOk("assignments." + strconv.Itoa(i) + ".model"); ok {
					Payload["model"] = v.(string)
				}
				if v, ok := d.GetOk("assignments." + strconv.Itoa(i) + ".position"); ok {
					Payload["position"] = v.(string)
				}
				if v, ok := d.GetOk("assignments." + strconv.Itoa(i) + ".targets"); ok {
					Payload["targets"] = v.(*schema.Set).List()
				}
				if v, ok := d.GetOk("assignments." + strconv.Itoa(i) + ".to_version"); ok {
					Payload["to-version"] = v.(string)
				}
				assignmentsPayload = append(assignmentsPayload, Payload)
			}
			defSetting["assignments"] = assignmentsPayload
		}
	}

	if v, ok := d.GetOk("color"); ok {
		defSetting["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		defSetting["comments"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		defSetting["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		defSetting["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		defSetting["ignore-errors"] = v.(bool)
	}

	log.Println("Create DefSetting - Map = ", defSetting)

	addDefSettingRes, err := client.ApiCall("add-def-setting", defSetting, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !addDefSettingRes.Success {
		if addDefSettingRes.ErrorMsg != "" {
			return fmt.Errorf(addDefSettingRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addDefSettingRes.GetData()["uid"].(string))

	return readManagementDefSetting(d, m)
}

func readManagementDefSetting(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showDefSettingRes, err := client.ApiCall("show-def-setting", payload, client.GetSessionID(), true, client.IsProxyUsed())
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

	if v := defSetting["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := defSetting["data-type"]; v != nil {
		_ = d.Set("data_type", v)
	}

	if defSetting["assignments"] != nil {

		assignmentsList := defSetting["assignments"].([]interface{})

		if len(assignmentsList) > 0 {

			var assignmentsListToReturn []map[string]interface{}

			for i := range assignmentsList {

				assignmentsMap := assignmentsList[i].(map[string]interface{})

				assignmentsMapToAdd := make(map[string]interface{})

				if v, _ := assignmentsMap["value"]; v != nil {
					assignmentsMapToAdd["value"] = v
				}
				if v, _ := assignmentsMap["description"]; v != nil {
					assignmentsMapToAdd["description"] = v
				}
				if v, _ := assignmentsMap["enabled"]; v != nil {
					assignmentsMapToAdd["enabled"] = v
				}
				if v, _ := assignmentsMap["from-version"]; v != nil {
					assignmentsMapToAdd["from_version"] = v
				}
				if v, _ := assignmentsMap["model"]; v != nil {
					assignmentsMapToAdd["model"] = v
				}
				if v, _ := assignmentsMap["position"]; v != nil {
					assignmentsMapToAdd["position"] = v
				}
				if v, _ := assignmentsMap["targets"]; v != nil {
					targetsJson, ok := v.([]interface{})
					if ok {
						targetsIds := make([]string, 0)
						if len(targetsJson) > 0 {
							for _, target := range targetsJson {
								target := target.(map[string]interface{})
								targetsIds = append(targetsIds, target["name"].(string))
							}
						}
						assignmentsMapToAdd["targets"] = targetsIds
					}
				}
				if v, _ := assignmentsMap["to-version"]; v != nil {
					assignmentsMapToAdd["to_version"] = v
				}
				assignmentsListToReturn = append(assignmentsListToReturn, assignmentsMapToAdd)
			}

			_ = d.Set("assignments", assignmentsListToReturn)
		} else {
			_ = d.Set("assignments", assignmentsList)
		}
	} else {
		_ = d.Set("assignments", nil)
	}

	if v := defSetting["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := defSetting["comments"]; v != nil {
		_ = d.Set("comments", v)
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

	if v := defSetting["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := defSetting["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementDefSetting(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	defSetting := make(map[string]interface{})

	defSetting["uid"] = d.Id()

	if ok := d.HasChange("name"); ok {
		if v, ok := d.GetOk("name"); ok {
			defSetting["new-name"] = v.(string)
		}
	}

	if ok := d.HasChange("data_type"); ok {
		defSetting["data-type"] = d.Get("data_type")
	}

	if d.HasChange("assignments") {

		if v, ok := d.GetOk("assignments"); ok {

			assignmentsList := v.([]interface{})

			var assignmentsPayload []map[string]interface{}

			for i := range assignmentsList {

				Payload := make(map[string]interface{})

				if v, ok := d.GetOk("assignments." + strconv.Itoa(i) + ".value"); ok {
					Payload["value"] = v
				}
				if v, ok := d.GetOk("assignments." + strconv.Itoa(i) + ".description"); ok {
					Payload["description"] = v
				}
				if v, ok := d.GetOkExists("assignments." + strconv.Itoa(i) + ".enabled"); ok {
					Payload["enabled"] = v
				}
				if v, ok := d.GetOk("assignments." + strconv.Itoa(i) + ".from_version"); ok {
					Payload["from-version"] = v
				}
				if v, ok := d.GetOk("assignments." + strconv.Itoa(i) + ".model"); ok {
					Payload["model"] = v
				}
				if v, ok := d.GetOk("assignments." + strconv.Itoa(i) + ".position"); ok {
					Payload["position"] = v
				}
				if v, ok := d.GetOk("assignments." + strconv.Itoa(i) + ".targets"); ok {
					Payload["targets"] = v
				}
				if v, ok := d.GetOk("assignments." + strconv.Itoa(i) + ".to_version"); ok {
					Payload["to-version"] = v
				}
				assignmentsPayload = append(assignmentsPayload, Payload)
			}
			defSetting["assignments"] = assignmentsPayload
		}
	}

	if ok := d.HasChange("color"); ok {
		defSetting["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		defSetting["comments"] = d.Get("comments")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			defSetting["tags"] = v.(*schema.Set).List()
		}
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		defSetting["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		defSetting["ignore-errors"] = v.(bool)
	}

	log.Println("Update DefSetting - Map = ", defSetting)

	updateDefSettingRes, err := client.ApiCall("set-def-setting", defSetting, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !updateDefSettingRes.Success {
		if updateDefSettingRes.ErrorMsg != "" {
			return fmt.Errorf(updateDefSettingRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementDefSetting(d, m)
}

func deleteManagementDefSetting(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	defSettingPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete DefSetting")

	deleteDefSettingRes, err := client.ApiCall("delete-def-setting", defSettingPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteDefSettingRes.Success {
		if deleteDefSettingRes.ErrorMsg != "" {
			return fmt.Errorf(deleteDefSettingRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
