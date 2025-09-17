---
layout: "checkpoint"
page_title: "checkpoint_management_simple_cluster"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-simple-cluster"
description: |-
This resource allows you to execute Check Point Simple Cluster.
---

# checkpoint_management_simple_cluster

This resource allows you to execute Check Point Simple Cluster.

## Example Usage


```hcl
resource "checkpoint_management_simple_cluster" "example" {
  name = "cluster1"
  color = "yellow"
  version = "R80.30"
  os_name = "Gaia"
  cluster_mode = "cluster-xl-ha"
  firewall = true
  ipv4_address = "17.23.5.1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `advanced_settings` - (Optional) N/Aadvanced_settings blocks are documented below.
* `anti_bot` - (Optional) Anti-Bot blade enabled. 
* `anti_virus` - (Optional) Anti-Virus blade enabled. 
* `application_control` - (Optional) Application Control blade enabled. 
* `application_control_and_url_filtering_settings` - (Optional) Gateway Application Control and URL filtering settings.application_control_and_url_filtering_settings blocks are documented below.
* `cluster_mode` - (Optional) Cluster mode. 
* `cluster_settings` - (Optional) ClusterXL and VRRP Settings.cluster_settings blocks are documented below.
* `content_awareness` - (Optional) Content Awareness blade enabled. 
* `enable_https_inspection` - (Optional) Enable HTTPS Inspection after defining an outbound inspection certificate. <br>To define the outbound certificate use outbound inspection certificate API. 
* `fetch_policy` - (Optional) Security management server(s) to fetch the policy from.fetch_policy blocks are documented below.
* `firewall` - (Optional) Firewall blade enabled. 
* `firewall_settings` - (Optional) N/Afirewall_settings blocks are documented below.
* `geo_mode` - (Optional) Cluster High Availability Geo mode.<br>This setting applies only to a cluster deployed in a cloud. Available when the cluster mode equals "cluster-xl-ha". 
* `hardware` - (Optional) Cluster platform hardware. 
* `hit_count` - (Optional) Hit count tracks the number of connections each rule matches. 
* `https_inspection` - (Optional) HTTPS inspection.https_inspection blocks are documented below.
* `identity_awareness` - (Optional) Identity awareness blade enabled. 
* `identity_awareness_settings` - (Optional) Gateway Identity Awareness settings.identity_awareness_settings blocks are documented below.
* `interfaces` - (Optional) Cluster interfaces.interfaces blocks are documented below.
* `ipv4_address` - (Optional) IPv4 address. 
* `ipv6_address` - (Optional) IPv6 address. 
* `ips` - (Optional) Intrusion Prevention System blade enabled. 
* `ips_settings` - (Optional) Cluster IPS settings. ips_settings blocks are documented below.
* `ips_update_policy` - (Optional) Specifies whether the IPS will be downloaded from the Management or directly to the Gateway. 
* `members` - (Optional) Cluster members list. Only new cluster member can be added. Adding existing gateway is not supported.members blocks are documented below.
* `nat_hide_internal_interfaces` - (Optional) Hide internal networks behind the Gateway's external IP. 
* `nat_settings` - (Optional) NAT settings.nat_settings blocks are documented below.
* `os_name` - (Optional) Cluster platform operating system. 
* `platform_portal_settings` - (Optional) Platform portal settings.platform_portal_settings blocks are documented below.
* `proxy_settings` - (Optional) Proxy Server for Gateway.proxy_settings blocks are documented below.
* `qos` - (Optional) QoS. 
* `send_alerts_to_server` - (Optional) Server(s) to send alerts to.send_alerts_to_server blocks are documented below.
* `send_logs_to_backup_server` - (Optional) Backup server(s) to send logs to.send_logs_to_backup_server blocks are documented below.
* `send_logs_to_server` - (Optional) Server(s) to send logs to.send_logs_to_server blocks are documented below.
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `threat_emulation` - (Optional) Threat Emulation blade enabled. 
* `threat_extraction` - (Optional) Threat Extraction blade enabled. 
* `threat_prevention_mode` - (Optional) The mode of Threat Prevention to use. When using Autonomous Threat Prevention, disabling the Threat Prevention blades is not allowed. 
* `url_filtering` - (Optional) URL Filtering blade enabled. 
* `usercheck_portal_settings` - (Optional) UserCheck portal settings.usercheck_portal_settings blocks are documented below.
* `version` - (Optional) Cluster platform version. 
* `vpn` - (Optional) VPN blade enabled. 
* `vpn_settings` - (Optional) Gateway VPN settings.vpn_settings blocks are documented below.
* `zero_phishing` - (Optional) Zero Phishing blade enabled. 
* `zero_phishing_fqdn` - (Optional) Zero Phishing gateway FQDN. 
* `show_portals_certificate` - (Optional) Indicates whether to show the portals certificate value in the reply. 
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `groups` - (Optional) Collection of group identifiers.groups blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


`advanced_settings` supports the following:

* `connection_persistence` - (Optional) Handling established connections when installing a new policy. 
* `sam` - (Optional) SAM.sam blocks are documented below.


`application_control_and_url_filtering_settings` supports the following:

* `global_settings_mode` - (Optional) Whether to override global settings or not. 
* `override_global_settings` - (Optional) override global settings object.override_global_settings blocks are documented below.


`cluster_settings` supports the following:

* `member_recovery_mode` - (Optional) In a High Availability cluster, each member is given a priority. The member with the highest priority serves as the gateway. If this gateway fails, control is passed to the member with the next highest priority. If that member fails, control is passed to the next, and so on. Upon gateway recovery, it is possible to:
Maintain current active Cluster Member (maintain-current-active) or
Switch to higher priority Cluster Member (according-to-priority). 
* `state_synchronization` - (Optional) Cluster State Synchronization settings.state_synchronization blocks are documented below.
* `track_changes_of_cluster_members` - (Optional) Track changes in the status of Cluster Members. 
* `use_virtual_mac` - (Optional) Use Virtual MAC. By enabling Virtual MAC in ClusterXL High Availability New mode, or Load Sharing Unicast mode, all cluster members associate the same Virtual MAC address with All Cluster Virtual Interfaces and the Virtual IP address. 


`firewall_settings` supports the following:

* `auto_calculate_connections_hash_table_size_and_memory_pool` - (Optional) N/A 
* `auto_maximum_limit_for_concurrent_connections` - (Optional) N/A 
* `connections_hash_size` - (Optional) N/A 
* `maximum_limit_for_concurrent_connections` - (Optional) N/A 
* `maximum_memory_pool_size` - (Optional) N/A 
* `memory_pool_size` - (Optional) N/A 


`https_inspection` supports the following:

* `bypass_on_failure` - (Optional) Set to be true in order to bypass all requests (Fail-open) in case of internal system error.bypass_on_failure blocks are documented below.
* `site_categorization_allow_mode` - (Optional) Set to 'background' in order to allowed requests until categorization is complete.site_categorization_allow_mode blocks are documented below.
* `deny_untrusted_server_cert` - (Optional) Set to be true in order to drop traffic from servers with untrusted server certificate.deny_untrusted_server_cert blocks are documented below.
* `deny_revoked_server_cert` - (Optional) Set to be true in order to drop traffic from servers with revoked server certificate (validate CRL).deny_revoked_server_cert blocks are documented below.
* `deny_expired_server_cert` - (Optional) Set to be true in order to drop traffic from servers with expired server certificate.deny_expired_server_cert blocks are documented below.


`identity_awareness_settings` supports the following:

* `browser_based_authentication` - (Optional) Enable Browser Based Authentication source. 
* `browser_based_authentication_settings` - (Optional) Browser Based Authentication settings.browser_based_authentication_settings blocks are documented below.
* `identity_agent` - (Optional) Enable Identity Agent source. 
* `identity_agent_settings` - (Optional) Identity Agent settings.identity_agent_settings blocks are documented below.
* `identity_collector` - (Optional) Enable Identity Collector source. 
* `identity_collector_settings` - (Optional) Identity Collector settings.identity_collector_settings blocks are documented below.
* `identity_sharing_settings` - (Optional) Identity sharing settings.identity_sharing_settings blocks are documented below.
* `proxy_settings` - (Optional) Identity-Awareness Proxy settings.proxy_settings blocks are documented below.
* `remote_access` - (Optional) Enable Remote Access Identity source. 


`interfaces` supports the following:

* `name` - (Optional) Object name. Must be unique in the domain. 
* `interface_type` - (Optional) Cluster interface type. 
* `ipv4_address` - (Optional) IPv4 address. 
* `ipv6_address` - (Optional) IPv6 address. 
* `network_mask` - (Optional) IPv4 or IPv6 network mask. If both masks are required use ipv4-network-mask and ipv6-network-mask fields explicitly. Instead of providing mask itself it is possible to specify IPv4 or IPv6 mask length in mask-length field. If both masks length are required use ipv4-mask-length and  ipv6-mask-length fields explicitly. 
* `ipv4_network_mask` - (Optional) IPv4 network address. 
* `ipv6_network_mask` - (Optional) IPv6 network address. 
* `ipv4_mask_length` - (Optional) IPv4 network mask length. 
* `ipv6_mask_length` - (Optional) IPv6 network mask length. 
* `anti_spoofing` - (Optional) N/A 
* `anti_spoofing_settings` - (Optional) N/Aanti_spoofing_settings blocks are documented below.
* `multicast_address` - (Optional) Multicast IP Address. 
* `multicast_address_type` - (Optional) Multicast Address Type. 
* `security_zone` - (Optional) N/A 
* `security_zone_settings` - (Optional) N/Asecurity_zone_settings blocks are documented below.
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `topology` - (Optional) N/A 
* `topology_settings` - (Optional) N/Atopology_settings blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 

`ips_settings` supports the following:

* `bypass_all_under_load` - (Optional) Disable/enable all IPS protections until CPU and memory levels are back to normal.
* `bypass_track_method` - (Optional) Track options when all IPS protections are disabled until CPU/memory levels are back to normal.
* `top_cpu_consuming_protections` - (Optional) Provides a way to reduce CPU levels on machines under load by disabling the top CPU consuming IPS protections.top_cpu_consuming_protections blocks are documented below.
* `activation_mode` - (Optional) Defines whether the IPS blade operates in Detect Only mode or enforces the configured IPS Policy.
* `cpu_usage_low_threshold` - (Optional) CPU usage low threshold percentage (1-99).
* `cpu_usage_high_threshold` - (Optional) CPU usage high threshold percentage (1-99).
* `memory_usage_low_threshold` - (Optional) Memory usage low threshold percentage (1-99).
* `memory_usage_high_threshold` - (Optional) Memory usage high threshold percentage (1-99).
* `send_threat_cloud_info` - (Optional) Help improve Check Point Threat Prevention product by sending anonymous information.
* `reject_on_cluster_fail_over` - (Optional) Define the IPS connections during fail over reject packets or accept packets.

`members` supports the following:

* `name` - (Optional) Object name. 
* `interfaces` - (Optional) Cluster Member network interfaces.interfaces blocks are documented below.
* `ipv4_address` - (Optional) IPv4 address. 
* `ipv6_address` - (Optional) IPv6 address. 
* `one_time_password` - (Optional) N/A 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `priority` - (Optional) In a High Availability New mode cluster each machine is given a priority. The highest priority machine serves as the gateway in normal circumstances. If this machine fails, control is passed to the next highest priority machine. If that machine fails, control is passed to the next machine, and so on.
In Load Sharing Unicast mode cluster, the highest priority is the pivot machine.
The values must be in a range from 1 to N, where N is number of cluster members. 
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


`nat_settings` supports the following:

* `auto_rule` - (Optional) Whether to add automatic address translation rules. 
* `ipv4_address` - (Optional) IPv4 address. 
* `ipv6_address` - (Optional) IPv6 address. 
* `hide_behind` - (Optional) Hide behind method. This parameter is forbidden in case "method" parameter is "static". 
* `install_on` - (Optional) Which gateway should apply the NAT translation. 
* `method` - (Optional) NAT translation method. 


`platform_portal_settings` supports the following:

* `portal_web_settings` - (Optional) Configuration of the portal web settings.portal_web_settings blocks are documented below.
* `certificate_settings` - (Optional) Configuration of the portal certificate settings.certificate_settings blocks are documented below.
* `accessibility` - (Optional) Configuration of the portal access settings.accessibility blocks are documented below.


`proxy_settings` supports the following:

* `use_custom_proxy` - (Optional) Use custom proxy settings for this network object. 
* `proxy_server` - (Optional) N/A 
* `port` - (Optional) N/A 


`usercheck_portal_settings` supports the following:

* `enabled` - (Optional) State of the web portal (enabled or disabled). The supported blades are: {'Application Control', 'URL Filtering', 'Data Loss Prevention', 'Anti Virus', 'Anti Bot', 'Threat Emulation', 'Threat Extraction', 'Data Awareness'}. 
* `portal_web_settings` - (Optional) Configuration of the portal web settings.portal_web_settings blocks are documented below.
* `certificate_settings` - (Optional) Configuration of the portal certificate settings.certificate_settings blocks are documented below.
* `accessibility` - (Optional) Configuration of the portal access settings.accessibility blocks are documented below.


`vpn_settings` supports the following:

* `authentication` - (Optional) Authentication.authentication blocks are documented below.
* `link_selection` - (Optional) Link Selection.link_selection blocks are documented below.
* `maximum_concurrent_ike_negotiations` - (Optional) N/A 
* `maximum_concurrent_tunnels` - (Optional) N/A 
* `office_mode` - (Optional) Office Mode.
Notation Wide Impact - Office Mode apply IPSec VPN Software Blade clients and to the Mobile Access Software Blade clients.office_mode blocks are documented below.
* `remote_access` - (Optional) Remote Access.remote_access blocks are documented below.
* `vpn_domain` - (Optional) Gateway VPN domain identified by the name or UID. 
* `vpn_domain_exclude_external_ip_addresses` - (Optional) Exclude the external IP addresses from the VPN domain of this Security Gateway. 
* `vpn_domain_type` - (Optional) Gateway VPN domain type. 


`sam` supports the following:

* `forward_to_other_sam_servers` - (Optional) Forward SAM clients' requests to other SAM servers. 
* `use_early_versions` - (Optional) Use early versions compatibility mode.use_early_versions blocks are documented below.
* `purge_sam_file` - (Optional) Purge SAM File.purge_sam_file blocks are documented below.


`override_global_settings` supports the following:

* `fail_mode` - (Optional) Fail mode - allow or block all requests. 
* `website_categorization` - (Optional) Website categorization object.website_categorization blocks are documented below.


`state_synchronization` supports the following:

* `delayed` - (Optional) Start synchronizing with delay of seconds, as defined by delayed-seconds, after connection initiation. Disabled when state-synchronization disabled. 
* `delayed_seconds` - (Optional) Start synchronizing X seconds after connection initiation
. The values must be in a range between 2 and 3600. 
* `enabled` - (Optional) Use State Synchronization. 


`bypass_on_failure` supports the following:

* `override_profile` - (Optional) Override profile of global configuration. 
* `value` - (Optional) Override value.<br><font color="red">Required only for</font> 'override-profile' is True. 


`site_categorization_allow_mode` supports the following:

* `override_profile` - (Optional) Override profile of global configuration. 
* `value` - (Optional) Override value.<br><font color="red">Required only for</font> 'override-profile' is True. 


`deny_untrusted_server_cert` supports the following:

* `override_profile` - (Optional) Override profile of global configuration. 
* `value` - (Optional) Override value.<br><font color="red">Required only for</font> 'override-profile' is True. 


`deny_revoked_server_cert` supports the following:

* `override_profile` - (Optional) Override profile of global configuration. 
* `value` - (Optional) Override value.<br><font color="red">Required only for</font> 'override-profile' is True. 


`deny_expired_server_cert` supports the following:

* `override_profile` - (Optional) Override profile of global configuration. 
* `value` - (Optional) Override value.<br><font color="red">Required only for</font> 'override-profile' is True. 


`browser_based_authentication_settings` supports the following:

* `authentication_settings` - (Optional) Authentication Settings for Browser Based Authentication.authentication_settings blocks are documented below.
* `browser_based_authentication_portal_settings` - (Optional) Browser Based Authentication portal settings.browser_based_authentication_portal_settings blocks are documented below.


`identity_agent_settings` supports the following:

* `agents_interval_keepalive` - (Optional) Agents send keepalive period (minutes). 
* `user_reauthenticate_interval` - (Optional) Agent reauthenticate time interval (minutes). 
* `authentication_settings` - (Optional) Authentication Settings for Identity Agent.authentication_settings blocks are documented below.
* `identity_agent_portal_settings` - (Optional) Identity Agent accessibility settings.identity_agent_portal_settings blocks are documented below.


`identity_collector_settings` supports the following:

* `authorized_clients` - (Optional) Authorized Clients.authorized_clients blocks are documented below.
* `authentication_settings` - (Optional) Authentication Settings for Identity Collector.authentication_settings blocks are documented below.
* `client_access_permissions` - (Optional) Identity Collector accessibility settings.client_access_permissions blocks are documented below.


`identity_sharing_settings` supports the following:

* `share_with_other_gateways` - (Optional) Enable identity sharing with other gateways. 
* `receive_from_other_gateways` - (Optional) Enable receiving identity from other gateways. 
* `receive_from` - (Optional) Gateway(s) to receive identity from.receive_from blocks are documented below.


`proxy_settings` supports the following:

* `detect_using_x_forward_for` - (Optional) Whether to use X-Forward-For HTTP header, which is added by the proxy server to keep track of the original source IP. 


`anti_spoofing_settings` supports the following:

* `action` - (Optional) If packets will be rejected (the Prevent option) or whether the packets will be monitored (the Detect option). 
* `exclude_packets` - (Optional) Don't check packets from excluded network. 
* `excluded_network_name` - (Optional) Excluded network name. 
* `excluded_network_uid` - (Optional) Excluded network UID. 
* `spoof_tracking` - (Optional) Spoof tracking. 


`security_zone_settings` supports the following:

* `auto_calculated` - (Optional) Security Zone is calculated according to where the interface leads to. 
* `specific_zone` - (Optional) Security Zone specified manually. 


`topology_settings` supports the following:

* `interface_leads_to_dmz` - (Optional) Whether this interface leads to demilitarized zone (perimeter network). 
* `specific_network` - (Optional) Network behind this interface. 


`interfaces` supports the following:

* `name` - (Optional) Object name. 
* `anti_spoofing` - (Optional) N/A 
* `anti_spoofing_settings` - (Optional) N/Aanti_spoofing_settings blocks are documented below.
* `ipv4_address` - (Optional) IPv4 address. 
* `ipv6_address` - (Optional) IPv6 address. 
* `network_mask` - (Optional) IPv4 or IPv6 network mask. If both masks are required use ipv4-network-mask and ipv6-network-mask fields explicitly. Instead of providing mask itself it is possible to specify IPv4 or IPv6 mask length in mask-length field. If both masks length are required use ipv4-mask-length and  ipv6-mask-length fields explicitly. 
* `ipv4_network_mask` - (Optional) IPv4 network address. 
* `ipv6_network_mask` - (Optional) IPv6 network address. 
* `ipv4_mask_length` - (Optional) IPv4 network mask length. 
* `ipv6_mask_length` - (Optional) IPv6 network mask length. 
* `security_zone` - (Optional) N/A 
* `security_zone_settings` - (Optional) N/Asecurity_zone_settings blocks are documented below.
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `topology` - (Optional) N/A 
* `topology_settings` - (Optional) N/Atopology_settings blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


`portal_web_settings` supports the following:

* `aliases` - (Optional) List of URL aliases that are redirected to the main portal URL.aliases blocks are documented below.
* `main_url` - (Optional) The main URL for the web portal. 


`certificate_settings` supports the following:

* `base64_certificate` - (Optional) The certificate file encoded in Base64 with padding. 
This file must be in the *.p12 format. 
* `base64_password` - (Optional) Password (encoded in Base64 with padding) for the certificate file. 


`accessibility` supports the following:

* `allow_access_from` - (Optional) Allowed access to the web portal (based on interfaces, or security policy). 
* `internal_access_settings` - (Optional) Configuration of the additional portal access settings for internal interfaces only.internal_access_settings blocks are documented below.


`portal_web_settings` supports the following:

* `aliases` - (Optional) List of URL aliases that are redirected to the main portal URL.aliases blocks are documented below.
* `main_url` - (Optional) The main URL for the web portal. 


`certificate_settings` supports the following:

* `base64_certificate` - (Optional) The certificate file encoded in Base64 with padding. 
This file must be in the *.p12 format. 
* `base64_password` - (Optional) Password (encoded in Base64 with padding) for the certificate file. 


`accessibility` supports the following:

* `allow_access_from` - (Optional) Allowed access to the web portal (based on interfaces, or security policy). 
* `internal_access_settings` - (Optional) Configuration of the additional portal access settings for internal interfaces only.internal_access_settings blocks are documented below.


`authentication` supports the following:

* `authentication_clients` - (Optional) Collection of VPN Authentication clients identified by the name or UID.authentication_clients blocks are documented below.


`link_selection` supports the following:

* `dns_resolving_hostname` - (Optional) DNS Resolving Hostname. Must be set when "ip-selection" was selected to be "dns-resolving-from-hostname". 


`office_mode` supports the following:

* `mode` - (Optional) Office Mode Permissions.
When selected to be "off", all the other definitions are irrelevant. 
* `group` - (Optional) Group. Identified by name or UID.
Must be set when "office-mode-permissions" was selected to be "group". 
* `allocate_ip_address_from` - (Optional) Allocate IP address Method.
Allocate IP address by sequentially trying the given methods until success.allocate_ip_address_from blocks are documented below.
* `support_multiple_interfaces` - (Optional) Support connectivity enhancement for gateways with multiple external interfaces. 
* `perform_anti_spoofing` - (Optional) Perform Anti-Spoofing on Office Mode addresses. 
* `anti_spoofing_additional_addresses` - (Optional) Additional IP Addresses for Anti-Spoofing.
Identified by name or UID.
Must be set when "perform-anti-spoofings" is true. 


`remote_access` supports the following:

* `support_l2tp` - (Optional) Support L2TP (relevant only when office mode is active). 
* `l2tp_auth_method` - (Optional) L2TP Authentication Method.
Must be set when "support-l2tp" is true. 
* `l2tp_certificate` - (Optional) L2TP Certificate.
Must be set when "l2tp-auth-method" was selected to be "certificate".
Insert "defaultCert" when you want to use the default certificate. 
* `allow_vpn_clients_to_route_traffic` - (Optional) Allow VPN clients to route traffic. 
* `support_nat_traversal_mechanism` - (Optional) Support NAT traversal mechanism (UDP encapsulation). 
* `nat_traversal_service` - (Optional) Allocated NAT traversal UDP service. Identified by name or UID.
Must be set when "support-nat-traversal-mechanism" is true. 
* `support_visitor_mode` - (Optional) Support Visitor Mode. 
* `visitor_mode_service` - (Optional) TCP Service for Visitor Mode. Identified by name or UID.
Must be set when "support-visitor-mode" is true. 
* `visitor_mode_interface` - (Optional) Interface for Visitor Mode.
Must be set when "support-visitor-mode" is true.
Insert IPV4 Address of existing interface or "All IPs" when you want all interfaces. 


`use_early_versions` supports the following:

* `enabled` - (Optional) Use early versions compatibility mode. 
* `compatibility_mode` - (Optional) Early versions compatibility mode. 


`purge_sam_file` supports the following:

* `enabled` - (Optional) Purge SAM File. 
* `purge_when_size_reaches_to` - (Optional) Purge SAM File When it Reaches to. 


`website_categorization` supports the following:

* `mode` - (Optional) Website categorization mode. 
* `custom_mode` - (Optional) Custom mode object.custom_mode blocks are documented below.


`authentication_settings` supports the following:

* `authentication_method` - (Optional) Authentication method. 
* `identity_provider` - (Optional) Identity provider object identified by the name or UID. Must be set when "authentication-method" was selected to be "identity provider".identity_provider blocks are documented below.
* `radius` - (Optional) Radius server object identified by the name or UID. Must be set when "authentication-method" was selected to be "radius". 
* `users_directories` - (Optional) Users directories.users_directories blocks are documented below.


`browser_based_authentication_portal_settings` supports the following:

* `portal_web_settings` - (Optional) Configuration of the portal web settings.portal_web_settings blocks are documented below.
* `certificate_settings` - (Optional) Configuration of the portal certificate settings.certificate_settings blocks are documented below.
* `accessibility` - (Optional) Configuration of the portal access settings.accessibility blocks are documented below.


`authentication_settings` supports the following:

* `authentication_method` - (Optional) Authentication method. 
* `radius` - (Optional) Radius server object identified by the name or UID. Must be set when "authentication-method" was selected to be "radius". 
* `users_directories` - (Optional) Users directories.users_directories blocks are documented below.


`identity_agent_portal_settings` supports the following:

* `accessibility` - (Optional) Configuration of the portal access settings.accessibility blocks are documented below.


`authorized_clients` supports the following:

* `client` - (Optional) Host / Network Group Name or UID. 
* `client_secret` - (Optional) Client Secret. 


`authentication_settings` supports the following:

* `users_directories` - (Optional) Users directories.users_directories blocks are documented below.


`client_access_permissions` supports the following:

* `accessibility` - (Optional) Configuration of the portal access settings.accessibility blocks are documented below.


`anti_spoofing_settings` supports the following:

* `action` - (Optional) If packets will be rejected (the Prevent option) or whether the packets will be monitored (the Detect option). 
* `exclude_packets` - (Optional) Don't check packets from excluded network. 
* `excluded_network_name` - (Optional) Excluded network name. 
* `excluded_network_uid` - (Optional) Excluded network UID. 
* `spoof_tracking` - (Optional) Spoof tracking. 


`security_zone_settings` supports the following:

* `auto_calculated` - (Optional) Security Zone is calculated according to where the interface leads to. 
* `specific_zone` - (Optional) Security Zone specified manually. 


`topology_settings` supports the following:

* `interface_leads_to_dmz` - (Optional) Whether this interface leads to demilitarized zone (perimeter network). 
* `specific_network` - (Optional) Network behind this interface. 


`internal_access_settings` supports the following:

* `undefined` - (Optional) Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'. 
* `dmz` - (Optional) Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'. 
* `vpn` - (Optional) Controls portal access settings for interfaces that are part of a VPN Encryption Domain. 


`internal_access_settings` supports the following:

* `undefined` - (Optional) Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'. 
* `dmz` - (Optional) Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'. 
* `vpn` - (Optional) Controls portal access settings for interfaces that are part of a VPN Encryption Domain. 


`allocate_ip_address_from` supports the following:

* `radius_server` - (Optional) Radius server used to authenticate the user. 
* `use_allocate_method` - (Optional) Use Allocate Method. 
* `allocate_method` - (Optional) Using either Manual (IP Pool) or Automatic (DHCP).
Must be set when "use-allocate-method" is true. 
* `manual_network` - (Optional) Manual Network. Identified by name or UID.
Must be set when "allocate-method" was selected to be "manual". 
* `dhcp_server` - (Optional) DHCP Server. Identified by name or UID.
Must be set when "allocate-method" was selected to be "automatic". 
* `virtual_ip_address` - (Optional) Virtual IPV4 address for DHCP server replies.
Must be set when "allocate-method" was selected to be "automatic". 
* `dhcp_mac_address` - (Optional) Calculated MAC address for DHCP allocation.
Must be set when "allocate-method" was selected to be "automatic". 
* `optional_parameters` - (Optional) This configuration applies to all Office Mode methods except Automatic (using DHCP) and ipassignment.conf entries which contain this data.optional_parameters blocks are documented below.


`custom_mode` supports the following:

* `social_networking_widgets` - (Optional) Social networking widgets mode. 
* `url_filtering` - (Optional) URL filtering mode. 


`users_directories` supports the following:

* `external_user_profile` - (Optional) External user profile. 
* `internal_users` - (Optional) Internal users. 
* `users_from_external_directories` - (Optional) Users from external directories. 
* `specific` - (Optional) LDAP AU objects identified by the name or UID. Must be set when "users-from-external-directories" was selected to be "specific".specific blocks are documented below.


`portal_web_settings` supports the following:

* `aliases` - (Optional) List of URL aliases that are redirected to the main portal URL.aliases blocks are documented below.
* `main_url` - (Optional) The main URL for the web portal. 


`certificate_settings` supports the following:

* `base64_certificate` - (Optional) The certificate file encoded in Base64 with padding. 
This file must be in the *.p12 format. 
* `base64_password` - (Optional) Password (encoded in Base64 with padding) for the certificate file. 


`accessibility` supports the following:

* `allow_access_from` - (Optional) Allowed access to the web portal (based on interfaces, or security policy). 
* `internal_access_settings` - (Optional) Configuration of the additional portal access settings for internal interfaces only.internal_access_settings blocks are documented below.


`users_directories` supports the following:

* `external_user_profile` - (Optional) External user profile. 
* `internal_users` - (Optional) Internal users. 
* `users_from_external_directories` - (Optional) Users from external directories. 
* `specific` - (Optional) LDAP AU objects identified by the name or UID. Must be set when "users-from-external-directories" was selected to be "specific".specific blocks are documented below.


`accessibility` supports the following:

* `allow_access_from` - (Optional) Allowed access to the web portal (based on interfaces, or security policy). 
* `internal_access_settings` - (Optional) Configuration of the additional portal access settings for internal interfaces only.internal_access_settings blocks are documented below.


`users_directories` supports the following:

* `external_user_profile` - (Optional) External user profile. 
* `internal_users` - (Optional) Internal users. 
* `users_from_external_directories` - (Optional) Users from external directories. 
* `specific` - (Optional) LDAP AU objects identified by the name or UID. Must be set when "users-from-external-directories" was selected to be "specific".specific blocks are documented below.


`accessibility` supports the following:

* `allow_access_from` - (Optional) Allowed access to the web portal (based on interfaces, or security policy). 
* `internal_access_settings` - (Optional) Configuration of the additional portal access settings for internal interfaces only.internal_access_settings blocks are documented below.


`optional_parameters` supports the following:

* `use_primary_dns_server` - (Optional) Use Primary DNS Server. 
* `primary_dns_server` - (Optional) Primary DNS Server. Identified by name or UID.
Must be set when "use-primary-dns-server" is true and can not be set when "use-primary-dns-server" is false. 
* `use_first_backup_dns_server` - (Optional) Use First Backup DNS Server. 
* `first_backup_dns_server` - (Optional) First Backup DNS Server. Identified by name or UID.
Must be set when "use-first-backup-dns-server" is true and can not be set when "use-first-backup-dns-server" is false. 
* `use_second_backup_dns_server` - (Optional) Use Second Backup DNS Server. 
* `second_backup_dns_server` - (Optional) Second Backup DNS Server. Identified by name or UID.
Must be set when "use-second-backup-dns-server" is true and can not be set when "use-second-backup-dns-server" is false. 
* `dns_suffixes` - (Optional) DNS Suffixes. 
* `use_primary_wins_server` - (Optional) Use Primary WINS Server. 
* `primary_wins_server` - (Optional) Primary WINS Server. Identified by name or UID.
Must be set when "use-primary-wins-server" is true and can not be set when "use-primary-wins-server" is false. 
* `use_first_backup_wins_server` - (Optional) Use First Backup WINS Server. 
* `first_backup_wins_server` - (Optional) First Backup WINS Server. Identified by name or UID.
Must be set when "use-first-backup-wins-server" is true and can not be set when "use-first-backup-wins-server" is false. 
* `use_second_backup_wins_server` - (Optional) Use Second Backup WINS Server. 
* `second_backup_wins_server` - (Optional) Second Backup WINS Server. Identified by name or UID.
Must be set when "use-second-backup-wins-server" is true and can not be set when "use-second-backup-wins-server" is false. 


`internal_access_settings` supports the following:

* `undefined` - (Optional) Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'. 
* `dmz` - (Optional) Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'. 
* `vpn` - (Optional) Controls portal access settings for interfaces that are part of a VPN Encryption Domain. 


`internal_access_settings` supports the following:

* `undefined` - (Optional) Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'. 
* `dmz` - (Optional) Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'. 
* `vpn` - (Optional) Controls portal access settings for interfaces that are part of a VPN Encryption Domain. 


`internal_access_settings` supports the following:

* `undefined` - (Optional) Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'. 
* `dmz` - (Optional) Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'. 
* `vpn` - (Optional) Controls portal access settings for interfaces that are part of a VPN Encryption Domain. 
