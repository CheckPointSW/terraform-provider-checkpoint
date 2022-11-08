---
layout: "checkpoint"
page_title: "checkpoint_management_radius_server"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-radius-server"
description: |-
Use this data source to get information on an existing Check Point Radius Server.
---

# Data Source: checkpoint_management_radius_server

Use this data source to get information on an existing Check Point Radius Server.

## Example Usage


```hcl
resource "checkpoint_management_host" "host" {
  name = "My Host"
  ipv4_address = "1.2.3.4"
}

resource "checkpoint_management_radius_server" "radius_server" {
  name = "New Radius Server"
  server = "${checkpoint_management_host.host.name}"
  shared_secret = "123"
}

data "checkpoint_management_radius_server" "data_radius_server" {
  name = "${checkpoint_management_radius_server.radius_server.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. Should be unique in the domain.
* `uid` - (Optional) Object unique identifier. 
* `server` - The UID or Name of the host that is the RADIUS Server.
* `service` - The UID or Name of the Service to which the RADIUS server listens.
* `version` - The version can be either RADIUS Version 1.0, which is RFC 2138 compliant, and RADIUS Version 2.0 which is RFC 2865 compliant.
* `protocol` - The type of authentication protocol that will be used when authenticating the user to the RADIUS server.
* `priority` - The priority of the RADIUS Server in case it is a member of a RADIUS Group.
* `accounting` - Accounting settings. accounting blocks are documented below.
* `tags` - Collection of tag objects identified by the name or UID.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.


`accounting` supports the following:

* `enable_ip_pool_management` - IP pool management, enables Accounting service.
* `accounting_service` - The UID or Name of the the accounting interface to notify the server when users login and logout which will then lock and release the IP addresses that the server allocated to those users.
