---
layout: "checkpoint"
page_title: "checkpoint_gaia_cluster_member"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-cluster-member"
description: |-
This resource allows you to execute Check Point Cluster Member.
---

# checkpoint_gaia_cluster_member

This resource allows you to execute Check Point Cluster Member.

## Example Usage


```hcl
resource "checkpoint_gaia_cluster_member" "example" {
  site_id = 1
  method = "serial-number"
  identifier = "LR201909015218"
}
```

## Argument Reference

The following arguments are supported:

* `method` - (Required) Method used for adding the member:         1. serial-number - Retrieve from new member using "show-serial-number"         2. hostname - Retrieve from new member using "show-hostname"         3. request-id - Retrieve from new member using "show-cluster-request-id" 
* `identifier` - (Required) Identifier of member 
* `site_id` - (Required) Site id to add member to 
* `member` - (Computed) Computed field, returned in the response. member blocks are documented below.


`member` supports the following:

* `hostname` - (Computed) Computed field, returned in the response. 
* `serial_number` - (Computed) Computed field, returned in the response. 
* `request_id` - (Computed) Computed field, returned in the response. 
* `site_id` - (Computed) Computed field, returned in the response. 
* `member_id` - (Computed) Computed field, returned in the response. 
* `model` - (Computed) Computed field, returned in the response. 
* `version` - (Computed) Computed field, returned in the response. 
* `member_status` - (Computed) Computed field, returned in the response. 
* `site_status` - (Computed) Computed field, returned in the response. 
* `state` - (Computed) Computed field, returned in the response. 
* `installed_jumbo_take` - (Computed) Computed field, returned in the response. 
