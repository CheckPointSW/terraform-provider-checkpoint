package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
)

func resourceManagementThreatIndicator() *schema.Resource {
	return &schema.Resource{
		Create: createManagementThreatIndicator,
		Read:   readManagementThreatIndicator,
		Update: updateManagementThreatIndicator,
		Delete: deleteManagementThreatIndicator,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name. Should be unique in the domain.",
			},
			"observables": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The indicator's observables.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Object name. Should be unique in the domain.",
						},
						"md5": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A valid MD5 sequence.",
						},
						"url": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A valid URL.",
						},
						"ip_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A valid IP-Address.",
						},
						"ip_address_first": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A valid IP-Address, the beginning of the range. If you configure this parameter with a value, you must also configure the value of the 'ip-address-last' parameter",
						},
						"ip_address_last": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A valid IP-Address, the end of the range. If you configure this parameter with a value, you must also configure the value of the 'ip-address-first' parameter.",
						},
						"domain": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of a domain.",
						},
						"mail_to": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A valid E-Mail address, recipient filed.",
						},
						"mail_from": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A valid E-Mail address, sender field.",
						},
						"mail_cc": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A valid E-Mail address, cc field.",
						},
						"mail_reply_to": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A valid E-Mail address, reply-to field.",
						},
						"mail_subject": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Subject of E-Mail.",
						},
						"confidence": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The confidence level the indicator has that a real threat has been uncovered.",
						},
						"product": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The software blade that processes the observable: AV - AntiVirus, AB - AntiBot.",
						},
						"severity": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The severity level of the threat.",
						},
					},
				},
			},
			"action": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The indicator's action.",
				Default:     "Prevent",
			},
			"profile_overrides": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Profiles in which to override the indicator's default action.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The indicator's action in this profile.",
						},
						"profile": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The profile in which to override the indicator's action.",
						},
					},
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
			"color": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Color of the object. Should be one of existing colors.",
				Default:     "black",
			},
			"comments": &schema.Schema{
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
		},
	}
}

