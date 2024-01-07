---
layout: "checkpoint"
page_title: "checkpoint_management_cme_management"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-cme-management"
description: |- Use this data source to get information on existing Check Point CME Management.
---

# checkpoint_management_cme_management

Use this data source to get information on existing Check Point CME Management.

## Example Usage

```hcl
data "checkpoint_management_cme_management" "mgmt" {
}
```

## Argument Reference

The following arguments are supported:

* `name` - The name of the management server.
* `domain` - The management's domain name in MDS environment.
* `host` - The host of the management server.
