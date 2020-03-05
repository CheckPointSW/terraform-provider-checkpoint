---
layout: "checkpoint"
page_title: "checkpoint_management_set_global_domain"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-set-global-domain"
description: |-
This resource allows you to execute Check Point Set Global Domain.
---

# checkpoint_management_set_global_domain

This resource allows you to execute Check Point Set Global Domain.

## Example Usage


```hcl
resource "checkpoint_management_set_global_domain" "example" {
  name = "Global2"
  comments = "this is my Global domain"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `servers` - (Optional) Multi Domain Servers. When the field is provided, 'set-global-domain' command is executed asynchronously.servers blocks are documented below.
* `tags` - (Optional) Collection of tag identifiers. Note: The list of tags can not be modified in a singlecommand together with the domain servers. To modify tags, please use the separate 'set-global-domain' command, without providing the list of domain servers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


`servers` supports the following:

* `add` - (Optional) Adds to collection of valuesadd blocks are documented below.


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

