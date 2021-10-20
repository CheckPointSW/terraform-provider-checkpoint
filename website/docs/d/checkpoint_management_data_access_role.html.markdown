---
layout: "checkpoint"
page_title: "checkpoint_management_data_access_role"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-access-role"
description: |- Use this data source to get information on an existing Check Point Access Role.
---

# Data Source: checkpoint_management_data_access_role

Use this data source to get information on an existing Check Point Access Role.

## Example Usage

```hcl
resource "checkpoint_management_access_role" "access_role" {
  name     = "My Access Role"
  comments = "my-Access-Role"
  machines {
    source    = "any"
    selection = ["any"]
  }
}

data "checkpoint_management_data_access_role" "data_access_role" {
  name = "${checkpoint_management_access_role.access_role.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.
* `machines` - Machines that can access the system. machines blocks are documented below.
* `networks` - Collection of Network objects identified by the name or UID that can access the system.
* `remote_access_clients` - Remote access clients identified by name or UID.
* `tags` - Collection of tag identifiers.
* `users` - Users that can access the system. users blocks are documented below.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.

`machines` supports the following:

* `source` - (Optional) any, all identified, Active Directory name or UID or Identity Tag. default = "any"
* `selection` - (Optional) Name or UID of an object selected from source. selection blocks are documented below. default
  = ["any"]
* `base_dn` - (Optional) When source is "Active Directory" use "base-dn" to refine the query in AD database.

`users` supports the following:

* `source` - (Optional) any, all identified, Active Directory name or UID or Identity Tag or Internal User Groups or
  LDAP groups or Guests. default value = "any"
* `selection` - (Optional) Name or UID of an object selected from source. selection blocks are documented below. default
  value = ["any"]
* `base_dn` - (Optional) When source is "Active Directory" use "base-dn" to refine the query in AD database. 
