---
layout: "checkpoint"
page_title: "checkpoint_management_limit"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-limit"
description: |-
This resource allows you to execute Check Point Limit.
---

# checkpoint_management_limit

This resource allows you to execute Check Point Limit.

## Example Usage


```hcl
resource "checkpoint_management_limit" "example" {
  name = "limit_obj"
  enable_download = true
  download_unit = "Gbps"
  download_rate = 3
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `enable_download` - (Optional) Enable throughput limit for downloads from the internet to the organization. 
* `download_rate` - (Optional) The Rate for the maximum permitted bandwidth. 
* `download_unit` - (Optional) The Unit for the maximum permitted bandwidth. 
* `enable_upload` - (Optional) Enable throughput limit for uploads from the organization to the internet. 
* `upload_rate` - (Optional) The Rate for the maximum permitted bandwidth. 
* `upload_unit` - (Optional) The Unit for the maximum permitted bandwidth. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
