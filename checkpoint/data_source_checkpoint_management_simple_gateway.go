package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementSimpleGateway() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementSimpleGatewayRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name. Should be unique in the domain.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object unique identifier.",
			},
			"ipv4_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IPv4 address.",
			},
			"ipv6_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IPv6 address.",
			},
			"interfaces": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Network interfaces.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object name. Should be unique in the domain.",
						},
						"ipv4_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv4 address.",
						},
						"ipv6_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv6 address.",
						},
						"ipv4_network_mask": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv4 network address.",
						},
						"ipv6_network_mask": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv6 network address.",
						},
						"ipv4_mask_length": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv4 network mask length.",
						},
						"ipv6_mask_length": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv6 network mask length.",
						},
						"anti_spoofing": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Anti spoofing.",
						},
						"anti_spoofing_settings": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Anti spoofing settings",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "If packets will be rejected (the Prevent option) or whether the packets will be monitored (the Detect option).",
									},
								},
							},
						},
						"security_zone": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Security zone.",
						},
						"security_zone_settings": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Security zone settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auto_calculated": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Security Zone is calculated according to where the interface leads to.",
									},
									"specific_zone": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Security Zone specified manually.",
									},
								},
							},
						},
						"topology": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Topology.",
						},
						"topology_settings": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Topology settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"interface_leads_to_dmz": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether this interface leads to demilitarized zone (perimeter network).",
									},
									"ip_address_behind_this_interface": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Ip address behind this interface.",
									},
									"specific_network": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Network behind this interface.",
									},
								},
							},
						},
						"topology_automatic_calculation": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Shows the automatic topology calculation.",
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
				},
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
			"content_awareness": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Content Awareness blade enabled.",
			},
			"firewall": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Firewall blade enabled.",
			},
			"firewall_settings": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Firewall settings",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_calculate_connections_hash_table_size_and_memory_pool": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Auto calculate connections hash table size and memory pool.",
						},
						"auto_maximum_limit_for_concurrent_connections": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Auto maximum limit for concurrent connections.",
						},
						"connections_hash_size": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Connections hash size.",
						},
						"maximum_limit_for_concurrent_connections": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Maximum limit for concurrent connections.",
						},
						"maximum_memory_pool_size": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Maximum memory pool size.",
						},
						"memory_pool_size": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Memory pool size.",
						},
					},
				},
			},
			"icap_server": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "ICAP Server enabled.",
			},
			"ips": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Intrusion Prevention System blade enabled.",
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
			"url_filtering": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "URL Filtering blade enabled.",
			},
			"dynamic_ip": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Dynamic IP address.",
			},
			"os_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "OS name.",
			},
			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Gateway platform version.",
			},
			"hardware": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Gateway platform hardware type.",
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
			"save_logs_locally": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Save logs locally.",
			},
			"send_alerts_to_server": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Server(s) to send alerts to.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"send_logs_to_backup_server": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Backup server(s) to send logs to.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"send_logs_to_server": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Server(s) to send logs to.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"logs_settings": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Logs settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"alert_when_free_disk_space_below": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable alert when free disk space is below threshold.",
						},
						"alert_when_free_disk_space_below_metrics": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Alert when free disk space below metrics.",
						},
						"alert_when_free_disk_space_below_threshold": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Alert when free disk space below threshold.",
						},
						"alert_when_free_disk_space_below_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Alert when free disk space below type.",
							Default:     "popup alert",
						},
						"before_delete_keep_logs_from_the_last_days": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable before delete keep logs from the last days.",
						},
						"before_delete_keep_logs_from_the_last_days_threshold": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Before delete keep logs from the last days threshold.",
						},
						"before_delete_run_script": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable Before delete run script.",
						},
						"before_delete_run_script_command": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Before delete run script command.",
						},
						"delete_index_files_older_than_days": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable delete index files older than days.",
						},
						"delete_index_files_older_than_days_threshold": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Delete index files older than days threshold.",
						},
						"delete_index_files_when_index_size_above": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable delete index files when index size above.",
						},
						"delete_index_files_when_index_size_above_metrics": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Delete index files when index size above metrics.",
						},
						"delete_index_files_when_index_size_above_threshold": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Delete index files when index size above threshold.",
						},
						"delete_when_free_disk_space_below": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable delete when free disk space below.",
						},
						"delete_when_free_disk_space_below_metrics": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Delete when free disk space below metric.",
						},
						"delete_when_free_disk_space_below_threshold": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Delete when free disk space below threshold.",
						},
						"detect_new_citrix_ica_application_names": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable detect new citrix ica application names.",
						},
						"forward_logs_to_log_server": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable forward logs to log server.",
						},
						"forward_logs_to_log_server_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Forward logs to log server name.",
						},
						"forward_logs_to_log_server_schedule_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Forward logs to log server schedule name.",
						},
						"free_disk_space_metrics": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Free disk space metrics.",
						},
						"perform_log_rotate_before_log_forwarding": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable perform log rotate before log forwarding.",
						},
						"reject_connections_when_free_disk_space_below_threshold": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable reject connections when free disk space below threshold.",
						},
						"reserve_for_packet_capture_metrics": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Reserve for packet capture metrics.",
						},
						"reserve_for_packet_capture_threshold": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Reserve for packet capture threshold.",
						},
						"rotate_log_by_file_size": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable rotate log by file size.",
						},
						"rotate_log_file_size_threshold": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Log file size threshold.",
						},
						"rotate_log_on_schedule": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable rotate log on schedule.",
						},
						"rotate_log_schedule_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Rotate log schedule name.",
						},
						"stop_logging_when_free_disk_space_below": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable stop logging when free disk space below.",
						},
						"stop_logging_when_free_disk_space_below_threshold": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Stop logging when free disk space below threshold.",
						},
						"turn_on_qos_logging": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable turn on qos logging.",
						},
						"update_account_log_every": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Update account log in every amount of seconds.",
						},
					},
				},
			},
			"vpn": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "VPN blade enabled.",
			},
			"vpn_settings": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Gateway VPN settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Authentication.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"authentication_clients": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "Collection of VPN Authentication clients identified by the name or UID.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"link_selection": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Link Selection.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_selection": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "IP selection",
									},
									"dns_resolving_hostname": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "DNS Resolving Hostname. Must be set when \"ip-selection\" was selected to be \"dns-resolving-from-hostname\".",
									},
									"ip_address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "IP Address. Must be set when \"ip-selection\" was selected to be \"use-selected-address-from-topology\" or \"use-statically-nated-ip\"",
									},
								},
							},
						},
						"maximum_concurrent_ike_negotiations": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Maximum concurrent ike negotiations",
						},
						"maximum_concurrent_tunnels": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Maximum concurrent tunnels",
						},
						"office_mode": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Office Mode. Notation Wide Impact - Office Mode apply IPSec VPN Software Blade clients and to the Mobile Access Software Blade clients.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"mode": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Office Mode Permissions. When selected to be \"off\", all the other definitions are irrelevant.",
									},
									"group": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Group. Identified by name or UID. Must be set when \"office-mode-permissions\" was selected to be \"group\".",
									},
									"allocate_ip_address_from": {
										Type:        schema.TypeMap,
										Computed:    true,
										Description: "Allocate IP address Method. Allocate IP address by sequentially trying the given methods until success.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"radius_server": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Radius server used to authenticate the user.",
												},
												"use_allocate_method": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Use Allocate Method.",
												},
												"allocate_method": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Using either Manual (IP Pool) or Automatic (DHCP). Must be set when \"use-allocate-method\" is true.",
												},
												"manual_network": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Manual Network. Identified by name or UID. Must be set when \"allocate-method\" was selected to be \"manual\".",
												},
												"dhcp_server": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "DHCP Server. Identified by name or UID. Must be set when \"allocate-method\" was selected to be \"automatic\".",
												},
												"virtual_ip_address": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Virtual IPV4 address for DHCP server replies. Must be set when \"allocate-method\" was selected to be \"automatic\".",
												},
												"dhcp_mac_address": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Calculated MAC address for DHCP allocation. Must be set when \"allocate-method\" was selected to be \"automatic\".",
												},
												"optional_parameters": {
													Type:        schema.TypeMap,
													Computed:    true,
													Description: "This configuration applies to all Office Mode methods except Automatic (using DHCP) and ipassignment.conf entries which contain this data.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"use_primary_dns_server": {
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Use Primary DNS Server.",
															},
															"primary_dns_server": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Primary DNS Server. Identified by name or UID. Must be set when \"use-primary-dns-server\" is true and can not be set when \"use-primary-dns-server\" is false.",
															},
															"use_first_backup_dns_server": {
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Use First Backup DNS Server.",
															},
															"first_backup_dns_server": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "First Backup DNS Server. Identified by name or UID. Must be set when \"use-first-backup-dns-server\" is true and can not be set when \"use-first-backup-dns-server\" is false.",
															},
															"use_second_backup_dns_server": {
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Use Second Backup DNS Server.",
															},
															"second_backup_dns_server": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Second Backup DNS Server. Identified by name or UID. Must be set when \"use-second-backup-dns-server\" is true and can not be set when \"use-second-backup-dns-server\" is false.",
															},
															"dns_suffixes": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "DNS Suffixes.",
															},
															"use_primary_wins_server": {
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Use Primary WINS Server.",
															},
															"primary_wins_server": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Primary WINS Server. Identified by name or UID. Must be set when \"use-primary-wins-server\" is true and can not be set when \"use-primary-wins-server\" is false.",
															},
															"use_first_backup_wins_server": {
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Use First Backup WINS Server.",
															},
															"first_backup_wins_server": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "First Backup WINS Server. Identified by name or UID. Must be set when \"use-first-backup-wins-server\" is true and can not be set when \"use-first-backup-wins-server\" is false.",
															},
															"use_second_backup_wins_server": {
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Use Second Backup WINS Server.",
															},
															"second_backup_wins_server": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Second Backup WINS Server. Identified by name or UID. Must be set when \"use-second-backup-wins-server\" is true and can not be set when \"use-second-backup-wins-server\" is false.",
															},
															"ip_lease_duration": {
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "IP Lease Duration in Minutes. The value must be in the range 2-32767.",
															},
														},
													},
												},
											},
										},
									},
									"support_multiple_interfaces": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Support connectivity enhancement for gateways with multiple external interfaces.",
									},
									"perform_anti_spoofing": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Perform Anti-Spoofing on Office Mode addresses.",
									},
									"anti_spoofing_additional_addresses": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Additional IP Addresses for Anti-Spoofing. Identified by name or UID. Must be set when \"perform-anti-spoofings\" is true.",
									},
								},
							},
						},
						"remote_access": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Remote Access.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"support_l2tp": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Support L2TP (relevant only when office mode is active).",
									},
									"l2tp_auth_method": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "L2TP Authentication Method. Must be set when \"support-l2tp\" is true.",
									},
									"l2tp_certificate": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "L2TP Certificate. Must be set when \"l2tp-auth-method\" was selected to be \"certificate\". Insert \"defaultCert\" when you want to use the default certificate.",
									},
									"allow_vpn_clients_to_route_traffic": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Allow VPN clients to route traffic.",
									},
									"support_nat_traversal_mechanism": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Support NAT traversal mechanism (UDP encapsulation).",
									},
									"nat_traversal_service": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Allocated NAT traversal UDP service. Identified by name or UID. Must be set when \"support-nat-traversal-mechanism\" is true.",
									},
									"support_visitor_mode": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Support Visitor Mode.",
									},
									"visitor_mode_service": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "TCP Service for Visitor Mode. Identified by name or UID. Must be set when \"support-visitor-mode\" is true.",
									},
									"visitor_mode_interface": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Interface for Visitor Mode. Must be set when \"support-visitor-mode\" is true. Insert IPV4 Address of existing interface or \"All IPs\" when you want all interfaces.",
									},
								},
							},
						},
						"vpn_domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Gateway VPN domain identified by the name or UID.",
						},
						"vpn_domain_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Gateway VPN domain type.",
						},
					},
				},
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
			"tags": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceManagementSimpleGatewayRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showGatewayRes, err := client.ApiCall("show-simple-gateway", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showGatewayRes.Success {
		return fmt.Errorf(showGatewayRes.ErrorMsg)
	}

	gateway := showGatewayRes.GetData()

	log.Println("Read Simple Gateway - Show JSON = ", gateway)

	if v := gateway["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := gateway["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := gateway["ipv4-address"]; v != nil {
		_ = d.Set("ipv4_address", v)
	}

	if v := gateway["ipv6-address"]; v != nil {
		_ = d.Set("ipv6_address", v)
	}

	if v := gateway["interfaces"]; v != nil {
		interfacesList := v.([]interface{})
		if len(interfacesList) > 0 {
			var interfacesListState []map[string]interface{}
			for i := range interfacesList {
				interfaceJson := interfacesList[i].(map[string]interface{})
				interfaceState := make(map[string]interface{})
				if v, _ := interfaceJson["name"]; v != nil {
					interfaceState["name"] = v
				}
				if v, _ := interfaceJson["ipv4-address"]; v != nil {
					interfaceState["ipv4_address"] = v
				}
				if v, _ := interfaceJson["ipv4-mask-length"]; v != nil {
					interfaceState["ipv4_mask_length"] = v
				}
				if v, _ := interfaceJson["ipv4-network-mask"]; v != nil {
					interfaceState["ipv4_network_mask"] = v
				}
				if v, _ := interfaceJson["ipv6-address"]; v != nil {
					interfaceState["ipv6_address"] = v
				}
				if v, _ := interfaceJson["ipv6-mask-length"]; v != nil {
					interfaceState["ipv6_mask_length"] = v
				}
				if v, _ := interfaceJson["ipv6-network-mask"]; v != nil {
					interfaceState["ipv6_network_mask"] = v
				}
				if v, _ := interfaceJson["anti-spoofing"]; v != nil {
					interfaceState["anti_spoofing"] = v
				}
				if v, _ := interfaceJson["anti-spoofing-settings"]; v != nil {
					antiSpoofingSettingsJson := v.(map[string]interface{})
					antiSpoofingSettingsState := make(map[string]interface{})
					if v, _ := antiSpoofingSettingsJson["action"]; v != nil {
						antiSpoofingSettingsState["action"] = v
					}
					interfaceState["anti_spoofing_settings"] = antiSpoofingSettingsState
				}
				if v, _ := interfaceJson["security-zone"]; v != nil {
					interfaceState["security_zone"] = v
				}
				if v, _ := interfaceJson["security-zone-settings"]; v != nil {
					securityZoneSettingsJson := v.(map[string]interface{})
					securityZoneSettingsState := make(map[string]interface{})
					if v, _ := securityZoneSettingsJson["auto-calculated"]; v != nil {
						securityZoneSettingsState["auto_calculated"] = v
					}
					if v, _ := securityZoneSettingsJson["specific-zone"]; v != nil {
						securityZoneSettingsState["specific_zone"] = v
					}
					interfaceState["security_zone_settings"] = securityZoneSettingsState
				}
				if v, _ := interfaceJson["topology"]; v != nil {
					interfaceState["topology"] = v
				}
				if v, _ := interfaceJson["topology-automatic-calculation"]; v != nil {
					interfaceState["topology_automatic_calculation"] = v
				}
				if v, _ := interfaceJson["topology-settings"]; v != nil {
					topologySettingsJson := v.(map[string]interface{})
					topologySettingsState := make(map[string]interface{})
					if v, _ := topologySettingsJson["interface-leads-to-dmz"]; v != nil {
						topologySettingsState["interface_leads_to_dmz"] = v
					}
					if v, _ := topologySettingsJson["ip-address-behind-this-interface"]; v != nil {
						topologySettingsState["ip_address_behind_this_interface"] = v
					}
					if v, _ := topologySettingsJson["specific-network"]; v != nil {
						topologySettingsState["specific_network"] = v
					}
					interfaceState["topology_settings"] = topologySettingsState
				}
				if v, _ := interfaceJson["color"]; v != nil {
					interfaceState["color"] = v
				}
				if v, _ := interfaceJson["comments"]; v != nil {
					interfaceState["comments"] = v
				}
				interfacesListState = append(interfacesListState, interfaceState)
			}
			_ = d.Set("interfaces", interfacesListState)
		} else {
			_ = d.Set("interfaces", interfacesList)
		}
	} else {
		_ = d.Set("interfaces", nil)
	}

	if v := gateway["anti-bot"]; v != nil {
		_ = d.Set("anti_bot", v)
	}

	if v := gateway["anti-virus"]; v != nil {
		_ = d.Set("anti_virus", v)
	}

	if v := gateway["application-control"]; v != nil {
		_ = d.Set("application_control", v)
	}

	if v := gateway["content-awareness"]; v != nil {
		_ = d.Set("content_awareness", v)
	}

	if v := gateway["dynamic-ip"]; v != nil {
		_ = d.Set("dynamic_ip", v)
	}

	if v := gateway["firewall"]; v != nil {
		_ = d.Set("firewall", v)
	}

	if v := gateway["icap-server"]; v != nil {
		_ = d.Set("icap_server", v)
	}

	if v := gateway["ips"]; v != nil {
		_ = d.Set("ips", v)
	}

	if v := gateway["threat-emulation"]; v != nil {
		_ = d.Set("threat_emulation", v)
	}

	if v := gateway["threat-extraction"]; v != nil {
		_ = d.Set("threat_extraction", v)
	}

	if v := gateway["url-filtering"]; v != nil {
		_ = d.Set("url_filtering", v)
	}

	if v := gateway["vpn"]; v != nil {
		_ = d.Set("vpn", v)
	}

	if v := gateway["os-name"]; v != nil {
		_ = d.Set("os_name", v)
	}

	if v := gateway["version"]; v != nil {
		_ = d.Set("version", v)
	}

	if v := gateway["hardware"]; v != nil {
		_ = d.Set("hardware", v)
	}

	if v := gateway["sic-name"]; v != nil {
		_ = d.Set("sic_name", v)
	}

	if v := gateway["sic-state"]; v != nil {
		_ = d.Set("sic_state", v)
	}

	if v := gateway["save-logs-locally"]; v != nil {
		_ = d.Set("save_logs_locally", v)
	}

	if v := gateway["send_alerts_to_server"]; v != nil {
		_ = d.Set("send_alerts_to_server", v)
	} else {
		_ = d.Set("send_alerts_to_server", nil)
	}

	if v := gateway["send-logs-to-backup-server"]; v != nil {
		_ = d.Set("send_logs_to_backup_server", v)
	} else {
		_ = d.Set("send_logs_to_backup_server", nil)
	}

	if v := gateway["send-logs-to-server"]; v != nil {
		_ = d.Set("send_logs_to_server", v)
	} else {
		_ = d.Set("send_logs_to_server", nil)
	}

	if v := gateway["logs-settings"]; v != nil {
		logSettingsJson := v.(map[string]interface{})
		logSettingsState := make(map[string]interface{})
		if v := logSettingsJson["alert-when-free-disk-space-below"]; v != nil {
			logSettingsState["alert_when_free_disk_space_below"] = v
		}
		if v := logSettingsJson["alert-when-free-disk-space-below-metrics"]; v != nil {
			logSettingsState["alert_when_free_disk_space_below_metrics"] = v
		}
		if v := logSettingsJson["alert_when_free_disk_space_below_threshold"]; v != nil {
			logSettingsState["alert_when_free_disk_space_below_threshold"] = v
		}
		if v := logSettingsJson["alert-when-free-disk-space-below-type"]; v != nil {
			logSettingsState["alert_when_free_disk_space_below_type"] = v
		}
		if v := logSettingsJson["before-delete-keep-logs-from-the-last-days"]; v != nil {
			logSettingsState["before_delete_keep_logs_from_the_last_days"] = v
		}
		if v := logSettingsJson["before-delete-keep-logs-from-the-last-days-threshold"]; v != nil {
			logSettingsState["before_delete_keep_logs_from_the_last_days_threshold"] = v
		}
		if v := logSettingsJson["before-delete-run-script"]; v != nil {
			logSettingsState["before_delete_run_script"] = v
		}
		if v := logSettingsJson["before-delete-run-script-command"]; v != nil {
			logSettingsState["before_delete_run_script_command"] = v
		}
		if v := logSettingsJson["delete-index-files-older-than-days"]; v != nil {
			logSettingsState["delete_index_files_older_than_days"] = v
		}
		if v := logSettingsJson["delete-index-files-older-than-days-threshold"]; v != nil {
			logSettingsState["delete_index_files_older_than_days_threshold"] = v
		}
		if v := logSettingsJson["delete-index-files-when-index-size-above"]; v != nil {
			logSettingsState["delete_index_files_when_index_size_above"] = v
		}
		if v := logSettingsJson["delete-index-files-when-index-size-above-metrics"]; v != nil {
			logSettingsState["delete_index_files_when_index_size_above_metrics"] = v
		}
		if v := logSettingsJson["delete-index-files-when-index-size-above-threshold"]; v != nil {
			logSettingsState["delete_index_files_when_index_size_above_threshold"] = v
		}
		if v := logSettingsJson["delete-when-free-disk-space-below"]; v != nil {
			logSettingsState["delete_when_free_disk_space_below"] = v
		}
		if v := logSettingsJson["delete-when-free-disk-space-below-metrics"]; v != nil {
			logSettingsState["delete_when_free_disk_space_below_metrics"] = v
		}
		if v := logSettingsJson["delete-when-free-disk-space-below-threshold"]; v != nil {
			logSettingsState["delete_when_free_disk_space_below_threshold"] = v
		}
		if v := logSettingsJson["detect-new-citrix-ica-application-names"]; v != nil {
			logSettingsState["detect_new_citrix_ica_application_names"] = v
		}
		if v := logSettingsJson["forward-logs-to-log-server"]; v != nil {
			logSettingsState["forward_logs_to_log_server"] = v
		}
		if v := logSettingsJson["forward-logs-to-log-server-name"]; v != nil {
			logSettingsState["forward_logs_to_log_server_name"] = v
		}
		if v := logSettingsJson["forward-logs-to-log-server-schedule-name"]; v != nil {
			logSettingsState["forward_logs_to_log_server_schedule_name"] = v
		}
		if v := logSettingsJson["perform-log-rotate-before-log-forwarding"]; v != nil {
			logSettingsState["perform_log_rotate_before_log_forwarding"] = v
		}
		if v := logSettingsJson["reject-connections-when-free-disk-space-below-threshold"]; v != nil {
			logSettingsState["reject_connections_when_free_disk_space_below_threshold"] = v
		}
		if v := logSettingsJson["reserve-for-packet-capture-metrics"]; v != nil {
			logSettingsState["reserve_for_packet_capture_metrics"] = v
		}
		if v := logSettingsJson["reserve-for-packet-capture-threshold"]; v != nil {
			logSettingsState["reserve_for_packet_capture_threshold"] = v
		}
		if v := logSettingsJson["rotate-log-by-file-size"]; v != nil {
			logSettingsState["rotate_log_by_file_size"] = v
		}
		if v := logSettingsJson["rotate-log-file-size-threshold"]; v != nil {
			logSettingsState["rotate_log_file_size_threshold"] = v
		}
		if v := logSettingsJson["rotate-log-on-schedule"]; v != nil {
			logSettingsState["rotate_log_on_schedule"] = v
		}
		if v := logSettingsJson["rotate-log-schedule-name"]; v != nil {
			logSettingsState["rotate_log_schedule_name"] = v
		}
		if v := logSettingsJson["stop-logging-when-free-disk-space-below"]; v != nil {
			logSettingsState["stop_logging_when_free_disk_space_below"] = v
		}
		if v := logSettingsJson["stop-logging-when-free-disk-space-below-metrics"]; v != nil {
			logSettingsState["stop_logging_when_free_disk_space_below_metrics"] = v
		}
		if v := logSettingsJson["stop-logging-when-free-disk-space-below-threshold"]; v != nil {
			logSettingsState["stop_logging_when_free_disk_space_below_threshold"] = v
		}
		if v := logSettingsJson["turn-on-qos-logging"]; v != nil {
			logSettingsState["turn_on_qos_logging"] = v
		}
		if v := logSettingsJson["update-account-log-every"]; v != nil {
			logSettingsState["update_account_log_every"] = v
		}
		_ = d.Set("logs_settings", logSettingsState)
		/*
			_, logsSettingsInConf := d.GetOk("logs_settings")
			defaultLogsSettings := map[string]interface{}{
				"rotate_log_by_file_size": false,
				"rotate_log_file_size_threshold": 1000,
				"rotate_log_on_schedule": false,
				"alert_when_free_disk_space_below_metrics": "mbytes",
				"alert_when_free_disk_space_below": true,
				"alert_when_free_disk_space_below_threshold": 20,
				"alert_when_free_disk_space_below_type": "popup alert",
				"delete_when_free_disk_space_below_metrics": "mbytes",
				"delete_when_free_disk_space_below": true,
				"delete_when_free_disk_space_below_threshold": 5000,
				"before_delete_keep_logs_from_the_last_days": false,
				"before_delete_keep_logs_from_the_last_days_threshold": 3664,
				"before_delete_run_script": false,
				"before_delete_run_script_command": "",
				"stop_logging_when_free_disk_space_below_metrics": "mbytes",
				"stop_logging_when_free_disk_space_below": true,
				"stop_logging_when_free_disk_space_below_threshold": 100,
				"reject_connections_when_free_disk_space_below_threshold": false,
				"reserve_for_packet_capture_metrics": "mbytes",
				"reserve_for_packet_capture_threshold": 500,
				"delete_index_files_when_index_size_above_metrics": "mbytes",
				"delete_index_files_when_index_size_above": false,
				"delete_index_files_when_index_size_above_threshold": 100000,
				"delete_index_files_older_than_days": false,
				"delete_index_files_older_than_days_threshold": 14,
				"forward_logs_to_log_server": false,
				"perform_log_rotate_before_log_forwarding": false,
				"update_account_log_every": 3600,
				"detect_new_citrix_ica_application_names": false,
				"turn_on_qos_logging": true,
			}
			if reflect.DeepEqual(defaultLogsSettings, logSettingsState) && !logsSettingsInConf {
				log.Println("[royl] simple GW using default logs settings")
				_ = d.Set("logs_settings", map[string]interface{}{})
			} else {
				log.Println("[royl] simple GW set current logs settings")
				_ = d.Set("logs_settings", logSettingsState)
			}
		*/
	} else {
		_ = d.Set("logs_settings", nil)
	}

	if v := gateway["firewall-settings"]; v != nil {
		firewallSettingsJson := v.(map[string]interface{})
		firewallSettingsState := make(map[string]interface{})
		if v := firewallSettingsJson["auto-calculate-connections-hash-table-size-and-memory-pool"]; v != nil {
			firewallSettingsState["auto_calculate_connections_hash_table_size_and_memory_pool"] = v
		}
		if v := firewallSettingsJson["auto-maximum-limit-for-concurrent-connections"]; v != nil {
			firewallSettingsState["auto_maximum_limit_for_concurrent_connections"] = v
		}
		if v := firewallSettingsJson["connections-hash-size"]; v != nil {
			firewallSettingsState["connections_hash_size"] = v
		}
		if v := firewallSettingsJson["maximum-limit-for-concurrent-connections"]; v != nil {
			firewallSettingsState["maximum_limit_for_concurrent_connections"] = v
		}
		if v := firewallSettingsJson["maximum-memory-pool-size"]; v != nil {
			firewallSettingsState["maximum_memory_pool_size"] = v
		}
		if v := firewallSettingsJson["memory-pool-size"]; v != nil {
			firewallSettingsState["memory_pool_size"] = v
		}
		_ = d.Set("firewall_settings", firewallSettingsState)
	} else {
		_ = d.Set("firewall_settings", nil)
	}

	if v := gateway["vpn-settings"]; v != nil {
		vpnSettingsJson := v.(map[string]interface{})
		vpnSettingsState := make(map[string]interface{})
		if v := vpnSettingsJson["authentication"]; v != nil {
			authenticationJson := v.(map[string]interface{})
			authenticationState := make(map[string]interface{})
			if v := authenticationJson["authentication-clients"]; v != nil {
				clientsJson := v.([]interface{})
				var clientsIds = make([]string, 0)
				if len(clientsJson) > 0 {
					for _, client := range clientsJson {
						clientsIds = append(clientsIds, client.(map[string]interface{})["name"].(string))
					}
				}
				authenticationState["authentication_clients"] = clientsIds
			}
			vpnSettingsState["authentication"] = authenticationState
		}

		if v := vpnSettingsJson["link-selection"]; v != nil {
			linkSelectionJson := v.(map[string]interface{})
			linkSelectionState := make(map[string]interface{})
			if v := linkSelectionJson["ip-selection"]; v != nil {
				linkSelectionState["ip_selection"] = v
			}
			if v := linkSelectionJson["dns-resolving-hostname"]; v != nil {
				linkSelectionState["dns_resolving_hostname"] = v
			}
			if v := linkSelectionJson["ip-address"]; v != nil {
				linkSelectionState["ip_address"] = v
			}
			vpnSettingsState["link_selection"] = linkSelectionState
		}
		if v := vpnSettingsJson["maximum-concurrent-ike-negotiations"]; v != nil {
			vpnSettingsState["maximum_concurrent_ike_negotiations"] = v
		}
		if v := vpnSettingsJson["maximum-concurrent-tunnels"]; v != nil {
			vpnSettingsState["maximum_concurrent_tunnels"] = v
		}
		if v := vpnSettingsJson["vpn-domain-type"]; v != nil {
			vpnSettingsState["vpn_domain_type"] = v
		}
		if v := vpnSettingsJson["vpn-domain"]; v != nil {
			vpnSettingsState["vpn_domain"] = v.(map[string]interface{})["name"]
		}
		if v := vpnSettingsJson["remote-access"]; v != nil {
			remoteAccessJson := v.(map[string]interface{})
			remoteAccessState := make(map[string]interface{})
			if v := remoteAccessJson["support-l2tp"]; v != nil {
				remoteAccessState["support_l2tp"] = v
			}
			if v := remoteAccessJson["l2tp-auth-method"]; v != nil {
				remoteAccessState["l2tp_auth_method"] = v
			}
			if v := remoteAccessJson["l2tp-certificate"]; v != nil {
				remoteAccessState["l2tp_certificate"] = v
			}
			if v := remoteAccessJson["allow-vpn-clients-to-route-traffic"]; v != nil {
				remoteAccessState["allow_vpn_clients_to_route_traffic"] = v
			}
			if v := remoteAccessJson["support-nat-traversal-mechanism"]; v != nil {
				remoteAccessState["support_nat_traversal_mechanism"] = v
			}
			if v := remoteAccessJson["nat-traversal-service"]; v != nil {
				remoteAccessState["nat_traversal_service"] = v.(map[string]interface{})["name"]
			}
			if v := remoteAccessJson["support-visitor-mode"]; v != nil {
				remoteAccessState["support_visitor_mode"] = v
			}
			if v := remoteAccessJson["visitor-mode-service"]; v != nil {
				remoteAccessState["visitor_mode_service"] = v.(map[string]interface{})["name"]
			}
			if v := remoteAccessJson["visitor-mode-interface"]; v != nil {
				remoteAccessState["visitor_mode_interface"] = v
			}
			vpnSettingsState["remote_access"] = remoteAccessState
		}

		if v := vpnSettingsJson["office-mode"]; v != nil {
			officeModeJson := v.(map[string]interface{})
			officeModeState := make(map[string]interface{})
			if v := officeModeJson["mode"]; v != nil {
				officeModeState["mode"] = v
			}
			if v := officeModeJson["group"]; v != nil {
				officeModeState["group"] = v.(map[string]interface{})["name"]
			}
			if v := officeModeJson["support-multiple-interfaces"]; v != nil {
				officeModeState["support_multiple_interfaces"] = v
			}
			if v := officeModeJson["perform-anti-spoofing"]; v != nil {
				officeModeState["perform_anti_spoofing"] = v
			}
			if v := officeModeJson["anti-spoofing-additional-addresses"]; v != nil {
				officeModeState["anti_spoofing_additional_addresses"] = v.(map[string]interface{})["name"]
			}
			if v := officeModeJson["allocate-ip-address-from"]; v != nil {
				allocateIpAddressFromJson := v.(map[string]interface{})
				allocateIpAddressFromState := make(map[string]interface{})
				if v := allocateIpAddressFromJson["radius-server"]; v != nil {
					allocateIpAddressFromState["radius_server"] = v
				}
				if v := allocateIpAddressFromJson["use-allocate-method"]; v != nil {
					allocateIpAddressFromState["use_allocate_method"] = v
				}
				if v := allocateIpAddressFromJson["allocate-method"]; v != nil {
					allocateIpAddressFromState["allocate_method"] = v
				}
				if v := allocateIpAddressFromJson["manual-network"]; v != nil {
					allocateIpAddressFromState["manual_network"] = v.(map[string]interface{})["name"]
				}
				if v := allocateIpAddressFromJson["dhcp-server"]; v != nil {
					allocateIpAddressFromState["dhcp_server"] = v.(map[string]interface{})["name"]
				}
				if v := allocateIpAddressFromJson["virtual-ip-address"]; v != nil {
					allocateIpAddressFromState["virtual_ip_address"] = v
				}
				if v := allocateIpAddressFromJson["dhcp-mac-address"]; v != nil {
					allocateIpAddressFromState["dhcp_mac_address"] = v
				}
				if v := allocateIpAddressFromJson["optional-parameters"]; v != nil {
					optionalParametersJson := v.(map[string]interface{})
					optionalParametersState := make(map[string]interface{})
					if v := optionalParametersJson["use-primary-dns-server"]; v != nil {
						optionalParametersState["use_primary_dns_server"] = v
					}
					if v := optionalParametersJson["primary-dns-server"]; v != nil {
						optionalParametersState["primary-dns-server"] = v.(map[string]interface{})["name"]
					}
					if v := optionalParametersJson["use-first-backup-dns-server"]; v != nil {
						optionalParametersState["use_first_backup_dns_server"] = v
					}
					if v := optionalParametersJson["first-backup-dns-server"]; v != nil {
						optionalParametersState["first_backup_dns_server"] = v.(map[string]interface{})["name"]
					}
					if v := optionalParametersJson["use-second-backup-dns-server"]; v != nil {
						optionalParametersState["use_second_backup_dns_server"] = v
					}
					if v := optionalParametersJson["second-backup-dns-server"]; v != nil {
						optionalParametersState["second_backup_dns_server"] = v.(map[string]interface{})["name"]
					}
					if v := optionalParametersJson["dns-suffixes"]; v != nil {
						optionalParametersState["dns_suffixes"] = v
					}
					if v := optionalParametersJson["use-primary-wins-server"]; v != nil {
						optionalParametersState["use_primary_wins_server"] = v
					}
					if v := optionalParametersJson["primary-wins-server"]; v != nil {
						optionalParametersState["primary_wins_server"] = v.(map[string]interface{})["name"]
					}
					if v := optionalParametersJson["use-first-backup-wins-server"]; v != nil {
						optionalParametersState["use_first_backup_wins_server"] = v
					}
					if v := optionalParametersJson["first-backup-wins-server"]; v != nil {
						optionalParametersState["first_backup_wins_server"] = v.(map[string]interface{})["name"]
					}
					if v := optionalParametersJson["use-second-backup-wins-server"]; v != nil {
						optionalParametersState["use_second_backup_wins_server"] = v
					}
					if v := optionalParametersJson["second-backup-wins-server"]; v != nil {
						optionalParametersState["second_backup_wins_server"] = v.(map[string]interface{})["name"]
					}
					if v := optionalParametersJson["ip-lease-duration"]; v != nil {
						optionalParametersState["ip_lease_duration"] = v
					}
					allocateIpAddressFromState["optional_parameters"] = optionalParametersState
				}
				officeModeState["allocate_ip_address_from"] = allocateIpAddressFromState
			}
			vpnSettingsState["office_mode"] = officeModeState
		}
		_ = d.Set("vpn-settings", vpnSettingsState)
	} else {
		_ = d.Set("vpn-settings", nil)
	}

	if v := gateway["tags"]; v != nil {
		tagsJson := v.([]interface{})
		var tagsIds = make([]string, 0)
		if len(tagsJson) > 0 {
			for _, tag := range tagsJson {
				tagsIds = append(tagsIds, tag.(map[string]interface{})["name"].(string))
			}
		}
		_ = d.Set("tags", tagsIds)
	} else {
		_ = d.Set("tags", nil)
	}

	if v := gateway["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := gateway["color"]; v != nil {
		_ = d.Set("color", v)
	}

	return nil
}
