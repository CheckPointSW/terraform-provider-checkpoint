package upgraders

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceManagementCommandCheckNetworkFeedV0 is the V0 schema where network_feed was TypeMap.
func ResourceManagementCommandCheckNetworkFeedV0() *schema.Resource {
	return &schema.Resource{
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

// ResourceManagementCommandCheckNetworkFeedStateUpgradeV0 converts network_feed from TypeMap to TypeList.
func ResourceManagementCommandCheckNetworkFeedStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	return UpgradeMapsToLists(rawState, "network_feed"), nil
}
