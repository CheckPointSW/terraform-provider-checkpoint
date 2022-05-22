package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementCheckNetworkFeed() *schema.Resource {
	return &schema.Resource{
		Create: createManagementCheckNetworkFeed,
		Read:   readManagementCheckNetworkFeed,
		Delete: deleteManagementCheckNetworkFeed,
		Schema: map[string]*schema.Schema{
			"network_feed": {
				Type:        schema.TypeMap,
				Required:    true,
				Description: "network feed parameters.",
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
						"certificate_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Certificate SHA-1 fingerprint to access the feed.",
						},
						"feed_format": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Feed file format.",
							Default:     "Flat List",
						},
						"feed_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Feed type to be enforced.",
							Default:     "IP Address",
						},
						"password": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "password for authenticating with the URL.",
						},
						"username": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "username for authenticating with the URL.",
						},
						"custom_header": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Headers to allow different authentication methods with the URL.",
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
						"update_interval": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Interval in minutes for updating the feed on the Security Gateway.",
						},
						"data_column": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Number of the column that contains the feed's data.",
							Default:     1,
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
						"json_query": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "JQ query to be parsed.",
						},
						"use_gateway_proxy": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Use the gateway's proxy for retrieving the feed.",
						},
						"domains_to_process": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Indicates which domains to process the commands on. It cannot be used with the details-level full, must be run from the System Domain only and with ignore-warnings true. Valid values are: CURRENT_DOMAIN, ALL_DOMAINS_ON_THIS_SERVER.",
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

func createManagementCheckNetworkFeed(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if _, ok := d.GetOk("network_feed"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("network_feed.name"); ok {
			res["name"] = v.(string)
		}
		if v, ok := d.GetOk("network_feed.feed_url"); ok {
			res["feed-url"] = v.(string)
		}
		if v, ok := d.GetOk("network_feed.certificate_id"); ok {
			res["certificate-id"] = v.(string)
		}
		if v, ok := d.GetOk("network_feed.feed_format"); ok {
			res["feed-format"] = v.(string)
		}
		if v, ok := d.GetOk("network_feed.feed_type"); ok {
			res["feed-type"] = v.(string)
		}
		if v, ok := d.GetOk("network_feed.password"); ok {
			res["password"] = v.(string)
		}
		if v, ok := d.GetOk("network_feed.username"); ok {
			res["username"] = v.(string)
		}
		if v, ok := d.GetOk("network_feed.custom_header"); ok {
			res["custom-header"] = v
		}
		if v, ok := d.GetOk("network_feed.update_interval"); ok {
			res["update-interval"] = v
		}
		if v, ok := d.GetOk("network_feed.data_column"); ok {
			res["data-column"] = v
		}
		if v, ok := d.GetOk("network_feed.fields_delimiter"); ok {
			res["fields-delimiter"] = v.(string)
		}
		if v, ok := d.GetOk("network_feed.ignore_lines_that_start_with"); ok {
			res["ignore-lines-that-start-with"] = v.(string)
		}
		if v, ok := d.GetOk("network_feed.json_query"); ok {
			res["json-query"] = v.(string)
		}
		if v, ok := d.GetOk("network_feed.use_gateway_proxy"); ok {
			res["use-gateway-proxy"] = v
		}
		if v, ok := d.GetOk("network_feed.domains_to_process"); ok {
			res["domains-to-process"] = v
		}
		if v, ok := d.GetOk("network_feed.ignore_warnings"); ok {
			res["ignore-warnings"] = v
		}
		if v, ok := d.GetOk("network_feed.ignore_errors"); ok {
			res["ignore-errors"] = v
		}
		payload["network-feed"] = res
	}

	if v, ok := d.GetOk("targets"); ok {
		payload["targets"] = v.(*schema.Set).List()
	}

	CheckNetworkFeedRes, _ := client.ApiCall("check-network-feed", payload, client.GetSessionID(), true, false)
	if !CheckNetworkFeedRes.Success {
		return fmt.Errorf(CheckNetworkFeedRes.ErrorMsg)
	}

	d.SetId("check-network-feed" + acctest.RandString(10))
	_ = d.Set("task_id", resolveTaskId(CheckNetworkFeedRes.GetData()))
	return readManagementCheckNetworkFeed(d, m)
}

func readManagementCheckNetworkFeed(d *schema.ResourceData, m interface{}) error {

	return nil
}

func deleteManagementCheckNetworkFeed(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
