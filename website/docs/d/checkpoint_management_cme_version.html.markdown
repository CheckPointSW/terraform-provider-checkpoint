---
layout: "checkpoint"
page_title: "checkpoint_management_cme_version"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-cme-version"
description: |- Use this data source to get information on existing Check Point CME version.
---

# Data Source: checkpoint_management_cme_version

Use this data source to get information on existing Check Point CME version.

Available in:

- Check Point Security Management/Multi-Domain Security Management Server R81.10 and higher.
- CME Take 255 and higher.

## Example Usage

```hcl
data "checkpoint_management_cme_version" "cme_version" {
}
```

## Argument Reference

These arguments are supported:

* `take` - CME Take number.
