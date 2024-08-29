---
layout: "checkpoint"
page_title: "checkpoint_management_cme_management"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-cme-management"
description: |- This resource allows you to update an existing Check Point CME Management.
---

# Resource: checkpoint_management_cme_management

This resource allows you to update an existing Check Point CME Management.

For details about the compatibility between the Terraform Release version and the CME API version, please refer to the section [Compatibility with CME](../index.html.markdown#compatibility-with-cme).


## Example Usage

```hcl
resource "checkpoint_management_cme_management" "mgmt" {
  name = "newManagement"
}
```

## Argument Reference

These arguments are supported:

* `name` - (Required) The name of the Management server.
* `domain` - (Optional) The management's domain name in Multi-Domain Security Management Server environment.
