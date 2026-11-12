---
layout: "checkpoint"
page_title: "checkpoint_gaia_mld_interface_static_group"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-mld-interface-static-group"
description: |-
This resource allows you to execute Check Point Mld Interface Static Group.
---

# checkpoint_gaia_mld_interface_static_group

This resource allows you to execute Check Point Mld Interface Static Group.

## Example Usage


```hcl
resource "checkpoint_gaia_mld_interface_static_group" "example" {
  interface    = "eth0"
  static_group = "ff02::feed"
  sources {
    source           = "11::11"
    source_count     = "3"
    source_increment = "::2"
  }
}
```

## Argument Reference

The following arguments are supported:

* `interface` - (Required) The name of the MLD interface 
* `static_group` - (Required) The statically configured group address that this MLD interface receives multicast data for 
* `group_count` - (Optional) The number of adjacent static groups 
* `group_increment` - (Optional) The increment between MLD static groups (default: ::1) 
* `sources` - (Optional) The list of IPv6 sources from which to receive traffic for this static group sources blocks are documented below.
* `source_all_off` - (Optional) Remove all sources of a static group 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


`sources` supports the following:

* `source` - (Optional) The IPv6 source from which to receive traffic for this static group 
* `source_count` - (Optional) The number of adjacent static group sources 
* `source_increment` - (Optional) The increment between MLD static group sources (default: ::1) 


## Import

Use a JSON object containing these lookup fields: `interface`, `static_group`.

```shell
terraform import checkpoint_gaia_mld_interface_static_group.example '{"interface":"eth1","static_group":"239.1.1.1"}'
```

To specify optional read context or filters (`member_id`), use a JSON object with Terraform attribute names and include all required lookup fields.

Create a matching resource block before importing, and review `terraform plan` afterward. Import reads existing configuration; it does not create the resource or recover write-only values such as passwords.
