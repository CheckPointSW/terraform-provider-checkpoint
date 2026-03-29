package upgraders

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceManagementLsvProfileV0 is the V0 schema where vpn_domain was TypeMap.
func ResourceManagementLsvProfileV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name. Must be unique in the domain.",
			},
			"certificate_authority": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Trusted Certificate authority for establishing trust between VPN peers, identified by name or UID.",
			},
			"allowed_ip_addresses": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of network objects identified by name or UID that represent IP addresses allowed in profile's VPN domain.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"restrict_allowed_addresses": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Indicate whether the IP addresses allowed in the VPN Domain will be restricted or not, according to allowed-ip-addresses field.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"vpn_domain": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "peers' VPN Domain properties.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"limit_peer_domain_size": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Use this parameter to limit the number of IP addresses in the VPN Domain of each peer according to the value in the max-allowed-addresses field.",
						},
						"max_allowed_addresses": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     256,
							Description: "Maximum number of IP addresses in the VPN Domain of each peer. This value will be enforced only when limit-peer-domain-size field is set to true. Select a value between 1 and 256. Default value is 256.",
						},
					},
				},
			},
			"color": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "black",
				Description: "Color of the object. Should be one of existing colors.",
			},
			"comments": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comments string.",
			},
			"ignore_warnings": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Apply changes ignoring warnings.",
			},
			"ignore_errors": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.",
			},
		},
	}
}

// ResourceManagementLsvProfileStateUpgradeV0 converts vpn_domain from TypeMap to TypeList.
func ResourceManagementLsvProfileStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	return UpgradeMapsToLists(rawState, "vpn_domain"), nil
}
