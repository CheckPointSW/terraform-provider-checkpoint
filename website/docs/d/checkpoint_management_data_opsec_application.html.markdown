---
layout: "checkpoint"
page_title: "checkpoint_management_data_opsec_application"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-opsec-application"
description: |-
  Use this data source to get information on an existing Check Point Opsec Application.
---

# checkpoint_management_data_opsec_application

Use this data source to get information on an existing Check Point Opsec Application.

## Example Usage


```hcl
resource "checkpoint_management_host" "myhost" {
    name = "myhost"
    ipv4_address = "1.2.3.4" 
}

resource "checkpoint_management_opsec_application" "opsec_application" {
    name = "OPSEC application"
    host = "${checkpoint_management_host.myhost.name}"
    cpmi = {
        enabled = true
        administrator_profile = "read only all"
        use_administrator_credentials = false
    }
    lea = {
        enabled = true
        access_permissions = "show all"
    }
}

data "checkpoint_management_data_opsec_application" "data_opsec_application" {
    name = "${checkpoint_management_opsec_application.opsec_application.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.  
* `cpmi` - Used to setup the CPMI client entity. cpmi blocks are documented below.
* `host` - The host where the server is running. Pre-define the host as a network object. 
* `lea` - Used to setup the LEA client entity. lea blocks are documented below.
* `tags` - Collection of tag identifiers.
* `color` - Color of the object. Should be one of existing colors. 
* `comments` - Comments string. 


`cpmi` supports the following:

* `administrator_profile` - A profile to set the log reading permissions by for the client entity. 
* `enabled` - Whether to enable this client entity on the Opsec Application. 
* `use_administrator_credentials` - Whether to use the Admin's credentials to login to the security management server. 


`lea` supports the following:

* `access_permissions` - Log reading permissions for the LEA client entity. 
* `administrator_profile` - A profile to set the log reading permissions by for the client entity. 
* `enabled` - Whether to enable this client entity on the Opsec Application. 
