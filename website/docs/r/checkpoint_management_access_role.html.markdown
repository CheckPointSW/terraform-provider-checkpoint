---
layout: "checkpoint"
page_title: "checkpoint_management_access_role"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-access-role"
description: |-
This resource allows you to execute Check Point Access Role.
---

# checkpoint_management_access_role

This resource allows you to execute Check Point Access Role.

## Example Usage


```hcl
resource "checkpoint_management_access_role" "example" {
  name = "New Access Role 1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `machines` - (Optional) Machines that can access the system.machines blocks are documented below.
* `networks` - (Optional) Collection of Network objects identified by the name or UID that can access the system.networks blocks are documented below.
* `remote_access_clients` - (Optional) Remote access clients identified by name or UID. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `users` - (Optional) Users that can access the system.users blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


`machines` supports the following:

* `source` - (Optional) Active Directory name or UID or Identity Tag. 
* `selection` - (Optional) Name or UID of an object selected from source.selection blocks are documented below.
* `base_dn` - (Optional) When source is "Active Directory" use "base-dn" to refine the query in AD database. 


`users` supports the following:

* `source` - (Optional) Active Directory name or UID or Identity Tag  or Internal User Groups or LDAP groups or Guests. 
* `selection` - (Optional) Name or UID of an object selected from source.selection blocks are documented below.
* `base_dn` - (Optional) When source is "Active Directory" use "base-dn" to refine the query in AD database. 
