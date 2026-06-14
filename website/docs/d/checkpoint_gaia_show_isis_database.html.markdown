---
layout: "checkpoint"
page_title: "checkpoint_gaia_show_isis_database"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-show-isis-database"
description: |-
This resource allows you to execute Check Point Show Isis Database.
---

# checkpoint_gaia_show_isis_database

This resource allows you to execute Check Point Show Isis Database.

## Example Usage


```hcl
resource "checkpoint_gaia_command_set_isis" "isis_setup" {
  system_id = "0101.0101.0101"
}

data "checkpoint_gaia_show_isis_database" "example" {
  protocol_instance = "default"

  depends_on = [checkpoint_gaia_command_set_isis.isis_setup]
}
```

## Argument Reference

The following arguments are supported:

* `protocol_instance` - (Optional) The instance to be queried 
* `level` - (Optional) Filter the results by IS-IS level 
* `lsp_type` - (Optional) Filter the results by lsp-type 
* `system_id` - (Optional) Filter the results by system-id 
* `limit` - (Optional) The maximum number of returned results 
* `offset` - (Optional) The number of results to initially skip 
* `order` - (Optional) Sorts the database entries by their level first, then their LSP IDs in either ascending or descending order. 
* `member_id` - (Optional) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

