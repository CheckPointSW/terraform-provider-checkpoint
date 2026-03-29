package upgraders

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceManagementAddressRangeV0 is the V0 schema where nat_settings was TypeMap.
func ResourceManagementAddressRangeV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name. Should be unique in the domain.",
			},
			"ipv4_address_first": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "First IPv4 address in the range.",
			},
			"ipv6_address_first": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "First IPv6 address in the range.",
			},
			"ipv4_address_last": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Last IPv4 address in the range.",
			},
			"ipv6_address_last": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Last IPv6 address in the range.",
			},
			"nat_settings": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "NAT settings.",
				//Default: map[string]interface{}{"auto_rule":false},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_rule": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Whether to add automatic address translation rules.",
						},
						"ipv4_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IPv4 address.",
						},
						"ipv6_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IPv6 address.",
						},
						"hide_behind": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Hide behind method. This parameter is not required in case \"method\" parameter is \"static\".",
						},
						"install_on": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Which gateway should apply the NAT translation.",
						},
						"method": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "NAT translation method.",
						},
					},
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
			"color": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Color of the object. Should be one of existing colors.",
				Default:     "black",
			},
			"comments": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comments string.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

// ResourceManagementAddressRangeStateUpgradeV0 converts nat_settings from TypeMap to TypeList.
func ResourceManagementAddressRangeStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	return UpgradeMapsToLists(rawState, "nat_settings"), nil
}
