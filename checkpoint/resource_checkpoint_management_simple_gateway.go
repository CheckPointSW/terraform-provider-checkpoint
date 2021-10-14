package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"reflect"
	"strconv"
)

func resourceManagementSimpleGateway() *schema.Resource {
	return &schema.Resource{
		Create: createManagementSimpleGateway,
		Read:   readManagementSimpleGateway,
		Update: updateManagementSimpleGateway,
		Delete: deleteManagementSimpleGateway,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name. Should be unique in the domain.",
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
			"interfaces": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Network interfaces.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Object name. Should be unique in the domain.",
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
						"ipv4_network_mask": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IPv4 network address.",
						},
						"ipv6_network_mask": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IPv6 network address.",
						},
						"ipv4_mask_length": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IPv4 network mask length.",
						},
						"ipv6_mask_length": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IPv6 network mask length.",
						},
						"anti_spoofing": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Anti spoofing.",
							Default:     true,
						},
						"anti_spoofing_settings": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Anti spoofing settings",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "If packets will be rejected (the Prevent option) or whether the packets will be monitored (the Detect option).",
									},
								},
							},
						},
						"security_zone": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Security zone.",
							Default:     false,
						},
						"security_zone_settings": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Security zone settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auto_calculated": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Security Zone is calculated according to where the interface leads to.",
									},
									"specific_zone": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Security Zone specified manually.",
									},
								},
							},
						},
						"topology": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Topology.",
							Default:     "automatic",
						},
						"topology_settings": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Topology settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"interface_leads_to_dmz": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether this interface leads to demilitarized zone (perimeter network).",
									},
									"ip_address_behind_this_interface": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Ip address behind this interface.",
									},
									"specific_network": {
										Type:        schema.TypeString,
										Optional:    true,
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
							Optional:    true,
							Default:     "black",
							Description: "Color of the object. Should be one of existing colors.",
						},
						"comments": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Comments string.",
						},
					},
				},
			},
			"anti_bot": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Anti-Bot blade enabled.",
			},
			"anti_virus": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Anti-Virus blade enabled.",
			},
			"application_control": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Application Control blade enabled.",
			},
			"content_awareness": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Content Awareness blade enabled.",
			},
			"firewall": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Firewall blade enabled.",
				Default:     true,
			},
			"firewall_settings": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Firewall settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_calculate_connections_hash_table_size_and_memory_pool": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Auto calculate connections hash table size and memory pool.",
						},
						"auto_maximum_limit_for_concurrent_connections": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Auto maximum limit for concurrent connections.",
						},
						"connections_hash_size": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Connections hash size.",
						},
						"maximum_limit_for_concurrent_connections": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Maximum limit for concurrent connections.",
						},
						"maximum_memory_pool_size": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Maximum memory pool size.",
						},
						"memory_pool_size": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Memory pool size.",
						},
					},
				},
			},
			"icap_server": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "ICAP Server enabled.",
			},
			"ips": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Intrusion Prevention System blade enabled.",
			},
			"threat_emulation": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Threat Emulation blade enabled.",
			},
			"threat_extraction": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Threat Extraction blade enabled.",
			},
			"url_filtering": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "URL Filtering blade enabled.",
			},
			"dynamic_ip": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Dynamic IP address.",
			},
			"os_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "OS name.",
				Default:     "Gaia",
			},
			"version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Gateway platform version.",
			},
			"hardware": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Gateway platform hardware type.",
			},
			"one_time_password": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "SIC one time password.",
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
				Optional:    true,
				Description: "Save logs locally.",
			},
			"send_alerts_to_server": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Server(s) to send alerts to.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"send_logs_to_backup_server": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Backup server(s) to send logs to.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"send_logs_to_server": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Server(s) to send logs to.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"logs_settings": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Logs settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"alert_when_free_disk_space_below": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable alert when free disk space is below threshold.",
						},
						"free_disk_space_metrics": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Free disk space metrics.",
						},
						"delete_index_files_when_index_size_above_metrics": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Delete index files when index size above metrics",
						},
						"delete_when_free_disk_space_below_metrics": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Delete when free disk space below metric.",
						},
						"stop_logging_when_free_disk_space_below_metrics": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Stop logging when free disk space below metrics",
						},
						"alert_when_free_disk_space_below_threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Alert when free disk space below threshold.",
						},
						"alert_when_free_disk_space_below_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Alert when free disk space below type.",
							Default:     "popup alert",
						},
						"before_delete_keep_logs_from_the_last_days": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable before delete keep logs from the last days.",
						},
						"before_delete_keep_logs_from_the_last_days_threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Before delete keep logs from the last days threshold.",
						},
						"before_delete_run_script": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable Before delete run script.",
						},
						"before_delete_run_script_command": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Before delete run script command.",
						},
						"delete_index_files_older_than_days": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable delete index files older than days.",
						},
						"delete_index_files_older_than_days_threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Delete index files older than days threshold.",
						},
						"delete_index_files_when_index_size_above": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable delete index files when index size above.",
						},
						"delete_index_files_when_index_size_above_threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Delete index files when index size above threshold.",
						},
						"delete_when_free_disk_space_below": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable delete when free disk space below.",
						},
						"delete_when_free_disk_space_below_threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Delete when free disk space below threshold.",
						},
						"detect_new_citrix_ica_application_names": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable detect new citrix ica application names.",
						},
						"forward_logs_to_log_server": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable forward logs to log server.",
						},
						"forward_logs_to_log_server_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Forward logs to log server name.",
						},
						"forward_logs_to_log_server_schedule_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Forward logs to log server schedule name.",
						},
						"perform_log_rotate_before_log_forwarding": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable perform log rotate before log forwarding.",
						},
						"reject_connections_when_free_disk_space_below_threshold": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable reject connections when free disk space below threshold.",
						},
						"reserve_for_packet_capture_metrics": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Reserve for packet capture metrics.",
						},
						"reserve_for_packet_capture_threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Reserve for packet capture threshold.",
						},
						"rotate_log_by_file_size": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable rotate log by file size.",
						},
						"rotate_log_file_size_threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Log file size threshold.",
						},
						"rotate_log_on_schedule": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable rotate log on schedule.",
						},
						"rotate_log_schedule_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Rotate log schedule name.",
						},
						"stop_logging_when_free_disk_space_below": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable stop logging when free disk space below.",
						},
						"stop_logging_when_free_disk_space_below_threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Stop logging when free disk space below threshold.",
						},
						"turn_on_qos_logging": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable turn on qos logging.",
						},
						"update_account_log_every": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Update account log in every amount of seconds.",
						},
					},
				},
			},
			"vpn": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "VPN blade enabled.",
			},
			"vpn_settings": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Gateway VPN settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Authentication.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"authentication_clients": {
										Type:        schema.TypeSet,
										Optional:    true,
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
							Optional:    true,
							Description: "Link Selection.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_selection": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "IP selection",
										Default:     "use-main-address",
									},
									"dns_resolving_hostname": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "DNS Resolving Hostname. Must be set when \"ip-selection\" was selected to be \"dns-resolving-from-hostname\".",
									},
									"ip_address": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "IP Address. Must be set when \"ip-selection\" was selected to be \"use-selected-address-from-topology\" or \"use-statically-nated-ip\"",
									},
								},
							},
						},
						"maximum_concurrent_ike_negotiations": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Maximum concurrent ike negotiations",
						},
						"maximum_concurrent_tunnels": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Maximum concurrent tunnels",
						},
						"office_mode": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Office Mode. Notation Wide Impact - Office Mode apply IPSec VPN Software Blade clients and to the Mobile Access Software Blade clients.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"mode": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Office Mode Permissions. When selected to be \"off\", all the other definitions are irrelevant.",
										Default:     "off",
									},
									"group": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Group. Identified by name or UID. Must be set when \"office-mode-permissions\" was selected to be \"group\".",
									},
									"allocate_ip_address_from": {
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "Allocate IP address Method. Allocate IP address by sequentially trying the given methods until success.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"radius_server": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Radius server used to authenticate the user.",
													Default:     false,
												},
												"use_allocate_method": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Use Allocate Method.",
													Default:     true,
												},
												"allocate_method": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Using either Manual (IP Pool) or Automatic (DHCP). Must be set when \"use-allocate-method\" is true.",
													Default:     "manual",
												},
												"manual_network": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Manual Network. Identified by name or UID. Must be set when \"allocate-method\" was selected to be \"manual\".",
												},
												"dhcp_server": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "DHCP Server. Identified by name or UID. Must be set when \"allocate-method\" was selected to be \"automatic\".",
												},
												"virtual_ip_address": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Virtual IPV4 address for DHCP server replies. Must be set when \"allocate-method\" was selected to be \"automatic\".",
												},
												"dhcp_mac_address": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Calculated MAC address for DHCP allocation. Must be set when \"allocate-method\" was selected to be \"automatic\".",
													Default:     "per-machine",
												},
												"optional_parameters": {
													Type:        schema.TypeMap,
													Optional:    true,
													Description: "This configuration applies to all Office Mode methods except Automatic (using DHCP) and ipassignment.conf entries which contain this data.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"use_primary_dns_server": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Use Primary DNS Server.",
																Default:     false,
															},
															"primary_dns_server": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Primary DNS Server. Identified by name or UID. Must be set when \"use-primary-dns-server\" is true and can not be set when \"use-primary-dns-server\" is false.",
															},
															"use_first_backup_dns_server": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Use First Backup DNS Server.",
																Default:     false,
															},
															"first_backup_dns_server": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "First Backup DNS Server. Identified by name or UID. Must be set when \"use-first-backup-dns-server\" is true and can not be set when \"use-first-backup-dns-server\" is false.",
															},
															"use_second_backup_dns_server": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Use Second Backup DNS Server.",
																Default:     false,
															},
															"second_backup_dns_server": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Second Backup DNS Server. Identified by name or UID. Must be set when \"use-second-backup-dns-server\" is true and can not be set when \"use-second-backup-dns-server\" is false.",
															},
															"dns_suffixes": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "DNS Suffixes.",
															},
															"use_primary_wins_server": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Use Primary WINS Server.",
																Default:     false,
															},
															"primary_wins_server": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Primary WINS Server. Identified by name or UID. Must be set when \"use-primary-wins-server\" is true and can not be set when \"use-primary-wins-server\" is false.",
															},
															"use_first_backup_wins_server": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Use First Backup WINS Server.",
																Default:     false,
															},
															"first_backup_wins_server": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "First Backup WINS Server. Identified by name or UID. Must be set when \"use-first-backup-wins-server\" is true and can not be set when \"use-first-backup-wins-server\" is false.",
															},
															"use_second_backup_wins_server": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Use Second Backup WINS Server.",
																Default:     false,
															},
															"second_backup_wins_server": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Second Backup WINS Server. Identified by name or UID. Must be set when \"use-second-backup-wins-server\" is true and can not be set when \"use-second-backup-wins-server\" is false.",
															},
															"ip_lease_duration": {
																Type:        schema.TypeInt,
																Optional:    true,
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
										Optional:    true,
										Description: "Support connectivity enhancement for gateways with multiple external interfaces.",
										Default:     false,
									},
									"perform_anti_spoofing": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Perform Anti-Spoofing on Office Mode addresses.",
										Default:     false,
									},
									"anti_spoofing_additional_addresses": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Additional IP Addresses for Anti-Spoofing. Identified by name or UID. Must be set when \"perform-anti-spoofings\" is true.",
										Default:     "None",
									},
								},
							},
						},
						"remote_access": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Remote Access.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"support_l2tp": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Support L2TP (relevant only when office mode is active).",
										Default:     false,
									},
									"l2tp_auth_method": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "L2TP Authentication Method. Must be set when \"support-l2tp\" is true.",
										Default:     "md5",
									},
									"l2tp_certificate": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "L2TP Certificate. Must be set when \"l2tp-auth-method\" was selected to be \"certificate\". Insert \"defaultCert\" when you want to use the default certificate.",
									},
									"allow_vpn_clients_to_route_traffic": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Allow VPN clients to route traffic.",
										Default:     false,
									},
									"support_nat_traversal_mechanism": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Support NAT traversal mechanism (UDP encapsulation).",
										Default:     true,
									},
									"nat_traversal_service": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Allocated NAT traversal UDP service. Identified by name or UID. Must be set when \"support-nat-traversal-mechanism\" is true.",
										Default:     "VPN1_IPSEC_encapsulation",
									},
									"support_visitor_mode": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Support Visitor Mode.",
										Default:     false,
									},
									"visitor_mode_service": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "TCP Service for Visitor Mode. Identified by name or UID. Must be set when \"support-visitor-mode\" is true.",
										Default:     "https",
									},
									"visitor_mode_interface": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Interface for Visitor Mode. Must be set when \"support-visitor-mode\" is true. Insert IPV4 Address of existing interface or \"All IPs\" when you want all interfaces.",
										Default:     "All IPs",
									},
								},
							},
						},
						"vpn_domain": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Gateway VPN domain identified by the name or UID.",
						},
						"vpn_domain_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Gateway VPN domain type.",
						},
					},
				},
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
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ignore_warnings": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Apply changes ignoring warnings.",
				Default:     false,
			},
		},
	}
}

