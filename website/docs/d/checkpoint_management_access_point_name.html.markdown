---
layout: "checkpoint"
page_title: "checkpoint_management_access_point_name"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-access-point-name"
description: |-
This resource allows you to execute Check Point Access Point Name.
---

# Data Source: checkpoint_management_access_point_name

This resource allows you to execute Check Point Access Point Name.

## Example Usage


```hcl
resource "checkpoint_management_access_point_name" "access_point_name" {
    name = "My Access Point Name"
    apn = "apn name"
    enforce_end_user_domain = true
    end_user_domain = "All_Internet"
}

data "checkpoint_management_access_point_name" "data_access_point_name" {
    name = "${checkpoint_management_access_point_name.access_point_name.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier. 
* `apn` - APN name. 
* `enforce_end_user_domain` - Enable enforce end user domain.
* `block_traffic_other_end_user_domains` - Block MS to MS traffic between this and other APN end user domains.
* `block_traffic_this_end_user_domain` - Block MS to MS traffic within this end user domain.
* `end_user_domain` - End user domain name or UID.
* `tags` - Collection of tag identifiers.
* `color` - Color of the object.
* `comments` - Comments string. 
