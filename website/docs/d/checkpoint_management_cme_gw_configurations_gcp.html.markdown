---
layout: "checkpoint"
page_title: "checkpoint_management_cme_gw_configurations_gcp"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-cme-gw-configurations-gcp"
description: |- Use this data source to get information on an existing Check Point CME GCP GW Configurations.
---

# Data Source: checkpoint_management_cme_gw_configurations_gcp

Use this data source to get information on an existing Check Point CME GCP GW Configurations.

Available in:

- Check Point Security Management/Multi Domain Management Server R81.10 and higher.
- CME take 255 and higher.

## Example Usage

```hcl
data "checkpoint_management_cme_gw_configurations_gcp" "gw_config_gcp" {
  name = "gcpGWConfigurations"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The GW configuration name.
* `version` - The GW version.
* `sic_key` - SIC key for trusted communication between management and GW.
* `policy` - Policy name to be installed on the GW.
* `related_account` - Related CME account name associated with the GW Configuration.
* `blades` - Dictionary of activated/deactivated blades on the GW. Supports the following:
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
* `repository_gateway_scripts` - List of objects that each contains name/UID of a script that exists in the scripts
  repository on the Management server. Supports the following:
    * `name` - The name of the script.
    * `uid` - The UID of the script.
    * `parameters` - Script parameters.
* `send_logs_to_server` - Comma separated list of Primary Log Servers names to which logs are sent.
* `send_logs_to_backup_server` - Comma separated list of Backup Log Servers names to which logs are sent in case Primary
  Log Servers are unavailable.
* `send_alerts_to_server` - Comma separated list of Alert Log Servers names to which alerts are sent.