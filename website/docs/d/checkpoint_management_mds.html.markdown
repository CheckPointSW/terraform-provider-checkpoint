---
layout: "checkpoint"
page_title: "checkpoint_management_mds"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-mds"
description: |-
This resource allows you to execute Check Point MDS.
---

# Data Source: checkpoint_management_mds

This resource allows you to execute Check Point MDS.

## Example Usage


```hcl
resource "checkpoint_management_mds" "mds" {
    name = "my mds"
    ipv4_address = "2.2.2.2"
}

data "checkpoint_management_mds" "data_mds" {
    name = "${checkpoint_management_mds.mds.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.  
* `ipv4_address` - IPv4 address. 
* `ipv6_address` - IPv6 address. 
* `hardware` - Hardware name. For example: Open server, Smart-1, Other. 
* `os` - Operating system name. For example: Gaia, Linux, SecurePlatform. 
* `version` - System version. 
* `ip_pool_first` - First IP address in the range. 
* `ip_pool_last` - Last IP address in the range. 
* `tags` - Collection of tag identifiers.
* `color` - Color of the object.
* `comments` - Comments string. 
* `server_type` - Type of the management server.
* `sic_name` - Name of the Secure Internal Connection Trust..
* `sic_state` - State the Secure Internal Connection Trust..
* `domains` - Collection of Domain objects identified by the name or UID.
* `global_domains` - Collection of Global domain objects identified by the name or UID.