---
layout: "checkpoint"
page_title: "checkpoint_management_cme_management"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-cme-management"
description: |- Use this data source to get information on existing Check Point CME Management.
---

# Data Source: checkpoint_management_cme_management

Use this data source to get information on existing Check Point CME Management.

Available in:

- Check Point Security Management/Multi-Domain Security Management Server R81.10 and higher.
- CME Take 255 and higher.

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
