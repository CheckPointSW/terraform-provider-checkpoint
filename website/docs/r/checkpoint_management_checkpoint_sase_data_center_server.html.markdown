---
layout: "checkpoint"
page_title: "checkpoint_management_checkpoint_sase_data_center_server"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-checkpoint-sase-data-center-server"
description: |- This resource allows you to execute Check Point SASE data center server.
---

# Resource: checkpoint_management_checkpoint_sase_data_center_server

This resource allows you to execute Check Point SASE Data Center Server.

## Example Usage

```hcl
resource "checkpoint_management_checkpoint_sase_data_center_server" "testCheckpointSase" {
  name = "MY-CHECKPOINT-SASE"
  connect_to = "connected-portal"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name.
* `connect_to` - (Required) connected-portal: Connect to the connected Check Point Portal Account. other-portal: Connect to a different Check Point Portal Account.
* `hostname` - (Optional) URL from Check Point Portal. Required for connect-to: other-portal.
* `client_id` - (Optional) Client ID for Check Point SASE account. Required for connect-to: other-portal.
* `secret_key` - (Optional) Secret key for Check Point SASE account. Required for connect-to: other-portal.
* `tags` - (Optional) Collection of tag identifiers.
* `color` - (Optional) Color of the object. Should be one of existing colors.
* `comments` - (Optional) Comments string.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. By Setting this parameter to 'true' test connection failure will be ignored.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.
