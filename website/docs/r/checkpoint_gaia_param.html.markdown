---
layout: "checkpoint"
page_title: "checkpoint_gaia_param"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-param"
description: |-
This resource allows you to execute Check Point Param.
---

# checkpoint_gaia_param

This resource allows you to execute Check Point Param.

## Example Usage


```hcl
resource "checkpoint_gaia_param" "example" {
  param_list {
    param_path  = "firewall.ipv4.acceleration.cphwd_agg_log_delay"
    use_default = true
  }
}
```

## Argument Reference

The following arguments are supported:

* `param_list` - (Required) List of parameters to be set param_list blocks are documented below.
* `param_path` - (Computed) Paramter full path or prefix 
* `virtual_system_id` - (Optional) VSX vs-id which present the context switch 
* `dry_run` - (Optional) run the request without saving to data base, return only with the changed values. 
* `use_regex` - (Optional) indicates that parameter param_path includes * 
* `filter_by` - (Computed) show parameters by filtering type in the following format: <fiter_type>=true when fiter_type can be one of (modified,with-comments,not-default), for example modified=true 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


`param_list` supports the following:

* `param_path` - (Optional) Parameter's full path 
* `value` - (Optional) Parameter's value. it can be sanitized-ascii, int or boolean or object. can not send it with use-default in the same request 
* `comments` - (Optional) Comments to be added to the parameter, length is limited to 256 characters 
* `use_default` - (Optional) Set the parameter back to its default value. can not send it with value in the same request 
* `volatile` - (Optional) Set parameter with the value specified untill the next reboot 
* `virtual_system_id` - (Optional) VSX vs-id which present the context id, can be 'all' or a string represent spesific VS Ids, for example: '2,4,7-10' 
