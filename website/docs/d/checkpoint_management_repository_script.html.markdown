---
layout: "checkpoint"
page_title: "checkpoint_management_repository_script"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-repository-script"
description: |-
Use this data source to get information on an existing Check Point Repository Script.
---

# Data Source: checkpoint_management_repository_script

Use this data source to get information on an existing Check Point Repository Script.

## Example Usage


```hcl
resource "checkpoint_management_repository_script" "example" {
  name = "New Script 1"
  script_body = "ls -l /"
}

data "checkpoint_management_repository_script" "data_repository_script" {
  name = "${checkpoint_management_repository_script.example.name}"
}
```

## Argument Reference

The following arguments are supported:

* `uid` - (Optional) Object unique identifier.
* `name` - (Optional) Object name.