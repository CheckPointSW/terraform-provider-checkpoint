package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"

	"strconv"
)

func resourceManagementThreatIocFeed() *schema.Resource {
	return &schema.Resource{
		Create: createManagementThreatIocFeed,
		Read:   readManagementThreatIocFeed,
		Update: updateManagementThreatIocFeed,
		Delete: deleteManagementThreatIocFeed,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"feed_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "URL of the feed. URL should be written as http or https.",
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The feed indicator's action.",
				Default:     "Prevent",
			},
			"certificate_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Certificate SHA-1 fingerprint to access the feed.",
			},
			"custom_comment": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Custom IOC feed - the column number of comment.",
			},
			"custom_confidence": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Custom IOC feed - the column number of confidence.",
			},
			"custom_header": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Custom HTTP headers.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"header_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the HTTP header we wish to add.",
						},
						"header_value": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the HTTP value we wish to add.",
						},
					},
				},
			},
			"custom_name": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Custom IOC feed - the column number of name.",
			},
			"custom_severity": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Custom IOC feed - the column number of severity.",
			},
			"custom_type": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Custom IOC feed - the column number of type in case a specific type is not chosen.",
			},
			"custom_value": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Custom IOC feed - the column number of value in case a specific type is chosen.",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Sets whether this indicator feed is enabled.",
				Default:     true,
			},
			"feed_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Feed type to be enforced.",
				Default:     "domain",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "password for authenticating with the URL.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"use_custom_feed_settings": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Set in order to configure a custom indicator feed.",
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "username for authenticating with the URL.",
			},
			"fields_delimiter": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The delimiter that separates between the columns in the feed.",
			},
			"ignore_lines_that_start_with": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A prefix that will determine which lines to ignore.",
			},
			"use_gateway_proxy": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Use the gateway's proxy for retrieving the feed.",
				Default:     true,
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

