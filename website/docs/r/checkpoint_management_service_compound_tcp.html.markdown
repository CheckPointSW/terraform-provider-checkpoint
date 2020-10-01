---
layout: "checkpoint"
page_title: "checkpoint_management_service_compound_tcp"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-service-compound-tcp"
description: |-
This resource allows you to execute Check Point Service Compound Tcp.
---

# Resource: checkpoint_management_service_compound_tcp

This resource allows you to execute Check Point Service Compound Tcp.

## Example Usage


```hcl
resource "checkpoint_management_service_compound_tcp" "example" {
  name = "mycompoundtcp"
  compound_service = "pointcast"
  keep_connections_open_after_policy_installation = true
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `compound_service` - (Optional) Compound service type. 
* `keep_connections_open_after_policy_installation` - (Optional) Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections. 
* `tags` - (Optional) Collection of tag identifiers.
* `color` - (Optional) Color of the object.
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.