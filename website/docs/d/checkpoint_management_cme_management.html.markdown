---
layout: "checkpoint"
page_title: "checkpoint_management_cme_management"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-cme-management"
description: |- Use this data source to get information on existing Check Point CME Management.
---

# Data Source: checkpoint_management_cme_management

Use this data source to get information on existing Check Point CME Management.

For details about the compatibility between the Terraform Release version and the CME API version, please refer to the section [Compatibility with CME](https://registry.terraform.io/providers/CheckPointSW/checkpoint/latest/docs#compatibility-with-cme).


## Example Usage

```hcl
data "checkpoint_management_cme_management" "mgmt" {
}
```

## Argument Reference

These arguments are supported:

* `name` - The name of the Management server.
* `domain` - The management's domain name in Multi-Domain Security Management Server environment.
* `host` - The host of the Management server.
