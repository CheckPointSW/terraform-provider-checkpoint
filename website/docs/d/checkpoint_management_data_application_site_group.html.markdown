---
layout: "checkpoint"
page_title: "checkpoint_management_data_application_site_group"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-application-site-group"
description: |-
  Use this data source to get information on an existing Check Point Application Site Group.
---

# checkpoint_management_data_application_site_group

  Use this data source to get information on an existing Check Point Application Site Group.

## Example Usage


```hcl
resource "checkpoint_management_application_site_group" "application_site_group" {
    name = "my application site group"
}

data "checkpoint_management_data_application_site_group" "data_application_site_group" {
    name = "${checkpoint_management_application_site_group.application_site_group.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. 
* `uid` - (Optional) Object unique identifier. 
* `members` - Collection of application and URL filtering objects identified by the name or UID.
* `tags` - Collection of tag identifiers.
* `color` - Color of the object. Should be one of existing colors. 
* `comments` - Comments string. 
* `groups` - Collection of group identifiers.
