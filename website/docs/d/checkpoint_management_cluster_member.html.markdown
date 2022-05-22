---
layout: "checkpoint"
page_title: "checkpoint_management_cluster_member"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-cluster-member"
description: |-
Use this data source to get information on an existing Check Point Cluster Member.
---

# Data Source: checkpoint_management_cluster_member

Use this data source to get information on an existing Check Point Cluster Member.

## Example Usage


```hcl
data "checkpoint_management_cluster_member" "data_cluster_member" {
  uid = "CLUSTER_MEMBER_UID"
  limit_interfaces = 20
}
```

## Argument Reference

The following arguments are supported:

* `uid` - (Required) Object unique identifier.
* `limit_interfaces` - (Optional) Limit number of cluster member interfaces to show.