---
layout: "checkpoint"
page_title: "checkpoint_management_dns_domain"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-dns-domain"
description: |-
This resource allows you to execute Check Point Dns Domain.
---

# checkpoint_management_dns_domain

This resource allows you to execute Check Point Dns Domain.

## Example Usage


```hcl
resource "checkpoint_management_dns_domain" "example" {
  name = ".www.example.com"
  is_sub_domain = false
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `is_sub_domain` - (Optional) Whether to match sub-domains in addition to the domain itself. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
