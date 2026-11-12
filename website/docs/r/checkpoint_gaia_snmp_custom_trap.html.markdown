---
layout: "checkpoint"
page_title: "checkpoint_gaia_snmp_custom_trap"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-snmp-custom-trap"
description: |-
This resource allows you to execute Check Point Snmp Custom Trap.
---

# checkpoint_gaia_snmp_custom_trap

This resource allows you to execute Check Point Snmp Custom Trap.

## Example Usage


```hcl
resource "checkpoint_gaia_snmp_custom_trap" "example" {
  name = "test"
  message = "Test"
  frequency = 6
  oid = "2.16.840.1.113883.3.3190.100"
  operator = "greater-than"
  threshold = "4"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Custom trap name 
* `oid` - (Required) OID (object identifier) 
* `operator` - (Required) Comparison operator 
* `threshold` - (Required) The value you want to compare to 
* `frequency` - (Required) Polling interval in seconds 
* `message` - (Required) Custom trap message 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


## Import

Use the existing resource's `name` as the import ID.

```shell
terraform import checkpoint_gaia_snmp_custom_trap.example 'myname'
```

To specify optional read context or filters (`member_id`), use a JSON object with Terraform attribute names and include all required lookup fields.

```shell
terraform import checkpoint_gaia_snmp_custom_trap.example '{"name":"myname","member_id":"1_1"}'
```

Create a matching resource block before importing, and review `terraform plan` afterward. Import reads existing configuration; it does not create the resource or recover write-only values such as passwords.
