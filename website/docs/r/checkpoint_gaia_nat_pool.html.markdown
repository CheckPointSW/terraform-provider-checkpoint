---
layout: "checkpoint"
page_title: "checkpoint_gaia_nat_pool"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-nat-pool"
description: |-
This resource allows you to execute Check Point Nat Pool.
---

# checkpoint_gaia_nat_pool

This resource allows you to execute Check Point Nat Pool.

## Example Usage


```hcl
resource "checkpoint_gaia_nat_pool" "example" {
  prefix = "10.10.10.0/24"
  comment = "Test Comment"
}
```

## Argument Reference

The following arguments are supported:

* `prefix` - (Required) Specifies the IPv4 or IPv6 destination prefix of a NAT pool to be configured.  Note: A prefix cannot be of type IPv6, if IPv6 capabilities are not enabled 
* `comment` - (Optional) Specifies a comment on a NAT pool. If the empty string is given, no comments will be added to the NAT pool.  Note: The length of the comment cannot exceed 100 characters 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


## Import

Use the existing resource's `prefix` as the import ID.

```shell
terraform import checkpoint_gaia_nat_pool.example '192.0.2.0/24'
```

To specify optional read context or filters (`member_id`), use a JSON object with Terraform attribute names and include all required lookup fields.

IDs beginning with `{` must use the JSON-object form; JSON syntax is reserved for structured import IDs.

`member_id` selects the member used for reads; it does not restrict later writes to that member. Gaia set operations retain the scope documented by the resource.

Create a matching resource block before importing, and review `terraform plan` afterward. The importer first validates and records the lookup fields; Terraform then calls the resource reader to fetch existing configuration. Import does not create the resource or recover write-only values such as passwords.
