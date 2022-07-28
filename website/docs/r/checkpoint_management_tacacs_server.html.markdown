---
layout: "checkpoint"
page_title: "checkpoint_management_tacacs_server"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-tacacs-server"
description: |-
This resource allows you to execute Check Point Tacacs Server.
---

# checkpoint_management_tacacs_server

This resource allows you to execute Check Point Tacacs Server.

## Example Usage


```hcl
resource "checkpoint_management_tacacs_server" "example" {
  name = "My Tacacs Server"
  server = "h1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `secret_key` - (Optional) The server's secret key.<br><font color="red">Required only when</font> "server-type" was selected to be "TACACS+". 
* `server` - (Required) The UID or Name of the host that is the TACACS Server. 
* `encryption` - (Optional) Is there a secret key defined on the server. Must be set true when "server-type" was selected to be "TACACS+". 
* `priority` - (Optional) The priority of the TACACS Server in case it is a member of a TACACS Group. 
* `server_type` - (Optional) Server type, TACACS or TACACS+. 
* `service` - (Optional) Server service, only relevant when "server-type" is TACACS. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
