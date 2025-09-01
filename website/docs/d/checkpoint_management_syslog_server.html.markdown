---
layout: "checkpoint"
page_title: "checkpoint_management_syslog_server"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-syslog-server"
description: |- Use this data source to get information on an existing Syslog Server.
---


# checkpoint_management_syslog_server

Use this data source to get information on an existing Syslog Server.

## Example Usage


```hcl
data "checkpoint_management_syslog_server" "data_syslog_server" {
  name = "syslog_server_example"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. 
* `uid` - (Optional) Object UID.
* `host` - Host server object identified by the name or UID. 
* `port` - Port number. 
* `version` - RFC version. 
* `color` - Color of the object. Should be one of existing colors. 
* `comments` - Comments string. 
* `tags` - Collection of tag identifiers.tags blocks are documented below.
* `icon` - Object icon.
