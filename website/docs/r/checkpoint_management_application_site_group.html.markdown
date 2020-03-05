---
layout: "checkpoint"
page_title: "checkpoint_management_application_site_group"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-application-site-group"
description: |-
This resource allows you to execute Check Point Application Site Group.
---

# checkpoint_management_application_site_group

This resource allows you to execute Check Point Application Site Group.

## Example Usage


```hcl
resource "checkpoint_management_application_site_group" "example" {
  name = "New Application Site Group 1"
  members = ["facebook", "Social Networking", "New Application Site 1", "New Application Site Category 1",]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `members` - (Optional) Collection of application and URL filtering objects identified by the name or UID.members blocks are documented below.
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `groups` - (Optional) Collection of group identifiers.groups blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
