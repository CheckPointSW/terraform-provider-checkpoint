---
layout: "checkpoint"
page_title: "checkpoint_management_cme_gw_configurations_aws"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-cme-gw-configurations-aws"
description: |- This resource allows you to add/update/delete Check Point CME AWS Gateway Configurations.
---

# Resource: checkpoint_management_cme_gw_configurations_aws

This resource allows you to add/update/delete Check Point CME AWS Gateway Configurations.

Available in:

- Check Point Security Management/Multi-Domain Security Management Server R81.10 and higher.
- CME Take 255 and higher.

## Example Usage

```hcl
resource "checkpoint_management_cme_gw_configurations_aws" "gw_config_aws" {
  name                       = "awsGWConfigurations"
  related_account            = "awsAccount"
  version                    = "R81"
  base64_sic_key             = "MTIzNDU2Nzg="
  policy                     = "Standard"
  tgw_spoke_routes           = ["192.168.100.0/24", "192.168.200.0/24"]
  tgw_static_routes          = ["10.0.0.0/16", "10.100.0.0/16"]
  send_logs_to_server        = ["PLS_A"]
  send_logs_to_backup_server = ["BLS_B"]
  send_alerts_to_server      = ["ALS_C"]
  section_name               = "my_section"
  x_forwarded_for            = true
  color                      = "blue"
  communication_with_servers_behind_nat = "translated-ip-only"
  repository_gateway_scripts {
    name       = "myScript"
    parameters = "ls -l"
  }
  blades {
    ips                          = true
    anti_bot                     = false
    anti_virus                   = false
    https_inspection             = false
    application_control          = false
    autonomous_threat_prevention = false
    content_awareness            = false
    identity_awareness           = false
    ipsec_vpn                    = false
    threat_emulation             = false
    url_filtering                = false
    vpn                          = false
  }
}
```

## Argument Reference

These arguments are supported:

* `name` - (Required) The Gateway configuration name without spaces.
* `version` - (Required) The Gateway version.
* `base64_sic_key` - (Required) Base64 key for trusted communication between the Management and the Gateway.
* `policy` - (Required) The policy name to install on the Gateway.
* `related_account` - (Required) The CME account name to associate with the Gateway Configuration.
* `blades` - (Required) Dictionary of activated/deactivated blades on the Gateway. Supports these blades:
    * `ips` - (Required) IPS blade.
    * `anti_bot` - (Required) Anti-Bot blade.
    * `anti_virus` - (Required) Anti-Virus blade.
    * `https_inspection` - (Required) HTTPS Inspection blade.
    * `application_control` - (Required) Application Control blade.
    * `autonomous_threat_prevention` - (Required) ATP blade.
    * `content_awareness` - (Required) Content Awareness blade.
    * `identity_awareness` - (Required) Identity Awareness blade.
    * `ipsec_vpn` - (Required) IPsec VPN blade.
    * `threat_emulation` - (Required) Threat Emulation blade.
    * `url_filtering` - (Required) URL Filtering blade.
    * `vpn` - (Required) VPN blade.
* `repository_gateway_scripts` - (Optional) List of objects that each contain the name/UID of a script that exists in
  the scripts repository on the Management server. Supports these parameters:
    * `name` - (Required) The name of the script.
    * `parameters` - (Optional) The parameters to pass to the script.
* `vpn_domain` - (Optional) The group object to set as the VPN domain for the VPN Gateway.
  An empty string automatically sets an empty group as the encryption domain. Must be an empty string for 'TGW'
  deployment type.
* `vpn_community` - (Optional) A star community to place the VPN Gateway as center.
* `deployment_type` - (Optional) The deployment type of the CloudGuard Security Gateways.
* `tgw_static_routes` - (Optional) Comma-separated list of CIDRs; for each CIDR a static route is created on each
  Gateway of the Transit Gateway auto-scaling group.
* `tgw_spoke_routes` - (Optional) Comma separated list of spoke CIDRs; each spoke CIDR that is learned from the Transit
  Gateway over BGP is re-advertised by the Gateways of the Transit Gateway auto-scaling group to the AWS Transit
  Gateway.
* `send_logs_to_server` - (Optional) Comma-separated list of Primary Log Servers names to which logs are sent.
  Configured Log Servers act as Log and Alert Servers. Must be defined as a part of Log Servers parameters.
* `send_logs_to_backup_server` - (Optional) Comma-separated list of Backup Log Servers names to which logs are sent if
  the Primary Log Servers are unavailable.
* `send_alerts_to_server` - (Optional) Comma-separated list of Alert Log Servers names to which alerts are sent.
* `section_name` - (Optional) Name of a rule section in the Access and NAT layers in the policy, where to insert the automatically generated rules.
* `x_forwarded_for` - (Optional) Enable XFF headers in HTTP / HTTPS requests.
* `color` - (Optional) Color of the gateways objects in SmartConsole.
* `communication_with_servers_behind_nat` - (Optional) Gateway behind NAT communications settings with the Check Point Servers(Management, Multi-Domain, Log Servers).


