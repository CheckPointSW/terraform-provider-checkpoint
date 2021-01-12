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
* `ipv4_address` - IPv4 address. 
* `ipv6_address` - IPv6 address.
* `interfaces` - Gateway interfaces. interfaces blocks are documented below.
* `anti_bot` - Anti-Bot blade enabled. 
* `anti_virus` - Anti-Virus blade enabled. 
* `application_control` - Application Control blade enabled.
* `content_awareness` - Content Awareness blade enabled.
* `icap_server` - ICAP Server enabled.
* `ips` - Intrusion Prevention System blade enabled.
* `threat_emulation` - Threat Emulation blade enabled.
* `threat_extraction` - Threat Extraction blade enabled.
* `url_filtering` - URL Filtering blade enabled.
* `firewall` - Firewall blade enabled.
* `firewall_settings` - Firewall settings. firewall_settings blocks are documented below.
* `vpn` - VPN blade enabled.
* `vpn_settings` - Gateway VPN settings. vpn_settings blocks are documented below.
* `dynamic_ip` - Dynamic IP address.
* `version` - Gateway platform version.
* `os_name` - Operating system name.
* `hardware` - Gateway platform hardware name.
* `one_time_password` - Secure internal connection one time password. 
* `sic_name` - Secure Internal Communication name.
* `sic_state` - Secure Internal Communication state.
* `save_logs_locally` - Enable save logs locally.
* `send_alerts_to_server` - Collection of Server(s) to send alerts to identified by the name.
* `send_logs_to_backup_server` - Collection of Backup server(s) to send logs to identified by the name.
* `send_logs_to_server` - Collection of Server(s) to send logs to identified by the name.
* `logs_settings` - Logs settings. logs_settings blocks are documented below.
* `color` - Color of the object.
* `comments` - Comments string. 
* `tags` - Collection of tags identified by name.

`interfaces` supports the following:
* `name` - Interface name. 
* `ipv4_address` - IPv4 address. 
* `ipv6_address` - IPv6 address. 
* `ipv4_network_mask` - IPv4 network address.
* `ipv6_network_mask` - IPv6 network address.
* `ipv4_mask_length` - IPv4 network mask length.
* `ipv6_mask_length` - IPv6 network mask length.
* `anti_spoofing` - Anti spoofing.
* `anti_spoofing_settings` - Anti spoofing settings. anti_spoofing_settings blocks are documented below.
* `security_zone` - Security zone.
* `security_zone_settings` - Security zone settings. security_zone_settings blocks are documented below.
* `topology` - Topology.
* `topology_settings` - Topology settings. topology_settings blocks are documented below.
* `topology_automatic_calculation` - Shows the automatic topology calculation..
* `color` - Color of the object. Should be one of existing colors. 
* `comments` - Comments string. 

`anti_spoofing_settings` supports the following:
* `action` - If packets will be rejected (the Prevent option) or whether the packets will be monitored (the Detect option).

`security_zone_settings` supports the following:
* `auto_calculated` - Security Zone is calculated according to where the interface leads to.
* `specific_zone` - Security Zone specified manually.

`topology_settings` supports the following:
* `interface_leads_to_dmz` - Whether this interface leads to demilitarized zone (perimeter network).
* `ip_address_behind_this_interface` - Ip address behind this interface.
* `specific_network` - Network behind this interface.

`firewall_settings` supports the following:
* `auto_calculate_connections_hash_table_size_and_memory_pool` - Auto calculate connections hash table size and memory pool. 
* `auto_maximum_limit_for_concurrent_connections` - Auto maximum limit for concurrent connections.
* `connections_hash_size` - Connections hash size.
* `maximum_limit_for_concurrent_connections` - Maximum limit for concurrent connections.
* `maximum_memory_pool_size` - Maximum memory pool size.
* `memory_pool_size` - Memory pool size.

`vpn_settings` supports the following:
* `authentication` - authentication blocks are documented below. 
* `link_selection` - Link selection blocks are documented below. 
* `maximum_concurrent_ike_negotiations` - Maximum concurrent ike negotiations.
* `maximum_concurrent_tunnels` - Maximum concurrent tunnels.
* `office_mode` - Office Mode. Notation Wide Impact - Office Mode apply IPSec VPN Software Blade clients and to the Mobile Access Software Blade clients. office_mode blocks are documented below.
* `remote_access` - remote_access blocks are documented below.
* `vpn_domain` - Gateway VPN domain identified by the name.
* `vpn_domain_type` - Gateway VPN domain type.

`authentication` supports the following:
* `authentication_clients` - Collection of VPN Authentication clients identified by the name. 

`link_selection` supports the following:
* `ip_selection` - IP selection. 
* `dns_resolving_hostname` - DNS Resolving Hostname. Must be set when "ip-selection" was selected to be "dns-resolving-from-hostname". 
* `ip_address` - IP Address. Must be set when "ip-selection" was selected to be "use-selected-address-from-topology" or "use-statically-nated-ip".

`office_mode` supports the following:
* `mode` - Office Mode Permissions. When selected to be "off", all the other definitions are irrelevant.
* `group` - Group. Identified by name. Must be set when "office-mode-permissions" was selected to be "group".
* `allocate_ip_address_from` - Allocate IP address Method. Allocate IP address by sequentially trying the given methods until success. allocate_ip_address_from blocks are documented below.
* `support_multiple_interfaces` - Support connectivity enhancement for gateways with multiple external interfaces.
* `perform_anti_spoofing` - Perform Anti-Spoofing on Office Mode addresses.
* `anti_spoofing_additional_addresses` - Additional IP Addresses for Anti-Spoofing. Identified by name. Must be set when "perform-anti-spoofings" is true.

