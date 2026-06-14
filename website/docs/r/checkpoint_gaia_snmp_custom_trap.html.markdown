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
