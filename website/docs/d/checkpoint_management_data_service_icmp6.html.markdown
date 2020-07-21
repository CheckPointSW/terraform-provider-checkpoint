---
layout: "checkpoint"
page_title: "checkpoint_management_data_service_icmp6"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-service-icmp6"
description: |-
  Use this data source to get information on an existing Check Point Service Icmp6.
---

# checkpoint_management_data_service_icmp6

Use this data source to get information on an existing Check Point Service Icmp6.

## Example Usage


```hcl
resource "checkpoint_management_service_icmp6" "service_icmp6" {
    name = "service icmp6"
}

data "checkpoint_management_data_service_icmp6" "data_service_icmp6" {
    name = "${checkpoint_management_service_icmp6.service_icmp6.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.  
* `icmp_code` - As listed in: <a href="http://www.iana.org/assignments/icmp-parameters" target="_blank">RFC 792</a>. 
* `icmp_type` - As listed in: <a href="http://www.iana.org/assignments/icmp-parameters" target="_blank">RFC 792</a>. 
* `keep_connections_open_after_policy_installation` - (Optional) Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections. 
* `tags` - Collection of tag identifiers.tags blocks are documented below.
* `color` - Color of the object. Should be one of existing colors. 
* `comments` - Comments string. 
* `groups` - Collection of group identifiers.