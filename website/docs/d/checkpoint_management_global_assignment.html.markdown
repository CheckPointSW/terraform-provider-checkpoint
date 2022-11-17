---
layout: "checkpoint"
page_title: "checkpoint_management_global_assignment"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-global-assignment"
description: |-
Use this data source to get information on an existing Check Point Global Assignment.
---

# Data Source: checkpoint_management_global_assignment

Use this data source to get information on an existing Check Point Global Assignment.

## Example Usage


```hcl
resource "checkpoint_management_global_assignment" "global_assignment" {
  global_domain = "Global"
  dependent_domain = "domain2"
  global_access_policy = "standard"
  global_threat_prevention_policy = "standard"
  manage_protection_actions = true
}

data "checkpoint_management_global_assignment" "data_global_assignment" {
  dependent_domain = "${checkpoint_management_global_assignment.global_assignment.dependent_domain}"
  global_domain = "${checkpoint_management_global_assignment.global_assignment.global_domain}"
}
```

## Argument Reference

The following arguments are supported:

* `uid` - (Optional) Object unique identifier.
* `dependent_domain` - (Optional)
* `global_domain` - (Optional)
* `assignment_status`
* `assignment_up_to_date` - The time when the assignment was assigned. assignment_up_to_date blocks are documented below.
* `global_access_policy` - Global domain access policy that is assigned to a dependent domain.
* `global_threat_prevention_policy` - Global domain threat prevention policy that is assigned to a dependent domain.
* `manage_protection_actions`
* `tags` - Collection of tag objects identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.