package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
)

func resourceManagementSetGlobalProperties() *schema.Resource {
	return &schema.Resource{
		Create: createManagementSetGlobalProperties,
		Read:   readManagementSetGlobalProperties,
		Delete: deleteManagementSetGlobalProperties,
		Schema: map[string]*schema.Schema{
			"firewall": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Add implied rules to or remove them from the Firewall Rule Base. Determine the position of the implied rules in the Rule Base, and whether or not to log them.",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"accept_control_connections": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Used for: Installing the security policy from the Security Management server to the gateways. Sending logs from the gateways to the Security Management server.Communication between SmartConsole clients and the Security Management Server.  Communication between Firewall daemons on different machines (Security Management Server, Security Gateway).< Connecting to OPSEC applications such as RADIUS and TACACS authentication servers. If you disable Accept Control Connections and you want Check Point components to communicate with each other and with OPSEC components, you must explicitly allow these connections in the Rule Base.",
							Default:     true,
						},
						"accept_ips1_management_connections": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Accepts IPS-1 connections. Available only if accept-control-connections is true.",
						},
						"accept_remote_access_control_connections": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Accepts Remote Access connections. Available only if accept-control-connections is true.",
							Default:     true,
						},
						"accept_smart_update_connections": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Accepts SmartUpdate connections.",
						},
						"accept_outgoing_packets_originating_from_gw": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Accepts all packets from connections that originate at the Check Point Security Gateway.",
							Default:     true,
						},
						"accept_outgoing_packets_originating_from_gw_position": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The position of the implied rules in the Rule Base. Available only if accept-outgoing-packets-originating-from-gw is false.",
							Default:     "before last",
						},
						"accept_outgoing_packets_originating_from_connectra_gw": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Accepts outgoing packets originating from Connectra gateway. Available only if accept-outgoing-packets-originating-from-gw is false.",
							Default:     true,
						},
						"accept_outgoing_packets_to_cp_online_services": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Allow Security Gateways to access Check Point online services. Supported for R80.10 Gateway and higher. Available only if accept-outgoing-packets-originating-from-gw is false.",
						},
						"accept_outgoing_packets_to_cp_online_services_position": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The position of the implied rules in the Rule Base. Available only if accept-outgoing-packets-to-cp-online-services is true.",
							Default:     "before last",
						},
						"accept_domain_name_over_tcp": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Accepts Domain Name (DNS) queries and replies over TCP, to allow downloading of the domain name-resolving tables used for zone transfers between servers. For clients, DNS over TCP is only used if the tables to be transferred are very large.",
							Default:     true,
						},
						"accept_domain_name_over_tcp_position": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The position of the implied rules in the Rule Base. Available only if accept-domain-name-over-tcp is true.",
							Default:     "first",
						},
						"accept_domain_name_over_udp": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Accepts Domain Name (DNS) queries and replies over UDP.",
							Default:     true,
						},
						"accept_domain_name_over_udp_position": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The position of the implied rules in the Rule Base. Available only if accept-domain-name-over-udp is true.",
							Default:     "first",
						},
						"accept_dynamic_addr_modules_outgoing_internet_connections": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Accept Dynamic Address modules' outgoing internet connections.Accepts DHCP traffic for DAIP (Dynamically Assigned IP Address) gateways. In Small Office Appliance gateways, this rule allows outgoing DHCP, PPP, PPTP and L2TP Internet connections (regardless of whether it is or is not a DAIP gateway).",
							Default:     true,
						},
						"accept_icmp_requests": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Accepts Internet Control Message Protocol messages.",
							Default:     true,
						},
						"accept_icmp_requests_position": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The position of the implied rules in the Rule Base. Available only if accept-icmp-requests is true.",
							Default:     "before last",
						},
						"accept_identity_awareness_control_connections": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Accepts traffic between Security Gateways in distributed environment configurations of Identity Awareness.",
							Default:     true,
						},
						"accept_identity_awareness_control_connections_position": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The position of the implied rules in the Rule Base.<br>Available only if accept-identity-awareness-control-connections is true.",
							Default:     "first",
						},
						"accept_incoming_traffic_to_dhcp_and_dns_services_of_gws": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Allows the Small Office Appliance gateway to provide DHCP relay, DHCP server and DNS proxy services regardless of the rule base.",
							Default:     true,
						},
						"accept_rip": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Accepts Routing Information Protocol (RIP), using UDP on port 520.",
						},
						"accept_rip_position": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The position of the implied rules in the Rule Base. Available only if accept-rip is true.",
							Default:     "first",
						},
						"accept_vrrp_packets_originating_from_cluster_members": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Selecting this option creates an implied rule in the security policy Rule Base that accepts VRRP inbound and outbound traffic to and from the members of the cluster.",
							Default:     true,
						},
						"accept_web_and_ssh_connections_for_gw_administration": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Accepts Web and SSH connections for Small Office Appliance gateways.",
							Default:     true,
						},
						"log_implied_rules": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Produces log records for communications that match the implied rules that are generated in the Rule Base from the properties defined in this window.",
						},
						"security_server": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Control the welcome messages that users will see when logging in to servers behind Check Point Security Gateways.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"client_auth_welcome_file": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Client authentication welcome file is the name of a file whose contents are to be displayed when a user begins a Client Authenticated session (optional) using the Manual Sign On Method. Client Authenticated Sessions initiated by Manual Sign On are not mediated by a security server.",
									},
									"ftp_welcome_msg_file": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "FTP welcome message file is the name of a file whose contents are to be displayed when a user begins an Authenticated FTP session.",
									},
									"rlogin_welcome_msg_file": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Rlogin welcome message file is the name of a file whose contents are to be displayed when a user begins an Authenticated RLOGIN session.",
									},
									"telnet_welcome_msg_file": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Telnet welcome message file is the name of a file whose contents are to be displayed when a user begins an Authenticated Telnet session.",
									},
									"mdq_welcome_msg": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "MDQ Welcome Message is the message to be displayed when a user begins an MDQ session. The MDQ Welcome Message should contain characters according to RFC 1035 and it must follow the ARPANET host name rules:<br>   - This message must begin with a number or letter. After the first letter or number character the remaining characters can be a letter, number, space, tab or hyphen.<br>   - This message must not end with a space or a tab and is limited to 63 characters.",
									},
									"smtp_welcome_msg": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "SMTP Welcome Message is the message to be displayed when a user begins an SMTP session.",
									},
									"http_next_proxy_host": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "HTTP next proxy host is the host name of the HTTP proxy behind the Check Point Security Gateway HTTP security server (if there is one). Changing the HTTP Next Proxy fields takes effect after the Security Gateway database is downloaded to the authenticating gateway, or after the security policy is re-installed. <br>These settings apply only to firewalled gateways prior to NG. For later versions, these settings should be defined in the Node Properties window.",
									},
									"http_next_proxy_port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "HTTP next proxy port is the port of the HTTP proxy behind the Check Point Security Gateway HTTP security server (if there is one). Changing the HTTP Next Proxy fields takes effect after the Security Gateway database is downloaded to the authenticating gateway, or after the security policy is re-installed. <br>These settings apply only to firewalled gateways prior to NG. For later versions, these settings should be defined in the Node Properties window.",
									},
									"http_servers": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "This list specifies the HTTP servers. Defining HTTP servers allows you to restrict incoming HTTP.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"logical_name": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Unique Logical Name of the HTTP Server.",
												},
												"host": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Host name of the HTTP Server.",
												},
												"port": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "Port number of the HTTP Server.",
													Default:     80,
												},
												"reauthentication": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Specify whether users must reauthenticate when accessing a specific server.",
													Default:     "Standard",
												},
											},
										},
									},
									"server_for_null_requests": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The Logical Name of a Null Requests Server from http-servers.",
									},
								},
							},
						},
					},
				},
			},
			"nat": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Configure settings that apply to all NAT connections.",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allow_bi_directional_nat": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Applies to automatic NAT rules in the NAT Rule Base, and allows two automatic NAT rules to match a connection. Without Bidirectional NAT, only one automatic NAT rule can match a connection.",
						},
						"auto_arp_conf": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Ensures that ARP requests for a translated (NATed) machine, network or address range are answered by the Check Point Security Gateway.",
						},
						"merge_manual_proxy_arp_conf": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Merges the automatic and manual ARP configurations. Manual proxy ARP configuration is required for manual Static NAT rules.<br>Available only if auto-arp-conf is true.",
						},
						"auto_translate_dest_on_client_side": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Applies to packets originating at the client, with the server as its destination. Static NAT for the server is performed on the client side.",
							Default:     true,
						},
						"manually_translate_dest_on_client_side": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Applies to packets originating at the client, with the server as its destination. Static NAT for the server is performed on the client side.",
							Default:     true,
						},
						"enable_ip_pool_nat": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Applies to packets originating at the client, with the server as its destination. Static NAT for the server is performed on the client side.",
						},
						"addr_alloc_and_release_track": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies whether to log each allocation and release of an IP address from the IP Pool. Available only if enable-ip-pool-nat is true.",
						},
						"addr_exhaustion_track": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies the action to take if the IP Pool is exhausted. Available only if enable-ip-pool-nat is true.",
						},
					},
				},
			},
			"authentication": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Define Authentication properties that are common to all users and to the various ways that the Check Point Security Gateway asks for passwords (User, Client and Session Authentication).",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_internal_users_with_specific_suffix": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enforce suffix for internal users authentication.",
							Default:     true,
						},
						"allowed_suffix_for_internal_users": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Suffix for internal users authentication.",
						},
						"max_days_before_expiration_of_non_pulled_user_certificates": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Users certificates which were initiated but not pulled will expire after the specified number of days. Any value from 1 to 60 days can be entered in this field.",
							Default:     14,
						},
						"max_client_auth_attempts_before_connection_termination": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Allowed Number of Failed Client Authentication Attempts Before Session Termination. Any value from 1 to 800 attempts can be entered in this field.",
							Default:     3,
						},
						"max_rlogin_attempts_before_connection_termination": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Allowed Number of Failed rlogin Attempts Before Session Termination. Any value from 1 to 800 attempts can be entered in this field.",
							Default:     3,
						},
						"max_session_auth_attempts_before_connection_termination": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Allowed Number of Failed Session Authentication Attempts Before Session Termination. Any value from 1 to 800 attempts can be entered in this field.",
							Default:     3,
						},
						"max_telnet_attempts_before_connection_termination": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Allowed Number of Failed telnet Attempts Before Session Termination. Any value from 1 to 800 attempts can be entered in this field.",
							Default:     3,
						},
						"enable_delayed_auth": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "all authentications other than certificate-based authentications will be delayed by the specified time. Applying this delay will stall brute force authentication attacks. The delay is applied for both failed and successful authentication attempts.",
							Default:     false,
						},
						"delay_each_auth_attempt_by": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Delay each authentication attempt by the specified number of milliseconds. Any value from 1 to 25000 can be entered in this field.",
							Default:     100,
						},
					},
				},
			},
			"vpn": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Configure settings relevant to VPN.",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vpn_conf_method": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Decide on Simplified or Traditional mode for all new security policies or decide which mode to use on a policy by policy basis.",
							Default:     "simplified",
						},
						"domain_name_for_dns_resolving": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Enter the domain name that will be used for gateways DNS lookup. The DNS host name that is used is \"gateway_name.domain_name\".",
						},
						"enable_backup_gw": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable Backup Gateway.",
						},
						"enable_decrypt_on_accept_for_gw_to_gw_traffic": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable decrypt on accept for gateway to gateway traffic. This is only relevant for policies in traditional mode. In Traditional Mode, the 'Accept' action determines that a connection is allowed, while the 'Encrypt' action determines that a connection is allowed and encrypted. Select whether VPN accepts an encrypted packet that matches a rule with an 'Accept' action or drops it.",
							Default:     true,
						},
						"enable_load_distribution_for_mep_conf": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable load distribution for Multiple Entry Points configurations (Site To Site connections). The VPN Multiple Entry Point (MEP) feature supplies high availability and load distribution for Check Point Security Gateways. MEP works in four modes:<br>&nbsp;&nbsp;&nbsp;&nbsp; <ul><li> First to Respond, in which the first gateway to reply to the peer gateway is chosen. An organization would choose this option if, for example, the organization has two gateways in a MEPed configuration - one in London, the other in New York. It makes sense for Check Point Security Gateway peers located in England to try the London gateway first and the NY gateway second. Being geographically closer to Check Point Security Gateway peers in England, the London gateway will be the first to respond, and becomes the entry point to the internal network.</li><br>&nbsp;&nbsp;&nbsp;&nbsp; <li> VPN Domain, is when the destination IP belongs to a particular VPN domain, the gateway of that domain becomes the chosen entry point. This gateway becomes the primary gateway while other gateways in the MEP configuration become its backup gateways.</li><br>&nbsp;&nbsp;&nbsp;&nbsp; <li> Random Selection, in which the remote Check Point Security Gateway peer randomly selects a gateway with which to open a VPN connection. For each IP source/destination address pair, a new gateway is randomly selected. An organization might have a number of machines with equal performance abilities. In this case, it makes sense to enable load distribution. The machines are used in a random and equal way.</li><br>&nbsp;&nbsp;&nbsp;&nbsp; <li> Manually set priority list, gateway priorities can be set manually for the entire community or for individual satellite gateways.</li></ul>.",
						},
						"enable_vpn_directional_match_in_vpn_column": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable VPN Directional Match in VPN Column.<br>Note: VPN Directional Match is supported only on Gaia, SecurePlatform, Linux and IPSO.",
							Default:     false,
						},
						"grace_period_after_the_crl_is_not_valid": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "When establishing VPN tunnels, the peer presents its certificate for authentication. The clock on the gateway machine must be synchronized with the clock on the Certificate Authority machine. Otherwise, the Certificate Revocation List (CRL) used for validating the peer's certificate may be considered invalid and thus the authentication fails. To resolve the issue of differing clock times, a Grace Period permits a wider window for CRL validity.",
							Default:     1800,
						},
						"grace_period_before_the_crl_is_valid": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "When establishing VPN tunnels, the peer presents its certificate for authentication. The clock on the gateway machine must be synchronized with the clock on the Certificate Authority machine. Otherwise, the Certificate Revocation List (CRL) used for validating the peer's certificate may be considered invalid and thus the authentication fails. To resolve the issue of differing clock times, a Grace Period permits a wider window for CRL validity.",
							Default:     7200,
						},
						"grace_period_extension_for_secure_remote_secure_client": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "When dealing with remote clients the Grace Period needs to be extended. The remote client sometimes relies on the peer gateway to supply the CRL. If the client's clock is not synchronized with the gateway's clock, a CRL that is considered valid by the gateway may be considered invalid by the client.",
							Default:     3600,
						},
						"support_ike_dos_protection_from_identified_src": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "When the number of IKE negotiations handled simultaneously exceeds a threshold above VPN's capacity, a gateway concludes that it is either under a high load or experiencing a Denial of Service attack. VPN can filter out peers that are the probable source of the potential Denial of Service attack. There are two kinds of protection:<br>&nbsp;&nbsp;&nbsp;&nbsp; <ul><li> Stateless - the peer has to respond to an IKE notification in a way that proves the peer's IP address is not spoofed. If the peer cannot prove this, VPN does not allocate resources for the IKE negotiation</li><br>&nbsp;&nbsp;&nbsp;&nbsp; <li> Puzzles - this is the same as Stateless, but in addition, the peer has to solve a mathematical puzzle. Solving this puzzle consumes peer CPU resources in a way that makes it difficult to initiate multiple IKE negotiations simultaneously.</li></ul>Puzzles is more secure then Stateless, but affects performance.<br>Since these kinds of attacks involve a new proprietary addition to the IKE protocol, enabling these protection mechanisms may cause difficulties with non Check Point VPN products or older versions of VPN.",
							Default:     "stateless",
						},
						"support_ike_dos_protection_from_unidentified_src": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "When the number of IKE negotiations handled simultaneously exceeds a threshold above VPN's capacity, a gateway concludes that it is either under a high load or experiencing a Denial of Service attack. VPN can filter out peers that are the probable source of the potential Denial of Service attack. There are two kinds of protection:<br>&nbsp;&nbsp;&nbsp;&nbsp; <ul><li> Stateless - the peer has to respond to an IKE notification in a way that proves the peer's IP address is not spoofed. If the peer cannot prove this, VPN does not allocate resources for the IKE negotiation</li><br>&nbsp;&nbsp;&nbsp;&nbsp; <li> Puzzles - this is the same as Stateless, but in addition, the peer has to solve a mathematical puzzle. Solving this puzzle consumes peer CPU resources in a way that makes it difficult to initiate multiple IKE negotiations simultaneously.</li></ul>Puzzles is more secure then Stateless, but affects performance.<br>Since these kinds of attacks involve a new proprietary addition to the IKE protocol, enabling these protection mechanisms may cause difficulties with non Check Point VPN products or older versions of VPN.",
							Default:     "puzzles",
						},
					},
				},
			},
			"remote_access": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Configure Remote Access properties.",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_back_connections": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Usually communication with remote clients must be initialized by the clients. However, once a client has opened a connection, the hosts behind VPN can open a return or back connection to the client. For a back connection, the client's details must be maintained on all the devices between the client and the gateway, and on the gateway itself. Determine whether the back connection is enabled.",
						},
						"keep_alive_packet_to_gw_interval": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Usually communication with remote clients must be initialized by the clients. However, once a client has opened a connection, the hosts behind VPN can open a return or back connection to the client. For a back connection, the client's details must be maintained on all the devices between the client and the gateway, and on the gateway itself. Determine frequency (in seconds) of the Keep Alive packets sent by the client in order to maintain the connection with the gateway.<br>Available only if enable-back-connections is true.",
						},
						"encrypt_dns_traffic": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "You can decide whether DNS queries sent by the remote client to a DNS server located on the corporate LAN are passed through the VPN tunnel or not. Disable this option if the client has to make DNS queries to the DNS server on the corporate LAN while connecting to the organization but without using the SecuRemote client.",
						},
						"simultaneous_login_mode": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Select the simultaneous login mode.",
						},
						"vpn_authentication_and_encryption": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "configure supported Encryption and Authentication methods for Remote Access clients.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"encryption_algorithms": {
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "Select the methods negotiated in IKE phase 2 and used in IPSec connections.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ike": {
													Type:        schema.TypeMap,
													Optional:    true,
													Description: "Configure the IKE Phase 1 settings.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"support_encryption_algorithms": {
																Type:        schema.TypeMap,
																Optional:    true,
																Description: "Select the encryption algorithms that will be supported with remote hosts.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"aes_128": {
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Select whether the AES-128 encryption algorithm will be supported with remote hosts.",
																		},
																		"aes_256": {
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Select whether the AES-256 encryption algorithm will be supported with remote hosts.",
																		},
																		"des": {
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Select whether the DES encryption algorithm will be supported with remote hosts.",
																		},
																		"tdes": {
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Select whether the Triple DES encryption algorithm will be supported with remote hosts.",
																		},
																	},
																},
															},
															"use_encryption_algorithm": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Choose the encryption algorithm that will have the highest priority of the selected algorithms. If given a choice of more that one encryption algorithm to use, the algorithm selected in this field will be used.",
															},
															"support_data_integrity": {
																Type:        schema.TypeMap,
																Optional:    true,
																Description: "Select the hash algorithms that will be supported with remote hosts to ensure data integrity.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"aes_xcbc": {
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Select whether the AES-XCBC hash algorithm will be supported with remote hosts to ensure data integrity.",
																		},
																		"md5": {
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Select whether the MD5 hash algorithm will be supported with remote hosts to ensure data integrity.",
																		},
																		"sha1": {
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Select whether the SHA1 hash algorithm will be supported with remote hosts to ensure data integrity.",
																		},
																		"sha256": {
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Select whether the SHA256 hash algorithm will be supported with remote hosts to ensure data integrity.",
																		},
																	},
																},
															},
															"use_data_integrity": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The hash algorithm chosen here will be given the highest priority if more than one choice is offered.",
															},
															"support_diffie_hellman_groups": {
																Type:        schema.TypeMap,
																Optional:    true,
																Description: "Select the Diffie-Hellman groups that will be supported with remote hosts.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"group1": {
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Select whether Diffie-Hellman Group 1 (768 bit) will be supported with remote hosts.",
																		},
																		"group14": {
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Select whether Diffie-Hellman Group 14 (2048 bit) will be supported with remote hosts.",
																		},
																		"group2": {
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Select whether Diffie-Hellman Group 2 (1024 bit) will be supported with remote hosts.",
																			Default:     true,
																		},
																		"group5": {
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Select whether Diffie-Hellman Group 5 (1536 bit) will be supported with remote hosts.",
																		},
																	},
																},
															},
															"use_diffie_hellman_group": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "SecureClient users utilize the Diffie-Hellman group selected in this field.",
																Default:     "Group 2",
															},
														},
													},
												},
												"ipsec": {
													Type:        schema.TypeMap,
													Optional:    true,
													Description: "Configure the IPSEC Phase 2 settings.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"support_encryption_algorithms": {
																Type:        schema.TypeMap,
																Optional:    true,
																Description: "Select the encryption algorithms that will be supported with remote hosts.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"aes_128": {
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Select whether the AES-128 encryption algorithm will be supported with remote hosts.",
																		},
																		"aes_256": {
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Select whether the AES-256 encryption algorithm will be supported with remote hosts.",
																		},
																		"des": {
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Select whether the DES encryption algorithm will be supported with remote hosts.",
																		},
																		"tdes": {
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Select whether the Triple DES encryption algorithm will be supported with remote hosts.",
																		},
																	},
																},
															},
															"use_encryption_algorithm": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Choose the encryption algorithm that will have the highest priority of the selected algorithms. If given a choice of more that one encryption algorithm to use, the algorithm selected in this field will be used.",
															},
															"support_data_integrity": {
																Type:        schema.TypeMap,
																Optional:    true,
																Description: "Select the hash algorithms that will be supported with remote hosts to ensure data integrity.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"aes_xcbc": {
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Select whether the AES-XCBC hash algorithm will be supported with remote hosts to ensure data integrity.",
																		},
																		"md5": {
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Select whether the MD5 hash algorithm will be supported with remote hosts to ensure data integrity.",
																		},
																		"sha1": {
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Select whether the SHA1 hash algorithm will be supported with remote hosts to ensure data integrity.",
																		},
																		"sha256": {
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Select whether the SHA256 hash algorithm will be supported with remote hosts to ensure data integrity.",
																		},
																	},
																},
															},
															"use_data_integrity": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The hash algorithm chosen here will be given the highest priority if more than one choice is offered.",
															},
															"enforce_encryption_alg_and_data_integrity_on_all_users": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Enforce Encryption Algorithm and Data Integrity on all users.",
															},
														},
													},
												},
											},
										},
									},
									"encryption_method": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Select the encryption method.",
										Default:     "ike_v1_only",
									},
									"pre_shared_secret": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "the user password is specified in the Authentication tab in the user's IKE properties (in the user properties window: Encryption tab > Edit).",
										Default:     false,
									},
									"support_legacy_auth_for_sc_l2tp_nokia_clients": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Support Legacy Authentication for SC (hybrid mode), L2TP (PAP) and Nokia clients (CRACK).",
										Default:     true,
									},
									"support_legacy_eap": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Support Legacy EAP (Extensible Authentication Protocol).",
										Default:     true,
									},
									"support_l2tp_with_pre_shared_key": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Use a centrally managed pre-shared key for IKE.",
										Default:     false,
									},
									"l2tp_pre_shared_key": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Type in the pre-shared key.<br>Available only if support-l2tp-with-pre-shared-key is set to true.",
									},
								},
							},
						},
						"vpn_advanced": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Configure encryption methods and interface resolution for remote access clients.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"allow_clear_traffic_to_encryption_domain_when_disconnected": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "SecuRemote/SecureClient behavior while disconnected - How traffic to the VPN domain is handled when the Remote Access VPN client is not connected to the site. Traffic can either be dropped or sent in clear without encryption.",
										Default:     true,
									},
									"enable_load_distribution_for_mep_conf": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Load distribution for Multiple Entry Points configurations - Remote access clients will randomly select a gateway from the list of entry points. Make sure to define the same VPN domain for all the Security Gateways you want to be entry points.",
									},
									"use_first_allocated_om_ip_addr_for_all_conn_to_the_gws_of_the_site": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Use first allocated Office Mode IP Address for all connections to the Gateways of the site.After a remote user connects and receives an Office Mode IP address from a gateway, every connection to that gateways encryption domain will go out with the Office Mode IP as the internal source IP. The Office Mode IP is what hosts in the encryption domain will recognize as the remote user's IP address. The Office Mode IP address assigned by a specific gateway can be used in its own encryption domain and in neighboring encryption domains as well. The neighboring encryption domains should reside behind gateways that are members of the same VPN community as the assigning gateway. Since the remote hosts connections are dependant on the Office Mode IP address it received, should the gateway that issued the IP become unavailable, all the connections to the site will terminate.",
									},
								},
							},
						},
						"scv": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Define properties of the Secure Configuration Verification process.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"apply_scv_on_simplified_mode_fw_policies": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Determine whether the gateway verifies that remote access clients are securely configured. This is set here only if the security policy is defined in the Simplified Mode. If the security policy is defined in the Traditional Mode, verification takes place per rule.",
									},
									"exceptions": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Specify the hosts that can be accessed using the selected services even if the client is not verified.<br>Available only if apply-scv-on-simplified-mode-fw-policies is true.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"hosts": {
													Type:        schema.TypeSet,
													Optional:    true,
													Description: "Specify the Hosts to be excluded from SCV.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"services": {
													Type:        schema.TypeSet,
													Optional:    true,
													Description: "Specify the services to be accessed.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"no_scv_for_unsupported_cp_clients": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Do not apply Secure Configuration Verification for connections from Check Point VPN clients that don't support it, such as SSL Network Extender, GO, Capsule VPN / Connect, Endpoint Connects lower than R75, or L2TP clients.<br>Available only if apply-scv-on-simplified-mode-fw-policies is true.",
									},
									"upon_verification_accept_and_log_client_connection": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If the gateway verifies the client's configuration, decide how the gateway should handle connections with clients that fail the Security Configuration Verification. It is possible to either drop the connection or Accept the connection and log it.",
									},
									"only_tcp_ip_protocols_are_used": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Most SCV checks are configured via the SCV policy. Specify whether to verify that  only TCP/IP protocols are used.",
									},
									"policy_installed_on_all_interfaces": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Most SCV checks are configured via the SCV policy. Specify whether to verify that  the Desktop Security Policy is installed on all the interfaces of the client.",
									},
									"generate_log": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If the client identifies that the secure configuration has been violated, select whether a log is generated by the remote access client and sent to the Security Management server.",
									},
									"notify_user": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If the client identifies that the secure configuration has been violated, select whether to user should be notified.",
									},
								},
							},
						},
						"ssl_network_extender": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Define properties for SSL Network Extender users.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"user_auth_method": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Wide Impact: Also applies for SecureClient Mobile devices and Check Point GO clients!<br>User authentication method indicates how the user will be authenticated by the gateway. Changes made here will also apply for SSL clients.<br>Legacy - Username and password only.<br>Certificate - Certificate only with an existing certificate.<br>Certificate with Enrollment - Allows you to obtain a new certificate and then use certificate authentication only.<br>Mixed - Can use either username and password or certificate.",
										Default:     "legacy",
									},
									"supported_encryption_methods": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Wide Impact: Also applies to SecureClient Mobile devices!<br>Select the encryption algorithms that will be supported for remote users. Changes made here will also apply for all SSL clients.",
										Default:     "d3des_or_rc4",
									},
									"client_upgrade_upon_connection": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "When a client connects to the gateway with SSL Network Extender, the client automatically checks for upgrade. Select whether the client should automatically upgrade.",
										Default:     "ask_user",
									},
									"client_uninstall_upon_disconnection": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Select whether the client should automatically uninstall SSL Network Extender when it disconnects from the gateway.",
										Default:     "dont_uninstall",
									},
									"re_auth_user_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Wide Impact: Applies for the SecureClient Mobile!<br>Select the interval that users will need to reauthenticate.",
										Default:     480,
									},
									"scan_ep_machine_for_compliance_with_ep_compliance_policy": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Set to true if you want endpoint machines to be scanned for compliance with the Endpoint Compliance Policy.",
										Default:     false,
									},
									"client_outgoing_keep_alive_packets_frequency": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Select the interval which the keep-alive packets are sent.",
										Default:     20,
									},
								},
							},
						},
						"secure_client_mobile": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Define properties for SecureClient Mobile.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"user_auth_method": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Wide Impact: Also applies for SSL Network Extender clients and Check Point GO clients.<br>How the user will be authenticated by the gateway.",
										Default:     "legacy",
									},
									"enable_password_caching": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "If the password entered to authenticate is saved locally on the user's machine.",
										Default:     "false",
									},
									"cache_password_timeout": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Cached password timeout (in minutes).",
										Default:     1440,
									},
									"re_auth_user_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Wide Impact: Also applies for SSL Network Extender clients!<br>The length of time (in minutes) until the user's credentials are resent to the gateway to verify authorization.",
										Default:     480,
									},
									"connect_mode": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Methods by which a connection to the gateway will be initiated:<br>Configured On Endpoint Client - the method used for initiating a connection to a gateway is determined by the endpoint client<br>Manual - VPN connections will not be initiated automatically.<br>Always connected - SecureClient Mobile will automatically establish a connection to the last connected gateway under the following circumstances: (a) the device has a valid IP address, (b) when the device \"wakes up\" from a low-power state or a soft-reset, or (c) after a condition that caused the device to automatically disconnect ceases to exist (for example, Device is out of PC Sync, Disconnect is not idle.).<br>On application request - Applications requiring access to resources through the VPN will be able to initiate a VPN connection.",
										Default:     "Configured On Endpoint Client",
									},
									"automatically_initiate_dialup": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "When selected, the client will initiate a GPRS dialup connection before attempting to establish the VPN connection. Note that if a local IP address is already available through another network interface, then the GPRS dialup is not initiated.",
										Default:     "client_decide",
									},
									"disconnect_when_device_is_idle": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Enabling this feature will disconnect users from the gateway if there is no traffic sent during the defined time period.",
										Default:     "client_decide",
									},
									"supported_encryption_methods": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Wide Impact: Also applies for SSL Network Extender clients!<br>Select the encryption algorithms that will be supported with remote users.",
										Default:     "d3des_or_rc4",
									},
									"route_all_traffic_to_gw": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Operates the client in Hub Mode, sending all traffic to the VPN server for routing, filtering, and processing.",
										Default:     "false",
									},
								},
							},
						},
						"endpoint_connect": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Configure global settings for Endpoint Connect. These settings apply to all gateways.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable_password_caching": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "If the password entered to authenticate is saved locally on the user's machine.",
										Default:     "false",
									},
									"cache_password_timeout": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Cached password timeout (in minutes).",
										Default:     1440,
									},
									"re_auth_user_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The length of time (in minutes) until the user's credentials are resent to the gateway to verify authorization.",
										Default:     480,
									},
									"connect_mode": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Methods by which a connection to the gateway will be initiated:<br>Manual - VPN connections will not be initiated automatically.<br>Always connected - Endpoint Connect will automatically establish a connection to the last connected gateway under the following circumstances: (a) the device has a valid IP address, (b) when the device \"wakes up\" from a low-power state or a soft-reset, or (c) after a condition that caused the device to automatically disconnect ceases to exist (for example, Device is out of PC Sync, Disconnect is not idle.).<br>Configured on endpoint client - the method used for initiating a connection to a gateway is determined by the endpoint client.",
										Default:     "Configured On Endpoint Client",
									},
									"network_location_awareness": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Wide Impact: Also applies for Check Point GO clients!<br>Endpoint Connect intelligently detects whether it is inside or outside of the VPN domain (Enterprise LAN), and automatically connects or disconnects as required. Select true and edit network-location-awareness-conf to configure this capability.",
										Default:     "client_decide",
									},
									"network_location_awareness_conf": {
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "Configure how the client determines its location in relation to the internal network.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"vpn_clients_are_considered_inside_the_internal_network_when_the_client": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "When a VPN client is within the internal network, the internal resources are available and the VPN tunnel should be disconnected. Determine when VPN clients are considered inside the internal network:<br>Connects to GW through internal interface - The client connects to the gateway through one of its internal interfaces (recommended).<br>Connects from network or group - The client connects from a network or group specified in network-or-group-of-conn-vpn-client.<br>Runs on computer with access to Active Directory domain - The client runs on a computer that can access its Active Directory domain.<br>Note: The VPN tunnel will resume automatically when the VPN client is no longer in the internal network and the client is set to \"Always connected\" mode.",
												},
												"network_or_group_of_conn_vpn_client": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Name or UID of Network or Group the VPN client is connected from.<br>Available only if vpn-clients-are-considered-inside-the-internal-network-when-the-client is set to \"Connects from network or group\".",
												},
												"consider_wireless_networks_as_external": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "The speed at which locations are classified as internal or external can be increased by creating a list of wireless networks that are known to be external. A wireless network is identified by its Service Set Identifier (SSID) a name used to identify a particular 802.11 wireless LAN.",
												},
												"excluded_internal_wireless_networks": {
													Type:        schema.TypeSet,
													Optional:    true,
													Description: "Excludes the specified internal networks names (SSIDs).<br>Available only if consider-wireless-networks-as-external is set to true.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"consider_undefined_dns_suffixes_as_external": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "The speed at which locations are classified as internal or external can be increased by creating a list of DNS suffixes that are known to be external. Enable this to be able to define DNS suffixes which won't be considered external.",
												},
												"dns_suffixes": {
													Type:        schema.TypeSet,
													Optional:    true,
													Description: "DNS suffixes not defined here will be considered as external. If this list is empty consider-undefined-dns-suffixes-as-external will automatically be set to false.<br>Available only if consider-undefined-dns-suffixes-as-external is set to true.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"remember_previously_detected_external_networks": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "The speed at which locations are classified as internal or external can be increased by caching (on the client side) names of networks that were previously determined to be external.",
												},
											},
										},
									},
									"disconnect_when_conn_to_network_is_lost": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Enabling this feature disconnects users from the gateway when connectivity to the network is lost.",
										Default:     "client_decide",
									},
									"disconnect_when_device_is_idle": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Enabling this feature will disconnect users from the gateway if there is no traffic sent during the defined time period.",
										Default:     "client_decide",
									},
									"route_all_traffic_to_gw": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Operates the client in Hub Mode, sending all traffic to the VPN server for routing, filtering, and processing.",
										Default:     "false",
									},
									"client_upgrade_mode": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Select an option to determine how the client is upgraded.",
										Default:     "ask_user",
									},
								},
							},
						},
						"hot_spot_and_hotel_registration": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Configure the settings for Wireless Hot Spot and Hotel Internet access registration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable_registration": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Set Enable registration to true in order to configure settings. Set Enable registration to false in order to cancel registration (the configurations below won't be available). When the feature is enabled, you have several minutes to complete registration.",
									},
									"local_subnets_access_only": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Local subnets access only.",
									},
									"registration_timeout": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Maximum time (in seconds) to complete registration.",
										Default:     600,
									},
									"track_log": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Track log.",
									},
									"max_ip_access_during_registration": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Maximum number of addresses to allow access to during registration.",
										Default:     5,
									},
									"ports": {
										Type:        schema.TypeSet,
										Optional:    true,
										Description: "Ports to be opened during registration (up to 10 ports).",
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
			"user_directory": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "User can enable LDAP User Directory as well as specify global parameters for LDAP. If LDAP User Directory is enabled, this means that users are managed on an external LDAP server and not on the internal Check Point Security Gateway users databases.",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_password_change_when_user_active_directory_expires": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "For organizations using MS Active Directory, this setting enables users whose passwords have expired to automatically create new passwords.",
							Default:     true,
						},
						"cache_size": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The maximum number of cached users allowed. The cache is FIFO (first-in, first-out). When a new user is added to a full cache, the first user is deleted to make room for the new user. The Check Point Security Gateway does not query the LDAP server for users already in the cache, unless the cache has timed out.",
							Default:     1000,
						},
						"enable_password_expiration_configuration": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable configuring of the number of days during which the password is valid.<br>If enable-password-change-when-user-active-directory-expires is true, the password expiration time is determined by the Active Directory. In this case it is recommended not to set this to true.",
						},
						"password_expires_after": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Specifies the number of days during which the password is valid. Users are authenticated using a special LDAP password. Should this password expire, a new password must be defined.<br>Available only if enable-password-expiration-configuration is true.",
							Default:     90,
						},
						"timeout_on_cached_users": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The period of time in which a cached user is timed out and will need to be fetched again from the LDAP server.",
							Default:     900,
						},
						"display_user_dn_at_login": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Decide whether or not you would like to display the user's DN when logging in. If you choose to display the user DN, you can select whether to display it, when the user is prompted for the password at login, or on the request of the authentication scheme. This property is a useful diagnostic tool when there is more than one user with the same name in an Account Unit. In this case, the first one is chosen and the others are ignored.",
							Default:     "no display",
						},
						"enforce_rules_for_user_mgmt_admins": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enforces password strength rules on LDAP users when you create or modify a Check Point Password.",
						},
						"min_password_length": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Specifies the minimum length (in characters) of the password.",
							Default:     6,
						},
						"password_must_include_a_digit": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Password must include a digit.",
						},
						"password_must_include_a_symbol": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Password must include a symbol.",
						},
						"password_must_include_lowercase_char": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Password must include a lowercase character.",
						},
						"password_must_include_uppercase_char": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Password must include an uppercase character.",
						},
					},
				},
			},
			"qos": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Define the general parameters of Quality of Service (QoS) and apply them to QoS rules.",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default_weight_of_rule": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Define a Weight at which bandwidth will be guaranteed. Set a default weight for a rule.<br>Note: Value will be applied to new rules only.",
							Default:     10,
						},
						"max_weight_of_rule": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Define a Weight at which bandwidth will be guaranteed. Set a maximum weight for a rule.",
							Default:     1000,
						},
						"unit_of_measure": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Define the Rate at which packets are transmitted, for which bandwidth will be guaranteed. Set a Unit of measure.",
							Default:     "Kbits-per-sec",
						},
						"authenticated_ip_expiration": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Define the Authentication time-out for QoS. This timeout is set in minutes. In an Authenticated IP all connections which are open in a specified time limit will be guaranteed bandwidth, but will not be guaranteed bandwidth after the time limit.",
							Default:     15,
						},
						"non_authenticated_ip_expiration": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Define the Authentication time-out for QoS. This timeout is set in minutes.",
							Default:     5,
						},
						"unanswered_queried_ip_expiration": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Define the Authentication time-out for QoS. This timeout is set in minutes.",
							Default:     3,
						},
					},
				},
			},
			"carrier_security": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Specify system-wide properties. Select GTP intra tunnel inspection options, including anti-spoofing; tracking and logging options, and integrity tests.",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"block_gtp_in_gtp": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Prevents GTP packets from being encapsulated inside GTP tunnels. When this option is checked, such packets are dropped and logged.",
							Default:     true,
						},
						"enforce_gtp_anti_spoofing": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "verifies that G-PDUs are using the end user IP address that has been agreed upon in the PDP context activation process. When this option is checked, packets that do not use this IP address are dropped and logged.",
							Default:     true,
						},
						"produce_extended_logs_on_unmatched_pdus": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "logs GTP packets not matched by previous rules with Carrier Security's extended GTP-related log fields. These logs are brown and their Action attribute is empty. The default setting is checked.",
							Default:     false,
						},
						"produce_extended_logs_on_unmatched_pdus_position": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Choose to place this implicit rule Before Last or as the Last rule.<br>Available only if produce-extended-logs-on-unmatched-pdus is true.",
							Default:     "before last",
						},
						"protocol_violation_track_option": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Set the appropriate track or alert option to be used when a protocol violation (malformed packet) is detected.",
							Default:     "log",
						},
						"enable_g_pdu_seq_number_check_with_max_deviation": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "If set to false, sequence checking is not enforced and all out-of-sequence G-PDUs will be accepted.<br>To enhance performance, disable this extended integrity test.",
							Default:     false,
						},
						"g_pdu_seq_number_check_max_deviation": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "specifies that a G-PDU is accepted only if the difference between its sequence number and the expected sequence number is less than or equal to the allowed deviation.<br>Available only ifenable-g-pdu-seq-number-check-with-max-deviation is true.",
							Default:     16,
						},
						"verify_flow_labels": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "See that each packet's flow label matches the flow labels defined by GTP signaling. This option is relevant for GTP version 0 only.<br>To enhance performance, disable this extended integrity test.",
							Default:     true,
						},
						"allow_ggsn_replies_from_multiple_interfaces": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Allows GTP signaling replies from an IP address different from the IP address to which the requests are sent (Relevant only for gateways below R80).",
							Default:     true,
						},
						"enable_reverse_connections": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Allows Carrier Security gateways to accept PDUs sent from the GGSN to the SGSN, on a previously established PDP context, even if these PDUs are sent over ports that do not match the ports of the established PDP context.",
							Default:     true,
						},
						"gtp_signaling_rate_limit_sampling_interval": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Works in correlation with the property Enforce GTP Signal packet rate limit found in the Carrier Security window of the GSN network object. For example, with the rate limit sampling interval default of 1 second, and the network object enforced a GTP signal packet rate limit of the default 2048 PDU per second, sampling will occur one time per second, or 2048 signaling PDUs between two consecutive samplings.",
							Default:     1,
						},
						"one_gtp_echo_on_each_path_frequency": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "sets the number of GTP Echo exchanges per path allowed per configured time period. Echo requests exceeding this rate are dropped and logged. Setting the value to 0 disables the feature and allows an unlimited number of echo requests per path at any interval.",
							Default:     5,
						},
						"aggressive_aging": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "If true, enables configuring aggressive aging thresholds and time out value.",
							Default:     false,
						},
						"aggressive_timeout": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Aggressive timeout. Available only if aggressive-aging is true.",
							Default:     3600,
						},
						"memory_activation_threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Memory activation threshold. Available only if aggressive-aging is true.",
							Default:     80,
						},
						"memory_deactivation_threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Memory deactivation threshold. Available only if aggressive-aging is true.",
							Default:     60,
						},
						"tunnel_activation_threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Tunnel activation threshold. Available only if aggressive-aging is true.",
							Default:     80,
						},
						"tunnel_deactivation_threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Tunnel deactivation threshold. Available only if aggressive-aging is true.",
							Default:     60,
						},
					},
				},
			},
			"user_accounts": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Set the expiration for a user account and configure \"about to expire\" warnings.",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"expiration_date_method": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Select an Expiration Date Method.<br>Expire at - Account expires on the date that you select.<br>Expire after - Account expires after the number of days that you select.",
						},
						"expiration_date": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specify an Expiration Date in the following format: YYYY-MM-DD.<br>Available only if expiration-date-method is set to \"expire at\".",
						},
						"days_until_expiration": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Account expires after the number of days that you select.<br>Available only if expiration-date-method is set to \"expire after\".",
						},
						"show_accounts_expiration_indication_days_in_advance": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Activates the Expired Accounts link, to open the Expired Accounts window.",
						},
					},
				},
			},
			"user_authority": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Decide whether to display and access the WebAccess rule base. This policy defines which users (that is, which Windows Domains) have access to the internal sites of the organization.",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_web_access_view": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Specify whether or not to display the WebAccess rule base. This rule base is used for UserAuthority.",
						},
						"windows_domains_to_trust": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "When matching Firewall usernames to Windows Domains usernames for Single Sign on, selectwhether to trust all or specify which Windows Domain should be trusted.<br>ALL - Enables you to allow all Windows domains to access the internal sites of the organization.<br>SELECTIVELY - Enables you to specify which Windows domains will have access to the internal sites of the organization.",
							Default:     "all",
						},
						"trust_only_following_windows_domains": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Specify which Windows domains will have access to the internal sites of the organization.<br>Available only if windows-domains-to-trust is set to SELECTIVELY.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"connect_control": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Configure settings that relate to ConnectControl server load balancing.",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"load_agents_port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Sets the port number on which load measuring agents communicate with ConnectControl.",
							Default:     18212,
						},
						"load_measurement_interval": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "sets how often (in seconds) the load measuring agents report their load status to ConnectControl.",
							Default:     20,
						},
						"persistence_server_timeout": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Sets the amount of time (in seconds) that a client, once directed to a particular server, will continue to be directed to that same server.",
							Default:     1800,
						},
						"server_availability_check_interval": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Sets how often (in seconds) ConnectControl checks to make sure the load balanced servers are running and responding to service requests.",
							Default:     20,
						},
						"server_check_retries": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Sets how many times ConnectControl attempts to contact a server before ceasing to direct traffic to it.",
							Default:     3,
						},
					},
				},
			},
			"stateful_inspection": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Adjust Stateful Inspection parameters.",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tcp_start_timeout": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "A TCP connection will be timed out if the interval between the arrival of the first packet and establishment of the connection (TCP three-way handshake) exceeds TCP start timeout seconds.",
							Default:     25,
						},
						"tcp_session_timeout": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The length of time (in seconds) an idle connection will remain in the Security Gateway connections table.",
							Default:     3600,
						},
						"tcp_end_timeout": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "A TCP connection will only terminate TCP end timeout seconds after two FIN packets (one in each direction: client-to-server, and server-to-client) or an RST packet. When a TCP connection ends (FIN packets sent or connection reset) the Check Point Security Gateway will keep the connection in the connections table for another TCP end timeout seconds, to allow for stray ACKs of the connection that arrive late.",
							Default:     20,
						},
						"tcp_end_timeout_r8020_gw_and_above": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "A TCP connection will only terminate TCP end timeout seconds after two FIN packets (one in each direction: client-to-server, and server-to-client) or an RST packet. When a TCP connection ends (FIN packets sent or connection reset) the Check Point Security Gateway will keep the connection in the connections table for another TCP end timeout seconds, to allow for stray ACKs of the connection that arrive late.",
							Default:     5,
						},
						"udp_virtual_session_timeout": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Specifies the amount of time (in seconds) a UDP reply channel may remain open without any packets being returned.",
							Default:     40,
						},
						"icmp_virtual_session_timeout": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "An ICMP virtual session will be considered to have timed out after this time period (in seconds).",
							Default:     30,
						},
						"other_ip_protocols_virtual_session_timeout": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "A virtual session of services which are not explicitly configured here will be considered to have timed out after this time period (in seconds).",
							Default:     60,
						},
						"sctp_start_timeout": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "SCTP connections will be timed out if the interval between the arrival of the first packet and establishment of the connection exceeds this value (in seconds).",
							Default:     30,
						},
						"sctp_session_timeout": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Time (in seconds) an idle connection will remain in the Security Gateway connections table.",
							Default:     3600,
						},
						"sctp_end_timeout": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "SCTP connections end after this number of seconds, after the connection ends or is reset, to allow for stray ACKs of the connection that arrive late.",
							Default:     20,
						},
						"accept_stateful_udp_replies_for_unknown_services": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Specifies if UDP replies are to be accepted for unknown services.",
						},
						"accept_stateful_icmp_errors": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Accept ICMP error packets which refer to another non-ICMP connection (for example, to an ongoing TCP or UDP connection) that was accepted by the Rule Base.",
							Default:     true,
						},
						"accept_stateful_icmp_replies": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Accept ICMP reply packets for ICMP requests that were accepted by the Rule Base.",
							Default:     true,
						},
						"accept_stateful_other_ip_protocols_replies_for_unknown_services": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Accept reply packets for other undefined services (that is, services which are not one of the following: TCP, UDP, ICMP).",
						},
						"drop_out_of_state_tcp_packets": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Drop TCP packets which are not consistent with the current state of the connection.",
						},
						"log_on_drop_out_of_state_tcp_packets": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Generates a log entry when these out of state TCP packets are dropped.<br>Available only if drop-out-of-state-tcp-packets is true.",
							Default:     true,
						},
						"tcp_out_of_state_drop_exceptions": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Name or uid of the gateways and clusters for which Out of State packets are allowed.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"drop_out_of_state_icmp_packets": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Drop ICMP packets which are not consistent with the current state of the connection.",
							Default:     true,
						},
						"log_on_drop_out_of_state_icmp_packets": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Generates a log entry when these out of state ICMP packets are dropped.<br>Available only if drop-out-of-state-icmp-packets is true.",
							Default:     true,
						},
						"drop_out_of_state_sctp_packets": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Drop SCTP packets which are not consistent with the current state of the connection.",
							Default:     true,
						},
						"log_on_drop_out_of_state_sctp_packets": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Generates a log entry when these out of state SCTP packets are dropped.<br>Available only if drop-out-of-state-sctp-packets is true.",
							Default:     true,
						},
					},
				},
			},
			"log_and_alert": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Define system-wide logging and alerting parameters.",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"administrative_notifications": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Administrative notifications specifies the action to be taken when an administrative event (for example, when a certificate is about to expire) occurs.",
							Default:     "Log",
						},
						"connection_matched_by_sam": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Connection matched by SAM specifies the action to be taken when a connection is blocked by SAM (Suspicious Activities Monitoring).",
							Default:     "Popup Alert",
						},
						"dynamic_object_resolution_failure": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Dynamic object resolution failure specifies the action to be taken when a dynamic object cannot be resolved.",
						},
						"packet_is_incorrectly_tagged": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Packet is incorrectly tagged.",
							Default:     "Log",
						},
						"packet_tagging_brute_force_attack": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Packet tagging brute force attack.",
							Default:     "Popup Alert",
						},
						"sla_violation": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "SLA violation specifies the action to be taken when an SLA violation occurs, as defined in the Virtual Links window.",
							Default:     "None",
						},
						"vpn_conf_and_key_exchange_errors": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "VPN configuration & key exchange errors specifies the action to be taken when logging configuration or key exchange errors occur, for example, when attempting to establish encrypted communication with a network object inside the same encryption domain.",
							Default:     "Log",
						},
						"vpn_packet_handling_error": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "VPN packet handling errors specifies the action to be taken when encryption or decryption errors occurs. A log entry contains the action performed (Drop or Reject) and a short description of the error cause, for example, scheme or method mismatch.",
							Default:     "Log",
						},
						"vpn_successful_key_exchange": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "VPN successful key exchange specifies the action to be taken when VPN keys are successfully exchanged.",
							Default:     "Log",
						},
						"log_every_authenticated_http_connection": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Log every authenticated HTTP connection specifies that a log entry should be generated for every authenticated HTTP connection.",
						},
						"log_traffic": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Log Traffic specifies whether or not to log traffic.",
							Default:     "Log",
						},
						"alerts": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Define the behavior of alert logs and the type of alert used for System Alert logs.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"send_popup_alert_to_smartview_monitor": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Send popup alert to SmartView Monitor when an alert is issued, it is also sent to SmartView Monitor.",
										Default:     true,
									},
									"popup_alert_script": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Run popup alert script the operating system script to be executed when an alert is issued. For example, set another form of notification, such as an email or a user-defined command.",
									},
									"send_mail_alert_to_smartview_monitor": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Send mail alert to SmartView Monitor when a mail alert is issued, it is also sent to SmartView Monitor.",
									},
									"mail_alert_script": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Run mail alert script the operating system script to be executed when Mail is specified as the Track in a rule. The default is internal_sendmail, which is not a script but an internal Security Gateway command.",
										Default:     "internal_sendmail -s alert -t mailer root",
									},
									"send_snmp_trap_alert_to_smartview_monitor": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Send SNMP trap alert to SmartView Monitor when an SNMP trap alert is issued, it is also sent to SmartView Monitor.",
									},
									"snmp_trap_alert_script": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Run SNMP trap alert script command to be executed when SNMP Trap is specified as the Track in a rule. By default the internal_snmp_trap is used. This command is executed by the fwd process.",
										Default:     "internal_snmp_trap localhost",
									},
									"send_user_defined_alert_num1_to_smartview_monitor": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Send user defined alert no. 1 to SmartView Monitor when an alert is issued, it is also sent to SmartView Monitor.",
										Default:     true,
									},
									"user_defined_script_num1": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Run user defined script the operating system script to be run when User-Defined is specified as the Track in a rule, or when User Defined Alert no. 1 is selected as a Track Option.",
									},
									"send_user_defined_alert_num2_to_smartview_monitor": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Send user defined alert no. 2 to SmartView Monitor when an alert is issued, it is also sent to SmartView Monitor.",
										Default:     true,
									},
									"user_defined_script_num2": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Run user defined 2 script the operating system script to be run when User-Defined is specified as the Track in a rule, or when User Defined Alert no. 2 is selected as a Track Option.",
									},
									"send_user_defined_alert_num3_to_smartview_monitor": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Send user defined alert no. 3 to SmartView Monitor when an alert is issued, it is also sent to SmartView Monitor.",
										Default:     true,
									},
									"user_defined_script_num3": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Run user defined 3 script the operating system script to be run when User-Defined is specified as the Track in a rule, or when User Defined Alert no. 3 is selected as a Track Option.",
									},
									"default_track_option_for_system_alerts": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Set the default track option for System Alerts.",
										Default:     "Popup Alert",
									},
								},
							},
						},
						"time_settings": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Configure the time settings associated with system-wide logging and alerting parameters.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"excessive_log_grace_period": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Specifies the minimum amount of time (in seconds) between consecutive logs of similar packets. Two packets are considered similar if they have the same source address, source port, destination address, and destination port; and the same protocol was used. After the first packet, similar packets encountered in the grace period will be acted upon according to the security policy, but only the first packet generates a log entry or an alert. Any value from 0 to 90 seconds can be entered in this field.<br>Note: This option only applies for DROP rules with logging.",
										Default:     62,
									},
									"logs_resolving_timeout": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Specifies the amount of time (in seconds), after which the log page is displayed without resolving names and while showing only IP addresses. Any value from 0 to 90 seconds can be entered in this field.",
										Default:     20,
									},
									"status_fetching_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Specifies the frequency at which the Security Management server queries the Check Point Security gateway, Check Point QoS and other gateways it manages for status information. Any value from 30 to 900 seconds can be entered in this field.",
										Default:     60,
									},
									"virtual_link_statistics_logging_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Specifies the frequency (in seconds) with which Virtual Link statistics will be logged. This parameter is relevant only for Virtual Links defined with SmartView Monitor statistics enabled in the SLA Parameters tab of the Virtual Link window. Any value from 60 to 3600 seconds can be entered in this field.",
										Default:     60,
									},
								},
							},
						},
					},
				},
			},
			"data_access_control": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Configure automatic downloads from Check Point and anonymously share product data. Options selected here apply to all Security Gateways, Clusters and VSX devices managed by this management server.",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_download_important_data": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Automatically download and install Software Blade Contracts, security updates and other important data (highly recommended).",
							Default:     true,
						},
						"auto_download_sw_updates_and_new_features": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Automatically download software updates and new features (highly recommended).<br>Available only if auto-download-important-data is set to true.",
							Default:     true,
						},
						"send_anonymous_info": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Help Check Point improve the product by sending anonymous information.",
							Default:     true,
						},
						"share_sensitive_info": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Approve sharing core dump files and other relevant crash data which might contain personal information. All shared data will be processed in accordance with Check Point's Privacy Policy.<br>Available only if send-anonymous-info is set to true.",
							Default:     false,
						},
					},
				},
			},
			"non_unique_ip_address_ranges": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specify Non Unique IP Address Ranges.",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The type of the IP Address.",
						},
						"first_ipv4_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The first IPV4 Address in the range.",
						},
						"first_ipv6_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The first IPV6 Address in the range.",
						},
						"last_ipv4_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The last IPV4 Address in the range.",
						},
						"last_ipv6_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The last IPV6 Address in the range.",
						},
					},
				},
			},
			"proxy": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Select whether a proxy server is used when servers, gateways, or clients need to access the internet for certain Check Point features and set the default proxy server that will be used.",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"use_proxy_server": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "If set to true, a proxy server is used when features need to access the internet.",
							Default:     false,
						},
						"proxy_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specify the URL or IP address of the proxy server.<br>Available only if use-proxy-server is set to true.",
						},
						"proxy_port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Specify the Port on which the server will be accessed.<br>Available only if use-proxy-server is set to true.",
							Default:     80,
						},
					},
				},
			},
			"user_check": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Set a language for the UserCheck message if the language setting in the user's browser cannot be determined.",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"preferred_language": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The preferred language for new UserCheck message.",
							Default:     "English",
						},
						"send_emails_using_mail_server": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Name or UID of mail server to send emails to.",
						},
					},
				},
			},
			"hit_count": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Enable the Hit Count feature that tracks the number of connections that each rule matches.",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_hit_count": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Select to enable or clear to disable all Security Gateways to monitor the number of connections each rule matches.",
							Default:     true,
						},
						"keep_hit_count_data_up_to": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Select one of the time range options. Data is kept in the Security Management Server database for this period and is shown in the Hits column.",
							Default:     "3 Months",
						},
					},
				},
			},
			"advanced_conf": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Configure advanced global attributes. It's highly recommended to consult with Check Point's Technical Support before modifying these values.",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certs_and_pki": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Configure Certificates and PKI properties.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cert_validation_enforce_key_size": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Enforce key length in certificate validation (R80+ gateways only).",
										Default:     "off",
									},
									"host_certs_ecdsa_key_size": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Select the key size for ECDSA of the host certificate.",
										Default:     "P-256",
									},
									"host_certs_key_size": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Select the key size of the host certificate.",
										Default:     "2048",
									},
								},
							},
						},
					},
				},
			},
			"allow_remote_registration_of_opsec_products": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "After installing an OPSEC application, the remote administration (RA) utility enables an OPSEC product to finish registering itself without having to access the SmartConsole. If set to true, any host including the application host can run the utility. Otherwise,  the RA utility can only be run from the Security Management host.",
			},
			"num_spoofing_errs_that_trigger_brute_force": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "Indicates how many incorrectly signed packets will be tolerated before assuming that there is an attack on the packet tagging and revoking the client's key.",
			},
			"domains_to_process": {
				Type:        schema.TypeSet,
				Optional:    true,
				ForceNew:    true,
				Description: "Indicates which domains to process the commands on. It cannot be used with the details-level full, must be run from the System Domain only and with ignore-warnings true. Valid values are: CURRENT_DOMAIN, ALL_DOMAINS_ON_THIS_SERVER.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ignore_warnings": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Apply changes ignoring warnings.",
			},
			"ignore_errors": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.",
			},
		},
	}
}

