---
layout: "checkpoint"
page_title: "checkpoint_management_command_set_global_properties"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-command-set-global-properties"
description: |-
This resource allows you to execute Check Point Set Global Properties.
---

# Resource: checkpoint_management_command_set_global_properties

This resource allows you to execute Check Point Set Global Properties.

## Example Usage


```hcl
resource "checkpoint_management_command_set_global_properties" "example" {
  hit_count = {
    enable_hit_count = false
  }
  data_access_control = {
    auto_download_important_data = false
  }
}
```

## Argument Reference

The following arguments are supported:

* `firewall` - (Optional) Add implied rules to or remove them from the Firewall Rule Base. Determine the position of the implied rules in the Rule Base, and whether or not to log them.firewall blocks are documented below.
* `nat` - (Optional) Configure settings that apply to all NAT connections.nat blocks are documented below.
* `authentication` - (Optional) Define Authentication properties that are common to all users and to the various ways that the Check Point Security Gateway asks for passwords (User, Client and Session Authentication).authentication blocks are documented below.
* `vpn` - (Optional) Configure settings relevant to VPN.vpn blocks are documented below.
* `remote_access` - (Optional) Configure Remote Access properties.remote_access blocks are documented below.
* `user_directory` - (Optional) User can enable LDAP User Directory as well as specify global parameters for LDAP. If LDAP User Directory is enabled, this means that users are managed on an external LDAP server and not on the internal Check Point Security Gateway users databases.user_directory blocks are documented below.
* `qos` - (Optional) Define the general parameters of Quality of Service (QoS) and apply them to QoS rules.qos blocks are documented below.
* `carrier_security` - (Optional) Specify system-wide properties. Select GTP intra tunnel inspection options, including anti-spoofing; tracking and logging options, and integrity tests.carrier_security blocks are documented below.
* `user_accounts` - (Optional) Set the expiration for a user account and configure "about to expire" warnings.user_accounts blocks are documented below.
* `user_authority` - (Optional) Decide whether to display and access the WebAccess rule base. This policy defines which users (that is, which Windows Domains) have access to the internal sites of the organization.user_authority blocks are documented below.
* `connect_control` - (Optional) Configure settings that relate to ConnectControl server load balancing.connect_control blocks are documented below.
* `stateful_inspection` - (Optional) Adjust Stateful Inspection parameters.stateful_inspection blocks are documented below.
* `log_and_alert` - (Optional) Define system-wide logging and alerting parameters.log_and_alert blocks are documented below.
* `data_access_control` - (Optional) Configure automatic downloads from Check Point and anonymously share product data. Options selected here apply to all Security Gateways, Clusters and VSX devices managed by this management server.data_access_control blocks are documented below.
* `non_unique_ip_address_ranges` - (Optional) Specify Non Unique IP Address Ranges.non_unique_ip_address_ranges blocks are documented below.
* `proxy` - (Optional) Select whether a proxy server is used when servers, gateways, or clients need to access the internet for certain Check Point features and set the default proxy server that will be used.proxy blocks are documented below.
* `user_check` - (Optional) Set a language for the UserCheck message if the language setting in the user's browser cannot be determined.user_check blocks are documented below.
* `hit_count` - (Optional) Enable the Hit Count feature that tracks the number of connections that each rule matches.hit_count blocks are documented below.
* `advanced_conf` - (Optional) Configure advanced global attributes. It's highly recommended to consult with Check Point's Technical Support before modifying these values.advanced_conf blocks are documented below.
* `allow_remote_registration_of_opsec_products` - (Optional) After installing an OPSEC application, the remote administration (RA) utility enables an OPSEC product to finish registering itself without having to access the SmartConsole. If set to true, any host including the application host can run the utility. Otherwise,  the RA utility can only be run from the Security Management host. 
* `num_spoofing_errs_that_trigger_brute_force` - (Optional) Indicates how many incorrectly signed packets will be tolerated before assuming that there is an attack on the packet tagging and revoking the client's key. 
* `domains_to_process` - (Optional) Indicates which domains to process the commands on. It cannot be used with the details-level full, must be run from the System Domain only and with ignore-warnings true. Valid values are: CURRENT_DOMAIN, ALL_DOMAINS_ON_THIS_SERVER.domains_to_process blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


`firewall` supports the following:

* `accept_control_connections` - (Optional) Used for:<br>&nbsp;&nbsp;&nbsp;&nbsp; <ul><li> Installing the security policy from the Security Management server to the gateways.</li><br>&nbsp;&nbsp;&nbsp;&nbsp; <li> Sending logs from the gateways to the Security Management server.</li><br>&nbsp;&nbsp;&nbsp;&nbsp; <li> Communication between SmartConsole clients and the Security Management Server</li><br>&nbsp;&nbsp;&nbsp;&nbsp; <li> Communication between Firewall daemons on different machines (Security Management Server, Security Gateway).</li><br>&nbsp;&nbsp;&nbsp;&nbsp; <li> Connecting to OPSEC applications such as RADIUS and TACACS authentication servers.</li></ul>If you disable Accept Control Connections and you want Check Point components to communicate with each other and with OPSEC components, you must explicitly allow these connections in the Rule Base. 
* `accept_ips1_management_connections` - (Optional) Accepts IPS-1 connections.<br>Available only if accept-control-connections is true. 
* `accept_remote_access_control_connections` - (Optional) Accepts Remote Access connections.<br>Available only if accept-control-connections is true. 
* `accept_smart_update_connections` - (Optional) Accepts SmartUpdate connections. 
* `accept_outgoing_packets_originating_from_gw` - (Optional) Accepts all packets from connections that originate at the Check Point Security Gateway. 
* `accept_outgoing_packets_originating_from_gw_position` - (Optional) The position of the implied rules in the Rule Base.<br>Available only if accept-outgoing-packets-originating-from-gw is false. 
* `accept_outgoing_packets_originating_from_connectra_gw` - (Optional) Accepts outgoing packets originating from Connectra gateway.<br>Available only if accept-outgoing-packets-originating-from-gw is false. 
* `accept_outgoing_packets_to_cp_online_services` - (Optional) Allow Security Gateways to access Check Point online services. Supported for R80.10 Gateway and higher.<br>Available only if accept-outgoing-packets-originating-from-gw is false. 
* `accept_outgoing_packets_to_cp_online_services_position` - (Optional) The position of the implied rules in the Rule Base.<br>Available only if accept-outgoing-packets-to-cp-online-services is true. 
* `accept_domain_name_over_tcp` - (Optional) Accepts Domain Name (DNS) queries and replies over TCP, to allow downloading of the domain name-resolving tables used for zone transfers between servers. For clients, DNS over TCP is only used if the tables to be transferred are very large. 
* `accept_domain_name_over_tcp_position` - (Optional) The position of the implied rules in the Rule Base.<br>Available only if accept-domain-name-over-tcp is true. 
* `accept_domain_name_over_udp` - (Optional) Accepts Domain Name (DNS) queries and replies over UDP. 
* `accept_domain_name_over_udp_position` - (Optional) The position of the implied rules in the Rule Base.<br>Available only if accept-domain-name-over-udp is true. 
* `accept_dynamic_addr_modules_outgoing_internet_connections` - (Optional) Accept Dynamic Address modules' outgoing internet connections.Accepts DHCP traffic for DAIP (Dynamically Assigned IP Address) gateways. In Small Office Appliance gateways, this rule allows outgoing DHCP, PPP, PPTP and L2TP Internet connections (regardless of whether it is or is not a DAIP gateway). 
* `accept_icmp_requests` - (Optional) Accepts Internet Control Message Protocol messages. 
* `accept_icmp_requests_position` - (Optional) The position of the implied rules in the Rule Base.<br>Available only if accept-icmp-requests is true. 
* `accept_identity_awareness_control_connections` - (Optional) Accepts traffic between Security Gateways in distributed environment configurations of Identity Awareness. 
* `accept_identity_awareness_control_connections_position` - (Optional) The position of the implied rules in the Rule Base.<br>Available only if accept-identity-awareness-control-connections is true. 
* `accept_incoming_traffic_to_dhcp_and_dns_services_of_gws` - (Optional) Allows the Small Office Appliance gateway to provide DHCP relay, DHCP server and DNS proxy services regardless of the rule base. 
* `accept_rip` - (Optional) Accepts Routing Information Protocol (RIP), using UDP on port 520. 
* `accept_rip_position` - (Optional) The position of the implied rules in the Rule Base.<br>Available only if accept-rip is true. 
* `accept_vrrp_packets_originating_from_cluster_members` - (Optional) Selecting this option creates an implied rule in the security policy Rule Base that accepts VRRP inbound and outbound traffic to and from the members of the cluster. 
* `accept_web_and_ssh_connections_for_gw_administration` - (Optional) Accepts Web and SSH connections for Small Office Appliance gateways. 
* `log_implied_rules` - (Optional) Produces log records for communications that match the implied rules that are generated in the Rule Base from the properties defined in this window. 
* `security_server` - (Optional) Control the welcome messages that users will see when logging in to servers behind Check Point Security Gateways.security_server blocks are documented below.


