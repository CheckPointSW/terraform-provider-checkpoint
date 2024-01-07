---
layout: "checkpoint"
page_title: "checkpoint_management_cme_gw_configurations_azure"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-cme-gw-configurations-azure"
description: |- This resource allows you to add/update/delete Check Point CME Azure GW Configurations.
---

# checkpoint_management_cme_gw_configurations_azure

This resource allows you to add/update/delete Check Point CME Azure GW Configurations.

## Example Usage

```hcl
resource "checkpoint_management_cme_gw_configurations_azure" "gw_config_azure" {
  name                       = "TestAzureTemplate"
  related_account            = "TestAzureController"
  version                    = "R81"
  base64_sic_key             = "MTIzNDU2Nzg="
  policy                     = "Standard"
  send_logs_to_server        = ["PLS_A"]
  send_logs_to_backup_server = ["BLS_B"]
  send_alerts_to_server      = ["ALS_C"]
  repository_gateway_scripts {
    name = "xyz"
    parameters = "param1 param2"
  }
  blades {
    ips                          = true
    anti_bot                     = true
    anti_virus                   = true
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

The following arguments are supported:

* `name` - (Required) The GW configuration name.
* `version` - (Required) The GW version.
* `base64_sic_key` - (Required) Base64 key for trusted communication between management and GW.
* `policy` - (Required) Policy name to be installed on the GW.
* `related_account` - (Required) The CME account to associate with the GW Configuration.
* `blades` - (Required) Dictionary of activated/deactivated blades on the GW. supports the following:
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
* `repository_gateway_scripts` - (Optional) List of objects that each contains name/UID of a script that exists in the scripts repository on the Management server. Supports the following:
    * `name` - (Required) The name of the script.
    * `parameters` - (Optional) The parameters to pass to the script.
* `send_logs_to_server` - (Optional) Comma separated list of Primary Log Servers names to which logs are sent. Defined Log Server will act as Log and Alert Servers. Must be defined as part of Log Servers parameters.
* `send_logs_to_backup_server` - (Optional) Comma separated list of Backup Log Servers names to which logs are sent in case Primary Log Servers are unavailable.
* `send_alerts_to_server` - (Optional) Comma separated list of Alert Log Servers names to which alerts are sent.

## How To Use

Make sure this resource creation will be executed in the right execution order. note: terraform execution is not sequential.  
For example, if you want to create a gateway configuration along with an account, the account should be created first.
For forcing this dependency, the "related_account" argument should use the account resource like in the following example:

```hcl
resource "checkpoint_management_cme_accounts_azure" "azure_account" {
   ///...
}
resource "checkpoint_management_cme_gw_configurations_azure" "gw_config_azure" {
  name                       = "TestAzureTemplate"
  related_account            = checkpoint_management_cme_accounts_azure.azure_account.name
  ///...
}
```
