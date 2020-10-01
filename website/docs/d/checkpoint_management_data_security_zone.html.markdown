---
layout: "checkpoint"
page_title: "checkpoint_management_data_security_zone"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-security-zone"
description: |-
  Use this data source to get information on an existing Check Point Security Zone.
---

# Data Source: checkpoint_management_data_security_zone

Use this data source to get information on an existing Check Point Security Zone.

## Example Usage


```hcl
resource "checkpoint_management_security_zone" "security_zone" {
    name = "Security Zone 1"
}

data "checkpoint_management_data_security_zone" "data_security_zone" {
    name = "${checkpoint_management_security_zone.security_zone.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier. 
* `tags` - Collection of tag identifiers.
* `color` - Color of the object. Should be one of existing colors. 
* `comments` - Comments string. 