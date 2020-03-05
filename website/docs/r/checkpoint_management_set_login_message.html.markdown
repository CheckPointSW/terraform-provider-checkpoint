---
layout: "checkpoint"
page_title: "checkpoint_management_set_login_message"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-set-login-message"
description: |-
This resource allows you to execute Check Point Set Login Message.
---

# checkpoint_management_set_login_message

This resource allows you to execute Check Point Set Login Message.

## Example Usage


```hcl
resource "checkpoint_management_set_login_message" "example" {
  show_message = true
  header = "Warning"
  message = "Unauthorized access of this server is prohibited and punished by law"
  warning = true
}
```

## Argument Reference

The following arguments are supported:

* `header` - (Optional) Login message header. 
* `message` - (Optional) Login message body. 
* `show_message` - (Optional) Whether to show login message. 
* `warning` - (Optional) Add warning sign. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

