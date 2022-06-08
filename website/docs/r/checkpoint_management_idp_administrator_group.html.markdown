---
layout: "checkpoint"
page_title: "checkpoint_management_idp_administrator_group"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-idp-administrator-group"
description: |-
This resource allows you to execute Check Point Idp Administrator Group.
---

# checkpoint_management_idp_administrator_group

This resource allows you to execute Check Point Idp Administrator Group.

## Example Usage


```hcl
resource "checkpoint_management_idp_administrator_group" "example" {
  name = "my super group"
  group_id = "it-team"
  multi_domain_profile = "domain super user"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `group_id` - (Required) Group ID or Name should be set base on the source attribute of 'groups' in the Saml Assertion. 
* `multi_domain_profile` - (Optional) Administrator multi-domain profile. 
* `permissions_profile` - (Optional) Administrator permissions profile. Permissions profile should not be provided when multi-domain-profile is set to "Multi-Domain Super User" or "Domain Super User".permissions_profile blocks are documented below.
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


`permissions_profile` supports the following:

* `domain` - (Optional) N/A 
* `profile` - (Optional) Permission profile. 
