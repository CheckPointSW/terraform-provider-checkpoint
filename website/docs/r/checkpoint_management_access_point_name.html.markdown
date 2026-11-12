---
layout: "checkpoint"
page_title: "checkpoint_management_access_point_name"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-access-point-name"
description: |-
This resource allows you to execute Check Point Access Point Name.
---

# Resource: checkpoint_management_access_point_name

This resource allows you to execute Check Point Access Point Name.

## Example Usage


```hcl
resource "checkpoint_management_access_point_name" "example" {
  name = "myaccesspointname"
  apn = "apnname"
  enforce_end_user_domain = true
  end_user_domain = "All_Internet"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `apn` - (Optional) APN name. 
* `enforce_end_user_domain` - (Optional) Enable enforce end user domain. 
* `block_traffic_other_end_user_domains` - (Optional) Block MS to MS traffic between this and other APN end user domains. 
* `block_traffic_this_end_user_domain` - (Optional) Block MS to MS traffic within this end user domain. 
* `end_user_domain` - (Optional) End user domain name or UID. 
* `tags` - (Optional) Collection of tag identifiers.
* `color` - (Optional) Color of the object. 
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.