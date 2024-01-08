---
layout: "checkpoint"
page_title: "checkpoint_management_cme_delay_cycle"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-cme-delay-cycle"
description: |- This resource allows you to update Check Point CME Delay Cycle.
---

# checkpoint_management_cme_delay_cycle

This resource allows you to update Check Point CME Delay Cycle.

## Example Usage

```hcl
resource "checkpoint_management_cme_delay_cycle" "delay_cycle" {
  delay_cycle = 30
}
```

## Argument Reference

The following arguments are supported:

* `delay_cycle` - (Required) Time to wait in seconds after each poll cycle.
