---
layout: "checkpoint"
page_title: "checkpoint_hostname"
sidebar_current: "docs-checkpoint-gaia-resource-checkpoint-hostname"
description: |-
  This resource allows you to set the hostname of a Check Point machine.
---

# Resource: checkpoint_hostname

This resource allows you to set the hostname of a Check Point machine.

## Example Usage


```hcl
resource "checkpoint_hostname" "hostname" {
      name = "terrahost"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) New hostname to change.














