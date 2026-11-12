---
layout: "checkpoint"
page_title: "checkpoint_gaia_static_mroute"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-static-mroute"
description: |-
This resource allows you to execute Check Point Static Mroute.
---

# checkpoint_gaia_static_mroute

This resource allows you to execute Check Point Static Mroute.

## Example Usage


```hcl
resource "checkpoint_gaia_static_mroute" "example" {
  address     = "40.40.40.0"
  mask_length = 24

  next_hop {
    gateway  = "172.23.22.1"
    priority = "default"
  }
}
```

## Argument Reference

The following arguments are supported:

* `address` - (Required) Address of the static-mroute to set configuration for. 
* `mask_length` - (Required) Mask length for the static-mroute. 
* `next_hop` - (Required) Static next-hop. Contains a list of next-hop gateways.  Each gateway is formatted in the following manner: {"gateway": IP address, "priority": default or integer 1-8} next_hop blocks are documented below.
* `ping` - (Optional) Configures ping monitoring of the given IPv4 static-mroute. Possible values: true, false  NOTE: Static-mroute ping is not supported in versions prior to R82.10 
* `virtual_system_id` - (Optional) Virtual System ID. Relevant for VSNext setups 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


`next_hop` supports the following:

* `gateway` - (Optional) IP address for the static next-hop gateway. 
* `priority` - (Optional) Priority defines which gateway to select as the next-hop: the lower the priority, the higher the preference. Possible values: default or integer 1-8 


## Import

Use a JSON object containing these lookup fields: `address`, `mask_length`.

```shell
terraform import checkpoint_gaia_static_mroute.example '{"address":"192.0.2.0","mask_length":24}'
```

To specify optional read context or filters (`virtual_system_id`, `member_id`), use a JSON object with Terraform attribute names and include all required lookup fields.

`member_id` selects the member used for reads; it does not restrict later writes to that member. Gaia set operations retain the scope documented by the resource.

Create a matching resource block before importing, and review `terraform plan` afterward. The importer first validates and records the lookup fields; Terraform then calls the resource reader to fetch existing configuration. Import does not create the resource or recover write-only values such as passwords.
