package upgraders

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceManagementCommandCheckThreatIocFeedV0 is the V0 schema where ioc_feed was TypeMap.
func ResourceManagementCommandCheckThreatIocFeedV0() *schema.Resource {
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

// ResourceManagementCommandCheckThreatIocFeedStateUpgradeV0 converts ioc_feed from TypeMap to TypeList.
func ResourceManagementCommandCheckThreatIocFeedStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	return UpgradeMapsToLists(rawState, "ioc_feed"), nil
}
