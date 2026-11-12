---
layout: "checkpoint"
page_title: "checkpoint_gaia_snmp_pre_defined_traps"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-snmp-pre-defined-traps"
description: |-
This resource allows you to execute Check Point Snmp Pre Defined Traps.
---

# checkpoint_gaia_snmp_pre_defined_traps

This resource allows you to execute Check Point Snmp Pre Defined Traps.

## Example Usage


```hcl
resource "checkpoint_gaia_snmp_pre_defined_traps" "example" {
}
```

## Argument Reference

The following arguments are supported:

* `authorizationerror` - (Optional) authorizationError Trap authorizationerror blocks are documented below.
* `biosfailure` - (Optional) biosFailure Trap biosfailure blocks are documented below.
* `configurationchange` - (Optional) configurationChange Trap configurationchange blocks are documented below.
* `configurationsave` - (Optional) configurationSave Trap configurationsave blocks are documented below.
* `fanfailure` - (Optional) fanFailure Trap fanfailure blocks are documented below.
* `highvoltage` - (Optional) highVoltage Trap highvoltage blocks are documented below.
* `linkuplinkdown` - (Optional) linkUpLinkDown Trap linkuplinkdown blocks are documented below.
* `clusterxlfailover` - (Optional) clusterXLFailover Trap clusterxlfailover blocks are documented below.
* `lowvoltage` - (Optional) lowVoltage Trap lowvoltage blocks are documented below.
* `overtemperature` - (Optional) overTemperature Trap overtemperature blocks are documented below.
* `powersupplyfailure` - (Optional) powerSupplyFailure Trap powersupplyfailure blocks are documented below.
* `raidvolumestate` - (Optional) raidVolumeState Trap raidvolumestate blocks are documented below.
* `vrrpv2authfailure` - (Optional) vrrpv2AuthFailure Trap vrrpv2authfailure blocks are documented below.
* `vrrpv2newmaster` - (Optional) vrrpv2NewMaster Trap vrrpv2newmaster blocks are documented below.
* `vrrpv3newmaster` - (Optional) vrrpv3NewMaster Trap vrrpv3newmaster blocks are documented below.
* `vrrpv3protoerror` - (Optional) vrrpv3ProtoError Trap vrrpv3protoerror blocks are documented below.
* `coldstart` - (Optional) ColdStart Trap coldstart blocks are documented below.
* `lowdiskspaceallpartitions` - (Optional) lowDiskSpaceAllPartitions Trap lowdiskspaceallpartitions blocks are documented below.
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


`authorizationError` supports the following:

* `enabled` - (Optional) Pre-defined trap state 


`biosFailure` supports the following:

* `enabled` - (Optional) Pre-defined trap state 


`configurationChange` supports the following:

* `enabled` - (Optional) Pre-defined trap state 


`configurationSave` supports the following:

* `enabled` - (Optional) Pre-defined trap state 


`fanFailure` supports the following:

* `enabled` - (Optional) Pre-defined trap state 


`highVoltage` supports the following:

* `enabled` - (Optional) Pre-defined trap state 


`linkUpLinkDown` supports the following:

* `enabled` - (Optional) Pre-defined trap state 


`clusterXLFailover` supports the following:

* `enabled` - (Optional) Pre-defined trap state 


`lowVoltage` supports the following:

* `enabled` - (Optional) Pre-defined trap state 


`overTemperature` supports the following:

* `enabled` - (Optional) Pre-defined trap state 


`powerSupplyFailure` supports the following:

* `enabled` - (Optional) Pre-defined trap state 


`raidVolumeState` supports the following:

* `enabled` - (Optional) Pre-defined trap state 


`vrrpv2AuthFailure` supports the following:

* `enabled` - (Optional) Pre-defined trap state 


`vrrpv2NewMaster` supports the following:

* `enabled` - (Optional) Pre-defined trap state 


`vrrpv3NewMaster` supports the following:

* `enabled` - (Optional) Pre-defined trap state 


`vrrpv3ProtoError` supports the following:

* `enabled` - (Optional) Pre-defined trap state 


`coldStart` supports the following:

* `enabled` - (Optional) Pre-defined trap state 
* `threshold` - (Optional) coldStart threshold (seconds), prevents sending coldStart trap when system up-time is greater than the threshold 
* `reboot_only` - (Optional) ColdStart reboot only, allows sending ColdStart trap only on reboot 


`lowDiskSpaceAllPartitions` supports the following:

* `enabled` - (Optional) Pre-defined trap state 


## Import

Use `settings` to import the existing configuration on the connected Gaia machine.

```shell
terraform import checkpoint_gaia_snmp_pre_defined_traps.example 'settings'
```

To specify optional read context or filters (`member_id`), use a JSON object with Terraform attribute names and include all required lookup fields.

`member_id` selects the member used for reads; it does not restrict later writes to that member. Gaia set operations retain the scope documented by the resource.

Create a matching resource block before importing, and review `terraform plan` afterward. The importer first validates and records the lookup fields; Terraform then calls the resource reader to fetch existing configuration. Import does not create the resource or recover write-only values such as passwords.