`nat` supports the following:

* `allow_bi_directional_nat` - (Optional) Applies to automatic NAT rules in the NAT Rule Base, and allows two automatic NAT rules to match a connection. Without Bidirectional NAT, only one automatic NAT rule can match a connection. 
* `auto_arp_conf` - (Optional) Ensures that ARP requests for a translated (NATed) machine, network or address range are answered by the Check Point Security Gateway. 
* `merge_manual_proxy_arp_conf` - (Optional) Merges the automatic and manual ARP configurations. Manual proxy ARP configuration is required for manual Static NAT rules.<br>Available only if auto-arp-conf is true. 
* `auto_translate_dest_on_client_side` - (Optional) Applies to packets originating at the client, with the server as its destination. Static NAT for the server is performed on the client side. 
* `manually_translate_dest_on_client_side` - (Optional) Applies to packets originating at the client, with the server as its destination. Static NAT for the server is performed on the client side. 
* `enable_ip_pool_nat` - (Optional) Applies to packets originating at the client, with the server as its destination. Static NAT for the server is performed on the client side. 
* `addr_alloc_and_release_track` - (Optional) Specifies whether to log each allocation and release of an IP address from the IP Pool.<br>Available only if enable-ip-pool-nat is true. 
* `addr_exhaustion_track` - (Optional) Specifies the action to take if the IP Pool is exhausted.<br>Available only if enable-ip-pool-nat is true. 


`authentication` supports the following:

* `auth_internal_users_with_specific_suffix` - (Optional) Enforce suffix for internal users authentication. 
* `allowed_suffix_for_internal_users` - (Optional) Suffix for internal users authentication. 
* `max_days_before_expiration_of_non_pulled_user_certificates` - (Optional) Users certificates which were initiated but not pulled will expire after the specified number of days. Any value from 1 to 60 days can be entered in this field. 
* `max_client_auth_attempts_before_connection_termination` - (Optional) Allowed Number of Failed Client Authentication Attempts Before Session Termination. Any value from 1 to 800 attempts can be entered in this field. 
* `max_rlogin_attempts_before_connection_termination` - (Optional) Allowed Number of Failed rlogin Attempts Before Session Termination. Any value from 1 to 800 attempts can be entered in this field. 
* `max_session_auth_attempts_before_connection_termination` - (Optional) Allowed Number of Failed Session Authentication Attempts Before Session Termination. Any value from 1 to 800 attempts can be entered in this field. 
* `max_telnet_attempts_before_connection_termination` - (Optional) Allowed Number of Failed telnet Attempts Before Session Termination. Any value from 1 to 800 attempts can be entered in this field. 
* `enable_delayed_auth` - (Optional) all authentications other than certificate-based authentications will be delayed by the specified time. Applying this delay will stall brute force authentication attacks. The delay is applied for both failed and successful authentication attempts. 
* `delay_each_auth_attempt_by` - (Optional) Delay each authentication attempt by the specified number of milliseconds. Any value from 1 to 25000 can be entered in this field. 


`vpn` supports the following:

