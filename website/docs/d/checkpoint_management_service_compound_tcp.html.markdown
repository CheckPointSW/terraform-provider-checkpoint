---
layout: "checkpoint"
page_title: "checkpoint_management_service_compound_tcp"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-service-compound-tcp"
description: |-
This resource allows you to execute Check Point Service Compound Tcp.
---

# Data Source: checkpoint_management_service_compound_tcp

This resource allows you to execute Check Point Service Compound Tcp.

## Example Usage


```hcl
resource "checkpoint_management_service_compound_tcp" "service_compound_tcp" {
    name = "service compound tcp"
    compound_service = "pointcast"
    keep_connections_open_after_policy_installation = true
}

data "checkpoint_management_service_compound_tcp" "test" {
    name = "${checkpoint_management_service_compound_tcp.service_compound_tcp.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.   
* `compound_service` - Compound service type. 
* `keep_connections_open_after_policy_installation` - Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections. 
* `tags` - Collection of tag identifiers.
* `color` - Color of the object.
* `comments` - Comments string. 
