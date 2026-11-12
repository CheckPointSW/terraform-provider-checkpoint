---
layout: "checkpoint"
page_title: "checkpoint_management_checkpoint_sase_data_center_server"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-checkpoint-sase-data-center-server"
description: |- Use this data source to get information on an existing Check Point SASE data center server.
---

# Data Source: checkpoint_management_checkpoint_sase_data_center_server

Use this data source to get information on an existing Check Point SASE Data Center Server.

## Example Usage

```hcl
resource "checkpoint_management_checkpoint_sase_data_center_server" "testCheckpointSase" {
  name = "MY-CHECKPOINT-SASE"
  connect_to = "connected-portal"
}

data "checkpoint_management_checkpoint_sase_data_center_server" "data_checkpoint_sase_data_center_server" {
  name = "${checkpoint_management_checkpoint_sase_data_center_server.testCheckpointSase.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.
* `connect_to` - (Computed) connected-portal: Connect to the connected Check Point Portal Account. other-portal: Connect to a different Check Point Portal Account.
* `hostname` - (Computed) URL from Check Point Portal. Required for connect-to: other-portal.
* `client_id` - (Computed) Client ID for Check Point SASE account. Required for connect-to: other-portal.
* `tags` - (Computed) Collection of tag objects identified by the name or UID.
* `color` - (Computed) Color of the object.
* `comments` - (Computed) Comments string.
