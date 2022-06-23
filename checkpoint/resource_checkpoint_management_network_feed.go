package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"

	"strconv"
)

func resourceManagementNetworkFeed() *schema.Resource {
	return &schema.Resource{
		Create: createManagementNetworkFeed,
		Read:   readManagementNetworkFeed,
		Update: updateManagementNetworkFeed,
		Delete: deleteManagementNetworkFeed,
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
				Sensitive:   true,
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
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
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
				Default:     60,
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
				Description: "The delimiter that separates between the columns in the feed. For feed format 'Flat List' default is '\n'(new line).",
				Default:     "Depends on the feed format",
			},
			"ignore_lines_that_start_with": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A prefix that will determine which lines to ignore.",
				Default:     "#",
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
	}
}

func createManagementNetworkFeed(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	networkFeed := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		networkFeed["name"] = v.(string)
	}

	if v, ok := d.GetOk("feed_url"); ok {
		networkFeed["feed-url"] = v.(string)
	}

	if v, ok := d.GetOk("certificate_id"); ok {
		networkFeed["certificate-id"] = v.(string)
	}

	if v, ok := d.GetOk("feed_format"); ok {
		networkFeed["feed-format"] = v.(string)
	}

	if v, ok := d.GetOk("feed_type"); ok {
		networkFeed["feed-type"] = v.(string)
	}

	if v, ok := d.GetOk("password"); ok {
		networkFeed["password"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		networkFeed["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("username"); ok {
		networkFeed["username"] = v.(string)
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
			networkFeed["customHeader"] = customHeaderPayload
		}
	}

	if v, ok := d.GetOk("update_interval"); ok {
		networkFeed["update-interval"] = v.(int)
	}

	if v, ok := d.GetOk("data_column"); ok {
		networkFeed["data-column"] = v.(int)
	}

	if v, ok := d.GetOk("fields_delimiter"); ok {
		networkFeed["fields-delimiter"] = v.(string)
	}

	if v, ok := d.GetOk("ignore_lines_that_start_with"); ok {
		networkFeed["ignore-lines-that-start-with"] = v.(string)
	}

	if v, ok := d.GetOk("json_query"); ok {
		networkFeed["json-query"] = v.(string)
	}

	if v, ok := d.GetOkExists("use_gateway_proxy"); ok {
		networkFeed["use-gateway-proxy"] = v.(bool)
	}

	if v, ok := d.GetOk("color"); ok {
		networkFeed["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		networkFeed["comments"] = v.(string)
	}

	if v, ok := d.GetOk("domains_to_process"); ok {
		networkFeed["domains-to-process"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		networkFeed["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		networkFeed["ignore-errors"] = v.(bool)
	}

	log.Println("Create NetworkFeed - Map = ", networkFeed)

	addNetworkFeedRes, err := client.ApiCall("add-network-feed", networkFeed, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !addNetworkFeedRes.Success {
		if addNetworkFeedRes.ErrorMsg != "" {
			return fmt.Errorf(addNetworkFeedRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addNetworkFeedRes.GetData()["uid"].(string))

	return readManagementNetworkFeed(d, m)
}

func readManagementNetworkFeed(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showNetworkFeedRes, err := client.ApiCall("show-network-feed", payload, client.GetSessionID(), true, client.IsProxyUsed())
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

	if v := networkFeed["password"]; v != nil {
		_ = d.Set("password", v)
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

	if v := networkFeed["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := networkFeed["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementNetworkFeed(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	networkFeed := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		networkFeed["name"] = oldName
		networkFeed["new-name"] = newName
	} else {
		networkFeed["name"] = d.Get("name")
	}

	if ok := d.HasChange("feed_url"); ok {
		networkFeed["feed-url"] = d.Get("feed_url")
	}

	if ok := d.HasChange("certificate_id"); ok {
		networkFeed["certificate-id"] = d.Get("certificate_id")
	}

	if ok := d.HasChange("feed_format"); ok {
		networkFeed["feed-format"] = d.Get("feed_format")
	}

	if ok := d.HasChange("feed_type"); ok {
		networkFeed["feed-type"] = d.Get("feed_type")
	}

	if ok := d.HasChange("password"); ok {
		networkFeed["password"] = d.Get("password")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			networkFeed["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			networkFeed["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("username"); ok {
		networkFeed["username"] = d.Get("username")
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
			networkFeed["custom-header"] = customHeaderPayload
		} else {
			oldcustomHeader, _ := d.GetChange("custom_header")
			var customHeaderToDelete []interface{}
			for _, i := range oldcustomHeader.([]interface{}) {
				customHeaderToDelete = append(customHeaderToDelete, i.(map[string]interface{})["name"].(string))
			}
			networkFeed["custom-header"] = map[string]interface{}{"remove": customHeaderToDelete}
		}
	}

	if ok := d.HasChange("update_interval"); ok {
		networkFeed["update-interval"] = d.Get("update_interval")
	}

	if ok := d.HasChange("data_column"); ok {
		networkFeed["data-column"] = d.Get("data_column")
	}

	if ok := d.HasChange("fields_delimiter"); ok {
		networkFeed["fields-delimiter"] = d.Get("fields_delimiter")
	}

	if ok := d.HasChange("ignore_lines_that_start_with"); ok {
		networkFeed["ignore-lines-that-start-with"] = d.Get("ignore_lines_that_start_with")
	}

	if ok := d.HasChange("json_query"); ok {
		networkFeed["json-query"] = d.Get("json_query")
	}

	if v, ok := d.GetOkExists("use_gateway_proxy"); ok {
		networkFeed["use-gateway-proxy"] = v.(bool)
	}

	if ok := d.HasChange("color"); ok {
		networkFeed["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		networkFeed["comments"] = d.Get("comments")
	}

	if d.HasChange("domains_to_process") {
		if v, ok := d.GetOk("domains_to_process"); ok {
			networkFeed["domains_to_process"] = v.(*schema.Set).List()
		} else {
			oldDomains_To_Process, _ := d.GetChange("domains_to_process")
			networkFeed["domains_to_process"] = map[string]interface{}{"remove": oldDomains_To_Process.(*schema.Set).List()}
		}
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		networkFeed["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		networkFeed["ignore-errors"] = v.(bool)
	}

	log.Println("Update NetworkFeed - Map = ", networkFeed)

	updateNetworkFeedRes, err := client.ApiCall("set-network-feed", networkFeed, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !updateNetworkFeedRes.Success {
		if updateNetworkFeedRes.ErrorMsg != "" {
			return fmt.Errorf(updateNetworkFeedRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementNetworkFeed(d, m)
}

func deleteManagementNetworkFeed(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	networkFeedPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete NetworkFeed")

	deleteNetworkFeedRes, err := client.ApiCall("delete-network-feed", networkFeedPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteNetworkFeedRes.Success {
		if deleteNetworkFeedRes.ErrorMsg != "" {
			return fmt.Errorf(deleteNetworkFeedRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