* `vpn_conf_method` - (Optional) Decide on Simplified or Traditional mode for all new security policies or decide which mode to use on a policy by policy basis. 
* `domain_name_for_dns_resolving` - (Optional) Enter the domain name that will be used for gateways DNS lookup. The DNS host name that is used is "gateway_name.domain_name". 
* `enable_backup_gw` - (Optional) Enable Backup Gateway. 
* `enable_decrypt_on_accept_for_gw_to_gw_traffic` - (Optional) Enable decrypt on accept for gateway to gateway traffic. This is only relevant for policies in traditional mode. In Traditional Mode, the 'Accept' action determines that a connection is allowed, while the 'Encrypt' action determines that a connection is allowed and encrypted. Select whether VPN accepts an encrypted packet that matches a rule with an 'Accept' action or drops it. 
* `enable_load_distribution_for_mep_conf` - (Optional) Enable load distribution for Multiple Entry Points configurations (Site To Site connections). The VPN Multiple Entry Point (MEP) feature supplies high availability and load distribution for Check Point Security Gateways. MEP works in four modes:<br>&nbsp;&nbsp;&nbsp;&nbsp; <ul><li> First to Respond, in which the first gateway to reply to the peer gateway is chosen. An organization would choose this option if, for example, the organization has two gateways in a MEPed configuration - one in London, the other in New York. It makes sense for Check Point Security Gateway peers located in England to try the London gateway first and the NY gateway second. Being geographically closer to Check Point Security Gateway peers in England, the London gateway will be the first to respond, and becomes the entry point to the internal network.</li><br>&nbsp;&nbsp;&nbsp;&nbsp; <li> VPN Domain, is when the destination IP belongs to a particular VPN domain, the gateway of that domain becomes the chosen entry point. This gateway becomes the primary gateway while other gateways in the MEP configuration become its backup gateways.</li><br>&nbsp;&nbsp;&nbsp;&nbsp; <li> Random Selection, in which the remote Check Point Security Gateway peer randomly selects a gateway with which to open a VPN connection. For each IP source/destination address pair, a new gateway is randomly selected. An organization might have a number of machines with equal performance abilities. In this case, it makes sense to enable load distribution. The machines are used in a random and equal way.</li><br>&nbsp;&nbsp;&nbsp;&nbsp; <li> Manually set priority list, gateway priorities can be set manually for the entire community or for individual satellite gateways.</li></ul>. 
* `enable_vpn_directional_match_in_vpn_column` - (Optional) Enable VPN Directional Match in VPN Column.<br>Note: VPN Directional Match is supported only on Gaia, SecurePlatform, Linux and IPSO. 
* `grace_period_after_the_crl_is_not_valid` - (Optional) When establishing VPN tunnels, the peer presents its certificate for authentication. The clock on the gateway machine must be synchronized with the clock on the Certificate Authority machine. Otherwise, the Certificate Revocation List (CRL) used for validating the peer's certificate may be considered invalid and thus the authentication fails. To resolve the issue of differing clock times, a Grace Period permits a wider window for CRL validity. 
* `grace_period_before_the_crl_is_valid` - (Optional) When establishing VPN tunnels, the peer presents its certificate for authentication. The clock on the gateway machine must be synchronized with the clock on the Certificate Authority machine. Otherwise, the Certificate Revocation List (CRL) used for validating the peer's certificate may be considered invalid and thus the authentication fails. To resolve the issue of differing clock times, a Grace Period permits a wider window for CRL validity. 
* `grace_period_extension_for_secure_remote_secure_client` - (Optional) When dealing with remote clients the Grace Period needs to be extended. The remote client sometimes relies on the peer gateway to supply the CRL. If the client's clock is not synchronized with the gateway's clock, a CRL that is considered valid by the gateway may be considered invalid by the client. 
* `support_ike_dos_protection_from_identified_src` - (Optional) When the number of IKE negotiations handled simultaneously exceeds a threshold above VPN's capacity, a gateway concludes that it is either under a high load or experiencing a Denial of Service attack. VPN can filter out peers that are the probable source of the potential Denial of Service attack. There are two kinds of protection:<br>&nbsp;&nbsp;&nbsp;&nbsp; <ul><li> Stateless - the peer has to respond to an IKE notification in a way that proves the peer's IP address is not spoofed. If the peer cannot prove this, VPN does not allocate resources for the IKE negotiation</li><br>&nbsp;&nbsp;&nbsp;&nbsp; <li> Puzzles - this is the same as Stateless, but in addition, the peer has to solve a mathematical puzzle. Solving this puzzle consumes peer CPU resources in a way that makes it difficult to initiate multiple IKE negotiations simultaneously.</li></ul>Puzzles is more secure then Stateless, but affects performance.<br>Since these kinds of attacks involve a new proprietary addition to the IKE protocol, enabling these protection mechanisms may cause difficulties with non Check Point VPN products or older versions of VPN. 
* `support_ike_dos_protection_from_unidentified_src` - (Optional) When the number of IKE negotiations handled simultaneously exceeds a threshold above VPN's capacity, a gateway concludes that it is either under a high load or experiencing a Denial of Service attack. VPN can filter out peers that are the probable source of the potential Denial of Service attack. There are two kinds of protection:<br>&nbsp;&nbsp;&nbsp;&nbsp; <ul><li> Stateless - the peer has to respond to an IKE notification in a way that proves the peer's IP address is not spoofed. If the peer cannot prove this, VPN does not allocate resources for the IKE negotiation</li><br>&nbsp;&nbsp;&nbsp;&nbsp; <li> Puzzles - this is the same as Stateless, but in addition, the peer has to solve a mathematical puzzle. Solving this puzzle consumes peer CPU resources in a way that makes it difficult to initiate multiple IKE negotiations simultaneously.</li></ul>Puzzles is more secure then Stateless, but affects performance.<br>Since these kinds of attacks involve a new proprietary addition to the IKE protocol, enabling these protection mechanisms may cause difficulties with non Check Point VPN products or older versions of VPN. 


`remote_access` supports the following:

* `enable_back_connections` - (Optional) Usually communication with remote clients must be initialized by the clients. However, once a client has opened a connection, the hosts behind VPN can open a return or back connection to the client. For a back connection, the client's details must be maintained on all the devices between the client and the gateway, and on the gateway itself. Determine whether the back connection is enabled. 
* `keep_alive_packet_to_gw_interval` - (Optional) Usually communication with remote clients must be initialized by the clients. However, once a client has opened a connection, the hosts behind VPN can open a return or back connection to the client. For a back connection, the client's details must be maintained on all the devices between the client and the gateway, and on the gateway itself. Determine frequency (in seconds) of the Keep Alive packets sent by the client in order to maintain the connection with the gateway.<br>Available only if enable-back-connections is true. 
* `encrypt_dns_traffic` - (Optional) You can decide whether DNS queries sent by the remote client to a DNS server located on the corporate LAN are passed through the VPN tunnel or not. Disable this option if the client has to make DNS queries to the DNS server on the corporate LAN while connecting to the organization but without using the SecuRemote client. 
* `simultaneous_login_mode` - (Optional) Select the simultaneous login mode. 
* `vpn_authentication_and_encryption` - (Optional) configure supported Encryption and Authentication methods for Remote Access clients.vpn_authentication_and_encryption blocks are documented below.
* `vpn_advanced` - (Optional) Configure encryption methods and interface resolution for remote access clients.vpn_advanced blocks are documented below.
* `scv` - (Optional) Define properties of the Secure Configuration Verification process.scv blocks are documented below.
* `ssl_network_extender` - (Optional) Define properties for SSL Network Extender users.ssl_network_extender blocks are documented below.
* `secure_client_mobile` - (Optional) Define properties for SecureClient Mobile.secure_client_mobile blocks are documented below.
* `endpoint_connect` - (Optional) Configure global settings for Endpoint Connect. These settings apply to all gateways.endpoint_connect blocks are documented below.
* `hot_spot_and_hotel_registration` - (Optional) Configure the settings for Wireless Hot Spot and Hotel Internet access registration.hot_spot_and_hotel_registration blocks are documented below.


`user_directory` supports the following:

* `enable_password_change_when_user_active_directory_expires` - (Optional) For organizations using MS Active Directory, this setting enables users whose passwords have expired to automatically create new passwords. 
* `cache_size` - (Optional) The maximum number of cached users allowed. The cache is FIFO (first-in, first-out). When a new user is added to a full cache, the first user is deleted to make room for the new user. The Check Point Security Gateway does not query the LDAP server for users already in the cache, unless the cache has timed out. 
* `enable_password_expiration_configuration` - (Optional) Enable configuring of the number of days during which the password is valid.<br>If enable-password-change-when-user-active-directory-expires is true, the password expiration time is determined by the Active Directory. In this case it is recommended not to set this to true. 
* `password_expires_after` - (Optional) Specifies the number of days during which the password is valid. Users are authenticated using a special LDAP password. Should this password expire, a new password must be defined.<br>Available only if enable-password-expiration-configuration is true. 
* `timeout_on_cached_users` - (Optional) The period of time in which a cached user is timed out and will need to be fetched again from the LDAP server. 
* `display_user_dn_at_login` - (Optional) Decide whether or not you would like to display the user's DN when logging in. If you choose to display the user DN, you can select whether to display it, when the user is prompted for the password at login, or on the request of the authentication scheme. This property is a useful diagnostic tool when there is more than one user with the same name in an Account Unit. In this case, the first one is chosen and the others are ignored. 
* `enforce_rules_for_user_mgmt_admins` - (Optional) Enforces password strength rules on LDAP users when you create or modify a Check Point Password. 
* `min_password_length` - (Optional) Specifies the minimum length (in characters) of the password. 
* `password_must_include_a_digit` - (Optional) Password must include a digit. 
* `password_must_include_a_symbol` - (Optional) Password must include a symbol. 
* `password_must_include_lowercase_char` - (Optional) Password must include a lowercase character. 
* `password_must_include_uppercase_char` - (Optional) Password must include an uppercase character. 


`qos` supports the following:

