---
layout: "checkpoint"
page_title: "checkpoint_management_def_setting"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-def-setting"
description: |-
This resource allows you to execute Check Point Def Setting.
---

# checkpoint_management_def_setting

This resource allows you to execute Check Point Def Setting.

## Example Usage


```hcl
resource "checkpoint_management_def_setting" "example" {
  name      = "My Boolean Def Setting"
  data_type = "boolean"
  assignments {
    value       = "true"
    description = "Default for Quantum gateways"
    model       = "quantum"
  }
  assignments {
    value       = "false"
    description = "Default for Spark gateways"
    model       = "spark"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `data_type` - (Required) The data type of the setting. 
* `assignments` - (Required) Assignments.assignments blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `tags` - (Optional) Collection of tag identifiers.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


`assignments` supports the following:

* `value` - (Required) The value of the setting. 
* `description` - (Optional) The description for this setting. 
* `enabled` - (Optional) If the setting is enabled. 
* `from_version` - (Optional) The gateway version this setting applies from. 
* `model` - (Optional) The gateway model this setting applies to. 
* `position` - (Optional) The position of the setting. 
* `targets` - (Optional) The Gateways or Clusters the assignment is applied to, identified by name or UID.
* `to_version` - (Optional) The gateway version this setting applies to. 
