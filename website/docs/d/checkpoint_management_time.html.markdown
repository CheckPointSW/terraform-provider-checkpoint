---
layout: "checkpoint"
page_title: "checkpoint_management_time"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-time"
description: |-
Use this data source to get information on an existing Check Point Time.
---

# Data Source: checkpoint_management_application_site

Use this data source to get information on an existing Check Point Time.

## Example Usage


```hcl
resource "checkpoint_management_time" "example" {
  name = "time1"
}
data "checkpoint_management_time" "data_time" {
  name = "${checkpoint_management_time.example.name}"
}
```

## Argument Reference

The following arguments are supported:

* `uid` - (Optional) Object unique identifier.
* `name` - (Optional) Object name.
