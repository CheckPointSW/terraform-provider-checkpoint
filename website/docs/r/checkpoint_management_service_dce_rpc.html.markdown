---
layout: "checkpoint"
page_title: "checkpoint_management_service_dce_rpc"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-service-dce-rpc"
description: |-
This resource allows you to execute Check Point Service Dce Rpc.
---

# checkpoint_management_service_dce_rpc

This resource allows you to execute Check Point Service Dce Rpc.

## Example Usage


```hcl
resource "checkpoint_management_service_dce_rpc" "example" {
  name = "New_DCE-RPC_Service_1"
  interface_uuid = "97aeb460-9aea-11d5-bd16-0090272ccb30"
  keep_connections_open_after_policy_installation = false
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `interface_uuid` - (Optional) Network interface UUID. 
* `keep_connections_open_after_policy_installation` - (Optional) Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `groups` - (Optional) Collection of group identifiers.groups blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
