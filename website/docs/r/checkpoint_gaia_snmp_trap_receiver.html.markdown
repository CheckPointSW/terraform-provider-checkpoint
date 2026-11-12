---
layout: "checkpoint"
page_title: "checkpoint_gaia_snmp_trap_receiver"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-snmp-trap-receiver"
description: |-
This resource allows you to execute Check Point Snmp Trap Receiver.
---

# checkpoint_gaia_snmp_trap_receiver

This resource allows you to execute Check Point Snmp Trap Receiver.

## Example Usage


```hcl
resource "checkpoint_gaia_snmp_trap_receiver" "example" {
  address = "4.4.4.5"
  community_string = "yy"
  version = "v3"
}
```

## Argument Reference

The following arguments are supported:

* `address` - (Required) Receiver address 
* `version` - (Required) Receiver version 
* `community_string` - (Optional) Receiver community - Required only in case of v1/v2 versions Trap Community String used by the trap receiver to determine which traps are accepted from a device. 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


## Import

Use the existing resource's `address` as the import ID.

```shell
terraform import checkpoint_gaia_snmp_trap_receiver.example '192.0.2.0'
```

To specify optional read context or filters (`member_id`), use a JSON object with Terraform attribute names and include all required lookup fields.

IDs beginning with `{` must use the JSON-object form; JSON syntax is reserved for structured import IDs.

`member_id` selects the member used for reads; it does not restrict later writes to that member. Gaia set operations retain the scope documented by the resource.

Create a matching resource block before importing, and review `terraform plan` afterward. The importer first validates and records the lookup fields; Terraform then calls the resource reader to fetch existing configuration. Import does not create the resource or recover write-only values such as passwords.
