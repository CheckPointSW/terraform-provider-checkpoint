---
layout: "checkpoint"
page_title: "checkpoint_gaia_custom_intelligence_interval"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-custom-intelligence-interval"
description: |-
This resource allows you to execute Check Point Custom Intelligence Interval.
---

# checkpoint_gaia_custom_intelligence_interval

This resource allows you to execute Check Point Custom Intelligence Interval.

## Example Usage


```hcl
resource "checkpoint_gaia_custom_intelligence_interval" "example" {
  interval = 900
}
```

## Argument Reference

The following arguments are supported:

* `interval` - (Required) Check for updates frequency 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 
