---
layout: "checkpoint"
page_title: "checkpoint_management_gsn_handover_group"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-gsn-handover-group"
description: |-
This resource allows you to execute Check Point Gsn Handover Group.
---

# Resource: checkpoint_management_gsn_handover_group

This resource allows you to execute Check Point Gsn Handover Group.

## Example Usage


```hcl
resource "checkpoint_management_gsn_handover_group" "example" {
  name = "gsn group"
  enforce_gtp = true
  gtp_rate = 2048
  members = ["All_Internet"]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `enforce_gtp` - (Optional) Enable enforce GTP signal packet rate limit from this group. 
* `gtp_rate` - (Optional) Limit of the GTP rate in PDU/sec. 
* `members` - (Optional) Collection of GSN handover group members identified by the name or UID.members blocks are documented below.
* `tags` - (Optional) Collection of tag identifiers.
* `color` - (Optional) Color of the object.
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.