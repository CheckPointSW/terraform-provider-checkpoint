---
layout: "checkpoint"
page_title: "checkpoint_gaia_lightshot_partition"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-lightshot-partition"
description: |-
This resource allows you to execute Check Point Lightshot Partition.
---

# checkpoint_gaia_lightshot_partition

This resource allows you to execute Check Point Lightshot Partition.

## Example Usage


```hcl
resource "checkpoint_gaia_lightshot_partition" "example" {
  size = 17
}
```

## Argument Reference

The following arguments are supported:

* `size` - (Required) New size (GB) for setting light shotpartition 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 
