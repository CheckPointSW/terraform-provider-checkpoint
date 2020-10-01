---
layout: "checkpoint"
page_title: "checkpoint_management_data_service_dce_rpc"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-service-dce-rpc"
description: |-
  Use this data source to get information on an existing Check Point Service Dce Rpc.
---

# Data Source: checkpoint_management_data_service_dce_rpc

Use this data source to get information on an existing Check Point Service Dce Rpc.

## Example Usage


```hcl
resource "checkpoint_management_service_dce_rpc" "service_dce_rpc" {
    name = "service dce rpc"
    interface_uuid = "97aeb460-9aea-11d5-bd16-0090272ccb30"
}

data "checkpoint_management_data_service_dce_rpc" "data_service_dce_rpc" {
    name = "${checkpoint_management_service_dce_rpc.service_dce_rpc.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.  
* `interface_uuid` - Network interface UUID. 
* `keep_connections_open_after_policy_installation` - Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections. 
* `tags` - Collection of tag identifiers.
* `color` - Color of the object. Should be one of existing colors. 
* `comments` - Comments string. 
* `groups` - Collection of group identifiers.