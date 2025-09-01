---
layout: "checkpoint"
page_title: "checkpoint_management_resource_mms"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-resource-mms"
description: |-
This resource allows you to execute Check Point Resource Mms.
---

# checkpoint_management_resource_mms

This resource allows you to execute Check Point Resource Mms.

## Example Usage


```hcl
resource "checkpoint_management_resource_mms" "example" {
  name = "newMmsResource"
  track = "log"
  action = "accept"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `track` - (Optional) Logs the activity when a packet matches on a Firewall Rule with the Resource. 
* `action` - (Optional) Accepts or Drops traffic that matches a Firewall Rule using the Resource. 
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
