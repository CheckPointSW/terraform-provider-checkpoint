package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementNetworkFeed() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementNetworkFeedRead,
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
			"certificate_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Certificate SHA-1 fingerprint to access the feed.",
			},
			"feed_format": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Feed file format.",
			},
			"feed_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Feed type to be enforced.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"username": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "username for authenticating with the URL.",
			},
			"custom_header": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Headers to allow different authentication methods with the URL.",
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
			"update_interval": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Interval in minutes for updating the feed on the Security Gateway.",
			},
			"data_column": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Number of the column that contains the feed's data.",
			},
			"fields_delimiter": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The delimiter that separates between the columns in the feed. For feed format 'Flat List' default is '\n'(new line).",
			},
			"ignore_lines_that_start_with": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A prefix that will determine which lines to ignore.",
			},
			"json_query": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "JQ query to be parsed.",
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
			"domains_to_process": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Indicates which domains to process the commands on. It cannot be used with the details-level full, must be run from the System Domain only and with ignore-warnings true. Valid values are: CURRENT_DOMAIN, ALL_DOMAINS_ON_THIS_SERVER.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceManagementNetworkFeedRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showNetworkFeedRes, err := client.ApiCall("show-network-feed", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showNetworkFeedRes.Success {
		if objectNotFound(showNetworkFeedRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showNetworkFeedRes.ErrorMsg)
	}

	networkFeed := showNetworkFeedRes.GetData()

	log.Println("Read NetworkFeed - Show JSON = ", networkFeed)

	if v := networkFeed["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := networkFeed["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := networkFeed["feed-url"]; v != nil {
		_ = d.Set("feed_url", v)
	}

	if v := networkFeed["certificate-id"]; v != nil {
		_ = d.Set("certificate_id", v)
	}

	if v := networkFeed["feed-format"]; v != nil {
		_ = d.Set("feed_format", v)
	}

	if v := networkFeed["feed-type"]; v != nil {
		_ = d.Set("feed_type", v)
	}

	if networkFeed["tags"] != nil {
		tagsJson, ok := networkFeed["tags"].([]interface{})
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

	if v := networkFeed["username"]; v != nil {
		_ = d.Set("username", v)
	}

	if networkFeed["custom-header"] != nil {

		customHeaderList, ok := networkFeed["custom-header"].([]interface{})

		if ok {

			if len(customHeaderList) > 0 {

				var customHeaderListToReturn []map[string]interface{}

				for i := range customHeaderList {

					customHeaderMap := customHeaderList[i].(map[string]interface{})

					customHeaderMapToAdd := make(map[string]interface{})

					if v, _ := customHeaderMap["header-name"]; v != nil {
						customHeaderMapToAdd["header_name"] = v
					}
					if v, _ := customHeaderMap["header-value"]; v != nil {
						customHeaderMapToAdd["header_value"] = v
					}
					customHeaderListToReturn = append(customHeaderListToReturn, customHeaderMapToAdd)
				}
			}
		}
	}

	if v := networkFeed["update-interval"]; v != nil {
		_ = d.Set("update_interval", v)
	}

	if v := networkFeed["data-column"]; v != nil {
		_ = d.Set("data_column", v)
	}

	if v := networkFeed["fields-delimiter"]; v != nil {
		_ = d.Set("fields_delimiter", v)
	}

	if v := networkFeed["ignore-lines-that-start-with"]; v != nil {
		_ = d.Set("ignore_lines_that_start_with", v)
	}

	if v := networkFeed["json-query"]; v != nil {
		_ = d.Set("json_query", v)
	}

	if v := networkFeed["use-gateway-proxy"]; v != nil {
		_ = d.Set("use_gateway_proxy", v)
	}

	if v := networkFeed["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := networkFeed["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if networkFeed["domains_to_process"] != nil {
		domainsToProcessJson, ok := networkFeed["domains_to_process"].([]interface{})
		if ok {
			domainsToProcessIds := make([]string, 0)
			if len(domainsToProcessJson) > 0 {
				for _, domains_to_process := range domainsToProcessJson {
					domains_to_process := domains_to_process.(map[string]interface{})
					domainsToProcessIds = append(domainsToProcessIds, domains_to_process["name"].(string))
				}
			}
			_ = d.Set("domains_to_process", domainsToProcessIds)
		}
	} else {
		_ = d.Set("domains_to_process", nil)
	}

	return nil
}
