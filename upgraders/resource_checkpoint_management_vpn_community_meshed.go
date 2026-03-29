package upgraders

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceManagementVpnCommunityMeshedV0 is the V0 schema where ike_phase_1, ike_phase_2,
// and granular_encryptions[*].ike_phase_1/2 were TypeMap.
func ResourceManagementVpnCommunityMeshedV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"disable_nat": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates whether to disable NAT inside the VPN Community.",
			},
			"encrypted_traffic": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Encrypted traffic settings.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Indicates whether to accept all encrypted traffic.",
						},
					},
				},
			},
			"encryption_method": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The encryption method to be used.",
			},
			"encryption_suite": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The encryption suite to be used.",
			},
			"excluded_services": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of services that are excluded from the community identified by the name or UID.<br> Connections with these services will not be encrypted and will not match rules specifying the community in the VPN community.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"gateways": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of Gateway objects identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ike_phase_1": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Ike Phase 1 settings. Only applicable when the encryption-suite is set to [custom].",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"data_integrity": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The hash algorithm to be used.",
							Default:     "sha1",
						},
						"diffie_hellman_group": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The Diffie-Hellman group to be used.",
							Default:     "group-2",
						},
						"encryption_algorithm": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The encryption algorithm to be used.",
							Default:     "aes-256",
						},
						"ike_p1_rekey_time": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Indicates the time interval for IKE phase 1 renegotiation.",
							Default:     1440,
						},
					},
				},
			},
			"ike_phase_2": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Ike Phase 2 settings. Only applicable when the encryption-suite is set to [custom].",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"data_integrity": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The hash algorithm to be used.",
							Default:     "sha1",
						},
						"encryption_algorithm": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The encryption algorithm to be used.",
							Default:     "aes-128",
						},
						"ike_p2_use_pfs": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Indicates whether Perfect Forward Secrecy (PFS) is being used for IKE phase 2.",
							Default:     false,
						},
						"ike_p2_pfs_dh_grp": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The Diffie-Hellman group to be used.",
							Default:     "group-2",
						},
						"ike_p2_rekey_time": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Indicates the time interval for IKE phase 2 renegotiation.",
							Default:     1440,
						},
					},
				},
			},
			"link_selection_mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Link Selection Mode.",
			},
			"override_interfaces": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Override the Enhanced Link Selection interfaces for each participant VPN peer.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gateway": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Participant VPN Peer.",
						},
						"interfaces": {
							Type:        schema.TypeList,
							Required:    true,
							Description: "Enhanced Link Selection Interfaces.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"interface_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The name of the interface.",
									},
									"next_hop_ip": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The IP address of the next hop.",
									},
									"static_nat_ip": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The NATed IPv4 address that hides the source IPv4 address of outgoing connections (applies only to IPv4).",
									},
									"priority": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Priority of a 'Backup' interface.",
									},
									"redundancy_mode": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Interface redundancy mode (Active/Backup).",
									},
								},
							},
						},
					},
				},
			},
			"override_vpn_domains": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The Overrides VPN Domains of the participants GWs.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gateway": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Participant gateway in override VPN domain identified by the name or UID.",
						},
						"vpn_domain": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "VPN domain network identified by the name or UID.",
						},
					},
				},
			},
			"shared_secrets": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Shared secrets for external gateways.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"external_gateway": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "External gateway identified by the name or UID.",
						},
						"shared_secret": {
							Type:        schema.TypeString,
							Optional:    true,
							Sensitive:   true,
							Description: "Shared secret.",
						},
					},
				},
			},
			"tunnel_granularity": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "VPN tunnel sharing option to be used.",
			},
			"granular_encryptions": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "VPN granular encryption settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"internal_gateway": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Internally managed Check Point gateway identified by name or UID, or 'Any' for all internal-gateways participants in this community.",
						},
						"external_gateway": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Externally managed or 3rd party gateway identified by name or UID.",
						},
						"encryption_method": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The encryption method to be used.",
						},
						"encryption_suite": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The encryption suite to be used.",
						},
						"ike_phase_1": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Ike Phase 1 settings. Only applicable when the encryption-suite is set to [custom].",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"data_integrity": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The hash algorithm to be used.",
										Default:     "sha1",
									},
									"diffie_hellman_group": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The Diffie-Hellman group to be used.",
										Default:     "group-2",
									},
									"encryption_algorithm": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The encryption algorithm to be used.",
										Default:     "aes-256",
									},
									"ike_p1_rekey_time": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Indicates the time interval for IKE phase 1 renegotiation.",
										Default:     1440,
									},
								},
							},
						},
						"ike_phase_2": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Ike Phase 2 settings. Only applicable when the encryption-suite is set to [custom].",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"data_integrity": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The hash algorithm to be used.",
										Default:     "sha1",
									},
									"encryption_algorithm": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The encryption algorithm to be used.",
										Default:     "aes-128",
									},
									"ike_p2_use_pfs": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Indicates whether Perfect Forward Secrecy (PFS) is being used for IKE phase 2.",
										Default:     false,
									},
									"ike_p2_pfs_dh_grp": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The Diffie-Hellman group to be used.",
										Default:     "group-2",
									},
									"ike_p2_rekey_time": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Indicates the time interval for IKE phase 2 renegotiation.",
										Default:     1440,
									},
								},
							},
						},
					},
				},
			},
			"permanent_tunnels": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Permanent tunnels properties.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"set_permanent_tunnels": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Indicates which tunnels to set as permanent.",
						},
						"gateways": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of gateways to set all their tunnels to permanent with specified track options. Will take effect only if set-permanent-tunnels-on is set to all-tunnels-of-specific-gateways.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"gateway": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Gateway to set all is tunnels to permanent with specified track options.<br> Identified by name or UID.",
									},
									"track_options": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Indicates whether to use the community track options or to override track options for the permanent tunnels.",
									},
									"override_tunnel_down_track": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Gateway tunnel down track option. Relevant only if the track-options is set to 'override track options'.",
									},
									"override_tunnel_up_track": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Gateway tunnel up track option. Relevant only if the track-options is set to 'override track options'.",
									},
								},
							},
						},
						"tunnels": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of tunnels to set as permanent with specified track options. Will take effect only if set-permanent-tunnels-on is set to specific-tunnels-in-the-community.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"first_tunnel_endpoint": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "First tunnel endpoint (center gateway). Identified by name or UID.",
									},
									"second_tunnel_endpoint": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Second tunnel endpoint (center gateway for meshed VPN community and satellitegateway for star VPN community).  Identified by name or UID.",
									},
									"track_options": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Indicates whether to use the community track options or to override track options for the permanent tunnels.",
									},
									"override_tunnel_down_track": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Gateway tunnel down track option. Relevant only if the track-options is set to 'override track options'.",
									},
									"override_tunnel_up_track": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Gateway tunnel up track option. Relevant only if the track-options is set to 'override track options'.",
									},
								},
							},
						},
						"rim": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Route Injection Mechanism settings.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Indicates whether Route Injection Mechanism is enabled.",
									},
									"enable_on_gateways": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Indicates whether to enable automatic Route Injection Mechanism for gateways.",
									},
									"route_injection_track": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Route injection track method.",
									},
								},
							},
						},
						"tunnel_down_track": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "VPN community permanent tunnels down track option.",
						},
						"tunnel_up_track": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Permanent tunnels up track option.",
						},
					},
				},
			},
			"wire_mode": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "VPN Community Wire mode properties.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allow_uninspected_encrypted_traffic": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Allow uninspected encrypted traffic between Wire mode interfaces of this Community members.",
						},
						"allow_uninspected_encrypted_routing": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Allow members to route uninspected encrypted traffic in VPN routing configurations.",
						},
					},
				},
			},
			"routing_mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "VPN Community Routing Mode.",
			},
			"advanced_properties": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Advanced properties.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"support_ip_compression": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Indicates whether to support IP compression.",
						},
						"use_aggressive_mode": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Indicates whether to use aggressive mode.",
						},
					},
				},
			},
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"use_shared_secret": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates whether the shared secret should be used for all external gateways.",
				Default:     false,
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

// ResourceManagementVpnCommunityMeshedStateUpgradeV0 converts ike_phase_1/2 from TypeMap to TypeList.
func ResourceManagementVpnCommunityMeshedStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	return UpgradeMapsToLists(rawState, "ike_phase_1", "ike_phase_2", "granular_encryptions"), nil
}