* `default_weight_of_rule` - (Optional) Define a Weight at which bandwidth will be guaranteed. Set a default weight for a rule.<br>Note: Value will be applied to new rules only. 
* `max_weight_of_rule` - (Optional) Define a Weight at which bandwidth will be guaranteed. Set a maximum weight for a rule. 
* `unit_of_measure` - (Optional) Define the Rate at which packets are transmitted, for which bandwidth will be guaranteed. Set a Unit of measure. 
* `authenticated_ip_expiration` - (Optional) Define the Authentication time-out for QoS. This timeout is set in minutes. In an Authenticated IP all connections which are open in a specified time limit will be guaranteed bandwidth, but will not be guaranteed bandwidth after the time limit. 
* `non_authenticated_ip_expiration` - (Optional) Define the Authentication time-out for QoS. This timeout is set in minutes. 
* `unanswered_queried_ip_expiration` - (Optional) Define the Authentication time-out for QoS. This timeout is set in minutes. 


`carrier_security` supports the following:

* `block_gtp_in_gtp` - (Optional) Prevents GTP packets from being encapsulated inside GTP tunnels. When this option is checked, such packets are dropped and logged. 
* `enforce_gtp_anti_spoofing` - (Optional) verifies that G-PDUs are using the end user IP address that has been agreed upon in the PDP context activation process. When this option is checked, packets that do not use this IP address are dropped and logged. 
* `produce_extended_logs_on_unmatched_pdus` - (Optional) logs GTP packets not matched by previous rules with Carrier Security's extended GTP-related log fields. These logs are brown and their Action attribute is empty. The default setting is checked. 
* `produce_extended_logs_on_unmatched_pdus_position` - (Optional) Choose to place this implicit rule Before Last or as the Last rule.<br>Available only if produce-extended-logs-on-unmatched-pdus is true. 
* `protocol_violation_track_option` - (Optional) Set the appropriate track or alert option to be used when a protocol violation (malformed packet) is detected. 
* `enable_g_pdu_seq_number_check_with_max_deviation` - (Optional) If set to false, sequence checking is not enforced and all out-of-sequence G-PDUs will be accepted.<br>To enhance performance, disable this extended integrity test. 
* `g_pdu_seq_number_check_max_deviation` - (Optional) specifies that a G-PDU is accepted only if the difference between its sequence number and the expected sequence number is less than or equal to the allowed deviation.<br>Available only ifenable-g-pdu-seq-number-check-with-max-deviation is true. 
* `verify_flow_labels` - (Optional) See that each packet's flow label matches the flow labels defined by GTP signaling. This option is relevant for GTP version 0 only.<br>To enhance performance, disable this extended integrity test. 
* `allow_ggsn_replies_from_multiple_interfaces` - (Optional) Allows GTP signaling replies from an IP address different from the IP address to which the requests are sent (Relevant only for gateways below R80). 
* `enable_reverse_connections` - (Optional) Allows Carrier Security gateways to accept PDUs sent from the GGSN to the SGSN, on a previously established PDP context, even if these PDUs are sent over ports that do not match the ports of the established PDP context. 
* `gtp_signaling_rate_limit_sampling_interval` - (Optional) Works in correlation with the property Enforce GTP Signal packet rate limit found in the Carrier Security window of the GSN network object. For example, with the rate limit sampling interval default of 1 second, and the network object enforced a GTP signal packet rate limit of the default 2048 PDU per second, sampling will occur one time per second, or 2048 signaling PDUs between two consecutive samplings. 
* `one_gtp_echo_on_each_path_frequency` - (Optional) sets the number of GTP Echo exchanges per path allowed per configured time period. Echo requests exceeding this rate are dropped and logged. Setting the value to 0 disables the feature and allows an unlimited number of echo requests per path at any interval. 
* `aggressive_aging` - (Optional) If true, enables configuring aggressive aging thresholds and time out value. 
* `aggressive_timeout` - (Optional) Aggressive timeout. Available only if aggressive-aging is true. 
* `memory_activation_threshold` - (Optional) Memory activation threshold. Available only if aggressive-aging is true. 
* `memory_deactivation_threshold` - (Optional) Memory deactivation threshold. Available only if aggressive-aging is true. 
* `tunnel_activation_threshold` - (Optional) Tunnel activation threshold. Available only if aggressive-aging is true. 
* `tunnel_deactivation_threshold` - (Optional) Tunnel deactivation threshold. Available only if aggressive-aging is true. 


`user_accounts` supports the following:

* `expiration_date_method` - (Optional) Select an Expiration Date Method.<br>Expire at - Account expires on the date that you select.<br>Expire after - Account expires after the number of days that you select. 
* `expiration_date` - (Optional) Specify an Expiration Date in the following format: YYYY-MM-DD.<br>Available only if expiration-date-method is set to "expire at". 
* `days_until_expiration` - (Optional) Account expires after the number of days that you select.<br>Available only if expiration-date-method is set to "expire after". 
* `show_accounts_expiration_indication_days_in_advance` - (Optional) Activates the Expired Accounts link, to open the Expired Accounts window. 


`user_authority` supports the following:

* `display_web_access_view` - (Optional) Specify whether or not to display the WebAccess rule base. This rule base is used for UserAuthority. 
* `windows_domains_to_trust` - (Optional) When matching Firewall usernames to Windows Domains usernames for Single Sign on, selectwhether to trust all or specify which Windows Domain should be trusted.<br>ALL - Enables you to allow all Windows domains to access the internal sites of the organization.<br>SELECTIVELY - Enables you to specify which Windows domains will have access to the internal sites of the organization. 
* `trust_only_following_windows_domains` - (Optional) Specify which Windows domains will have access to the internal sites of the organization.<br>Available only if windows-domains-to-trust is set to SELECTIVELY.trust_only_following_windows_domains blocks are documented below.


`connect_control` supports the following:

* `load_agents_port` - (Optional) Sets the port number on which load measuring agents communicate with ConnectControl. 
* `load_measurement_interval` - (Optional) sets how often (in seconds) the load measuring agents report their load status to ConnectControl. 
* `persistence_server_timeout` - (Optional) Sets the amount of time (in seconds) that a client, once directed to a particular server, will continue to be directed to that same server. 
* `server_availability_check_interval` - (Optional) Sets how often (in seconds) ConnectControl checks to make sure the load balanced servers are running and responding to service requests. 
* `server_check_retries` - (Optional) Sets how many times ConnectControl attempts to contact a server before ceasing to direct traffic to it. 


`stateful_inspection` supports the following:

