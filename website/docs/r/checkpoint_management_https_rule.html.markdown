---
layout: "checkpoint"
page_title: "checkpoint_management_https_rule"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-https-rule"
description: |-
This resource allows you to execute Check Point Https Rule.
---

# checkpoint_management_https_rule

This resource allows you to execute Check Point Https Rule.

## Example Usage


```hcl
resource "checkpoint_management_https_rule" "example" {
  name = "FirstRule"
  position = {top = "top"}
  layer = "MyLayer"
}
```

## Argument Reference

The following arguments are supported:

* `rule_number` - (Required) Rule number. 
* `layer` - (Required) Layer that holds the Object. Identified by the Name or UID. 
* `name` - (Optional) HTTPS rule name. 
* `destination` - (Optional) Collection of Network objects identified by Name or UID that represents connection destination.destination blocks are documented below.
* `service` - (Optional) Collection of Network objects identified by Name or UID that represents connection service.service blocks are documented below.
* `source` - (Optional) Collection of Network objects identified by Name or UID that represents connection source.source blocks are documented below.
* `action` - (Optional) Rule inspect level. "Bypass" or "Inspect". 
* `blade` - (Optional) Blades for HTTPS Inspection. Identified by Name or UID to enable the inspection for.
"Anti Bot","Anti Virus","Application Control","Data Awareness","DLP","IPS","Threat Emulation","Url Filtering".blade blocks are documented below.
* `certificate` - (Optional) Internal Server Certificate identified by Name or UID,
otherwise, "Outbound Certificate" is a default value. 
* `destination_negate` - (Optional) TRUE if "negate" value is set for Destination. 
* `enabled` - (Optional) Enable/Disable the rule. 
* `install_on` - (Optional) Which Gateways identified by the name or UID to install the policy on.install_on blocks are documented below.
* `service_negate` - (Optional) TRUE if "negate" value is set for Service. 
* `site_category` - (Optional) Collection of Site Categories objects identified by the name or UID.site_category blocks are documented below.
* `site_category_negate` - (Optional) TRUE if "negate" value is set for Site Category. 
* `source_negate` - (Optional) TRUE if "negate" value is set for Source. 
* `track` - (Optional) "None","Log","Alert","Mail","SNMP trap","Mail","User Alert", "User Alert 2", "User Alert 3". 
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
* `position` - (Required) Position in the rulebase. 
