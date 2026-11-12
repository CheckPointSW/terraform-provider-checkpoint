---
layout: "checkpoint"
page_title: "checkpoint_management_def_setting"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-def-setting"
description: |-
Use this data source to get information on an existing Check Point Def Setting.
---

# Data Source: checkpoint_management_def_setting

Use this data source to get information on an existing Check Point Def Setting.

## Example Usage

```hcl
data "checkpoint_management_def_setting" "data_test" {
    name = "def-setting1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.
* `assignments` - (Computed) Assignments. assignments blocks are documented below.
* `custom` - (Computed) Whether this is a user-created custom setting or a predefined setting.
* `data_type` - (Computed) The data type of the setting.
* `global` - (Computed) Whether this setting applies globally to all gateways.
* `color` - (Computed) Color of the object. Should be one of existing colors.
* `comments` - (Computed) Comments string.
* `icon` - (Computed) Object icon.
* `tags` - (Computed) Collection of tag objects identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.


`assignments` supports the following:

* `description` - (Computed) The description for this setting.
* `enabled` - (Computed) If the setting is enabled.
* `from_version` - (Computed) The gateway version this setting applies from.
* `model` - (Computed) The gateway model this setting applies to.
* `position` - (Computed) The position of the setting.
* `targets` - (Computed) Collection of Gateways identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.
* `to_version` - (Computed) The gateway version this setting applies to.
* `value` - (Computed) The value of the setting.
