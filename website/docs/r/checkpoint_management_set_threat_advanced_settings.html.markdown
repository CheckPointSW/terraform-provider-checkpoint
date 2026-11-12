---
layout: "checkpoint"
page_title: "checkpoint_management_command_set_threat_advanced_settings"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-command-set-threat-advanced-settings"
description: |-
This resource allows you to execute Check Point Set Threat Advanced Settings.
---

# Resource: checkpoint_management_command_set_threat_advanced_settings

This resource allows you to execute Check Point Set Threat Advanced Settings.

## Example Usage


```hcl
resource "checkpoint_management_command_set_threat_advanced_settings" "example" {
  internal_error_fail_mode = "allow connections"
  log_unification_timeout = 600
  feed_retrieving_interval = "00:05"
  httpi_non_standard_ports = true
}
```

## Argument Reference

The following arguments are supported:

* `feed_retrieving_interval` - (Optional) Feed retrieving intervals of External Feed, in the form of HH:MM. 
* `httpi_non_standard_ports` - (Optional) Enable HTTP Inspection on non standard ports for Threat Prevention blades. 
* `internal_error_fail_mode` - (Optional) In case of internal system error, allow or block all connections. 
* `log_unification_timeout` - (Optional) Session unification timeout for logs (minutes). 
* `resource_classification` - (Optional) Allow (Background) or Block (Hold) requests until categorization is complete.resource_classification blocks are documented below. resource_classification is type list.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


`resource_classification` supports the following:

* `custom_settings` - (Optional) On Custom mode, custom resources classification per service.custom_settings blocks are documented below. custom_settings is type list.
* `mode` - (Optional) Set all services to the same mode or choose a custom mode. 
* `web_service_fail_mode` - (Optional) Block connections when the web service is unavailable. 


`custom_settings` supports the following:

* `anti_bot` - (Optional) Custom Settings for Anti Bot Blade. 
* `anti_virus` - (Optional) Custom Settings for Anti Virus Blade. 
* `zero_phishing` - (Optional) Custom Settings for Zero Phishing Blade. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

