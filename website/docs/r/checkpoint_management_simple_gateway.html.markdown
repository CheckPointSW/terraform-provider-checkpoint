---
layout: "checkpoint"
page_title: "checkpoint_management_simple_gateway"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-simple-gateway"
description: |-
This resource allows you to execute Check Point Simple Gateway.
---

# Resource: checkpoint_management_simple_gateway

This resource allows you to execute Check Point Simple Gateway.

## Example Usage


```hcl
resource "checkpoint_management_simple_gateway" "gateway" {
    name = "mygateway"
    ipv4_address = "1.2.3.4"
    version = "R81"
    send_logs_to_server = ["mylogserver"]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `ipv4_address` - (Optional) IPv4 address. 
* `ipv6_address` - (Optional) IPv6 address.
* `interfaces` - (Optional) Gateway interfaces. interfaces blocks are documented below.
* `anti_bot` - (Optional) Anti-Bot blade enabled. 
* `anti_virus` - (Optional) Anti-Virus blade enabled. 
* `application_control` - (Optional) Application Control blade enabled.
* `content_awareness` - (Optional) Content Awareness blade enabled.
* `icap_server` - (Optional) ICAP Server enabled.
* `ips` - (Optional) Intrusion Prevention System blade enabled.
* `threat_emulation` - (Optional) Threat Emulation blade enabled.
* `threat_extraction` - (Optional) Threat Extraction blade enabled.
* `url_filtering` - (Optional) URL Filtering blade enabled.
* `firewall` - (Optional) Firewall blade enabled.
* `firewall_settings` - (Optional) Firewall settings. firewall_settings blocks are documented below.
* `vpn` - (Optional) VPN blade enabled.
* `vpn_settings` - (Optional) Gateway VPN settings. vpn_settings blocks are documented below.
* `dynamic_ip` - (Computed) Dynamic IP address.
* `version` - (Optional) Gateway platform version.
* `os_name` - (Optional) Operating system name.
* `hardware` - (Computed) Gateway platform hardware name.
* `one_time_password` - (Optional) Secure internal connection one time password. 
* `sic_name` - (Computed) Secure Internal Communication name.
* `sic_state` - (Computed) Secure Internal Communication state.
* `save_logs_locally` - (Optional) Enable save logs locally.
* `send_alerts_to_server` - (Optional) Collection of Server(s) to send alerts to identified by the name.
* `send_logs_to_backup_server` - (Optional) Collection of Backup server(s) to send logs to identified by the name.
* `send_logs_to_server` - (Optional) Collection of Server(s) to send logs to identified by the name.
* `logs_settings` - (Optional) Logs settings. logs_settings blocks are documented below.
* `color` - (Optional) Color of the object.
* `comments` - (Optional) Comments string. 
* `tags` - (Optional) Collection of tags identified by name.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 

`interfaces` supports the following:
* `name` - (Optional) Interface name. 
* `ipv4_address` - (Optional) IPv4 address. 
* `ipv6_address` - (Optional) IPv6 address. 
* `ipv4_network_mask` - (Optional) IPv4 network address.
* `ipv6_network_mask` - (Optional) IPv6 network address.
* `ipv4_mask_length` - (Optional) IPv4 network mask length.
* `ipv6_mask_length` - (Optional) IPv6 network mask length.
* `anti_spoofing` - (Optional) Anti spoofing.
* `anti_spoofing_settings` - (Optional) Anti spoofing settings. anti_spoofing_settings blocks are documented below.
* `security_zone` - (Optional) Security zone.
* `security_zone_settings` - (Optional) Security zone settings. security_zone_settings blocks are documented below.
* `topology` - (Optional) Topology.
* `topology_settings` - (Optional) Topology settings. topology_settings blocks are documented below.
* `topology_automatic_calculation` - (Computed) Shows the automatic topology calculation..
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 

`anti_spoofing_settings` supports the following:
* `action` - (Optional) If packets will be rejected (the Prevent option) or whether the packets will be monitored (the Detect option).

`security_zone_settings` supports the following:
* `auto_calculated` - (Optional) Security Zone is calculated according to where the interface leads to.
* `specific_zone` - (Optional) Security Zone specified manually.

`topology_settings` supports the following:
* `interface_leads_to_dmz` - (Optional) Whether this interface leads to demilitarized zone (perimeter network).
* `ip_address_behind_this_interface` - (Optional) Ip address behind this interface.
* `specific_network` - (Optional) Network behind this interface.

`firewall_settings` supports the following:
* `auto_calculate_connections_hash_table_size_and_memory_pool` - (Optional) Auto calculate connections hash table size and memory pool. 
* `auto_maximum_limit_for_concurrent_connections` - (Optional) Auto maximum limit for concurrent connections.
* `connections_hash_size` - (Optional) Connections hash size.
* `maximum_limit_for_concurrent_connections` - (Optional) Maximum limit for concurrent connections.
* `maximum_memory_pool_size` - (Optional) Maximum memory pool size.
* `memory_pool_size` - (Optional) Memory pool size.

`vpn_settings` supports the following:
* `authentication` - (Optional) authentication blocks are documented below. 
* `link_selection` - (Optional) Link selection blocks are documented below. 
* `maximum_concurrent_ike_negotiations` - (Optional) Maximum concurrent ike negotiations.
* `maximum_concurrent_tunnels` - (Optional) Maximum concurrent tunnels.
* `office_mode` - (Optional) Office Mode. Notation Wide Impact - Office Mode apply IPSec VPN Software Blade clients and to the Mobile Access Software Blade clients. office_mode blocks are documented below.
* `remote_access` - (Optional) remote_access blocks are documented below.
* `vpn_domain` - (Optional) Gateway VPN domain identified by the name.
* `vpn_domain_type` - (Optional) Gateway VPN domain type.

`authentication` supports the following:
* `authentication_clients` - (Optional) Collection of VPN Authentication clients identified by the name. 

`link_selection` supports the following:
* `ip_selection` - (Optional) IP selection. 
* `dns_resolving_hostname` - (Optional) DNS Resolving Hostname. Must be set when "ip-selection" was selected to be "dns-resolving-from-hostname". 
* `ip_address` - (Optional) IP Address. Must be set when "ip-selection" was selected to be "use-selected-address-from-topology" or "use-statically-nated-ip".

`office_mode` supports the following:
* `mode` - (Optional) Office Mode Permissions. When selected to be "off", all the other definitions are irrelevant.
* `group` - (Optional) Group. Identified by name. Must be set when "office-mode-permissions" was selected to be "group".
* `allocate_ip_address_from` - (Optional) Allocate IP address Method. Allocate IP address by sequentially trying the given methods until success. allocate_ip_address_from blocks are documented below.
* `support_multiple_interfaces` - (Optional) Support connectivity enhancement for gateways with multiple external interfaces.
* `perform_anti_spoofing` - (Optional) Perform Anti-Spoofing on Office Mode addresses.
* `anti_spoofing_additional_addresses` - (Optional) Additional IP Addresses for Anti-Spoofing. Identified by name. Must be set when "perform-anti-spoofings" is true.

`allocate_ip_address_from` supports the following:
* `radius_server` - (Optional) Radius server used to authenticate the user.
* `use_allocate_method` - (Optional) Use Allocate Method.
* `allocate_method` - (Optional) Using either Manual (IP Pool) or Automatic (DHCP). Must be set when "use-allocate-method" is true.
* `manual_network` - (Optional) Manual Network. Identified by name. Must be set when "allocate-method" was selected to be "manual".
* `dhcp_server` - (Optional) DHCP Server. Identified by name. Must be set when "allocate-method" was selected to be "automatic".
* `virtual_ip_address` - (Optional) Virtual IPV4 address for DHCP server replies. Must be set when "allocate-method" was selected to be "automatic".
* `dhcp_mac_address` - (Optional) Calculated MAC address for DHCP allocation. Must be set when "allocate-method" was selected to be "automatic".
* `optional_parameters` - (Optional) This configuration applies to all Office Mode methods except Automatic (using DHCP) and ipassignment.conf entries which contain this data. optional_parameters blocks are documented below.

`optional_parameters` supports the following:
* `use_primary_dns_server` - (Optional) Use Primary DNS Server.
* `primary_dns_server` - (Optional) Primary DNS Server. Identified by name. Must be set when "use-primary-dns-server" is true and can not be set when "use-primary-dns-server" is false.
* `use_first_backup_dns_server` - (Optional) Use First Backup DNS Server.
* `first_backup_dns_server` - (Optional) First Backup DNS Server. Identified by name. Must be set when "use-first-backup-dns-server" is true and can not be set when "use-first-backup-dns-server" is false.
* `use_second_backup_dns_server` - (Optional) Use Second Backup DNS Server.
* `second_backup_dns_server` - (Optional) Second Backup DNS Server. Identified by name. Must be set when "use-second-backup-dns-server" is true and can not be set when "use-second-backup-dns-server" is false.
* `dns_suffixes` - (Optional) DNS Suffixes.
* `use_primary_wins_server` - (Optional) Use Primary WINS Server.
* `primary_wins_server` - (Optional) Primary WINS Server. Identified by name. Must be set when "use-primary-wins-server" is true and can not be set when "use-primary-wins-server" is false.
* `use_first_backup_wins_server` - (Optional) Use First Backup WINS Server.
* `first_backup_wins_server` - (Optional) First Backup WINS Server. Identified by name. Must be set when "use-first-backup-wins-server" is true and can not be set when "use-first-backup-wins-server" is false.
* `use_second_backup_wins_server` - (Optional) Use Second Backup WINS Server.
* `second_backup_wins_server` - (Optional) Second Backup WINS Server. Identified by name. Must be set when "use-second-backup-wins-server" is true and can not be set when "use-second-backup-wins-server" is false.
* `ip_lease_duration` - (Optional) IP Lease Duration in Minutes. The value must be in the range 2-32767.

`remote_access` supports the following:
* `support_l2tp` - (Optional) Support L2TP (relevant only when office mode is active).
* `l2tp_auth_method` - (Optional) L2TP Authentication Method. Must be set when "support-l2tp" is true.
* `l2tp_certificate` - (Optional) L2TP Certificate. Must be set when "l2tp-auth-method" was selected to be "certificate". Insert "defaultCert" when you want to use the default certificate.
* `allow_vpn_clients_to_route_traffic` - (Optional) Allow VPN clients to route traffic.
* `support_nat_traversal_mechanism` - (Optional) Support NAT traversal mechanism (UDP encapsulation).
* `nat_traversal_service` - (Optional) Allocated NAT traversal UDP service. Identified by name. Must be set when "support-nat-traversal-mechanism" is true.
* `support_visitor_mode` - (Optional) Support Visitor Mode.
* `visitor_mode_service` - (Optional) TCP Service for Visitor Mode. Identified by name. Must be set when "support-visitor-mode" is true.
* `visitor_mode_interface` - (Optional) Interface for Visitor Mode. Must be set when "support-visitor-mode" is true. Insert IPV4 Address of existing interface or "All IPs" when you want all interfaces.

`logs_settings` supports the following:
* `alert_when_free_disk_space_below` - (Optional) Enable alert when free disk space is below threshold.
* `alert_when_free_disk_space_below_metrics` - (Optional) Free disk space metrics. 
* `alert_when_free_disk_space_below_threshold` - (Optional) Alert when free disk space below threshold. 
* `alert_when_free_disk_space_below_type` - (Optional) Alert when free disk space below type. 
* `before_delete_keep_logs_from_the_last_days` - (Optional) Enable before delete keep logs from the last days. 
* `before_delete_keep_logs_from_the_last_days_threshold` - (Optional) Before delete keep logs from the last days threshold. 
* `before_delete_run_script` - (Optional) Enable Before delete run script. 
* `before_delete_run_script_command` - (Optional) Before delete run script command. 
* `delete_index_files_older_than_days` - (Optional) Enable delete index files older than days. 
* `delete_index_files_older_than_days_threshold` - (Optional) Delete index files older than days threshold. 
* `delete_index_files_when_index_size_above` - (Optional) Enable delete index files when index size is above.
* `delete_index_files_when_index_size_above_threshold` - (Optional) Delete index files when index size is above threshold.
* `delete_when_free_disk_space_below` - (Optional) Enable delete when free disk space below. 
* `delete_when_free_disk_space_below_threshold` - (Optional) Delete when free disk space below threshold. 
* `detect_new_citrix_ica_application_names` - (Optional) Enable detect new citrix ica application names. 
* `enable_log_indexing` - (Optional) Enable log indexing. 
* `forward_logs_to_log_server` - (Optional) Enable forward logs to log server.
* `perform_log_rotate_before_log_forwarding` - (Optional) Enable perform log rotate before log forwarding.
* `reject_connections_when_free_disk_space_below_threshold` - (Optional) Enable reject connections when free disk space below threshold.
* `reserve_for_packet_capture_metrics` - (Optional) Reserve for packet capture metrics.
* `reserve_for_packet_capture_threshold` - (Optional) Reserve for packet capture threshold.
* `rotate_log_by_file_size` - (Optional) Enable rotate log by file size. 
* `rotate_log_file_size_threshold` - (Optional) Log file size threshold. 
* `rotate_log_on_schedule` - (Optional) Enable rotate log on schedule.
* `stop_logging_when_free_disk_space_below` - (Optional) Enable stop logging when free disk space below. 
* `stop_logging_when_free_disk_space_below_threshold` - (Optional) Stop logging when free disk space below threshold. 
* `turn_on_qos_logging` - (Optional) Enable turn on qos loggig. 
* `update_account_log_every` - (Optional) Update account log in every amount of seconds.
