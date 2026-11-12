---
layout: "checkpoint"
page_title: "checkpoint_management_syslog_server"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-securid-server"
description: |- Use this data source to get information on an existing SecurID Server.
---


# checkpoint_management_securid_server

Use this data source to get information on an existing SecurID Server.

## Example Usage


```hcl
data "checkpoint_management_securid_server" "data_securid_server" {
  name = "securid_server_example"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object uid.
* `config_file_name` - Configuration file name.
* `base64_config_file_content` - Base64 encoded configuration file for authentication.<br>If no SecurID file was selected - this feature will work properly only if a SecurID file was already deployed manually on the target machine.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.
* `tags` - Collection of tag identifiers.tags blocks are documented below.
