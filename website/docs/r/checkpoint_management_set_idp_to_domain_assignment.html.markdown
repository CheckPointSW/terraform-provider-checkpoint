---
layout: "checkpoint"
page_title: "checkpoint_management_set_idp_to_domain_assignment"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-set-idp-to-domain-assignment"
description: |-
This resource allows you to execute Check Point Set Idp To Domain Assignment.
---

# checkpoint_management_set_idp_to_domain_assignment

This resource allows you to execute Check Point Set Idp To Domain Assignment.

## Example Usage


```hcl
resource "checkpoint_management_set_idp_to_domain_assignment" "example" {
  assigned_domain = "BSMS"
  identity_provider = "okta"
}
```

## Argument Reference

The following arguments are supported:

* `uid` - (Optional) Object unique identifier..
* `assigned_domain` - (Optional) Represents the Domain assigned by 'idp-to-domain-assignment', need to be domain name or UID. 
* `identity_provider` - (Optional) Represents the Identity Provider to be used for Login by this assignment. Must be set when "using-default" was set to be false. 
* `using_default` - (Optional) Is this assignment override by 'idp-default-assignment'. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

