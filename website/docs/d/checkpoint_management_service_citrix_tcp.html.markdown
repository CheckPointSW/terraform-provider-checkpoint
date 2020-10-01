---
layout: "checkpoint"
page_title: "checkpoint_management_service_citrix_tcp"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-service-citrix-tcp"
description: |-
This resource allows you to execute Check Point Service Citrix Tcp.
---

# Data Source: checkpoint_management_service_citrix_tcp

This resource allows you to execute Check Point Service Citrix Tcp.

## Example Usage


```hcl
resource "checkpoint_management_service_citrix_tcp" "service_citrix_tcp" {
     name = "citrix tcp"
     application = "app name"
}

data "checkpoint_management_service_citrix_tcp" "test" {
    name = "${checkpoint_management_service_citrix_tcp.service_citrix_tcp.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.  
* `application` - Citrix application name. 
* `tags` - Collection of tag identifiers.
* `color` - Color of the object.
* `comments` - Comments string. 