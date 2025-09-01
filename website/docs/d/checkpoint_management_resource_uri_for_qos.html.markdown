---
layout: "checkpoint"
page_title: "checkpoint_management_resource_uri_for_qos"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-resource-uri-for-qos"
description: |- Use this data source to get information on an existing Uri For QoS resource.
---


# checkpoint_management_resource_uri_for_qos

Use this data source to get information on an existing Uri For QoS resource.

## Example Usage


```hcl
data "checkpoint_management_resource_uri_for_qos" "data_uri_for_qos" {
  name = "uri_for_qos_example"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object uid.
* `search_for_url` - URL string that will be matched to an HTTP connection.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.
* `tags` -  Collection of tag identifiers.tags blocks are documented below.
