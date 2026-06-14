---
layout: "checkpoint"
page_title: "checkpoint_gaia_show_snmp_oid"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-show-snmp-oid"
description: |-
This resource allows you to execute Check Point Show Snmp Oid.
---

# checkpoint_gaia_show_snmp_oid

This resource allows you to execute Check Point Show Snmp Oid.

## Example Usage


```hcl
data "checkpoint_gaia_show_snmp_oid" "example" {
  oid = ".1.3.6.1.4.1.2620.1.6.7.3.3.0"
  snmp_sid = "9839921201005492114921201019992120101549212098549212057999212010052921205710239"
}
```

## Argument Reference

The following arguments are supported:

* `oid` - (Required) OID (object identifier) 
* `snmp_sid` - (Required) SNMP session id returned from set-snmp-session 
* `member_id` - (Optional) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

