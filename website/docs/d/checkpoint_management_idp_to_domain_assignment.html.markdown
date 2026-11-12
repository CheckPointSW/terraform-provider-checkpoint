---
layout: "checkpoint"
page_title: "checkpoint_management_idp_to_domain_assignment"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-idp-to-domain-assignment"
description: |-
Use this data source to get information on an existing Check Point Idp To Domain Assignment.
---

# Data Source: checkpoint_management_idp_to_domain_assignment

Use this data source to get information on an existing Check Point Idp To Domain Assignment.

## Example Usage


```hcl
data "checkpoint_management_idp_to_domain_assignment" "example" {
  assigned_domain = "SMS"
}
```

## Argument Reference

The following arguments are supported:
* `uid` - (Optional) Object unique identifier.
* `assigned_domain` - (Optional) Represents the Domain assigned by 'idp-to-domain-assignment', need to be domain name or UID.

## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