* `tcp_start_timeout` - (Optional) A TCP connection will be timed out if the interval between the arrival of the first packet and establishment of the connection (TCP three-way handshake) exceeds TCP start timeout seconds. 
* `tcp_session_timeout` - (Optional) The length of time (in seconds) an idle connection will remain in the Security Gateway connections table. 
* `tcp_end_timeout` - (Optional) A TCP connection will only terminate TCP end timeout seconds after two FIN packets (one in each direction: client-to-server, and server-to-client) or an RST packet. When a TCP connection ends (FIN packets sent or connection reset) the Check Point Security Gateway will keep the connection in the connections table for another TCP end timeout seconds, to allow for stray ACKs of the connection that arrive late. 
* `tcp_end_timeout_r8020_gw_and_above` - (Optional) A TCP connection will only terminate TCP end timeout seconds after two FIN packets (one in each direction: client-to-server, and server-to-client) or an RST packet. When a TCP connection ends (FIN packets sent or connection reset) the Check Point Security Gateway will keep the connection in the connections table for another TCP end timeout seconds, to allow for stray ACKs of the connection that arrive late. 
* `udp_virtual_session_timeout` - (Optional) Specifies the amount of time (in seconds) a UDP reply channel may remain open without any packets being returned. 
* `icmp_virtual_session_timeout` - (Optional) An ICMP virtual session will be considered to have timed out after this time period (in seconds). 
* `other_ip_protocols_virtual_session_timeout` - (Optional) A virtual session of services which are not explicitly configured here will be considered to have timed out after this time period (in seconds). 
* `sctp_start_timeout` - (Optional) SCTP connections will be timed out if the interval between the arrival of the first packet and establishment of the connection exceeds this value (in seconds). 
* `sctp_session_timeout` - (Optional) Time (in seconds) an idle connection will remain in the Security Gateway connections table. 
* `sctp_end_timeout` - (Optional) SCTP connections end after this number of seconds, after the connection ends or is reset, to allow for stray ACKs of the connection that arrive late. 
* `accept_stateful_udp_replies_for_unknown_services` - (Optional) Specifies if UDP replies are to be accepted for unknown services. 
* `accept_stateful_icmp_errors` - (Optional) Accept ICMP error packets which refer to another non-ICMP connection (for example, to an ongoing TCP or UDP connection) that was accepted by the Rule Base. 
* `accept_stateful_icmp_replies` - (Optional) Accept ICMP reply packets for ICMP requests that were accepted by the Rule Base. 
* `accept_stateful_other_ip_protocols_replies_for_unknown_services` - (Optional) Accept reply packets for other undefined services (that is, services which are not one of the following: TCP, UDP, ICMP). 
* `drop_out_of_state_tcp_packets` - (Optional) Drop TCP packets which are not consistent with the current state of the connection. 
* `log_on_drop_out_of_state_tcp_packets` - (Optional) Generates a log entry when these out of state TCP packets are dropped.<br>Available only if drop-out-of-state-tcp-packets is true. 
* `tcp_out_of_state_drop_exceptions` - (Optional) Name or uid of the gateways and clusters for which Out of State packets are allowed.tcp_out_of_state_drop_exceptions blocks are documented below.
* `drop_out_of_state_icmp_packets` - (Optional) Drop ICMP packets which are not consistent with the current state of the connection. 
* `log_on_drop_out_of_state_icmp_packets` - (Optional) Generates a log entry when these out of state ICMP packets are dropped.<br>Available only if drop-out-of-state-icmp-packets is true. 
* `drop_out_of_state_sctp_packets` - (Optional) Drop SCTP packets which are not consistent with the current state of the connection. 
* `log_on_drop_out_of_state_sctp_packets` - (Optional) Generates a log entry when these out of state SCTP packets are dropped.<br>Available only if drop-out-of-state-sctp-packets is true. 


`log_and_alert` supports the following:

* `administrative_notifications` - (Optional) Administrative notifications specifies the action to be taken when an administrative event (for example, when a certificate is about to expire) occurs. 
* `connection_matched_by_sam` - (Optional) Connection matched by SAM specifies the action to be taken when a connection is blocked by SAM (Suspicious Activities Monitoring). 
* `dynamic_object_resolution_failure` - (Optional) Dynamic object resolution failure specifies the action to be taken when a dynamic object cannot be resolved. 
* `ip_options_drop` - (Optional) IP Options drop specifies the action to take when a packet with IP Options is encountered. The Check Point Security Gateway always drops these packets, but you can log them or issue an alert. 
* `packet_is_incorrectly_tagged` - (Optional) Packet is incorrectly tagged. 
* `packet_tagging_brute_force_attack` - (Optional) Packet tagging brute force attack. 
* `sla_violation` - (Optional) SLA violation specifies the action to be taken when an SLA violation occurs, as defined in the Virtual Links window. 
* `vpn_conf_and_key_exchange_errors` - (Optional) VPN configuration & key exchange errors specifies the action to be taken when logging configuration or key exchange errors occur, for example, when attempting to establish encrypted communication with a network object inside the same encryption domain. 
* `vpn_packet_handling_error` - (Optional) VPN packet handling errors specifies the action to be taken when encryption or decryption errors occurs. A log entry contains the action performed (Drop or Reject) and a short description of the error cause, for example, scheme or method mismatch. 
* `vpn_successful_key_exchange` - (Optional) VPN successful key exchange specifies the action to be taken when VPN keys are successfully exchanged. 
* `log_every_authenticated_http_connection` - (Optional) Log every authenticated HTTP connection specifies that a log entry should be generated for every authenticated HTTP connection. 
* `log_traffic` - (Optional) Log Traffic specifies whether or not to log traffic. 
* `alerts` - (Optional) Define the behavior of alert logs and the type of alert used for System Alert logs.alerts blocks are documented below.
* `time_settings` - (Optional) Configure the time settings associated with system-wide logging and alerting parameters.time_settings blocks are documented below.


`data_access_control` supports the following:

* `auto_download_important_data` - (Optional) Automatically download and install Software Blade Contracts, security updates and other important data (highly recommended). 
* `auto_download_sw_updates_and_new_features` - (Optional) Automatically download software updates and new features (highly recommended).<br>Available only if auto-download-important-data is set to true. 
* `send_anonymous_info` - (Optional) Help Check Point improve the product by sending anonymous information. 
* `share_sensitive_info` - (Optional) Approve sharing core dump files and other relevant crash data which might contain personal information. All shared data will be processed in accordance with Check Point's Privacy Policy.<br>Available only if send-anonymous-info is set to true. 


`non_unique_ip_address_ranges` supports the following:

* `address_type` - (Optional) The type of the IP Address. 
* `first_ipv4_address` - (Optional) The first IPV4 Address in the range. 
* `first_ipv6_address` - (Optional) The first IPV6 Address in the range. 
* `last_ipv4_address` - (Optional) The last IPV4 Address in the range. 
* `last_ipv6_address` - (Optional) The last IPV6 Address in the range. 


`proxy` supports the following:

* `use_proxy_server` - (Optional) If set to true, a proxy server is used when features need to access the internet. 
* `proxy_address` - (Optional) Specify the URL or IP address of the proxy server.<br>Available only if use-proxy-server is set to true. 
* `proxy_port` - (Optional) Specify the Port on which the server will be accessed.<br>Available only if use-proxy-server is set to true. 


`user_check` supports the following:

* `preferred_language` - (Optional) The preferred language for new UserCheck message. 
* `send_emails_using_mail_server` - (Optional) Name or UID of mail server to send emails to. 


`hit_count` supports the following:

* `enable_hit_count` - (Optional) Select to enable or clear to disable all Security Gateways to monitor the number of connections each rule matches. 
* `keep_hit_count_data_up_to` - (Optional) Select one of the time range options. Data is kept in the Security Management Server database for this period and is shown in the Hits column. 


`advanced_conf` supports the following:

* `certs_and_pki` - (Optional) Configure Certificates and PKI properties.certs_and_pki blocks are documented below.


`security_server` supports the following:

