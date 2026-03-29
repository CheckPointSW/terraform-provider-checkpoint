package upgraders

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceManagementNetworkProbeV0 is the V0 schema where http_options and icmp_options were TypeMap.
func ResourceManagementNetworkProbeV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"http_options": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Additional options when [protocol] is set to \"http\".",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destination": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The destination URL.",
						},
					},
				},
			},
			"icmp_options": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Additional options when [protocol] is set to \"icmp\".",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destination": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "One of these:<br>- Name or UID of an existing object with a unicast IPv4 address (Host, Security Gateway, and so on).<br>- A unicast IPv4 address string (if you do not want to create such an object).",
						},
						"source": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "One of these:<br>- The string \"main-ip\" (the probe uses the main IPv4 address of the Security Gateway objects you specified in the parameter [install-on]).<br>- Name or UID of an existing object of type 'Host' with a unicast IPv4 address.<br>- A unicast IPv4 address string (if you do not want to create such an object).",
							Default:     "main-ip",
						},
					},
				},
			},
			"install_on": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "Collection of Check Point Security Gateways that generate the probe, identified by name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The probing protocol to use.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The time interval (in seconds) between each probe request.<br>Best Practice - The interval value should be lower than the timeout value.",
				Default:     10,
			},
			"timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The probe expiration timeout (in seconds). If there is not a single reply within this time, the status of the probe changes to \"Down\".",
				Default:     20,
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

// ResourceManagementNetworkProbeStateUpgradeV0 converts http_options and icmp_options from TypeMap to TypeList.
func ResourceManagementNetworkProbeStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	return UpgradeMapsToLists(rawState, "http_options", "icmp_options"), nil
}
