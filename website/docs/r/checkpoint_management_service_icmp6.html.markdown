---
layout: "checkpoint"
page_title: "checkpoint_management_service_icmp6"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-service-icmp6"
description: |-
This resource allows you to execute Check Point Service Icmp6.
---

# checkpoint_management_service_icmp6

This resource allows you to execute Check Point Service Icmp6.

## Example Usage


```hcl
resource "checkpoint_management_service_icmp6" "example" {
  name = "Icmp1"
  icmp_type = 5
  icmp_code = 7
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `icmp_code` - (Optional) As listed in: <a href="http://www.iana.org/assignments/icmp-parameters" target="_blank">RFC 792</a>. 
* `icmp_type` - (Optional) As listed in: <a href="http://www.iana.org/assignments/icmp-parameters" target="_blank">RFC 792</a>. 
* `keep_connections_open_after_policy_installation` - (Optional) Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `groups` - (Optional) Collection of group identifiers.groups blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
