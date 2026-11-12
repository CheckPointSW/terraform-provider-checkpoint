---
layout: "checkpoint"
page_title: "checkpoint_management_login_message"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-login-message"
description: |-
Use this data source to get information on an existing Check Point Login Message.
---

# Data Source: checkpoint_management_login_message

Use this data source to get information on an existing Check Point Login Message.

## Example Usage


```hcl
data "checkpoint_management_login_message" "data_login_message" {

}
```

## Argument Reference

The following arguments are supported:

* `header` - Login message header.
* `message` - Login message body.
* `show_message` - Whether to show login message.
* `warning` - Add warning sign.