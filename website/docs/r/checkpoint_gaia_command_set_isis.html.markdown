---
layout: "checkpoint"
page_title: "checkpoint_gaia_command_set_isis"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-command-set-isis"
description: |-
This resource allows you to execute Check Point Set Isis.
---

# checkpoint_gaia_command_set_isis

This resource allows you to execute Check Point Set Isis.

## Example Usage


```hcl
resource "checkpoint_gaia_command_set_isis" "example" {
  adjacency_check = true
  area_list = ["45", "58",]
  authentication {
    encrypted_secret = "Bs1/h84ILBk="
    level = "2"
    type = "plaintext"
  }
  authentication {
    active_key = "default"
    level = "1"
    type = "MD5"
  }
  authentication_ignore {
    ignore_all = true
    ignore_csnp = true
    ignore_hello = true
    ignore_lsp = false
    ignore_none = false
    ignore_psnp = false
    level = "1"
  }
  authentication_ignore {
    ignore_all = true
    ignore_csnp = true
    ignore_hello = true
    ignore_lsp = true
    ignore_none = false
    ignore_psnp = false
    level = "2"
  }
  default_metric {
    level = "1-2"
    metric = "12345"
  }
  dynamic_hostname = false
  ignore_attached_bit = true
  is_type = "level-1-2"
  max_areas = "125"
  metric_type {
    level = "1-2"
    type = "wide"
  }
  overload_bit = true
  prc_interval {
    initial_offset = "55"
    level = "1-2"
    max_interval = "119"
    second_offset = "68"
  }
  spf {
    initial_offset = "123"
    level = "1"
    max_interval = "115"
    second_offset = "777"
  }
  spf {
    initial_offset = "987"
    level = "2"
    max_interval = "114"
    second_offset = "323"
  }
  system_id = "aaaa.aaaa.aaaa"
}
```

## Argument Reference

The following arguments are supported:

* `adjacency_check` - (Optional) Enable or disable strict protocol checks with neighbors 
* `area_list` - (Optional) Add or remove an IS-IS area area_list blocks are documented below.
* `default_metric` - (Optional) Set IS-IS default metric default_metric blocks are documented below.
* `is_type` - (Optional) Configure which level this IS-IS router should operate on 
* `max_areas` - (Optional) Set the maximum number of areas 
* `overload_bit` - (Optional) Set whether this IS-IS router should include an overload bit 
* `system_id` - (Optional) Configure IS System ID. Note that this can not be done when IS-IS is configured and running 
* `ignore_attached_bit` - (Optional) Set to ignore attached bits set by level 2 connected routers 
* `dynamic_hostname` - (Optional) Enable or disable dyanmic hostname mapping for system IDs 
* `hello` - (Optional) Configure how IS-IS sends and receives hello messages hello blocks are documented below.
* `metric_type` - (Optional) Configure how IS-IS sends metric messages metric_type blocks are documented below.
* `lsp` - (Optional) Configure IS-IS LSP lsp blocks are documented below.
* `spf` - (Optional) Configure IS-IS SPF spf blocks are documented below.
* `prc_interval` - (Optional) Configure IS-IS PRC Interval prc_interval blocks are documented below.
* `authentication_ignore` - (Optional) Ignore settings for authentication authentication_ignore blocks are documented below.
* `authentication` - (Optional) Configure IS-IS authentication authentication blocks are documented below.
* `ipv6` - (Optional) Configure IS-IS ipv6 options ipv6 blocks are documented below.


`area_list` supports the following:



`default_metric` supports the following:

* `level` - (Optional) Set the default metric for level 1 
* `metric` - (Optional) Set the default metric for level 2 


`hello` supports the following:

* `interface_point_to_point` - (Optional) Set hello type for point to point interfaces 
* `interface_broadcast` - (Optional) Set hello type for broadcast interfacec 


`metric_type` supports the following:

* `type` - (Optional) Set metric type 
* `level` - (Optional) Set level for which this metric type applies 


`lsp` supports the following:

* `lifetime` - (Optional) Set the lifetime of the LSP 
* `mtu` - (Optional) Set the mtu of the LSP 
* `refresh_interval` - (Optional) Set the refresh interval of the LSP 
* `gen_interval` - (Optional) Set the gen interval  gen_interval blocks are documented below.


`spf` supports the following:

* `level` - (Optional) The is type for spf 
* `max_interval` - (Optional) Set the max interval  
* `initial_offset` - (Optional) Set the intial offset 
* `second_offset` - (Optional) Set the second offset 


`prc_interval` supports the following:

* `level` - (Optional) The is type for the prc interval 
* `max_interval` - (Optional) Set the level 1 max interval  
* `initial_offset` - (Optional) Set the level 1 intial offset 
* `second_offset` - (Optional) Set the level 1 second offset 


`authentication_ignore` supports the following:

* `ignore_all` - (Optional) Ignore all setting 
* `ignore_csnp` - (Optional) Ignore csnp setting 
* `ignore_hello` - (Optional) Ignore hello setting 
* `ignore_lsp` - (Optional) Ignore lsp setting 
* `ignore_psnp` - (Optional) Ignore psnp setting 
* `ignore_none` - (Optional) Ignore none setting 
* `level` - (Optional) The is type for the auth interval 


`authentication` supports the following:

* `type` - (Optional) Authentication type 
* `level` - (Optional) The is type for the auth interval 
* `encrypted_secret` - (Optional) Encrypted secret 
* `secret` - (Optional) Not encrypted secret 
* `keys` - (Optional) Authentication key keys blocks are documented below.
* `active_key` - (Optional) Active key 


`ipv6` supports the following:

* `overload_bit` - (Optional) Set whether this IS-IS router should include an overload bit 
* `ignore_attached_bit` - (Optional) Set to ignore attached bits set by level 2 connected routers 
* `multi_topology` - (Optional) Configure IS-IS ipv6 multi topology 
* `prc_interval` - (Optional) Configure IS-IS PRC Interval prc_interval blocks are documented below.
* `spf` - (Optional) Configure IS-IS SPF spf blocks are documented below.
* `default_metric` - (Optional) Set IS-IS default metric default_metric blocks are documented below.


`gen_interval` supports the following:

* `level` - (Optional) Set the level for this gen interval 
* `max_interval` - (Optional) Set the level 1 max gen interval  
* `initial_offset` - (Optional) Set the level 1 intial offset 
* `second_offset` - (Optional) Set the level 1 second offset 


`keys` supports the following:

* `encrypted_secret` - (Optional) Encrypted secret 
* `secret` - (Optional) Not encrypted secret 
* `resource_id` - (Optional) Authentication key 
* `algorithm` - (Optional) Authentication algorithm 


`prc_interval` supports the following:

* `level` - (Optional) The is type for the prc interval 
* `max_interval` - (Optional) Set the level 1 max interval  
* `initial_offset` - (Optional) Set the level 1 intial offset 
* `second_offset` - (Optional) Set the level 1 second offset 


`spf` supports the following:

* `level` - (Optional) The is type for spf 
* `max_interval` - (Optional) Set the max interval  
* `initial_offset` - (Optional) Set the intial offset 
* `second_offset` - (Optional) Set the second offset 


`default_metric` supports the following:

* `level` - (Optional) Set the default metric for level 1 
* `metric` - (Optional) Set the default metric for level 2 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

