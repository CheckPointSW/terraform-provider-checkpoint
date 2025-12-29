---
layout: "checkpoint"
page_title: "checkpoint_management_syslog_server"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-syslog-server"
description: |-
This resource allows you to execute Check Point Syslog Server.
---

# checkpoint_management_syslog_server

This resource allows you to execute Check Point Syslog Server.

## Example Usage


```hcl
resource "checkpoint_management_syslog_server" "example" {
  name = "newSyslogServer"
  host = "syslogServerHost"
  port = 18889
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `host` - (Required) Host server object identified by the name or UID. 
* `port` - (Optional) Port number. 
* `version` - (Optional) RFC version. 
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
