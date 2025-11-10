---
layout: "checkpoint"
page_title: "checkpoint_management_simple_cluster"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-simple-cluster"
description: |-
This resource allows you to execute Check Point Simple Cluster.
---

# Data Source: checkpoint_management_simple_cluster

This resource allows you to execute Check Point Simple Cluster.

## Example Usage

```hcl
resource "checkpoint_management_simple_cluster" "simple_cluster" {
    name = "mycluster"
    ipv4_address = "1.2.3.4"
    version = "R81"
    hardware = "Open server"
    send_logs_to_server = ["mylogserver"]
    firewall = true
}

data "checkpoint_management_simple_cluster" "simple_cluster" {
    name = "${checkpoint_management_simple_cluster.test.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.
* `advanced_settings` - N/Aadvanced_settings blocks are documented below.
* `anti_bot` - Anti-Bot blade enabled.
* `anti_virus` - Anti-Virus blade enabled.
* `application_control` - Application Control blade enabled.
* `application_control_and_url_filtering_settings` - Gateway Application Control and URL filtering settings.application_control_and_url_filtering_settings blocks are documented below.
* `cluster_mode` - Cluster mode.
* `cluster_settings` - ClusterXL and VRRP Settings.cluster_settings blocks are documented below.
* `content_awareness` - Content Awareness blade enabled.
* `enable_https_inspection` - Enable HTTPS Inspection after defining an outbound inspection certificate. <br>To define the outbound certificate use outbound inspection certificate API.
* `fetch_policy` - Security management server(s) to fetch the policy from.fetch_policy blocks are documented below.
* `firewall` - Firewall blade enabled.
* `firewall_settings` - N/Afirewall_settings blocks are documented below.
* `geo_mode` - Cluster High Availability Geo mode.<br>This setting applies only to a cluster deployed in a cloud. Available when the cluster mode equals "cluster-xl-ha".
* `hardware` - Cluster platform hardware.
* `hit_count` - Hit count tracks the number of connections each rule matches.
* `https_inspection` - HTTPS inspection.https_inspection blocks are documented below.
* `identity_awareness` - Identity awareness blade enabled.
* `identity_awareness_settings` - Gateway Identity Awareness settings.identity_awareness_settings blocks are documented below.
* `interfaces` - Cluster interfaces.interfaces blocks are documented below.
* `ipv4_address` - IPv4 address.
* `ipv6_address` - IPv6 address.
* `ips` - Intrusion Prevention System blade enabled.
* `ips_settings` - Cluster IPS settings. ips_settings blocks are documented below.
* `ips_update_policy` - Specifies whether the IPS will be downloaded from the Management or directly to the Gateway.
* `members` - Cluster members list. Only new cluster member can be added. Adding existing gateway is not supported.members blocks are documented below.
* `nat_hide_internal_interfaces` - Hide internal networks behind the Gateway's external IP.
* `nat_settings` - NAT settings.nat_settings blocks are documented below.
* `os_name` - Cluster platform operating system.
* `platform_portal_settings` - Platform portal settings.platform_portal_settings blocks are documented below.
* `proxy_settings` - Proxy Server for Gateway.proxy_settings blocks are documented below.
* `qos` - QoS.
* `send_alerts_to_server` - Server(s) to send alerts to.send_alerts_to_server blocks are documented below.
* `send_logs_to_backup_server` - Backup server(s) to send logs to.send_logs_to_backup_server blocks are documented below.
* `send_logs_to_server` - Server(s) to send logs to.send_logs_to_server blocks are documented below.
* `tags` - Collection of tag identifiers.tags blocks are documented below.
* `threat_emulation` - Threat Emulation blade enabled.
* `threat_extraction` - Threat Extraction blade enabled.
* `threat_prevention_mode` - The mode of Threat Prevention to use. When using Autonomous Threat Prevention, disabling the Threat Prevention blades is not allowed.
* `url_filtering` - URL Filtering blade enabled.
* `usercheck_portal_settings` - UserCheck portal settings.usercheck_portal_settings blocks are documented below.
* `version` - Cluster platform version.
* `vpn` - VPN blade enabled.
* `vpn_settings` - Gateway VPN settings.vpn_settings blocks are documented below.
* `zero_phishing` - Zero Phishing blade enabled.
* `zero_phishing_fqdn` - Zero Phishing gateway FQDN.
* `show_portals_certificate` - Indicates whether to show the portals certificate value in the reply.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.
* `groups` - Collection of group identifiers.groups blocks are documented below.
* `ignore_warnings` - Apply changes ignoring warnings.
* `ignore_errors` - Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.


`advanced_settings` supports the following:

* `connection_persistence` - Handling established connections when installing a new policy.
* `sam` - SAM.sam blocks are documented below.


`application_control_and_url_filtering_settings` supports the following:

* `global_settings_mode` - Whether to override global settings or not.
* `override_global_settings` - override global settings object.override_global_settings blocks are documented below.


`cluster_settings` supports the following:

* `member_recovery_mode` - In a High Availability cluster, each member is given a priority. The member with the highest priority serves as the gateway. If this gateway fails, control is passed to the member with the next highest priority. If that member fails, control is passed to the next, and so on. Upon gateway recovery, it is possible to:
  Maintain current active Cluster Member (maintain-current-active) or
  Switch to higher priority Cluster Member (according-to-priority).
* `state_synchronization` - Cluster State Synchronization settings.state_synchronization blocks are documented below.
* `track_changes_of_cluster_members` - Track changes in the status of Cluster Members.
* `use_virtual_mac` - Use Virtual MAC. By enabling Virtual MAC in ClusterXL High Availability New mode, or Load Sharing Unicast mode, all cluster members associate the same Virtual MAC address with All Cluster Virtual Interfaces and the Virtual IP address.


`firewall_settings` supports the following:

* `auto_calculate_connections_hash_table_size_and_memory_pool` - N/A
* `auto_maximum_limit_for_concurrent_connections` - N/A
* `connections_hash_size` - N/A
* `maximum_limit_for_concurrent_connections` - N/A
* `maximum_memory_pool_size` - N/A
* `memory_pool_size` - N/A


`https_inspection` supports the following:

* `bypass_on_failure` - Set to be true in order to bypass all requests (Fail-open) in case of internal system error.bypass_on_failure blocks are documented below.
* `site_categorization_allow_mode` - Set to 'background' in order to allowed requests until categorization is complete.site_categorization_allow_mode blocks are documented below.
* `deny_untrusted_server_cert` - Set to be true in order to drop traffic from servers with untrusted server certificate.deny_untrusted_server_cert blocks are documented below.
* `deny_revoked_server_cert` - Set to be true in order to drop traffic from servers with revoked server certificate (validate CRL).deny_revoked_server_cert blocks are documented below.
* `deny_expired_server_cert` - Set to be true in order to drop traffic from servers with expired server certificate.deny_expired_server_cert blocks are documented below.


`identity_awareness_settings` supports the following:

* `browser_based_authentication` - Enable Browser Based Authentication source.
* `browser_based_authentication_settings` - Browser Based Authentication settings.browser_based_authentication_settings blocks are documented below.
* `identity_agent` - Enable Identity Agent source.
* `identity_agent_settings` - Identity Agent settings.identity_agent_settings blocks are documented below.
* `identity_collector` - Enable Identity Collector source.
* `identity_collector_settings` - Identity Collector settings.identity_collector_settings blocks are documented below.
* `identity_sharing_settings` - Identity sharing settings.identity_sharing_settings blocks are documented below.
* `proxy_settings` - Identity-Awareness Proxy settings.proxy_settings blocks are documented below.
* `remote_access` - Enable Remote Access Identity source.


`interfaces` supports the following:

* `name` - Object name. Must be unique in the domain.
* `interface_type` - Cluster interface type.
* `ipv4_address` - IPv4 address.
* `ipv6_address` - IPv6 address.
* `network_mask` - IPv4 or IPv6 network mask. If both masks are required use ipv4-network-mask and ipv6-network-mask fields explicitly. Instead of providing mask itself it is possible to specify IPv4 or IPv6 mask length in mask-length field. If both masks length are required use ipv4-mask-length and  ipv6-mask-length fields explicitly.
* `ipv4_network_mask` - IPv4 network address.
* `ipv6_network_mask` - IPv6 network address.
* `ipv4_mask_length` - IPv4 network mask length.
* `ipv6_mask_length` - IPv6 network mask length.
* `anti_spoofing` - N/A
* `anti_spoofing_settings` - N/Aanti_spoofing_settings blocks are documented below.
* `multicast_address` - Multicast IP Address.
* `multicast_address_type` - Multicast Address Type.
* `security_zone` - N/A
* `security_zone_settings` - N/Asecurity_zone_settings blocks are documented below.
* `tags` - Collection of tag identifiers.tags blocks are documented below.
* `topology` - N/A
* `topology_settings` - N/Atopology_settings blocks are documented below.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.
* `ignore_warnings` - Apply changes ignoring warnings.
* `ignore_errors` - Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.

`ips_settings` supports the following:

* `bypass_all_under_load` - Disable/enable all IPS protections until CPU and memory levels are back to normal.
* `bypass_track_method` - Track options when all IPS protections are disabled until CPU/memory levels are back to normal.
* `top_cpu_consuming_protections` - Provides a way to reduce CPU levels on machines under load by disabling the top CPU consuming IPS protections.top_cpu_consuming_protections blocks are documented below.
* `activation_mode` - Defines whether the IPS blade operates in Detect Only mode or enforces the configured IPS Policy.
* `cpu_usage_low_threshold` - CPU usage low threshold percentage (1-99).
* `cpu_usage_high_threshold` - CPU usage high threshold percentage (1-99).
* `memory_usage_low_threshold` - Memory usage low threshold percentage (1-99).
* `memory_usage_high_threshold` - Memory usage high threshold percentage (1-99).
* `send_threat_cloud_info` - Help improve Check Point Threat Prevention product by sending anonymous information.
* `reject_on_cluster_fail_over` - Define the IPS connections during fail over reject packets or accept packets.

`members` supports the following:

* `name` - Object name.
* `interfaces` - Cluster Member network interfaces.interfaces blocks are documented below.
* `ipv4_address` - IPv4 address.
* `ipv6_address` - IPv6 address.
* `one_time_password` - N/A
* `tags` - Collection of tag identifiers.tags blocks are documented below.
* `priority` - In a High Availability New mode cluster each machine is given a priority. The highest priority machine serves as the gateway in normal circumstances. If this machine fails, control is passed to the next highest priority machine. If that machine fails, control is passed to the next machine, and so on.
  In Load Sharing Unicast mode cluster, the highest priority is the pivot machine.
  The values must be in a range from 1 to N, where N is number of cluster members.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.
* `ignore_warnings` - Apply changes ignoring warnings.
* `ignore_errors` - Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.


`nat_settings` supports the following:

* `auto_rule` - Whether to add automatic address translation rules.
* `ipv4_address` - IPv4 address.
* `ipv6_address` - IPv6 address.
* `hide_behind` - Hide behind method. This parameter is forbidden in case "method" parameter is "static".
* `install_on` - Which gateway should apply the NAT translation.
* `method` - NAT translation method.


`platform_portal_settings` supports the following:

* `portal_web_settings` - Configuration of the portal web settings.portal_web_settings blocks are documented below.
* `certificate_settings` - Configuration of the portal certificate settings.certificate_settings blocks are documented below.
* `accessibility` - Configuration of the portal access settings.accessibility blocks are documented below.


`proxy_settings` supports the following:

* `use_custom_proxy` - Use custom proxy settings for this network object.
* `proxy_server` - N/A
* `port` - N/A


`usercheck_portal_settings` supports the following:

* `enabled` - State of the web portal (enabled or disabled). The supported blades are: {'Application Control', 'URL Filtering', 'Data Loss Prevention', 'Anti Virus', 'Anti Bot', 'Threat Emulation', 'Threat Extraction', 'Data Awareness'}.
* `portal_web_settings` - Configuration of the portal web settings.portal_web_settings blocks are documented below.
* `certificate_settings` - Configuration of the portal certificate settings.certificate_settings blocks are documented below.
* `accessibility` - Configuration of the portal access settings.accessibility blocks are documented below.


`vpn_settings` supports the following:

* `authentication` - Authentication.authentication blocks are documented below.
* `link_selection` - Link Selection.link_selection blocks are documented below.
* `maximum_concurrent_ike_negotiations` - N/A
* `maximum_concurrent_tunnels` - N/A
* `office_mode` - Office Mode.
  Notation Wide Impact - Office Mode apply IPSec VPN Software Blade clients and to the Mobile Access Software Blade clients.office_mode blocks are documented below.
* `remote_access` - Remote Access.remote_access blocks are documented below.
* `vpn_domain` - Gateway VPN domain identified by the name or UID.
* `vpn_domain_exclude_external_ip_addresses` - Exclude the external IP addresses from the VPN domain of this Security Gateway.
* `vpn_domain_type` - Gateway VPN domain type.


`sam` supports the following:

* `forward_to_other_sam_servers` - Forward SAM clients' requests to other SAM servers.
* `use_early_versions` - Use early versions compatibility mode.use_early_versions blocks are documented below.
* `purge_sam_file` - Purge SAM File.purge_sam_file blocks are documented below.


`override_global_settings` supports the following:

* `fail_mode` - Fail mode - allow or block all requests.
* `website_categorization` - Website categorization object.website_categorization blocks are documented below.


`state_synchronization` supports the following:

* `delayed` - Start synchronizing with delay of seconds, as defined by delayed-seconds, after connection initiation. Disabled when state-synchronization disabled.
* `delayed_seconds` - Start synchronizing X seconds after connection initiation
  . The values must be in a range between 2 and 3600.
* `enabled` - Use State Synchronization.


`bypass_on_failure` supports the following:

* `override_profile` - Override profile of global configuration.
* `value` - Override value.<br><font color="red">Required only for</font> 'override-profile' is True.


`site_categorization_allow_mode` supports the following:

* `override_profile` - Override profile of global configuration.
* `value` - Override value.<br><font color="red">Required only for</font> 'override-profile' is True.


`deny_untrusted_server_cert` supports the following:

* `override_profile` - Override profile of global configuration.
* `value` - Override value.<br><font color="red">Required only for</font> 'override-profile' is True.


`deny_revoked_server_cert` supports the following:

* `override_profile` - Override profile of global configuration.
* `value` - Override value.<br><font color="red">Required only for</font> 'override-profile' is True.


`deny_expired_server_cert` supports the following:

* `override_profile` - Override profile of global configuration.
* `value` - Override value.<br><font color="red">Required only for</font> 'override-profile' is True.


`browser_based_authentication_settings` supports the following:

* `authentication_settings` - Authentication Settings for Browser Based Authentication.authentication_settings blocks are documented below.
* `browser_based_authentication_portal_settings` - Browser Based Authentication portal settings.browser_based_authentication_portal_settings blocks are documented below.


`identity_agent_settings` supports the following:

* `agents_interval_keepalive` - Agents send keepalive period (minutes).
* `user_reauthenticate_interval` - Agent reauthenticate time interval (minutes).
* `authentication_settings` - Authentication Settings for Identity Agent.authentication_settings blocks are documented below.
* `identity_agent_portal_settings` - Identity Agent accessibility settings.identity_agent_portal_settings blocks are documented below.


`identity_collector_settings` supports the following:

* `authorized_clients` - Authorized Clients.authorized_clients blocks are documented below.
* `authentication_settings` - Authentication Settings for Identity Collector.authentication_settings blocks are documented below.
* `client_access_permissions` - Identity Collector accessibility settings.client_access_permissions blocks are documented below.


`identity_sharing_settings` supports the following:

* `share_with_other_gateways` - Enable identity sharing with other gateways.
* `receive_from_other_gateways` - Enable receiving identity from other gateways.
* `receive_from` - Gateway(s) to receive identity from.receive_from blocks are documented below.


`proxy_settings` supports the following:

* `detect_using_x_forward_for` - Whether to use X-Forward-For HTTP header, which is added by the proxy server to keep track of the original source IP.


`anti_spoofing_settings` supports the following:

* `action` - If packets will be rejected (the Prevent option) or whether the packets will be monitored (the Detect option).
* `exclude_packets` - Don't check packets from excluded network.
* `excluded_network_name` - Excluded network name.
* `excluded_network_uid` - Excluded network UID.
* `spoof_tracking` - Spoof tracking.


`security_zone_settings` supports the following:

* `auto_calculated` - Security Zone is calculated according to where the interface leads to.
* `specific_zone` - Security Zone specified manually.


`topology_settings` supports the following:

* `interface_leads_to_dmz` - Whether this interface leads to demilitarized zone (perimeter network).
* `specific_network` - Network behind this interface.


`interfaces` supports the following:

* `name` - Object name.
* `anti_spoofing` - N/A
* `anti_spoofing_settings` - N/Aanti_spoofing_settings blocks are documented below.
* `ipv4_address` - IPv4 address.
* `ipv6_address` - IPv6 address.
* `network_mask` - IPv4 or IPv6 network mask. If both masks are required use ipv4-network-mask and ipv6-network-mask fields explicitly. Instead of providing mask itself it is possible to specify IPv4 or IPv6 mask length in mask-length field. If both masks length are required use ipv4-mask-length and  ipv6-mask-length fields explicitly.
* `ipv4_network_mask` - IPv4 network address.
* `ipv6_network_mask` - IPv6 network address.
* `ipv4_mask_length` - IPv4 network mask length.
* `ipv6_mask_length` - IPv6 network mask length.
* `security_zone` - N/A
* `security_zone_settings` - N/Asecurity_zone_settings blocks are documented below.
* `tags` - Collection of tag identifiers.tags blocks are documented below.
* `topology` - N/A
* `topology_settings` - N/Atopology_settings blocks are documented below.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.
* `ignore_warnings` - Apply changes ignoring warnings.
* `ignore_errors` - Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.


`portal_web_settings` supports the following:

* `aliases` - List of URL aliases that are redirected to the main portal URL.aliases blocks are documented below.
* `main_url` - The main URL for the web portal.


`certificate_settings` supports the following:

* `base64_certificate` - The certificate file encoded in Base64 with padding.
  This file must be in the *.p12 format.
* `base64_password` - Password (encoded in Base64 with padding) for the certificate file.


`accessibility` supports the following:

* `allow_access_from` - Allowed access to the web portal (based on interfaces, or security policy).
* `internal_access_settings` - Configuration of the additional portal access settings for internal interfaces only.internal_access_settings blocks are documented below.


`portal_web_settings` supports the following:

* `aliases` - List of URL aliases that are redirected to the main portal URL.aliases blocks are documented below.
* `main_url` - The main URL for the web portal.


`certificate_settings` supports the following:

* `base64_certificate` - The certificate file encoded in Base64 with padding.
  This file must be in the *.p12 format.
* `base64_password` - Password (encoded in Base64 with padding) for the certificate file.


`accessibility` supports the following:

* `allow_access_from` - Allowed access to the web portal (based on interfaces, or security policy).
* `internal_access_settings` - Configuration of the additional portal access settings for internal interfaces only.internal_access_settings blocks are documented below.


`authentication` supports the following:

* `authentication_clients` - Collection of VPN Authentication clients identified by the name or UID.authentication_clients blocks are documented below.


`link_selection` supports the following:

* `dns_resolving_hostname` - DNS Resolving Hostname. Must be set when "ip-selection" was selected to be "dns-resolving-from-hostname".


`office_mode` supports the following:

* `mode` - Office Mode Permissions.
  When selected to be "off", all the other definitions are irrelevant.
* `group` - Group. Identified by name or UID.
  Must be set when "office-mode-permissions" was selected to be "group".
* `allocate_ip_address_from` - Allocate IP address Method.
  Allocate IP address by sequentially trying the given methods until success.allocate_ip_address_from blocks are documented below.
* `support_multiple_interfaces` - Support connectivity enhancement for gateways with multiple external interfaces.
* `perform_anti_spoofing` - Perform Anti-Spoofing on Office Mode addresses.
* `anti_spoofing_additional_addresses` - Additional IP Addresses for Anti-Spoofing.
  Identified by name or UID.
  Must be set when "perform-anti-spoofings" is true.


`remote_access` supports the following:

* `support_l2tp` - Support L2TP (relevant only when office mode is active).
* `l2tp_auth_method` - L2TP Authentication Method.
  Must be set when "support-l2tp" is true.
* `l2tp_certificate` - L2TP Certificate.
  Must be set when "l2tp-auth-method" was selected to be "certificate".
  Insert "defaultCert" when you want to use the default certificate.
* `allow_vpn_clients_to_route_traffic` - Allow VPN clients to route traffic.
* `support_nat_traversal_mechanism` - Support NAT traversal mechanism (UDP encapsulation).
* `nat_traversal_service` - Allocated NAT traversal UDP service. Identified by name or UID.
  Must be set when "support-nat-traversal-mechanism" is true.
* `support_visitor_mode` - Support Visitor Mode.
* `visitor_mode_service` - TCP Service for Visitor Mode. Identified by name or UID.
  Must be set when "support-visitor-mode" is true.
* `visitor_mode_interface` - Interface for Visitor Mode.
  Must be set when "support-visitor-mode" is true.
  Insert IPV4 Address of existing interface or "All IPs" when you want all interfaces.


`use_early_versions` supports the following:

* `enabled` - Use early versions compatibility mode.
* `compatibility_mode` - Early versions compatibility mode.


`purge_sam_file` supports the following:

* `enabled` - Purge SAM File.
* `purge_when_size_reaches_to` - Purge SAM File When it Reaches to.


`website_categorization` supports the following:

* `mode` - Website categorization mode.
* `custom_mode` - Custom mode object.custom_mode blocks are documented below.


`authentication_settings` supports the following:

* `authentication_method` - Authentication method.
* `identity_provider` - Identity provider object identified by the name or UID. Must be set when "authentication-method" was selected to be "identity provider".identity_provider blocks are documented below.
* `radius` - Radius server object identified by the name or UID. Must be set when "authentication-method" was selected to be "radius".
* `users_directories` - Users directories.users_directories blocks are documented below.


`browser_based_authentication_portal_settings` supports the following:

* `portal_web_settings` - Configuration of the portal web settings.portal_web_settings blocks are documented below.
* `certificate_settings` - Configuration of the portal certificate settings.certificate_settings blocks are documented below.
* `accessibility` - Configuration of the portal access settings.accessibility blocks are documented below.


`authentication_settings` supports the following:

* `authentication_method` - Authentication method.
* `radius` - Radius server object identified by the name or UID. Must be set when "authentication-method" was selected to be "radius".
* `users_directories` - Users directories.users_directories blocks are documented below.


`identity_agent_portal_settings` supports the following:

* `accessibility` - Configuration of the portal access settings.accessibility blocks are documented below.


`authorized_clients` supports the following:

* `client` - Host / Network Group Name or UID.
* `client_secret` - Client Secret.


`authentication_settings` supports the following:

* `users_directories` - Users directories.users_directories blocks are documented below.


`client_access_permissions` supports the following:

* `accessibility` - Configuration of the portal access settings.accessibility blocks are documented below.


`anti_spoofing_settings` supports the following:

* `action` - If packets will be rejected (the Prevent option) or whether the packets will be monitored (the Detect option).
* `exclude_packets` - Don't check packets from excluded network.
* `excluded_network_name` - Excluded network name.
* `excluded_network_uid` - Excluded network UID.
* `spoof_tracking` - Spoof tracking.


`security_zone_settings` supports the following:

* `auto_calculated` - Security Zone is calculated according to where the interface leads to.
* `specific_zone` - Security Zone specified manually.


`topology_settings` supports the following:

* `interface_leads_to_dmz` - Whether this interface leads to demilitarized zone (perimeter network).
* `specific_network` - Network behind this interface.


`internal_access_settings` supports the following:

* `undefined` - Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'.
* `dmz` - Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'.
* `vpn` - Controls portal access settings for interfaces that are part of a VPN Encryption Domain.


`internal_access_settings` supports the following:

* `undefined` - Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'.
* `dmz` - Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'.
* `vpn` - Controls portal access settings for interfaces that are part of a VPN Encryption Domain.


`allocate_ip_address_from` supports the following:

* `radius_server` - Radius server used to authenticate the user.
* `use_allocate_method` - Use Allocate Method.
* `allocate_method` - Using either Manual (IP Pool) or Automatic (DHCP).
  Must be set when "use-allocate-method" is true.
* `manual_network` - Manual Network. Identified by name or UID.
  Must be set when "allocate-method" was selected to be "manual".
* `dhcp_server` - DHCP Server. Identified by name or UID.
  Must be set when "allocate-method" was selected to be "automatic".
* `virtual_ip_address` - Virtual IPV4 address for DHCP server replies.
  Must be set when "allocate-method" was selected to be "automatic".
* `dhcp_mac_address` - Calculated MAC address for DHCP allocation.
  Must be set when "allocate-method" was selected to be "automatic".
* `optional_parameters` - This configuration applies to all Office Mode methods except Automatic (using DHCP) and ipassignment.conf entries which contain this data.optional_parameters blocks are documented below.


`custom_mode` supports the following:

* `social_networking_widgets` - Social networking widgets mode.
* `url_filtering` - URL filtering mode.


`users_directories` supports the following:

* `external_user_profile` - External user profile.
* `internal_users` - Internal users.
* `users_from_external_directories` - Users from external directories.
* `specific` - LDAP AU objects identified by the name or UID. Must be set when "users-from-external-directories" was selected to be "specific".specific blocks are documented below.


`portal_web_settings` supports the following:

* `aliases` - List of URL aliases that are redirected to the main portal URL.aliases blocks are documented below.
* `main_url` - The main URL for the web portal.


`certificate_settings` supports the following:

* `base64_certificate` - The certificate file encoded in Base64 with padding.
  This file must be in the *.p12 format.
* `base64_password` - Password (encoded in Base64 with padding) for the certificate file.


`accessibility` supports the following:

* `allow_access_from` - Allowed access to the web portal (based on interfaces, or security policy).
* `internal_access_settings` - Configuration of the additional portal access settings for internal interfaces only.internal_access_settings blocks are documented below.


`users_directories` supports the following:

* `external_user_profile` - External user profile.
* `internal_users` - Internal users.
* `users_from_external_directories` - Users from external directories.
* `specific` - LDAP AU objects identified by the name or UID. Must be set when "users-from-external-directories" was selected to be "specific".specific blocks are documented below.


`accessibility` supports the following:

* `allow_access_from` - Allowed access to the web portal (based on interfaces, or security policy).
* `internal_access_settings` - Configuration of the additional portal access settings for internal interfaces only.internal_access_settings blocks are documented below.


`users_directories` supports the following:

* `external_user_profile` - External user profile.
* `internal_users` - Internal users.
* `users_from_external_directories` - Users from external directories.
* `specific` - LDAP AU objects identified by the name or UID. Must be set when "users-from-external-directories" was selected to be "specific".specific blocks are documented below.


`accessibility` supports the following:

* `allow_access_from` - Allowed access to the web portal (based on interfaces, or security policy).
* `internal_access_settings` - Configuration of the additional portal access settings for internal interfaces only.internal_access_settings blocks are documented below.


`optional_parameters` supports the following:

* `use_primary_dns_server` - Use Primary DNS Server.
* `primary_dns_server` - Primary DNS Server. Identified by name or UID.
  Must be set when "use-primary-dns-server" is true and can not be set when "use-primary-dns-server" is false.
* `use_first_backup_dns_server` - Use First Backup DNS Server.
* `first_backup_dns_server` - First Backup DNS Server. Identified by name or UID.
  Must be set when "use-first-backup-dns-server" is true and can not be set when "use-first-backup-dns-server" is false.
* `use_second_backup_dns_server` - Use Second Backup DNS Server.
* `second_backup_dns_server` - Second Backup DNS Server. Identified by name or UID.
  Must be set when "use-second-backup-dns-server" is true and can not be set when "use-second-backup-dns-server" is false.
* `dns_suffixes` - DNS Suffixes.
* `use_primary_wins_server` - Use Primary WINS Server.
* `primary_wins_server` - Primary WINS Server. Identified by name or UID.
  Must be set when "use-primary-wins-server" is true and can not be set when "use-primary-wins-server" is false.
* `use_first_backup_wins_server` - Use First Backup WINS Server.
* `first_backup_wins_server` - First Backup WINS Server. Identified by name or UID.
  Must be set when "use-first-backup-wins-server" is true and can not be set when "use-first-backup-wins-server" is false.
* `use_second_backup_wins_server` - Use Second Backup WINS Server.
* `second_backup_wins_server` - Second Backup WINS Server. Identified by name or UID.
  Must be set when "use-second-backup-wins-server" is true and can not be set when "use-second-backup-wins-server" is false.


`internal_access_settings` supports the following:

* `undefined` - Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'.
* `dmz` - Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'.
* `vpn` - Controls portal access settings for interfaces that are part of a VPN Encryption Domain.


`internal_access_settings` supports the following:

* `undefined` - Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'.
* `dmz` - Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'.
* `vpn` - Controls portal access settings for interfaces that are part of a VPN Encryption Domain.


`internal_access_settings` supports the following:

* `undefined` - Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'.
* `dmz` - Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'.
* `vpn` - Controls portal access settings for interfaces that are part of a VPN Encryption Domain. 
