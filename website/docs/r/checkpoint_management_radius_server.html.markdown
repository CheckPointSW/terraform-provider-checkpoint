---
layout: "checkpoint"
page_title: "checkpoint_management_radius_server"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-radius-server"
description: |-
This resource allows you to add/update/delete Check Point Radius Server.
---

# Resource: checkpoint_management_radius_server

This resource allows you to add/update/delete Check Point Radius Server.

## Example Usage


```hcl
resource "checkpoint_management_host" "host" {
  name = "My Host"
  ipv4_address = "1.2.3.4"
}

resource "checkpoint_management_radius_server" "example" {
  name = "New Radius Server"
  server = "${checkpoint_management_host.host.name}"
  shared_secret = "123"
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. Should be unique in the domain.
* `server` - (Required) The UID or Name of the host that is the RADIUS Server.
* `shared_secret` - (Required) The secret between the RADIUS server and the Security Gateway.
* `service` - (Optional) The UID or Name of the Service to which the RADIUS server listens.
* `version` - (Optional) The version can be either RADIUS Version 1.0, which is RFC 2138 compliant, and RADIUS Version 2.0 which is RFC 2865 compliant.
* `protocol` - (Optional) The type of authentication protocol that will be used when authenticating the user to the RADIUS server.
* `priority` - (Optional) The priority of the RADIUS Server in case it is a member of a RADIUS Group.
* `accounting` - (Optional) Accounting settings. accounting blocks are documented below.
* `tags` - (Optional) Collection of tag identifiers.
* `color` - (Optional) Color of the object. Should be one of existing colors.
* `comments` - (Optional) Comments string.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.


`accounting` supports the following:

* `enable_ip_pool_management` - (Optional) IP pool management, enables Accounting service.
* `accounting_service` - (Optional) The UID or Name of the the accounting interface to notify the server when users login and logout which will then lock and release the IP addresses that the server allocated to those users.
