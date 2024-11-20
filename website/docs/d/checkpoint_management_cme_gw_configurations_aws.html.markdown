---
layout: "checkpoint"
page_title: "checkpoint_management_cme_gw_configurations_aws"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-cme-gw-configurations-aws"
description: |- Use this data source to get information on an existing Check Point CME AWS Gateway Configurations.
---

# Data Source: checkpoint_management_cme_gw_configurations_aws

Use this data source to get information on an existing Check Point CME AWS Gateway Configurations.

For details about the compatibility between the Terraform Release version and the CME API version, please refer to the section [Compatibility with CME](../index.html.markdown#compatibility-with-cme).


## Example Usage

```hcl
data "checkpoint_management_cme_gw_configurations_aws" "gw_config_aws" {
  name = "awsGWConfigurations"
}
```

## Argument Reference

These arguments are supported:

* `name` - (Required) The Gateway configuration name.
* `version` - The Gateway version.
* `sic_key` - SIC key for trusted communication between the Management and the Gateway.
* `policy` - The policy name to install on the Gateway.
* `related_account` - The related CME account name associated with the Gateway configuration.
* `blades` - Dictionary of activated/deactivated blades on the Gateway. Supports these blades:
    * `ips` - IPS blade.
    * `anti_bot` - Anti-Bot blade.
    * `anti_virus` - Anti-Virus blade.
    * `https_inspection` - HTTPS Inspection blade.
    * `application_control` - Application Control blade.
    * `autonomous_threat_prevention` - ATP blade.
    * `content_awareness` - Content Awareness blade.
    * `identity_awareness` - Identity Awareness blade.
    * `ipsec_vpn` - IPsec VPN blade.
    * `threat_emulation` - Threat Emulation blade.
    * `url_filtering` - URL Filtering blade.
    * `vpn` - VPN blade.
* `identity_awareness_settings` - Dictionary of Identity Awareness settings that can be configured on the gateway:
    * `enable_cloudguard_controller` - Enable the Web API identity source for CloudGuard Controller.
    * `receive_identities_from` - List of PDP gateway names from which to receive identities through Identity Sharing.
* `repository_gateway_scripts` - List of objects that each contain the name/UID of a script that exists in the scripts
  repository on the Management server. Supports these parameters:
    * `name` - The name of the script.
    * `uid` - The UID of the script.
    * `parameters` - Script parameters.
* `vpn_domain` - The group object is set as the VPN domain for the VPN Gateway.
* `vpn_community` - A star community to place the VPN Gateway as center.
* `deployment_type` - The deployment type of the CloudGuard Security Gateways.
* `tgw_static_routes` - Comma-separated list of CIDRs; for each CIDR, a static route is created on each Gateway of the
  Transit Gateway auto-scaling group.
* `tgw_spoke_routes` - Comma-separated list of spoke CIDRs; each spoke CIDR that is learned from the Transit Gateway
  over BGP is re-advertised by the Gateways of the Transit Gateway auto-scaling group to the AWS Transit Gateway.
* `send_logs_to_server` - Comma-separated list of Primary Log Servers names to which logs are sent.
* `send_logs_to_backup_server` - Comma-separated list of Backup Log Servers names to which logs are sent if the Primary
  Log Servers are unavailable.
* `send_alerts_to_server` - Comma-separated list of Alert Log Servers names to which alerts are sent.
* `section_name` - Name of a rule section in the Access and NAT layers in the policy, where to insert the automatically generated rules.
* `x_forwarded_for` - Enable XFF headers in HTTP / HTTPS requests.
* `color` - Color of the gateways objects in SmartConsole.
* `communication_with_servers_behind_nat` - Gateway behind NAT communications settings with the Check Point Servers(Management, Multi-Domain, Log Servers).