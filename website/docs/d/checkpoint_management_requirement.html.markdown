---
layout: "checkpoint"
page_title: "checkpoint_management_requirement"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-requirement"
description: |-
Use this data source to get information on an existing Check Point Requirement.
---

# Data Source: checkpoint_management_requirement

Use this data source to get information on an existing Check Point Requirement.

## Example Usage

```hcl
data "checkpoint_management_requirement" "data_test" {
    name = "requirement1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Requirement name.
* `uid` - (Optional) Object unique identifier.
* `regulation` - (Optional) The relevant regulation of the requirement, identified by name.
* `best_practices` - (Computed) The list of the best practices that make up the requirement. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.
* `score` - (Computed) The score of the requirement.
* `score_level` - (Computed) The score level of the requirement.
* `user_defined` - (Computed) Shows if the requirement is user defined.
* `color` - (Computed) Color of the object. Should be one of existing colors.
* `comments` - (Computed) The requirement comments.
* `icon` - (Computed) Object icon.
* `tags` - (Computed) Collection of tag objects identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.
