---
layout: "checkpoint"
page_title: "checkpoint_management_data_service_icmp"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-service-icmp"
description: |-
  Use this data source to get information on an existing Check Point Service Icmp.
---

# Data Source: checkpoint_management_data_service_icmp

Use this data source to get information on an existing Check Point Service Icmp.

## Example Usage


```hcl
resource "checkpoint_management_service_icmp" "service_icmp" {
    name = "icmp group"
}

data "checkpoint_management_data_service_icmp" "data_service_icmp" {
    name = "${checkpoint_management_service_icmp.service_icmp.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name.
* `uid` - (Optional) Object unique identifier.  
* `icmp_code` - As listed in: <a href="http://www.iana.org/assignments/icmp-parameters" target="_blank">RFC 792</a>. 
* `icmp_type` - As listed in: <a href="http://www.iana.org/assignments/icmp-parameters" target="_blank">RFC 792</a>. 
* `keep_connections_open_after_policy_installation` - Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections. 
* `tags` - Collection of tag identifiers.
* `color` - Color of the object. Should be one of existing colors. 
* `comments` - Comments string. 
* `groups` - Collection of group identifiers.