* `client_auth_welcome_file` - (Optional) Client authentication welcome file is the name of a file whose contents are to be displayed when a user begins a Client Authenticated session (optional) using the Manual Sign On Method. Client Authenticated Sessions initiated by Manual Sign On are not mediated by a security server. 
* `ftp_welcome_msg_file` - (Optional) FTP welcome message file is the name of a file whose contents are to be displayed when a user begins an Authenticated FTP session. 
* `rlogin_welcome_msg_file` - (Optional) Rlogin welcome message file is the name of a file whose contents are to be displayed when a user begins an Authenticated RLOGIN session. 
* `telnet_welcome_msg_file` - (Optional) Telnet welcome message file is the name of a file whose contents are to be displayed when a user begins an Authenticated Telnet session. 
* `mdq_welcome_msg` - (Optional) MDQ Welcome Message is the message to be displayed when a user begins an MDQ session. The MDQ Welcome Message should contain characters according to RFC 1035 and it must follow the ARPANET host name rules:<br>   - This message must begin with a number or letter. After the first letter or number character the remaining characters can be a letter, number, space, tab or hyphen.<br>   - This message must not end with a space or a tab and is limited to 63 characters. 
* `smtp_welcome_msg` - (Optional) SMTP Welcome Message is the message to be displayed when a user begins an SMTP session. 
* `http_next_proxy_host` - (Optional) HTTP next proxy host is the host name of the HTTP proxy behind the Check Point Security Gateway HTTP security server (if there is one). Changing the HTTP Next Proxy fields takes effect after the Security Gateway database is downloaded to the authenticating gateway, or after the security policy is re-installed. <br>These settings apply only to firewalled gateways prior to NG. For later versions, these settings should be defined in the Node Properties window. 
* `http_next_proxy_port` - (Optional) HTTP next proxy port is the port of the HTTP proxy behind the Check Point Security Gateway HTTP security server (if there is one). Changing the HTTP Next Proxy fields takes effect after the Security Gateway database is downloaded to the authenticating gateway, or after the security policy is re-installed. <br>These settings apply only to firewalled gateways prior to NG. For later versions, these settings should be defined in the Node Properties window. 
* `http_servers` - (Optional) This list specifies the HTTP servers. Defining HTTP servers allows you to restrict incoming HTTP.http_servers blocks are documented below.
* `server_for_null_requests` - (Optional) The Logical Name of a Null Requests Server from http-servers. 


`vpn_authentication_and_encryption` supports the following:

* `encryption_algorithms` - (Optional) Select the methods negotiated in IKE phase 2 and used in IPSec connections.encryption_algorithms blocks are documented below.
* `encryption_method` - (Optional) Select the encryption method. 
* `pre_shared_secret` - (Optional) the user password is specified in the Authentication tab in the user's IKE properties (in the user properties window: Encryption tab > Edit). 
* `support_legacy_auth_for_sc_l2tp_nokia_clients` - (Optional) Support Legacy Authentication for SC (hybrid mode), L2TP (PAP) and Nokia clients (CRACK). 
* `support_legacy_eap` - (Optional) Support Legacy EAP (Extensible Authentication Protocol). 
* `support_l2tp_with_pre_shared_key` - (Optional) Use a centrally managed pre-shared key for IKE. 
* `l2tp_pre_shared_key` - (Optional) Type in the pre-shared key.<br>Available only if support-l2tp-with-pre-shared-key is set to true. 


`vpn_advanced` supports the following:

* `allow_clear_traffic_to_encryption_domain_when_disconnected` - (Optional) SecuRemote/SecureClient behavior while disconnected - How traffic to the VPN domain is handled when the Remote Access VPN client is not connected to the site. Traffic can either be dropped or sent in clear without encryption. 
* `enable_load_distribution_for_mep_conf` - (Optional) Load distribution for Multiple Entry Points configurations - Remote access clients will randomly select a gateway from the list of entry points. Make sure to define the same VPN domain for all the Security Gateways you want to be entry points. 
* `use_first_allocated_om_ip_addr_for_all_conn_to_the_gws_of_the_site` - (Optional) Use first allocated Office Mode IP Address for all connections to the Gateways of the site.After a remote user connects and receives an Office Mode IP address from a gateway, every connection to that gateways encryption domain will go out with the Office Mode IP as the internal source IP. The Office Mode IP is what hosts in the encryption domain will recognize as the remote user's IP address. The Office Mode IP address assigned by a specific gateway can be used in its own encryption domain and in neighboring encryption domains as well. The neighboring encryption domains should reside behind gateways that are members of the same VPN community as the assigning gateway. Since the remote hosts connections are dependant on the Office Mode IP address it received, should the gateway that issued the IP become unavailable, all the connections to the site will terminate. 


`scv` supports the following:

* `apply_scv_on_simplified_mode_fw_policies` - (Optional) Determine whether the gateway verifies that remote access clients are securely configured. This is set here only if the security policy is defined in the Simplified Mode. If the security policy is defined in the Traditional Mode, verification takes place per rule. 
* `exceptions` - (Optional) Specify the hosts that can be accessed using the selected services even if the client is not verified.<br>Available only if apply-scv-on-simplified-mode-fw-policies is true.exceptions blocks are documented below.
* `no_scv_for_unsupported_cp_clients` - (Optional) Do not apply Secure Configuration Verification for connections from Check Point VPN clients that don't support it, such as SSL Network Extender, GO, Capsule VPN / Connect, Endpoint Connects lower than R75, or L2TP clients.<br>Available only if apply-scv-on-simplified-mode-fw-policies is true. 
* `upon_verification_accept_and_log_client_connection` - (Optional) If the gateway verifies the client's configuration, decide how the gateway should handle connections with clients that fail the Security Configuration Verification. It is possible to either drop the connection or Accept the connection and log it. 
* `only_tcp_ip_protocols_are_used` - (Optional) Most SCV checks are configured via the SCV policy. Specify whether to verify that  only TCP/IP protocols are used. 
* `policy_installed_on_all_interfaces` - (Optional) Most SCV checks are configured via the SCV policy. Specify whether to verify that  the Desktop Security Policy is installed on all the interfaces of the client. 
* `generate_log` - (Optional) If the client identifies that the secure configuration has been violated, select whether a log is generated by the remote access client and sent to the Security Management server. 
* `notify_user` - (Optional) If the client identifies that the secure configuration has been violated, select whether to user should be notified. 


`ssl_network_extender` supports the following:

* `user_auth_method` - (Optional) Wide Impact: Also applies for SecureClient Mobile devices and Check Point GO clients!<br>User authentication method indicates how the user will be authenticated by the gateway. Changes made here will also apply for SSL clients.<br>Legacy - Username and password only.<br>Certificate - Certificate only with an existing certificate.<br>Certificate with Enrollment - Allows you to obtain a new certificate and then use certificate authentication only.<br>Mixed - Can use either username and password or certificate. 
* `supported_encryption_methods` - (Optional) Wide Impact: Also applies to SecureClient Mobile devices!<br>Select the encryption algorithms that will be supported for remote users. Changes made here will also apply for all SSL clients. 
* `client_upgrade_upon_connection` - (Optional) When a client connects to the gateway with SSL Network Extender, the client automatically checks for upgrade. Select whether the client should automatically upgrade. 
* `client_uninstall_upon_disconnection` - (Optional) Select whether the client should automatically uninstall SSL Network Extender when it disconnects from the gateway. 
* `re_auth_user_interval` - (Optional) Wide Impact: Applies for the SecureClient Mobile!<br>Select the interval that users will need to reauthenticate. 
* `scan_ep_machine_for_compliance_with_ep_compliance_policy` - (Optional) Set to true if you want endpoint machines to be scanned for compliance with the Endpoint Compliance Policy. 
* `client_outgoing_keep_alive_packets_frequency` - (Optional) Select the interval which the keep-alive packets are sent. 


`secure_client_mobile` supports the following:

