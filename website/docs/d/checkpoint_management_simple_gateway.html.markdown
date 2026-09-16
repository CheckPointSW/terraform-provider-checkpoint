---
layout: "checkpoint"
page_title: "checkpoint_management_simple_gateway"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-simple-gateway"
description: |-
This resource allows you to execute Check Point Simple Gateway.
---

# Data Source: checkpoint_management_simple_gateway

This resource allows you to execute Check Point Simple Gateway.

## Example Usage


```hcl
resource "checkpoint_management_simple_gateway" "simple_gateway" {
    name = "mygateway"
    ipv4_address = "1.2.3.4"
    version = "R81"
    send_logs_to_server = ["logserver"]
}

data "checkpoint_management_simple_gateway" "simple_gateway" {
    name = "${checkpoint_management_simple_gateway.test.name}"
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
* `content_awareness` - Content Awareness blade enabled.
* `enable_https_inspection` - Enable HTTPS Inspection after defining an outbound inspection certificate. <br>To define the outbound certificate use outbound inspection certificate API.
* `fetch_policy` - Security management server(s) to fetch the policy from.fetch_policy blocks are documented below.
* `firewall` - Firewall blade enabled.
* `firewall_settings` - N/Afirewall_settings blocks are documented below.
* `hit_count` - Hit count tracks the number of connections each rule matches.
* `https_inspection` - HTTPS inspection.https_inspection blocks are documented below.
* `icap_server` - ICAP Server enabled.
* `identity_awareness` - Identity awareness blade enabled.
* `identity_awareness_settings` - Gateway Identity Awareness settings.identity_awareness_settings blocks are documented below.
* `interfaces` - Network interfaces.interfaces blocks are documented below.
* `ipv4_address` - IPv4 address.
* `ipv6_address` - IPv6 address.
* `ips` - Intrusion Prevention System blade enabled.
* `ips_settings` - Gateway IPS settings.ips_settings blocks are documented below.
* `ips_update_policy` - Specifies whether the IPS will be downloaded from the Management or directly to the Gateway.
* `nat_hide_internal_interfaces` - Hide internal networks behind the Gateway's external IP.
* `nat_settings` - NAT settings.nat_settings blocks are documented below.
* `one_time_password` - N/A
* `os_name` - Gateway platform operating system.
* `platform_portal_settings` - Platform portal settings.platform_portal_settings blocks are documented below.
* `proxy_settings` - Proxy Server for Gateway.proxy_settings blocks are documented below.
* `qos` - QoS.
* `save_logs_locally` - Save logs locally on the gateway.
* `send_alerts_to_server` - Server(s) to send alerts to.send_alerts_to_server blocks are documented below.
* `send_logs_to_backup_server` - Backup server(s) to send logs to.send_logs_to_backup_server blocks are documented below.
* `send_logs_to_server` - Server(s) to send logs to.send_logs_to_server blocks are documented below.
* `tags` - Collection of tag identifiers.tags blocks are documented below.
* `threat_emulation` - Threat Emulation blade enabled.
* `threat_extraction` - Threat Extraction blade enabled.
* `threat_prevention_mode` - The mode of Threat Prevention to use. When using Autonomous Threat Prevention, disabling the Threat Prevention blades is not allowed.
* `url_filtering` - URL Filtering blade enabled.
* `usercheck_portal_settings` - UserCheck portal settings.usercheck_portal_settings blocks are documented below.
* `version` - Gateway platform version.
* `vpn` - VPN blade enabled.
* `vpn_settings` - Gateway VPN settings.vpn_settings blocks are documented below.
* `zero_phishing` - Zero Phishing blade enabled.
* `zero_phishing_fqdn` - Zero Phishing gateway FQDN. **Deprecated** - use `zero_phishing_settings.manual_fqdn` instead.
* `logs_settings` - Logs settings that apply to Quantum Security Gateways that run Gaia OS.logs_settings blocks are documented below.
* `show_portals_certificate` - Indicates whether to show the portals certificate value in the reply.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.
* `groups` - Collection of group identifiers.groups blocks are documented below.
* `ignore_errors` - Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.
* `accept_syslog_messages` - Enables the gateway accept syslog messages, relevant only when Logging and Status blade is enabled.
* `anti_spam_and_email_security` - Anti-Spam & Email-Security blade enabled.
* `auto_generate_ip` - Use an automatically generated IP address for the Gateway object (applies only to Smart-1 Cloud).
* `auto_topology_custom_recalculation_time` - Auto topology custom recalculation time (seconds).
* `auto_topology_use_custom_recalculation_time` - Auto topology to use custom recalculation time instead of default.
* `autonomous_system_number` - The Autonomous System Number (ASN) for this Check Point Security Gateway.
* `communication_with_servers_behind_nat` - Gateway behind NAT communications settings with the server.communication_with_servers_behind_nat blocks are documented below.
* `data_loss_prevention` - Data Loss Prevention.
* `dns_server` - DNS Server.
* `enable_log_indexing` - Enable log indexing, The Log Indexing uses more storage to provide fast log queries. Relevant only when Logging and Status blade is enabled.
* `export_logs_to_servers` - Export logs to syslog/SIEM servers. NOTE:After you configure a Log Exporter, you must run Install Database.
* `externally_managed` - Externally Managed Check Point Gateway.
* `fetch_policy_scheduler` - Fetch policy functionality settings.fetch_policy_scheduler blocks are documented below.
* `hardware_subtype` - Gateway type (relevant only for Spark gateways).
* `install_policy_without_push` - Specifies whether the policy is pushed to the gateway during policy installation, or whether the gateway should fetch the policy.
* `interfaces_topology_settings` - Topology setting for all interfaces on a Security Gateway. Default for Security Gateways that run Gaia OS: 'per interface'. Default for Quantum Spark ...
* `legacy_url_filtering` - Legacy URL Filtering enabled.
* `log_server` - Logging & Status.
* `mobile_access` - Mobile-Access blade enabled.
* `monitoring` - Monitoring blade enabled.
* `network_policy_management` - Management blade enabled.
* `policy_server` - Policy-Server blade enabled.
* `rtm_counters_report` - Enables monitoring blades system counters report (e.g CPU Usage,Memory Usage).
* `rtm_traffic_report` - Enables monitoring blades traffic report.
* `rtm_traffic_report_per_connection` - Enables Monitoring blade traffic report per connection.
* `sic_message` - Secure Internal Communication message.
* `smart_event_intro_correlation_unit` - Enables the gateway use SmartEvent intro correlation unit with one Security Gateway Software Blade, relevant only when Logging and Status blade is ena...
* `smb_logs_settings` - Logs settings that apply to Quantum Spark Appliances that run Gaia Embedded OS.smb_logs_settings blocks are documented below.
* `trust_details` - Details for trusted communication.trust_details blocks are documented below.
* `trust_method` - Trust method that was used for establishing communication.
* `workforce_ai` - Workforce AI Security blade enabled. Requires content awareness blade and version R82.20 or higher to be enabled.
* `zero_phishing_settings` - Fqdn settings.zero_phishing_settings blocks are documented below.


`advanced_settings` supports the following:

* `connection_persistence` - Handling established connections when installing a new policy.
* `sam` - SAM.sam blocks are documented below.


`application_control_and_url_filtering_settings` supports the following:

* `global_settings_mode` - Whether to override global settings or not.
* `override_global_settings` - override global settings object.override_global_settings blocks are documented below.


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
* `bypass_on_client_failure` - Bypass HTTPS inspection on client failure.bypass_on_client_failure blocks are documented below.
* `bypass_under_load` - Bypass HTTPS inspection under load.bypass_under_load blocks are documented below.
* `deployment_mode` - HTTPS inspection deployment mode.
* `outbound_certificate` - Outbound HTTPS inspection certificate.outbound_certificate blocks are documented below.


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
* `ad_query` - AD Query source enabled.
* `collecting_identities` - This gateway collects identities.
* `identity_based_enforcement` - Configures this object as a PEP-only object - identity-based enforcement (PEP).
* `identity_web_api` - Identity Web API source enabled.
* `identity_web_api_settings` - Identity Web API settings.identity_web_api_settings blocks are documented below.
* `radius_accounting` - Radius Accounting source enabled.
* `terminal_servers` - Terminal Servers source enabled.


`interfaces` supports the following:

* `name` - Object name. Must be unique in the domain.
* `ipv4_address` - IPv4 address.
* `ipv6_address` - IPv6 address.
* `network_mask` - IPv4 or IPv6 network mask. If both masks are required use ipv4-network-mask and ipv6-network-mask fields explicitly. Instead of providing mask itself it is possible to specify IPv4 or IPv6 mask length in mask-length field. If both masks length are required use ipv4-mask-length and  ipv6-mask-length fields explicitly.
* `ipv4_network_mask` - IPv4 network address.
* `ipv6_network_mask` - IPv6 network address.
* `ipv4_mask_length` - IPv4 network mask length.
* `ipv6_mask_length` - IPv6 network mask length.
* `anti_spoofing` - N/A
* `anti_spoofing_settings` - N/Aanti_spoofing_settings blocks are documented below.
* `security_zone` - N/A
* `security_zone_settings` - N/Asecurity_zone_settings blocks are documented below.
* `tags` - Collection of tag identifiers.tags blocks are documented below.
* `topology` - N/A
* `topology_settings` - N/Atopology_settings blocks are documented below.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.
* `ignore_errors` - Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.
* `dynamic_ip` - Enable dynamic interface.
* `network_interface_type` - Type of network interface.

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

`nat_settings` supports the following:

* `auto_rule` - Whether to add automatic address translation rules.
* `ipv4_address` - IPv4 address.
* `ipv6_address` - IPv6 address.
* `hide_behind` - Hide behind method. This parameter is forbidden in case "method" parameter is "static".
* `install_on` - Which gateway should apply the NAT translation.
* `method` - NAT translation method.
* `apply_control_connections` - This option performs NAT on VPN control connections to and from this object.


`platform_portal_settings` supports the following:

* `portal_web_settings` - Configuration of the portal web settings.portal_web_settings blocks are documented below.
* `certificate_settings` - Configuration of the portal certificate settings.certificate_settings blocks are documented below.
* `accessibility` - Configuration of the portal access settings.accessibility blocks are documented below.
* `enabled` - State of the web portal (enabled or disabled).


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
* `advanced` - Advanced VPN settings.advanced blocks are documented below.
* `certificates` - VPN certificates.certificates blocks are documented below.
* `clientless_vpn_settings` - Clientless VPN settings.clientless_vpn_settings blocks are documented below.
* `enable_clientless_vpn` - Enable clientless VPN.
* `exported_routes` - Exported routes.exported_routes blocks are documented below.
* `interfaces` - VPN link selection interfaces.interfaces blocks are documented below.
* `saml_portal_settings` - SAML portal settings.saml_portal_settings blocks are documented below.
* `vpn_clients` - VPN clients settings.vpn_clients blocks are documented below.


`logs_settings` supports the following:

* `alert_when_free_disk_space_below` - Enable alert when free disk space is below threshold.
* `alert_when_free_disk_space_below_threshold` - Alert when free disk space below threshold.
* `alert_when_free_disk_space_below_type` - Alert when free disk space below type.
* `before_delete_keep_logs_from_the_last_days` - Enable before delete keep logs from the last days.
* `before_delete_keep_logs_from_the_last_days_threshold` - Before delete keep logs from the last days threshold.
* `before_delete_run_script` - Enable Before delete run script.
* `before_delete_run_script_command` - Before delete run script command.
* `delete_index_files_older_than_days` - Enable delete index files older than days.
* `delete_index_files_older_than_days_threshold` - Delete index files older than days threshold.
* `delete_index_files_when_index_size_above` - Enable delete index files when index size above.
* `delete_index_files_when_index_size_above_threshold` - Delete index files when index size above threshold.
* `delete_when_free_disk_space_below` - Enable delete when free disk space below.
* `delete_when_free_disk_space_below_threshold` - Delete when free disk space below threshold.
* `detect_new_citrix_ica_application_names` - Enable detect new Citrix ICA application names.
* `distribute_logs_between_all_active_servers` - Distribute logs between all active servers.
* `forward_logs_to_log_server` - Enable forward logs to log server.
* `forward_logs_to_log_server_name` - Forward logs to log server name.
* `forward_logs_to_log_server_schedule_name` - Forward logs to log server schedule name.
* `free_disk_space_metrics` - Free disk space metrics.
* `perform_log_rotate_before_log_forwarding` - Enable perform log rotate before log forwarding.
* `reject_connections_when_free_disk_space_below_threshold` - Enable reject connections when free disk space below threshold.
* `reserve_for_packet_capture_metrics` - Reserve for packet capture metrics.
* `reserve_for_packet_capture_threshold` - Reserve for packet capture threshold.
* `rotate_log_by_file_size` - Enable rotate log by file size.
* `rotate_log_file_size_threshold` - Log file size threshold.
* `rotate_log_on_schedule` - Enable rotate log on schedule.
* `rotate_log_schedule_name` - Rotate log schedule name.
* `stop_logging_when_free_disk_space_below` - Enable stop logging when free disk space below.
* `stop_logging_when_free_disk_space_below_threshold` - Stop logging when free disk space below threshold.
* `turn_on_qos_logging` - Enable turn on QoS Logging.
* `update_account_log_every` - Update account log in every amount of seconds.
* `alert_when_free_disk_space_below_metrics` - Alert when free disk space below metrics.
* `include_tcp_state_information` - Include TCP state information. Relevant only when Firewall blade is enabled.


`sam` supports the following:

* `forward_to_other_sam_servers` - Forward SAM clients' requests to other SAM servers.
* `use_early_versions` - Use early versions compatibility mode.use_early_versions blocks are documented below.
* `purge_sam_file` - Purge SAM File.purge_sam_file blocks are documented below.


`override_global_settings` supports the following:

* `fail_mode` - Fail mode - allow or block all requests.
* `website_categorization` - Website categorization object.website_categorization blocks are documented below.


`bypass_on_failure` supports the following:

* `override_profile` - Override profile of global configuration.
* `value` - Override value.<br><font color="red">Required only for</font> 'override-profile' is True.
* `profile_value` - The value inherited from the profile.


`site_categorization_allow_mode` supports the following:

* `override_profile` - Override profile of global configuration.
* `value` - Override value.<br><font color="red">Required only for</font> 'override-profile' is True.
* `profile_value` - The value inherited from the profile.


`deny_untrusted_server_cert` supports the following:

* `override_profile` - Override profile of global configuration.
* `value` - Override value.<br><font color="red">Required only for</font> 'override-profile' is True.
* `profile_value` - The value inherited from the profile.


`deny_revoked_server_cert` supports the following:

* `override_profile` - Override profile of global configuration.
* `value` - Override value.<br><font color="red">Required only for</font> 'override-profile' is True.
* `profile_value` - The value inherited from the profile.


`deny_expired_server_cert` supports the following:

* `override_profile` - Override profile of global configuration.
* `value` - Override value.<br><font color="red">Required only for</font> 'override-profile' is True.
* `profile_value` - The value inherited from the profile.


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
* `cache_mode` - Identity cache mode.cache_mode blocks are documented below.
* `cache_mode_duration` - Identity cache mode duration.cache_mode_duration blocks are documented below.
* `receive_from_infinity_identity` - Whether to receive identities from Infinity Identity.
* `scaled_sharing` - Whether scaled identity sharing is enabled.


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


`portal_web_settings` supports the following:

* `aliases` - List of URL aliases that are redirected to the main portal URL.aliases blocks are documented below.
* `main_url` - The main URL for the web portal.
* `ip_address` - Optional IP address to be used for the portal URL.


`certificate_settings` supports the following:

* `base64_certificate` - The certificate file encoded in Base64 with padding.
  This file must be in the *.p12 format.
* `base64_password` - Password (encoded in Base64 with padding) for the certificate file.
* `certificate` - The certificate.
* `certificate_dn` - The certificate distinguished name.
* `certificate_valid_from` - The date from which the certificate is valid.
* `certificate_valid_to` - The date until which the certificate is valid.


`accessibility` supports the following:

* `allow_access_from` - Allowed access to the web portal (based on interfaces, or security policy).
* `internal_access_settings` - Configuration of the additional portal access settings for internal interfaces only.internal_access_settings blocks are documented below.


`portal_web_settings` supports the following:

* `aliases` - List of URL aliases that are redirected to the main portal URL.aliases blocks are documented below.
* `main_url` - The main URL for the web portal.
* `ip_address` - Optional: IP address for the web portal to use, if your DNS server fails to resolve the main portal URL. Note: If your DNS server resolves the main po...


`certificate_settings` supports the following:

* `base64_certificate` - The certificate file encoded in Base64 with padding.
  This file must be in the *.p12 format.
* `base64_password` - Password (encoded in Base64 with padding) for the certificate file.
* `certificate` - The certificate.
* `certificate_dn` - The DN (Distinguished Name) of the certificate.
* `certificate_valid_from` - The date, from which the certificate is valid.
* `certificate_valid_to` - The certificate expiration date.


`accessibility` supports the following:

* `allow_access_from` - Allowed access to the web portal (based on interfaces, or security policy).
* `internal_access_settings` - Configuration of the additional portal access settings for internal interfaces only.internal_access_settings blocks are documented below.


`authentication` supports the following:

* `authentication_clients` - Collection of VPN Authentication clients identified by the name or UID.authentication_clients blocks are documented below.
* `dynamic_id_settings` - Dynamic ID settings, relevant only when "override-global-dynamic-id-settings" is true.dynamic_id_settings blocks are documented below.
* `override_global_dynamic_id_settings` - Override global dynamic ID settings.
* `send_machine_certificate` - Configure when to send machine certificate.
* `single_authentication_client` - Settings for clients that support only single authentication method.single_authentication_client blocks are documented below.


`link_selection` supports the following:

* `dns_resolving_hostname` - DNS Resolving Hostname. Must be set when "ip-selection" was selected to be "dns-resolving-from-hostname".
* `outgoing_link_tracking` - Outgoing link tracking method.
* `probing_settings` - Probing settings configuration. Only available when "ip-selection" is "use-probing-with-high-availability" or "use-probing-with-load-sharing".probing_settings blocks are documented below.
* `responding_traffic` - Responding traffic route selection method.
* `route_selection_method` - Outgoing route selection method when initiating a tunnel.
* `selected_ip` - Selected IP address. Must be set when "source-ip-selection" was selected to be "manual".
* `source_ip_selection` - Source IP address selection method for outgoing traffic.


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
* `certificate_settings` - Configuration of the portal certificate.certificate_settings blocks are documented below.
* `portal_web_settings` - Configuration of the portal web settings.portal_web_settings blocks are documented below.


`authorized_clients` supports the following:

* `client` - Host / Network Group Name or UID.
* `client_secret` - Client Secret.


`authentication_settings` supports the following:

* `users_directories` - Users directories.users_directories blocks are documented below.


`client_access_permissions` supports the following:

* `accessibility` - Configuration of the portal access settings.accessibility blocks are documented below.
* `certificate_settings` - Configuration of the portal certificate.certificate_settings blocks are documented below.
* `portal_web_settings` - Configuration of the portal web settings.portal_web_settings blocks are documented below.


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
* `ip_address` - Optional: IP address for the web portal to use, if your DNS server fails to resolve the main portal URL. Note: If your DNS server resolves the main po...


`certificate_settings` supports the following:

* `base64_certificate` - The certificate file encoded in Base64 with padding.
  This file must be in the *.p12 format.
* `base64_password` - Password (encoded in Base64 with padding) for the certificate file.
* `certificate` - The certificate.
* `certificate_dn` - The DN (Distinguished Name) of the certificate.
* `certificate_valid_from` - The date, from which the certificate is valid.
* `certificate_valid_to` - The certificate expiration date.


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


`communication_with_servers_behind_nat` supports the following:

* `override_profile` - Whether to override the Server (Check Point Host) object configuration.
* `value` - according-to-topology: Use the original or translated IP address of the server based on the Topology of Security Gateway interfaces.<br>original-ip-on...


`fetch_policy_scheduler` supports the following:

* `enabled` - Indicates if the Security Gateway will fetch policy according to a schedule (true) or manually (false).
* `schedule` - Scheduled Event for fetching policy by.<br><font color='red'>Will be applied only</font> when the `enabled` field is set to true.<br>When not defined ...


`bypass_on_client_failure` supports the following:

* `override_profile` - Whether to override the value inherited from the profile.
* `profile_value` - The value inherited from the profile.
* `value` - Whether to bypass on client failure.


`bypass_under_load` supports the following:

* `value` - Whether to bypass under load.


`outbound_certificate` supports the following:

* `override_profile` - Whether to override the value inherited from the profile.
* `profile_value` - The value inherited from the profile.
* `value` - Outbound certificate identified by the name or UID.


`certificate_settings` supports the following:

* `certificate` - The certificate.
* `certificate_dn` - The certificate distinguished name.
* `certificate_valid_from` - The date from which the certificate is valid.
* `certificate_valid_to` - The date until which the certificate is valid.


`portal_web_settings` supports the following:

* `aliases` - List of URL aliases that are redirected to the main portal URL.
* `ip_address` - Optional IP address to be used for the portal URL.
* `main_url` - The main URL for the portal.


`certificate_settings` supports the following:

* `certificate` - The certificate.
* `certificate_dn` - The certificate distinguished name.
* `certificate_valid_from` - The date from which the certificate is valid.
* `certificate_valid_to` - The date until which the certificate is valid.


`portal_web_settings` supports the following:

* `aliases` - List of URL aliases that are redirected to the main portal URL.
* `ip_address` - Optional IP address to be used for the portal URL.
* `main_url` - The main URL for the portal.


`cache_mode` supports the following:

* `override_profile` - Whether to override the value inherited from the profile.
* `profile_value` - The value inherited from the profile.
* `value` - Whether the identity cache is enabled.


`cache_mode_duration` supports the following:

* `override_profile` - Whether to override the value inherited from the profile.
* `profile_value` - The duration inherited from the profile, in minutes.
* `value` - Identity cache duration in minutes.


`identity_web_api_settings` supports the following:

* `authentication_settings` - Authentication settings for Identity Web API.authentication_settings blocks are documented below.
* `authorized_clients` - Authorized clients.authorized_clients blocks are documented below.
* `client_access_permissions` - Identity Web API accessibility settings.client_access_permissions blocks are documented below.


`authentication_settings` supports the following:

* `users_directories` - Users directories.users_directories blocks are documented below.


`users_directories` supports the following:

* `external_user_profile` - External user profile.
* `internal_users` - Internal users.
* `specific` - LDAP AU objects identified by the name or UID.
* `users_from_external_directories` - Users from external directories.


`authorized_clients` supports the following:

* `client` - Host / Network Group Name or UID.


`client_access_permissions` supports the following:

* `accessibility` - Configuration of the portal access settings.accessibility blocks are documented below.
* `certificate_settings` - Configuration of the portal certificate.certificate_settings blocks are documented below.
* `portal_web_settings` - Configuration of the portal web settings.portal_web_settings blocks are documented below.


`accessibility` supports the following:

* `allow_access_from` - Allowed access to the web portal (based on interfaces, or security policy).
* `internal_access_settings` - Configuration of the additional portal access settings for internal interfaces only.internal_access_settings blocks are documented below.


`internal_access_settings` supports the following:

* `dmz` - Controls portal access settings for internal interfaces, whose topology is set to "DMZ".
* `undefined` - Controls portal access settings for internal interfaces, whose topology is set to "Undefined".
* `vpn` - Controls portal access settings for interfaces that are part of a VPN Encryption Domain.


`certificate_settings` supports the following:

* `certificate` - The certificate.
* `certificate_dn` - The certificate distinguished name.
* `certificate_valid_from` - The date from which the certificate is valid.
* `certificate_valid_to` - The date until which the certificate is valid.


`portal_web_settings` supports the following:

* `aliases` - List of URL aliases that are redirected to the main portal URL.
* `ip_address` - Optional IP address to be used for the portal URL.
* `main_url` - The main URL for the portal.


`smb_logs_settings` supports the following:

* `alert_when_queue_is_full` - Alert when queue is full enabled.
* `alert_when_queue_is_full_type` - Alert when queue is full type.
* `detect_new_citrix_ica_application_names` - Detect new citrix ica application names enabled.
* `stop_logging_when_queue_reaches_maximal_capacity` - Stop logging when queue reaches maximal capacity enabled.
* `stop_logging_when_queue_reaches_maximal_capacity_threshold` - Stop logging when queue reaches maximal capacity threshold.
* `turn_on_qos_logging` - Turn on qos logging enabled.
* `update_account_log_every` - Update account log in every amount of seconds.


`trust_details` supports the following:

* `authentication_token` - Authentication token to use on the Gateway side to establish the communication between the Gateway and the Management Server (applies only to Smart-1 ...
* `cloud_communication_details` - Details about the communication status with cloud (applies only to Smart-1 Cloud).cloud_communication_details blocks are documented below.
* `gateway_mac_address` - Use the Security Gateway MAC address, relevant for the gateway_mac_address identification-method.
* `identification_method` - How to identify the gateway (relevant for DAIP gateways only).
* `status` - Status of the trusted communication with the Security Gateway.
* `token_expiration_date` - Details about the communication status with cloud (applies only to Smart-1 Cloud).


`cloud_communication_details` supports the following:

* `ip` - IP address used for communication between the Gateway and the Management Server (used when 'auto-generate-ip=true').
* `status` - Status of communication between the Gateway and the Management Server.


`advanced` supports the following:

* `enable_nat_traversal` - Enable NAT traversal.
* `enable_wire_mode` - Enable wire mode.
* `enable_wire_mode_log_traffic` - Log traffic in wire mode.
* `shutdown_on_gateway_restart` - Shutdown VPN tunnels on gateway restart.
* `tunnel_sharing_mode` - Tunnel sharing mode.
* `wire_mode_interfaces` - Wire mode interfaces.


`dynamic_id_settings` supports the following:

* `advanced_settings` - Advanced Dynamic ID configuration settings.advanced_settings blocks are documented below.
* `sms_provider_and_email_settings` - SMS provider and email configuration.
* `sms_provider_credentials` - SMS provider credentials configuration.sms_provider_credentials blocks are documented below.


`advanced_settings` supports the following:

* `country_code` - Country code for SMS services.
* `dynamic_id_message` - Dynamic ID message displayed to users.
* `enable_display_user_details` - Enable display of user details.
* `otp_settings` - One Time Password configuration settings.otp_settings blocks are documented below.
* `user_details_retrieval` - User details retrieval method.


`otp_settings` supports the following:

* `expiration` - One time password expiration (in minutes).
* `length` - Length of one time password.
* `max_attempts` - Number of times users can attempt to enter the one time password before the entire authentication process restarts.


`sms_provider_credentials` supports the following:

* `api_id` - SMS provider API ID.
* `username` - SMS provider username.


`single_authentication_client` supports the following:

* `allow_multiple_authentication_clients` - Allow clients that support multiple authentication methods to connect.
* `client_display_settings` - Client display configuration settings.client_display_settings blocks are documented below.
* `display_name` - Display name for the authentication method.
* `enabled` - Allow clients that support only single authentication method.
* `method` - Authentication method type.
* `personal_certificate` - Personal certificate authentication settings, relevant only when method is "personal-certificate".personal_certificate blocks are documented below.
* `radius` - RADIUS authentication settings, relevant only when method is "radius".radius blocks are documented below.
* `secur_id` - SecurID authentication settings, relevant only when method is "secur-id".secur_id blocks are documented below.


`client_display_settings` supports the following:

* `headline` - Display headline for authentication dialog.
* `password_label` - Label for password field.
* `username_label` - Label for username field.


`personal_certificate` supports the following:

* `dn_concurrence` - DN part occurrence number.
* `dn_part` - DN part to extract.
* `fetch_username_from` - Fetch username from.
* `source` - Certificate source field.
* `storage_type` - Certificate storage type.


`radius` supports the following:

* `ask_user_password` - Ask user for password during authentication.
* `server` - Server object identified by the name or UID.


`secur_id` supports the following:

* `server` - Server object identified by the name or UID.
* `token_card_type` - Token card type.


`certificates` supports the following:

* `base64_certificate` - The certificate encoded in Base64.
* `certificate_authority` - Certificate authority identified by the name or UID.
* `distinguished_name` - Certificate distinguished name.
* `expiration_date` - Certificate expiration date.expiration_date blocks are documented below.
* `name` - Certificate name.
* `status` - Certificate status.
* `stored_at` - Where the certificate is stored.


`expiration_date` supports the following:

* `iso_8601` - Date and time represented in international ISO 8601 format.
* `posix` - Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.


`clientless_vpn_settings` supports the following:

* `accept_only_3des` - Accept only 3DES.
* `certificate_gateway_authentication` - Certificate gateway authentication.
* `client_authentication` - Client authentication.
* `concurrent_servers_or_processes` - Number of concurrent servers or processes.


`exported_routes` supports the following:

* `custom_routes` - Export custom routes.
* `custom_routes_object` - Custom routes object identified by the name or UID.
* `internal_interfaces` - Export internal interfaces.
* `static_routes` - Export static routes.


`interfaces` supports the following:

* `interface_name` - The name of the interface.
* `ip_version` - The IP version of the interface's IP address.
* `next_hop_ip` - The IP address of the next hop.
* `priority` - Priority of a "Backup" interface.
* `redundancy_mode` - Interface redundancy mode (Active/Backup).
* `static_nat_ip` - The NATed IPv4 address that hides the source IPv4 address of outgoing connections.


`probing_settings` supports the following:

* `primary_address` - Primary IP address to use. Must be one of the addresses from "probed-interface-list". Required when "use-primary-address" is true.
* `probed_interface_list` - List of specific IP addresses to probe. Only relevant when "probed-interfaces" is set to "specific".
* `probed_interfaces` - Specifies whether to probe all addresses defined in the topology tab or specific addresses.
* `probing_method` - Probing method.
* `use_primary_address` - Whether to use a primary address for high availability probing.


`saml_portal_settings` supports the following:

* `accessibility` - Configuration of the portal access settings.accessibility blocks are documented below.
* `certificate_settings` - Configuration of the SAML portal certificate.certificate_settings blocks are documented below.
* `enabled` - Whether the SAML portal is enabled.
* `portal_web_settings` - Configuration of the SAML portal web settings.portal_web_settings blocks are documented below.


`accessibility` supports the following:

* `allow_access_from` - Allowed access to the SAML portal.
* `internal_access_settings` - Configuration of the additional portal access settings for internal interfaces only.internal_access_settings blocks are documented below.


`internal_access_settings` supports the following:

* `dmz` - Controls portal access settings for internal interfaces, whose topology is set to "DMZ".
* `undefined` - Controls portal access settings for internal interfaces, whose topology is set to "Undefined".
* `vpn` - Controls portal access settings for interfaces that are part of a VPN Encryption Domain.


`certificate_settings` supports the following:

* `certificate` - The certificate.
* `certificate_dn` - The certificate distinguished name.
* `certificate_valid_from` - The date from which the certificate is valid.
* `certificate_valid_to` - The date until which the certificate is valid.


`portal_web_settings` supports the following:

* `aliases` - List of URL aliases that are redirected to the main portal URL.
* `ip_address` - Optional IP address to be used for the portal URL.
* `main_url` - The main URL for the portal.


`vpn_clients` supports the following:

* `enable_capsule_vpn_connect` - Enable Capsule VPN Connect client.
* `enable_cp_mobile_for_windows` - Enable Check Point Mobile for Windows client.
* `enable_endpoint_security_vpn` - Enable Endpoint Security VPN client.
* `enable_secu_remote` - Enable SecuRemote client.
* `enable_ssl_network_extender` - Enable SSL Network Extender client.
* `gateway_authentication_certificate` - Gateway authentication certificate.


`zero_phishing_settings` supports the following:

* `gateway_fqdn_mode` - Manual Fqdn.
* `manual_fqdn` - Zero Phishing gateway FQDN.
