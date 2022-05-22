---
layout: "checkpoint"
page_title: "checkpoint_management_threat_ioc_feed"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-threat-ioc-feed"
description: |-
Use this data source to get information on an existing Check Point Threat Ioc Feed.
---

# Data Source: checkpoint_management_threat_ioc_feed

Use this data source to get information on an existing Check Point Threat Ioc Feed.

## Example Usage


```hcl
resource "checkpoint_management_threat_ioc_feed" "example" {
  name = "ioc_feed"
  feed_url = "https://www.feedsresource.com/resource"
  action = "Prevent"
}

data "checkpoint_management_threat_ioc_feed" "data_threat_ioc_feed" {
  name = "${checkpoint_management_threat_ioc_feed.example.name}"
}
```

## Argument Reference

The following arguments are supported:

* `uid` - (Optional) Object unique identifier.
* `name` - (Optional) Object name.