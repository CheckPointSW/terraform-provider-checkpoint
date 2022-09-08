---
layout: "checkpoint"
page_title: "checkpoint_management_policy_settings"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-policy-settings"
description: |-
Use this data source to get information on an existing Check Point Policy Settings.
---

# Data Source: checkpoint_management_policy_settings

Use this data source to get information on an existing Check Point Policy Settings.

## Example Usage

```hcl
data "checkpoint_management_policy_settings" "data_policy_settings" {

}
```

## Argument Reference

The following arguments are supported:

* `last_in_cell` - Added object after removing the last object in cell.
* `none_object_behavior` - 'None' object behavior. Rules with object 'None' will never be matched.
* `security_access_defaults` - Access Policy default values. security_accesses_defaults blocks are documented below.


`security_access_defaults` supports the following:

* `destination` - Destination default value identified by name.
* `service` - Service and Applications default value identified by name.
* `source` - Source default value identified by name.