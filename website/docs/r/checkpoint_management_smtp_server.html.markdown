---
layout: "checkpoint"
page_title: "checkpoint_management_smtp_server"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-smtp-server"
description: |-
This resource allows you to execute Check Point Smtp Server.
---

# checkpoint_management_smtp_server

This resource allows you to execute Check Point Smtp Server.

## Example Usage


```hcl
resource "checkpoint_management_smtp_server" "example" {
  name = "SMTP1"
  server = "smtp.example.com"
  port = 25
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `port` - (Required) The SMTP port to use. 
* `server` - (Required) The SMTP server address. 
* `authentication` - (Optional) Does the mail server requires authentication. 
* `encryption` - (Optional) Encryption type. 
* `password` - (Optional) A password for the SMTP server.Required only if authentication is set to true. 
* `username` - (Optional) A username for the SMTP server.Required only if authentication is set to true. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `domains_to_process` - (Optional) Indicates which domains to process the commands on. It cannot be used with the details-level full, must be run from the System Domain only and with ignore-warnings true. Valid values are: CURRENT_DOMAIN, ALL_DOMAINS_ON_THIS_SERVER.domains_to_process blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
