---
layout: "checkpoint"
page_title: "checkpoint_management_application_site_category"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-application-site-category"
description: |-
This resource allows you to execute Check Point Application Site Category.
---

# checkpoint_management_application_site_category

This resource allows you to execute Check Point Application Site Category.

## Example Usage


```hcl
resource "checkpoint_management_application_site_category" "example" {
  name = "New Application Site Category 1"
  description = "My Application Site category"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `description` - (Optional) N/A 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `groups` - (Optional) Collection of group identifiers.groups blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
