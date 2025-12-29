package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"reflect"
	"strconv"
)

func resourceManagementVpnCommunityStar() *schema.Resource {
	return &schema.Resource{
		Create: createManagementVpnCommunityStar,
		Read:   readManagementVpnCommunityStar,
		Update: updateManagementVpnCommunityStar,
		Delete: deleteManagementVpnCommunityStar,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"center_gateways": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of Gateway objects representing center gateways identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"disable_nat": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates whether to disable NAT inside the VPN Community.",
			},
			"disable_nat_on": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Indicates on which gateways to disable NAT inside the VPN Community.",
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
						"community_members": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Indicates on which community members to accept all encrypted traffic.",
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
			"mep": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Multiple Entry Point properties.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable center gateways as Multiple Entry Points.",
						},
						"entry_point_selection_mechanism": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The method by which the entry point gateway will be chosen from the gateways in the center.",
						},
						"entry_point_final_selection_mechanism": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The method by which the final entry point gateway will be chosen when the chosen mechanism returns more than one optional entry point.",
						},
						"tracking": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Tracking option for the MEP.",
						},
						"default_priority_rule": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Priority rule for all satellite gateways. Relevant only if 'entry-point-selection-mechanism' is set to 'manual'.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"first_priority_center_gateways": {
										Type:        schema.TypeSet,
										Optional:    true,
										Description: "Collection of first priority center gateways identified by the name or UID.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"second_priority_center_gateways": {
										Type:        schema.TypeSet,
										Optional:    true,
										Description: "Collection of second priority center gateways identified by the name or UID.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"third_priority_center_gateways": {
										Type:        schema.TypeSet,
										Optional:    true,
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
							Optional:    true,
							Description: "Exception priority rules for specific satellites gateways. Relevant only if 'entry-point-selection-mechanism' is set to 'manual'.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"satellite_gateways": {
										Type:        schema.TypeSet,
										Optional:    true,
										Description: "Collection of satellite gateways to apply priority rules on identified by the name or UID.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"first_priority_center_gateways": {
										Type:        schema.TypeSet,
										Optional:    true,
										Description: "Collection of first priority center gateways identified by the name or UID.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"second_priority_center_gateways": {
										Type:        schema.TypeSet,
										Optional:    true,
										Description: "Collection of second priority center gateways identified by the name or UID.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"third_priority_center_gateways": {
										Type:        schema.TypeSet,
										Optional:    true,
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
				Optional:    true,
				Description: "Indicates whether the meshed community is in center.",
				Default:     false,
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
			"satellite_gateways": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of Gateway objects representing satellite gateways identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
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
									"enable_on_center_gateways": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Indicates whether to enable automatic Route Injection Mechanism on center gateways.",
									},
									"enable_on_satellite_gateways": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Indicates whether to enable automatic Route Injection Mechanism on satellite gateways.",
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
			"vpn_routing": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Enable VPN routing to satellites.",
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

