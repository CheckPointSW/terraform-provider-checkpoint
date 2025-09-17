package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"reflect"
	"strconv"
)

func dataSourceManagementVpnCommunityStar() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementVpnCommunityStarRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object unique identifier.",
			},
			"center_gateways": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of Gateway objects representing center gateways identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"disable_nat": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether to disable NAT inside the VPN Community.",
			},
			"disable_nat_on": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Indicates on which gateways to disable NAT inside the VPN Community.",
			},
			"encrypted_traffic": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Encrypted traffic settings.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether to accept all encrypted traffic.",
						},
						"community_members": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Indicates on which community members to accept all encrypted traffic.",
						},
					},
				},
			},
			"encryption_method": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The encryption method to be used.",
			},
			"encryption_suite": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The encryption suite to be used.",
			},
			"excluded_services": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of services that are excluded from the community identified by the name or UID.<br> Connections with these services will not be encrypted and will not match rules specifying the community in the VPN community.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ike_phase_1": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Ike Phase 1 settings. Only applicable when the encryption-suite is set to [custom].",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"data_integrity": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The hash algorithm to be used.",
						},
						"diffie_hellman_group": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Diffie-Hellman group to be used.",
						},
						"encryption_algorithm": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The encryption algorithm to be used.",
						},
						"ike_p1_rekey_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Indicates the time interval for IKE phase 1 renegotiation.",
						},
					},
				},
			},
			"ike_phase_2": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Ike Phase 2 settings. Only applicable when the encryption-suite is set to [custom].",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"data_integrity": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The hash algorithm to be used.",
						},
						"encryption_algorithm": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The encryption algorithm to be used.",
						},
						"ike_p2_use_pfs": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Indicates whether Perfect Forward Secrecy (PFS) is being used for IKE phase 2.",
						},
						"ike_p2_pfs_dh_grp": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Diffie-Hellman group to be used.",
						},
						"ike_p2_rekey_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Indicates the time interval for IKE phase 2 renegotiation.",
						},
					},
				},
			},
			"link_selection_mode": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Link Selection Mode.",
			},
			"override_interfaces": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Override the Enhanced Link Selection interfaces for each participant VPN peer.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gateway": {
							Type:        schema.TypeString,
							Computed:    true,
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
										Computed:    true,
										Description: "The name of the interface.",
									},
									"next_hop_ip": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The IP address of the next hop.",
									},
									"static_nat_ip": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The NATed IPv4 address that hides the source IPv4 address of outgoing connections (applies only to IPv4).",
									},
									"priority": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Priority of a 'Backup' interface.",
									},
									"redundancy_mode": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Interface redundancy mode (Active/Backup).",
									},
								},
							},
						},
					},
				},
			},
			"mep": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Multiple Entry Point properties.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable center gateways as Multiple Entry Points.",
						},
						"entry_point_selection_mechanism": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The method by which the entry point gateway will be chosen from the gateways in the center.",
						},
						"entry_point_final_selection_mechanism": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The method by which the final entry point gateway will be chosen when the chosen mechanism returns more than one optional entry point.",
						},
						"tracking": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Tracking option for the MEP.",
						},
						"default_priority_rule": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Priority rule for all satellite gateways. Relevant only if 'entry-point-selection-mechanism' is set to 'manual'.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"first_priority_center_gateways": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "Collection of first priority center gateways identified by the name or UID.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"second_priority_center_gateways": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "Collection of second priority center gateways identified by the name or UID.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"third_priority_center_gateways": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "Collection of third priority center gateways identified by the name or UID.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"exception_priority_rules": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Exception priority rules for specific satellites gateways. Relevant only if 'entry-point-selection-mechanism' is set to 'manual'.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"satellite_gateways": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "Collection of satellite gateways to apply priority rules on identified by the name or UID.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"first_priority_center_gateways": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "Collection of first priority center gateways identified by the name or UID.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"second_priority_center_gateways": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "Collection of second priority center gateways identified by the name or UID.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"third_priority_center_gateways": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "Collection of third priority center gateways identified by the name or UID.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
					},
				},
			},
			"mesh_center_gateways": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether the meshed community is in center.",
			},
			"override_vpn_domains": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The Overrides VPN Domains of the participants GWs.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gateway": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Participant gateway in override VPN domain identified by the name or UID.",
						},
						"vpn_domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "VPN domain network identified by the name or UID.",
						},
					},
				},
			},
			"satellite_gateways": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of Gateway objects representing satellite gateways identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"shared_secrets": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Shared secrets for external gateways.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"external_gateway": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "External gateway identified by the name or UID.",
						},
						"shared_secret": {
							Type:        schema.TypeString,
							Computed:    true,
							Sensitive:   true,
							Description: "Shared secret.",
						},
					},
				},
			},
			"tunnel_granularity": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "VPN tunnel sharing option to be used.",
			},
			"granular_encryptions": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "VPN granular encryption settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"internal_gateway": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Internally managed Check Point gateway identified by name or UID, or 'Any' for all internal-gateways participants in this community.",
						},
						"external_gateway": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Externally managed or 3rd party gateway identified by name or UID.",
						},
						"encryption_method": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The encryption method to be used.",
						},
						"encryption_suite": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The encryption suite to be used.",
						},
						"ike_phase_1": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Ike Phase 1 settings. Only applicable when the encryption-suite is set to [custom].",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"data_integrity": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The hash algorithm to be used.",
									},
									"diffie_hellman_group": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Diffie-Hellman group to be used.",
									},
									"encryption_algorithm": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The encryption algorithm to be used.",
									},
									"ike_p1_rekey_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Indicates the time interval for IKE phase 1 renegotiation.",
									},
								},
							},
						},
						"ike_phase_2": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Ike Phase 2 settings. Only applicable when the encryption-suite is set to [custom].",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"data_integrity": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The hash algorithm to be used.",
									},
									"encryption_algorithm": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The encryption algorithm to be used.",
									},
									"ike_p2_use_pfs": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Indicates whether Perfect Forward Secrecy (PFS) is being used for IKE phase 2.",
									},
									"ike_p2_pfs_dh_grp": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Diffie-Hellman group to be used.",
									},
									"ike_p2_rekey_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Indicates the time interval for IKE phase 2 renegotiation.",
									},
								},
							},
						},
					},
				},
			},
			"permanent_tunnels": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Permanent tunnels properties.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"set_permanent_tunnels": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Indicates which tunnels to set as permanent.",
						},
						"gateways": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of gateways to set all their tunnels to permanent with specified track options. Will take effect only if set-permanent-tunnels-on is set to all-tunnels-of-specific-gateways.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"gateway": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Gateway to set all is tunnels to permanent with specified track options.<br> Identified by name or UID.",
									},
									"track_options": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Indicates whether to use the community track options or to override track options for the permanent tunnels.",
									},
									"override_tunnel_down_track": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Gateway tunnel down track option. Relevant only if the track-options is set to 'override track options'.",
									},
									"override_tunnel_up_track": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Gateway tunnel up track option. Relevant only if the track-options is set to 'override track options'.",
									},
								},
							},
						},
						"tunnels": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of tunnels to set as permanent with specified track options. Will take effect only if set-permanent-tunnels-on is set to specific-tunnels-in-the-community.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"first_tunnel_endpoint": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "First tunnel endpoint (center gateway). Identified by name or UID.",
									},
									"second_tunnel_endpoint": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Second tunnel endpoint (center gateway for meshed VPN community and satellitegateway for star VPN community).  Identified by name or UID.",
									},
									"track_options": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Indicates whether to use the community track options or to override track options for the permanent tunnels.",
									},
									"override_tunnel_down_track": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Gateway tunnel down track option. Relevant only if the track-options is set to 'override track options'.",
									},
									"override_tunnel_up_track": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Gateway tunnel up track option. Relevant only if the track-options is set to 'override track options'.",
									},
								},
							},
						},
						"rim": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Route Injection Mechanism settings.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Indicates whether Route Injection Mechanism is enabled.",
									},
									"enable_on_center_gateways": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Indicates whether to enable automatic Route Injection Mechanism on center gateways.",
									},
									"enable_on_satellite_gateways": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Indicates whether to enable automatic Route Injection Mechanism on satellite gateways.",
									},
									"route_injection_track": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Route injection track method.",
									},
								},
							},
						},
						"tunnel_down_track": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "VPN community permanent tunnels down track option.",
						},
						"tunnel_up_track": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Permanent tunnels up track option.",
						},
					},
				},
			},
			"vpn_routing": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Enable VPN routing to satellites.",
			},
			"wire_mode": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "VPN Community Wire mode properties.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allow_uninspected_encrypted_traffic": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Allow uninspected encrypted traffic between Wire mode interfaces of this Community members.",
						},
						"allow_uninspected_encrypted_routing": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Allow members to route uninspected encrypted traffic in VPN routing configurations.",
						},
					},
				},
			},
			"routing_mode": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "VPN Community Routing Mode.",
			},
			"advanced_properties": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Advanced properties.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"support_ip_compression": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether to support IP compression.",
						},
						"use_aggressive_mode": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether to use aggressive mode.",
						},
					},
				},
			},
			"tags": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"use_shared_secret": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether the shared secret should be used for all external gateways.",
			},
			"color": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Color of the object. Should be one of existing colors.",
			},
			"comments": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Comments string.",
			},
		},
	}
}

func dataSourceManagementVpnCommunityStarRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showVpnCommunityStarRes, err := client.ApiCall("show-vpn-community-star", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showVpnCommunityStarRes.Success {
		return fmt.Errorf(showVpnCommunityStarRes.ErrorMsg)
	}

	vpnCommunityStar := showVpnCommunityStarRes.GetData()

	log.Println("Read VpnCommunityStar - Show JSON = ", vpnCommunityStar)

	if v := vpnCommunityStar["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := vpnCommunityStar["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if vpnCommunityStar["center-gateways"] != nil {
		centerGatewaysJson, ok := vpnCommunityStar["center-gateways"].([]interface{})
		if ok {
			centerGatewaysIds := make([]string, 0)
			if len(centerGatewaysJson) > 0 {
				for _, center_gateways := range centerGatewaysJson {
					center_gateways := center_gateways.(map[string]interface{})
					centerGatewaysIds = append(centerGatewaysIds, center_gateways["name"].(string))
				}
			}
			_ = d.Set("center_gateways", centerGatewaysIds)
		}
	} else {
		_ = d.Set("center_gateways", nil)
	}

	if v := vpnCommunityStar["disable-nat"]; v != nil {
		_ = d.Set("disable_nat", v)
	}

	if v := vpnCommunityStar["disable-nat-on"]; v != nil {
		_ = d.Set("disable_nat_on", v)
	}

	if vpnCommunityStar["encrypted-traffic"] != nil {

		encryptedTrafficMap := vpnCommunityStar["encrypted-traffic"].(map[string]interface{})

		encryptedTrafficMapToReturn := make(map[string]interface{})

		if v := encryptedTrafficMap["enabled"]; v != nil {
			encryptedTrafficMapToReturn["enabled"] = v
		}
		if v := encryptedTrafficMap["community-members"]; v != nil {
			encryptedTrafficMapToReturn["community_members"] = v
		}
		_ = d.Set("encrypted_traffic", []interface{}{encryptedTrafficMapToReturn})

	} else {
		_ = d.Set("encrypted_traffic", nil)
	}

	if v := vpnCommunityStar["encryption-method"]; v != nil {
		_ = d.Set("encryption_method", v)
	}

	if v := vpnCommunityStar["encryption-suite"]; v != nil {
		_ = d.Set("encryption_suite", v)
	}

	if vpnCommunityStar["excluded_services"] != nil {
		excludedServicesJson, ok := vpnCommunityStar["excluded_services"].([]interface{})
		if ok {
			excludedServicesIds := make([]string, 0)
			if len(excludedServicesJson) > 0 {
				for _, excluded_services := range excludedServicesJson {
					excluded_services := excluded_services.(map[string]interface{})
					excludedServicesIds = append(excludedServicesIds, excluded_services["name"].(string))
				}
			}
			_ = d.Set("excluded_services", excludedServicesIds)
		}
	} else {
		_ = d.Set("excluded_services", nil)
	}

	if vpnCommunityStar["ike-phase-1"] != nil {

		ikePhase1Map := vpnCommunityStar["ike-phase-1"].(map[string]interface{})

		ikePhase1MapToReturn := make(map[string]interface{})

		if v, _ := ikePhase1Map["data-integrity"]; v != nil {
			ikePhase1MapToReturn["data_integrity"] = v
		}
		if v, _ := ikePhase1Map["diffie-hellman-group"]; v != nil {
			ikePhase1MapToReturn["diffie_hellman_group"] = v
		}
		if v, _ := ikePhase1Map["encryption-algorithm"]; v != nil {
			ikePhase1MapToReturn["encryption_algorithm"] = v
		}
		if v := ikePhase1Map["ike-p1-rekey-time"]; v != nil {
			ikePhase1MapToReturn["ike_p1_rekey_time"] = strconv.Itoa(int(v.(float64)))
		}
		_, ikePhase1InConf := d.GetOk("ike_phase_1")
		defaultIkePhase1 := map[string]interface{}{"encryption_algorithm": "aes-256", "diffie_hellman_group": "group-2", "data_integrity": "sha1"}
		if reflect.DeepEqual(defaultIkePhase1, ikePhase1MapToReturn) && !ikePhase1InConf {
			_ = d.Set("ike_phase_1", map[string]interface{}{})
		} else {
			_ = d.Set("ike_phase_1", ikePhase1MapToReturn)
		}

	} else {
		_ = d.Set("ike_phase_1", nil)
	}

	if vpnCommunityStar["ike-phase-2"] != nil {

		ikePhase2Map := vpnCommunityStar["ike-phase-2"].(map[string]interface{})

		ikePhase2MapToReturn := make(map[string]interface{})

		if v, _ := ikePhase2Map["data-integrity"]; v != nil {
			ikePhase2MapToReturn["data_integrity"] = v
		}
		if v, _ := ikePhase2Map["encryption-algorithm"]; v != nil {
			ikePhase2MapToReturn["encryption_algorithm"] = v
		}
		if v := ikePhase2Map["ike-p2-use-pfs"]; v != nil {
			ikePhase2MapToReturn["ike_p2_use_pfs"] = strconv.FormatBool(v.(bool))
		}
		if v := ikePhase2Map["ike-p2-pfs-dh-grp"]; v != nil {
			ikePhase2MapToReturn["ike_p2_pfs_dh_grp"] = v
		}
		if v := ikePhase2Map["ike-p2-rekey-time"]; v != nil {
			ikePhase2MapToReturn["ike_p2_rekey_time"] = strconv.Itoa(int(v.(float64)))
		}
		_, ikePhase2InConf := d.GetOk("ike_phase_2")
		defaultIkePhase2 := map[string]interface{}{"encryption_algorithm": "aes-128", "data_integrity": "sha1"}
		if reflect.DeepEqual(defaultIkePhase2, ikePhase2MapToReturn) && !ikePhase2InConf {
			_ = d.Set("ike_phase_2", map[string]interface{}{})
		} else {
			_ = d.Set("ike_phase_2", ikePhase2MapToReturn)
		}

	} else {
		_ = d.Set("ike_phase_2", nil)
	}

	if v := vpnCommunityStar["link-selection-mode"]; v != nil {
		_ = d.Set("link_selection_mode", v)
	}

	if vpnCommunityStar["override-interfaces"] != nil {

		overrideInterfacesList, ok := vpnCommunityStar["override-interfaces"].([]interface{})

		if ok {

			if len(overrideInterfacesList) > 0 {

				var overrideInterfacesListToReturn []map[string]interface{}

				for i := range overrideInterfacesList {

					overrideInterfacesMap := overrideInterfacesList[i].(map[string]interface{})

					overrideInterfacesMapToAdd := make(map[string]interface{})

					if v, _ := overrideInterfacesMap["gateway"]; v != nil {
						overrideInterfacesMapToAdd["gateway"] = v
					}
					if v := overrideInterfacesMap["interfaces"]; v != nil {
						interfacesShow := v.(map[string]interface{})
						interfacesState := make(map[string]interface{})
						if v := interfacesShow["interface-name"]; v != nil {
							interfacesState["interface_name"] = v
						}
						if v := interfacesShow["next-hop-ip"]; v != nil {
							interfacesState["next_hop_ip"] = v
						}
						if v := interfacesShow["static-nat-ip"]; v != nil {
							interfacesState["static_nat_ip"] = v
						}
						if v := interfacesShow["priority"]; v != nil {
							interfacesState["priority"] = v
						}
						if v := interfacesShow["redundancy-mode"]; v != nil {
							interfacesState["redundancy_mode"] = v
						}
						overrideInterfacesMapToAdd["interfaces"] = interfacesState
					}
					overrideInterfacesListToReturn = append(overrideInterfacesListToReturn, overrideInterfacesMapToAdd)
				}

				_ = d.Set("override_interfaces", overrideInterfacesListToReturn)
			} else {
				_ = d.Set("override_interfaces", overrideInterfacesList)
			}
		}
	} else {
		_ = d.Set("override_interfaces", nil)
	}

	if vpnCommunityStar["mep"] != nil {

		mepMap := vpnCommunityStar["mep"].(map[string]interface{})

		mepMapToReturn := make(map[string]interface{})

		if v := mepMap["enabled"]; v != nil {
			mepMapToReturn["enabled"] = v
		}
		if v := mepMap["entry-point-selection-mechanism"]; v != nil {
			mepMapToReturn["entry_point_selection_mechanism"] = v
		}
		if v := mepMap["entry-point-final-selection-mechanism"]; v != nil {
			mepMapToReturn["entry_point_final_selection_mechanism"] = v
		}
		if v := mepMap["tracking"]; v != nil {
			mepMapToReturn["tracking"] = v
		}
		if v := mepMap["default-priority-rule"]; v != nil {

			defaultPriorityRuleMap := v.(map[string]interface{})
			defaultPriorityRuleMapToReturn := make(map[string]interface{})

			if v, _ := defaultPriorityRuleMap["first-priority-center-gateways"]; v != nil {
				defaultPriorityRuleMapToReturn["first_priority_center_gateways"] = v
			}
			if v, _ := defaultPriorityRuleMap["second-priority-center-gateways"]; v != nil {
				defaultPriorityRuleMapToReturn["second_priority_center_gateways"] = v
			}
			if v, _ := defaultPriorityRuleMap["third-priority-center-gateways"]; v != nil {
				defaultPriorityRuleMapToReturn["third_priority_center_gateways"] = v
			}
			mepMapToReturn["default_priority_rule"] = []interface{}{defaultPriorityRuleMapToReturn}
		}

		if v := mepMap["exception-priority-rules"]; v != nil {

			exceptionPriorityRulesList := v.([]interface{})

			if len(exceptionPriorityRulesList) > 0 {

				var exceptionPriorityRulesListToReturn []map[string]interface{}

				for i := range exceptionPriorityRulesList {

					exceptionPriorityRulesMap := exceptionPriorityRulesList[i].(map[string]interface{})

					exceptionPriorityRulesMapToAdd := make(map[string]interface{})

					if v, _ := exceptionPriorityRulesMap["satellite-gateways"]; v != nil {
						exceptionPriorityRulesMapToAdd["satellite_gateways"] = v
					}
					if v, _ := exceptionPriorityRulesMap["first-priority-center-gateways"]; v != nil {
						exceptionPriorityRulesMapToAdd["first_priority_center_gateways"] = v
					}
					if v, _ := exceptionPriorityRulesMap["second-priority-center-gateways"]; v != nil {
						exceptionPriorityRulesMapToAdd["second_priority_center_gateways"] = v
					}
					if v, _ := exceptionPriorityRulesMap["third-priority-center-gateways"]; v != nil {
						exceptionPriorityRulesMapToAdd["third_priority_center_gateways"] = v
					}
					exceptionPriorityRulesListToReturn = append(exceptionPriorityRulesListToReturn, exceptionPriorityRulesMapToAdd)
				}
				mepMapToReturn["exception_priority_rules"] = exceptionPriorityRulesListToReturn
			} else {
				mepMapToReturn["exception_priority_rules"] = exceptionPriorityRulesList
			}
		} else {
			mepMapToReturn["exception_priority_rules"] = nil
		}

		_ = d.Set("mep", []interface{}{mepMapToReturn})

	} else {
		_ = d.Set("mep", nil)
	}

	if v := vpnCommunityStar["mesh-center-gateways"]; v != nil {
		_ = d.Set("mesh_center_gateways", v)
	}

	if vpnCommunityStar["override-vpn-domains"] != nil {
		overrideVpnDomainsList := vpnCommunityStar["override-vpn-domains"].([]interface{})
		var overrideVpnDomainsListToReturn []map[string]interface{}
		if len(overrideVpnDomainsList) > 0 {
			for i := range overrideVpnDomainsList {

				overrideVpnDomainsMap := overrideVpnDomainsList[i].(map[string]interface{})

				overrideVpnDomainsMapToAdd := make(map[string]interface{})

				if v, _ := overrideVpnDomainsMap["gateway"]; v != nil {
					overrideVpnDomainsMapToAdd["gateway"] = v.(map[string]interface{})["name"].(string)
				}
				if v, _ := overrideVpnDomainsMap["vpn-domain"]; v != nil {
					overrideVpnDomainsMapToAdd["vpn_domain"] = v.(map[string]interface{})["name"].(string)
				}
				overrideVpnDomainsListToReturn = append(overrideVpnDomainsListToReturn, overrideVpnDomainsMapToAdd)
			}
		}
		_ = d.Set("override_vpn_domains", overrideVpnDomainsListToReturn)
	} else {
		_ = d.Set("override_vpn_domains", nil)
	}

	if vpnCommunityStar["satellite-gateways"] != nil {
		satelliteGatewaysJson, ok := vpnCommunityStar["satellite-gateways"].([]interface{})
		if ok {
			satelliteGatewaysIds := make([]string, 0)
			if len(satelliteGatewaysJson) > 0 {
				for _, satellite_gateways := range satelliteGatewaysJson {
					satellite_gateways := satellite_gateways.(map[string]interface{})
					satelliteGatewaysIds = append(satelliteGatewaysIds, satellite_gateways["name"].(string))
				}
			}
			_ = d.Set("satellite_gateways", satelliteGatewaysIds)
		}
	} else {
		_ = d.Set("satellite_gateways", nil)
	}

	if vpnCommunityStar["shared-secrets"] != nil {
		sharedSecretsList := vpnCommunityStar["shared-secrets"].([]interface{})
		var sharedSecretsListToReturn []map[string]interface{}
		if len(sharedSecretsList) > 0 {
			for i := range sharedSecretsList {
				sharedSecretsMap := sharedSecretsList[i].(map[string]interface{})
				externalGateway := ""
				sharedSecret := "N/A"
				if v, _ := sharedSecretsMap["external-gateway"]; v != nil {
					externalGateway = v.(map[string]interface{})["name"].(string)
					if val, ok := d.GetOk("shared_secrets"); ok {
						sharedSecretsList := val.([]interface{})
						if len(sharedSecretsList) > 0 {
							for i := range sharedSecretsList {
								if v, ok := d.GetOk("shared_secrets." + strconv.Itoa(i) + ".external_gateway"); ok {
									if externalGateway == v.(string) {
										sharedSecret = d.Get("shared_secrets." + strconv.Itoa(i) + ".shared_secret").(string)
										break
									}
								}
							}
						}
					}
				}
				if externalGateway != "" {
					sharedSecretsMapToAdd := make(map[string]interface{})
					sharedSecretsMapToAdd["external_gateway"] = externalGateway
					sharedSecretsMapToAdd["shared_secret"] = sharedSecret
					sharedSecretsListToReturn = append(sharedSecretsListToReturn, sharedSecretsMapToAdd)
				}
			}
		}
		_ = d.Set("shared_secrets", sharedSecretsListToReturn)
	} else {
		_ = d.Set("shared_secrets", nil)
	}

	if v := vpnCommunityStar["tunnel-granularity"]; v != nil {
		_ = d.Set("tunnel_granularity", v)
	}

	if vpnCommunityStar["granular-encryptions"] != nil {
		granularEncryptions, ok := vpnCommunityStar["granular-encryptions"].([]interface{})
		if ok {
			if len(granularEncryptions) > 0 {
				var granularEncryptionsState []map[string]interface{}
				for i := range granularEncryptions {
					granularEncryptionShow := granularEncryptions[i].(map[string]interface{})
					granularEncryptionState := make(map[string]interface{})
					if granularEncryptionShow["internal-gateway"] != nil {
						var internalGatewayName string
						v := granularEncryptionShow["internal-gateway"]
						if obj, ok := v.(map[string]interface{}); ok {
							if obj["name"] != nil {
								internalGatewayName = obj["name"].(string)
							}
						} else if val, ok := v.(string); ok {
							internalGatewayName = val
						}
						granularEncryptionState["internal_gateway"] = internalGatewayName
					}

					if granularEncryptionShow["external-gateway"] != nil {
						var externalGatewayName string
						v := granularEncryptionShow["external-gateway"]
						if obj, ok := v.(map[string]interface{}); ok {
							if obj["name"] != nil {
								externalGatewayName = obj["name"].(string)
							}
						} else if val, ok := v.(string); ok {
							externalGatewayName = val
						}
						granularEncryptionState["external_gateway"] = externalGatewayName
					}

					if v := granularEncryptionShow["encryption-method"]; v != nil {
						granularEncryptionState["encryption_method"] = v
					}

					if v := granularEncryptionShow["encryption-suite"]; v != nil {
						granularEncryptionState["encryption_suite"] = v
					}

					if v := granularEncryptionShow["ike-phase-1"]; v != nil {
						ikePhase1Show := v.(map[string]interface{})
						ikePhase1State := make(map[string]interface{})
						if v := ikePhase1Show["encryption-algorithm"]; v != nil {
							ikePhase1State["encryption_algorithm"] = v
						}
						if v := ikePhase1Show["data-integrity"]; v != nil {
							ikePhase1State["data_integrity"] = v
						}
						if v := ikePhase1Show["diffie-hellman-group"]; v != nil {
							ikePhase1State["diffie_hellman_group"] = v
						}
						if v := ikePhase1Show["ike-p1-rekey-time"]; v != nil {
							ikePhase1State["ike_p1_rekey_time"] = strconv.Itoa(int(v.(float64)))
						}
						granularEncryptionState["ike_phase_1"] = ikePhase1State
					}

					if v := granularEncryptionShow["ike-phase-2"]; v != nil {
						ikePhase2Show := v.(map[string]interface{})
						ikePhase2State := make(map[string]interface{})
						if v := ikePhase2Show["encryption-algorithm"]; v != nil {
							ikePhase2State["encryption_algorithm"] = v
						}
						if v := ikePhase2Show["data-integrity"]; v != nil {
							ikePhase2State["data_integrity"] = v
						}
						if v := ikePhase2Show["ike-p2-use-pfs"]; v != nil {
							ikePhase2State["ike_p2_use_pfs"] = strconv.FormatBool(v.(bool))
						}
						if v := ikePhase2Show["ike-p2-pfs-dh-grp"]; v != nil {
							ikePhase2State["ike_p2_pfs_dh_grp"] = v
						}
						if v := ikePhase2Show["ike-p2-rekey-time"]; v != nil {
							ikePhase2State["ike_p2_rekey_time"] = strconv.Itoa(int(v.(float64)))
						}
						granularEncryptionState["ike_phase_2"] = ikePhase2State
					}
					granularEncryptionsState = append(granularEncryptionsState, granularEncryptionState)
				}
				_ = d.Set("granular_encryptions", granularEncryptionsState)
			} else {
				_ = d.Set("granular_encryptions", nil)
			}
		}
	}

	if vpnCommunityStar["permanent-tunnels"] != nil {

		permanentTunnelsMap := vpnCommunityStar["permanent-tunnels"].(map[string]interface{})

		permanentTunnelsMapToReturn := make(map[string]interface{})

		if v := permanentTunnelsMap["set-permanent-tunnels"]; v != nil {
			permanentTunnelsMapToReturn["set_permanent_tunnels"] = v
		}
		if v := permanentTunnelsMap["gateways"]; v != nil {

			gatewaysList := v.([]interface{})

			if len(gatewaysList) > 0 {

				var gatewaysListToReturn []map[string]interface{}

				for i := range gatewaysList {

					gatewaysMap := gatewaysList[i].(map[string]interface{})

					gatewaysMapToAdd := make(map[string]interface{})

					if v := gatewaysMap["gateway"]; v != nil {
						gatewayObj := v.(map[string]interface{})
						if v := gatewayObj["name"]; v != nil {
							gatewaysMapToAdd["gateway"] = v.(string)
						}
					}
					if v, _ := gatewaysMap["track-options"]; v != nil {
						gatewaysMapToAdd["track_options"] = v
					}
					if v, _ := gatewaysMap["override-tunnel-down-track"]; v != nil {
						gatewaysMapToAdd["override_tunnel_down_track"] = v
					}
					if v, _ := gatewaysMap["override-tunnel-up-track"]; v != nil {
						gatewaysMapToAdd["override_tunnel_up_track"] = v
					}
					gatewaysListToReturn = append(gatewaysListToReturn, gatewaysMapToAdd)
				}
				permanentTunnelsMapToReturn["gateways"] = gatewaysListToReturn
			} else {
				permanentTunnelsMapToReturn["gateways"] = gatewaysList
			}
		} else {
			permanentTunnelsMapToReturn["gateways"] = nil
		}

		if v := permanentTunnelsMap["tunnels"]; v != nil {

			tunnelsList := v.([]interface{})

			if len(tunnelsList) > 0 {

				var tunnelsListToReturn []map[string]interface{}

				for i := range tunnelsList {

					tunnelsMap := tunnelsList[i].(map[string]interface{})

					tunnelsMapToAdd := make(map[string]interface{})

					if v, _ := tunnelsMap["first-tunnel-endpoint"]; v != nil {
						tunnelsMapToAdd["first_tunnel_endpoint"] = v
					}
					if v, _ := tunnelsMap["second-tunnel-endpoint"]; v != nil {
						tunnelsMapToAdd["second_tunnel_endpoint"] = v
					}
					if v, _ := tunnelsMap["track-options"]; v != nil {
						tunnelsMapToAdd["track_options"] = v
					}
					if v, _ := tunnelsMap["override-tunnel-down-track"]; v != nil {
						tunnelsMapToAdd["override_tunnel_down_track"] = v
					}
					if v, _ := tunnelsMap["override-tunnel-up-track"]; v != nil {
						tunnelsMapToAdd["override_tunnel_up_track"] = v
					}
					tunnelsListToReturn = append(tunnelsListToReturn, tunnelsMapToAdd)
				}
				permanentTunnelsMapToReturn["tunnels"] = tunnelsListToReturn
			} else {
				permanentTunnelsMapToReturn["tunnels"] = tunnelsList
			}
		} else {
			permanentTunnelsMapToReturn["tunnels"] = nil
		}

		if v := permanentTunnelsMap["rim"]; v != nil {

			rimMap := v.(map[string]interface{})
			rimMapToReturn := make(map[string]interface{})

			if v, _ := rimMap["enabled"]; v != nil {
				rimMapToReturn["enabled"] = v
			}
			if v, _ := rimMap["enable-on-center-gateways"]; v != nil {
				rimMapToReturn["enable_on_center_gateways"] = v
			}
			if v, _ := rimMap["enable-on-satellite-gateways"]; v != nil {
				rimMapToReturn["enable_on_satellite_gateways"] = v
			}
			if v, _ := rimMap["route-injection-track"]; v != nil {
				rimMapToReturn["route_injection_track"] = v
			}
			permanentTunnelsMapToReturn["rim"] = []interface{}{rimMapToReturn}
		}

		if v := permanentTunnelsMap["tunnel-down-track"]; v != nil {
			permanentTunnelsMapToReturn["tunnel_down_track"] = v
		}
		if v := permanentTunnelsMap["tunnel-up-track"]; v != nil {
			permanentTunnelsMapToReturn["tunnel_up_track"] = v
		}
		_ = d.Set("permanent_tunnels", []interface{}{permanentTunnelsMapToReturn})

	} else {
		_ = d.Set("permanent_tunnels", nil)
	}

	if v := vpnCommunityStar["vpn-routing"]; v != nil {
		_ = d.Set("vpn_routing", v)
	}

	if vpnCommunityStar["wire-mode"] != nil {

		wireModeMap := vpnCommunityStar["wire-mode"].(map[string]interface{})

		wireModeMapToReturn := make(map[string]interface{})

		if v := wireModeMap["allow-uninspected-encrypted-traffic"]; v != nil {
			wireModeMapToReturn["allow_uninspected_encrypted_traffic"] = v
		}
		if v := wireModeMap["allow-uninspected-encrypted-routing"]; v != nil {
			wireModeMapToReturn["allow_uninspected_encrypted_routing"] = v
		}
		_ = d.Set("wire_mode", []interface{}{wireModeMapToReturn})

	} else {
		_ = d.Set("wire_mode", nil)
	}

	if v := vpnCommunityStar["routing-mode"]; v != nil {
		_ = d.Set("routing_mode", v)
	}

	if vpnCommunityStar["advanced-properties"] != nil {

		advancedPropertiesMap := vpnCommunityStar["advanced-properties"].(map[string]interface{})

		advancedPropertiesMapToReturn := make(map[string]interface{})

		if v := advancedPropertiesMap["support-ip-compression"]; v != nil {
			advancedPropertiesMapToReturn["support_ip_compression"] = v
		}
		if v := advancedPropertiesMap["use-aggressive-mode"]; v != nil {
			advancedPropertiesMapToReturn["use_aggressive_mode"] = v
		}
		_ = d.Set("advanced_properties", []interface{}{advancedPropertiesMapToReturn})

	} else {
		_ = d.Set("advanced_properties", nil)
	}

	if vpnCommunityStar["tags"] != nil {
		tagsJson, ok := vpnCommunityStar["tags"].([]interface{})
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

	if v := vpnCommunityStar["use-shared-secret"]; v != nil {
		_ = d.Set("use_shared_secret", v)
	}

	if v := vpnCommunityStar["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := vpnCommunityStar["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil
}
