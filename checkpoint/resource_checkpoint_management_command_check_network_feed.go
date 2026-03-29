package checkpoint

import (
	"github.com/CheckPointSW/terraform-provider-checkpoint/upgraders"
	"fmt"
	"strconv"

	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceManagementCheckNetworkFeed() *schema.Resource {
	return &schema.Resource{
		Create: createManagementCheckNetworkFeed,
		Read:   readManagementCheckNetworkFeed,
		Delete: deleteManagementCheckNetworkFeed,
		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    upgraders.ResourceManagementCommandCheckNetworkFeedV0().CoreConfigSchema().ImpliedType(),
				Upgrade: upgraders.ResourceManagementCommandCheckNetworkFeedStateUpgradeV0,
				Version: 0,
			},
		},
		Schema: map[string]*schema.Schema{
			"network_feed": {
				Type:        schema.TypeList,
				MaxItems:    1,
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
	if v, ok := d.GetOk("network_feed"); ok {

		networkFeedList := v.([]interface{})

		if len(networkFeedList) > 0 {

			networkFeedPayload := make(map[string]interface{})

			if v, ok := d.GetOk("network_feed.0.name"); ok {
				networkFeedPayload["name"] = v.(string)
			}
			if v, ok := d.GetOk("network_feed.0.feed_url"); ok {
				networkFeedPayload["feed-url"] = v.(string)
			}
			if v, ok := d.GetOk("network_feed.0.certificate_id"); ok {
				networkFeedPayload["certificate-id"] = v.(string)
			}
			if v, ok := d.GetOk("network_feed.0.feed_format"); ok {
				networkFeedPayload["feed-format"] = v.(string)
			}
			if v, ok := d.GetOk("network_feed.0.feed_type"); ok {
				networkFeedPayload["feed-type"] = v.(string)
			}
			if v, ok := d.GetOk("network_feed.0.password"); ok {
				networkFeedPayload["password"] = v.(string)
			}
			if v, ok := d.GetOk("network_feed.0.username"); ok {
				networkFeedPayload["username"] = v.(string)
			}
			if v, ok := d.GetOk("network_feed.0.custom_header"); ok {

				customHeaderList := v.([]interface{})

				if len(customHeaderList) > 0 {

					var customHeaderPayload []map[string]interface{}

					for j := range customHeaderList {

						customHeaderMapToAdd := make(map[string]interface{})

						if v, ok := d.GetOk("network_feed.0.custom_header." + strconv.Itoa(j) + ".header_name"); ok {
							customHeaderMapToAdd["header-name"] = v.(string)
						}
						if v, ok := d.GetOk("network_feed.0.custom_header." + strconv.Itoa(j) + ".header_value"); ok {
							customHeaderMapToAdd["header-value"] = v.(string)
						}
						customHeaderPayload = append(customHeaderPayload, customHeaderMapToAdd)
					}
					networkFeedPayload["custom-header"] = customHeaderPayload
				}
			}
			if v, ok := d.GetOk("network_feed.0.update_interval"); ok {
				networkFeedPayload["update-interval"] = v.(int)
			}
			if v, ok := d.GetOk("network_feed.0.data_column"); ok {
				networkFeedPayload["data-column"] = v.(int)
			}
			if v, ok := d.GetOk("network_feed.0.fields_delimiter"); ok {
				networkFeedPayload["fields-delimiter"] = v.(string)
			}
			if v, ok := d.GetOk("network_feed.0.ignore_lines_that_start_with"); ok {
				networkFeedPayload["ignore-lines-that-start-with"] = v.(string)
			}
			if v, ok := d.GetOk("network_feed.0.json_query"); ok {
				networkFeedPayload["json-query"] = v.(string)
			}
			if v, ok := d.GetOkExists("network_feed.0.use_gateway_proxy"); ok {
				networkFeedPayload["use-gateway-proxy"] = v.(bool)
			}
			if v, ok := d.GetOk("network_feed.0.domains_to_process"); ok {
				networkFeedPayload["domains-to-process"] = v
			}
			if v, ok := d.GetOkExists("network_feed.0.ignore_warnings"); ok {
				networkFeedPayload["ignore-warnings"] = v.(bool)
			}
			if v, ok := d.GetOkExists("network_feed.0.ignore_errors"); ok {
				networkFeedPayload["ignore-errors"] = v.(bool)
			}
			payload["network-feed"] = networkFeedPayload
		}
	}

	if v, ok := d.GetOk("targets"); ok {
		payload["targets"] = v.(*schema.Set).List()
	}

	CheckNetworkFeedRes, _ := client.ApiCall("check-network-feed", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if !CheckNetworkFeedRes.Success {
		return fmt.Errorf("%s", CheckNetworkFeedRes.ErrorMsg)
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
