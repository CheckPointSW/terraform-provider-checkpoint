package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementCheckThreatIocFeed() *schema.Resource {
	return &schema.Resource{
		Create: createManagementCheckThreatIocFeed,
		Read:   readManagementCheckThreatIocFeed,
		Delete: deleteManagementCheckThreatIocFeed,
		Schema: map[string]*schema.Schema{
			"ioc_feed": {
				Type:        schema.TypeMap,
				Required:    true,
				Description: "threat ioc feed parameters.",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
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
						},
						"feed_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Feed type to be enforced.",
						},
						"password": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "password for authenticating with the URL.",
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
				},
			},
			"targets": {
				Type:        schema.TypeSet,
				Required:    true,
				ForceNew:    true,
				Description: "On what targets to execute this command. Targets may be identified by their name, or object unique identifier.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Command asynchronous task unique identifier.",
			},
		},
	}
}

func createManagementCheckThreatIocFeed(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if _, ok := d.GetOk("ioc_feed"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("ioc_feed.name"); ok {
			res["name"] = v.(string)
		}
		if v, ok := d.GetOk("ioc_feed.feed_url"); ok {
			res["feed-url"] = v.(string)
		}
		if v, ok := d.GetOk("ioc_feed.action"); ok {
			res["action"] = v.(string)
		}
		if v, ok := d.GetOk("ioc_feed.certificate_id"); ok {
			res["certificate-id"] = v.(string)
		}
		if v, ok := d.GetOk("ioc_feed.custom_comment"); ok {
			res["custom-comment"] = v
		}
		if v, ok := d.GetOk("ioc_feed.custom_confidence"); ok {
			res["custom-confidence"] = v
		}
		if v, ok := d.GetOk("ioc_feed.custom_header"); ok {
			res["custom-header"] = v
		}
		if v, ok := d.GetOk("ioc_feed.custom_name"); ok {
			res["custom-name"] = v
		}
		if v, ok := d.GetOk("ioc_feed.custom_severity"); ok {
			res["custom-severity"] = v
		}
		if v, ok := d.GetOk("ioc_feed.custom_type"); ok {
			res["custom-type"] = v
		}
		if v, ok := d.GetOk("ioc_feed.custom_value"); ok {
			res["custom-value"] = v
		}
		if v, ok := d.GetOk("ioc_feed.enabled"); ok {
			res["enabled"] = v
		}
		if v, ok := d.GetOk("ioc_feed.feed_type"); ok {
			res["feed-type"] = v.(string)
		}
		if v, ok := d.GetOk("ioc_feed.password"); ok {
			res["password"] = v.(string)
		}
		if v, ok := d.GetOk("ioc_feed.use_custom_feed_settings"); ok {
			res["use-custom-feed-settings"] = v
		}
		if v, ok := d.GetOk("ioc_feed.username"); ok {
			res["username"] = v.(string)
		}
		if v, ok := d.GetOk("ioc_feed.fields_delimiter"); ok {
			res["fields-delimiter"] = v.(string)
		}
		if v, ok := d.GetOk("ioc_feed.ignore_lines_that_start_with"); ok {
			res["ignore-lines-that-start-with"] = v.(string)
		}
		if v, ok := d.GetOk("ioc_feed.use_gateway_proxy"); ok {
			res["use-gateway-proxy"] = v
		}
		if v, ok := d.GetOk("ioc_feed.ignore_warnings"); ok {
			res["ignore-warnings"] = v
		}
		if v, ok := d.GetOk("ioc_feed.ignore_errors"); ok {
			res["ignore-errors"] = v
		}
		payload["ioc-feed"] = res
	}

	if v, ok := d.GetOk("targets"); ok {
		payload["targets"] = v.(*schema.Set).List()
	}

	CheckThreatIocFeedRes, _ := client.ApiCall("check-threat-ioc-feed", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if !CheckThreatIocFeedRes.Success {
		return fmt.Errorf(CheckThreatIocFeedRes.ErrorMsg)
	}

	d.SetId("check-threat-ioc-feed" + acctest.RandString(10))
	_ = d.Set("task_id", resolveTaskId(CheckThreatIocFeedRes.GetData()))
	return readManagementCheckThreatIocFeed(d, m)
}

func readManagementCheckThreatIocFeed(d *schema.ResourceData, m interface{}) error {

	return nil
}

func deleteManagementCheckThreatIocFeed(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
