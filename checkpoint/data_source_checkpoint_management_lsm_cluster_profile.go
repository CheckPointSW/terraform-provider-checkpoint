package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"math"
	"strconv"
)

func dataSourceManagementLsmClusterProfile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementLsmClusterProfileRead,
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
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object type.",
			},
			"anti_bot": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Anti-Bot blade enabled.",
			},
			"anti_virus": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Anti-Virus blade enabled.",
			},
			"application_control": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Application Control blade enabled.",
			},
			"application_control_and_url_filtering_settings": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Computed:    true,
				Description: "Gateway Application Control and URL filtering settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"global_settings_mode": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Whether to override global settings or not.",
						},
						"override_global_settings": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "override global settings object.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"fail_mode": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Fail mode - allow or block all requests.",
									},
									"website_categorization": {
										Type:        schema.TypeMap,
										Computed:    true,
										Description: "Website categorization object.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"mode": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Website categorization mode.",
												},
												"custom_mode": {
													Type:        schema.TypeMap,
													Computed:    true,
													Description: "Custom mode object.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"social_networking_widgets": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Social networking widgets mode.",
															},
															"url_filtering": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "URL filtering mode.",
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
					},
				},
			},
			"advanced_settings": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Computed:    true,
				Description: "N/A",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connection_persistence": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Handling established connections when installing a new policy.",
						},
						"sam": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Computed:    true,
							Description: "SAM.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"forward_to_other_sam_servers": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Forward SAM clients' requests to other SAM servers.",
									},
									"use_early_versions": {
										Type:        schema.TypeMap,
										Computed:    true,
										Description: "N/A",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enabled": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Use early versions compatibility mode",
												},
												"compatibility_mode": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Early versions compatibility mode.",
												},
											},
										},
									},
									"purge_sam_file": {
										Type:        schema.TypeMap,
										Computed:    true,
										Description: "Purge SAM File.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enabled": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Purge SAM File.",
												},
												"purge_when_size_reaches_to": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "Purge SAM File When it Reaches to.",
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
			"cluster_interfaces": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "N/A",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object name. Must be unique in the domain.",
						},
						"network_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "N/A",
						},
						"network_mask": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "N/A",
						},
						"network_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "N/A",
						},
						"topology": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "N/A",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"anti_spoofing": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "N/A",
									},
									"anti_spoofing_settings": {
										Type:        schema.TypeMap,
										Computed:    true,
										Description: "N/A",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"action": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "N/A",
												},
												"do_not_check_specific_packets": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "N/A",
												},
												"do_not_check_specific_packets_from": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "N/A",
												},
												"spoof_tracking": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "N/A",
												},
											},
										},
									},
									"interface_leads_to_dmz": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "N/A",
									},
									"ip_addresses_behind_this_interface": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "N/A",
									},
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "N/A",
									},
								},
							},
						},
					},
				},
			},
			"cluster_members": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "N/A",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object name. Must be unique in the domain.",
						},
						"uid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object unique identifier.",
						},
						"ip_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "N/A",
						},
						"comments": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Comments string.",
						},
					},
				},
			},
			"content_awareness": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Content Awareness blade enabled.",
			},
			"data_loss_prevention": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Data Loss Prevention.",
			},
			"dynamic_ip": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Dynamic IP address.",
			},
			"enable_https_inspection": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable HTTPS Inspection after defining an outbound inspection certificate.\nTo define the outbound certificate use \"set outbound-inspection-certificate\".",
			},
			"firewall": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Firewall blade enabled.",
			},
			"groups": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object name. Must be unique in the domain.",
						},
						"uid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object unique identifier.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object type.",
						},
						"color": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Color of the object. Should be one of existing colors.",
						},
					},
				},
			},
			"hit_count": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Hit count tracks the number of connections each rule matches.",
			},
			"https_inspection": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Computed:    true,
				Description: "HTTPS inspection.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bypass_on_failure": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Set to be true in order to bypass all requests (Fail-open) in case of internal system error.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"override_profile": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Override profile of global configuration.",
									},
									"profile_value": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Override profile value.",
									},
									"value": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Override value.",
									},
								},
							},
						},
						"site_categorization_allow_mode": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Set to 'background' in order to allowed requests until categorization is complete.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"override_profile": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Override profile of global configuration.",
									},
									"profile_value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Override profile value.",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Override value.",
									},
								},
							},
						},
						"deny_untrusted_server_cert": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Action settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"override_profile": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Override profile of global configuration.",
									},
									"profile_value": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Override profile value.",
									},
									"value": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Override value.",
									},
								},
							},
						},
						"deny_revoked_server_cert": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Action settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"override_profile": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Override profile of global configuration.",
									},
									"profile_value": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Override profile value.",
									},
									"value": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Override value.",
									},
								},
							},
						},
						"deny_expired_server_cert": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Action settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"override_profile": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Override profile of global configuration.",
									},
									"profile_value": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Override profile value.",
									},
									"value": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Override value.",
									},
								},
							},
						},
					},
				},
			},
			"interfaces": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Cluster network interfaces.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ips": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Intrusion Prevention System blade enabled.",
			},
			"nat_hide_internal_interfaces": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Hide internal networks behind the Gateway's external IP.",
			},
			"nat_settings": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "NAT settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_rule": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether to add automatic address translation rules.",
						},
						"hide_behind": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Hide behind method. This parameter is forbidden in case \"method\" parameter is \"static\".",
						},
						"install_on": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Which gateway should apply the NAT translation.",
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
						"method": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "NAT translation method.",
						},
					},
				},
			},
			"os_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Gateway platform operating system.",
			},
			"proxy_settings": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "N/A",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"use_custom_proxy": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Use custom proxy settings for this network object.",
						},
						"proxy_server": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "N/A",
						},
						"port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "N/A",
						},
					},
				},
			},
			"qos": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "QoS.",
			},
			"save_logs_locally": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Save logs locally on the gateway.",
			},
			"send_alerts_to_server": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Server(s) to send alerts to.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"send_logs_to_backup_server": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Backup server(s) to send logs to.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"send_logs_to_server": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Servers(s) to send logs to.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"tags": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of tag objects identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"threat_emulation": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Threat Emulation blade enabled.",
			},
			"threat_extraction": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Threat Extraction blade enabled.",
			},
			"threat_prevention_mode": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The mode of Threat Prevention to use. When using Autonomous Threat Prevention, disabling the Threat Prevention blades is not allowed.",
			},
			"url_filtering": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "URL Filtering blade enabled.",
			},
			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Gateway platform version.",
			},
			"vpn": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "VPN blade enabled.",
			},
			"zero_phishing": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Zero Phishing blade enabled.",
			},
			"zero_phishing_fqdn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Zero Phishing gateway FQDN.",
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

func dataSourceManagementLsmClusterProfileRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showLsmClusterProfileRes, err := client.ApiCall("show-lsm-cluster-profile", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showLsmClusterProfileRes.Success {
		if objectNotFound(showLsmClusterProfileRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showLsmClusterProfileRes.ErrorMsg)
	}

	lsmClusterProfile := showLsmClusterProfileRes.GetData()

	log.Println("Read LsmClusterProfile - Show JSON = ", lsmClusterProfile)

	if v := lsmClusterProfile["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := lsmClusterProfile["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := lsmClusterProfile["type"]; v != nil {
		_ = d.Set("type", v)
	}

	if v := lsmClusterProfile["anti-bot"]; v != nil {
		_ = d.Set("anti_bot", v)
	}

	if v := lsmClusterProfile["anti-virus"]; v != nil {
		_ = d.Set("anti_virus", v)
	}

	if v := lsmClusterProfile["application-control"]; v != nil {
		_ = d.Set("application_control", v)
	}

	if v := lsmClusterProfile["content-awareness"]; v != nil {
		_ = d.Set("content_awareness", v)
	}

	if v := lsmClusterProfile["data-loss-prevention"]; v != nil {
		_ = d.Set("data_loss_prevention", v)
	}

	if v := lsmClusterProfile["dynamic-ip"]; v != nil {
		_ = d.Set("dynamic_ip", v)
	}

	if v := lsmClusterProfile["enable-https-inspection"]; v != nil {
		_ = d.Set("enable_https_inspection", v)
	}

	if v := lsmClusterProfile["firewall"]; v != nil {
		_ = d.Set("firewall", v)
	}

	if v := lsmClusterProfile["hit-count"]; v != nil {
		_ = d.Set("hit_count", v)
	}

	if v := lsmClusterProfile["ips"]; v != nil {
		_ = d.Set("ips", v)
	}

	if v := lsmClusterProfile["nat-hide-internal-interfaces"]; v != nil {
		_ = d.Set("nat_hide_internal_interfaces", v)
	}

	if v := lsmClusterProfile["qos"]; v != nil {
		_ = d.Set("qos", v)
	}

	if v := lsmClusterProfile["save-logs-locally"]; v != nil {
		_ = d.Set("save_logs_locally", v)
	}

	if v := lsmClusterProfile["threat-emulation"]; v != nil {
		_ = d.Set("threat_emulation", v)
	}

	if v := lsmClusterProfile["threat-extraction"]; v != nil {
		_ = d.Set("threat_extraction", v)
	}

	if v := lsmClusterProfile["url-filtering"]; v != nil {
		_ = d.Set("url_filtering", v)
	}

	if v := lsmClusterProfile["vpn"]; v != nil {
		_ = d.Set("vpn", v)
	}

	if v := lsmClusterProfile["zero-phishing"]; v != nil {
		_ = d.Set("zero_phishing", v)
	}

	if v := lsmClusterProfile["threat-prevention-mode"]; v != nil {
		_ = d.Set("threat_prevention_mode", v)
	}

	if v := lsmClusterProfile["version"]; v != nil {
		_ = d.Set("version", v)
	}

	if v := lsmClusterProfile["zero-phishing-fqdn"]; v != nil {
		_ = d.Set("zero_phishing_fqdn", v)
	}

	if v := lsmClusterProfile["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := lsmClusterProfile["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if lsmClusterProfile["groups"] != nil {

		interfacesList, ok := lsmClusterProfile["groups"].([]interface{})

		var interfacesListToReturn []map[string]interface{}

		if ok {

			if len(interfacesList) > 0 {

				for i := range interfacesList {

					interfacesMap := interfacesList[i].(map[string]interface{})

					interfacesMapToAdd := make(map[string]interface{})

					if v, _ := interfacesMap["name"]; v != nil {
						interfacesMapToAdd["name"] = v
					}
					if v, _ := interfacesMap["uid"]; v != nil {
						interfacesMapToAdd["uid"] = v
					}
					if v, _ := interfacesMap["type"]; v != nil {
						interfacesMapToAdd["type"] = v
					}
					if v, _ := interfacesMap["color"]; v != nil {
						interfacesMapToAdd["color"] = v
					}
					interfacesListToReturn = append(interfacesListToReturn, interfacesMapToAdd)
				}
			}
		}
		_ = d.Set("groups", interfacesListToReturn)
	}

	if lsmClusterProfile["cluster-members"] != nil {

		interfacesList, ok := lsmClusterProfile["cluster-members"].([]interface{})

		var interfacesListToReturn []map[string]interface{}

		if ok {

			if len(interfacesList) > 0 {

				for i := range interfacesList {

					interfacesMap := interfacesList[i].(map[string]interface{})

					interfacesMapToAdd := make(map[string]interface{})

					if v, _ := interfacesMap["name"]; v != nil {
						interfacesMapToAdd["name"] = v
					}
					if v, _ := interfacesMap["uid"]; v != nil {
						interfacesMapToAdd["uid"] = v
					}
					if v, _ := interfacesMap["ip-address"]; v != nil {
						interfacesMapToAdd["ip_address"] = v
					}
					if v, _ := interfacesMap["comments"]; v != nil {
						interfacesMapToAdd["comments"] = v
					}
					interfacesListToReturn = append(interfacesListToReturn, interfacesMapToAdd)
				}
			}
		}
		_ = d.Set("cluster_members", interfacesListToReturn)
	}

	if lsmClusterProfile["cluster-interfaces"] != nil {

		interfacesList, ok := lsmClusterProfile["cluster-interfaces"].([]interface{})

		var interfacesListToReturn []map[string]interface{}

		if ok {

			if len(interfacesList) > 0 {

				for i := range interfacesList {

					interfacesMap := interfacesList[i].(map[string]interface{})

					interfacesMapToAdd := make(map[string]interface{})

					if v, _ := interfacesMap["name"]; v != nil {
						interfacesMapToAdd["name"] = v
					}
					if v, _ := interfacesMap["network-address"]; v != nil {
						interfacesMapToAdd["network_address"] = v
					}
					if v, _ := interfacesMap["network-mask"]; v != nil {
						interfacesMapToAdd["network_mask"] = v
					}
					if v, _ := interfacesMap["network-type"]; v != nil {
						interfacesMapToAdd["network_type"] = v
					}
					if topology, _ := interfacesMap["topology"]; topology != nil {
						overrideGlobalSettingsMap := topology.(map[string]interface{})
						topologyMapToAdd := make(map[string]interface{})

						if v, _ := overrideGlobalSettingsMap["anti-spoofing"]; v != nil {
							topologyMapToAdd["anti_spoofing"] = v
						}
						if v, _ := overrideGlobalSettingsMap["interface-leads-to-dmz"]; v != nil {
							topologyMapToAdd["interface_leads_to_dmz"] = v
						}
						if v, _ := overrideGlobalSettingsMap["ip-addresses-behind-this-interface"]; v != nil {
							topologyMapToAdd["ip_addresses_behind_this_interface"] = v
						}
						if v, _ := overrideGlobalSettingsMap["type"]; v != nil {
							topologyMapToAdd["type"] = v
						}
						if antiSpoofing, _ := overrideGlobalSettingsMap["anti-spoofing-settings"]; antiSpoofing != nil {
							antiSpoofingSettingsMap := antiSpoofing.(map[string]interface{})
							antiSpoofingMapToAdd := make(map[string]interface{})

							if v, _ := antiSpoofingSettingsMap["action"]; v != nil {
								antiSpoofingMapToAdd["action"] = v
							}
							if v, _ := antiSpoofingSettingsMap["do-not-check-specific-packets"]; v != nil {
								antiSpoofingMapToAdd["do_not_check_specific_packets"] = v
							}
							if v, _ := antiSpoofingSettingsMap["do-not-check-specific-packets-from"]; v != nil {
								antiSpoofingMapToAdd["do_not_check_specific_packets_from"] = v
							}
							if v, _ := antiSpoofingSettingsMap["spoof-tracking"]; v != nil {
								antiSpoofingMapToAdd["spoof_tracking"] = v
							}
							topologyMapToAdd["anti_spoofing_settings"] = antiSpoofingMapToAdd
						}
						interfacesMapToAdd["topology"] = topologyMapToAdd
					}
					interfacesListToReturn = append(interfacesListToReturn, interfacesMapToAdd)
				}
			}
		}
		_ = d.Set("cluster_interfaces", interfacesListToReturn)
	}

	if lsmClusterProfile["advanced-settings"] != nil {

		advancedSettingsMap := lsmClusterProfile["advanced-settings"].(map[string]interface{})

		advancedSettingsMapToReturn := make(map[string]interface{})

		if v, _ := advancedSettingsMap["connection-persistence"]; v != nil {
			advancedSettingsMapToReturn["connection_persistence"] = v
		}

		if sam, _ := advancedSettingsMap["sam"]; sam != nil {
			samSettingsMap := sam.(map[string]interface{})
			samMapToReturn := make(map[string]interface{})
			if v, _ := samSettingsMap["forward-to-other-sam-servers"]; v != nil {
				samMapToReturn["forward_to_other_sam_servers"] = v
			}

			if useEarlyVersion, _ := samSettingsMap["use-early-versions"]; useEarlyVersion != nil {
				useEarlyVersionMap := useEarlyVersion.(map[string]interface{})
				useEarlyVersionMapToReturn := make(map[string]interface{})

				if v, _ := useEarlyVersionMap["enabled"]; v != nil {
					useEarlyVersionMapToReturn["enabled"] = strconv.FormatBool(v.(bool))
				}

				if v, _ := useEarlyVersionMap["compatibility-mode"]; v != nil {
					useEarlyVersionMapToReturn["compatibility_mode"] = v
				}
				samMapToReturn["use_early_versions"] = useEarlyVersionMapToReturn
			}
			if purgeSamFile, _ := samSettingsMap["purge-sam-file"]; purgeSamFile != nil {
				purgeSamFileMap := purgeSamFile.(map[string]interface{})
				purgeSamFileMapToReturn := make(map[string]interface{})

				if v, _ := purgeSamFileMap["enabled"]; v != nil {
					purgeSamFileMapToReturn["enabled"] = strconv.FormatBool(v.(bool))
				}

				if v, _ := purgeSamFileMap["purge-when-size-reaches-to"]; v != nil {
					purgeSamFileMapToReturn["purge_when_size_reaches_to"] = strconv.Itoa(int(math.Round(v.(float64))))
				}
				samMapToReturn["purge_sam_file"] = purgeSamFileMapToReturn
			}
			advancedSettingsMapToReturn["sam"] = []interface{}{samMapToReturn}
		}

		_ = d.Set("advanced_settings", []interface{}{advancedSettingsMapToReturn})
	} else {
		_ = d.Set("advanced_settings", []interface{}{})
	}

	if lsmClusterProfile["application-control-and-url-filtering-settings"] != nil {

		applicationControlSettingsMap := lsmClusterProfile["application-control-and-url-filtering-settings"].(map[string]interface{})

		applicationControlSettingsMapToReturn := make(map[string]interface{})

		if v, _ := applicationControlSettingsMap["global-settings-mode"]; v != nil {
			applicationControlSettingsMapToReturn["global_settings_mode"] = v
		}

		if overrideGlobal, _ := applicationControlSettingsMap["override-global-settings"]; overrideGlobal != nil {
			overrideGlobalSettingsMap := overrideGlobal.(map[string]interface{})
			overrideGlobalMapToReturn := make(map[string]interface{})
			if v, _ := overrideGlobalSettingsMap["fail-mode"]; v != nil {
				overrideGlobalMapToReturn["fail_mode"] = v
			}

			if websiteCategorization, _ := overrideGlobalSettingsMap["website-categorization"]; websiteCategorization != nil {
				websiteCategorizationMap := websiteCategorization.(map[string]interface{})
				websiteCategorizationMapToReturn := make(map[string]interface{})

				if v, _ := websiteCategorizationMap["mode"]; v != nil {
					websiteCategorizationMapToReturn["mode"] = v
				}

				if customMode, _ := websiteCategorizationMap["custom-mode"]; customMode != nil {
					customModeMap := customMode.(map[string]interface{})
					customModeMapToReturn := make(map[string]interface{})
					if v, _ := customModeMap["social-networking-widgets"]; v != nil {
						customModeMapToReturn["social_networking_widgets"] = v
					}

					if v, _ := customModeMap["url-filtering"]; v != nil {
						customModeMapToReturn["url_filtering"] = v
					}
					websiteCategorizationMapToReturn["custom_mode"] = customModeMapToReturn
				}
				overrideGlobalMapToReturn["website_categorization"] = websiteCategorizationMapToReturn
			}
			applicationControlSettingsMapToReturn["override_global_settings"] = overrideGlobalMapToReturn
		}

		_ = d.Set("application_control_and_url_filtering_settings", []interface{}{applicationControlSettingsMapToReturn})
	} else {
		_ = d.Set("application_control_and_url_filtering_settings", []interface{}{})
	}

	if lsmClusterProfile["https-inspection"] != nil {

		actionSettingsMap := lsmClusterProfile["https-inspection"].(map[string]interface{})

		actionSettingsMapToReturn := make(map[string]interface{})

		if v, _ := actionSettingsMap["bypass-on-failure"]; v != nil {
			httpsSettingsMap := actionSettingsMap["bypass-on-failure"].(map[string]interface{})
			httpsMapToReturn := make(map[string]interface{})
			if v, _ := httpsSettingsMap["override-profile"]; v != nil {
				httpsMapToReturn["override_profile"] = strconv.FormatBool(v.(bool))
			}

			if v, _ := httpsSettingsMap["profile-value"]; v != nil {
				httpsMapToReturn["profile_value"] = strconv.FormatBool(v.(bool))
			}

			if v, _ := httpsSettingsMap["value"]; v != nil {
				httpsMapToReturn["value"] = strconv.FormatBool(v.(bool))
			}
			actionSettingsMapToReturn["bypass_on_failure"] = httpsMapToReturn
		}

		if v, _ := actionSettingsMap["site-categorization-allow-mode"]; v != nil {
			httpsSettingsMap := actionSettingsMap["site-categorization-allow-mode"].(map[string]interface{})
			httpsMapToReturn := make(map[string]interface{})
			if v, _ := httpsSettingsMap["override-profile"]; v != nil {
				httpsMapToReturn["override_profile"] = strconv.FormatBool(v.(bool))
			}

			if v, _ := httpsSettingsMap["profile-value"]; v != nil {
				httpsMapToReturn["profile_value"] = v
			}

			if v, _ := httpsSettingsMap["value"]; v != nil {
				httpsMapToReturn["value"] = v
			}
			actionSettingsMapToReturn["site_categorization_allow_mode"] = httpsMapToReturn
		}

		if v, _ := actionSettingsMap["deny-untrusted-server-cert"]; v != nil {
			httpsSettingsMap := actionSettingsMap["deny-untrusted-server-cert"].(map[string]interface{})
			httpsMapToReturn := make(map[string]interface{})
			if v, _ := httpsSettingsMap["override-profile"]; v != nil {
				httpsMapToReturn["override_profile"] = strconv.FormatBool(v.(bool))
			}

			if v, _ := httpsSettingsMap["profile-value"]; v != nil {
				httpsMapToReturn["profile_value"] = strconv.FormatBool(v.(bool))
			}

			if v, _ := httpsSettingsMap["value"]; v != nil {
				httpsMapToReturn["value"] = strconv.FormatBool(v.(bool))
			}
			actionSettingsMapToReturn["deny_untrusted_server_cert"] = httpsMapToReturn
		}

		if v, _ := actionSettingsMap["deny-revoked-server-cert"]; v != nil {
			httpsSettingsMap := actionSettingsMap["deny-revoked-server-cert"].(map[string]interface{})
			httpsMapToReturn := make(map[string]interface{})
			if v, _ := httpsSettingsMap["override-profile"]; v != nil {
				httpsMapToReturn["override_profile"] = strconv.FormatBool(v.(bool))
			}

			if v, _ := httpsSettingsMap["profile-value"]; v != nil {
				httpsMapToReturn["profile_value"] = strconv.FormatBool(v.(bool))
			}

			if v, _ := httpsSettingsMap["value"]; v != nil {
				httpsMapToReturn["value"] = strconv.FormatBool(v.(bool))
			}
			actionSettingsMapToReturn["deny_revoked_server_cert"] = httpsMapToReturn
		}

		if v, _ := actionSettingsMap["deny-expired-server-cert"]; v != nil {
			httpsSettingsMap := actionSettingsMap["deny-expired-server-cert"].(map[string]interface{})
			httpsMapToReturn := make(map[string]interface{})
			if v, _ := httpsSettingsMap["override-profile"]; v != nil {
				httpsMapToReturn["override_profile"] = strconv.FormatBool(v.(bool))
			}

			if v, _ := httpsSettingsMap["profile-value"]; v != nil {
				httpsMapToReturn["profile_value"] = strconv.FormatBool(v.(bool))
			}

			if v, _ := httpsSettingsMap["value"]; v != nil {
				httpsMapToReturn["value"] = strconv.FormatBool(v.(bool))
			}
			actionSettingsMapToReturn["deny_expired_server_cert"] = httpsMapToReturn
		}
		err = d.Set("https_inspection", []interface{}{actionSettingsMapToReturn})
	} else {
		_ = d.Set("https_inspection", nil)
	}

	if lsmClusterProfile["nat-settings"] != nil {

		actionSettingsMap := lsmClusterProfile["nat-settings"].(map[string]interface{})

		actionSettingsMapToReturn := make(map[string]interface{})

		if v, _ := actionSettingsMap["auto-rule"]; v != nil {
			actionSettingsMapToReturn["auto_rule"] = strconv.FormatBool(v.(bool))
		}

		if v, _ := actionSettingsMap["hide-behind"]; v != nil {
			actionSettingsMapToReturn["hide_behind"] = v
		}

		if v, _ := actionSettingsMap["install-on"]; v != nil {
			actionSettingsMapToReturn["install_on"] = v
		}

		if v, _ := actionSettingsMap["ipv4-address"]; v != nil {
			actionSettingsMapToReturn["ipv4_address"] = v
		}

		if v, _ := actionSettingsMap["ipv6-address"]; v != nil {
			actionSettingsMapToReturn["ipv6_address"] = v
		}

		if v, _ := actionSettingsMap["method"]; v != nil {
			actionSettingsMapToReturn["method"] = v
		}

		_ = d.Set("nat_settings", actionSettingsMapToReturn)
	} else {
		_ = d.Set("nat_settings", nil)
	}

	if lsmClusterProfile["proxy-settings"] != nil {

		actionSettingsMap := lsmClusterProfile["proxy-settings"].(map[string]interface{})

		actionSettingsMapToReturn := make(map[string]interface{})

		if v, _ := actionSettingsMap["use-custom-proxy"]; v != nil {
			actionSettingsMapToReturn["use_custom_proxy"] = strconv.FormatBool(v.(bool))
		}

		if v, _ := actionSettingsMap["proxy-server"]; v != nil {
			actionSettingsMapToReturn["proxy_server"] = v
		}

		if v, _ := actionSettingsMap["port"]; v != nil {
			actionSettingsMapToReturn["port"] = strconv.Itoa(int(math.Round(v.(float64))))
		}

		_ = d.Set("proxy_settings", actionSettingsMapToReturn)
	} else {
		_ = d.Set("proxy_settings", nil)
	}

	if lsmClusterProfile["interfaces"] != nil {
		serversJson, ok := lsmClusterProfile["interfaces"].([]interface{})
		if ok {
			serversIds := make([]string, 0)
			if len(serversJson) > 0 {
				for _, server := range serversJson {
					serversIds = append(serversIds, server.(string))
				}
			}
			_ = d.Set("interfaces", serversIds)
		}
	} else {
		_ = d.Set("interfaces", nil)
	}

	if lsmClusterProfile["send-alerts-to-server"] != nil {
		serversJson, ok := lsmClusterProfile["send-alerts-to-server"].([]interface{})
		if ok {
			serversIds := make([]string, 0)
			if len(serversJson) > 0 {
				for _, server := range serversJson {
					serversIds = append(serversIds, server.(string))
				}
			}
			_ = d.Set("send_alerts_to_server", serversIds)
		}
	} else {
		_ = d.Set("send_alerts_to_server", nil)
	}

	if lsmClusterProfile["send-logs-to-backup-server"] != nil {
		serversJson, ok := lsmClusterProfile["send-logs-to-backup-server"].([]interface{})
		if ok {
			serversIds := make([]string, 0)
			if len(serversJson) > 0 {
				for _, server := range serversJson {
					serversIds = append(serversIds, server.(string))
				}
			}
			_ = d.Set("send_logs_to_backup_server", serversIds)
		}
	} else {
		_ = d.Set("send_logs_to_backup_server", nil)
	}

	if lsmClusterProfile["send-logs-to-server"] != nil {
		serversJson, ok := lsmClusterProfile["send-logs-to-server"].([]interface{})
		if ok {
			serversIds := make([]string, 0)
			if len(serversJson) > 0 {
				for _, server := range serversJson {
					serversIds = append(serversIds, server.(string))
				}
			}
			_ = d.Set("send_logs_to_server", serversIds)
		}
	} else {
		_ = d.Set("send_logs_to_server", nil)
	}

	if lsmClusterProfile["tags"] != nil {
		tagsJson, ok := lsmClusterProfile["tags"].([]interface{})
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

	return nil

}
