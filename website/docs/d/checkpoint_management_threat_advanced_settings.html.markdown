---
layout: "checkpoint"
page_title: "checkpoint_management_threat_advanced_settings"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-threat-advanced-settings"
description: |-
Use this data source to get information on an existing Check Point Threat Advanced Settings.
---

# Data Source: checkpoint_management_threat_advanced_settings

Use this data source to get information on an existing Check Point Threat Advanced Settings.

## Example Usage


```hcl
data "checkpoint_management_threat_advanced_settings" "data_threat_advanced_settings" {

}
```

## Argument Reference

The following arguments are supported:

* `uid` - Object unique identifier.
* `feed_retrieving_interval` - Feed retrieving intervals of External Feed, in the form of HH:MM.
* `httpi_non_standard_ports` - Enable HTTP Inspection on non standard ports for Threat Prevention blades.
* `internal_error_fail_mode` - In case of internal system error, allow or block all connections.
* `log_unification_timeout` - Session unification timeout for logs (minutes).
* `resource_classification` - Allow (Background) or Block (Hold) requests until categorization is complete. resource_classification blocks are documented below.


`resource_classification` supports the following:

* `custom_settings` - Custom resources classification per service. custom_settings blocks are documented below.
* `mode` - Set all services to the same mode or choose a custom mode.
* `web_service_fail_mode` - Block connections when the web service is unavailable.


`custom_settings` supports the following:

* `anti_bot` - Custom Settings for Anti Bot Blade.
* `anti_virus` - Custom Settings for Anti Virus Blade.
* `zero_phishing` - Custom Settings for Zero Phishing Blade.