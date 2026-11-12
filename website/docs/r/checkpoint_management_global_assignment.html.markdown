---
layout: "checkpoint"
page_title: "checkpoint_management_global_assignment"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-global-assignment"
description: |-
This resource allows you to execute Check Point Global Assignment.
---

# Resource: checkpoint_management_global_assignment

This resource allows you to execute Check Point Global Assignment.

## Example Usage


```hcl
resource "checkpoint_management_global_assignment" "example" {
  global_domain = "Global"
  dependent_domain = "domain2"
  global_access_policy = "standard"
  global_threat_prevention_policy = "standard"
  manage_protection_actions = true
}
```

## Argument Reference

The following arguments are supported:

* `dependent_domain` - (Optional) N/A 
* `global_access_policy` - (Optional) Global domain access policy that is assigned to a dependent domain. 
* `global_domain` - (Optional) N/A 
* `global_threat_prevention_policy` - (Optional) Global domain threat prevention policy that is assigned to a dependent domain. 
* `manage_protection_actions` - (Optional) N/A 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
* `assignment_status`
* `assignment_up_to_date` - The time when the assignment was assigned. assignment_up_to_date blocks are documented below.


`assignment_up_to_date` supports the follwoing:

* `iso_8601` - Date and time represented in international ISO 8601 format.
* `posix` - Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.