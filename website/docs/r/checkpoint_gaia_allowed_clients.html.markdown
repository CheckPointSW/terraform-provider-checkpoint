---
layout: "checkpoint"
page_title: "checkpoint_gaia_allowed_clients"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-allowed-clients"
description: |-
This resource allows you to execute Check Point Allowed Clients.
---

# checkpoint_gaia_allowed_clients

This resource allows you to execute Check Point Allowed Clients.

## Example Usage


```hcl
resource "checkpoint_gaia_allowed_clients" "example" {
  allowed_any_host = true
}
```

## Argument Reference

The following arguments are supported:

* `allowed_networks` - (Optional)  allowed_networks blocks are documented below.
* `allowed_hosts` - (Optional)  allowed_hosts blocks are documented below.
* `allowed_any_host` - (Optional)  
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


`allowed_networks` supports the following:

* `subnet` - (Optional) The network subnet 
* `mask_length` - (Optional) The network mask length 


`allowed_hosts` supports the following:



## Import

Use `settings` to import the existing configuration on the connected Gaia machine.

```shell
terraform import checkpoint_gaia_allowed_clients.example 'settings'
```

To specify optional read context or filters (`member_id`), use a JSON object with Terraform attribute names and include all required lookup fields.

`member_id` selects the member used for reads; it does not restrict later writes to that member. Gaia set operations retain the scope documented by the resource.

Create a matching resource block before importing, and review `terraform plan` afterward. The importer first validates and records the lookup fields; Terraform then calls the resource reader to fetch existing configuration. Import does not create the resource or recover write-only values such as passwords.
