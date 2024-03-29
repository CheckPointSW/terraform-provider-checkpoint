---
layout: "checkpoint"
page_title: "checkpoint_management_cme_delay_cycle"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-cme-delay-cycle"
description: |- This resource allows you to update an existing Check Point CME Delay Cycle.
---

# Resource: checkpoint_management_cme_delay_cycle

This resource allows you to update an existing Check Point CME Delay Cycle.

Available in:

- Check Point Security Management/Multi-Domain Security Management Server R81.10 and higher.
- CME Take 255 and higher.

## Example Usage

```hcl
resource "checkpoint_management_cme_delay_cycle" "delay_cycle" {
  delay_cycle = 30
}
```

## Argument Reference

These arguments are supported:

* `delay_cycle` - (Required) Time to wait in seconds after each poll cycle.