* `user_auth_method` - (Optional) Wide Impact: Also applies for SSL Network Extender clients and Check Point GO clients.<br>How the user will be authenticated by the gateway. 
* `enable_password_caching` - (Optional) If the password entered to authenticate is saved locally on the user's machine. 
* `cache_password_timeout` - (Optional) Cached password timeout (in minutes). 
* `re_auth_user_interval` - (Optional) Wide Impact: Also applies for SSL Network Extender clients!<br>The length of time (in minutes) until the user's credentials are resent to the gateway to verify authorization. 
* `connect_mode` - (Optional) Methods by which a connection to the gateway will be initiated:<br>Configured On Endpoint Client - the method used for initiating a connection to a gateway is determined by the endpoint client<br>Manual - VPN connections will not be initiated automatically.<br>Always connected - SecureClient Mobile will automatically establish a connection to the last connected gateway under the following circumstances: (a) the device has a valid IP address, (b) when the device "wakes up" from a low-power state or a soft-reset, or (c) after a condition that caused the device to automatically disconnect ceases to exist (for example, Device is out of PC Sync, Disconnect is not idle.).<br>On application request - Applications requiring access to resources through the VPN will be able to initiate a VPN connection. 
* `automatically_initiate_dialup` - (Optional) When selected, the client will initiate a GPRS dialup connection before attempting to establish the VPN connection. Note that if a local IP address is already available through another network interface, then the GPRS dialup is not initiated. 
* `disconnect_when_device_is_idle` - (Optional) Enabling this feature will disconnect users from the gateway if there is no traffic sent during the defined time period. 
* `supported_encryption_methods` - (Optional) Wide Impact: Also applies for SSL Network Extender clients!<br>Select the encryption algorithms that will be supported with remote users. 
* `route_all_traffic_to_gw` - (Optional) Operates the client in Hub Mode, sending all traffic to the VPN server for routing, filtering, and processing. 


`endpoint_connect` supports the following:

* `enable_password_caching` - (Optional) If the password entered to authenticate is saved locally on the user's machine. 
* `cache_password_timeout` - (Optional) Cached password timeout (in minutes). 
* `re_auth_user_interval` - (Optional) The length of time (in minutes) until the user's credentials are resent to the gateway to verify authorization. 
* `connect_mode` - (Optional) Methods by which a connection to the gateway will be initiated:<br>Manual - VPN connections will not be initiated automatically.<br>Always connected - Endpoint Connect will automatically establish a connection to the last connected gateway under the following circumstances: (a) the device has a valid IP address, (b) when the device "wakes up" from a low-power state or a soft-reset, or (c) after a condition that caused the device to automatically disconnect ceases to exist (for example, Device is out of PC Sync, Disconnect is not idle.).<br>Configured on endpoint client - the method used for initiating a connection to a gateway is determined by the endpoint client. 
* `network_location_awareness` - (Optional) Wide Impact: Also applies for Check Point GO clients!<br>Endpoint Connect intelligently detects whether it is inside or outside of the VPN domain (Enterprise LAN), and automatically connects or disconnects as required. Select true and edit network-location-awareness-conf to configure this capability. 
* `network_location_awareness_conf` - (Optional) Configure how the client determines its location in relation to the internal network.network_location_awareness_conf blocks are documented below.
* `disconnect_when_conn_to_network_is_lost` - (Optional) Enabling this feature disconnects users from the gateway when connectivity to the network is lost. 
* `disconnect_when_device_is_idle` - (Optional) Enabling this feature will disconnect users from the gateway if there is no traffic sent during the defined time period. 
* `route_all_traffic_to_gw` - (Optional) Operates the client in Hub Mode, sending all traffic to the VPN server for routing, filtering, and processing. 
* `client_upgrade_mode` - (Optional) Select an option to determine how the client is upgraded. 


`hot_spot_and_hotel_registration` supports the following:

