---
layout: "checkpoint"
page_title: "checkpoint_management_cme_version"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-cme-version"
description: |- Use this data source to get information on existing Check Point CME version.
---

# checkpoint_management_cme_version

Use this data source to get information on existing Check Point CME version.

## Example Usage

```hcl
data "checkpoint_management_cme_version" "cme_version" {
}
```

## Argument Reference

The following arguments are supported:

* `take` - CME take number.