func createManagementThreatIndicator(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	threatIndicator := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		threatIndicator["name"] = v.(string)
	}

	if v, ok := d.GetOk("observables"); ok {

		observablesList := v.([]interface{})
		if len(observablesList) > 0 {

			var observablesPayload []map[string]interface{}

			for i := range observablesList {

				payload := make(map[string]interface{})

				if v, ok := d.GetOk("observables." + strconv.Itoa(i) + ".name"); ok {
					payload["name"] = v.(string)
				}
				if v, ok := d.GetOk("observables." + strconv.Itoa(i) + ".md5"); ok {
					payload["md5"] = v.(string)
				}
				if v, ok := d.GetOk("observables." + strconv.Itoa(i) + ".url"); ok {
					payload["url"] = v.(string)
				}
				if v, ok := d.GetOk("observables." + strconv.Itoa(i) + ".ip_address"); ok {
					payload["ip-address"] = v.(string)
				}
				if v, ok := d.GetOk("observables." + strconv.Itoa(i) + ".ip_address_first"); ok {
					payload["ip-address-first"] = v.(string)
				}
				if v, ok := d.GetOk("observables." + strconv.Itoa(i) + ".ip_address_last"); ok {
					payload["ip-address-last"] = v.(string)
				}
				if v, ok := d.GetOk("observables." + strconv.Itoa(i) + ".domain"); ok {
					payload["domain"] = v.(string)
				}
				if v, ok := d.GetOk("observables." + strconv.Itoa(i) + ".mail_to"); ok {
					payload["mail-to"] = v.(string)
				}
				if v, ok := d.GetOk("observables." + strconv.Itoa(i) + ".mail_from"); ok {
					payload["mail-from"] = v.(string)
				}
				if v, ok := d.GetOk("observables." + strconv.Itoa(i) + ".mail_cc"); ok {
					payload["mail-cc"] = v.(string)
				}
				if v, ok := d.GetOk("observables." + strconv.Itoa(i) + ".mail_reply_to"); ok {
					payload["mail-reply-to"] = v.(string)
				}
				if v, ok := d.GetOk("observables." + strconv.Itoa(i) + ".mail_subject"); ok {
					payload["mail-subject"] = v.(string)
				}
				if v, ok := d.GetOk("observables." + strconv.Itoa(i) + ".confidence"); ok {
					payload["confidence"] = v.(string)
				}
				if v, ok := d.GetOk("observables." + strconv.Itoa(i) + ".product"); ok {
					payload["product"] = v.(string)
				}
				if v, ok := d.GetOk("observables." + strconv.Itoa(i) + ".severity"); ok {
					payload["severity"] = v.(string)
				}
				observablesPayload = append(observablesPayload, payload)
			}
			threatIndicator["observables"] = observablesPayload
		}
	}

	if v, ok := d.GetOk("action"); ok {
		threatIndicator["action"] = v.(string)
	}

	if v, ok := d.GetOk("profile_overrides"); ok {

		profileOverridesList := v.([]interface{})
		if len(profileOverridesList) > 0 {

			var profileOverridesPayload []map[string]interface{}

			for i := range profileOverridesList {

				payload := make(map[string]interface{})

				if v, ok := d.GetOk("profile_overrides." + strconv.Itoa(i) + ".action"); ok {
					payload["action"] = v.(string)
				}
				if v, ok := d.GetOk("profile_overrides." + strconv.Itoa(i) + ".profile"); ok {
					payload["profile"] = v.(string)
				}
				profileOverridesPayload = append(profileOverridesPayload, payload)
			}
			threatIndicator["profile-overrides"] = profileOverridesPayload
		}
	}

	if val, ok := d.GetOk("comments"); ok {
		threatIndicator["comments"] = val.(string)
	}
	if val, ok := d.GetOk("tags"); ok {
		threatIndicator["tags"] = val.(*schema.Set).List()
	}
	if val, ok := d.GetOk("color"); ok {
		threatIndicator["color"] = val.(string)
	}
	if val, ok := d.GetOkExists("ignore_errors"); ok {
		threatIndicator["ignore-errors"] = val.(bool)
	}
	if val, ok := d.GetOkExists("ignore_warnings"); ok {
		threatIndicator["ignore-warnings"] = val.(bool)
	}

	log.Println("Create Threat Indicator - Map = ", threatIndicator)

	threatIndicatorRes, err := client.ApiCall("add-threat-indicator", threatIndicator, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !threatIndicatorRes.Success {
		if threatIndicatorRes.ErrorMsg != "" {
			return fmt.Errorf(threatIndicatorRes.ErrorMsg)
		}

		taskDetails := threatIndicatorRes.GetData()["tasks"].([]interface{})[0].(map[string]interface{})["task-details"].([]interface{})
		errStr := "Status: " + taskDetails[0].(map[string]interface{})["request-status"].(string)
		errStr += "\nDescription: " + taskDetails[0].(map[string]interface{})["request-status-description"].(string)
		return fmt.Errorf(errStr)
	}

	//special section because of the unique type of threat indicator
	//task-id is back but we need object ID
	payload := map[string]interface{}{
		"name": d.Get("name"),
	}

	showThreatIndicatorRes, err := client.ApiCall("show-threat-indicator", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showThreatIndicatorRes.Success {
		// Handle delete resource from other clients
		if objectNotFound(showThreatIndicatorRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showThreatIndicatorRes.ErrorMsg)
	}
	d.SetId(showThreatIndicatorRes.GetData()["uid"].(string))

	return readManagementThreatIndicator(d, m)
}

func readManagementThreatIndicator(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"name": d.Get("name"),
	}

	showThreatIndicatorRes, err := client.ApiCall("show-threat-indicator", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showThreatIndicatorRes.Success {
		// Handle delete resource from other clients
		if objectNotFound(showThreatIndicatorRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showThreatIndicatorRes.ErrorMsg)
	}

	threatIndicator := showThreatIndicatorRes.GetData()

	log.Println("Read Threat Indicator - Show JSON = ", threatIndicator)

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

func updateManagementThreatIndicator(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	threatIndicator := make(map[string]interface{})

	if d.HasChange("name") {
		oldName, newName := d.GetChange("name")
		threatIndicator["name"] = oldName
		threatIndicator["new-name"] = newName
	} else {
		threatIndicator["name"] = d.Get("name")
	}

	if ok := d.HasChange("action"); ok {
		threatIndicator["action"] = d.Get("action")
	}

	if d.HasChange("profile_overrides") {

		if v, ok := d.GetOk("profile_overrides"); ok {

			profileOverridesList := v.([]interface{})

			if len(profileOverridesList) > 0 {

				var profileOverridesPayload []map[string]interface{}

				for i := range profileOverridesList {

					payload := make(map[string]interface{})

					if v, ok := d.GetOk("profile_overrides." + strconv.Itoa(i) + ".action"); ok {
						payload["action"] = v.(string)
					}
					if v, ok := d.GetOk("profile_overrides." + strconv.Itoa(i) + ".profile"); ok {
						payload["profile"] = v.(string)
					}
					profileOverridesPayload = append(profileOverridesPayload, payload)
				}
				threatIndicator["profile-overrides"] = profileOverridesPayload
			}

		} else { //delete all of the list
			oldProfileOverride, _ := d.GetChange("profile_overrides")
			var profileOverridesToDelete []interface{}
			for _, profile := range oldProfileOverride.([]interface{}) {
				profileOverridesToDelete = append(profileOverridesToDelete, profile.(map[string]interface{})["profile"].(string))
			}
			threatIndicator["profile-overrides"] = map[string]interface{}{"remove": profileOverridesToDelete}
		}
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		threatIndicator["ignore-errors"] = v.(bool)
	}
	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		threatIndicator["ignore-warnings"] = v.(bool)
	}

	if ok := d.HasChange("comments"); ok {
		threatIndicator["comments"] = d.Get("comments")
	}
	if ok := d.HasChange("color"); ok {
		threatIndicator["color"] = d.Get("color")
	}

	if ok := d.HasChange("tags"); ok {
		if v, ok := d.GetOk("tags"); ok {
			threatIndicator["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			threatIndicator["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	log.Println("Update Threat Indicator - Map = ", threatIndicator)
	updateThreatIndicatorRes, err := client.ApiCall("set-threat-indicator", threatIndicator, client.GetSessionID(), true, false)
	if err != nil || !updateThreatIndicatorRes.Success {
		if updateThreatIndicatorRes.ErrorMsg != "" {
			return fmt.Errorf(updateThreatIndicatorRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementThreatIndicator(d, m)
}

func deleteManagementThreatIndicator(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	threatIndicatorPayload := map[string]interface{}{
		"name": d.Get("name"),
	}

	deleteThreatIndicatorRes, err := client.ApiCall("delete-threat-indicator", threatIndicatorPayload, client.GetSessionID(), true, false)
	if err != nil || !deleteThreatIndicatorRes.Success {
		if deleteThreatIndicatorRes.ErrorMsg != "" {
			return fmt.Errorf(deleteThreatIndicatorRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