func createManagementVpnCommunityStar(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	vpnCommunityStar := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		vpnCommunityStar["name"] = v.(string)
	}

	if v, ok := d.GetOk("center_gateways"); ok {
		vpnCommunityStar["center-gateways"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("disable_nat"); ok {
		vpnCommunityStar["disable-nat"] = v.(bool)
	}

	if v, ok := d.GetOk("disable_nat_on"); ok {
		vpnCommunityStar["disable-nat-on"] = v.(string)
	}

	if v, ok := d.GetOk("encrypted_traffic"); ok {

		encryptedTrafficList := v.([]interface{})

		if len(encryptedTrafficList) > 0 {

			encryptedTrafficPayload := make(map[string]interface{})

			if v, ok := d.GetOk("encrypted_traffic.0.enabled"); ok {
				encryptedTrafficPayload["enabled"] = v.(bool)
			}
			if v, ok := d.GetOk("encrypted_traffic.0.community_members"); ok {
				encryptedTrafficPayload["community-members"] = v.(string)
			}
			vpnCommunityStar["encrypted-traffic"] = encryptedTrafficPayload
		}
	}

	if v, ok := d.GetOk("encryption_method"); ok {
		vpnCommunityStar["encryption-method"] = v.(string)
	}

	if v, ok := d.GetOk("encryption_suite"); ok {
		vpnCommunityStar["encryption-suite"] = v.(string)
	}

	if v, ok := d.GetOk("excluded_services"); ok {
		vpnCommunityStar["excluded-services"] = v.(*schema.Set).List()
	}

	if _, ok := d.GetOk("ike_phase_1"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("ike_phase_1.data_integrity"); ok {
			res["data-integrity"] = v.(string)
		}
		if v, ok := d.GetOk("ike_phase_1.diffie_hellman_group"); ok {
			res["diffie-hellman-group"] = v.(string)
		}
		if v, ok := d.GetOk("ike_phase_1.encryption_algorithm"); ok {
			res["encryption-algorithm"] = v.(string)
		}
		if v, ok := d.GetOk("ike_phase_1.ike_p1_rekey_time"); ok {
			res["ike-p1-rekey-time"] = v.(string)
		}
		vpnCommunityStar["ike-phase-1"] = res
	}

	if _, ok := d.GetOk("ike_phase_2"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("ike_phase_2.data_integrity"); ok {
			res["data-integrity"] = v.(string)
		}
		if v, ok := d.GetOk("ike_phase_2.encryption_algorithm"); ok {
			res["encryption-algorithm"] = v.(string)
		}
		if v, ok := d.GetOk("ike_phase_2.ike_p2_use_pfs"); ok {
			res["ike-p2-use-pfs"] = v.(string)
		}
		if v, ok := d.GetOk("ike_phase_2.ike_p2_pfs_dh_grp"); ok {
			res["ike-p2-pfs-dh-grp"] = v.(string)
		}
		if v, ok := d.GetOk("ike_phase_2.ike_p2_rekey_time"); ok {
			res["ike-p2-rekey-time"] = v.(string)
		}
		vpnCommunityStar["ike-phase-2"] = res
	}

	if v, ok := d.GetOk("link_selection_mode"); ok {
		vpnCommunityStar["link-selection-mode"] = v.(string)
	}

	if v, ok := d.GetOk("override_interfaces"); ok {

		overrideInterfacesList := v.([]interface{})

		if len(overrideInterfacesList) > 0 {

			var overrideInterfacesPayload []map[string]interface{}

			for i := range overrideInterfacesList {

				Payload := make(map[string]interface{})

				if v, ok := d.GetOk("override_interfaces." + strconv.Itoa(i) + ".gateway"); ok {
					Payload["gateway"] = v.(string)
				}
				if v, ok := d.GetOk("override_interfaces." + strconv.Itoa(i) + ".interfaces"); ok {

					overrideInterfacesInterfacesList := v.([]interface{})

					if len(overrideInterfacesInterfacesList) > 0 {

						var overrideInterfacesInterfacesPayload []map[string]interface{}

						for j := range overrideInterfacesInterfacesList {

							interfacesPayload := make(map[string]interface{})

							if v, ok := d.GetOk("override_interfaces." + strconv.Itoa(i) + ".interfaces." + strconv.Itoa(j) + ".interface_name"); ok {
								interfacesPayload["interface-name"] = v.(string)
							}
							if v, ok := d.GetOk("override_interfaces." + strconv.Itoa(i) + ".interfaces." + strconv.Itoa(j) + ".next_hop_ip"); ok {
								interfacesPayload["next-hop-ip"] = v.(string)
							}
							if v, ok := d.GetOk("override_interfaces." + strconv.Itoa(i) + ".interfaces." + strconv.Itoa(j) + ".static_nat_ip"); ok {
								interfacesPayload["static-nat-ip"] = v.(string)
							}
							if v, ok := d.GetOk("override_interfaces." + strconv.Itoa(i) + ".interfaces." + strconv.Itoa(j) + ".priority"); ok {
								interfacesPayload["priority"] = v.(int)
							}
							if v, ok := d.GetOk("override_interfaces." + strconv.Itoa(i) + ".interfaces." + strconv.Itoa(j) + ".redundancy_mode"); ok {
								interfacesPayload["redundancy-mode"] = v.(string)
							}
							overrideInterfacesInterfacesPayload = append(overrideInterfacesInterfacesPayload, interfacesPayload)
						}
						Payload["interfaces"] = overrideInterfacesInterfacesPayload
					}
				}

				overrideInterfacesPayload = append(overrideInterfacesPayload, Payload)
			}
			vpnCommunityStar["overrideInterfaces"] = overrideInterfacesPayload
		}
	}

	if v, ok := d.GetOk("mep"); ok {

		mepList := v.([]interface{})

		if len(mepList) > 0 {

			mepPayload := make(map[string]interface{})

			if v, ok := d.GetOk("mep.0.enabled"); ok {
				mepPayload["enabled"] = v.(bool)
			}
			if v, ok := d.GetOk("mep.0.entry_point_selection_mechanism"); ok {
				mepPayload["entry-point-selection-mechanism"] = v.(string)
			}
			if v, ok := d.GetOk("mep.0.entry_point_final_selection_mechanism"); ok {
				mepPayload["entry-point-final-selection-mechanism"] = v.(string)
			}
			if v, ok := d.GetOk("mep.0.tracking"); ok {
				mepPayload["tracking"] = v.(string)
			}
			if _, ok := d.GetOk("mep.0.default_priority_rule"); ok {

				defaultPriorityRulePayload := make(map[string]interface{})

				if v, ok := d.GetOk("mep.0.default_priority_rule.0.first_priority_center_gateways"); ok {
					defaultPriorityRulePayload["first-priority-center-gateways"] = v.(*schema.Set).List()
				}
				if v, ok := d.GetOk("mep.0.default_priority_rule.0.second_priority_center_gateways"); ok {
					defaultPriorityRulePayload["second-priority-center-gateways"] = v.(*schema.Set).List()
				}
				if v, ok := d.GetOk("mep.0.default_priority_rule.0.third_priority_center_gateways"); ok {
					defaultPriorityRulePayload["third-priority-center-gateways"] = v.(*schema.Set).List()
				}
				mepPayload["default-priority-rule"] = defaultPriorityRulePayload
			}
			if _, ok := d.GetOk("mep.0.exception_priority_rules"); ok {

				exceptionPriorityRulesPayload := make(map[string]interface{})

				if v, ok := d.GetOk("mep.0.exception_priority_rules.0.satellite_gateways"); ok {
					exceptionPriorityRulesPayload["satellite-gateways"] = v.(*schema.Set).List()
				}
				if v, ok := d.GetOk("mep.0.exception_priority_rules.0.first_priority_center_gateways"); ok {
					exceptionPriorityRulesPayload["first-priority-center-gateways"] = v.(*schema.Set).List()
				}
				if v, ok := d.GetOk("mep.0.exception_priority_rules.0.second_priority_center_gateways"); ok {
					exceptionPriorityRulesPayload["second-priority-center-gateways"] = v.(*schema.Set).List()
				}
				if v, ok := d.GetOk("mep.0.exception_priority_rules.0.third_priority_center_gateways"); ok {
					exceptionPriorityRulesPayload["third-priority-center-gateways"] = v.(*schema.Set).List()
				}
				mepPayload["exception-priority-rules"] = exceptionPriorityRulesPayload
			}
			vpnCommunityStar["mep"] = mepPayload
		}
	}

	if v, ok := d.GetOkExists("mesh_center_gateways"); ok {
		vpnCommunityStar["mesh-center-gateways"] = v.(bool)
	}

	if v, ok := d.GetOk("override_vpn_domains"); ok {

		overrideVpnDomainsList := v.([]interface{})

		if len(overrideVpnDomainsList) > 0 {

			var overrideVpnDomainsPayload []map[string]interface{}

			for i := range overrideVpnDomainsList {

				Payload := make(map[string]interface{})

				if v, ok := d.GetOk("override_vpn_domains." + strconv.Itoa(i) + ".gateway"); ok {
					Payload["gateway"] = v.(string)
				}
				if v, ok := d.GetOk("override_vpn_domains." + strconv.Itoa(i) + ".vpn_domain"); ok {
					Payload["vpn-domain"] = v.(string)
				}
				overrideVpnDomainsPayload = append(overrideVpnDomainsPayload, Payload)
			}
			vpnCommunityStar["override-vpn-domains"] = overrideVpnDomainsPayload
		}
	}

	if v, ok := d.GetOk("satellite_gateways"); ok {
		vpnCommunityStar["satellite-gateways"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("shared_secrets"); ok {

		sharedSecretsList := v.([]interface{})

		if len(sharedSecretsList) > 0 {

			var sharedSecretsPayload []map[string]interface{}

			for i := range sharedSecretsList {

				Payload := make(map[string]interface{})

				if v, ok := d.GetOk("shared_secrets." + strconv.Itoa(i) + ".external_gateway"); ok {
					Payload["external-gateway"] = v.(string)
				}
				if v, ok := d.GetOk("shared_secrets." + strconv.Itoa(i) + ".shared_secret"); ok {
					Payload["shared-secret"] = v.(string)
				}
				sharedSecretsPayload = append(sharedSecretsPayload, Payload)
			}
			vpnCommunityStar["shared-secrets"] = sharedSecretsPayload
		}
	}

	if v, ok := d.GetOk("tunnel_granularity"); ok {
		vpnCommunityStar["tunnel-granularity"] = v.(string)
	}

	if v, ok := d.GetOk("granular_encryptions"); ok {
		granularEncryptions := v.([]interface{})
		if len(granularEncryptions) > 0 {
			var granularEncryptionsPayload []map[string]interface{}

			for i := range granularEncryptions {
				payload := make(map[string]interface{})
				if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".internal_gateway"); ok {
					payload["internal-gateway"] = v.(string)
				}
				if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".external_gateway"); ok {
					payload["external-gateway"] = v.(string)
				}
				if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".encryption_method"); ok {
					payload["encryption-method"] = v.(string)
				}
				if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".encryption_suite"); ok {
					payload["encryption-suite"] = v.(string)
				}
				if _, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_1"); ok {
					ikePhase1Payload := make(map[string]interface{})
					if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_1.encryption_algorithm"); ok {
						ikePhase1Payload["encryption-algorithm"] = v.(string)
					}
					if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_1.data_integrity"); ok {
						ikePhase1Payload["data-integrity"] = v.(string)
					}
					if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_1.diffie_hellman_group"); ok {
						ikePhase1Payload["diffie-hellman-group"] = v.(string)
					}
					if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_1.ike_p1_rekey_time"); ok {
						ikePhase1Payload["ike-p1-rekey-time"] = v.(string)
					}
					payload["ike-phase-1"] = ikePhase1Payload
				}
				if _, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_2"); ok {
					ikePhase2Payload := make(map[string]interface{})
					if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_2.encryption_algorithm"); ok {
						ikePhase2Payload["encryption-algorithm"] = v.(string)
					}
					if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_2.data_integrity"); ok {
						ikePhase2Payload["data-integrity"] = v.(string)
					}
					if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_2.ike_p2_use_pfs"); ok {
						ikePhase2Payload["ike-p2-use-pfs"] = v.(string)
					}
					if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_2.ike_p2_pfs_dh_grp"); ok {
						ikePhase2Payload["ike-p2-pfs-dh-grp"] = v.(string)
					}
					if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_2.ike_p2_rekey_time"); ok {
						ikePhase2Payload["ike-p2-rekey-time"] = v.(string)
					}
					payload["ike-phase-2"] = ikePhase2Payload
				}
				granularEncryptionsPayload = append(granularEncryptionsPayload, payload)
			}
			vpnCommunityStar["granular-encryptions"] = granularEncryptionsPayload
		}
	}

	if v, ok := d.GetOk("permanent_tunnels"); ok {

		permanentTunnelsList := v.([]interface{})

		if len(permanentTunnelsList) > 0 {

			permanentTunnelsPayload := make(map[string]interface{})

			if v, ok := d.GetOk("permanent_tunnels.0.set_permanent_tunnels"); ok {
				permanentTunnelsPayload["set-permanent-tunnels"] = v.(string)
			}
			if _, ok := d.GetOk("permanent_tunnels.0.gateways"); ok {

				gatewaysPayload := make(map[string]interface{})

				if v, ok := d.GetOk("permanent_tunnels.0.gateways.0.gateway"); ok {
					gatewaysPayload["gateway"] = v.(string)
				}
				if v, ok := d.GetOk("permanent_tunnels.0.gateways.0.track_options"); ok {
					gatewaysPayload["track-options"] = v.(string)
				}
				if v, ok := d.GetOk("permanent_tunnels.0.gateways.0.override_tunnel_down_track"); ok {
					gatewaysPayload["override-tunnel-down-track"] = v.(string)
				}
				if v, ok := d.GetOk("permanent_tunnels.0.gateways.0.override_tunnel_up_track"); ok {
					gatewaysPayload["override-tunnel-up-track"] = v.(string)
				}
				permanentTunnelsPayload["gateways"] = gatewaysPayload
			}
			if _, ok := d.GetOk("permanent_tunnels.0.tunnels"); ok {

				tunnelsPayload := make(map[string]interface{})

				if v, ok := d.GetOk("permanent_tunnels.0.tunnels.0.first_tunnel_endpoint"); ok {
					tunnelsPayload["first-tunnel-endpoint"] = v.(string)
				}
				if v, ok := d.GetOk("permanent_tunnels.0.tunnels.0.second_tunnel_endpoint"); ok {
					tunnelsPayload["second-tunnel-endpoint"] = v.(string)
				}
				if v, ok := d.GetOk("permanent_tunnels.0.tunnels.0.track_options"); ok {
					tunnelsPayload["track-options"] = v.(string)
				}
				if v, ok := d.GetOk("permanent_tunnels.0.tunnels.0.override_tunnel_down_track"); ok {
					tunnelsPayload["override-tunnel-down-track"] = v.(string)
				}
				if v, ok := d.GetOk("permanent_tunnels.0.tunnels.0.override_tunnel_up_track"); ok {
					tunnelsPayload["override-tunnel-up-track"] = v.(string)
				}
				permanentTunnelsPayload["tunnels"] = tunnelsPayload
			}
			if _, ok := d.GetOk("permanent_tunnels.0.rim"); ok {

				rimPayload := make(map[string]interface{})

				if v, ok := d.GetOk("permanent_tunnels.0.rim.0.enabled"); ok {
					rimPayload["enabled"] = strconv.FormatBool(v.(bool))
				}
				if v, ok := d.GetOk("permanent_tunnels.0.rim.0.enable_on_center_gateways"); ok {
					rimPayload["enable-on-center-gateways"] = strconv.FormatBool(v.(bool))
				}
				if v, ok := d.GetOk("permanent_tunnels.0.rim.0.enable_on_satellite_gateways"); ok {
					rimPayload["enable-on-satellite-gateways"] = strconv.FormatBool(v.(bool))
				}
				if v, ok := d.GetOk("permanent_tunnels.0.rim.0.route_injection_track"); ok {
					rimPayload["route-injection-track"] = v.(string)
				}
				permanentTunnelsPayload["rim"] = rimPayload
			}
			if v, ok := d.GetOk("permanent_tunnels.0.tunnel_down_track"); ok {
				permanentTunnelsPayload["tunnel-down-track"] = v.(string)
			}
			if v, ok := d.GetOk("permanent_tunnels.0.tunnel_up_track"); ok {
				permanentTunnelsPayload["tunnel-up-track"] = v.(string)
			}
			vpnCommunityStar["permanent-tunnels"] = permanentTunnelsPayload
		}
	}

	if v, ok := d.GetOk("vpn_routing"); ok {
		vpnCommunityStar["vpn-routing"] = v.(string)
	}

	if v, ok := d.GetOk("wire_mode"); ok {

		wireModeList := v.([]interface{})

		if len(wireModeList) > 0 {

			wireModePayload := make(map[string]interface{})

			if v, ok := d.GetOk("wire_mode.0.allow_uninspected_encrypted_traffic"); ok {
				wireModePayload["allow-uninspected-encrypted-traffic"] = v.(bool)
			}
			if v, ok := d.GetOk("wire_mode.0.allow_uninspected_encrypted_routing"); ok {
				wireModePayload["allow-uninspected-encrypted-routing"] = v.(bool)
			}
			vpnCommunityStar["wire-mode"] = wireModePayload
		}
	}
	if v, ok := d.GetOk("routing_mode"); ok {
		vpnCommunityStar["routing-mode"] = v.(string)
	}

	if v, ok := d.GetOk("advanced_properties"); ok {

		advancedPropertiesList := v.([]interface{})

		if len(advancedPropertiesList) > 0 {

			advancedPropertiesPayload := make(map[string]interface{})

			if v, ok := d.GetOk("advanced_properties.0.support_ip_compression"); ok {
				advancedPropertiesPayload["support-ip-compression"] = v.(bool)
			}
			if v, ok := d.GetOk("advanced_properties.0.use_aggressive_mode"); ok {
				advancedPropertiesPayload["use-aggressive-mode"] = v.(bool)
			}
			vpnCommunityStar["advanced-properties"] = advancedPropertiesPayload
		}
	}

	if v, ok := d.GetOk("tags"); ok {
		vpnCommunityStar["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("use_shared_secret"); ok {
		vpnCommunityStar["use-shared-secret"] = v.(bool)
	}

	if v, ok := d.GetOk("color"); ok {
		vpnCommunityStar["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		vpnCommunityStar["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		vpnCommunityStar["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		vpnCommunityStar["ignore-errors"] = v.(bool)
	}

	log.Println("Create VpnCommunityStar - Map = ", vpnCommunityStar)

	addVpnCommunityStarRes, err := client.ApiCall("add-vpn-community-star", vpnCommunityStar, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !addVpnCommunityStarRes.Success {
		if addVpnCommunityStarRes.ErrorMsg != "" {
			return fmt.Errorf(addVpnCommunityStarRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addVpnCommunityStarRes.GetData()["uid"].(string))

	return readManagementVpnCommunityStar(d, m)
}

func readManagementVpnCommunityStar(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showVpnCommunityStarRes, err := client.ApiCall("show-vpn-community-star", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showVpnCommunityStarRes.Success {
		if objectNotFound(showVpnCommunityStarRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showVpnCommunityStarRes.ErrorMsg)
	}

	vpnCommunityStar := showVpnCommunityStarRes.GetData()

	log.Println("Read VpnCommunityStar - Show JSON = ", vpnCommunityStar)

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
					if v := tunnelsMap["first-tunnel-endpoint"]; v != nil {
						tunnelObj := v.(map[string]interface{})
						if v := tunnelObj["name"]; v != nil {
							tunnelsMapToAdd["first_tunnel_endpoint"] = v.(string)
						}
					}
					if v := tunnelsMap["second-tunnel-endpoint"]; v != nil {
						tunnelObj := v.(map[string]interface{})
						if v := tunnelObj["name"]; v != nil {
							tunnelsMapToAdd["second_tunnel_endpoint"] = v.(string)
						}
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

	if v := vpnCommunityStar["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := vpnCommunityStar["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementVpnCommunityStar(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	vpnCommunityStar := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		vpnCommunityStar["name"] = oldName
		vpnCommunityStar["new-name"] = newName
	} else {
		vpnCommunityStar["name"] = d.Get("name")
	}

	if d.HasChange("center_gateways") {
		if v, ok := d.GetOk("center_gateways"); ok {
			vpnCommunityStar["center-gateways"] = v.(*schema.Set).List()
		} else {
			oldCenter_Gateways, _ := d.GetChange("center_gateways")
			vpnCommunityStar["center-gateways"] = map[string]interface{}{"remove": oldCenter_Gateways.(*schema.Set).List()}
		}
	}

	if v, ok := d.GetOkExists("disable_nat"); ok {
		vpnCommunityStar["disable-nat"] = v.(bool)
	}

	if ok := d.HasChange("disable_nat_on"); ok {
		vpnCommunityStar["disable-nat-on"] = d.Get("disable_nat_on")
	}

	if d.HasChange("encrypted_traffic") {

		if v, ok := d.GetOk("encrypted_traffic"); ok {

			encryptedTrafficList := v.([]interface{})

			if len(encryptedTrafficList) > 0 {

				encryptedTrafficPayload := make(map[string]interface{})

				if v, ok := d.GetOk("encrypted_traffic.0.enabled"); ok {
					encryptedTrafficPayload["enabled"] = v.(bool)
				}
				if v, ok := d.GetOk("encrypted_traffic.0.community_members"); ok {
					encryptedTrafficPayload["community-members"] = v.(string)
				}
				vpnCommunityStar["encrypted-traffic"] = encryptedTrafficPayload
			}
		}
	}

	if ok := d.HasChange("encryption_method"); ok {
		vpnCommunityStar["encryption-method"] = d.Get("encryption_method")
	}

	if ok := d.HasChange("encryption_suite"); ok {
		vpnCommunityStar["encryption-suite"] = d.Get("encryption_suite")
	}

	if d.HasChange("excluded_services") {
		if v, ok := d.GetOk("excluded_services"); ok {
			vpnCommunityStar["excluded_services"] = v.(*schema.Set).List()
		}
	}

	if d.HasChange("ike_phase_1") {

		if _, ok := d.GetOk("ike_phase_1"); ok {

			res := make(map[string]interface{})

			if d.HasChange("ike_phase_1.data_integrity") {
				res["data-integrity"] = d.Get("ike_phase_1.data_integrity")
			}
			if d.HasChange("ike_phase_1.diffie_hellman_group") {
				res["diffie-hellman-group"] = d.Get("ike_phase_1.diffie_hellman_group")
			}
			if d.HasChange("ike_phase_1.encryption_algorithm") {
				res["encryption-algorithm"] = d.Get("ike_phase_1.encryption_algorithm")
			}
			if d.HasChange("ike_phase_1.ike_p1_rekey_time") {
				res["ike-p1-rekey-time"] = d.Get("ike_phase_1.ike_p1_rekey_time")
			}
			vpnCommunityStar["ike-phase-1"] = res
		} else {
			vpnCommunityStar["ike-phase-1"] = map[string]interface{}{"encryption-algorithm": "aes-256", "diffie-hellman-group": "group-2", "data-integrity": "sha1"}
		}
	}

	if d.HasChange("ike_phase_2") {

		if _, ok := d.GetOk("ike_phase_2"); ok {

			res := make(map[string]interface{})

			if d.HasChange("ike_phase_2.data_integrity") {
				res["data-integrity"] = d.Get("ike_phase_2.data_integrity")
			}
			if d.HasChange("ike_phase_2.encryption_algorithm") {
				res["encryption-algorithm"] = d.Get("ike_phase_2.encryption_algorithm")
			}
			if d.HasChange("ike_phase_2.ike_p2_use_pfs") {
				res["ike-p2-use-pfs"] = d.Get("ike_phase_2.ike_p2_use_pfs")
			}
			if d.HasChange("ike_phase_2.ike_p2_pfs_dh_grp") {
				res["ike-p2-pfs-dh-grp"] = d.Get("ike_phase_2.ike_p2_pfs_dh_grp")
			}
			if d.HasChange("ike_phase_2.ike_p2_rekey_time") {
				res["ike-p2-rekey-time"] = d.Get("ike_phase_2.ike_p2_rekey_time")
			}
			vpnCommunityStar["ike-phase-2"] = res
		} else {
			vpnCommunityStar["ike-phase-2"] = map[string]interface{}{"encryption-algorithm": "aes-128", "data-integrity": "sha1"}
		}
	}

	if ok := d.HasChange("link_selection_mode"); ok {
		vpnCommunityStar["link-selection-mode"] = d.Get("link_selection_mode")
	}

	if d.HasChange("override_interfaces") {

		if v, ok := d.GetOk("override_interfaces"); ok {

			overrideInterfacesList := v.([]interface{})

			var overrideInterfacesPayload []map[string]interface{}

			for i := range overrideInterfacesList {

				Payload := make(map[string]interface{})

				if v, ok := d.GetOk("override_interfaces." + strconv.Itoa(i) + ".gateway"); ok {
					Payload["gateway"] = v
				}
				if d.HasChange("override_interfaces." + strconv.Itoa(i) + ".interfaces") {

					if v, ok := d.GetOk("override_interfaces." + strconv.Itoa(i) + ".interfaces"); ok {

						overrideInterfacesInterfacesList := v.([]interface{})

						var overrideInterfacesInterfacesPayload []map[string]interface{}

						for j := range overrideInterfacesInterfacesList {

							interfacesPayload := make(map[string]interface{})

							if v, ok := d.GetOk("override_interfaces." + strconv.Itoa(i) + ".interfaces." + strconv.Itoa(j) + ".interface_name"); ok {
								interfacesPayload["interface-name"] = v
							}
							if v, ok := d.GetOk("override_interfaces." + strconv.Itoa(i) + ".interfaces." + strconv.Itoa(j) + ".next_hop_ip"); ok {
								interfacesPayload["next-hop-ip"] = v
							}
							if v, ok := d.GetOk("override_interfaces." + strconv.Itoa(i) + ".interfaces." + strconv.Itoa(j) + ".static_nat_ip"); ok {
								interfacesPayload["static-nat-ip"] = v
							}
							if v, ok := d.GetOk("override_interfaces." + strconv.Itoa(i) + ".interfaces." + strconv.Itoa(j) + ".priority"); ok {
								interfacesPayload["priority"] = v
							}
							if v, ok := d.GetOk("override_interfaces." + strconv.Itoa(i) + ".interfaces." + strconv.Itoa(j) + ".redundancy_mode"); ok {
								interfacesPayload["redundancy-mode"] = v
							}
							overrideInterfacesInterfacesPayload = append(overrideInterfacesInterfacesPayload, interfacesPayload)
						}
						Payload["interfaces"] = overrideInterfacesInterfacesPayload
					}
				}

				overrideInterfacesPayload = append(overrideInterfacesPayload, Payload)
			}
			vpnCommunityStar["override-interfaces"] = overrideInterfacesPayload
		}
	}

	if d.HasChange("mep") {

		if v, ok := d.GetOk("mep"); ok {

			mepList := v.([]interface{})

			if len(mepList) > 0 {

				mepPayload := make(map[string]interface{})

				if v, ok := d.GetOkExists("mep.0.enabled"); ok {
					mepPayload["enabled"] = v.(bool)
				}
				if v, ok := d.GetOk("mep.0.entry_point_selection_mechanism"); ok {
					mepPayload["entry-point-selection-mechanism"] = v.(string)
				}
				if v, ok := d.GetOk("mep.0.entry_point_final_selection_mechanism"); ok {
					mepPayload["entry-point-final-selection-mechanism"] = v.(string)
				}
				if v, ok := d.GetOk("mep.0.tracking"); ok {
					mepPayload["tracking"] = v.(string)
				}
				if _, ok := d.GetOk("mep.0.default_priority_rule"); ok {
					defaultPriorityRulePayload := make(map[string]interface{})

					if v, ok := d.GetOk("mep.0.default_priority_rule.0.first_priority_center_gateways"); ok {
						defaultPriorityRulePayload["first-priority-center-gateways"] = v.(*schema.Set).List()
					}
					if v, ok := d.GetOk("mep.0.default_priority_rule.0.second_priority_center_gateways"); ok {
						defaultPriorityRulePayload["second-priority-center-gateways"] = v.(*schema.Set).List()
					}
					if v, ok := d.GetOk("mep.0.default_priority_rule.0.third_priority_center_gateways"); ok {
						defaultPriorityRulePayload["third-priority-center-gateways"] = v.(*schema.Set).List()
					}
					mepPayload["default-priority-rule"] = defaultPriorityRulePayload
				}
				if _, ok := d.GetOk("mep.0.exception_priority_rules"); ok {
					exceptionPriorityRulesPayload := make(map[string]interface{})

					if v, ok := d.GetOk("mep.0.exception_priority_rules.0.satellite_gateways"); ok {
						exceptionPriorityRulesPayload["satellite-gateways"] = v.(*schema.Set).List()
					}
					if v, ok := d.GetOk("mep.0.exception_priority_rules.0.first_priority_center_gateways"); ok {
						exceptionPriorityRulesPayload["first-priority-center-gateways"] = v.(*schema.Set).List()
					}
					if v, ok := d.GetOk("mep.0.exception_priority_rules.0.second_priority_center_gateways"); ok {
						exceptionPriorityRulesPayload["second-priority-center-gateways"] = v.(*schema.Set).List()
					}
					if v, ok := d.GetOk("mep.0.exception_priority_rules.0.third_priority_center_gateways"); ok {
						exceptionPriorityRulesPayload["third-priority-center-gateways"] = v.(*schema.Set).List()
					}
					mepPayload["exception-priority-rules"] = exceptionPriorityRulesPayload
				}
				vpnCommunityStar["mep"] = mepPayload
			}
		}
	}

	if v, ok := d.GetOkExists("mesh_center_gateways"); ok {
		vpnCommunityStar["mesh-center-gateways"] = v.(bool)
	}

	if d.HasChange("override_vpn_domains") {

		if v, ok := d.GetOk("override_vpn_domains"); ok {

			overrideVpnDomainsList := v.([]interface{})

			var overrideVpnDomainsPayload []map[string]interface{}

			for i := range overrideVpnDomainsList {

				Payload := make(map[string]interface{})

				if d.HasChange("override_vpn_domains." + strconv.Itoa(i) + ".gateway") {
					Payload["gateway"] = d.Get("override_vpn_domains." + strconv.Itoa(i) + ".gateway")
				}
				if d.HasChange("override_vpn_domains." + strconv.Itoa(i) + ".vpn_domain") {
					Payload["vpn-domain"] = d.Get("override_vpn_domains." + strconv.Itoa(i) + ".vpn_domain")
				}
				overrideVpnDomainsPayload = append(overrideVpnDomainsPayload, Payload)
			}
			vpnCommunityStar["override-vpn-domains"] = overrideVpnDomainsPayload
		} else {
			oldoverrideVpnDomains, _ := d.GetChange("override_vpn_domains")
			var overrideVpnDomainsToDelete []interface{}
			for _, i := range oldoverrideVpnDomains.([]interface{}) {
				overrideVpnDomainsToDelete = append(overrideVpnDomainsToDelete, i.(map[string]interface{})["name"].(string))
			}
			vpnCommunityStar["override-vpn-domains"] = map[string]interface{}{"remove": overrideVpnDomainsToDelete}
		}
	}

	if d.HasChange("satellite_gateways") {
		if v, ok := d.GetOk("satellite_gateways"); ok {
			vpnCommunityStar["satellite-gateways"] = v.(*schema.Set).List()
		} else {
			oldSatellite_Gateways, _ := d.GetChange("satellite_gateways")
			vpnCommunityStar["satellite-gateways"] = map[string]interface{}{"remove": oldSatellite_Gateways.(*schema.Set).List()}
		}
	}

	if d.HasChange("shared_secrets") {

		if v, ok := d.GetOk("shared_secrets"); ok {

			sharedSecretsList := v.([]interface{})

			var sharedSecretsPayload []map[string]interface{}

			for i := range sharedSecretsList {

				Payload := make(map[string]interface{})

				if d.HasChange("shared_secrets." + strconv.Itoa(i) + ".external_gateway") {
					Payload["external-gateway"] = d.Get("shared_secrets." + strconv.Itoa(i) + ".external_gateway")
				}
				if d.HasChange("shared_secrets." + strconv.Itoa(i) + ".shared_secret") {
					Payload["shared-secret"] = d.Get("shared_secrets." + strconv.Itoa(i) + ".shared_secret")
				}
				sharedSecretsPayload = append(sharedSecretsPayload, Payload)
			}
			vpnCommunityStar["shared-secrets"] = sharedSecretsPayload
		} else {
			oldsharedSecrets, _ := d.GetChange("shared_secrets")
			var sharedSecretsToDelete []interface{}
			for _, i := range oldsharedSecrets.([]interface{}) {
				sharedSecretsToDelete = append(sharedSecretsToDelete, i.(map[string]interface{})["name"].(string))
			}
			vpnCommunityStar["shared-secrets"] = map[string]interface{}{"remove": sharedSecretsToDelete}
		}
	}

	if d.HasChange("tunnel_granularity") {
		vpnCommunityStar["tunnel-granularity"] = d.Get("tunnel_granularity")
	}

	if d.HasChange("granular_encryptions") {
		if v, ok := d.GetOk("granular_encryptions"); ok {
			granularEncryptions := v.([]interface{})
			if len(granularEncryptions) > 0 {
				var granularEncryptionsPayload []map[string]interface{}

				for i := range granularEncryptions {
					payload := make(map[string]interface{})
					if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".internal_gateway"); ok {
						payload["internal-gateway"] = v.(string)
					}
					if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".external_gateway"); ok {
						payload["external-gateway"] = v.(string)
					}
					if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".encryption_method"); ok {
						payload["encryption-method"] = v.(string)
					}
					if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".encryption_suite"); ok {
						payload["encryption-suite"] = v.(string)
					}
					if _, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_1"); ok {
						ikePhase1Payload := make(map[string]interface{})
						if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_1.encryption_algorithm"); ok {
							ikePhase1Payload["encryption-algorithm"] = v.(string)
						}
						if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_1.data_integrity"); ok {
							ikePhase1Payload["data-integrity"] = v.(string)
						}
						if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_1.diffie_hellman_group"); ok {
							ikePhase1Payload["diffie-hellman-group"] = v.(string)
						}
						if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_1.ike_p1_rekey_time"); ok {
							ikePhase1Payload["ike-p1-rekey-time"] = v.(string)
						}
						payload["ike-phase-1"] = ikePhase1Payload
					}
					if _, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_2"); ok {
						ikePhase2Payload := make(map[string]interface{})
						if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_2.encryption_algorithm"); ok {
							ikePhase2Payload["encryption-algorithm"] = v.(string)
						}
						if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_2.data_integrity"); ok {
							ikePhase2Payload["data-integrity"] = v.(string)
						}
						if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_2.ike_p2_use_pfs"); ok {
							ikePhase2Payload["ike-p2-use-pfs"] = v.(string)
						}
						if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_2.ike_p2_pfs_dh_grp"); ok {
							ikePhase2Payload["ike-p2-pfs-dh-grp"] = v.(string)
						}
						if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_2.ike_p2_rekey_time"); ok {
							ikePhase2Payload["ike-p2-rekey-time"] = v.(string)
						}
						payload["ike-phase-2"] = ikePhase2Payload
					}
					granularEncryptionsPayload = append(granularEncryptionsPayload, payload)
				}
				vpnCommunityStar["granular-encryptions"] = granularEncryptionsPayload
			}
		} else {
			granularEncryptions, _ := d.GetChange("granular_encryptions")
			oldValues := granularEncryptions.([]interface{})
			if len(oldValues) > 0 {
				var toRemove []interface{}
				for _, v := range oldValues {
					obj := make(map[string]interface{})
					obj["internal-gateway"] = v.(map[string]interface{})["internal_gateway"].(string)
					obj["external-gateway"] = v.(map[string]interface{})["external_gateway"].(string)
					toRemove = append(toRemove, obj)
				}
				vpnCommunityStar["granular-encryptions"] = map[string]interface{}{"remove": toRemove}
			}
		}
	}

	if d.HasChange("permanent_tunnels") {

		if v, ok := d.GetOk("permanent_tunnels"); ok {

			permanentTunnelsList := v.([]interface{})

			if len(permanentTunnelsList) > 0 {

				permanentTunnelsPayload := make(map[string]interface{})

				if v, ok := d.GetOk("permanent_tunnels.0.set_permanent_tunnels"); ok {
					permanentTunnelsPayload["set-permanent-tunnels"] = v.(string)
				}
				if _, ok := d.GetOk("permanent_tunnels.0.gateways"); ok {
					gatewaysPayload := make(map[string]interface{})

					if v, ok := d.GetOk("permanent_tunnels.0.gateways.0.gateway"); ok {
						gatewaysPayload["gateway"] = v.(string)
					}
					if v, ok := d.GetOk("permanent_tunnels.0.gateways.0.track_options"); ok {
						gatewaysPayload["track-options"] = v.(string)
					}
					if v, ok := d.GetOk("permanent_tunnels.0.gateways.0.override_tunnel_down_track"); ok {
						gatewaysPayload["override-tunnel-down-track"] = v.(string)
					}
					if v, ok := d.GetOk("permanent_tunnels.0.gateways.0.override_tunnel_up_track"); ok {
						gatewaysPayload["override-tunnel-up-track"] = v.(string)
					}
					permanentTunnelsPayload["gateways"] = gatewaysPayload
				}
				if _, ok := d.GetOk("permanent_tunnels.0.tunnels"); ok {
					tunnelsPayload := make(map[string]interface{})

					if v, ok := d.GetOk("permanent_tunnels.0.tunnels.0.first_tunnel_endpoint"); ok {
						tunnelsPayload["first-tunnel-endpoint"] = v.(string)
					}
					if v, ok := d.GetOk("permanent_tunnels.0.tunnels.0.second_tunnel_endpoint"); ok {
						tunnelsPayload["second-tunnel-endpoint"] = v.(string)
					}
					if v, ok := d.GetOk("permanent_tunnels.0.tunnels.0.track_options"); ok {
						tunnelsPayload["track-options"] = v.(string)
					}
					if v, ok := d.GetOk("permanent_tunnels.0.tunnels.0.override_tunnel_down_track"); ok {
						tunnelsPayload["override-tunnel-down-track"] = v.(string)
					}
					if v, ok := d.GetOk("permanent_tunnels.0.tunnels.0.override_tunnel_up_track"); ok {
						tunnelsPayload["override-tunnel-up-track"] = v.(string)
					}
					permanentTunnelsPayload["tunnels"] = tunnelsPayload
				}
				if _, ok := d.GetOk("permanent_tunnels.0.rim"); ok {
					rimPayload := make(map[string]interface{})

					if v, ok := d.GetOk("permanent_tunnels.0.rim.0.enabled"); ok {
						rimPayload["enabled"] = strconv.FormatBool(v.(bool))
					}
					if v, ok := d.GetOk("permanent_tunnels.0.rim.0.enable_on_center_gateways"); ok {
						rimPayload["enable-on-center-gateways"] = strconv.FormatBool(v.(bool))
					}
					if v, ok := d.GetOk("permanent_tunnels.0.rim.0.enable_on_satellite_gateways"); ok {
						rimPayload["enable-on-satellite-gateways"] = strconv.FormatBool(v.(bool))
					}
					if v, ok := d.GetOk("permanent_tunnels.0.rim.0.route_injection_track"); ok {
						rimPayload["route-injection-track"] = v.(string)
					}
					permanentTunnelsPayload["rim"] = rimPayload
				}
				if v, ok := d.GetOk("permanent_tunnels.0.tunnel_down_track"); ok {
					permanentTunnelsPayload["tunnel-down-track"] = v.(string)
				}
				if v, ok := d.GetOk("permanent_tunnels.0.tunnel_up_track"); ok {
					permanentTunnelsPayload["tunnel-up-track"] = v.(string)
				}
				vpnCommunityStar["permanent-tunnels"] = permanentTunnelsPayload
			}
		}
	}

	if ok := d.HasChange("vpn_routing"); ok {
		vpnCommunityStar["vpn-routing"] = d.Get("vpn_routing")
	}

	if d.HasChange("wire_mode") {

		if v, ok := d.GetOk("wire_mode"); ok {

			wireModeList := v.([]interface{})

			if len(wireModeList) > 0 {

				wireModePayload := make(map[string]interface{})

				if v, ok := d.GetOkExists("wire_mode.0.allow_uninspected_encrypted_traffic"); ok {
					wireModePayload["allow-uninspected-encrypted-traffic"] = v.(bool)
				}
				if v, ok := d.GetOkExists("wire_mode.0.allow_uninspected_encrypted_routing"); ok {
					wireModePayload["allow-uninspected-encrypted-routing"] = v.(bool)
				}
				vpnCommunityStar["wire-mode"] = wireModePayload
			}
		}
	}

	if ok := d.HasChange("routing_mode"); ok {
		vpnCommunityStar["routing-mode"] = d.Get("routing_mode")
	}

	if d.HasChange("advanced_properties") {

		if v, ok := d.GetOk("advanced_properties"); ok {

			advancedPropertiesList := v.([]interface{})

			if len(advancedPropertiesList) > 0 {

				advancedPropertiesPayload := make(map[string]interface{})

				if v, ok := d.GetOk("advanced_properties.0.support_ip_compression"); ok {
					advancedPropertiesPayload["support-ip-compression"] = v.(bool)
				}
				if v, ok := d.GetOk("advanced_properties.0.use_aggressive_mode"); ok {
					advancedPropertiesPayload["use-aggressive-mode"] = v.(bool)
				}
				vpnCommunityStar["advanced-properties"] = advancedPropertiesPayload
			}
		}
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			vpnCommunityStar["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			vpnCommunityStar["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if v, ok := d.GetOkExists("use_shared_secret"); ok {
		vpnCommunityStar["use-shared-secret"] = v.(bool)
	}

	if ok := d.HasChange("color"); ok {
		vpnCommunityStar["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		vpnCommunityStar["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		vpnCommunityStar["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		vpnCommunityStar["ignore-errors"] = v.(bool)
	}

	log.Println("Update VpnCommunityStar - Map = ", vpnCommunityStar)

	updateVpnCommunityStarRes, err := client.ApiCall("set-vpn-community-star", vpnCommunityStar, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !updateVpnCommunityStarRes.Success {
		if updateVpnCommunityStarRes.ErrorMsg != "" {
			return fmt.Errorf(updateVpnCommunityStarRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementVpnCommunityStar(d, m)
}

func deleteManagementVpnCommunityStar(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	vpnCommunityStarPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		vpnCommunityStarPayload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		vpnCommunityStarPayload["ignore-errors"] = v.(bool)
	}
	log.Println("Delete VpnCommunityStar")

	deleteVpnCommunityStarRes, err := client.ApiCall("delete-vpn-community-star", vpnCommunityStarPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteVpnCommunityStarRes.Success {
		if deleteVpnCommunityStarRes.ErrorMsg != "" {
			return fmt.Errorf(deleteVpnCommunityStarRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
