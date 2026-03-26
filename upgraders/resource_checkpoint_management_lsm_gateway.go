package upgraders

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceManagementLsmGatewayV0 is the V0 schema where provisioning_settings and sic were TypeMap.
func ResourceManagementLsmGatewayV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"security_profile": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "LSM profile.",
			},
			"device_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Device ID.",
			},
			"dynamic_objects": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Dynamic Objects.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"comments": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Comments.",
						},
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Object name. Must be unique in the domain.",
						},
						"uid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "UID.",
						},
						"resolved_ip_addresses": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Single IP-address or a range of addresses.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv4_address": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "IPv4 Address.",
									},
									"ipv4_address_range": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "IPv4 Address range.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"from_ipv4_address": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "First IPv4 address of the IP address range.",
												},
												"to_ipv4_address": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Last IPv4 address of the IP address range.",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"ip_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IP address",
			},
			"os_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Device platform operating system.",
			},
			"provisioning_settings": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Provisioning settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"provisioning_profile": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Provisioning profile.",
							Default:     "No Provisioning Profile",
						},
					},
				},
			},
			"provisioning_state": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Provisioning state. By default the state is 'manual'- enable provisioning but not attach to profile. If 'using-profile' state is provided a provisioning profile must be provided in provisioning-settings.",
			},

			"sic": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Secure Internal Communication.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"one_time_password": {
							Type:        schema.TypeString,
							Optional:    true,
							Sensitive:   true,
							Description: "One-time password. When one-time password is provided without ip-address- trusted communication is automatically initiated  when the gateway connects to the Security Management server for the first time.",
						},
						"ip_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IP address. When IP address is provided- initiate trusted communication immediately using this IP address",
						},
					},
				},
			},
			"sic_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Secure Internal Communication name.",
			},
			"sic_state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Secure Internal Communication state.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"topology": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Topology.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"manual_vpn_domain": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "A list of IP-addresses ranges, defined the VPN community network.This field is relevant only when 'manual' option of vpn-domain is checked.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"comments": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Comments string.",
									},
									"from_ipv4_address": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "First IPv4 address of the IP address range.",
									},
									"to_ipv4_address": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Last IPv4 address of the IP address range.",
									},
								},
							},
						},
						"vpn_domain": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "VPN Domain type.  'external-interfaces-only' is relevnt only for Gaia devices. 'hide-behind-gateway-external-ip-address' is relevant only for SMB devices.",
						},
					},
				},
			},
			"version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Device platform version.",
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

// ResourceManagementLsmGatewayStateUpgradeV0 converts provisioning_settings and sic from TypeMap to TypeList.
func ResourceManagementLsmGatewayStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	return UpgradeMapsToLists(rawState, "provisioning_settings", "sic"), nil
}
