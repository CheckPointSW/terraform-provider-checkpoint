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