func createManagementSimpleGateway(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	gateway := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		gateway["name"] = v.(string)
	}

	if v, ok := d.GetOk("ipv4_address"); ok {
		gateway["ipv4-address"] = v.(string)
	}

	if v, ok := d.GetOk("ipv6_address"); ok {
		gateway["ipv6-address"] = v.(string)
	}

	if v, ok := d.GetOk("interfaces"); ok {
		interfacesList := v.([]interface{})
		if len(interfacesList) > 0 {
			var interfacesPayload []map[string]interface{}
			for i := range interfacesList {

				interfacePayload := make(map[string]interface{})

				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".name"); ok {
					interfacePayload["name"] = v.(string)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".ipv4_address"); ok {
					interfacePayload["ipv4-address"] = v.(string)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".ipv6_address"); ok {
					interfacePayload["ipv6-address"] = v.(string)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".ipv4_network_mask"); ok {
					interfacePayload["ipv4-network-mask"] = v.(string)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".ipv6_network_mask"); ok {
					interfacePayload["ipv6-network-mask"] = v.(string)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".ipv4_mask_length"); ok {
					interfacePayload["ipv4-mask-length"] = v.(string)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".ipv6_mask_length"); ok {
					interfacePayload["ipv6-mask-length"] = v.(string)
				}
				if v, ok := d.GetOkExists("interfaces." + strconv.Itoa(i) + ".anti_spoofing"); ok {
					interfacePayload["anti-spoofing"] = v
				}
				if _, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".anti_spoofing_settings"); ok {
					antiSpoofingSettings := make(map[string]interface{})
					if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".anti_spoofing_settings.action"); ok {
						antiSpoofingSettings["action"] = v.(string)
					}
					interfacePayload["anti-spoofing-settings"] = antiSpoofingSettings
				}
				if v, ok := d.GetOkExists("interfaces." + strconv.Itoa(i) + ".security_zone"); ok {
					interfacePayload["security-zone"] = v
				}

				if _, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".security_zone_settings"); ok {
					securityZoneSettings := make(map[string]interface{})
					if v, ok := d.GetOkExists("interfaces." + strconv.Itoa(i) + ".security_zone_settings.auto_calculated"); ok {
						securityZoneSettings["auto-calculated"] = v
					}
					if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".security_zone_settings.specific_zone"); ok {
						securityZoneSettings["specific-zone"] = v.(string)
					}
					interfacePayload["security-zone-settings"] = securityZoneSettings
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".topology"); ok {
					interfacePayload["topology"] = v.(string)
				}
				if _, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".topology_settings"); ok {
					topologySettings := make(map[string]interface{})

					if v, ok := d.GetOkExists("interfaces." + strconv.Itoa(i) + ".topology_settings.interface_leads_to_dmz"); ok {
						topologySettings["interface-leads-to-dmz"] = v
					}
					if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".topology_settings.ip_address_behind_this_interface"); ok {
						topologySettings["ip-address-behind-this-interface"] = v.(string)
					}
					if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".topology_settings.specific_network"); ok {
						topologySettings["specific-network"] = v.(string)
					}
					interfacePayload["topology-settings"] = topologySettings
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".color"); ok {
					interfacePayload["color"] = v.(string)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".comments"); ok {
					interfacePayload["comments"] = v.(string)
				}
				interfacesPayload = append(interfacesPayload, interfacePayload)
			}
			gateway["interfaces"] = interfacesPayload
		}
	}

	// Platform
	if v, ok := d.GetOk("one_time_password"); ok {
		gateway["one-time-password"] = v.(string)
	}

	if v, ok := d.GetOk("os_name"); ok {
		gateway["os-name"] = v.(string)
	}

	if v, ok := d.GetOk("version"); ok {
		gateway["version"] = v.(string)
	}

	// Blades
	if v, ok := d.GetOkExists("anti_bot"); ok {
		gateway["anti-bot"] = v
	}

	if v, ok := d.GetOkExists("anti_virus"); ok {
		gateway["anti-virus"] = v
	}

	if v, ok := d.GetOkExists("application_control"); ok {
		gateway["application-control"] = v
	}

	if v, ok := d.GetOkExists("content_awareness"); ok {
		gateway["content-awareness"] = v
	}

	if v, ok := d.GetOkExists("icap_server"); ok {
		gateway["icap-server"] = v
	}

	if v, ok := d.GetOkExists("ips"); ok {
		gateway["ips"] = v
	}

	if v, ok := d.GetOkExists("threat_emulation"); ok {
		gateway["threat-emulation"] = v
	}

	if v, ok := d.GetOkExists("threat_extraction"); ok {
		gateway["threat-extraction"] = v
	}

	if v, ok := d.GetOkExists("url_filtering"); ok {
		gateway["url-filtering"] = v
	}

	if v, ok := d.GetOkExists("vpn"); ok {
		gateway["vpn"] = v
	}

	if v, ok := d.GetOkExists("firewall"); ok {
		gateway["firewall"] = v
	}

	if _, ok := d.GetOk("firewall_settings"); ok {
		firewallSettings := make(map[string]interface{})
		if v, ok := d.GetOkExists("firewall_settings.auto_calculate_connections_hash_table_size_and_memory_pool"); ok {
			firewallSettings["auto-calculate-connections-hash-table-size-and-memory-pool"] = v
		}
		if v, ok := d.GetOkExists("firewall_settings.auto_maximum_limit_for_concurrent_connections"); ok {
			firewallSettings["auto-maximum-limit-for-concurrent-connections"] = v
		}
		if v, ok := d.GetOk("firewall_settings.connections_hash_size"); ok {
			firewallSettings["connections-hash-size"] = v.(int)
		}
		if v, ok := d.GetOk("firewall_settings.maximum_limit_for_concurrent_connections"); ok {
			firewallSettings["maximum-limit-for-concurrent-connections"] = v.(int)
		}
		if v, ok := d.GetOk("firewall_settings.maximum_memory_pool_size"); ok {
			firewallSettings["maximum-memory-pool-size"] = v.(int)
		}
		if v, ok := d.GetOk("firewall_settings.memory_pool_size"); ok {
			firewallSettings["memory-pool-size"] = v.(int)
		}
		gateway["firewall-settings"] = firewallSettings
	}

	// VPN settings
	if _, ok := d.GetOk("vpn_settings"); ok {
		vpnSettings := make(map[string]interface{})

		if _, ok := d.GetOk("vpn_settings.authentication"); ok {
			authentication := make(map[string]interface{})
			if v, ok := d.GetOk("vpn_settings.authentication.authentication_clients"); ok {
				authentication["authentication-clients"] = v.(*schema.Set).List()
			}
			vpnSettings["authentication"] = authentication
		}

		if _, ok := d.GetOk("vpn_settings.link_selection"); ok {
			linkSelection := make(map[string]interface{})
			if v, ok := d.GetOk("vpn_settings.link_selection.ip_selection"); ok {
				linkSelection["ip-selection"] = v.(string)
			}
			if v, ok := d.GetOk("vpn_settings.link_selection.dns_resolving_hostname"); ok {
				linkSelection["dns-resolving-hostname"] = v.(string)
			}
			if v, ok := d.GetOk("vpn_settings.link_selection.ip_address"); ok {
				linkSelection["ip-address"] = v.(string)
			}
			vpnSettings["link-selection"] = linkSelection
		}

		if v, ok := d.GetOk("vpn_settings.maximum_concurrent_ike_negotiations"); ok {
			vpnSettings["maximum-concurrent-ike-negotiations"] = v.(int)
		}
		if v, ok := d.GetOk("vpn_settings.maximum_concurrent_tunnels"); ok {
			vpnSettings["maximum-concurrent-tunnels"] = v.(int)
		}

		if _, ok := d.GetOk("vpn_settings.office_mode"); ok {
			officeMode := make(map[string]interface{})
			if v, ok := d.GetOk("vpn_settings.office_mode.mode"); ok {
				officeMode["mode"] = v.(string)
			}
			if v, ok := d.GetOk("vpn_settings.office_mode.group"); ok {
				officeMode["group"] = v.(string)
			}
			if v, ok := d.GetOkExists("vpn_settings.office_mode.support_multiple_interfaces"); ok {
				officeMode["support-multiple-interfaces"] = v
			}
			if v, ok := d.GetOkExists("vpn_settings.office_mode.perform_anti_spoofing"); ok {
				officeMode["perform-anti-spoofing"] = v
			}
			if v, ok := d.GetOk("vpn_settings.office_mode.anti_spoofing_additional_addresses"); ok {
				officeMode["anti-spoofing-additional-addresses"] = v.(string)
			}
			if _, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from"); ok {
				allocateIpAddressFrom := make(map[string]interface{})
				if v, ok := d.GetOkExists("vpn_settings.office_mode.allocate_ip_address_from.radius_server"); ok {
					allocateIpAddressFrom["radius-server"] = v
				}
				if v, ok := d.GetOkExists("vpn_settings.office_mode.allocate_ip_address_from.use_allocate_method"); ok {
					allocateIpAddressFrom["use-allocate-method"] = v
				}
				if v, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from.allocate_method"); ok {
					allocateIpAddressFrom["allocate-method"] = v.(string)
				}
				if v, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from.manual_network"); ok {
					allocateIpAddressFrom["manual-network"] = v.(string)
				}
				if v, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from.dhcp_server"); ok {
					allocateIpAddressFrom["dhcp-server"] = v.(string)
				}
				if v, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from.virtual_ip_address"); ok {
					allocateIpAddressFrom["virtual-ip-address"] = v.(string)
				}
				if v, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from.dhcp_mac_address"); ok {
					allocateIpAddressFrom["dhcp-mac-address"] = v.(string)
				}
				if _, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters"); ok {
					optionalParameters := make(map[string]interface{})
					if v, ok := d.GetOkExists("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.use_primary_dns_server"); ok {
						optionalParameters["use-primary-dns-server"] = v
					}
					if v, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.primary_dns_server"); ok {
						optionalParameters["primary-dns-server"] = v.(string)
					}
					if v, ok := d.GetOkExists("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.use_first_backup_dns_server"); ok {
						optionalParameters["use-first-backup-dns-server"] = v
					}
					if v, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.first_backup_dns_server"); ok {
						optionalParameters["first-backup-dns-server"] = v.(string)
					}
					if v, ok := d.GetOkExists("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.use_second_backup_dns_server"); ok {
						optionalParameters["use-second-backup-dns-server"] = v
					}
					if v, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.second_backup_dns_server"); ok {
						optionalParameters["second-backup-dns-server"] = v.(string)
					}
					if v, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.dns_suffixes"); ok {
						optionalParameters["dns-suffixes"] = v.(string)
					}
					if v, ok := d.GetOkExists("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.use_primary_wins_server"); ok {
						optionalParameters["use-primary-wins-server"] = v
					}
					if v, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.primary_wins_server"); ok {
						optionalParameters["primary-wins-server"] = v.(string)
					}
					if v, ok := d.GetOkExists("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.use_first_backup_wins_server"); ok {
						optionalParameters["use-first-backup-wins-server"] = v
					}
					if v, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.first_backup_wins_server"); ok {
						optionalParameters["first-backup-wins-server"] = v.(string)
					}
					if v, ok := d.GetOkExists("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.use_second_backup_wins_server"); ok {
						optionalParameters["use-second-backup-wins-server"] = v
					}
					if v, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.second_backup_wins_server"); ok {
						optionalParameters["second-backup-wins-server"] = v.(string)
					}
					if v, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.ip_lease_duration"); ok {
						optionalParameters["ip-lease-duration"] = v.(int)
					}
					allocateIpAddressFrom["optional-parameters"] = optionalParameters
				}
				officeMode["allocate-ip-address-from"] = allocateIpAddressFrom
			}
			vpnSettings["office-mode"] = officeMode
		}

		if _, ok := d.GetOk("vpn_settings.remote_access"); ok {
			remoteAccess := make(map[string]interface{})
			if v, ok := d.GetOkExists("vpn_settings.remote_access.support_l2tp"); ok {
				remoteAccess["support-l2tp"] = v
			}
			if v, ok := d.GetOk("vpn_settings.remote_access.l2tp_auth_method"); ok {
				remoteAccess["l2tp-auth-method"] = v.(string)
			}
			if v, ok := d.GetOk("vpn_settings.remote_access.l2tp_certificate"); ok {
				remoteAccess["l2tp-certificate"] = v.(string)
			}
			if v, ok := d.GetOkExists("vpn_settings.remote_access.allow_vpn_clients_to_route_traffic"); ok {
				remoteAccess["allow-vpn-clients-to-route-traffic"] = v
			}
			if v, ok := d.GetOkExists("vpn_settings.remote_access.support_nat_traversal_mechanism"); ok {
				remoteAccess["support-nat-traversal-mechanism"] = v
			}
			if v, ok := d.GetOk("vpn_settings.remote_access.nat_traversal_service"); ok {
				remoteAccess["nat-traversal-service"] = v.(string)
			}
			if v, ok := d.GetOkExists("vpn_settings.remote_access.support_visitor_mode"); ok {
				remoteAccess["support-visitor-mode"] = v
			}
			if v, ok := d.GetOkExists("vpn_settings.remote_access.visitor_mode_service"); ok {
				remoteAccess["visitor-mode-service"] = v.(string)
			}
			if v, ok := d.GetOkExists("vpn_settings.remote_access.visitor_mode_interface"); ok {
				remoteAccess["visitor-mode-interface"] = v.(string)
			}
			vpnSettings["remote-access"] = remoteAccess
		}

		if v, ok := d.GetOk("vpn_settings.vpn_domain"); ok {
			vpnSettings["vpn-domain"] = v.(string)
		}
		if v, ok := d.GetOk("vpn_settings.vpn_domain_type"); ok {
			vpnSettings["vpn-domain-type"] = v.(string)
		}
		gateway["vpn-settings"] = vpnSettings
	}

	// Logs
	if v, ok := d.GetOkExists("save_logs_locally"); ok {
		gateway["save-logs-locally"] = v
	}

	if v, ok := d.GetOk("send_alerts_to_server"); ok {
		gateway["send-alerts-to-server"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("send_logs_to_backup_server"); ok {
		gateway["send-logs-to-backup-server"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("send_logs_to_server"); ok {
		gateway["send-logs-to-server"] = v.(*schema.Set).List()
	}

	if _, ok := d.GetOk("logs_settings"); ok {
		logsSettings := make(map[string]interface{})
		if v, ok := d.GetOkExists("logs_settings.alert_when_free_disk_space_below"); ok {
			logsSettings["alert-when-free-disk-space-below"] = v
		}
		if v, ok := d.GetOk("logs_settings.alert_when_free_disk_space_below_threshold"); ok {
			logsSettings["alert-when-free-disk-space-below-threshold"] = v.(int)
		}
		if v, ok := d.GetOk("logs_settings.alert_when_free_disk_space_below_type"); ok {
			logsSettings["alert-when-free-disk-space-below-type"] = v.(string)
		}
		if v, ok := d.GetOkExists("logs_settings.before_delete_keep_logs_from_the_last_days"); ok {
			logsSettings["before-delete-keep-logs-from-the-last-days"] = v
		}
		if v, ok := d.GetOk("logs_settings.before_delete_keep_logs_from_the_last_days_threshold"); ok {
			logsSettings["before-delete-keep-logs-from-the-last-days-threshold"] = v.(int)
		}
		if v, ok := d.GetOkExists("logs_settings.before_delete_run_script"); ok {
			logsSettings["before-delete-run-script"] = v
		}
		if v, ok := d.GetOk("logs_settings.before_delete_run_script_command"); ok {
			logsSettings["before-delete-run-script-command"] = v.(string)
		}
		if v, ok := d.GetOkExists("logs_settings.delete_index_files_older_than_days"); ok {
			logsSettings["delete-index-files-older-than-days"] = v
		}
		if v, ok := d.GetOk("logs_settings.delete_index_files_older_than_days_threshold"); ok {
			logsSettings["delete-index-files-older-than-days-threshold"] = v.(int)
		}
		if v, ok := d.GetOkExists("logs_settings.delete_index_files_when_index_size_above"); ok {
			logsSettings["delete-index-files-when-index-size-above"] = v
		}
		if v, ok := d.GetOk("logs_settings.delete_index_files_when_index_size_above_threshold"); ok {
			logsSettings["delete-index-files-when-index-size-above-threshold"] = v.(int)
		}
		if v, ok := d.GetOkExists("logs_settings.delete_when_free_disk_space_below"); ok {
			logsSettings["delete-when-free-disk-space-below"] = v
		}
		if v, ok := d.GetOk("logs_settings.delete_when_free_disk_space_below_threshold"); ok {
			logsSettings["delete-when-free-disk-space-below-threshold"] = v.(int)
		}
		if v, ok := d.GetOkExists("logs_settings.detect_new_citrix_ica_application_names"); ok {
			logsSettings["detect-new-citrix-ica-application-names"] = v
		}
		if v, ok := d.GetOkExists("logs_settings.forward_logs_to_log_server"); ok {
			logsSettings["forward-logs-to-log-server"] = v
		}
		if v, ok := d.GetOk("logs_settings.forward_logs_to_log_server_name"); ok {
			logsSettings["forward-logs-to-log-server-name"] = v.(string)
		}
		if v, ok := d.GetOk("logs_settings.forward_logs_to_log_server_schedule_name"); ok {
			logsSettings["forward-logs-to-log-server-schedule-name"] = v.(string)
		}
		if v, ok := d.GetOk("logs_settings.free_disk_space_metrics"); ok {
			logsSettings["free-disk-space-metrics"] = v.(string)
		}
		if v, ok := d.GetOkExists("logs_settings.perform_log_rotate_before_log_forwarding"); ok {
			logsSettings["perform-log-rotate-before-log-forwarding"] = v
		}
		if v, ok := d.GetOkExists("logs_settings.reject_connections_when_free_disk_space_below_threshold"); ok {
			logsSettings["reject-connections-when-free-disk-space-below-threshold"] = v
		}
		if v, ok := d.GetOk("logs_settings.reserve_for_packet_capture_metrics"); ok {
			logsSettings["reserve-for-packet-capture-metrics"] = v.(string)
		}
		if v, ok := d.GetOk("logs_settings.reserve_for_packet_capture_threshold"); ok {
			logsSettings["reserve-for-packet-capture-threshold"] = v.(int)
		}
		if v, ok := d.GetOkExists("logs_settings.rotate_log_by_file_size"); ok {
			logsSettings["rotate-log-by-file-size"] = v
		}
		if v, ok := d.GetOk("logs_settings.rotate_log_file_size_threshold"); ok {
			logsSettings["rotate-log-file-size-threshold"] = v.(int)
		}
		if v, ok := d.GetOkExists("logs_settings.rotate_log_on_schedule"); ok {
			logsSettings["rotate-log-on-schedule"] = v
		}
		if v, ok := d.GetOk("logs_settings.rotate_log_schedule_name"); ok {
			logsSettings["rotate-log-schedule-name"] = v.(string)
		}
		if v, ok := d.GetOkExists("logs_settings.stop_logging_when_free_disk_space_below"); ok {
			logsSettings["stop-logging-when-free-disk-space-below"] = v
		}
		if v, ok := d.GetOk("logs_settings.stop_logging_when_free_disk_space_below_threshold"); ok {
			logsSettings["stop-logging-when-free-disk-space-below-threshold"] = v.(int)
		}
		if v, ok := d.GetOkExists("logs_settings.turn_on_qos_logging"); ok {
			logsSettings["turn-on-qos-logging"] = v
		}
		if v, ok := d.GetOk("logs_settings.update_account_log_every"); ok {
			logsSettings["update-account-log-every"] = v.(int)
		}
		gateway["logs-settings"] = logsSettings
	}

	// General
	if v, ok := d.GetOk("tags"); ok {
		gateway["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("comments"); ok {
		gateway["comments"] = v.(string)
	}

	if v, ok := d.GetOk("color"); ok {
		gateway["color"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		gateway["ignore-warnings"] = v
	}

	log.Println("Create Simple Gateway - Map = ", gateway)

	addGatewayRes, err := client.ApiCall("add-simple-gateway", gateway, client.GetSessionID(), true, false)
	if err != nil || !addGatewayRes.Success {
		if addGatewayRes.ErrorMsg != "" {
			return fmt.Errorf(addGatewayRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addGatewayRes.GetData()["uid"].(string))

	return readManagementSimpleGateway(d, m)
}

func readManagementSimpleGateway(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showGatewayRes, err := client.ApiCall("show-simple-gateway", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showGatewayRes.Success {
		if objectNotFound(showGatewayRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showGatewayRes.ErrorMsg)
	}

	gateway := showGatewayRes.GetData()

	log.Println("Read Simple Gateway - Show JSON = ", gateway)

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
			logSettingsState["free_disk_space_metrics"] = v
		}
		if v := logSettingsJson["delete-index-files-when-index-size-above-metrics"]; v != nil {
			logSettingsState["delete_index_files_when_index_size_above_metrics"] = v
		}
		if v := logSettingsJson["delete-when-free-disk-space-below-metrics"]; v != nil {
			logSettingsState["delete_when_free_disk_space_below_metrics"] = v
		}
		if v := logSettingsJson["stop-logging-when-free-disk-space-below-metrics"]; v != nil {
			logSettingsState["stop_logging_when_free_disk_space_below_metrics"] = v
		}
		if v := logSettingsJson["alert-when-free-disk-space-below-threshold"]; v != nil {
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
		if v := logSettingsJson["delete-index-files-when-index-size-above-threshold"]; v != nil {
			logSettingsState["delete_index_files_when_index_size_above_threshold"] = v
		}
		if v := logSettingsJson["delete-when-free-disk-space-below"]; v != nil {
			logSettingsState["delete_when_free_disk_space_below"] = v
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
		if v := logSettingsJson["stop-logging-when-free-disk-space-below-threshold"]; v != nil {
			logSettingsState["stop_logging_when_free_disk_space_below_threshold"] = v
		}
		if v := logSettingsJson["turn-on-qos-logging"]; v != nil {
			logSettingsState["turn_on_qos_logging"] = v
		}
		if v := logSettingsJson["update-account-log-every"]; v != nil {
			logSettingsState["update_account_log_every"] = v
		}

		_, logsSettingsInConf := d.GetOk("logs_settings")
		defaultLogsSettings := map[string]interface{}{
			"alert_when_free_disk_space_below":                        true,
			"free_disk_space_metrics":                                 "mbytes",
			"delete_index_files_when_index_size_above_metrics":        "mbytes",
			"delete_when_free_disk_space_below_metrics":               "mbytes",
			"stop_logging_when_free_disk_space_below_metrics":         "mbytes",
			"alert_when_free_disk_space_below_type":                   "popup alert",
			"alert_when_free_disk_space_below_threshold":              20,
			"before_delete_keep_logs_from_the_last_days":              false,
			"before_delete_keep_logs_from_the_last_days_threshold":    3664,
			"before_delete_run_script":                                false,
			"before_delete_run_script_command":                        "",
			"delete_index_files_older_than_days":                      false,
			"delete_index_files_older_than_days_threshold":            14,
			"delete_index_files_when_index_size_above":                false,
			"delete_index_files_when_index_size_above_threshold":      100000,
			"delete_when_free_disk_space_below":                       true,
			"delete_when_free_disk_space_below_threshold":             5000,
			"detect_new_citrix_ica_application_names":                 false,
			"forward_logs_to_log_server":                              false,
			"perform_log_rotate_before_log_forwarding":                false,
			"reject_connections_when_free_disk_space_below_threshold": false,
			"reserve_for_packet_capture_metrics":                      "mbytes",
			"reserve_for_packet_capture_threshold":                    500,
			"rotate_log_by_file_size":                                 false,
			"rotate_log_file_size_threshold":                          1000,
			"rotate_log_on_schedule":                                  false,
			"stop_logging_when_free_disk_space_below":                 true,
			"stop_logging_when_free_disk_space_below_threshold":       100,
			"turn_on_qos_logging":                                     true,
			"update_account_log_every":                                3600,
		}
		if reflect.DeepEqual(defaultLogsSettings, logSettingsState) && !logsSettingsInConf {
			_ = d.Set("logs_settings", map[string]interface{}{})
		} else {
			_ = d.Set("logs_settings", logSettingsState)
		}
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

func updateManagementSimpleGateway(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	gateway := make(map[string]interface{})

	gateway["uid"] = d.Id()

	if d.HasChange("name") {
		gateway["new-name"] = d.Get("name")
	}

	if ok := d.HasChange("ipv4_address"); ok {
		gateway["ipv4-address"] = d.Get("ipv4_address")
	}

	if ok := d.HasChange("ipv6_address"); ok {
		gateway["ipv6-address"] = d.Get("ipv6_address")
	}

	if d.HasChange("interfaces") {
		if v, ok := d.GetOk("interfaces"); ok {
			interfacesList := v.([]interface{})
			var interfacesPayload []map[string]interface{}
			for i := range interfacesList {
				interfacePayload := make(map[string]interface{})
				interfacePayload["name"] = d.Get("interfaces." + strconv.Itoa(i) + ".name").(string)
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".ipv4_address"); ok {
					interfacePayload["ipv4-address"] = v.(string)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".ipv6_address"); ok {
					interfacePayload["ipv6-address"] = v.(string)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".ipv4_network_mask"); ok {
					interfacePayload["ipv4-network-mask"] = v.(string)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".ipv6_network_mask"); ok {
					interfacePayload["ipv6-network-mask"] = v.(string)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".ipv4_mask_length"); ok {
					interfacePayload["ipv4-mask-length"] = v.(string)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".ipv6_mask_length"); ok {
					interfacePayload["ipv6-mask-length"] = v.(string)
				}
				if v, ok := d.GetOkExists("interfaces." + strconv.Itoa(i) + ".anti_spoofing"); ok {
					interfacePayload["anti-spoofing"] = v
				}
				if _, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".anti_spoofing_settings"); ok {
					antiSpoofingSettings := make(map[string]interface{})
					if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".anti_spoofing_settings.action"); ok {
						antiSpoofingSettings["action"] = v.(string)
					}
					interfacePayload["anti-spoofing-settings"] = antiSpoofingSettings
				}
				if v, ok := d.GetOkExists("interfaces." + strconv.Itoa(i) + ".security_zone"); ok {
					interfacePayload["security-zone"] = v
				}

				if _, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".security_zone_settings"); ok {
					securityZoneSettings := make(map[string]interface{})
					if v, ok := d.GetOkExists("interfaces." + strconv.Itoa(i) + ".security_zone_settings.auto_calculated"); ok {
						securityZoneSettings["auto-calculated"] = v
					}
					if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".security_zone_settings.specific_zone"); ok {
						securityZoneSettings["specific-zone"] = v.(string)
					}
					interfacePayload["security-zone-settings"] = securityZoneSettings
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".topology"); ok {
					interfacePayload["topology"] = v.(string)
				}
				if _, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".topology_settings"); ok {
					topologySettings := make(map[string]interface{})

					if v, ok := d.GetOkExists("interfaces." + strconv.Itoa(i) + ".topology_settings.interface_leads_to_dmz"); ok {
						topologySettings["interface-leads-to-dmz"] = v
					}
					if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".topology_settings.ip_address_behind_this_interface"); ok {
						topologySettings["ip-address-behind-this-interface"] = v.(string)
					}
					if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".topology_settings.specific_network"); ok {
						topologySettings["specific-network"] = v.(string)
					}
					interfacePayload["topology-settings"] = topologySettings
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".color"); ok {
					interfacePayload["color"] = v.(string)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".comments"); ok {
					interfacePayload["comments"] = v.(string)
				}
				interfacesPayload = append(interfacesPayload, interfacePayload)
			}
			gateway["interfaces"] = interfacesPayload
		}
	}

	if ok := d.HasChange("one_time_password"); ok {
		gateway["one-time-password"] = d.Get("one_time_password").(string)
	}

	if ok := d.HasChange("os_name"); ok {
		gateway["os-name"] = d.Get("os_name").(string)
	}

	if ok := d.HasChange("version"); ok {
		gateway["version"] = d.Get("version").(string)
	}

	// Blades
	if ok := d.HasChange("anti_bot"); ok {
		gateway["anti-bot"] = d.Get("anti_bot")
	}

	if ok := d.HasChange("anti_virus"); ok {
		gateway["anti-virus"] = d.Get("anti_virus")
	}

	if ok := d.HasChange("application_control"); ok {
		gateway["application-control"] = d.Get("application_control")
	}

	if ok := d.HasChange("content_awareness"); ok {
		gateway["content-awareness"] = d.Get("content_awareness")
	}

	if ok := d.HasChange("icap_server"); ok {
		gateway["icap-server"] = d.Get("icap_server")
	}

	if ok := d.HasChange("ips"); ok {
		gateway["ips"] = d.Get("ips")
	}

	if ok := d.HasChange("threat_emulation"); ok {
		gateway["threat-emulation"] = d.Get("threat_emulation")
	}

	if ok := d.HasChange("threat_extraction"); ok {
		gateway["threat-extraction"] = d.Get("threat_extraction")
	}

	if ok := d.HasChange("url_filtering"); ok {
		gateway["url-filtering"] = d.Get("url_filtering")
	}

	if ok := d.HasChange("vpn"); ok {
		gateway["vpn"] = d.Get("vpn")
	}

	if ok := d.HasChange("firewall"); ok {
		gateway["firewall"] = d.Get("firewall")
	}

	if ok := d.HasChange("firewall_settings"); ok {
		if _, ok := d.GetOk("firewall_settings"); ok {
			firewallSettings := make(map[string]interface{})
			if ok := d.HasChange("firewall_settings.auto_calculate_connections_hash_table_size_and_memory_pool"); ok {
				firewallSettings["auto-calculate-connections-hash-table-size-and-memory-pool"] = d.Get("firewall_settings.auto_calculate_connections_hash_table_size_and_memory_pool")
			}
			if ok := d.HasChange("firewall_settings.auto_maximum_limit_for_concurrent_connections"); ok {
				firewallSettings["auto-maximum-limit-for-concurrent-connections"] = d.Get("firewall_settings.auto_maximum_limit_for_concurrent_connections")
			}
			if ok := d.HasChange("firewall_settings.connections_hash_size"); ok {
				firewallSettings["connections-hash-size"] = d.Get("firewall_settings.connections_hash_size").(int)
			}
			if ok := d.HasChange("firewall_settings.maximum_limit_for_concurrent_connections"); ok {
				firewallSettings["maximum-limit-for-concurrent-connections"] = d.Get("firewall_settings.maximum_limit_for_concurrent_connections").(int)
			}
			if ok := d.HasChange("firewall_settings.maximum_memory_pool_size"); ok {
				firewallSettings["maximum-memory-pool-size"] = d.Get("firewall_settings.maximum_memory_pool_size").(int)
			}
			if ok := d.HasChange("firewall_settings.memory_pool_size"); ok {
				firewallSettings["memory-pool-size"] = d.Get("firewall_settings.memory_pool_size").(int)
			}
			gateway["firewall-settings"] = firewallSettings
		}
	}

	// VPN settings
	if ok := d.HasChange("vpn_settings"); ok {
		if _, ok := d.GetOk("vpn_settings"); ok {
			vpnSettings := make(map[string]interface{})

			if ok := d.HasChange("vpn_settings.authentication"); ok {
				if _, ok := d.GetOk("vpn_settings.authentication"); ok {
					authentication := make(map[string]interface{})
					if ok := d.HasChange("vpn_settings.authentication.authentication_clients"); ok {
						if v, ok := d.GetOk("vpn_settings.authentication.authentication_clients"); ok {
							authentication["authentication-clients"] = v.(*schema.Set).List()
						} else {
							oldValues, _ := d.GetChange("vpn_settings.authentication.authentication_clients")
							authentication["authentication-clients"] = map[string]interface{}{"remove": oldValues.(*schema.Set).List()}
						}
					}
					vpnSettings["authentication"] = authentication
				}
			}

			if ok := d.HasChange("vpn_settings.link_selection"); ok {
				if _, ok := d.GetOk("vpn_settings.link_selection"); ok {
					linkSelection := make(map[string]interface{})
					if ok := d.HasChange("vpn_settings.link_selection.ip_selection"); ok {
						linkSelection["ip-selection"] = d.Get("vpn_settings.link_selection.ip_selection").(string)
					}
					if ok := d.HasChange("vpn_settings.link_selection.dns_resolving_hostname"); ok {
						linkSelection["dns-resolving-hostname"] = d.Get("vpn_settings.link_selection.dns_resolving_hostname").(string)
					}
					if ok := d.HasChange("vpn_settings.link_selection.ip_address"); ok {
						linkSelection["ip-address"] = d.Get("vpn_settings.link_selection.ip_address").(string)
					}
					vpnSettings["link-selection"] = linkSelection
				}
			}

			if ok := d.HasChange("vpn_settings.maximum_concurrent_ike_negotiations"); ok {
				vpnSettings["maximum-concurrent-ike-negotiations"] = d.Get("vpn_settings.maximum_concurrent_ike_negotiations").(int)
			}

			if ok := d.HasChange("vpn_settings.maximum_concurrent_tunnels"); ok {
				vpnSettings["maximum-concurrent-tunnels"] = d.Get("vpn_settings.maximum_concurrent_tunnels").(int)
			}

			if ok := d.HasChange("vpn_settings.office_mode"); ok {
				if _, ok := d.GetOk("vpn_settings.office_mode"); ok {
					officeMode := make(map[string]interface{})

					if ok := d.HasChange("vpn_settings.office_mode.mode"); ok {
						officeMode["mode"] = d.Get("vpn_settings.office_mode.mode").(string)
					}
					if ok := d.HasChange("vpn_settings.office_mode.group"); ok {
						officeMode["group"] = d.Get("vpn_settings.office_mode.group").(string)
					}
					if ok := d.HasChange("vpn_settings.office_mode.support_multiple_interfaces"); ok {
						officeMode["support-multiple-interfaces"] = d.Get("vpn_settings.office_mode.support_multiple_interfaces")
					}
					if ok := d.HasChange("vpn_settings.office_mode.perform_anti_spoofing"); ok {
						officeMode["perform-anti-spoofing"] = d.Get("vpn_settings.office_mode.perform_anti_spoofing")
					}
					if ok := d.HasChange("vpn_settings.office_mode.anti_spoofing_additional_addresses"); ok {
						officeMode["anti-spoofing-additional-addresses"] = d.Get("vpn_settings.office_mode.anti_spoofing_additional_addresses").(string)
					}

					if ok := d.HasChange("vpn_settings.office_mode.allocate_ip_address_from"); ok {
						if _, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from"); ok {
							allocateIpAddressFrom := make(map[string]interface{})

							if ok := d.HasChange("vpn_settings.office_mode.allocate_ip_address_from.radius_server"); ok {
								allocateIpAddressFrom["radius-server"] = d.Get("vpn_settings.office_mode.allocate_ip_address_from.radius_server")
							}
							if ok := d.HasChange("vpn_settings.office_mode.allocate_ip_address_from.use_allocate_method"); ok {
								allocateIpAddressFrom["use-allocate-method"] = d.Get("vpn_settings.office_mode.allocate_ip_address_from.use_allocate_method")
							}
							if ok := d.HasChange("vpn_settings.office_mode.allocate_ip_address_from.allocate_method"); ok {
								allocateIpAddressFrom["allocate-method"] = d.Get("vpn_settings.office_mode.allocate_ip_address_from.allocate_method").(string)
							}
							if ok := d.HasChange("vpn_settings.office_mode.allocate_ip_address_from.manual_network"); ok {
								allocateIpAddressFrom["manual-network"] = d.Get("vpn_settings.office_mode.allocate_ip_address_from.manual_network").(string)
							}
							if ok := d.HasChange("vpn_settings.office_mode.allocate_ip_address_from.dhcp_server"); ok {
								allocateIpAddressFrom["dhcp-server"] = d.Get("vpn_settings.office_mode.allocate_ip_address_from.dhcp_server").(string)
							}
							if ok := d.HasChange("vpn_settings.office_mode.allocate_ip_address_from.virtual_ip_address"); ok {
								allocateIpAddressFrom["virtual-ip-address"] = d.Get("vpn_settings.office_mode.allocate_ip_address_from.virtual_ip_address").(string)
							}
							if ok := d.HasChange("vpn_settings.office_mode.allocate_ip_address_from.dhcp_mac_address"); ok {
								allocateIpAddressFrom["dhcp-mac-address"] = d.Get("vpn_settings.office_mode.allocate_ip_address_from.dhcp_mac_address").(string)
							}
							if ok := d.HasChange("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters"); ok {
								if _, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters"); ok {
									optionalParameters := make(map[string]interface{})

									if ok := d.HasChange("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.use_primary_dns_server"); ok {
										optionalParameters["use-primary-dns-server"] = d.Get("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.use_primary_dns_server")
									}
									if ok := d.HasChange("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.primary_dns_server"); ok {
										optionalParameters["primary-dns-server"] = d.Get("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.primary_dns_server").(string)
									}
									if ok := d.HasChange("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.use_first_backup_dns_server"); ok {
										optionalParameters["use-first-backup-dns-server"] = d.Get("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.use_first_backup_dns_server")
									}
									if ok := d.HasChange("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.first_backup_dns_server"); ok {
										optionalParameters["first-backup-dns-server"] = d.Get("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.first_backup_dns_server").(string)
									}
									if ok := d.HasChange("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.use_second_backup_dns_server"); ok {
										optionalParameters["use-second-backup-dns-server"] = d.Get("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.use_second_backup_dns_server")
									}
									if ok := d.HasChange("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.second_backup_dns_server"); ok {
										optionalParameters["second-backup-dns-server"] = d.Get("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.second_backup_dns_server").(string)
									}
									if ok := d.HasChange("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.dns_suffixes"); ok {
										optionalParameters["dns-suffixes"] = d.Get("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.dns_suffixes").(string)
									}
									if ok := d.HasChange("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.use_primary_wins_server"); ok {
										optionalParameters["use-primary-wins-server"] = d.Get("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.use_primary_wins_server")
									}
									if ok := d.HasChange("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.primary_wins_server"); ok {
										optionalParameters["primary-wins-server"] = d.Get("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.primary_wins_server").(string)
									}
									if ok := d.HasChange("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.use_first_backup_wins_server"); ok {
										optionalParameters["use-first-backup-wins-server"] = d.Get("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.use_first_backup_wins_server")
									}
									if ok := d.HasChange("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.first_backup_wins_server"); ok {
										optionalParameters["first-backup-wins-server"] = d.Get("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.first_backup_wins_server").(string)
									}
									if ok := d.HasChange("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.use_second_backup_wins_server"); ok {
										optionalParameters["use-second-backup-wins-server"] = d.Get("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.use_second_backup_wins_server")
									}
									if ok := d.HasChange("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.second_backup_wins_server"); ok {
										optionalParameters["second-backup-wins-server"] = d.Get("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.second_backup_wins_server").(string)
									}
									if ok := d.HasChange("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.ip_lease_duration"); ok {
										optionalParameters["ip-lease-duration"] = d.Get("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.ip_lease_duration").(int)
									}
									allocateIpAddressFrom["optional-parameters"] = optionalParameters
								}
							}
							officeMode["allocate-ip-address-from"] = allocateIpAddressFrom
						}
					}
					vpnSettings["office-mode"] = officeMode
				}
			}

			if ok := d.HasChange("vpn_settings.remote_access"); ok {
				if _, ok := d.GetOk("vpn_settings.remote_access"); ok {
					remoteAccess := make(map[string]interface{})
					if ok := d.HasChange("vpn_settings.remote_access.support_l2tp"); ok {
						remoteAccess["support-l2tp"] = d.Get("vpn_settings.remote_access.support_l2tp")
					}
					if ok := d.HasChange("vpn_settings.remote_access.l2tp_auth_method"); ok {
						remoteAccess["l2tp-auth-method"] = d.Get("vpn_settings.remote_access.l2tp_auth_method").(string)
					}
					if ok := d.HasChange("vpn_settings.remote_access.l2tp_certificate"); ok {
						remoteAccess["l2tp-certificate"] = d.Get("vpn_settings.remote_access.l2tp_certificate").(string)
					}
					if ok := d.HasChange("vpn_settings.remote_access.allow_vpn_clients_to_route_traffic"); ok {
						remoteAccess["allow-vpn-clients-to-route-traffic"] = d.Get("vpn_settings.remote_access.allow_vpn_clients_to_route_traffic")
					}
					if ok := d.HasChange("vpn_settings.remote_access.support_nat_traversal_mechanism"); ok {
						remoteAccess["support-nat-traversal-mechanism"] = d.Get("vpn_settings.remote_access.support_nat_traversal_mechanism")
					}
					if ok := d.HasChange("vpn_settings.remote_access.nat_traversal_service"); ok {
						remoteAccess["nat-traversal-service"] = d.Get("vpn_settings.remote_access.nat_traversal_service").(string)
					}
					if ok := d.HasChange("vpn_settings.remote_access.support_visitor_mode"); ok {
						remoteAccess["support-visitor-mode"] = d.Get("vpn_settings.remote_access.support_visitor_mode")
					}
					if ok := d.HasChange("vpn_settings.remote_access.visitor_mode_service"); ok {
						remoteAccess["visitor-mode-service"] = d.Get("vpn_settings.remote_access.visitor_mode_service").(string)
					}
					if ok := d.HasChange("vpn_settings.remote_access.visitor_mode_interface"); ok {
						remoteAccess["visitor-mode-interface"] = d.Get("vpn_settings.remote_access.visitor_mode_interface").(string)
					}
					vpnSettings["remote-access"] = remoteAccess
				}
			}

			if ok := d.HasChange("vpn_settings.vpn_domain"); ok {
				vpnSettings["vpn-domain"] = d.Get("vpn_settings.vpn_domain").(string)
			}
			if ok := d.HasChange("vpn_settings.vpn_domain_type"); ok {
				vpnSettings["vpn-domain-type"] = d.Get("vpn_settings.vpn_domain_type").(string)
			}
			gateway["vpn-settings"] = vpnSettings
		}
	}

	// Logs
	if ok := d.HasChange("save_logs_locally"); ok {
		gateway["save-logs-locally"] = d.Get("save_logs_locally")
	}

	if ok := d.HasChange("send_alerts_to_server"); ok {
		if v, ok := d.GetOk("send_alerts_to_server"); ok {
			gateway["send-alerts-to-server"] = v.(*schema.Set).List()
		} else {
			oldValues, _ := d.GetChange("send_alerts_to_server")
			gateway["send-alerts-to-server"] = map[string]interface{}{"remove": oldValues.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("send_logs_to_backup_server"); ok {
		if v, ok := d.GetOk("send_logs_to_backup_server"); ok {
			gateway["send-logs-to-backup-server"] = v.(*schema.Set).List()
		} else {
			oldValues, _ := d.GetChange("send_logs_to_backup_server")
			gateway["send-logs-to-backup-server"] = map[string]interface{}{"remove": oldValues.(*schema.Set).List()}
		}
	}
	if ok := d.HasChange("send_logs_to_server"); ok {
		if v, ok := d.GetOk("send_logs_to_server"); ok {
			gateway["send-logs-to-server"] = v.(*schema.Set).List()
		} else {
			oldValues, _ := d.GetChange("send_logs_to_server")
			gateway["send-logs-to-server"] = map[string]interface{}{"remove": oldValues.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("logs_settings"); ok {
		if _, ok := d.GetOk("logs_settings"); ok {
			logsSettings := make(map[string]interface{})
			if ok := d.HasChange("logs_settings.alert_when_free_disk_space_below"); ok {
				logsSettings["alert-when-free-disk-space-below"] = d.Get("logs_settings.alert_when_free_disk_space_below")
			}
			if ok := d.HasChange("logs_settings.alert_when_free_disk_space_below_threshold"); ok {
				logsSettings["alert-when-free-disk-space-below-threshold"] = d.Get("logs_settings.alert_when_free_disk_space_below_threshold").(int)
			}
			if ok := d.HasChange("logs_settings.alert_when_free_disk_space_below_type"); ok {
				logsSettings["alert-when-free-disk-space-below-type"] = d.Get("logs_settings.alert_when_free_disk_space_below_type").(string)
			}
			if ok := d.HasChange("logs_settings.before_delete_keep_logs_from_the_last_days"); ok {
				logsSettings["before-delete-keep-logs-from-the-last-days"] = d.Get("logs_settings.before_delete_keep_logs_from_the_last_days")
			}
			if ok := d.HasChange("logs_settings.before_delete_keep_logs_from_the_last_days_threshold"); ok {
				logsSettings["before-delete-keep-logs-from-the-last-days-threshold"] = d.Get("logs_settings.before_delete_keep_logs_from_the_last_days_threshold").(int)
			}
			if ok := d.HasChange("logs_settings.before_delete_run_script"); ok {
				logsSettings["before-delete-run-script"] = d.Get("logs_settings.before_delete_run_script")
			}
			if ok := d.HasChange("logs_settings.before_delete_run_script_command"); ok {
				logsSettings["before-delete-run-script-command"] = d.Get("logs_settings.before_delete_run_script_command").(string)
			}
			if ok := d.HasChange("logs_settings.delete_index_files_older_than_days"); ok {
				logsSettings["delete-index-files-older-than-days"] = d.Get("logs_settings.delete_index_files_older_than_days")
			}
			if ok := d.HasChange("logs_settings.delete_index_files_older_than_days_threshold"); ok {
				logsSettings["delete-index-files-older-than-days-threshold"] = d.Get("logs_settings.delete_index_files_older_than_days_threshold").(int)
			}
			if ok := d.HasChange("logs_settings.delete_index_files_when_index_size_above"); ok {
				logsSettings["delete-index-files-when-index-size-above"] = d.Get("logs_settings.delete_index_files_when_index_size_above")
			}
			if ok := d.HasChange("logs_settings.delete_index_files_when_index_size_above_threshold"); ok {
				logsSettings["delete-index-files-when-index-size-above-threshold"] = d.Get("logs_settings.delete_index_files_when_index_size_above_threshold").(int)
			}
			if ok := d.HasChange("logs_settings.delete_when_free_disk_space_below"); ok {
				logsSettings["delete-when-free-disk-space-below"] = d.Get("logs_settings.delete_when_free_disk_space_below")
			}
			if ok := d.HasChange("logs_settings.delete_when_free_disk_space_below_threshold"); ok {
				logsSettings["delete-when-free-disk-space-below-threshold"] = d.Get("logs_settings.delete_when_free_disk_space_below_threshold").(int)
			}
			if ok := d.HasChange("logs_settings.detect_new_citrix_ica_application_names"); ok {
				logsSettings["detect-new-citrix-ica-application-names"] = d.Get("logs_settings.detect_new_citrix_ica_application_names")
			}
			if ok := d.HasChange("logs_settings.forward_logs_to_log_server"); ok {
				logsSettings["forward-logs-to-log-server"] = d.Get("logs_settings.forward_logs_to_log_server")
			}
			if ok := d.HasChange("logs_settings.forward_logs_to_log_server_name"); ok {
				logsSettings["forward-logs-to-log-server-name"] = d.Get("logs_settings.forward_logs_to_log_server_name").(string)
			}
			if ok := d.HasChange("logs_settings.forward_logs_to_log_server_schedule_name"); ok {
				logsSettings["forward-logs-to-log-server-schedule-name"] = d.Get("logs_settings.forward_logs_to_log_server_schedule_name").(string)
			}
			if ok := d.HasChange("logs_settings.free_disk_space_metrics"); ok {
				logsSettings["free-disk-space-metrics"] = d.Get("logs_settings.free_disk_space_metrics").(string)
			}
			if ok := d.HasChange("logs_settings.perform_log_rotate_before_log_forwarding"); ok {
				logsSettings["perform-log-rotate-before-log-forwarding"] = d.Get("logs_settings.perform_log_rotate_before_log_forwarding")
			}
			if ok := d.HasChange("logs_settings.reject_connections_when_free_disk_space_below_threshold"); ok {
				logsSettings["reject-connections-when-free-disk-space-below-threshold"] = d.Get("logs_settings.reject_connections_when_free_disk_space_below_threshold")
			}
			if ok := d.HasChange("logs_settings.reserve_for_packet_capture_metrics"); ok {
				logsSettings["reserve-for-packet-capture-metrics"] = d.Get("logs_settings.reserve_for_packet_capture_metrics").(string)
			}
			if ok := d.HasChange("logs_settings.reserve_for_packet_capture_threshold"); ok {
				logsSettings["reserve-for-packet-capture-threshold"] = d.Get("logs_settings.reserve_for_packet_capture_threshold").(int)
			}
			if ok := d.HasChange("logs_settings.rotate_log_by_file_size"); ok {
				logsSettings["rotate-log-by-file-size"] = d.Get("logs_settings.rotate_log_by_file_size")
			}
			if ok := d.HasChange("logs_settings.rotate_log_file_size_threshold"); ok {
				logsSettings["rotate-log-file-size-threshold"] = d.Get("logs_settings.rotate_log_file_size_threshold").(int)
			}
			if ok := d.HasChange("logs_settings.rotate_log_on_schedule"); ok {
				logsSettings["rotate-log-on-schedule"] = d.Get("logs_settings.rotate_log_on_schedule")
			}
			if ok := d.HasChange("logs_settings.rotate_log_schedule_name"); ok {
				logsSettings["rotate-log-schedule-name"] = d.Get("logs_settings.rotate_log_schedule_name").(string)
			}
			if ok := d.HasChange("logs_settings.stop_logging_when_free_disk_space_below"); ok {
				logsSettings["stop-logging-when-free-disk-space-below"] = d.Get("logs_settings.stop_logging_when_free_disk_space_below")
			}
			if ok := d.HasChange("logs_settings.stop_logging_when_free_disk_space_below_threshold"); ok {
				logsSettings["stop-logging-when-free-disk-space-below-threshold"] = d.Get("logs_settings.stop_logging_when_free_disk_space_below_threshold").(int)
			}
			if ok := d.HasChange("logs_settings.turn_on_qos_logging"); ok {
				logsSettings["turn-on-qos-logging"] = d.Get("logs_settings.turn_on_qos_logging")
			}
			if ok := d.HasChange("logs_settings.update_account_log_every"); ok {
				logsSettings["update-account-log-every"] = d.Get("logs_settings.update_account_log_every").(int)
			}
			gateway["logs-settings"] = logsSettings
		}
	}

	if ok := d.HasChange("tags"); ok {
		if v, ok := d.GetOk("tags"); ok {
			gateway["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			gateway["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("comments"); ok {
		gateway["comments"] = d.Get("comments").(string)
	}

	if ok := d.HasChange("color"); ok {
		gateway["color"] = d.Get("color").(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		gateway["ignore-warnings"] = v
	}

	log.Println("Update Simple Gateway - Map = ", gateway)
	updateSimpleGatewayRes, err := client.ApiCall("set-simple-gateway", gateway, client.GetSessionID(), true, false)
	if err != nil || !updateSimpleGatewayRes.Success {
		if updateSimpleGatewayRes.ErrorMsg != "" {
			return fmt.Errorf(updateSimpleGatewayRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementSimpleGateway(d, m)
}

func deleteManagementSimpleGateway(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	gatewayPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	deleteGatewayRes, err := client.ApiCall("delete-simple-gateway", gatewayPayload, client.GetSessionID(), true, false)
	if err != nil || !deleteGatewayRes.Success {
		if deleteGatewayRes.ErrorMsg != "" {
			return fmt.Errorf(deleteGatewayRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
