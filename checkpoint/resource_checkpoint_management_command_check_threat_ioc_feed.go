package checkpoint

import (
	"github.com/CheckPointSW/terraform-provider-checkpoint/upgraders"
	"fmt"
	"strconv"

	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceManagementCheckThreatIocFeed() *schema.Resource {
	return &schema.Resource{
		Create: createManagementCheckThreatIocFeed,
		Read:   readManagementCheckThreatIocFeed,
		Delete: deleteManagementCheckThreatIocFeed,
		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    upgraders.ResourceManagementCommandCheckThreatIocFeedV0().CoreConfigSchema().ImpliedType(),
				Upgrade: upgraders.ResourceManagementCommandCheckThreatIocFeedStateUpgradeV0,
				Version: 0,
			},
		},
		Schema: map[string]*schema.Schema{
			"ioc_feed": {
				Type:        schema.TypeList,
				MaxItems:    1,
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
	if v, ok := d.GetOk("ioc_feed"); ok {

		iocFeedList := v.([]interface{})

		if len(iocFeedList) > 0 {

			iocFeedPayload := make(map[string]interface{})

			if v, ok := d.GetOk("ioc_feed.0.name"); ok {
				iocFeedPayload["name"] = v.(string)
			}
			if v, ok := d.GetOk("ioc_feed.0.feed_url"); ok {
				iocFeedPayload["feed-url"] = v.(string)
			}
			if v, ok := d.GetOk("ioc_feed.0.action"); ok {
				iocFeedPayload["action"] = v.(string)
			}
			if v, ok := d.GetOk("ioc_feed.0.certificate_id"); ok {
				iocFeedPayload["certificate-id"] = v.(string)
			}
			if v, ok := d.GetOk("ioc_feed.0.confidence"); ok {
				iocFeedPayload["confidence"] = v.(int)
			}
			if v, ok := d.GetOk("ioc_feed.0.custom_comment"); ok {
				iocFeedPayload["custom-comment"] = v.(int)
			}
			if v, ok := d.GetOk("ioc_feed.0.custom_confidence"); ok {
				iocFeedPayload["custom-confidence"] = v.(int)
			}
			if v, ok := d.GetOk("ioc_feed.0.custom_header"); ok {

				customHeaderList := v.([]interface{})

				if len(customHeaderList) > 0 {

					var customHeaderPayload []map[string]interface{}

					for j := range customHeaderList {

						customHeaderMapToAdd := make(map[string]interface{})

						if v, ok := d.GetOk("ioc_feed.0.custom_header." + strconv.Itoa(j) + ".header_name"); ok {
							customHeaderMapToAdd["header-name"] = v.(string)
						}
						if v, ok := d.GetOk("ioc_feed.0.custom_header." + strconv.Itoa(j) + ".header_value"); ok {
							customHeaderMapToAdd["header-value"] = v.(string)
						}
						customHeaderPayload = append(customHeaderPayload, customHeaderMapToAdd)
					}
					iocFeedPayload["custom-header"] = customHeaderPayload
				}
			}
			if v, ok := d.GetOk("ioc_feed.0.custom_name"); ok {
				iocFeedPayload["custom-name"] = v.(int)
			}
			if v, ok := d.GetOk("ioc_feed.0.custom_severity"); ok {
				iocFeedPayload["custom-severity"] = v.(int)
			}
			if v, ok := d.GetOk("ioc_feed.0.custom_type"); ok {
				iocFeedPayload["custom-type"] = v.(int)
			}
			if v, ok := d.GetOk("ioc_feed.0.custom_value"); ok {
				iocFeedPayload["custom-value"] = v.(int)
			}
			if v, ok := d.GetOkExists("ioc_feed.0.enabled"); ok {
				iocFeedPayload["enabled"] = v.(bool)
			}
			if v, ok := d.GetOk("ioc_feed.0.feed_type"); ok {
				iocFeedPayload["feed-type"] = v.(string)
			}
			if v, ok := d.GetOk("ioc_feed.0.password"); ok {
				iocFeedPayload["password"] = v.(string)
			}
			if v, ok := d.GetOk("ioc_feed.0.performance_impact"); ok {
				iocFeedPayload["performance-impact"] = v.(int)
			}
			if v, ok := d.GetOk("ioc_feed.0.severity"); ok {
				iocFeedPayload["severity"] = v.(int)
			}
			if v, ok := d.GetOkExists("ioc_feed.0.use_custom_feed_settings"); ok {
				iocFeedPayload["use-custom-feed-settings"] = v.(bool)
			}
			if v, ok := d.GetOk("ioc_feed.0.use_snort_format"); ok {
				iocFeedPayload["use-snort-format"] = v.(bool)
			}
			if v, ok := d.GetOk("ioc_feed.0.username"); ok {
				iocFeedPayload["username"] = v.(string)
			}
			if v, ok := d.GetOk("ioc_feed.0.fields_delimiter"); ok {
				iocFeedPayload["fields-delimiter"] = v.(string)
			}
			if v, ok := d.GetOk("ioc_feed.0.ignore_lines_that_start_with"); ok {
				iocFeedPayload["ignore-lines-that-start-with"] = v.(string)
			}
			if v, ok := d.GetOkExists("ioc_feed.0.use_gateway_proxy"); ok {
				iocFeedPayload["use-gateway-proxy"] = v.(bool)
			}
			if v, ok := d.GetOkExists("ioc_feed.0.ignore_warnings"); ok {
				iocFeedPayload["ignore-warnings"] = v.(bool)
			}
			if v, ok := d.GetOkExists("ioc_feed.0.ignore_errors"); ok {
				iocFeedPayload["ignore-errors"] = v.(bool)
			}
			payload["ioc-feed"] = iocFeedPayload
		}
	}

	if v, ok := d.GetOk("targets"); ok {
		payload["targets"] = v.(*schema.Set).List()
	}

	CheckThreatIocFeedRes, _ := client.ApiCall("check-threat-ioc-feed", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if !CheckThreatIocFeedRes.Success {
		return fmt.Errorf("%s", CheckThreatIocFeedRes.ErrorMsg)
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
