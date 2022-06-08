---
layout: "checkpoint"
page_title: "checkpoint_management_lsm_cluster_profile"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-lsm-cluster-profile"
description: |-
Use this data source to get information on an existing Check Point Lsm Cluster Profile.
---

# Data Source: checkpoint_management_lsm_cluster_profile

Use this data source to get information on an existing Check Point Lsm Cluster Profile.

## Example Usage


```hcl
data "checkpoint_management_lsm_cluster_profile" "example" {
  name = "cluster_profile"
}
```

## Argument Reference

The following arguments are supported:

* `uid` - (Optional) Object unique identifier.
* `name` - (Optional) Object name.

## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

