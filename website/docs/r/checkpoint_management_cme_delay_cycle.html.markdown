---
layout: "checkpoint"
page_title: "checkpoint_management_cme_delay_cycle"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-cme-delay-cycle"
description: |- This resource allows you to update an existing Check Point CME Delay Cycle.
---

# Resource: checkpoint_management_cme_delay_cycle

This resource allows you to update an existing Check Point CME Delay Cycle.

For details about the compatibility between the Terraform Release version and the CME API version, please refer to the section [Compatibility with CME](../index.html.markdown#compatibility-with-cme).


## Example Usage

```hcl
resource "checkpoint_management_cme_delay_cycle" "delay_cycle" {
  delay_cycle = 30
}
```

## Argument Reference

These arguments are supported:

* `delay_cycle` - (Required) Time to wait in seconds after each poll cycle.
