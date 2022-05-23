package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementThreatIocFeed() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementThreatIocFeedRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"feed_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL of the feed. URL should be written as http or https.",
			},
			"action": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The feed indicator's action.",
			},
			"certificate_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Certificate SHA-1 fingerprint to access the feed.",
			},
			"custom_comment": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Custom IOC feed - the column number of comment.",
			},
			"custom_confidence": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Custom IOC feed - the column number of confidence.",
			},
			"custom_headers": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Custom HTTP headers.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"header_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the HTTP header we wish to add.",
						},
						"header_value": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the HTTP value we wish to add.",
						},
					},
				},
			},
			"custom_name": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Custom IOC feed - the column number of name.",
			},
			"custom_severity": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Custom IOC feed - the column number of severity.",
			},
			"custom_type": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Custom IOC feed - the column number of type in case a specific type is not chosen.",
			},
			"custom_value": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Custom IOC feed - the column number of value in case a specific type is chosen.",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Sets whether this indicator feed is enabled.",
			},
			"feed_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Feed type to be enforced.",
			},
			"password": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "password for authenticating with the URL.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"use_custom_feed_settings": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Set in order to configure a custom indicator feed.",
			},
			"username": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "username for authenticating with the URL.",
			},
			"fields_delimiter": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The delimiter that separates between the columns in the feed.",
			},
			"ignore_lines_that_start_with": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A prefix that will determine which lines to ignore.",
			},
			"use_gateway_proxy": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Use the gateway's proxy for retrieving the feed.",
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
		},
	}
}

func dataSourceManagementThreatIocFeedRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showThreatIocFeedRes, err := client.ApiCall("show-threat-ioc-feed", payload, client.GetSessionID(), true, client.IsProxyUsed())
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

	if v := threatIocFeed["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

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

			if len(customHeaderList) > 0 {

				var customHeaderListToReturn []map[string]interface{}

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
				_ = d.Set("custom_headers", customHeaderListToReturn)
			}
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
