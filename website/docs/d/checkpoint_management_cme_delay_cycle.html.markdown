---
layout: "checkpoint"
page_title: "checkpoint_management_cme_delay_cycle"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-cme-delay-cycle"
description: |- Use this data source to get information on existing Check Point CME Delay Cycle.
---

# Data Source: checkpoint_management_cme_management

Use this data source to get information on existing Check Point CME Delay Cycle.

## Example Usage

```hcl
data "checkpoint_management_cme_delay_cycle" "delay_cycle" {
}
```

## Argument Reference

The following arguments are supported:

* `delay_cycle` - Time to wait in seconds after each poll cycle.
