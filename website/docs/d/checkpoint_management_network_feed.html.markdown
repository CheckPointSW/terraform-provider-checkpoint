---
layout: "checkpoint"
page_title: "checkpoint_management_network_feed"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-network-feed"
description: |-
Use this data source to get information on an existing Check Point Network Feed.
---

# Data Source: checkpoint_management_network_feed

Use this data source to get information on an existing Check Point Network Feed.

## Example Usage


```hcl
resource "checkpoint_management_network_feed" "example" {
  name = "network_feed"
  feed_url = "https://www.feedsresource.com/resource"
  username = "feed_username"
  password = "feed_password"
  feed_format = "Flat List"
  feed_type = "IP Address"
  update_interval = 60
  data_column = 1
  use_gateway_proxy = false
  fields_delimiter = "	"
  ignore_lines_that_start_with = "!"
}

data "checkpoint_management_network_feed" "data_network_feed" {
  name = "${checkpoint_management_network_feed.example.name}"
}
```

## Argument Reference

The following arguments are supported:

* `uid` - (Optional) Object unique identifier.
* `name` - (Optional) Object name.