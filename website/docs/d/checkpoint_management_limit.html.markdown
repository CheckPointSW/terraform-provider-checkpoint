---
layout: "checkpoint"
page_title: "checkpoint_management_limit"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-limit"
description: |-
Use this data source to get information on an existing Check Point Limit.
---

# checkpoint_management_limit

Use this data source to get information on an existing Check Point Limit.

## Example Usage


```hcl
resource "checkpoint_management_limit" "example" {
  name = "limit_obj"
  enable_download = true
  download_unit = "Gbps"
  download_rate = 3
}

data "checkpoint_management_limit" "data" {
  name = "${checkpoint_management_limit.example.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. 
* `uid` - (Optional) Object unique identifier.
* `enable_download` -  Enable throughput limit for downloads from the internet to the organization. 
* `download_rate` -  The Rate for the maximum permitted bandwidth. 
* `download_unit` - The Unit for the maximum permitted bandwidth. 
* `enable_upload` - Enable throughput limit for uploads from the organization to the internet. 
* `upload_rate` - The Rate for the maximum permitted bandwidth. 
* `upload_unit` - The Unit for the maximum permitted bandwidth. 
* `tags` - Collection of tag identifiers.tags blocks are documented below.
* `color` - Color of the object. Should be one of existing colors. 
* `comments` - Comments string. 

