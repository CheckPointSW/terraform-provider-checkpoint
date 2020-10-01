---
layout: "checkpoint"
page_title: "checkpoint_management_mds"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-mds"
description: |-
This resource allows you to execute Check Point MDS.
---

# Resource: checkpoint_management_mds

This resource allows you to execute Check Point MDS.

## Example Usage


```hcl
resource "checkpoint_management_mds" "example" {
  name = "mymds"
  server_type = "multi-domain server"
  os = "Gaia"
  hardware = "Open server"
  ipv4_address = "1.1.1.1"
  ip_pool_first = "2.2.2.2"
  ip_pool_last = "3.3.3.3"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `ipv4_address` - (Optional) IPv4 address. 
* `ipv6_address` - (Optional) IPv6 address. 
* `hardware` - (Optional) Hardware name. For example: Open server, Smart-1, Other. 
* `os` - (Optional) Operating system name. For example: Gaia, Linux, SecurePlatform. 
* `version` - (Optional) System version. 
* `one_time_password` - (Optional) Secure internal connection one time password. 
* `ip_pool_first` - (Optional) First IP address in the range. 
* `ip_pool_last` - (Optional) Last IP address in the range. 
* `tags` - (Optional) Collection of tag identifiers.
* `color` - (Optional) Color of the object.
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
* `server_type` - (Optional) Type of the management server.
* `sic_name` - (Computed) Name of the Secure Internal Connection Trust.
* `sic_state` - (Computed) State the Secure Internal Connection Trust.
* `domains` - (Computed) Collection of Domain objects identified by the name or UID.
* `global_domains` - (Computed) Collection of Global domain objects identified by the name or UID.