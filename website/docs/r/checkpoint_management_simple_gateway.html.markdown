---
layout: "checkpoint"
page_title: "checkpoint_management_simple_gateway"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-simple-gateway"
description: |-
This resource allows you to execute Check Point Simple Gateway.
---

# checkpoint_management_simple_gateway

This resource allows you to execute Check Point Simple Gateway.

## Example Usage


```hcl
resource "checkpoint_management_simple_gateway" "example" {
  name = "gw1"
  ipv4_address = "192.0.2.1"
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
* `content_awareness` - (Optional) Content Awareness blade enabled. 
* `enable_https_inspection` - (Optional) Enable HTTPS Inspection after defining an outbound inspection certificate. <br>To define the outbound certificate use outbound inspection certificate API. 
* `fetch_policy` - (Optional) Security management server(s) to fetch the policy from.fetch_policy blocks are documented below.
* `firewall` - (Optional) Firewall blade enabled. 
* `firewall_settings` - (Optional) N/Afirewall_settings blocks are documented below.
* `hit_count` - (Optional) Hit count tracks the number of connections each rule matches. 
* `https_inspection` - (Optional) HTTPS inspection.https_inspection blocks are documented below.
* `icap_server` - (Optional) ICAP Server enabled. 
* `identity_awareness` - (Optional) Identity awareness blade enabled. 
* `identity_awareness_settings` - (Optional) Gateway Identity Awareness settings.identity_awareness_settings blocks are documented below.
* `interfaces` - (Optional) Network interfaces.interfaces blocks are documented below.
* `ipv4_address` - (Optional) IPv4 address. 
* `ipv6_address` - (Optional) IPv6 address. 
* `ips` - (Optional) Intrusion Prevention System blade enabled.
* `ips_settings` - (Optional) Gateway IPS settings.ips_settings blocks are documented below.
* `ips_update_policy` - (Optional) Specifies whether the IPS will be downloaded from the Management or directly to the Gateway. 
* `nat_hide_internal_interfaces` - (Optional) Hide internal networks behind the Gateway's external IP. 
* `nat_settings` - (Optional) NAT settings.nat_settings blocks are documented below.
* `one_time_password` - (Optional) N/A 
* `os_name` - (Optional) Gateway platform operating system. 
* `platform_portal_settings` - (Optional) Platform portal settings.platform_portal_settings blocks are documented below.
* `proxy_settings` - (Optional) Proxy Server for Gateway.proxy_settings blocks are documented below.
* `qos` - (Optional) QoS. 
* `save_logs_locally` - (Optional) Save logs locally on the gateway. 
* `send_alerts_to_server` - (Optional) Server(s) to send alerts to.send_alerts_to_server blocks are documented below.
* `send_logs_to_backup_server` - (Optional) Backup server(s) to send logs to.send_logs_to_backup_server blocks are documented below.
* `send_logs_to_server` - (Optional) Server(s) to send logs to.send_logs_to_server blocks are documented below.
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `threat_emulation` - (Optional) Threat Emulation blade enabled. 
* `threat_extraction` - (Optional) Threat Extraction blade enabled. 
* `threat_prevention_mode` - (Optional) The mode of Threat Prevention to use. When using Autonomous Threat Prevention, disabling the Threat Prevention blades is not allowed. 
* `url_filtering` - (Optional) URL Filtering blade enabled. 
* `usercheck_portal_settings` - (Optional) UserCheck portal settings.usercheck_portal_settings blocks are documented below.
* `version` - (Optional) Gateway platform version. 
* `vpn` - (Optional) VPN blade enabled. 
* `vpn_settings` - (Optional) Gateway VPN settings.vpn_settings blocks are documented below.
* `zero_phishing` - (Optional) Zero Phishing blade enabled. 
* `zero_phishing_fqdn` - (Optional) Zero Phishing gateway FQDN. **Deprecated** - use `zero_phishing_settings.manual_fqdn` instead.
* `logs_settings` - (Optional) Logs settings that apply to Quantum Security Gateways that run Gaia OS.logs_settings blocks are documented below.
* `show_portals_certificate` - (Optional) Indicates whether to show the portals certificate value in the reply. 
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `groups` - (Optional) Collection of group identifiers.groups blocks are documented below.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
* `accept_syslog_messages` - (Optional) Enables the gateway accept syslog messages, relevant only when Logging and Status/Network Policy Management blades are enabled.
* `anti_spam_and_email_security` - (Optional) Enables Anti-Spam & Email-Security blade.
* `auto_generate_ip` - (Optional) Use an automatically generated IP address for the Gateway object (applies only to Smart-1 Cloud).
* `auto_topology_custom_recalculation_time` - (Optional) Auto topology custom recalculation time (seconds).
* `auto_topology_use_custom_recalculation_time` - (Optional) Auto topology to use custom recalculation time instead of default.
* `autonomous_system_number` - (Optional) The Autonomous System Number (ASN). It is automatically fetched from the Security Gateway object. You can change this value only for externally manage...
* `communication_with_servers_behind_nat` - (Optional) Gateway behind NAT communications settings with the server.communication_with_servers_behind_nat blocks are documented below.
* `data_loss_prevention` - (Optional) Data Loss Prevention blade.
* `dns_server` - (Optional) DNS Server.
* `enable_log_indexing` - (Optional) Enable log indexing, The Log Indexing uses more storage to provide fast log queries, relevant only when Logging and Status blade is enabled.
* `export_logs_to_servers` - (Optional) Export logs to syslog/SIEM servers. NOTE:After you configure a Log Exporter, you must run Install Database. Relevant only when Logging and Status/Netw...
* `fetch_policy_scheduler` - (Optional) Fetch policy functionality settings.fetch_policy_scheduler blocks are documented below.
* `hardware_subtype` - (Optional) Gateway type (relevant only for Spark gateways).
* `install_policy_without_push` - (Optional) Specifies whether the policy is pushed to the gateway during policy installation, or whether the gateway should fetch the policy.
* `interfaces_topology_settings` - (Optional) Topology setting for all interfaces on a Security Gateway. Default for Security Gateways that run Gaia OS: 'per interface'. Default for Quantum Spark gateways t.
* `mobile_access` - (Optional) Mobile Access blade.
* `monitoring` - (Optional) Enables Real Time Monitoring blade.
* `policy_server` - (Optional) Policy Server blade.
* `rtm_counters_report` - (Optional) Enables monitoring blades system counters report (e.g CPU Usage,Memory Usage).
* `rtm_traffic_report` - (Optional) Enables monitoring blades traffic report.
* `rtm_traffic_report_per_connection` - (Optional) Enables Monitoring blade traffic report per connection.
* `smart_event_intro_correlation_unit` - (Optional) Enables the gateway use SmartEvent intro correlation unit with one Security Gateway Software Blade. Relevant only when the Logging and Status blade is enabled,.
* `trust_method` - (Optional) Establish the trust communication method.
* `trust_settings` - (Optional) Settings for the trusted communication establishment.trust_settings blocks are documented below.
* `workforce_ai` - (Optional) Workforce AI Security blade enabled. Requires content awareness blade and version R82.20 or higher to be enabled.
* `zero_phishing_settings` - (Optional) Fqdn settings.zero_phishing_settings blocks are documented below.


`advanced_settings` supports the following:

* `connection_persistence` - (Optional) Handling established connections when installing a new policy. 
* `sam` - (Optional) SAM.sam blocks are documented below.


`application_control_and_url_filtering_settings` supports the following:

* `global_settings_mode` - (Optional) Whether to override global settings or not. 
* `override_global_settings` - (Optional) override global settings object.override_global_settings blocks are documented below.


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
* `bypass_on_client_failure` - (Optional) Bypass HTTPS inspection on client failure.bypass_on_client_failure blocks are documented below.
* `bypass_under_load` - (Optional) Bypass HTTPS inspection under load.bypass_under_load blocks are documented below.
* `deployment_mode` - (Optional) HTTPS inspection deployment mode.
* `outbound_certificate` - (Optional) Outbound HTTPS inspection certificate.outbound_certificate blocks are documented below.


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
* `identity_based_enforcement` - (Optional) ON: Configures this object as a PEP-only object - identity-based enforcement (PEP) is enabled.<br>OFF: Configures this object as a PDP-only object - identity-ba.
* `identity_web_api` - (Optional) Enable Identity Web API source.
* `identity_web_api_settings` - (Optional) Identity Web API settings.identity_web_api_settings blocks are documented below.


`interfaces` supports the following:

* `name` - (Optional) Object name. Must be unique in the domain. 
* `ipv4_address` - (Optional) IPv4 address. 
* `ipv6_address` - (Optional) IPv6 address. 
* `network_mask` - (Optional) IPv4 or IPv6 network mask. If both masks are required use ipv4-network-mask and ipv6-network-mask fields explicitly. Instead of providing mask itself it is possible to specify IPv4 or IPv6 mask length in mask-length field. If both masks length are required use ipv4-mask-length and  ipv6-mask-length fields explicitly. 
* `ipv4_network_mask` - (Optional) IPv4 network address. 
* `ipv6_network_mask` - (Optional) IPv6 network address. 
* `ipv4_mask_length` - (Optional) IPv4 network mask length. 
* `ipv6_mask_length` - (Optional) IPv6 network mask length. 
* `anti_spoofing` - (Optional) N/A 
* `anti_spoofing_settings` - (Optional) N/Aanti_spoofing_settings blocks are documented below.
* `security_zone` - (Optional) N/A 
* `security_zone_settings` - (Optional) N/Asecurity_zone_settings blocks are documented below.
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `topology` - (Optional) N/A 
* `topology_settings` - (Optional) N/Atopology_settings blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
* `dynamic_ip` - (Optional) The Topology of interface with Dynamic IP is set to Automatic - External.

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

`nat_settings` supports the following:

* `auto_rule` - (Optional) Whether to add automatic address translation rules. 
* `ipv4_address` - (Optional) IPv4 address. 
* `ipv6_address` - (Optional) IPv6 address. 
* `hide_behind` - (Optional) Hide behind method. This parameter is forbidden in case "method" parameter is "static". 
* `install_on` - (Optional) Which gateway should apply the NAT translation. 
* `method` - (Optional) NAT translation method. 
* `apply_control_connections` - (Optional) This option performs NAT on VPN control connections to and from this object.


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
* `advanced` - (Optional) Advanced VPN settings.advanced blocks are documented below.
* `clientless_vpn_settings` - (Optional) Clientless VPN settings.clientless_vpn_settings blocks are documented below.
* `enable_clientless_vpn` - (Optional) Enable clientless VPN.
* `exported_routes` - (Optional) Exported routes.exported_routes blocks are documented below.
* `interfaces` - (Optional) Enhanced Link Selection Interfaces.interfaces blocks are documented below.
* `saml_portal_settings` - (Optional) SAML portal settings.saml_portal_settings blocks are documented below.
* `vpn_clients` - (Optional) VPN clients settings.vpn_clients blocks are documented below.


`logs_settings` supports the following:

* `alert_when_free_disk_space_below` - (Optional) Enable alert when free disk space is below threshold. 
* `alert_when_free_disk_space_below_threshold` - (Optional) Alert when free disk space below threshold. 
* `alert_when_free_disk_space_below_type` - (Optional) Alert when free disk space below type. 
* `before_delete_keep_logs_from_the_last_days` - (Optional) Enable before delete keep logs from the last days. 
* `before_delete_keep_logs_from_the_last_days_threshold` - (Optional) Before delete keep logs from the last days threshold. 
* `before_delete_run_script` - (Optional) Enable Before delete run script. 
* `before_delete_run_script_command` - (Optional) Before delete run script command. 
* `delete_index_files_older_than_days` - (Optional) Enable delete index files older than days. 
* `delete_index_files_older_than_days_threshold` - (Optional) Delete index files older than days threshold. 
* `delete_index_files_when_index_size_above` - (Optional) Enable delete index files when index size above. 
* `delete_index_files_when_index_size_above_threshold` - (Optional) Delete index files when index size above threshold. 
* `delete_when_free_disk_space_below` - (Optional) Enable delete when free disk space below. 
* `delete_when_free_disk_space_below_threshold` - (Optional) Delete when free disk space below threshold. 
* `detect_new_citrix_ica_application_names` - (Optional) Enable detect new Citrix ICA application names. 
* `distribute_logs_between_all_active_servers` - (Optional) Distribute logs between all active servers. 
* `forward_logs_to_log_server` - (Optional) Enable forward logs to log server. 
* `forward_logs_to_log_server_name` - (Optional) Forward logs to log server name. 
* `forward_logs_to_log_server_schedule_name` - (Optional) Forward logs to log server schedule name. 
* `free_disk_space_metrics` - (Optional) Free disk space metrics. 
* `perform_log_rotate_before_log_forwarding` - (Optional) Enable perform log rotate before log forwarding. 
* `reject_connections_when_free_disk_space_below_threshold` - (Optional) Enable reject connections when free disk space below threshold. 
* `reserve_for_packet_capture_metrics` - (Optional) Reserve for packet capture metrics. 
* `reserve_for_packet_capture_threshold` - (Optional) Reserve for packet capture threshold. 
* `rotate_log_by_file_size` - (Optional) Enable rotate log by file size. 
* `rotate_log_file_size_threshold` - (Optional) Log file size threshold. 
* `rotate_log_on_schedule` - (Optional) Enable rotate log on schedule. 
* `rotate_log_schedule_name` - (Optional) Rotate log schedule name. 
* `stop_logging_when_free_disk_space_below` - (Optional) Enable stop logging when free disk space below. 
* `stop_logging_when_free_disk_space_below_threshold` - (Optional) Stop logging when free disk space below threshold. 
* `turn_on_qos_logging` - (Optional) Enable turn on QoS Logging. 
* `update_account_log_every` - (Optional) Update account log in every amount of seconds. 
* `include_tcp_state_information` - (Optional) Include TCP state information. Relevant only when Firewall blade is enabled.


`sam` supports the following:

* `forward_to_other_sam_servers` - (Optional) Forward SAM clients' requests to other SAM servers. 
* `use_early_versions` - (Optional) Use early versions compatibility mode.use_early_versions blocks are documented below.
* `purge_sam_file` - (Optional) Purge SAM File.purge_sam_file blocks are documented below.


`override_global_settings` supports the following:

* `fail_mode` - (Optional) Fail mode - allow or block all requests. 
* `website_categorization` - (Optional) Website categorization object.website_categorization blocks are documented below.


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
* `cache_mode` - (Optional) Identity cache mode.cache_mode blocks are documented below.
* `cache_mode_duration` - (Optional) Identity cache mode duration.cache_mode_duration blocks are documented below.
* `receive_from_infinity_identity` - (Optional) Whether to receive identities from Infinity Identity.
* `scaled_sharing` - (Optional) Whether scaled identity sharing is enabled.


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


`portal_web_settings` supports the following:

* `aliases` - (Optional) List of URL aliases that are redirected to the main portal URL.aliases blocks are documented below.
* `main_url` - (Optional) The main URL for the web portal. 
* `ip_address` - (Optional) Optional IP address to be used for the portal URL.


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
* `ip_address` - (Optional) Optional IP address to be used for the portal URL.


`certificate_settings` supports the following:

* `base64_certificate` - (Optional) The certificate file encoded in Base64 with padding. 
This file must be in the *.p12 format. 
* `base64_password` - (Optional) Password (encoded in Base64 with padding) for the certificate file. 


`accessibility` supports the following:

* `allow_access_from` - (Optional) Allowed access to the web portal (based on interfaces, or security policy). 
* `internal_access_settings` - (Optional) Configuration of the additional portal access settings for internal interfaces only.internal_access_settings blocks are documented below.


`authentication` supports the following:

* `authentication_clients` - (Optional) Collection of VPN Authentication clients identified by the name or UID.authentication_clients blocks are documented below.
* `dynamic_id_settings` - (Optional) Dynamic ID settings, relevant only when "override-global-dynamic-id-settings" is true.dynamic_id_settings blocks are documented below.
* `override_global_dynamic_id_settings` - (Optional) Override global dynamic ID settings.
* `send_machine_certificate` - (Optional) Configure when to send machine certificate.
* `single_authentication_client` - (Optional) Settings for clients that support only single authentication method.single_authentication_client blocks are documented below.


`link_selection` supports the following:

* `dns_resolving_hostname` - (Optional) DNS Resolving Hostname. Must be set when "ip-selection" was selected to be "dns-resolving-from-hostname". 
* `outgoing_link_tracking` - (Optional) Outgoing link tracking method.
* `probing_settings` - (Optional) Probing settings configuration. Only available when "ip-selection" is "use-probing-with-high-availability" or "use-probing-with-load-sharing".probing_settings blocks are documented below.
* `responding_traffic` - (Optional) Responding traffic route selection method.
* `route_selection_method` - (Optional) Outgoing route selection method when initiating a tunnel.
* `selected_ip` - (Optional) Selected IP address. Must be set when "source-ip-selection" was selected to be "manual".
* `source_ip_selection` - (Optional) Source IP address selection method for outgoing traffic.


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
* `ip_address` - (Optional) Optional IP address to be used for the portal URL.


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


`communication_with_servers_behind_nat` supports the following:

* `override_profile` - (Optional) Whether to override the Server (Check Point Host) object configuration.
* `value` - (Optional) according-to-topology: Use the original or translated IP address of the server based on the Topology of Security Gateway interfaces.<br>original-ip-only: Use on.


`fetch_policy_scheduler` supports the following:

* `enabled` - (Optional) Indicates if the Security Gateway will fetch policy according to a schedule (true) or manually (false).
* `schedule` - (Optional) Scheduled Event for fetching policy by.<br><font color='red'>Will be applied only</font> when the `enabled` field is set to true.<br>When not defined ...


`bypass_on_client_failure` supports the following:

* `override_profile` - (Optional) Whether to override the value inherited from the profile.
* `value` - (Optional) Whether to bypass on client failure.


`bypass_under_load` supports the following:

* `value` - (Optional) Whether to bypass under load.


`outbound_certificate` supports the following:

* `override_profile` - (Optional) Whether to override the value inherited from the profile.
* `value` - (Optional) Outbound certificate identified by the name or UID.


`cache_mode` supports the following:

* `override_profile` - (Optional) Whether to override the value inherited from the profile.
* `value` - (Optional) Whether the identity cache is enabled.


`cache_mode_duration` supports the following:

* `override_profile` - (Optional) Whether to override the value inherited from the profile.
* `value` - (Optional) Identity cache duration in minutes.


`identity_web_api_settings` supports the following:

* `authentication_settings` - (Optional) Authentication Settings for Identity Web Api.authentication_settings blocks are documented below.
* `authorized_clients` - (Optional) Authorized Clients.authorized_clients blocks are documented below.
* `client_access_permissions` - (Optional) Identity Web Api accessibility settings.client_access_permissions blocks are documented below.


`authentication_settings` supports the following:

* `users_directories` - (Optional) Users directories.users_directories blocks are documented below.


`users_directories` supports the following:

* `external_user_profile` - (Optional) External user profile.
* `internal_users` - (Optional) Internal users.
* `specific` - (Optional) LDAP AU objects identified by the name or UID. Must be set when 'users-from-external-directories' was selected to be 'specific'.
* `users_from_external_directories` - (Optional) Users from external directories.


`authorized_clients` supports the following:

* `client` - (Optional) Host / Network Group Name or UID.
* `client_secret` - (Optional) Client Secret.


`client_access_permissions` supports the following:

* `accessibility` - (Optional) Configuration of the portal access settings.accessibility blocks are documented below.


`accessibility` supports the following:

* `allow_access_from` - (Optional) Allowed access to the web portal (based on interfaces, or security policy).
* `internal_access_settings` - (Optional) Configuration of the additional portal access settings for internal interfaces only.internal_access_settings blocks are documented below.


`internal_access_settings` supports the following:

* `dmz` - (Optional) Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'.
* `undefined` - (Optional) Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'.
* `vpn` - (Optional) Controls portal access settings for interfaces that are part of a VPN Encryption Domain.


`trust_settings` supports the following:

* `gateway_mac_address` - (Optional) Use the Security Gateway MAC address, relevant for the gateway_mac_address identification-method.
* `identification_method` - (Optional) How to identify the gateway (relevant for DAIP gateways only).
* `initiation_phase` - (Optional) Push the certificate to the Security Gateway immediately, or wait for the Security Gateway to pull the certificate. Default value for Spark Gateway is...


`advanced` supports the following:

* `enable_nat_traversal` - (Optional) Enable NAT traversal.
* `enable_wire_mode` - (Optional) Enable wire mode.
* `enable_wire_mode_log_traffic` - (Optional) Log traffic in wire mode.
* `shutdown_on_gateway_restart` - (Optional) Shutdown VPN tunnels on gateway restart.
* `tunnel_sharing_mode` - (Optional) Tunnel sharing mode.
* `wire_mode_interfaces` - (Optional) Wire mode interfaces.


`dynamic_id_settings` supports the following:

* `advanced_settings` - (Optional) Advanced Dynamic ID configuration settings.advanced_settings blocks are documented below.
* `sms_provider_and_email_settings` - (Optional) SMS provider and email configuration.
* `sms_provider_credentials` - (Optional) SMS provider credentials configuration.sms_provider_credentials blocks are documented below.


`advanced_settings` supports the following:

* `country_code` - (Optional) Country code for SMS services.
* `dynamic_id_message` - (Optional) Dynamic ID message displayed to users.
* `enable_display_user_details` - (Optional) Enable display of user details.
* `otp_settings` - (Optional) One Time Password configuration settings.otp_settings blocks are documented below.
* `user_details_retrieval` - (Optional) User details retrieval method.


`otp_settings` supports the following:

* `expiration` - (Optional) One time password expiration (in minutes).
* `length` - (Optional) Length of one time password.
* `max_attempts` - (Optional) Number of times users can attempt to enter the one time password before the entire authentication process restarts.


`sms_provider_credentials` supports the following:

* `api_id` - (Optional) SMS provider API ID.
* `password` - (Optional) SMS provider password.
* `username` - (Optional) SMS provider username.


`single_authentication_client` supports the following:

* `allow_multiple_authentication_clients` - (Optional) Allow clients that support multiple authentication methods to connect.
* `client_display_settings` - (Optional) Client display configuration settings.client_display_settings blocks are documented below.
* `display_name` - (Optional) Display name for the authentication method.
* `enabled` - (Optional) Allow clients that support only single authentication method.
* `method` - (Optional) Authentication method type.
* `personal_certificate` - (Optional) Personal certificate authentication settings, relevant only when method is "personal-certificate".personal_certificate blocks are documented below.
* `radius` - (Optional) RADIUS authentication settings, relevant only when method is "radius".radius blocks are documented below.
* `secur_id` - (Optional) SecurID authentication settings, relevant only when method is "secur-id".secur_id blocks are documented below.


`client_display_settings` supports the following:

* `headline` - (Optional) Display headline for authentication dialog.
* `password_label` - (Optional) Label for password field.
* `username_label` - (Optional) Label for username field.


`personal_certificate` supports the following:

* `dn_concurrence` - (Optional) DN part occurrence number.
* `dn_part` - (Optional) DN part to extract.
* `fetch_username_from` - (Optional) Fetch username from.
* `source` - (Optional) Certificate source field.
* `storage_type` - (Optional) Certificate storage type.


`radius` supports the following:

* `ask_user_password` - (Optional) Ask user for password during authentication.
* `server` - (Optional) Server object identified by the name or UID.


`secur_id` supports the following:

* `server` - (Optional) Server object identified by the name or UID.
* `token_card_type` - (Optional) Token card type.


`clientless_vpn_settings` supports the following:

* `accept_only_3des` - (Optional) Accept only 3DES.
* `certificate_gateway_authentication` - (Optional) Certificate gateway authentication.
* `client_authentication` - (Optional) Client authentication.
* `concurrent_servers_or_processes` - (Optional) Number of concurrent servers or processes.


`exported_routes` supports the following:

* `custom_routes` - (Optional) Export custom routes.
* `custom_routes_object` - (Optional) Custom routes object identified by the name or UID.
* `internal_interfaces` - (Optional) Export internal interfaces.
* `static_routes` - (Optional) Export static routes.


`interfaces` supports the following:

* `interface_name` - (Optional) The name of the interface.
* `ip_version` - (Optional) The IP version of the interface's IP address (IPv4/IPv6).
* `next_hop_ip` - (Optional) The IP address of the next hop.
* `priority` - (Optional) Priority of a 'Backup' interface.
* `redundancy_mode` - (Optional) Interface redundancy mode (Active/Backup).
* `static_nat_ip` - (Optional) The NATed IPv4 address that hides the source IPv4 address of outgoing connections (applies only to IPv4).


`probing_settings` supports the following:

* `primary_address` - (Optional) Primary IP address to use. Must be one of the addresses from "probed-interface-list". Required when "use-primary-address" is true.
* `probed_interface_list` - (Optional) List of specific IP addresses to probe. Only relevant when "probed-interfaces" is set to "specific".
* `probed_interfaces` - (Optional) Specifies whether to probe all addresses defined in the topology tab or specific addresses.
* `probing_method` - (Optional) Probing method.
* `use_primary_address` - (Optional) Whether to use a primary address for high availability probing.


`saml_portal_settings` supports the following:

* `accessibility` - (Optional) Configuration of the portal access settings.accessibility blocks are documented below.
* `certificate_settings` - (Optional) Configuration of the SAML portal certificate.certificate_settings blocks are documented below.
* `portal_web_settings` - (Optional) Configuration of the SAML portal web settings.portal_web_settings blocks are documented below.


`accessibility` supports the following:

* `allow_access_from` - (Optional) Allowed access to the SAML portal.
* `internal_access_settings` - (Optional) Configuration of the additional portal access settings for internal interfaces only.internal_access_settings blocks are documented below.


`internal_access_settings` supports the following:

* `dmz` - (Optional) Controls portal access settings for internal interfaces, whose topology is set to "DMZ".
* `undefined` - (Optional) Controls portal access settings for internal interfaces, whose topology is set to "Undefined".
* `vpn` - (Optional) Controls portal access settings for interfaces that are part of a VPN Encryption Domain.


`certificate_settings` supports the following:

* `base64_certificate` - (Optional) The certificate file encoded in Base64 with padding.
* `base64_password` - (Optional) Certificate file password.


`portal_web_settings` supports the following:

* `aliases` - (Optional) List of URL aliases that are redirected to the main portal URL.
* `ip_address` - (Optional) Optional IP address to be used for the portal URL.
* `main_url` - (Optional) The main URL for the portal.


`vpn_clients` supports the following:

* `enable_capsule_vpn_connect` - (Optional) Enable Capsule VPN Connect client.
* `enable_cp_mobile_for_windows` - (Optional) Enable Check Point Mobile for Windows client.
* `enable_endpoint_security_vpn` - (Optional) Enable Endpoint Security VPN client.
* `enable_secu_remote` - (Optional) Enable SecuRemote client.
* `enable_ssl_network_extender` - (Optional) Enable SSL Network Extender client.
* `gateway_authentication_certificate` - (Optional) Gateway authentication certificate.


`zero_phishing_settings` supports the following:

* `gateway_fqdn_mode` - (Optional) Manual Fqdn.
* `manual_fqdn` - (Optional) Zero Phishing gateway FQDN.
