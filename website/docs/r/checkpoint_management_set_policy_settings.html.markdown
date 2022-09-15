---
layout: "checkpoint"
page_title: "checkpoint_management_command_set_policy_settings"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-command-set-policy-settings"
description: |-
This resource allows you to execute Check Point Set Policy Settings.
---

# Resource: checkpoint_management_command_set_policy_settings

This resource allows you to execute Check Point Set Policy Settings.

## Example Usage


```hcl
resource "checkpoint_management_command_set_policy_settings" "example" {
  last_in_cell = "none"
  none_object_behavior = "warning"
}
```

## Argument Reference

The following arguments are supported:

* `last_in_cell` - (Optional) Added object after removing the last object in cell. 
* `none_object_behavior` - (Optional) 'None' object behavior. Rules with object 'None' will never be matched. 
* `security_access_defaults` - (Optional) Access Policy default values. security_access_defaults blocks are documented below.


`security_access_defaults` supports the following:

* `destination` - (Optional) Destination default value for new rule creation. Any or None. 
* `service` - (Optional) Service and Applications default value for new rule creation. Any or None. 
* `source` - (Optional) Source default value for new rule creation. Any or None. 



