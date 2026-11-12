---
layout: "checkpoint"
page_title: "checkpoint_management_regulation"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-regulation"
description: |-
Use this data source to get information on an existing Check Point Regulation.
---

# Data Source: checkpoint_management_regulation

Use this data source to get information on an existing Check Point Regulation.

## Example Usage

```hcl
data "checkpoint_management_regulation" "data_test" {
    name = "regulation1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Regulation name.
* `uid` - (Computed) Object unique identifier.
* `full_name` - (Optional) Regulation full name.
* `show_requirements` - (Optional) Show the requirements of the regulation.
* `enabled` - (Computed) Shows if the regulation is enabled.
* `score` - (Computed) The regulation score.
* `user_defined` - (Computed) Shows if the regulation is user defined.
* `comments` - (Computed) The Compliance Regulation comments.
* `color` - (Computed) Color of the object. Should be one of existing colors.
* `icon` - (Computed) Object icon.
* `tags` - (Computed) Collection of tag objects identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.
* `requirements` - (Computed) The requirements of the regulation, identified by name. Appears only when the value of the 'show-requirements' parameter is set to 'true'.
