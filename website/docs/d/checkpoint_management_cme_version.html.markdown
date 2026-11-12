---
layout: "checkpoint"
page_title: "checkpoint_management_cme_version"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-cme-version"
description: |- Use this data source to get information on existing Check Point CME version.
---

# Data Source: checkpoint_management_cme_version

Use this data source to get information on existing Check Point CME version.

For details about the compatibility between the Terraform Release version and the CME API version, please refer to the section [Compatibility with CME](https://registry.terraform.io/providers/CheckPointSW/checkpoint/latest/docs#compatibility-with-cme).


## Example Usage

```hcl
data "checkpoint_management_cme_version" "cme_version" {
}
```

## Argument Reference

These arguments are supported:

* `take` - CME Take number.
