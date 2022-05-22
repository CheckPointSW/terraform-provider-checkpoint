---
layout: "checkpoint"
page_title: "checkpoint_management_trusted-client"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-trusted-client"
description: |-
This resource allows you to execute Check Point Trusted Client.
---

# Resource: checkpoint_management_trusted_client

This resource allows you to execute Check Point Trusted Client.

## Example Usage


```hcl
resource "checkpoint_management_trustedClient" "example" {
  name = "New TrustedClient 1"
  ipv4_address = "192.168.2.1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `ipv4_address` - (Optional) IPv4 address. 
* `ipv6_address` - (Optional) IPv6 address. 
* `ipv4_address_first` - (Optional) First IPv4 address in the range.
* `ipv6_address_first` - (Optional) First IPv6 address in the range.
* `ipv4_address_last` - (Optional) Last IPv4 address in the range.
* `ipv6_address_last` - (Optional) Last IPv6 address in the range.
* `domains_assignment` - (Optional) Domains to be added to this profile. Use domain name only. See example below: "add-trusted-client (with domain)".
* `mask_length4` - (Optional) IPv4 mask length.
* `mask_length6` - (Optional) IPv6 mask length.
* `multi_domain_server_trusted_client` - (Optional) Let this trusted client connect to all Multi-Domain Servers in the deployment.
* `type` - (Optional) Trusted client type.
* `wild_card` - (Optional) IP wild card (e.g. 192.0.2.*).
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
