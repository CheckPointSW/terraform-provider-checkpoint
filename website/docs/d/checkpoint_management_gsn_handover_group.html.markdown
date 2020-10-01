---
layout: "checkpoint"
page_title: "checkpoint_management_gsn_handover_group"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-gsn-handover-group"
description: |-
This resource allows you to execute Check Point Gsn Handover Group.
---

# Data Source: checkpoint_management_gsn_handover_group

This resource allows you to execute Check Point Gsn Handover Group.

## Example Usage


```hcl
resource "checkpoint_management_gsn_handover_group" "test" {
    name = "gsn group"
    enforce_gtp = true
    gtp_rate = 2048
    members = ["All_Internet"]
}

data "checkpoint_management_gsn_handover_group" "data_test" {
    name = "${checkpoint_management_gsn_handover_group.test.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier. 
* `enforce_gtp` - Enable enforce GTP signal packet rate limit from this group. 
* `gtp_rate` - Limit of the GTP rate in PDU/sec. 
* `members` - Collection of GSN handover group members identified by the name or UID.members blocks are documented below.
* `tags` - Collection of tag identifiers.
* `color` - Color of the object.
* `comments` - Comments string.