* `enable_registration` - (Optional) Set Enable registration to true in order to configure settings. Set Enable registration to false in order to cancel registration (the configurations below won't be available). When the feature is enabled, you have several minutes to complete registration. 
* `local_subnets_access_only` - (Optional) Local subnets access only. 
* `registration_timeout` - (Optional) Maximum time (in seconds) to complete registration. 
* `track_log` - (Optional) Track log. 
* `max_ip_access_during_registration` - (Optional) Maximum number of addresses to allow access to during registration. 
* `ports` - (Optional) Ports to be opened during registration (up to 10 ports).ports blocks are documented below.


`alerts` supports the following:

* `send_popup_alert_to_smartview_monitor` - (Optional) Send popup alert to SmartView Monitor when an alert is issued, it is also sent to SmartView Monitor. 
* `popup_alert_script` - (Optional) Run popup alert script the operating system script to be executed when an alert is issued. For example, set another form of notification, such as an email or a user-defined command. 
* `send_mail_alert_to_smartview_monitor` - (Optional) Send mail alert to SmartView Monitor when a mail alert is issued, it is also sent to SmartView Monitor. 
* `mail_alert_script` - (Optional) Run mail alert script the operating system script to be executed when Mail is specified as the Track in a rule. The default is internal_sendmail, which is not a script but an internal Security Gateway command. 
* `send_snmp_trap_alert_to_smartview_monitor` - (Optional) Send SNMP trap alert to SmartView Monitor when an SNMP trap alert is issued, it is also sent to SmartView Monitor. 
* `snmp_trap_alert_script` - (Optional) Run SNMP trap alert script command to be executed when SNMP Trap is specified as the Track in a rule. By default the internal_snmp_trap is used. This command is executed by the fwd process. 
* `send_user_defined_alert_num1_to_smartview_monitor` - (Optional) Send user defined alert no. 1 to SmartView Monitor when an alert is issued, it is also sent to SmartView Monitor. 
* `user_defined_script_num1` - (Optional) Run user defined script the operating system script to be run when User-Defined is specified as the Track in a rule, or when User Defined Alert no. 1 is selected as a Track Option. 
* `send_user_defined_alert_num2_to_smartview_monitor` - (Optional) Send user defined alert no. 2 to SmartView Monitor when an alert is issued, it is also sent to SmartView Monitor. 
* `user_defined_script_num2` - (Optional) Run user defined 2 script the operating system script to be run when User-Defined is specified as the Track in a rule, or when User Defined Alert no. 2 is selected as a Track Option. 
* `send_user_defined_alert_num3_to_smartview_monitor` - (Optional) Send user defined alert no. 3 to SmartView Monitor when an alert is issued, it is also sent to SmartView Monitor. 
* `user_defined_script_num3` - (Optional) Run user defined 3 script the operating system script to be run when User-Defined is specified as the Track in a rule, or when User Defined Alert no. 3 is selected as a Track Option. 
* `default_track_option_for_system_alerts` - (Optional) Set the default track option for System Alerts. 


`time_settings` supports the following:

* `excessive_log_grace_period` - (Optional) Specifies the minimum amount of time (in seconds) between consecutive logs of similar packets. Two packets are considered similar if they have the same source address, source port, destination address, and destination port; and the same protocol was used. After the first packet, similar packets encountered in the grace period will be acted upon according to the security policy, but only the first packet generates a log entry or an alert. Any value from 0 to 90 seconds can be entered in this field.<br>Note: This option only applies for DROP rules with logging. 
* `logs_resolving_timeout` - (Optional) Specifies the amount of time (in seconds), after which the log page is displayed without resolving names and while showing only IP addresses. Any value from 0 to 90 seconds can be entered in this field. 
* `status_fetching_interval` - (Optional) Specifies the frequency at which the Security Management server queries the Check Point Security gateway, Check Point QoS and other gateways it manages for status information. Any value from 30 to 900 seconds can be entered in this field. 
* `virtual_link_statistics_logging_interval` - (Optional) Specifies the frequency (in seconds) with which Virtual Link statistics will be logged. This parameter is relevant only for Virtual Links defined with SmartView Monitor statistics enabled in the SLA Parameters tab of the Virtual Link window. Any value from 60 to 3600 seconds can be entered in this field. 


`certs_and_pki` supports the following:

* `cert_validation_enforce_key_size` - (Optional) Enforce key length in certificate validation (R80+ gateways only). 
* `host_certs_ecdsa_key_size` - (Optional) Select the key size for ECDSA of the host certificate. 
* `host_certs_key_size` - (Optional) Select the key size of the host certificate. 


`http_servers` supports the following:

* `logical_name` - (Optional) Unique Logical Name of the HTTP Server. 
* `host` - (Optional) Host name of the HTTP Server. 
* `port` - (Optional) Port number of the HTTP Server. 
* `reauthentication` - (Optional) Specify whether users must reauthenticate when accessing a specific server. 


`encryption_algorithms` supports the following:

* `ike` - (Optional) Configure the IKE Phase 1 settings.ike blocks are documented below.
* `ipsec` - (Optional) Configure the IPSEC Phase 2 settings.ipsec blocks are documented below.


`exceptions` supports the following:

* `hosts` - (Optional) Specify the Hosts to be excluded from SCV.hosts blocks are documented below.
* `services` - (Optional) Specify the services to be accessed.services blocks are documented below.


`network_location_awareness_conf` supports the following:

* `vpn_clients_are_considered_inside_the_internal_network_when_the_client` - (Optional) When a VPN client is within the internal network, the internal resources are available and the VPN tunnel should be disconnected. Determine when VPN clients are considered inside the internal network:<br>Connects to GW through internal interface - The client connects to the gateway through one of its internal interfaces (recommended).<br>Connects from network or group - The client connects from a network or group specified in network-or-group-of-conn-vpn-client.<br>Runs on computer with access to Active Directory domain - The client runs on a computer that can access its Active Directory domain.<br>Note: The VPN tunnel will resume automatically when the VPN client is no longer in the internal network and the client is set to "Always connected" mode. 
* `network_or_group_of_conn_vpn_client` - (Optional) Name or UID of Network or Group the VPN client is connected from.<br>Available only if vpn-clients-are-considered-inside-the-internal-network-when-the-client is set to "Connects from network or group". 
* `consider_wireless_networks_as_external` - (Optional) The speed at which locations are classified as internal or external can be increased by creating a list of wireless networks that are known to be external. A wireless network is identified by its Service Set Identifier (SSID) a name used to identify a particular 802.11 wireless LAN. 
* `excluded_internal_wireless_networks` - (Optional) Excludes the specified internal networks names (SSIDs).<br>Available only if consider-wireless-networks-as-external is set to true.excluded_internal_wireless_networks blocks are documented below.
* `consider_undefined_dns_suffixes_as_external` - (Optional) The speed at which locations are classified as internal or external can be increased by creating a list of DNS suffixes that are known to be external. Enable this to be able to define DNS suffixes which won't be considered external. 
* `dns_suffixes` - (Optional) DNS suffixes not defined here will be considered as external. If this list is empty consider-undefined-dns-suffixes-as-external will automatically be set to false.<br>Available only if consider-undefined-dns-suffixes-as-external is set to true.dns_suffixes blocks are documented below.
* `remember_previously_detected_external_networks` - (Optional) The speed at which locations are classified as internal or external can be increased by caching (on the client side) names of networks that were previously determined to be external. 


`ike` supports the following:

* `support_encryption_algorithms` - (Optional) Select the encryption algorithms that will be supported with remote hosts.support_encryption_algorithms blocks are documented below.
* `use_encryption_algorithm` - (Optional) Choose the encryption algorithm that will have the highest priority of the selected algorithms. If given a choice of more that one encryption algorithm to use, the algorithm selected in this field will be used. 
* `support_data_integrity` - (Optional) Select the hash algorithms that will be supported with remote hosts to ensure data integrity.support_data_integrity blocks are documented below.
* `use_data_integrity` - (Optional) The hash algorithm chosen here will be given the highest priority if more than one choice is offered. 
* `support_diffie_hellman_groups` - (Optional) Select the Diffie-Hellman groups that will be supported with remote hosts.support_diffie_hellman_groups blocks are documented below.
* `use_diffie_hellman_group` - (Optional) SecureClient users utilize the Diffie-Hellman group selected in this field. 


`ipsec` supports the following:

* `support_encryption_algorithms` - (Optional) Select the encryption algorithms that will be supported with remote hosts.support_encryption_algorithms blocks are documented below.
* `use_encryption_algorithm` - (Optional) Choose the encryption algorithm that will have the highest priority of the selected algorithms. If given a choice of more that one encryption algorithm to use, the algorithm selected in this field will be used. 
* `support_data_integrity` - (Optional) Select the hash algorithms that will be supported with remote hosts to ensure data integrity.support_data_integrity blocks are documented below.
* `use_data_integrity` - (Optional) The hash algorithm chosen here will be given the highest priority if more than one choice is offered. 
* `enforce_encryption_alg_and_data_integrity_on_all_users` - (Optional) Enforce Encryption Algorithm and Data Integrity on all users. 


`support_encryption_algorithms` supports the following:

* `aes_128` - (Optional) Select whether the AES-128 encryption algorithm will be supported with remote hosts. 
* `aes_256` - (Optional) Select whether the AES-256 encryption algorithm will be supported with remote hosts. 
* `des` - (Optional) Select whether the DES encryption algorithm will be supported with remote hosts. 
* `tdes` - (Optional) Select whether the Triple DES encryption algorithm will be supported with remote hosts. 


`support_data_integrity` supports the following:

* `aes_xcbc` - (Optional) Select whether the AES-XCBC hash algorithm will be supported with remote hosts to ensure data integrity. 
* `md5` - (Optional) Select whether the MD5 hash algorithm will be supported with remote hosts to ensure data integrity. 
* `sha1` - (Optional) Select whether the SHA1 hash algorithm will be supported with remote hosts to ensure data integrity. 
* `sha256` - (Optional) Select whether the SHA256 hash algorithm will be supported with remote hosts to ensure data integrity. 


`support_diffie_hellman_groups` supports the following:

* `group1` - (Optional) Select whether Diffie-Hellman Group 1 (768 bit) will be supported with remote hosts. 
* `group14` - (Optional) Select whether Diffie-Hellman Group 14 (2048 bit) will be supported with remote hosts. 
* `group2` - (Optional) Select whether Diffie-Hellman Group 2 (1024 bit) will be supported with remote hosts. 
* `group5` - (Optional) Select whether Diffie-Hellman Group 5 (1536 bit) will be supported with remote hosts. 


`support_encryption_algorithms` supports the following:

* `aes_128` - (Optional) Select whether the AES-128 encryption algorithm will be supported with remote hosts. 
* `aes_256` - (Optional) Select whether the AES-256 encryption algorithm will be supported with remote hosts. 
* `des` - (Optional) Select whether the DES encryption algorithm will be supported with remote hosts. 
* `tdes` - (Optional) Select whether the Triple DES encryption algorithm will be supported with remote hosts. 


`support_data_integrity` supports the following:

* `aes_xcbc` - (Optional) Select whether the AES-XCBC hash algorithm will be supported with remote hosts to ensure data integrity. 
* `md5` - (Optional) Select whether the MD5 hash algorithm will be supported with remote hosts to ensure data integrity. 
* `sha1` - (Optional) Select whether the SHA1 hash algorithm will be supported with remote hosts to ensure data integrity. 
* `sha256` - (Optional) Select whether the SHA256 hash algorithm will be supported with remote hosts to ensure data integrity. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