func createManagementSetGlobalProperties(d *schema.ResourceData, m interface{}) error {
	return readManagementSetGlobalProperties(d, m)
}

func readManagementSetGlobalProperties(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if _, ok := d.GetOk("firewall"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("firewall.accept_control_connections"); ok {
			res["accept-control-connections"] = v
		}
		if v, ok := d.GetOk("firewall.accept_ips1_management_connections"); ok {
			res["accept-ips1-management-connections"] = v
		}
		if v, ok := d.GetOk("firewall.accept_remote_access_control_connections"); ok {
			res["accept-remote-access-control-connections"] = v
		}
		if v, ok := d.GetOk("firewall.accept_smart_update_connections"); ok {
			res["accept-smart-update-connections"] = v
		}
		if v, ok := d.GetOk("firewall.accept_outgoing_packets_originating_from_gw"); ok {
			res["accept-outgoing-packets-originating-from-gw"] = v
		}
		if v, ok := d.GetOk("firewall.accept_outgoing_packets_originating_from_gw_position"); ok {
			res["accept-outgoing-packets-originating-from-gw-position"] = v.(string)
		}
		if v, ok := d.GetOk("firewall.accept_outgoing_packets_originating_from_connectra_gw"); ok {
			res["accept-outgoing-packets-originating-from-connectra-gw"] = v
		}
		if v, ok := d.GetOk("firewall.accept_outgoing_packets_to_cp_online_services"); ok {
			res["accept-outgoing-packets-to-cp-online-services"] = v
		}
		if v, ok := d.GetOk("firewall.accept_outgoing_packets_to_cp_online_services_position"); ok {
			res["accept-outgoing-packets-to-cp-online-services-position"] = v.(string)
		}
		if v, ok := d.GetOk("firewall.accept_domain_name_over_tcp"); ok {
			res["accept-domain-name-over-tcp"] = v
		}
		if v, ok := d.GetOk("firewall.accept_domain_name_over_tcp_position"); ok {
			res["accept-domain-name-over-tcp-position"] = v.(string)
		}
		if v, ok := d.GetOk("firewall.accept_domain_name_over_udp"); ok {
			res["accept-domain-name-over-udp"] = v
		}
		if v, ok := d.GetOk("firewall.accept_domain_name_over_udp_position"); ok {
			res["accept-domain-name-over-udp-position"] = v.(string)
		}
		if v, ok := d.GetOk("firewall.accept_dynamic_addr_modules_outgoing_internet_connections"); ok {
			res["accept-dynamic-addr-modules-outgoing-internet-connections"] = v
		}
		if v, ok := d.GetOk("firewall.accept_icmp_requests"); ok {
			res["accept-icmp-requests"] = v
		}
		if v, ok := d.GetOk("firewall.accept_icmp_requests_position"); ok {
			res["accept-icmp-requests-position"] = v.(string)
		}
		if v, ok := d.GetOk("firewall.accept_identity_awareness_control_connections"); ok {
			res["accept-identity-awareness-control-connections"] = v
		}
		if v, ok := d.GetOk("firewall.accept_identity_awareness_control_connections_position"); ok {
			res["accept-identity-awareness-control-connections-position"] = v.(string)
		}
		if v, ok := d.GetOk("firewall.accept_incoming_traffic_to_dhcp_and_dns_services_of_gws"); ok {
			res["accept-incoming-traffic-to-dhcp-and-dns-services-of-gws"] = v
		}
		if v, ok := d.GetOk("firewall.accept_rip"); ok {
			res["accept-rip"] = v
		}
		if v, ok := d.GetOk("firewall.accept_rip_position"); ok {
			res["accept-rip-position"] = v.(string)
		}
		if v, ok := d.GetOk("firewall.accept_vrrp_packets_originating_from_cluster_members"); ok {
			res["accept-vrrp-packets-originating-from-cluster-members"] = v
		}
		if v, ok := d.GetOk("firewall.accept_web_and_ssh_connections_for_gw_administration"); ok {
			res["accept-web-and-ssh-connections-for-gw-administration"] = v
		}
		if v, ok := d.GetOk("firewall.log_implied_rules"); ok {
			res["log-implied-rules"] = v
		}
		if v, ok := d.GetOk("firewall.security_server"); ok {
			res["security-server"] = v
		}
		payload["firewall"] = res
	}

	if _, ok := d.GetOk("nat"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("nat.allow_bi_directional_nat"); ok {
			res["allow-bi-directional-nat"] = v
		}
		if v, ok := d.GetOk("nat.auto_arp_conf"); ok {
			res["auto-arp-conf"] = v
		}
		if v, ok := d.GetOk("nat.merge_manual_proxy_arp_conf"); ok {
			res["merge-manual-proxy-arp-conf"] = v
		}
		if v, ok := d.GetOk("nat.auto_translate_dest_on_client_side"); ok {
			res["auto-translate-dest-on-client-side"] = v
		}
		if v, ok := d.GetOk("nat.manually_translate_dest_on_client_side"); ok {
			res["manually-translate-dest-on-client-side"] = v
		}
		if v, ok := d.GetOk("nat.enable_ip_pool_nat"); ok {
			res["enable-ip-pool-nat"] = v
		}
		if v, ok := d.GetOk("nat.addr_alloc_and_release_track"); ok {
			res["addr-alloc-and-release-track"] = v.(string)
		}
		if v, ok := d.GetOk("nat.addr_exhaustion_track"); ok {
			res["addr-exhaustion-track"] = v.(string)
		}
		payload["nat"] = res
	}

	if _, ok := d.GetOk("authentication"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("authentication.auth_internal_users_with_specific_suffix"); ok {
			res["auth-internal-users-with-specific-suffix"] = v
		}
		if v, ok := d.GetOk("authentication.allowed_suffix_for_internal_users"); ok {
			res["allowed-suffix-for-internal-users"] = v.(string)
		}
		if v, ok := d.GetOk("authentication.max_days_before_expiration_of_non_pulled_user_certificates"); ok {
			res["max-days-before-expiration-of-non-pulled-user-certificates"] = v
		}
		if v, ok := d.GetOk("authentication.max_client_auth_attempts_before_connection_termination"); ok {
			res["max-client-auth-attempts-before-connection-termination"] = v
		}
		if v, ok := d.GetOk("authentication.max_rlogin_attempts_before_connection_termination"); ok {
			res["max-rlogin-attempts-before-connection-termination"] = v
		}
		if v, ok := d.GetOk("authentication.max_session_auth_attempts_before_connection_termination"); ok {
			res["max-session-auth-attempts-before-connection-termination"] = v
		}
		if v, ok := d.GetOk("authentication.max_telnet_attempts_before_connection_termination"); ok {
			res["max-telnet-attempts-before-connection-termination"] = v
		}
		if v, ok := d.GetOk("authentication.enable_delayed_auth"); ok {
			res["enable-delayed-auth"] = v
		}
		if v, ok := d.GetOk("authentication.delay_each_auth_attempt_by"); ok {
			res["delay-each-auth-attempt-by"] = v
		}
		payload["authentication"] = res
	}

	if _, ok := d.GetOk("vpn"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("vpn.vpn_conf_method"); ok {
			res["vpn-conf-method"] = v.(string)
		}
		if v, ok := d.GetOk("vpn.domain_name_for_dns_resolving"); ok {
			res["domain-name-for-dns-resolving"] = v.(string)
		}
		if v, ok := d.GetOk("vpn.enable_backup_gw"); ok {
			res["enable-backup-gw"] = v
		}
		if v, ok := d.GetOk("vpn.enable_decrypt_on_accept_for_gw_to_gw_traffic"); ok {
			res["enable-decrypt-on-accept-for-gw-to-gw-traffic"] = v
		}
		if v, ok := d.GetOk("vpn.enable_load_distribution_for_mep_conf"); ok {
			res["enable-load-distribution-for-mep-conf"] = v
		}
		if v, ok := d.GetOk("vpn.enable_vpn_directional_match_in_vpn_column"); ok {
			res["enable-vpn-directional-match-in-vpn-column"] = v
		}
		if v, ok := d.GetOk("vpn.grace_period_after_the_crl_is_not_valid"); ok {
			res["grace-period-after-the-crl-is-not-valid"] = v
		}
		if v, ok := d.GetOk("vpn.grace_period_before_the_crl_is_valid"); ok {
			res["grace-period-before-the-crl-is-valid"] = v
		}
		if v, ok := d.GetOk("vpn.grace_period_extension_for_secure_remote_secure_client"); ok {
			res["grace-period-extension-for-secure-remote-secure-client"] = v
		}
		if v, ok := d.GetOk("vpn.support_ike_dos_protection_from_identified_src"); ok {
			res["support-ike-dos-protection-from-identified-src"] = v.(string)
		}
		if v, ok := d.GetOk("vpn.support_ike_dos_protection_from_unidentified_src"); ok {
			res["support-ike-dos-protection-from-unidentified-src"] = v.(string)
		}
		payload["vpn"] = res
	}

	if _, ok := d.GetOk("remote_access"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("remote_access.enable_back_connections"); ok {
			res["enable-back-connections"] = v
		}
		if v, ok := d.GetOk("remote_access.keep_alive_packet_to_gw_interval"); ok {
			res["keep-alive-packet-to-gw-interval"] = v
		}
		if v, ok := d.GetOk("remote_access.encrypt_dns_traffic"); ok {
			res["encrypt-dns-traffic"] = v
		}
		if v, ok := d.GetOk("remote_access.simultaneous_login_mode"); ok {
			res["simultaneous-login-mode"] = v.(string)
		}
		if v, ok := d.GetOk("remote_access.vpn_authentication_and_encryption"); ok {
			res["vpn-authentication-and-encryption"] = v
		}
		if v, ok := d.GetOk("remote_access.vpn_advanced"); ok {
			res["vpn-advanced"] = v
		}
		if v, ok := d.GetOk("remote_access.scv"); ok {
			res["scv"] = v
		}
		if v, ok := d.GetOk("remote_access.ssl_network_extender"); ok {
			res["ssl-network-extender"] = v
		}
		if v, ok := d.GetOk("remote_access.secure_client_mobile"); ok {
			res["secure-client-mobile"] = v
		}
		if v, ok := d.GetOk("remote_access.endpoint_connect"); ok {
			res["endpoint-connect"] = v
		}
		if v, ok := d.GetOk("remote_access.hot_spot_and_hotel_registration"); ok {
			res["hot-spot-and-hotel-registration"] = v
		}
		payload["remote-access"] = res
	}

	if _, ok := d.GetOk("user_directory"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("user_directory.enable_password_change_when_user_active_directory_expires"); ok {
			res["enable-password-change-when-user-active-directory-expires"] = v
		}
		if v, ok := d.GetOk("user_directory.cache_size"); ok {
			res["cache-size"] = v
		}
		if v, ok := d.GetOk("user_directory.enable_password_expiration_configuration"); ok {
			res["enable-password-expiration-configuration"] = v
		}
		if v, ok := d.GetOk("user_directory.password_expires_after"); ok {
			res["password-expires-after"] = v
		}
		if v, ok := d.GetOk("user_directory.timeout_on_cached_users"); ok {
			res["timeout-on-cached-users"] = v
		}
		if v, ok := d.GetOk("user_directory.display_user_dn_at_login"); ok {
			res["display-user-dn-at-login"] = v.(string)
		}
		if v, ok := d.GetOk("user_directory.enforce_rules_for_user_mgmt_admins"); ok {
			res["enforce-rules-for-user-mgmt-admins"] = v
		}
		if v, ok := d.GetOk("user_directory.min_password_length"); ok {
			res["min-password-length"] = v
		}
		if v, ok := d.GetOk("user_directory.password_must_include_a_digit"); ok {
			res["password-must-include-a-digit"] = v
		}
		if v, ok := d.GetOk("user_directory.password_must_include_a_symbol"); ok {
			res["password-must-include-a-symbol"] = v
		}
		if v, ok := d.GetOk("user_directory.password_must_include_lowercase_char"); ok {
			res["password-must-include-lowercase-char"] = v
		}
		if v, ok := d.GetOk("user_directory.password_must_include_uppercase_char"); ok {
			res["password-must-include-uppercase-char"] = v
		}
		payload["user-directory"] = res
	}

	if _, ok := d.GetOk("qos"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("qos.default_weight_of_rule"); ok {
			res["default-weight-of-rule"] = v
		}
		if v, ok := d.GetOk("qos.max_weight_of_rule"); ok {
			res["max-weight-of-rule"] = v
		}
		if v, ok := d.GetOk("qos.unit_of_measure"); ok {
			res["unit-of-measure"] = v.(string)
		}
		if v, ok := d.GetOk("qos.authenticated_ip_expiration"); ok {
			res["authenticated-ip-expiration"] = v
		}
		if v, ok := d.GetOk("qos.non_authenticated_ip_expiration"); ok {
			res["non-authenticated-ip-expiration"] = v
		}
		if v, ok := d.GetOk("qos.unanswered_queried_ip_expiration"); ok {
			res["unanswered-queried-ip-expiration"] = v
		}
		payload["qos"] = res
	}

	if _, ok := d.GetOk("carrier_security"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("carrier_security.block_gtp_in_gtp"); ok {
			res["block-gtp-in-gtp"] = v
		}
		if v, ok := d.GetOk("carrier_security.enforce_gtp_anti_spoofing"); ok {
			res["enforce-gtp-anti-spoofing"] = v
		}
		if v, ok := d.GetOk("carrier_security.produce_extended_logs_on_unmatched_pdus"); ok {
			res["produce-extended-logs-on-unmatched-pdus"] = v
		}
		if v, ok := d.GetOk("carrier_security.produce_extended_logs_on_unmatched_pdus_position"); ok {
			res["produce-extended-logs-on-unmatched-pdus-position"] = v.(string)
		}
		if v, ok := d.GetOk("carrier_security.protocol_violation_track_option"); ok {
			res["protocol-violation-track-option"] = v.(string)
		}
		if v, ok := d.GetOk("carrier_security.enable_g_pdu_seq_number_check_with_max_deviation"); ok {
			res["enable-g-pdu-seq-number-check-with-max-deviation"] = v
		}
		if v, ok := d.GetOk("carrier_security.g_pdu_seq_number_check_max_deviation"); ok {
			res["g-pdu-seq-number-check-max-deviation"] = v
		}
		if v, ok := d.GetOk("carrier_security.verify_flow_labels"); ok {
			res["verify-flow-labels"] = v
		}
		if v, ok := d.GetOk("carrier_security.allow_ggsn_replies_from_multiple_interfaces"); ok {
			res["allow-ggsn-replies-from-multiple-interfaces"] = v
		}
		if v, ok := d.GetOk("carrier_security.enable_reverse_connections"); ok {
			res["enable-reverse-connections"] = v
		}
		if v, ok := d.GetOk("carrier_security.gtp_signaling_rate_limit_sampling_interval"); ok {
			res["gtp-signaling-rate-limit-sampling-interval"] = v
		}
		if v, ok := d.GetOk("carrier_security.one_gtp_echo_on_each_path_frequency"); ok {
			res["one-gtp-echo-on-each-path-frequency"] = v
		}
		if v, ok := d.GetOk("carrier_security.aggressive_aging"); ok {
			res["aggressive-aging"] = v
		}
		if v, ok := d.GetOk("carrier_security.aggressive_timeout"); ok {
			res["aggressive-timeout"] = v
		}
		if v, ok := d.GetOk("carrier_security.memory_activation_threshold"); ok {
			res["memory-activation-threshold"] = v
		}
		if v, ok := d.GetOk("carrier_security.memory_deactivation_threshold"); ok {
			res["memory-deactivation-threshold"] = v
		}
		if v, ok := d.GetOk("carrier_security.tunnel_activation_threshold"); ok {
			res["tunnel-activation-threshold"] = v
		}
		if v, ok := d.GetOk("carrier_security.tunnel_deactivation_threshold"); ok {
			res["tunnel-deactivation-threshold"] = v
		}
		payload["carrier-security"] = res
	}

	if _, ok := d.GetOk("user_accounts"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("user_accounts.expiration_date_method"); ok {
			res["expiration-date-method"] = v.(string)
		}
		if v, ok := d.GetOk("user_accounts.expiration_date"); ok {
			res["expiration-date"] = v.(string)
		}
		if v, ok := d.GetOk("user_accounts.days_until_expiration"); ok {
			res["days-until-expiration"] = v
		}
		if v, ok := d.GetOk("user_accounts.show_accounts_expiration_indication_days_in_advance"); ok {
			res["show-accounts-expiration-indication-days-in-advance"] = v
		}
		payload["user-accounts"] = res
	}

	if _, ok := d.GetOk("user_authority"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("user_authority.display_web_access_view"); ok {
			res["display-web-access-view"] = v
		}
		if v, ok := d.GetOk("user_authority.windows_domains_to_trust"); ok {
			res["windows-domains-to-trust"] = v.(string)
		}
		if v, ok := d.GetOk("user_authority.trust_only_following_windows_domains"); ok {
			res["trust-only-following-windows-domains"] = v
		}
		payload["user-authority"] = res
	}

	if _, ok := d.GetOk("connect_control"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("connect_control.load_agents_port"); ok {
			res["load-agents-port"] = v
		}
		if v, ok := d.GetOk("connect_control.load_measurement_interval"); ok {
			res["load-measurement-interval"] = v
		}
		if v, ok := d.GetOk("connect_control.persistence_server_timeout"); ok {
			res["persistence-server-timeout"] = v
		}
		if v, ok := d.GetOk("connect_control.server_availability_check_interval"); ok {
			res["server-availability-check-interval"] = v
		}
		if v, ok := d.GetOk("connect_control.server_check_retries"); ok {
			res["server-check-retries"] = v
		}
		payload["connect-control"] = res
	}

	if _, ok := d.GetOk("stateful_inspection"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("stateful_inspection.tcp_start_timeout"); ok {
			res["tcp-start-timeout"] = v
		}
		if v, ok := d.GetOk("stateful_inspection.tcp_session_timeout"); ok {
			res["tcp-session-timeout"] = v
		}
		if v, ok := d.GetOk("stateful_inspection.tcp_end_timeout"); ok {
			res["tcp-end-timeout"] = v
		}
		if v, ok := d.GetOk("stateful_inspection.tcp_end_timeout_r8020_gw_and_above"); ok {
			res["tcp-end-timeout-r8020-gw-and-above"] = v
		}
		if v, ok := d.GetOk("stateful_inspection.udp_virtual_session_timeout"); ok {
			res["udp-virtual-session-timeout"] = v
		}
		if v, ok := d.GetOk("stateful_inspection.icmp_virtual_session_timeout"); ok {
			res["icmp-virtual-session-timeout"] = v
		}
		if v, ok := d.GetOk("stateful_inspection.other_ip_protocols_virtual_session_timeout"); ok {
			res["other-ip-protocols-virtual-session-timeout"] = v
		}
		if v, ok := d.GetOk("stateful_inspection.sctp_start_timeout"); ok {
			res["sctp-start-timeout"] = v
		}
		if v, ok := d.GetOk("stateful_inspection.sctp_session_timeout"); ok {
			res["sctp-session-timeout"] = v
		}
		if v, ok := d.GetOk("stateful_inspection.sctp_end_timeout"); ok {
			res["sctp-end-timeout"] = v
		}
		if v, ok := d.GetOk("stateful_inspection.accept_stateful_udp_replies_for_unknown_services"); ok {
			res["accept-stateful-udp-replies-for-unknown-services"] = v
		}
		if v, ok := d.GetOk("stateful_inspection.accept_stateful_icmp_errors"); ok {
			res["accept-stateful-icmp-errors"] = v
		}
		if v, ok := d.GetOk("stateful_inspection.accept_stateful_icmp_replies"); ok {
			res["accept-stateful-icmp-replies"] = v
		}
		if v, ok := d.GetOk("stateful_inspection.accept_stateful_other_ip_protocols_replies_for_unknown_services"); ok {
			res["accept-stateful-other-ip-protocols-replies-for-unknown-services"] = v
		}
		if v, ok := d.GetOk("stateful_inspection.drop_out_of_state_tcp_packets"); ok {
			res["drop-out-of-state-tcp-packets"] = v
		}
		if v, ok := d.GetOk("stateful_inspection.log_on_drop_out_of_state_tcp_packets"); ok {
			res["log-on-drop-out-of-state-tcp-packets"] = v
		}
		if v, ok := d.GetOk("stateful_inspection.tcp_out_of_state_drop_exceptions"); ok {
			res["tcp-out-of-state-drop-exceptions"] = v
		}
		if v, ok := d.GetOk("stateful_inspection.drop_out_of_state_icmp_packets"); ok {
			res["drop-out-of-state-icmp-packets"] = v
		}
		if v, ok := d.GetOk("stateful_inspection.log_on_drop_out_of_state_icmp_packets"); ok {
			res["log-on-drop-out-of-state-icmp-packets"] = v
		}
		if v, ok := d.GetOk("stateful_inspection.drop_out_of_state_sctp_packets"); ok {
			res["drop-out-of-state-sctp-packets"] = v
		}
		if v, ok := d.GetOk("stateful_inspection.log_on_drop_out_of_state_sctp_packets"); ok {
			res["log-on-drop-out-of-state-sctp-packets"] = v
		}
		payload["stateful-inspection"] = res
	}

	if _, ok := d.GetOk("log_and_alert"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("log_and_alert.administrative_notifications"); ok {
			res["administrative-notifications"] = v.(string)
		}
		if v, ok := d.GetOk("log_and_alert.connection_matched_by_sam"); ok {
			res["connection-matched-by-sam"] = v.(string)
		}
		if v, ok := d.GetOk("log_and_alert.dynamic_object_resolution_failure"); ok {
			res["dynamic-object-resolution-failure"] = v.(string)
		}
		if v, ok := d.GetOk("log_and_alert.ip_options_drop"); ok {
			res["ip-options-drop"] = v.(string)
		}
		if v, ok := d.GetOk("log_and_alert.packet_is_incorrectly_tagged"); ok {
			res["packet-is-incorrectly-tagged"] = v.(string)
		}
		if v, ok := d.GetOk("log_and_alert.packet_tagging_brute_force_attack"); ok {
			res["packet-tagging-brute-force-attack"] = v.(string)
		}
		if v, ok := d.GetOk("log_and_alert.sla_violation"); ok {
			res["sla-violation"] = v.(string)
		}
		if v, ok := d.GetOk("log_and_alert.vpn_conf_and_key_exchange_errors"); ok {
			res["vpn-conf-and-key-exchange-errors"] = v.(string)
		}
		if v, ok := d.GetOk("log_and_alert.vpn_packet_handling_error"); ok {
			res["vpn-packet-handling-error"] = v.(string)
		}
		if v, ok := d.GetOk("log_and_alert.vpn_successful_key_exchange"); ok {
			res["vpn-successful-key-exchange"] = v.(string)
		}
		if v, ok := d.GetOk("log_and_alert.log_every_authenticated_http_connection"); ok {
			res["log-every-authenticated-http-connection"] = v
		}
		if v, ok := d.GetOk("log_and_alert.log_traffic"); ok {
			res["log-traffic"] = v.(string)
		}
		if v, ok := d.GetOk("log_and_alert.alerts"); ok {
			res["alerts"] = v
		}
		if v, ok := d.GetOk("log_and_alert.time_settings"); ok {
			res["time-settings"] = v
		}
		payload["log-and-alert"] = res
	}

	if _, ok := d.GetOk("data_access_control"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("data_access_control.auto_download_important_data"); ok {
			res["auto-download-important-data"] = v
		}
		if v, ok := d.GetOk("data_access_control.auto_download_sw_updates_and_new_features"); ok {
			res["auto-download-sw-updates-and-new-features"] = v
		}
		if v, ok := d.GetOk("data_access_control.send_anonymous_info"); ok {
			res["send-anonymous-info"] = v
		}
		if v, ok := d.GetOk("data_access_control.share_sensitive_info"); ok {
			res["share-sensitive-info"] = v
		}
		payload["data-access-control"] = res
	}

	if v, ok := d.GetOk("non_unique_ip_address_ranges"); ok {

		nonUniqueIpAddressRangesList := v.([]interface{})

		if len(nonUniqueIpAddressRangesList) > 0 {

			var nonUniqueIpAddressRangesPayload []map[string]interface{}

			for i := range nonUniqueIpAddressRangesList {

				Payload := make(map[string]interface{})

				if v, ok := d.GetOk("non_unique_ip_address_ranges." + strconv.Itoa(i) + ".address_type"); ok {
					Payload["address-type"] = v.(string)
				}
				if v, ok := d.GetOk("non_unique_ip_address_ranges." + strconv.Itoa(i) + ".first_ipv4_address"); ok {
					Payload["first-ipv4-address"] = v.(string)
				}
				if v, ok := d.GetOk("non_unique_ip_address_ranges." + strconv.Itoa(i) + ".first_ipv6_address"); ok {
					Payload["first-ipv6-address"] = v.(string)
				}
				if v, ok := d.GetOk("non_unique_ip_address_ranges." + strconv.Itoa(i) + ".last_ipv4_address"); ok {
					Payload["last-ipv4-address"] = v.(string)
				}
				if v, ok := d.GetOk("non_unique_ip_address_ranges." + strconv.Itoa(i) + ".last_ipv6_address"); ok {
					Payload["last-ipv6-address"] = v.(string)
				}
				nonUniqueIpAddressRangesPayload = append(nonUniqueIpAddressRangesPayload, Payload)
			}
			payload["nonUniqueIpAddressRanges"] = nonUniqueIpAddressRangesPayload
		}
	}

	if _, ok := d.GetOk("proxy"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("proxy.use_proxy_server"); ok {
			res["use-proxy-server"] = v
		}
		if v, ok := d.GetOk("proxy.proxy_address"); ok {
			res["proxy-address"] = v.(string)
		}
		if v, ok := d.GetOk("proxy.proxy_port"); ok {
			res["proxy-port"] = v
		}
		payload["proxy"] = res
	}

	if _, ok := d.GetOk("user_check"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("user_check.preferred_language"); ok {
			res["preferred-language"] = v.(string)
		}
		if v, ok := d.GetOk("user_check.send_emails_using_mail_server"); ok {
			res["send-emails-using-mail-server"] = v.(string)
		}
		payload["user-check"] = res
	}

	if _, ok := d.GetOk("hit_count"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("hit_count.enable_hit_count"); ok {
			res["enable-hit-count"] = v
		}
		if v, ok := d.GetOk("hit_count.keep_hit_count_data_up_to"); ok {
			res["keep-hit-count-data-up-to"] = v.(string)
		}
		payload["hit-count"] = res
	}

	if _, ok := d.GetOk("advanced_conf"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("advanced_conf.certs_and_pki"); ok {
			res["certs-and-pki"] = v
		}
		payload["advanced-conf"] = res
	}

	if v, ok := d.GetOkExists("allow_remote_registration_of_opsec_products"); ok {
		payload["allow-remote-registration-of-opsec-products"] = v.(bool)
	}

	if v, ok := d.GetOk("num_spoofing_errs_that_trigger_brute_force"); ok {
		payload["num-spoofing-errs-that-trigger-brute-force"] = v.(int)
	}

	if v, ok := d.GetOk("domains_to_process"); ok {
		payload["domains-to-process"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		payload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		payload["ignore-errors"] = v.(bool)
	}

	SetGlobalPropertiesRes, _ := client.ApiCall("set-global-properties", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if !SetGlobalPropertiesRes.Success {
		return fmt.Errorf(SetGlobalPropertiesRes.ErrorMsg)
	}

	d.SetId("set-global-properties-" + acctest.RandString(10))
	return nil
}

func deleteManagementSetGlobalProperties(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
