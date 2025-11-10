---
layout: "checkpoint"
page_title: "checkpoint_management_renew_scaled_sharing_server_certificate"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-renew-scaled-sharing-server-certificate"
description: |-
 This resource allows you to execute Check Point Renew Scaled Sharing Server Certificate.
---

# checkpoint_management_renew_scaled_sharing_server_certificate

This resource allows you to execute Check Point Renew Scaled Sharing Server Certificate.

## Example Usage


```hcl
resource "checkpoint_management_renew_scaled_sharing_server_certificate" "example" {
  name = "gw1"
}
```

## Argument Reference

The following arguments are supported:

* `uid` - (Optional) Gateway or cluster unique identifier.
* `name` - (Optional) Gateway or cluster name. 
* `message` - Operation status.


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