func createManagementThreatIocFeed(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	threatIocFeed := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		threatIocFeed["name"] = v.(string)
	}

	if v, ok := d.GetOk("feed_url"); ok {
		threatIocFeed["feed-url"] = v.(string)
	}

	if v, ok := d.GetOk("action"); ok {
		threatIocFeed["action"] = v.(string)
	}

	if v, ok := d.GetOk("certificate_id"); ok {
		threatIocFeed["certificate-id"] = v.(string)
	}

	if v, ok := d.GetOk("custom_comment"); ok {
		threatIocFeed["custom-comment"] = v.(int)
	}

	if v, ok := d.GetOk("custom_confidence"); ok {
		threatIocFeed["custom-confidence"] = v.(int)
	}

	if v, ok := d.GetOk("custom_header"); ok {

		customHeaderList := v.([]interface{})

		if len(customHeaderList) > 0 {

			var customHeaderPayload []map[string]interface{}

			for i := range customHeaderList {

				Payload := make(map[string]interface{})

				if v, ok := d.GetOk("custom_header." + strconv.Itoa(i) + ".header_name"); ok {
					Payload["header-name"] = v.(string)
				}
				if v, ok := d.GetOk("custom_header." + strconv.Itoa(i) + ".header_value"); ok {
					Payload["header-value"] = v.(string)
				}
				customHeaderPayload = append(customHeaderPayload, Payload)
			}
			threatIocFeed["custom-header"] = customHeaderPayload
		}
	}

	if v, ok := d.GetOk("custom_name"); ok {
		threatIocFeed["custom-name"] = v.(int)
	}

	if v, ok := d.GetOk("custom_severity"); ok {
		threatIocFeed["custom-severity"] = v.(int)
	}

	if v, ok := d.GetOk("custom_type"); ok {
		threatIocFeed["custom-type"] = v.(int)
	}

	if v, ok := d.GetOk("custom_value"); ok {
		threatIocFeed["custom-value"] = v.(int)
	}

	if v, ok := d.GetOkExists("enabled"); ok {
		threatIocFeed["enabled"] = v.(bool)
	}

	if v, ok := d.GetOk("feed_type"); ok {
		threatIocFeed["feed-type"] = v.(string)
	}

	if v, ok := d.GetOk("password"); ok {
		threatIocFeed["password"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		threatIocFeed["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("use_custom_feed_settings"); ok {
		threatIocFeed["use-custom-feed-settings"] = v.(bool)
	}

	if v, ok := d.GetOk("username"); ok {
		threatIocFeed["username"] = v.(string)
	}

	if v, ok := d.GetOk("fields_delimiter"); ok {
		threatIocFeed["fields-delimiter"] = v.(string)
	}

	if v, ok := d.GetOk("ignore_lines_that_start_with"); ok {
		threatIocFeed["ignore-lines-that-start-with"] = v.(string)
	}

	if v, ok := d.GetOkExists("use_gateway_proxy"); ok {
		threatIocFeed["use-gateway-proxy"] = v.(bool)
	}

	if v, ok := d.GetOk("color"); ok {
		threatIocFeed["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		threatIocFeed["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		threatIocFeed["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		threatIocFeed["ignore-errors"] = v.(bool)
	}

	log.Println("Create ThreatIocFeed - Map = ", threatIocFeed)

	addThreatIocFeedRes, err := client.ApiCall("add-threat-ioc-feed", threatIocFeed, client.GetSessionID(), true, false)
	if err != nil || !addThreatIocFeedRes.Success {
		if addThreatIocFeedRes.ErrorMsg != "" {
			return fmt.Errorf(addThreatIocFeedRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addThreatIocFeedRes.GetData()["uid"].(string))

	return readManagementThreatIocFeed(d, m)
}

func readManagementThreatIocFeed(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showThreatIocFeedRes, err := client.ApiCall("show-threat-ioc-feed", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showThreatIocFeedRes.Success {
		if objectNotFound(showThreatIocFeedRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showThreatIocFeedRes.ErrorMsg)
	}

	threatIocFeed := showThreatIocFeedRes.GetData()

	log.Println("Read ThreatIocFeed - Show JSON = ", threatIocFeed)

	if v := threatIocFeed["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := threatIocFeed["feed-url"]; v != nil {
		_ = d.Set("feed_url", v)
	}

	if v := threatIocFeed["action"]; v != nil {
		_ = d.Set("action", v)
	}

	if v := threatIocFeed["certificate-id"]; v != nil {
		_ = d.Set("certificate_id", v)
	}

	if v := threatIocFeed["custom-comment"]; v != nil {
		_ = d.Set("custom_comment", v)
	}

	if v := threatIocFeed["custom-confidence"]; v != nil {
		_ = d.Set("custom_confidence", v)
	}
	if threatIocFeed["custom-headers"] != nil {
		customHeaderList, ok := threatIocFeed["custom-headers"].([]interface{})

		if ok {
			var customHeaderListToReturn []map[string]interface{}
			if len(customHeaderList) > 0 {

				for i := range customHeaderList {

					customHeaderMap := customHeaderList[i].(map[string]interface{})

					customHeaderMapToAdd := make(map[string]interface{})

					if v, _ := customHeaderMap["headerName"]; v != nil {
						customHeaderMapToAdd["header_name"] = v
					}
					if v, _ := customHeaderMap["headerValue"]; v != nil {
						customHeaderMapToAdd["header_value"] = v
					}
					customHeaderListToReturn = append(customHeaderListToReturn, customHeaderMapToAdd)
				}
			}
			err = d.Set("custom_header", customHeaderListToReturn)
		}
	}

	if v := threatIocFeed["custom-name"]; v != nil {
		_ = d.Set("custom_name", v)
	}

	if v := threatIocFeed["custom-severity"]; v != nil {
		_ = d.Set("custom_severity", v)
	}

	if v := threatIocFeed["custom-type"]; v != nil {
		_ = d.Set("custom_type", v)
	}

	if v := threatIocFeed["custom-value"]; v != nil {
		_ = d.Set("custom_value", v)
	}

	if v := threatIocFeed["enabled"]; v != nil {
		_ = d.Set("enabled", v)
	}

	if v := threatIocFeed["feed-type"]; v != nil {
		_ = d.Set("feed_type", v)
	}

	if v := threatIocFeed["password"]; v != nil {
		_ = d.Set("password", v)
	}

	if threatIocFeed["tags"] != nil {
		tagsJson, ok := threatIocFeed["tags"].([]interface{})
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

	if v := threatIocFeed["use-custom-feed-settings"]; v != nil {
		_ = d.Set("use_custom_feed_settings", v)
	}

	if v := threatIocFeed["username"]; v != nil {
		_ = d.Set("username", v)
	}

	if v := threatIocFeed["fields-delimiter"]; v != nil {
		_ = d.Set("fields_delimiter", v)
	}

	if v := threatIocFeed["ignore-lines-that-start-with"]; v != nil {
		_ = d.Set("ignore_lines_that_start_with", v)
	}

	if v := threatIocFeed["use-gateway-proxy"]; v != nil {
		_ = d.Set("use_gateway_proxy", v)
	}

	if v := threatIocFeed["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := threatIocFeed["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := threatIocFeed["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := threatIocFeed["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementThreatIocFeed(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	threatIocFeed := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		threatIocFeed["name"] = oldName
		threatIocFeed["new-name"] = newName
	} else {
		threatIocFeed["name"] = d.Get("name")
	}

	if ok := d.HasChange("feed_url"); ok {
		threatIocFeed["feed-url"] = d.Get("feed_url")
	}

	if ok := d.HasChange("action"); ok {
		threatIocFeed["action"] = d.Get("action")
	}

	if ok := d.HasChange("certificate_id"); ok {
		threatIocFeed["certificate-id"] = d.Get("certificate_id")
	}

	if ok := d.HasChange("custom_comment"); ok {
		threatIocFeed["custom-comment"] = d.Get("custom_comment")
	}

	if ok := d.HasChange("custom_confidence"); ok {
		threatIocFeed["custom-confidence"] = d.Get("custom_confidence")
	}

	if d.HasChange("custom_header") {

		if v, ok := d.GetOk("custom_header"); ok {

			customHeaderList := v.([]interface{})

			var customHeaderPayload []map[string]interface{}

			for i := range customHeaderList {

				Payload := make(map[string]interface{})

				if d.HasChange("custom_header." + strconv.Itoa(i) + ".header_name") {
					Payload["header-name"] = d.Get("custom_header." + strconv.Itoa(i) + ".header_name")
				}
				if d.HasChange("custom_header." + strconv.Itoa(i) + ".header_value") {
					Payload["header-value"] = d.Get("custom_header." + strconv.Itoa(i) + ".header_value")
				}
				customHeaderPayload = append(customHeaderPayload, Payload)
			}
			threatIocFeed["custom-header"] = customHeaderPayload
		} else {
			oldcustomHeader, _ := d.GetChange("custom_header")
			var customHeaderToDelete []interface{}
			for _, i := range oldcustomHeader.([]interface{}) {
				customHeaderToDelete = append(customHeaderToDelete, i.(map[string]interface{})["name"].(string))
			}
			threatIocFeed["custom-header"] = map[string]interface{}{"remove": customHeaderToDelete}
		}
	}

	if ok := d.HasChange("custom_name"); ok {
		threatIocFeed["custom-name"] = d.Get("custom_name")
	}

	if ok := d.HasChange("custom_severity"); ok {
		threatIocFeed["custom-severity"] = d.Get("custom_severity")
	}

	if ok := d.HasChange("custom_type"); ok {
		threatIocFeed["custom-type"] = d.Get("custom_type")
	}

	if ok := d.HasChange("custom_value"); ok {
		threatIocFeed["custom-value"] = d.Get("custom_value")
	}

	if v, ok := d.GetOkExists("enabled"); ok {
		threatIocFeed["enabled"] = v.(bool)
	}

	if ok := d.HasChange("feed_type"); ok {
		threatIocFeed["feed-type"] = d.Get("feed_type")
	}

	if ok := d.HasChange("password"); ok {
		threatIocFeed["password"] = d.Get("password")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			threatIocFeed["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			threatIocFeed["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if v, ok := d.GetOkExists("use_custom_feed_settings"); ok {
		threatIocFeed["use-custom-feed-settings"] = v.(bool)
	}

	if ok := d.HasChange("username"); ok {
		threatIocFeed["username"] = d.Get("username")
	}

	if ok := d.HasChange("fields_delimiter"); ok {
		threatIocFeed["fields-delimiter"] = d.Get("fields_delimiter")
	}

	if ok := d.HasChange("ignore_lines_that_start_with"); ok {
		threatIocFeed["ignore-lines-that-start-with"] = d.Get("ignore_lines_that_start_with")
	}

	if v, ok := d.GetOkExists("use_gateway_proxy"); ok {
		threatIocFeed["use-gateway-proxy"] = v.(bool)
	}

	if ok := d.HasChange("color"); ok {
		threatIocFeed["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		threatIocFeed["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		threatIocFeed["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		threatIocFeed["ignore-errors"] = v.(bool)
	}

	log.Println("Update ThreatIocFeed - Map = ", threatIocFeed)

	updateThreatIocFeedRes, err := client.ApiCall("set-threat-ioc-feed", threatIocFeed, client.GetSessionID(), true, false)
	if err != nil || !updateThreatIocFeedRes.Success {
		if updateThreatIocFeedRes.ErrorMsg != "" {
			return fmt.Errorf(updateThreatIocFeedRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementThreatIocFeed(d, m)
}

func deleteManagementThreatIocFeed(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	threatIocFeedPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete ThreatIocFeed")

	deleteThreatIocFeedRes, err := client.ApiCall("delete-threat-ioc-feed", threatIocFeedPayload, client.GetSessionID(), true, false)
	if err != nil || !deleteThreatIocFeedRes.Success {
		if deleteThreatIocFeedRes.ErrorMsg != "" {
			return fmt.Errorf(deleteThreatIocFeedRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
