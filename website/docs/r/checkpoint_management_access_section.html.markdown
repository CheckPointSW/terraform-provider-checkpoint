---
layout: "checkpoint"
page_title: "checkpoint_management_access_section"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-access-section"
description: |-
  This resource allows you to execute Check Point Access Section.
---

# Resource: checkpoint_management_access_section

This resource allows you to execute Check Point Access Section.

## Example Usage


```hcl
resource "checkpoint_management_access_section" "example" {
  name = "New Section 1"
  position = {top = "top"}
  layer = "Network"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `layer` - (Required) Layer that the rule belongs to identified by the name or UID. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
* `position` - (Required) Position in the rulebase. 

## Import

`checkpoint_management_access_section` can be imported by using the following format: LAYER_NAME;SECTION_UID

```
$ terraform import checkpoint_management_access_section.example "Network;354e184c-2f42-485c-b62d-ff9b3d29ee3e"
```