---
layout: "checkpoint"
page_title: "checkpoint_management_api_settings"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-api-settings"
description: |-
Use this data source to get information on an existing Check Point Api Settings.
---

# Data Source: checkpoint_management_api_settings

Use this data source to get information on an existing Check Point Api Settings.

## Example Usage

```hcl
data "checkpoint_management_api_settings" "data_api_settings" {
  
}
```

## Argument Reference

The following arguments are supported:

* `name` - 	Object name.
* `uid` - Object unique identifier.
* `accepted_api_calls_from` - Clients allowed to connect to the API Server.
* `automatic_start` - MGMT API will start after server will start.
* `tags` - Collection of tag identifiers.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.