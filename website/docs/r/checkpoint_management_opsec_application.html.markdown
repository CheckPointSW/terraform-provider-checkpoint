---
layout: "checkpoint"
page_title: "checkpoint_management_opsec_application"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-opsec-application"
description: |-
This resource allows you to execute Check Point Opsec Application.
---

# checkpoint_management_opsec_application

This resource allows you to execute Check Point Opsec Application.

## Example Usage


```hcl
resource "checkpoint_management_opsec_application" "example" {
  name = "MyOpsecApplication"
  host = "SomeHost"
  one_time_password = "SomePassword"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `cpmi` - (Optional) Used to setup the CPMI client entity.cpmi blocks are documented below.
* `host` - (Optional) The host where the server is running. Pre-define the host as a network object. 
* `lea` - (Optional) Used to setup the LEA client entity.lea blocks are documented below.
* `one_time_password` - (Optional) A password required for establishing a Secure Internal Communication (SIC). 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


`cpmi` supports the following:

* `administrator_profile` - (Optional) A profile to set the log reading permissions by for the client entity. 
* `enabled` - (Optional) Whether to enable this client entity on the Opsec Application. 
* `use_administrator_credentials` - (Optional) Whether to use the Admin's credentials to login to the security management server. 


`lea` supports the following:

* `access_permissions` - (Optional) Log reading permissions for the LEA client entity. 
* `administrator_profile` - (Optional) A profile to set the log reading permissions by for the client entity. 
* `enabled` - (Optional) Whether to enable this client entity on the Opsec Application. 
