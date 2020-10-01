---
layout: "checkpoint"
page_title: "checkpoint_management_service_citrix_tcp"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-service-citrix-tcp"
description: |-
This resource allows you to execute Check Point Service Citrix Tcp.
---

# Resource: checkpoint_management_service_citrix_tcp

This resource allows you to execute Check Point Service Citrix Tcp.

## Example Usage


```hcl
resource "checkpoint_management_service_citrix_tcp" "example" {
  name = "mycitrixtcp"
  application = "My Citrix Application"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `application` - (Optional) Citrix application name. 
* `tags` - (Optional) Collection of tag identifiers.
* `color` - (Optional) Color of the object.
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.