---
layout: "checkpoint"
page_title: "checkpoint_management_delete_threat_protections"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-delete-threat-protections"
description: |-
This resource allows you to execute Check Point Delete Threat Protections.
---

# checkpoint_management_delete_threat_protections

This resource allows you to execute Check Point Delete Threat Protections.

## Example Usage


```hcl
resource "checkpoint_management_delete_threat_protections" "example" {
  package_format = "snort"
}
```

## Argument Reference

The following arguments are supported:

* `package_format` - (Optional) Protections package format. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

