---
layout: "checkpoint"
page_title: "checkpoint_management_domain"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-domain"
description: |- Use this data source to get information on an existing Check Point Domain.
---

# Data Source: checkpoint_management_domain

Use this data source to get information on an existing Check Point Domain.

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

data "checkpoint_management_domain" "data_domain" {
  name = "${checkpoint_management_domain.example.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. Must be unique in the domain.
* `uid` - (Optional) Object unique identifier.
