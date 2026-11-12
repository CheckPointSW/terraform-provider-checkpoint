---
layout: "checkpoint"
page_title: "checkpoint_management_securid_server"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-securid-server"
description: |-
This resource allows you to execute Check Point Securid Server.
---

# checkpoint_management_securid_server

This resource allows you to execute Check Point Securid Server.

## Example Usage


```hcl
resource "checkpoint_management_securid_server" "example" {
  name = "TestSecurIdServer"
  config_file_name = "configFile"
  base64_config_file_content = "Q0xJRU5UX0lQPSAxLjEuMS4xMQ=="
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `config_file_name` - (Optional) Configuration file name. <font color="red">Required only when</font> 'base64-config-file-content' is not empty. 
* `base64_config_file_content` - (Optional) Base64 encoded configuration file for authentication.<br>If no SecurID file was selected - this feature will work properly only if a SecurID file was already deployed manually on the target machine. 
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
