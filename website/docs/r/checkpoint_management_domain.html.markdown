---
layout: "checkpoint"
page_title: "checkpoint_management_domain"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-domain"
description: |- This resource allows you to execute Check Point Domain.
---

# Resource: checkpoint_management_domain

This resource allows you to execute Check Point Domain.

## Example Usage

```hcl
resource "checkpoint_management_domain" "example" {
    name = "domain1"
    servers {
      name = "domain1_ManagementServer_1"
      ipv4_address = "192.0.2.1"
      multi_domain_server = "MDM_Server"
    }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. Must be unique in the domain.
* `servers` - (Required) Domain servers.
* `color` - (Optional) Color of the object. Should be one of existing colors.
* `comments` - (Optional) Comments string.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If
  ignore-warnings flag was omitted - warnings will also be ignored.

`servers` supports the following:

* `name` - (Required) Object name. Must be unique in the domain.
* `ipv4_address` - (Optional) IPv4 address.
* `ipv6_address` - (Optional) IPv6 address.
* `multi_domain_server` - (Required) Multi Domain server name or UID.
* `active` - (Optional) Activate domain server. Only one domain server is allowed to be active.
* `skip_start_domain_server` - (Optional) Set this value to be true to prevent starting the new created domain.
* `type` - (Optional) Domain server type.
