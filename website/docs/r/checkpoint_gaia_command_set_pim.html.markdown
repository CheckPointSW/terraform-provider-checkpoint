---
layout: "checkpoint"
page_title: "checkpoint_gaia_command_set_pim"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-command-set-pim"
description: |-
This resource allows you to execute Check Point Set Pim.
---

# checkpoint_gaia_command_set_pim

This resource allows you to execute Check Point Set Pim.

## Example Usage


```hcl
resource "checkpoint_gaia_command_set_pim" "example" {
  mode              = "sparse"
  data_interval     = "23"
  assert_interval   = "24"
  hello_interval    = "26"
  jp_delay_interval = "27"
  jp_interval       = "28"
  spt_threshold {
    multicast_group = "224.88.88.88/32"
    threshold       = "infinity"
  }
  spt_threshold {
    multicast_group = "224.99.99.99/32"
    threshold       = "infinity"
  }
}
```

## Argument Reference

The following arguments are supported:

* `assert_interval` - (Optional) Specifies the number of seconds that assert state should be maintained in the absence of a refreshing assert message. 
* `assert_rank` - (Optional) Assert rank defines the cost of a routing protocolrelative to other protocols. assert_rank blocks are documented below.
* `bootstrap_candidate` - (Optional) Configures candidate Bootstrap Router (candidate BSR) options. bootstrap_candidate blocks are documented below.
* `candidate_rp` - (Optional) Configures candidate Rendezvous Point (candidate RP) options. candidate_rp blocks are documented below.
* `data_interval` - (Optional) Configure the Data Interval. 
* `hello_interval` - (Optional) Configure the Hello Interval. 
* `jp_delay_interval` - (Optional) Configure the Random Delay Join/Prune Interval 
* `jp_interval` - (Optional) Configure the Join/Prune Interval (PIM-SM/SSM only). 
* `mode` - (Optional) Configure Dense Mode/Sparse Mode/SSM Mode 
* `register_suppress_interval` - (Optional) Configure Register-Suppression Interval 
* `enable_state_refresh` - (Optional) Configure State Refresh 
* `state_refresh_interval` - (Optional) Configure State Refresh Interval 
* `state_refresh_ttl` - (Optional) Configure State Refresh TTL 
* `static_rp` - (Optional) Configure Static Rendezvous Point static_rp blocks are documented below.
* `spt_threshold` - (Optional) Configure SPT Threshold spt_threshold blocks are documented below.
* `custom_ssm_prefix` - (Optional) Configure Custom SSM Prefix custom_ssm_prefix blocks are documented below.


`assert_rank` supports the following:

* `protocol` - (Optional) Configure the assert rank of a protocol. 
* `rank` - (Optional) The cost metric. 


`bootstrap_candidate` supports the following:

* `local_address` - (Optional) Configures the local address to use for this candidate Bootstrap Router (candidate BSR). 
* `priority` - (Optional) Configures the candidate Bootstrap Router (candidate BSR) priority. 
* `enable` - (Optional) Configures this router to be a candidate Bootstrap Router (candidate BSR). 


`candidate_rp` supports the following:

* `advertise_interval` - (Optional) Configure the Advertisement Interval 
* `local_address` - (Optional) Configure the Candidate RP Local Address. 
* `enable` - (Optional) Configure a Candidate Rendezvous Point 
* `priority` - (Optional) Configure Candidate Rendezvous Point Priority. 
* `multicast_group` - (Optional) Configure a Candidate RP Multicast Group. multicast_group blocks are documented below.


`static_rp` supports the following:

* `rp_address` - (Optional) Adds the given static Rendezvous Point (static RP). 
* `enable` - (Optional) Configure a static Rendezvous Point (static RP). 
* `multicast_group` - (Optional) Configure a Static RP Multicast Group. multicast_group blocks are documented below.


`spt_threshold` supports the following:

* `multicast_group` - (Optional) The multicast group prefix/mask, in CIDR notation. 
* `threshold` - (Optional) The threshold for the multicast group. 


`custom_ssm_prefix` supports the following:

* `address` - (Optional) The multicast group prefix/mask, in CIDR notation. 


`multicast_group` supports the following:

* `address` - (Optional) The multicast group prefix/mask, in CIDR notation. 


`multicast_group` supports the following:

* `address` - (Optional) The multicast group prefix/mask, in CIDR notation. 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