`allocate_ip_address_from` supports the following:
* `radius_server` - Radius server used to authenticate the user.
* `use_allocate_method` - Use Allocate Method.
* `allocate_method` - Using either Manual (IP Pool) or Automatic (DHCP). Must be set when "use-allocate-method" is true.
* `manual_network` - Manual Network. Identified by name. Must be set when "allocate-method" was selected to be "manual".
* `dhcp_server` - DHCP Server. Identified by name. Must be set when "allocate-method" was selected to be "automatic".
* `virtual_ip_address` - Virtual IPV4 address for DHCP server replies. Must be set when "allocate-method" was selected to be "automatic".
* `dhcp_mac_address` - Calculated MAC address for DHCP allocation. Must be set when "allocate-method" was selected to be "automatic".
* `optional_parameters` - This configuration applies to all Office Mode methods except Automatic (using DHCP) and ipassignment.conf entries which contain this data. optional_parameters blocks are documented below.

`optional_parameters` supports the following:
* `use_primary_dns_server` - Use Primary DNS Server.
* `primary_dns_server` - Primary DNS Server. Identified by name. Must be set when "use-primary-dns-server" is true and can not be set when "use-primary-dns-server" is false.
* `use_first_backup_dns_server` - Use First Backup DNS Server.
* `first_backup_dns_server` - First Backup DNS Server. Identified by name. Must be set when "use-first-backup-dns-server" is true and can not be set when "use-first-backup-dns-server" is false.
* `use_second_backup_dns_server` - Use Second Backup DNS Server.
* `second_backup_dns_server` - Second Backup DNS Server. Identified by name. Must be set when "use-second-backup-dns-server" is true and can not be set when "use-second-backup-dns-server" is false.
* `dns_suffixes` - DNS Suffixes.
* `use_primary_wins_server` - Use Primary WINS Server.
* `primary_wins_server` - Primary WINS Server. Identified by name. Must be set when "use-primary-wins-server" is true and can not be set when "use-primary-wins-server" is false.
* `use_first_backup_wins_server` - Use First Backup WINS Server.
* `first_backup_wins_server` - First Backup WINS Server. Identified by name. Must be set when "use-first-backup-wins-server" is true and can not be set when "use-first-backup-wins-server" is false.
* `use_second_backup_wins_server` - Use Second Backup WINS Server.
* `second_backup_wins_server` - Second Backup WINS Server. Identified by name. Must be set when "use-second-backup-wins-server" is true and can not be set when "use-second-backup-wins-server" is false.
* `ip_lease_duration` - IP Lease Duration in Minutes. The value must be in the range 2-32767.

`remote_access` supports the following:
* `support_l2tp` - Support L2TP (relevant only when office mode is active).
* `l2tp_auth_method` - L2TP Authentication Method. Must be set when "support-l2tp" is true.
* `l2tp_certificate` - L2TP Certificate. Must be set when "l2tp-auth-method" was selected to be "certificate". Insert "defaultCert" when you want to use the default certificate.
* `allow_vpn_clients_to_route_traffic` - Allow VPN clients to route traffic.
* `support_nat_traversal_mechanism` - Support NAT traversal mechanism (UDP encapsulation).
* `nat_traversal_service` - Allocated NAT traversal UDP service. Identified by name. Must be set when "support-nat-traversal-mechanism" is true.
* `support_visitor_mode` - Support Visitor Mode.
* `visitor_mode_service` - TCP Service for Visitor Mode. Identified by name. Must be set when "support-visitor-mode" is true.
* `visitor_mode_interface` - Interface for Visitor Mode. Must be set when "support-visitor-mode" is true. Insert IPV4 Address of existing interface or "All IPs" when you want all interfaces.

`logs_settings` supports the following:
* `free_disk_space_metrics` - Free disk space metrics. 
* `accept_syslog_messages` - Enable accept syslog messages. 
* `alert_when_free_disk_space_below` - Enable alert when free disk space is below threshold. 
* `alert_when_free_disk_space_below_threshold` - Alert when free disk space below threshold. 
* `alert_when_free_disk_space_below_type` - Alert when free disk space below type. 
* `before_delete_keep_logs_from_the_last_days` - Enable before delete keep logs from the last days. 
* `before_delete_keep_logs_from_the_last_days_threshold` - Before delete keep logs from the last days threshold. 
* `before_delete_run_script` - Enable Before delete run script. 
* `before_delete_run_script_command` - Before delete run script command. 
* `delete_index_files_older_than_days` - Enable delete index files older than days. 
* `delete_index_files_older_than_days_threshold` - Delete index files older than days threshold. 
* `delete_when_free_disk_space_below` - Enable delete when free disk space below. 
* `delete_when_free_disk_space_below_threshold` - Delete when free disk space below threshold. 
* `detect_new_citrix_ica_application_names` - Enable detect new citrix ica application names. 
* `enable_log_indexing` - Enable log indexing. 
* `forward_logs_to_log_server` - Enable forward logs to log server. 
* `forward_logs_to_log_server_name` - Forward logs to log server name. 
* `forward_logs_to_log_server_schedule_name` - Forward logs to log server schedule name. 
* `rotate_log_by_file_size` - Enable rotate log by file size. 
* `rotate_log_file_size_threshold` - Log file size threshold. 
* `rotate_log_on_schedule` - Enable rotate log on schedule. 
* `rotate_log_schedule_name` - Rotate log schedule name. 
* `stop_logging_when_free_disk_space_below` - Enable stop logging when free disk space below. 
* `stop_logging_when_free_disk_space_below_threshold` - Stop logging when free disk space below threshold. 
* `turn_on_qos_logging` - Enable turn on qos logging. 
* `update_account_log_every` - Update account log in every amount of seconds.