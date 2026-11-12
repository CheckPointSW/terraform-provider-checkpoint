---
layout: "checkpoint"
page_title: "checkpoint_management_data_dns_domain"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-dns-domain"
description: |-
  Use this data source to get information on an existing Check Point Dns Domain.
---

# Data Source: checkpoint_management_data_dns_domain

Use this data source to get information on an existing Check Point Dns Domain.

## Example Usage


```hcl
resource "checkpoint_management_dns_domain" "dns_domain" {
        name = "My DNS domain"
		is_sub_domain = true
}

data "checkpoint_management_data_dns_domain" "data_dns_domain" {
    name = "${checkpoint_management_dns_domain.dns_domain.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. 
* `uid` - (Optional) Object unique identifier. 
* `is_sub_domain` - Whether to match sub-domains in addition to the domain itself. 
* `tags` - Collection of tag identifiers.
* `color` - Color of the object. Should be one of existing colors. 
* `comments` - Comments string. 
