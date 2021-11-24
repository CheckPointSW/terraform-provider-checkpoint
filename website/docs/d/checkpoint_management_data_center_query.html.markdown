---
layout: "checkpoint"
page_title: "checkpoint_management_data_center_query"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-data-center-query"
description: |- Use this data source to get information on an existing Check Point Data Center Query.
---

# checkpoint_management_data_center_query

Use this data source to get information on an existing Check Point Data Center Query.

## Example Usage

```hcl
resource "checkpoint_management_data_center_query" "testQuery" {
  name         = "myQuery"
  data_centers = ["All"]
  query_rules {
    key_type = "predefined"
    key      = "name-in-data-center"
    values   = ["firstVal", "secondVal"]
  }
}

data "checkpoint_management_data_center_query" "data_center_query" {
  name = "${checkpoint_management_data_center_query.testQuery.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.
