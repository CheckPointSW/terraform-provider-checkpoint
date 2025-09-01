---
layout: "checkpoint"
page_title: "checkpoint_management_resource_uri_for_qos"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-resource-uri-for-qos"
description: |-
This resource allows you to execute Check Point Resource Uri For Qos.
---

# checkpoint_management_resource_uri_for_qos

This resource allows you to execute Check Point Resource Uri For Qos.

## Example Usage


```hcl
resource "checkpoint_management_resource_uri_for_qos" "example" {
  name = "newUriForQosResource"
  search_for_url = "www.checkpoint.com"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `search_for_url` - (Optional) URL string that will be matched to an HTTP connection. 
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
