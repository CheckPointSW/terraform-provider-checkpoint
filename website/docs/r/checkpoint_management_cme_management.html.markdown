---
layout: "checkpoint"
page_title: "checkpoint_management_cme_management"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-cme-management"
description: |- This resource allows you to update Check Point CME Management.
---

# Resource: checkpoint_management_cme_management

This resource allows you to update Check Point CME Management.

Available in:

- Check Point Security Management/Multi Domain Management Server R81.10 and higher.
- CME take 255 and higher.

## Example Usage

```hcl
resource "checkpoint_management_cme_management" "mgmt" {
  name = "newManagement"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the management server.
* `domain` - (Optional) The management's domain name in MDS environment.
