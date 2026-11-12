---
layout: "checkpoint"
page_title: "checkpoint_management_export_access_rulebase"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-export-access-rulebase"
description: |-
This resource allows you to execute Check Point Export Access Rulebase.
---

# checkpoint_management_export_access_rulebase

This resource allows you to execute Check Point Export Access Rulebase.

## Example Usage


```hcl
resource "checkpoint_management_export_access_rulebase" "example" {
  name = "Corp-Access"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. Must be unique in the domain.
* `uid`  - (Optional) Object unique identifier.
* `package` - (Optional) Name of the package. 
* `show_expiration_settings` - (Optional) Indicates whether to calculate and show "expiration date settings" field in reply. 
* `show_hits` - (Optional) Show hitcount data. 
* `use_object_dictionary` - (Optional) N/A 
* `hits_settings` - (Optional) Hitcount settings, define the range if hits to show.hits_settings blocks are documented below.
* `dereference_group_members` - (Optional) Indicates whether to dereference "members" field by details level for every object in reply. 
* `show_membership` - (Optional) Indicates whether to calculate and show "groups" field for every object in reply. 


`hits_settings` supports the following:

* `from_date` - (Optional) Format: YYYY-MM-DD, YYYY-mm-ddThh:mm:ss. 
* `target` - (Optional) Target gateway name or UID. 
* `to_date` - (Optional) Format: YYYY-MM-DD, YYYY-mm-ddThh:mm:ss.
