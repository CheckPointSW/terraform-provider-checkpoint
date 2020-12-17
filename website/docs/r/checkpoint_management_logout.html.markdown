---
layout: "checkpoint"
page_title: "checkpoint_management_logout "
sidebar_current: "docs-checkpoint-resource-checkpoint-management-logout"
description: |-
  Log out from the current session. After logging out the session id is not valid any more.
---

# Resource: checkpoint_management_logout

This command resource allows you to log out from the current session. After logging out the session id is not valid any more.

## Example Usage

```hcl
resource "checkpoint_management_logout" "example" {}
```

## Argument Reference

* `triggers` - (Optional) Triggers a logout if there are any changes to objects in this list.

## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.    



